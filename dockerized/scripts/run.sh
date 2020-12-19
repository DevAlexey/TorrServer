#!/bin/bash

ipaddr=$(hostname -I | awk '{print $1}')

while true; do curl "http://$ipaddr:4044/scripts/scan.lua"; sleep 10; done &
./xupnpd2/xupnpd &
./dist/TorrServer-$(go env GOHOSTOS)-$(go env GOHOSTARCH)
