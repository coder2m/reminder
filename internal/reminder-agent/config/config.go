/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/5 0:04
 */
package config

import (
	"fmt"
	"github.com/myxy99/reminder/pkg/client/email"
	"github.com/myxy99/reminder/pkg/client/rabbitmq"
	"github.com/spf13/viper"
)

const (
	// DefaultConfigurationName is the default name of configuration
	defaultConfigurationName = "config"

	// DefaultConfigurationPath the default location of the configuration file
	defaultConfigurationPath = "./config"
)

type Cfg struct {
	RabbitMq *rabbitmq.Options `yaml:"rabbitMQ"`
	Email    *email.Options    `yaml:"email"`
}

func New() *Cfg {
	return &Cfg{
		RabbitMq: rabbitmq.NewRabbitMQOptions(),
		Email:    email.NewEmailOptions(),
	}
}

func TryLoadFromDisk() (*Cfg, error) {
	viper.SetConfigName(defaultConfigurationName)
	viper.AddConfigPath(defaultConfigurationPath)

	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, err
		} else {
			return nil, fmt.Errorf("error parsing configuration file %s", err)
		}
	}

	conf := New()

	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
