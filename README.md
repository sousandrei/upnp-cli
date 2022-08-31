## UPnP Cli

Simple cli to open ports on your router using UPnP.

### Installation

```bash
go install github.com/sousandrei/upnp-cli
```

### Usage

```
➜ ./upnp --help
Usage of ./build/upnp:
  -d uint
        duration in seconds (default 30)
  -e string
        external ip:port
  -i string
        internal ip:port
  -n string
        rule description/name
  -u    is udp, otherwise tcp
```
