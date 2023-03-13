package config

var CLI struct {
	Verbose bool   `help:"Verbose mode."`
	Config  string `help:"Config file." name:"config" type:"file" default:"config.json"`
}
