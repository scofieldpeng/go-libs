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
func RandomString(lenght ...int) string {
	if len(lenght) == 0 {
		lenght = []int{8}
	}

	rand.Seed(time.Now().Unix())
	res := Md5([]byte(strconv.Itoa(rand.Int())))

	return res[0:lenght[0]]
}

// AppDir 应用目录绝对路径
func AppDir() string{
	return filepath.Dir(os.Args[0])
}