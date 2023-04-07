package cxrand

import (
	"crypto/rand"
	"time"

	fastrand "github.com/bytedance/gopkg/lang/fastrand"
)

func RandomBytes(b []byte) {
	n := len(b)
	var randNum uint32
	for i := 0; i < n; i += 4 {
		randNum = fastrand.Uint32()
		if i+3 < n {
			b[i] = byte(randNum)
			b[i+1] = byte(randNum >> 8)
			b[i+2] = byte(randNum >> 16)
			b[i+3] = byte(randNum >> 24)
		} else {
			for j := 0; j < n-i; j++ {
				b[i+j] = byte(randNum >> (8 * j))
			}
		}
	}
}

func CRandomBytes(b []byte) {
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
}

func RandomString(b []byte) {
	length := len(b)
	var t int64
	for i := 0; i < length; i++ {
		t = time.Now().UnixNano()
		randByte := byte(t & 0xff)
		if randByte < 32 || randByte > 126 {
			randByte = 32
		}
		b[i] = randByte
	}
}
