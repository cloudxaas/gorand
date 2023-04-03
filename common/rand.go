package cxrand

import (
    crand "github.com/lukechampine/frand"
    fastrand "github.com/bytedance/gopkg/lang/fastrand"
)

//0a
func CRandomBytes(n int) []byte {
    bytes := make([]byte, n)
    crand.Read(bytes)
    return bytes
}

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
