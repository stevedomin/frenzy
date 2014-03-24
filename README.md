frenzy
======

Experimental Vagrant clone written in Go and using Docker instead of VBox.

## Why?

Vagrant is an awesome tool but because it's using VMs it's quite slow.
If you use it to, for example, test your Chef cookbooks (my case), you quickly loose a lot of time during the day, waiting.
So I wondered if I could build something that would do the same job but using containers instead of VMs.

I could have use docker-provider or even vagrant-lxc but to be honest I also wanted to have some fun :)

Keep in mind this is **experimental** so things will probably break, and you can't do much yet.

## Install

### From source

```bash
$ go get github.com/stevedomin/frenzy
$ cd $GOPATH/src/github.com/stevedomin/frenzy/frenzy
$ go install
```

### Binary releases

For now you'll have to build it yourself.

## Requirements

### OS X

```bash
$ # Install [boot2docker](https://github.com/boot2docker/boot2docker)
$ boot2docker stop
$ # Forward VM port range that Docker uses to host
$ for i in {49000..49900}; do
$   VBoxManage modifyvm "boot2docker-vm" --natpf1 "tcp-port$i,tcp,,$i,,$i";
$   VBoxManage modifyvm "boot2docker-vm" --natpf1 "udp-port$i,udp,,$i,,$i";
$ done
$ boot2docker start
```

### Linux

Make sure you have Docker [installed](http://docs.docker.io/en/latest/installation/)

## Known caveats

* You need to be able to execute docker commands without sudo

## TODO

* Some testing would be good :)
* Move to Docker API instead of Docker CLI
* In progress: Stop command (commit container with image name == node name)
* In progress: Chef provisioner
* Support for volumes
* More networking options
