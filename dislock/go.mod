module github.com/phpdi/ant/dislock

go 1.13

require (
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.15.2 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/phpdi/ant/redisutil v0.0.0-20210710113852-bb13b8ea732f
	github.com/prometheus/client_golang v1.8.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/genproto v0.0.0-20201026171402-d4b8fe4fd877 // indirect
	google.golang.org/grpc v1.33.1 // indirect
	google.golang.org/grpc/examples v0.0.0-20201028002921-15a78f19307d // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace github.com/phpdi/ant => ../

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
