package biq

import (
	"math/big"
)

// Int 扩展了 math/big.Int 的功能
type Int struct {
	*big.Int
}

// NewInt 创建新的 biq.Int 实例
func NewInt(x int64) *Int {
	return &Int{big.NewInt(x)}
}

// SetByte 将整数设置为单个字节的值 (0-255)
// 这是对标准库 big.Int 的功能扩展
func (z *Int) SetByte(b byte) *Int {
	if z.Int == nil {
		z.Int = new(big.Int)
	}
	z.SetUint64(uint64(b))
	return z
}

// SetBytes 重写标准方法以返回 biq.Int 类型
// 保持与标准库兼容性的同时支持链式调用
// func (z *Int)
func (z *Int) SetBytes(buf []byte) *big.Int {
	if z.Int == nil {
		z.Int = big.NewInt(1)
	}

	z.Int.Div(big.NewInt(0).SetBytes(buf), z.Int)
	return z.Int
}

// 以下方法覆盖所有 big.Int 的 Set 方法
// 确保链式调用时返回 biq.Int 类型

func (z *Int) Set(x *big.Int) *Int {
	z.Int.Set(x)
	return z
}

func (z *Int) SetInt64(x int64) *Int {
	z.Int.SetInt64(x)
	return z
}

func (z *Int) SetUint64(x uint64) *Int {
	z.Int.SetUint64(x)
	return z
}

func (z *Int) SetString(s string, base int) (*Int, bool) {
	_, ok := z.Int.SetString(s, base)
	return z, ok
}

func (z *Int) Cmp(x *big.Int) (r int) {
	return z.Int.Cmp(x)
}
