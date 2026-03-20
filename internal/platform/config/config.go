package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Snowflake SnowflakeConfig `mapstructure:"snowflake"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type SnowflakeConfig struct {
	NodeID int64 `mapstructure:"node_id"`
}

// LoadConfig 解析 yaml 文件与环境变量
func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")    // 本地运行能找到
	viper.AddConfigPath("/app/configs") // 容器生产运行能找到
	viper.AddConfigPath(".")            // 备用路径

	// 支持环境变量，配置的下划线分隔，比如 SERVER_PORT
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("反序列化配置失败: %v", err)
	}

	return &cfg
}
