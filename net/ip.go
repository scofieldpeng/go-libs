package net

import(
    "strconv"
    "strings"
    "errors"
    "fmt"
)

// IPv4ToUint 将IPv4地址转化为uint64数字
func IPv4ToUint(ip string) (uint64,error) {
	ipSlice := strings.Split(ip,".")
	if len(ipSlice) != 4 {
		return 0,errors.New("invalid ip")
	}

	ipUintSlice := make([]uint64,len(ipSlice))
	var err error
	for i,ipSplit := range ipSlice {
		ipUintSlice[i],err = strconv.ParseUint(ipSplit,10,64)
		if err != nil {
			return 0,errors.New("invalid ip")
		}
	}

	return ipUintSlice[0] << 24 + ipUintSlice[1] << 16 + ipUintSlice[2] << 8 + ipUintSlice[3],nil
}

// UintToIPv4 将uint64类型的ip数字类型转化为字符串
func UintToIPv4(ipNum uint64) string {
	return fmt.Sprintf("%d.%d.%d.%d",ipNum & 0xFF000000 >> 24, ipNum & 0x00FF0000 >> 16, ipNum & 0x0000FF00 >> 8, ipNum & 0x000000FF)
}
