package configs

var (
	DefaultConfig = []byte(`
service_name: "metadata"
version: "v1.0.0"
grpc_port: 11000
db:
  type: "postgreSQL"
  username: "postgres"
  password: "postgres"
  host: "localhost"
  port: 5432
  name: "metadata"
  is_ssl_enable: false
`)
)
