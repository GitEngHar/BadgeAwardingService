package echo

type Repository struct {
	config *Config
}

func NewEchoRepository(config *Config) *Repository {
	return &Repository{config: config}
}

func (er Repository) Run() {
	config := er.config
	server := config.server
	config.router.Do()
	server.Logger.Fatal(server.Start(config.port))
}
