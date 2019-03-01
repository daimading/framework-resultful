package utils

import (
	"math/rand"
	"time"
)

const (
	RandKindNum   = 1 << 0 // 纯数字
	RandKindLower = 1 << 1 // 小写字母
	RandKindUpper = 1 << 2 // 大写字母
	RandKindAll   = RandKindNum | RandKindLower | RandKindUpper
)

var randSource = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

// 随机字符串
func RandStringK(size int, kind int) []byte {
	scope := make([][]int, 0, 3)
	if kind&RandKindNum != 0 {
		scope = append(scope, []int{10, 48})
	}
	if kind&RandKindLower != 0 {
		scope = append(scope, []int{26, 97})
	}
	if kind&RandKindUpper != 0 {
		scope = append(scope, []int{26, 65})
	}

	result := make([]byte, size)
	l := len(scope)
	for i := 0; i < size; i++ {
		index := randSource.Intn(l)
		s, base := scope[index][0], scope[index][1]
		result[i] = uint8(base + randSource.Intn(s))
	}
	return result
}

func RandString(len int) string {
	return string(RandStringK(len, RandKindLower))
}
