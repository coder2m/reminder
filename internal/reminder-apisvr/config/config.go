/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:23
 */
package config

import (
	"fmt"
	"github.com/myxy99/reminder/internal/reminder-apisvr/config/server"
	"github.com/myxy99/reminder/pkg/client/database"
	"github.com/myxy99/reminder/pkg/client/rabbitmq"
	"github.com/myxy99/reminder/pkg/reminder"
	"github.com/spf13/viper"
)

const (
	// DefaultConfigurationName is the default name of configuration
	defaultConfigurationName = "config"

	// DefaultConfigurationPath the default location of the configuration file
	defaultConfigurationPath = "./config"
)

type Cfg struct {
	Database *database.Options `yaml:"database"`
	Server   *server.Options   `yaml:"server"`
	Reminder *reminder.Options `yaml:"reminder"`
	RabbitMq *rabbitmq.Options `yaml:"rabbitMQ"`
}

func New() *Cfg {
	return &Cfg{
		Database: database.NewDatabaseOptions(),
		Server:   server.NewDefault(),
		Reminder: reminder.NewReminderOptions(),
		RabbitMq: rabbitmq.NewRabbitMQOptions(),
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
