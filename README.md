# lynette

Cotainer runtime learning project.

## How to build ?

```sh
$ mage -v build
Running target: Build
Running dependency: downloadRootfs
Running dependency: buildCoverage
exec: mkdir "-p" "build/rootfs"
Running dependency: buildGeneric
exec: wget "https://cdimage.ubuntu.com/ubuntu-base/releases/22.04/release/ubuntu-base-22.04-base-amd64.tar.gz" "-O" "build/ubuntu.tar.gz"
--2023-10-23 09:12:49--  https://cdimage.ubuntu.com/ubuntu-base/releases/22.04/release/ubuntu-base-22.04-base-amd64.tar.gz
Resolving cdimage.ubuntu.com (cdimage.ubuntu.com)... 91.189.91.123, 91.189.91.124, 185.125.190.37, ...
Connecting to cdimage.ubuntu.com (cdimage.ubuntu.com)|91.189.91.123|:443... exec: go "build" "-cover" "-o" "./build/lynette_coverage" "cmd/lynette/lynette.go"
exec: go "build" "-o" "./build/lynette" "cmd/lynette/lynette.go"
connected.
HTTP request sent, awaiting response... 200 OK
Length: 29824980 (28M) [application/x-gzip]
Saving to: ‘build/ubuntu.tar.gz’

build/ubuntu.tar.gz                                      100%[==================================================================================================================================>]  28.44M  13.6MB/s    in 2.1s    

2023-10-23 09:12:52 (13.6 MB/s) - ‘build/ubuntu.tar.gz’ saved [29824980/29824980]

exec: tar "xf" "build/ubuntu.tar.gz" "-C" "build/rootfs"
```

## How to run ?

```sh
$ ./build/lynette run "$(pwd)/build/rootfs/" bash
time=2023-10-23T09:11:02.750+02:00 level=DEBUG msg="Executing command" command="/proc/self/exe run /home/jdaniecki/projects/lynette/build/rootfs/ bash"
time=2023-10-23T09:11:02.753+02:00 level=DEBUG msg="Setting up hostname" hostname=container
time=2023-10-23T09:11:02.754+02:00 level=DEBUG msg="Setting up rootfs" rootfs=/home/jdaniecki/projects/lynette/build/rootfs/
time=2023-10-23T09:11:02.754+02:00 level=DEBUG msg="Executing command" command=/usr/bin/bash
root@container:/# 
```