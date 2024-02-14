package config

type Config struct {
	App    App
	Server Server
	DB     DB
}

type App struct {
	Name string
}

type Server struct {
	Host                        string
	Port                        string
	MemoryLimitForMultipartForm int64
}

type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}
