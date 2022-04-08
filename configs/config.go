package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

type (
	Config struct {
		Environment string
		Postgres    PostgresConfig
		GRPC        GRPCConfig
	}

	PostgresConfig struct {
		Port     string
		Sslmode  string
		Host     string
		Username string
		Dbname   string
		Password string
	}

	GRPCConfig struct {
		Port string
	}
)

func Init(configsDir string) (*Config, error) {

	if err := parseConfigFile(configsDir); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {

	if err := viper.UnmarshalKey("db", &cfg.Postgres); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("grpc", &cfg.GRPC); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	_ = godotenv.Load()
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")
}
