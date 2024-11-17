package config

var (
	ListenPort int
	DB         database
	Session    string
	Semester   int
	StartTime  string
	Week       int
	File       string
)

type database struct {
	Type     string
	Path     string
	Port     int
	User     string
	Password string
	Name     string
}
