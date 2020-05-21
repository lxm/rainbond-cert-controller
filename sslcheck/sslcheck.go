package sslcheck

import (
	"crypto/tls"
	"net"
	"time"
)

func GetCertsExpire(domain string, port string) (int, error) {
	dialer := new(net.Dialer)
	dialer.Timeout = 5 * time.Second
	conn, err := tls.DialWithDialer(dialer, "tcp", domain+":"+port, nil)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	if err = conn.VerifyHostname(domain); err != nil {
		return 0, err
	}
	return int(conn.ConnectionState().PeerCertificates[0].NotAfter.Unix() - time.Now().Unix()), nil
}
