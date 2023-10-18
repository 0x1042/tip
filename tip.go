package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"tailscale.com/net/stun"
)

const (
	NETWORK    = "udp"
	NETWORK4   = "udp4"
	ENVHOSTKEY = "STUN_HOST"
	ENVPORTKEY = "STUN_PORT"
)

func main() {
	app := cli.NewApp()
	app.Name = "tip"
	app.Usage = "Get Public IP address by stun server."
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		{
			Name:  "wei.xuan",
			Email: "adamweixuan@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "host",
			Required: false,
			Usage:    "stun host",
			Value:    "stun.qq.com",
			Aliases:  []string{"t"},
			EnvVars:  []string{ENVHOSTKEY},
		},
		&cli.StringFlag{
			Name:     "port",
			Required: false,
			Usage:    "stun port",
			Value:    "3478",
			Aliases:  []string{"p"},
			EnvVars:  []string{ENVPORTKEY},
		},
		&cli.BoolFlag{
			Name:     "ipv4",
			Required: false,
			Value:    false,
			Usage:    "use ipv4 only",
			Aliases:  []string{"4"},
		},
	}

	app.Action = func(ctx *cli.Context) error {
		host := ctx.String("host")
		addr := net.JoinHostPort(host, ctx.String("port"))
		v4 := ctx.Bool("ipv4")
		return run(addr, v4)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(addr string, v4 bool) (err error) {
	network := NETWORK
	if v4 {
		network = NETWORK4
	}
	uaddr, err := net.ResolveUDPAddr(network, addr)
	if err != nil {
		return
	}

	ln, err := net.ListenUDP(network, nil)
	if err != nil {
		return
	}

	if _, err = ln.WriteToUDP(stun.Request(stun.NewTxID()), uaddr); err != nil {
		return
	}

	var buf [1024]byte
	n, _, err := ln.ReadFromUDPAddrPort(buf[:])
	if err != nil {
		return
	}

	_, saddr, err := stun.ParseResponse(buf[:n])
	if err != nil {
		return
	}

	println(saddr.Addr().String())
	return
}
