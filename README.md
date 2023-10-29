# lynette

Cotainer runtime learning project.

## How to build ?

```sh
$ mage -v build
Running target: Build
Running dependency: buildGeneric
exec: go "build" "-o" "./build/lynette" "cmd/lynette/lynette.go"
Running dependency: buildCoverage
exec: go "build" "-cover" "-o" "./build/lynette_coverage" "cmd/lynette/lynette.go"
Running dependency: ensureRootfs
exec: mkdir "-p" "build/rootfs"
exec: sh "-c" "docker export $(docker create busybox) | tar -C build/rootfs -xf -"
```

## How to run ?

```sh
$ mage -v run
Running target: Run
exec: build/lynette "run" "build/rootfs" "sh"
time=2023-10-29T11:22:05.169+01:00 level=DEBUG msg="Creating bridge..." name=lynette0
time=2023-10-29T11:22:05.170+01:00 level=DEBUG msg="Executing command" command="/proc/self/exe run build/rootfs sh"
time=2023-10-29T11:22:05.171+01:00 level=DEBUG msg="Setting up hostname" hostname=container
time=2023-10-29T11:22:05.171+01:00 level=DEBUG msg="Setting up rootfs..."
time=2023-10-29T11:22:05.172+01:00 level=DEBUG msg="Changing root" root=build/rootfs
time=2023-10-29T11:22:05.172+01:00 level=DEBUG msg="Mounting proc" proc=/proc
time=2023-10-29T11:22:05.172+01:00 level=DEBUG msg="Executing command" command=/bin/sh
/ # ps
PID   USER     TIME  COMMAND
    1 root      0:00 /proc/self/exe run build/rootfs sh
    7 root      0:00 sh
    8 root      0:00 ps
```