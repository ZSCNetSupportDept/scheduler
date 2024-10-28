package config

var (
	ListenPort int
	File       string
	DB         database
)

type database struct {
	Path string
	//Port     string
	//User     string
	//Password string

	// enable if you want use an instance other than SQLite
	Type string
}
