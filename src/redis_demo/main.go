package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

// 普通连接模式
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default db
		PoolSize: 100, //连接池大小
	})
	_, err = rdb.Ping().Result()
	return err
}

// 连接Redis哨兵模式
func initClient1() (err error) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// 连接Redis集群
func initClient2() (err error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	// val := rdb.Get("score").Value()
	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score: ", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
	//rdb.HGetAll()
	//rdb.HGet()
	//rdb.HMGet()
}

func hgetDemo() {
	v, err := rdb.HGetAll("user").Result()
	if err != nil {
		// redis.Nil
		fmt.Printf("hgetall failed, err:%v\n", err)
	}
	fmt.Println(v)

	v2 := rdb.HMGet("user", "name", "age").Val()
	fmt.Println(v2)

	v3 := rdb.HGet("user", "age").Val()
	fmt.Println(v3)
}

func redisExample2() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{
			Score:  90.0,
			Member: "Golang",
		},
		redis.Z{
			Score:  98.0,
			Member: "Java",
		},
		redis.Z{
			Score:  95.0,
			Member: "python",
		},
		redis.Z{
			Score:  97.0,
			Member: "JavaScript",
		},
		redis.Z{
			Score:  99.0,
			Member: "C/C++",
		},
	}
	// ZADD
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v", err)
		return
	}
	fmt.Printf("zadd %d success.\n", num)

	//Golang's score add 10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("ZIncrBy failed, err:%v", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err := rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("ZRangeByScore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

// 有多条命令要执行时，就可以考虑试用pipeline优化
func pipeline() {
	pipe := rdb.Pipeline()

	incr := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}

func pipelined() {
	var incr *redis.IntCmd
	_, err := rdb.Pipelined(func(pipe redis.Pipeliner) error {
		incr = pipe.Incr("pipelined_counter")
		pipe.Expire("pipelined_counter", time.Hour)
		return nil
	})
	fmt.Println(incr.Val(), err)
}

func watchDemo() {
	key := "watch_count"
	err := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipeliner redis.Pipeliner) error {
			time.Sleep(time.Second * 5)
			pipeliner.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)
	if err != nil {
		fmt.Printf("tx exec failed, err:%v\n", err)
	}
	fmt.Println("tx exec success")
}

func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis client failed, err:%v", err)
	} else {
		fmt.Println("connect success")
	}
	// 程序退出时释放相关资源
	defer rdb.Close()
	//redisExample()
	//hgetDemo()
	//redisExample2()
	watchDemo()
}
