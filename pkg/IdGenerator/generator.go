package IdGenerator

import (
	"crypto/rand"
	"math/big"
	"sync"
)

/*
	全局 ID 生成算法：薄雾算法，详情见：github.com/asyncins/mist
*/

type Mist struct {
	sync.Mutex
	increase int64
	saltA    int64
	saltB    int64
}

const (
	saltBit       = uint(8)
	saltShift     = uint(8)
	increaseShift = saltBit + saltShift
)

func (m *Mist) GenerateID() int64 {
	m.Lock()
	m.increase++
	// 获取随机因子数值 ｜ 使用真随机函数提高性能
	randA, _ := rand.Int(rand.Reader, big.NewInt(255))
	m.saltA = randA.Int64()
	randB, _ := rand.Int(rand.Reader, big.NewInt(255))
	m.saltB = randB.Int64()
	// 通过位运算实现自动占位
	mist := (m.increase << increaseShift) | (m.saltA << saltShift) | m.saltB
	m.Unlock()
	return mist
}

func NewMist() *Mist {
	mist := Mist{increase: 1}
	return &mist
}
