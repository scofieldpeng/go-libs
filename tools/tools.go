package tools

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"math/rand"
	"strconv"
	"time"
	"path/filepath"
	"os"
	"crypto/sha1"
)

// Crc32 将指定的数据加密和处理为crc32
func Crc32(data []byte) (uint32, error) {
	hash := crc32.NewIEEE()
	_, err := hash.Write(data)
	if err == nil {
		return hash.Sum32(), nil
	}
	return 0, err
}

// Md5 对[]byte字符进行md5操作,返回md5的字符串(16进制)
func Md5(ori []byte) string {
	return fmt.Sprintf("%x", md5.Sum(ori))
}

// RandomString 返回一个随机字符串,默认8位长度,如果需要指定位数,传入要指定的位数,最长不超过32位
func RandomString(length ...int) string {
	if len(length) == 0 {
		length = make([]int,1)
		length[0] = 32
	}
	if length[0] > 32 || length[0] < 1{
		length[0] = 32
	}
	rand.Seed(time.Now().UnixNano())
	randStr := fmt.Sprintf("%x",sha1.Sum([]byte(strconv.Itoa(rand.Int()))))
	return randStr[0:length[0]]
}

// AppDir 应用目录绝对路径
func AppDir() string{
	return filepath.Dir(os.Args[0])
}