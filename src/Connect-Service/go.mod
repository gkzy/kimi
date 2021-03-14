module github.com/gkzy/kimi/src/Connect-Service

go 1.15

replace github.com/gkzy/kimi/common => /Users/samsong/mygo/src/github.com/gkzy/kimi/common

require (
	github.com/gkzy/gow v0.5.1
	github.com/gkzy/kimi/common v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/websocket v1.4.2
	google.golang.org/grpc v1.36.0
)
