// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package loadtest

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type LoadTestConfig struct {
	LoadtestEnviromentConfig  LoadtestEnviromentConfig
	ConnectionConfiguration   ConnectionConfiguration
	UserEntitiesConfiguration UserEntitiesConfiguration
	ResultsConfiguration      ResultsConfiguration
	LogSettings               LoggerSettings
}

type UserEntitiesConfiguration struct {
	TestLengthMinutes                 int
	NumActiveEntities                 int
	ActionRateMilliseconds            int
	ActionRateMaxVarianceMilliseconds int
	EnableRequestTiming               bool
	ChannelLinkChance                 float64
	UploadImageChance                 float64
	DoStatusPolling                   bool
	RandomizeEntitySelection          bool
}

type ConnectionConfiguration struct {
	ServerURL            string
	WebsocketURL         string
	PProfURL             string
	DriverName           string
	DataSource           string
	DBEndpoint           string // deprecated
	LocalCommands        bool
	SSHHostnamePort      string
	SSHUsername          string
	SSHPassword          string
	SSHKey               string
	MattermostInstallDir string
	ConfigFileLoc        string
	AdminEmail           string
	AdminPassword        string
	SkipBulkload         bool
	WaitForServerStart   bool
}

type ResultsConfiguration struct {
	CustomReportText     string
	SendReportToMMServer bool
	ResultsServerURL     string
	ResultsChannelId     string
	ResultsUsername      string
	ResultsPassword      string
	PProfDelayMinutes    int
	PProfLength          int
}

type LoggerSettings struct {
	EnableConsole bool
	ConsoleJson   bool
	ConsoleLevel  string
	EnableFile    bool
	FileJson      bool
	FileLevel     string
	FileLocation  string
}

func ReadConfig() error {
	viper.SetConfigName("loadtestconfig")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")
	viper.SetEnvPrefix("mmloadtest")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("LogSettings.EnableConsole", true)
	viper.SetDefault("LogSettings.ConsoleLevel", "INFO")
	viper.SetDefault("LogSettings.ConsoleJson", true)
	viper.SetDefault("LogSettings.EnableFile", true)
	viper.SetDefault("LogSettings.FileLevel", "INFO")
	viper.SetDefault("LogSettings.FileJson", true)
	viper.SetDefault("LogSettings.FileLocation", "loadtest.log")

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "unable to read configuration file")
	}

	return nil
}

func GetConfig() (*LoadTestConfig, error) {
	var cfg *LoadTestConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
