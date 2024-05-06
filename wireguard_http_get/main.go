package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/netip"
	"os"

	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun/netstack"
)

func main() {
	if len(os.Args) != 6 {
		log.Panic("Usage: wireguard_http_get <privateKey> <privateIP> <peerPublicKey> <peerEndpoint> <httpEndpoint>")
	}

	privateKey := os.Args[1]
	privateIP := os.Args[2]
	peerPublicKey := os.Args[3]
	fmt.Println(peerPublicKey)
	peerEndpoint := os.Args[4]
	httpEndpoint := os.Args[5]

	tun, tnet, err := netstack.CreateNetTUN(
		[]netip.Addr{netip.MustParseAddr(privateIP)},
		[]netip.Addr{},
		1420)
	if err != nil {
		log.Panic(err)
	}
	dev := device.NewDevice(tun, conn.NewDefaultBind(), device.NewLogger(device.LogLevelVerbose, ""))
	// echo SHpP1UAoEGCHE1hxMr93NZnqUw1gva5seycgNXqSa30= | base64 -d | xxd -p -c32
	err = dev.IpcSet(`private_key=` + privateKey + `
public_key=` + peerPublicKey + `
allowed_ip=0.0.0.0/0
endpoint=` + peerEndpoint + `
persistent_keepalive_interval=25
`)
	if err != nil {
		log.Panic(err)
	}
	err = dev.Up()
	if err != nil {
		log.Panic(err)
	}

	client := http.Client{
		Transport: &http.Transport{
			DialContext: tnet.DialContext,
		},
	}
	resp, err := client.Get(httpEndpoint)
	if err != nil {
		log.Panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(body))
}
