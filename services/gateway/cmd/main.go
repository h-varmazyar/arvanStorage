package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/arvanStorage/services/gateway/configs"
	"github.com/h-varmazyar/arvanStorage/services/gateway/docs"
	"github.com/h-varmazyar/arvanStorage/services/gateway/internal/metadata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a Arvan cloud API document.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	ctx := context.Background()
	logger := log.New()

	defaultConf := flag.Bool("default-configs", false, "run program with default config")
	flag.Parse()

	conf, err := loadConfigs(*defaultConf)
	if err != nil {
		log.Panicf("failed to read configs: %v", err)
	}

	initializeDocs(conf)

	logger.Infof("running %v(%v)", conf.ServiceName, conf.Version)

	initializeAndRegisterApps(ctx, logger, conf)
}

func loadConfigs(defaultConfig bool) (*Configs, error) {
	log.Infof("reding configs...")

	if defaultConfig {
		viper.SetConfigType("yaml")
		log.Infof("reading deafult configs")
		err := viper.ReadConfig(bytes.NewBuffer(configs.DefaultConfig))
		if err != nil {
			log.WithError(err).Error("read from default configs failed")
			return nil, err
		}
	} else {
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Warnf("failed to read from env: %v", err)
			viper.AddConfigPath("./configs")  //path for docker compose configs
			viper.AddConfigPath("../configs") //path for local configs
			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
			if err = viper.ReadInConfig(); err != nil {
				log.Warnf("failed to read from yaml: %v", err)
				localErr := viper.ReadConfig(bytes.NewBuffer(configs.DefaultConfig))
				if localErr != nil {
					log.WithError(localErr).Error("read from default configs failed")
					return nil, localErr
				}
			}
		}
	}

	conf := new(Configs)
	if err := viper.Unmarshal(conf); err != nil {
		log.Errorf("faeiled unmarshal")
		return nil, err
	}

	return conf, nil
}

func initializeDocs(configs *Configs) {
	docs.SwaggerInfo.Title = "The Arvan Storage API document"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = configs.Version
	docs.SwaggerInfo.Host = configs.ApiExternalAddress
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func initializeAndRegisterApps(_ context.Context, logger *log.Logger, configs *Configs) {
	engine := gin.Default()
	metadata.RegisterRoutes(logger, engine, configs.MetadataAddress)

	docs.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := engine.Run(fmt.Sprintf(":%v", configs.HttpPort))
	if err != nil {
		log.WithError(err).Panicf("failed to start controller")
	}
}
