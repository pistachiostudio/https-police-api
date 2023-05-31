package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/syumai/workers"
)

type CheckRequest struct {
	Domain string `json:"domain"`
}

type CheckResponse struct {
	Issuer     string `json:"issuer"`
	Expiry     string `json:"expiry"`
	TlsVersion string `json:"tlsVersion"`
}

var tlsVersions = map[uint16]string{
	tls.VersionSSL30: "SSL",
	tls.VersionTLS10: "TLS 1.0",
	tls.VersionTLS11: "TLS 1.1",
	tls.VersionTLS12: "TLS 1.2",
	tls.VersionTLS13: "TLS 1.3",
}

func httpsCheck(domain string) (CheckResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	d := tls.Dialer{}
	conn, err := d.DialContext(ctx, "tcp", domain+":443")
	cancel()
	if err != nil {
		return CheckResponse{}, err
	}
	defer conn.Close()

	tlsConn := conn.(*tls.Conn)

	err = tlsConn.VerifyHostname(domain)
	if err != nil {
		return CheckResponse{}, err
	}
	expiry := tlsConn.ConnectionState().PeerCertificates[0].NotAfter

	res := CheckResponse{
		Issuer:     tlsConn.ConnectionState().PeerCertificates[0].Issuer.String(),
		Expiry:     expiry.Format(time.RFC850),
		TlsVersion: tlsVersions[tlsConn.ConnectionState().Version],
	}
	return res, nil
}

func CheckHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(fmt.Sprintf("request format is invalid: %s", req.Method)))
		return
	}

	var checkReq CheckRequest
	if err := json.NewDecoder(req.Body).Decode(&checkReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("request format is invalid"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res, err := httpsCheck(checkReq.Domain)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(fmt.Sprintf("failed to check: %s", err.Error())))
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode response: %w\n", err)
	}
}

func main() {
	http.HandleFunc("/", CheckHandler)
	workers.Serve(nil) // use http.DefaultServeMux
}
