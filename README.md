# C0verter
Integrated tunneling tool with covert channel support, transmitting TCP/UDP over IPv6,ICMP,DNS,HTTP,HTTPS.

## Usage
Command line:
- in: address inside the tunnel, local or remote address.
- out: target port outside the tunnel, will use command line IO if not defined.
- covert: which type of covert channel to use, none as default.
- proxy: use socks5(?) proxy, set to false by default.
```shellscript
C0 --in TunnelAddr [--out TargetAddr] [--covert http/https/icmp/dns] [--proxy false/true]
C0 -i TunnelAddr [-o TargetAddr] [-c http/https/icmp/dns] [-p false/true]
```
Examples:
```shellscript
C0 -i localhost:8080 -o web:80 
C0 -i vps:5000 -o internal:3389 --covert https --crt server.crt --key server.key
C0 -i :5000 -o cmd.exe
```
