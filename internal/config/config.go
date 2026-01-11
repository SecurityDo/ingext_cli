package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds the global configuration
type Config struct {
	Cluster   string `mapstructure:"cluster"`
	Provider  string `mapstructure:"provider"`
	Context   string `mapstructure:"context"`
	Namespace string `mapstructure:"namespace"`
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in ~/.ingext directory with name "config.yaml"
	configPath := filepath.Join(home, ".ingext")
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	_ = viper.ReadInConfig()
}

// SaveConfig writes the current viper configuration to disk
func SaveConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	
	configDir := filepath.Join(home, ".ingext")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		_ = os.Mkdir(configDir, 0755)
	}

	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// WriteConfigAs ensures we save to the specific path
	return viper.WriteConfigAs(filepath.Join(configDir, "config.yaml"))
}
