package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/golang-module/carbon/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Load() {

	i, err := Init()
	if err != nil {
		panic(err)
	}
	Default = *i
	fmt.Println(Default.String())

	carbon.SetDefault(carbon.Default{
		Layout:       carbon.DateTimeLayout,
		Timezone:     carbon.PRC,
		WeekStartsAt: carbon.Monday,
		Locale:       "zh-CN", // 取值范围：lang 目录下翻译文件名，不包含文件后缀
	})

}

func Init() (*Config, error) {

	v := viper.New()

	//解析命令行传过来的参数
	pflag.StringP("config", "c", "", "Path to the configuration file (required).")
	pflag.String("app.name", "ZSCNetworkSupport Scheduler", "Name of the application")
	pflag.String("app.listenpath", ":25005", "HTTP listen path")
	pflag.String("app.memberfile", "members.csv", "Path to the member file")
	pflag.String("app.frontenddir", "./FrontEnd", "Path to the frontend directory")
	pflag.String("app.templatedir", "./template", "Path to the template directory")

	pflag.Bool("option.databaseautomigrate", false, "Enable automatic database migration")
	pflag.Bool("option.debug", false, "Enable debug mode")

	pflag.String("db.type", "SQLite", "Database type (e.g., mysql, sqlite)")

	pflag.String("business.year", "", "Business year")
	pflag.Int("business.semester", 0, "Current semester")
	pflag.String("business.starttime", "", "Start time of the semester")
	pflag.Int("business.week", 0, "Total weeks in the semester")

	pflag.Parse()
	if err := v.BindPFlags(pflag.CommandLine); err != nil {
		return nil, fmt.Errorf("failed to bind command-line flags: %w", err)
	}

	//加载环境变量
	v.BindEnv("db.Path")
	v.BindEnv("db.User")
	v.BindEnv("db.Port")
	v.BindEnv("db.Name")
	v.BindEnv("db.Passwd")
	v.SetEnvPrefix("SCHEDULER")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	//加载配置文件
	configFile := v.GetString("config")
	if configFile == "" {
		pflag.Usage()
		return nil, errors.New("the --config flag is required")
	}
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file '%s': %w", configFile, err)
	}

	// 导出配置
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return &cfg, nil
}
