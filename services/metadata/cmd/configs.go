package main

import gormext "github.com/h-varmazyar/gopack/gorm"

type Configs struct {
	ServiceName string          `mapstructure:"service_name"`
	Version     string          `mapstructure:"version"`
	GRPCPort    uint16          `mapstructure:"grpc_port"`
	DB          gormext.Configs `mapstructure:"db"`
}
