#!/bin/bash
go build
sudo setcap 'cap_net_raw,cap_net_admin=eip' wave-to-mqtt
source variables.sh
./wave-to-mqtt