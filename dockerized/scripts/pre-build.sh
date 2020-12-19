#!/bin/bash
set -e

sudo apt update && sudo apt install -y software-properties-common
sudo add-apt-repository -y ppa:longsleep/golang-backports
sudo apt install -y golang-1.15 golang-go
