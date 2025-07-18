// 系统配置
package config

import (
	"fmt"
	"strings"
)

type Config struct {
	App struct {
		Name        string `mapstructure:"Name"`
		ListenPath  string `mapstructure:"ListenPath"`
		MemberFile  string `mapstructure:"MemberFile"`
		FrontEndDir string `mapstructure:"FrontEndDir"`
		TemplateDir string `mapstructure:"TemplateDir"`
	} `mapstructure:"app"`
	Option struct {
		DatabaseAutoMigrate bool `mapstructure:"DatabaseAutoMigrate"`
		Debug               bool `mapstructure:"Debug"`
	} `mapstructure:"option"`
	DB struct {
		Type   string `mapstructure:"Type"`
		Path   string `mapstructure:"Path"`
		Port   int    `mapstructure:"Port"`
		Name   string `mapstructure:"Name"`
		User   string `mapstructure:"User"`
		Passwd string `mapstructure:"Passwd"`
	} `mapstructure:"DB"`
	Business struct {
		Year      string `mapstructure:"Year"`
		Semester  int    `mapstructure:"Semester"`
		StartTime string `mapstructure:"StartTime"`
		Week      int    `mapstructure:"Week"`
	} `mapstructure:"business"`
}

var pathToConfigure string //配置文件的路径
var Default Config         //系统的默认配置

func (c *Config) String() string {
	var builder strings.Builder

	builder.WriteString("\n\n\n\n\n")

	builder.WriteString("Configuration:\n")
	builder.WriteString("--------------\n")

	// App Section
	builder.WriteString("[App]\n")
	builder.WriteString(fmt.Sprintf("  Name:         %s\n", c.App.Name))
	builder.WriteString(fmt.Sprintf("  ListenPath:   %s\n", c.App.ListenPath))
	builder.WriteString(fmt.Sprintf("  MemberFile:   %s\n", c.App.MemberFile))
	builder.WriteString(fmt.Sprintf("  FrontEndDir:  %s\n", c.App.FrontEndDir))
	builder.WriteString(fmt.Sprintf("  TemplateDir:  %s\n", c.App.TemplateDir))
	builder.WriteString("\n")

	// Option Section
	builder.WriteString("[Option]\n")
	builder.WriteString(fmt.Sprintf("  Auto Migrate: %t\n", c.Option.DatabaseAutoMigrate))
	builder.WriteString(fmt.Sprintf("  Debug Mode:   %t\n", c.Option.Debug))
	builder.WriteString("\n")

	// DB Section
	builder.WriteString("[Database]\n")
	builder.WriteString(fmt.Sprintf("  Type:         %s\n", c.DB.Type))
	builder.WriteString(fmt.Sprintf("  Path:         %s\n", c.DB.Path))
	builder.WriteString(fmt.Sprintf("  Port:         %d\n", c.DB.Port))
	builder.WriteString(fmt.Sprintf("  Name:         %s\n", c.DB.Name))
	builder.WriteString(fmt.Sprintf("  User:         %s\n", c.DB.User))
	builder.WriteString(fmt.Sprintf("  Password:     %s\n", c.DB.Passwd))
	builder.WriteString("\n")

	// Business Section
	builder.WriteString("[Business]\n")
	builder.WriteString(fmt.Sprintf("  Year:         %s\n", c.Business.Year))
	builder.WriteString(fmt.Sprintf("  Semester:     %d\n", c.Business.Semester))
	builder.WriteString(fmt.Sprintf("  StartTime:    %s\n", c.Business.StartTime))
	builder.WriteString(fmt.Sprintf("  Week:         %d\n", c.Business.Week))
	builder.WriteString("--------------\n\n\n\n\n")

	return builder.String()
}
