package config

type Config struct {
	Database database
	Web      web
}

type database struct {
	Type     string
	Server   string
	Port     string
	Database string
	User     string
	Password string
}

type web struct {
	IP         string
	Port       string
	Secret     string
	CSRF       string
	CSRFSecure bool
}
