package cxrand

func RandBytes(n int) []byte {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[fastrand.Uint32n(uint32(len(letterBytes)))]
	}
	return b
}


