package main

import (
	"fmt"
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
	app.Usage = "Get public IP address by stun server."
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		{
			Name:  "wei.xuan",
			Email: "adamweixuan@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "stunhost",
			Required: false,
			Usage:    "stun host",
			Value:    "stun.qq.com",
			Aliases:  []string{"s"},
			EnvVars:  []string{ENVHOSTKEY},
		},
		&cli.UintFlag{
			Name:     "stunport",
			Required: false,
			Usage:    "stun port",
			Value:    3478,
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
		host := ctx.String("stunhost")
		port := fmt.Sprintf("%d", ctx.Uint("stunport"))
		addr := net.JoinHostPort(host, port)
		v4Only := ctx.Bool("ipv4")
		return run(addr, v4Only)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(addr string, onlyV4 bool) error {
	network := ifelse(onlyV4, NETWORK4, NETWORK)
	uaddr, err := net.ResolveUDPAddr(network, addr)
	if err != nil {
		return err
	}

	ln, err := net.ListenUDP(network, nil)
	if err != nil {
		return err
	}

	if _, err = ln.WriteToUDP(stun.Request(stun.NewTxID()), uaddr); err != nil {
		return err
	}

	var buf [1024]byte
	n, _, err := ln.ReadFromUDPAddrPort(buf[:])
	if err != nil {
		return err
	}

	_, saddr, err := stun.ParseResponse(buf[:n])
	if err != nil {
		return err
	}

	println(saddr.Addr().String())
	return nil
}

func ifelse(flag bool, trueValue, falseValue string) string {
	if flag {
		return trueValue
	}
	return falseValue
}
