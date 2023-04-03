package cxrand

import (
    "crypto/rand"
)

func RandomBytes(n int) []byte {
    bytes := make([]byte, n)
    rand.Read(bytes)
    return bytes
}
