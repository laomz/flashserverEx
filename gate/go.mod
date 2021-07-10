module gate

go 1.16

replace (
	common => ../common
	framework => ../framework
)

require (
	common v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.2 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
