#!/bin/bash
set -e

./deps.sh
PLATFORMS=$(go env GOHOSTOS)/$(go env GOHOSTARCH) ./build-all.sh

sudo apt update
sudo apt install -y git uuid-dev curl build-essential

git clone https://github.com/clark15b/xupnpd2.git
cd xupnpd2
git checkout 0da5ca9c4a796b9ac9da4a706569a3665f75e35e
rm -r media/*
make -f Makefile.linux
