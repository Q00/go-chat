package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type ServerInfo struct {
	Port int64
}

type Context struct {
	Timeout int
}

type LogInfo struct {
	Level string
}

type PhaseInfo struct {
	Level string
}

type AWSInfo struct {
	Region      string   `mapstructure:"region"`
	AccessKeyID string   `mapstructure:"accessKeyID"`
	SecretKey   string   `mapstructure:"secretKey"`
	Dynamo      DynamoDB `mapstructure:"dynamoDB"`
}

type DynamoDB struct {
	Room DynamoInfo `mapstructure:"room"`
	Chat DynamoInfo `mapstructure:"chat"`
}

type DynamoInfo struct {
	TableName          string `mapstructure:"tableName"`
	ReadCapacityUnits  int    `mapstructure:"readCapacityUnits"`
	WriteCapacityUnits int    `mapstructure:"writeCapacityUnits"`
}

type Configuration struct {
	Phase  PhaseInfo  `mapstructure:"phase"`
	Server ServerInfo `mapstructure:"server"`
	Log    LogInfo    `mapstructure:"logging"`
	AWS    AWSInfo    `mapstructure:"aws"`
}

func loadSingleton() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		log.Fatal("config read error")
		panic(err)
	}
}

// LoadConfig comment something
func LoadConfig() *Configuration {
	loadSingleton()
	config := &Configuration{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("config unmarshal error")
		panic(err)
	}

	return config
}
