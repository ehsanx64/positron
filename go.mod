module github.com/ehsanx64/positron

go 1.22.0

// replace github.com/ehsanx64/positron/infra/shared => ../positron/infra/shared

require github.com/eclipse/paho.mqtt.golang v1.5.0

require (
	github.com/gorilla/websocket v1.5.3 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
)
