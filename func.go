package helper

import (
	"net"
	"net/http"
	"time"
)

func GetRealIp() string {
	var req = &http.Request{}
	var ip = req.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = req.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = req.RemoteAddr
	}
	if net.ParseIP(ip) != nil {
		return ip
	}
	host, _, _ := net.SplitHostPort(ip)
	return host
}

func pickUnusedPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0, err
	}
	return port, nil
}

func NanotimeToDatetime(nanoTime int64) string {
	return time.Unix(nanoTime/1e9, 0).Format("2006-01-02 15:04:05")
}