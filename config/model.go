package config

type Config struct {
	App struct {
		Name       string `mapstructure:"Name"`
		ListenPort int    `mapstructure:"ListenPort"`
		File       string `mapstructure:"File"`
	} `mapstructure:"app"`
	DB struct {
		Type     string `mapstructure:"Type"`
		Path     string `mapstructure:"Path"`
		Port     int    `mapstructure:"Port"`
		User     string `mapstructure:"User"`
		Password string `mapstructure:"Password"`
		Name     string `mapstructure:"Name"`
	} `mapstructure:"DB"`
	Business struct {
		Session   string `mapstructure:"Session"`
		Semester  int    `mapstructure:"Semester"`
		StartTime string `mapstructure:"StartTime"`
		Week      int    `mapstructure:"Week"`
	} `mapstructure:"business"`
}

var pathToConfigure string
var Default Config
