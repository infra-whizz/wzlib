module github.com/infra-whizz/wzlib

go 1.13

require (
	github.com/elastic/go-sysinfo v1.3.0
	github.com/nats-io/nats-server/v2 v2.1.4 // indirect
	github.com/nats-io/nats.go v1.9.1
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/oklog/ulid v1.3.1
	github.com/vmihailenco/msgpack/v4 v4.3.11
	golang.org/x/crypto v0.0.0-20200219234226-1ad67e1f0ef4 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/infra-whizz/wzlib => /home/bo/work/golang/infra-whizz/wzlib
