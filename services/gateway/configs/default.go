package configs

var (
	DefaultConfig = []byte(`
service_name: "gateway"
version: "v1.0.0"
http_port: 8787
host: "localhost"
api_external_Address: "localhost:8787"
metadata_address: ":11000"
`)
)
