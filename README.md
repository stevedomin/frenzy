frenzy
======

# Install

## OS X

```bash
$ # Install [boot2docker](https://github.com/boot2docker/boot2docker)
$
$ # Forward VM port range that Docker uses to host
$ for i in {49000..49900}; do
$   VBoxManage modifyvm "boot2docker-vm" --natpf1 "tcp-port$i,tcp,,$i,,$i";
$   VBoxManage modifyvm "boot2docker-vm" --natpf1 "udp-port$i,udp,,$i,,$i";
$ done
$
$ # Install frenzy
$
```

## Linux

```bash
$ # Make sure you have Docker [installed](http://docs.docker.io/en/latest/installation/)
$
$ # Install Frenzy
$
```

## From source

```bash
$ go get github.com/stevedomin/frenzy
$ cd $GOPATH/src/github.com/stevedomin/frenzy/frenzy
$ go install
```

# TODO

* Some tests :)
* In progress: Stop command (commit container with image name == node name)
* In progress: Chef provisioner
