
# Info
```console
docker run -it golang:latest
```
```
# go version go1.15.6 linux/amd64
```
### Build Go standard library as shared
```
go install -buildmode=shared std
```
### Confirm libstd presence
```
root@2532c987ea01:/go# ls -l /usr/local/go/pkg/linux_amd64_dynlink/libstd.so
-rw-r--r-- 1 root root 39730448 Jan 19 14:33 /usr/local/go/pkg/linux_amd64_dynlink/libstd.so
```
### Build the Go library as a shared module
```
go get -u github.com/octaine-dev/lib
go install -buildmode=shared -linkshared github.com/octaine-dev/lib
```
```
root@2532c987ea01:/go# ldd /go/pkg/linux_amd64_dynlink/libgithub.com-octaine-dev-lib.so
        linux-vdso.so.1 (0x00007ffeaf379000)
        libstd.so => /usr/local/go/pkg/linux_amd64_dynlink/libstd.so (0x00007fc55e61f000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fc55e45a000)
        libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007fc55e455000)
        libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007fc55e434000)
        /lib64/ld-linux-x86-64.so.2 (0x00007fc560d5d000)
```
# Build the Go executable that depends on the lib
```
go get github.com/octaine-dev/host
go install -linkshared github.com/octaine-dev/host
```
```
root@2532c987ea01:~# ldd /go/bin/host
        linux-vdso.so.1 (0x00007fffcd387000)
        libstd.so => /usr/local/go/pkg/linux_amd64_dynlink/libstd.so (0x00007f37dd250000)
        libgithub.com-octaine-dev-lib.so => /go/pkg/linux_amd64_dynlink/libgithub.com-octaine-dev-lib.so (0x00007f37dd248000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f37dd083000)
        libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007f37dd07e000)
        libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f37dd05d000)
        /lib64/ld-linux-x86-64.so.2 (0x00007f37df98d000)
```
```
root@2532c987ea01:~# /go/bin/host
RNG callback  [1 2 3]
rng.Next():  [4 5 6]
Response from lib  &{100 true [random engine data] [outcome data]}
```
