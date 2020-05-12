package sslcheck
import (
	"crypto/tls"
	"fmt"
	"time"
)
func GetCertsExpire(domain string, port string) (int, error){
	conn, err := tls.Dial("tcp", domain + ":" + port, nil)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	if err = conn.VerifyHostname(domain); err != nil {
		return 0, err
	}
	return int(conn.ConnectionState().PeerCertificates[0].NotAfter.Unix() - time.Now().Unix()), nil
}