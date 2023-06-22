package object

type Configs struct {
	RedisAddress  string `mapstructure:"redis_address"`
	RedisPassword string `mapstructure:"redis_password"`
}
