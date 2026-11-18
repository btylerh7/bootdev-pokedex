package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config, args ...string) error
}

type Config struct {
	Commands      map[string]CliCommand
	NextUrl       *string
	PrevUrl       *string
	CaughtPokemon []string
}
