module github.com/infra-whizz/wzlib

go 1.13

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/antonfisher/nested-logrus-formatter v1.0.3
	github.com/davecgh/go-spew v1.1.1
	github.com/elastic/go-sysinfo v1.3.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/google/uuid v1.2.0
	github.com/jinzhu/gorm v1.9.12
	github.com/lib/pq v1.3.0 // indirect
	github.com/nats-io/nats-server/v2 v2.1.4 // indirect
	github.com/nats-io/nats.go v1.9.1
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/oklog/ulid v1.3.1
	github.com/shirou/gopsutil v2.20.2+incompatible
	github.com/sirupsen/logrus v1.5.0
	github.com/stretchr/testify v1.3.0
	github.com/vmihailenco/msgpack/v4 v4.3.11
	golang.org/x/crypto v0.0.0-20200219234226-1ad67e1f0ef4 // indirect
	golang.org/x/sys v0.0.0-20191025021431-6c3a3bfe00ae
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/infra-whizz/wzlib => ../wzlib
