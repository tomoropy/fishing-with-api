package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// マッピング用の構造体
type Config struct {
	DB   DB   `yaml:"db"`
	Auth Auth `yaml:"auth"`
}

type DB struct {
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Auth struct {
	SecretKey string `yaml:"secret_key"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")  // 設定ファイル名を指定
	viper.SetConfigType("yaml")    // 設定ファイルの形式を指定
	viper.AddConfigPath("config/") // ファイルのpathを指定

	err := viper.ReadInConfig() // 設定ファイルを探索して読み取る
	if err != nil {
		return nil, fmt.Errorf("設定ファイル読み込みエラー: %s \n", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s \n", err)
	}

	return &cfg, nil
}
