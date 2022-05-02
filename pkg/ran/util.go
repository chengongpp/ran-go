package ran

import (
	"math/rand"
	"net"
)

func probeLocalAddresses() ([]net.Addr, error) {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func randomString(n int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	name := make([]byte, n)
	for i := 0; i < n; i++ {
		name[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(name)
}
