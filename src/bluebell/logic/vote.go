package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"
)

// 投票功能：
// 1、用户投票的数据
// 2、状态

// 使用简化版的投票分数算法
// 投一票就加432分 86400/200 -> 需要200张赞成票可以给帖子续一天 -> 《redis实战》

/* 投票的几种情况：
direction = 1，两种情况：
	1、之前没有投过票，现在投赞成票 --> 更新分数和投票记录
	2、之前投反对票，现在改投赞成票 --> 更新分数和投票记录
direction = 0，两种情况：
	1、之前投赞成票，现在要取消投票 --> 更新分数和投票记录
	2、之前投反对票，现在要取消投票 --> 更新分数和投票记录
direction = -1，两种情况：
	1、之前没有投过票，现在投反对票 --> 更新分数和投票记录
	2、之前投赞成票，现在改投反对票 --> 更新分数和投票记录

投票的限制：
每个帖子自发表之日起，一个星期之内允许用户投票，超过就不允许再投票了
	1、到期之后将redis中保存的赞成票数和反对票数保存到mysql表中
	2、到期之后删除 KeyPostVoteZSetPf
*/

// VoteForPost 为帖子投票的函数
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteFotPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction),
	)

	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
