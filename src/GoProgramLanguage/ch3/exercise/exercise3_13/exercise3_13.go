package main

import "fmt"

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

const (
	Byte = 1
	KB   = Byte * 1000
	MB   = KB * 1000
	GB   = MB * 1000
	TB   = GB * 1000
	PB   = TB * 1000
	EB   = PB * 1000
	//ZB   = EB * 1000
	//YB   = ZB * 1000
)

func main() {
	fmt.Println(YiB / ZiB)
	fmt.Printf("Byte: %d\n", Byte)
	fmt.Printf("KB: %d\n", KB)
	fmt.Printf("MB: %d\n", MB)
	fmt.Printf("GB: %d\n", GB)
	fmt.Printf("TB: %d\n", TB)
	fmt.Printf("PB: %d\n", PB)
	fmt.Printf("EB: %d\n", EB)
	//fmt.Printf("ZB: %s\n", ZB)
	//fmt.Printf("YB: %s\n", YB)

}
