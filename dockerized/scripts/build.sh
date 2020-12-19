#!/bin/bash

./deps.sh
PLATFORMS=$(go env GOHOSTOS)/$(go env GOHOSTARCH) ./build-all.sh

apt update
apt install -y git uuid-dev curl

git clone https://github.com/clark15b/xupnpd2.git
cd xupnpd2
git checkout 0da5ca9c4a796b9ac9da4a706569a3665f75e35e
rm -r media/*
make -f Makefile.linux
