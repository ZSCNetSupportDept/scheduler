package config

import (
	"fmt"
	"os"

	"github.com/golang-module/carbon/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Load() {

	parseArgs()
	readconfig()
	overrides()
	fmt.Printf("%+v\n", Default)

	carbon.SetDefault(carbon.Default{
		Layout:       carbon.DateTimeLayout,
		Timezone:     carbon.PRC,
		WeekStartsAt: carbon.Monday,
		Locale:       "zh-CN", // 取值范围：lang 目录下翻译文件名，不包含文件后缀
	})

}

func readconfig() {
	viper.SetConfigFile(pathToConfigure)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(&Default); err != nil {
		panic(fmt.Errorf("映射配置到结构体失败: %s", err))
	}
}

func parseArgs() {
	pflag.String("config", "./config.yaml", "the path to config file.")
	pflag.Bool("init-db", false, "whether to initialize the database on starting,useful when migrating to a new one.")
	pflag.String("csv-path", "./member.csv", "the CSV file containing member information")
	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()
	pathToConfigure = viper.GetString("config")
	InitDB = viper.GetBool("init-db")
	CSVPath = viper.GetString("csv-path")
}

func overrides() {
	if CSVPath != "" {
		Default.App.File = CSVPath
	}
}
