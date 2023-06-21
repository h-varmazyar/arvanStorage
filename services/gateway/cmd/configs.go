package main

type Configs struct {
	HttpPort           uint16 `mapstructure:"http_port"`
	MetadataAddress    string `mapstructure:"metadata_address"`
	ApiExternalAddress string `mapstructure:"api_external_address"`
	Version            string `mapstructure:"version"`
	ServiceName        string `mapstructure:"service_name"`
}
