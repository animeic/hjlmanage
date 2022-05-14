package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// 生成6位随机数
func MakeEmailCode() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(900000))
	randn := n.Int64()
	rn := randn + 100000
	res := strconv.FormatInt(rn, 10)
	return res
}
