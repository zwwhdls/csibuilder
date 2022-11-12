module {{ .Repo }}

go {{ .GoVersion }}

require (
	github.com/container-storage-interface/spec v1.6.0
	google.golang.org/grpc v1.26.0
	k8s.io/klog v1.0.0
)

require (
	github.com/golang/protobuf v1.3.2 // indirect
	golang.org/x/net v0.0.0-20190311183353-d8887717615a // indirect
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55 // indirect
)
