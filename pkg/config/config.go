package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// StaiConfig the STAI config.yaml
type StaiConfig struct {
	StaiRoot        string
	DaemonPort      uint16          `yaml:"daemon_port"`
	DaemonSSL       SSLConfig       `yaml:"daemon_ssl"`
	Farmer          FarmerConfig    `yaml:"farmer"`
	FullNode        FullNodeConfig  `yaml:"full_node"`
	Harvester       HarvesterConfig `yaml:"harvester"`
	Wallet          WalletConfig    `yaml:"wallet"`
	Seeder          SeederConfig    `yaml:"seeder"`
	SelectedNetwork string          `yaml:"selected_network"`
}

// FarmerConfig farmer configuration section
type FarmerConfig struct {
	PortConfig `yaml:",inline"`
	SSL        SSLConfig `yaml:"ssl"`
}

// FullNodeConfig full node configuration section
type FullNodeConfig struct {
	PortConfig      `yaml:",inline"`
	SSL             SSLConfig `yaml:"ssl"`
	SelectedNetwork string    `yaml:"selected_network"`
	DatabasePath    string    `yaml:"database_path"`
}

// HarvesterConfig harvester configuration section
type HarvesterConfig struct {
	PortConfig `yaml:",inline"`
	SSL        SSLConfig `yaml:"ssl"`
}

// WalletConfig wallet configuration section
type WalletConfig struct {
	PortConfig `yaml:",inline"`
	SSL        SSLConfig `yaml:"ssl"`
}

// SeederConfig seeder configuration section
type SeederConfig struct {
	CrawlerConfig CrawlerConfig `yaml:"crawler"`
}

// CrawlerConfig is the subsection of the seeder config specific to the crawler
type CrawlerConfig struct {
	PortConfig `yaml:",inline"`
	SSL        SSLConfig `yaml:"ssl"`
}

// PortConfig common port settings found in many sections of the config
type PortConfig struct {
	Port    uint16 `yaml:"port"`
	RPCPort uint16 `yaml:"rpc_port"`
}

// SSLConfig common ssl settings found in many sections of the config
type SSLConfig struct {
	PrivateCRT string `yaml:"private_crt"`
	PrivateKey string `yaml:"private_key"`
	PublicCRT  string `yaml:"public_crt"`
	PublicKey  string `yaml:"public_key"`
}

// GetStaiConfig returns a struct containing the config.yaml values
func GetStaiConfig() (*StaiConfig, error) {
	rootPath, err := GetStaiRootPath()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(rootPath, "config", "config.yaml")
	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found")
	}

	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	config := &StaiConfig{}

	err = yaml.Unmarshal(configBytes, config)
	if err != nil {
		return nil, err
	}

	config.StaiRoot = rootPath
	config.fillDatabasePath()

	return config, nil
}

// GetStaiRootPath returns the root path for the STAI installation
func GetStaiRootPath() (string, error) {
	if root, ok := os.LookupEnv("STAI_ROOT"); ok {
		return root, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	root := filepath.Join(home, ".stai", "mainnet")

	return root, nil
}

// GetFullPath returns the full path to a particular filename within STAI_ROOT
func (c *StaiConfig) GetFullPath(filename string) string {
	return filepath.Join(c.StaiRoot, filename)
}

func (c *StaiConfig) fillDatabasePath() {
	c.FullNode.DatabasePath = strings.Replace(c.FullNode.DatabasePath, "CHALLENGE", c.FullNode.SelectedNetwork, 1)
}
