module github.com/amcpherran/wave-to-mqtt

go 1.12

replace github.com/AMcPherran/go-wave => /home/andrew/go/src/github.com/amcpherran/go-wave

require (
	github.com/AMcPherran/go-wave v0.0.0-00010101000000-000000000000
	github.com/eclipse/paho.mqtt.golang v1.2.0
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b // indirect
)
