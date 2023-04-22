# tip
Get public IP address by stun server.

# Install
```shell
go install github.com/adamweixuan/tip@latest
```

# Usage
```shell
./tip -h
NAME:
   tip - Get public IP address by stun server.

USAGE:
   tip [global options] command [command options] [arguments...]

AUTHOR:
   wei.xuan <adamweixuan@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --stunhost value, -s value  stun host (default: "stun.qq.com") [$STUN_HOST]
   --stunport value, -p value  stun port (default: 3478) [$STUN_PORT]
   --ipv4, -4                  use ipv4 only (default: false)
   --help, -h                  show help
```

## use default stun server
```shell
./tip
101.204.117.54
```
## custom stun server
> from flag
```shell
./tip -s stun.sipnet.net -p 3478
```

> from env
```shell
STUN_HOST=stun.l.google.com STUN_PORT=19302 ./tip
```
# public stun server 

[https://gist.github.com/mondain/b0ec1cf5f60ae726202e](https://gist.github.com/mondain/b0ec1cf5f60ae726202e)