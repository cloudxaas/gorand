package cxrand

import (
    "unsafe"
    "time"
    crand "lukechampine.com/frand"
    fastrand "github.com/bytedance/gopkg/lang/fastrand"
)

//0a
func RandomBytes(n int) []byte {
    bytes := make([]byte, n)
    var randNum uint32
    for i := 0; i < n; i += 4 {
        randNum = fastrand.Uint32()
        if i+3 < n {
            bytes[i] = byte(randNum)
            bytes[i+1] = byte(randNum >> 8)
            bytes[i+2] = byte(randNum >> 16)
            bytes[i+3] = byte(randNum >> 24)
        } else {
            for j := 0; j < n-i; j++ {
                bytes[i+j] = byte(randNum >> (8 * j))
            }
        }
    }
    return bytes
}

//0a, cryptographic random bytes
func CRandomBytes(n int) []byte {
    bytes := make([]byte, n)
    crand.Read(bytes)
    return bytes
}

//0a, random string
func RandomString(length int) string {
    b := make([]byte, length)
    var t int64
    for i := 0; i < length; i++ {
        t = time.Now().UnixNano()
        randByte := byte(t & 0xff)
        if randByte < 32 || randByte > 126 {
            randByte = 32
        }
        b[i] = randByte
    }
    return *(*string)(unsafe.Pointer(&b))
}
