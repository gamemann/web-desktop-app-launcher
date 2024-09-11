package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gamemann/web-desktop-app-launcher/config"
	"github.com/gamemann/web-desktop-app-launcher/web"
)

const HELPMENU = `
Usage: wdal-web --cfgpath <config path> --list --version --help\n\n
\t--cfgpath => Path to config file (default /etc/wdal/conf.json).\n
\t--list -l => Lists config options.\n
\t--version -v => Prints current version.\n
\t--help => Prints help menu.\n
`

const VERSION = "1.0.0"

func main() {
	// Parse command line arguments.
	var list bool
	var version bool
	var help bool

	flag.BoolVar(&list, "l", false, "List config option.")
	flag.BoolVar(&list, "list", false, "List config option.")

	flag.BoolVar(&version, "v", false, "Prints version.")
	flag.BoolVar(&version, "version", false, "Prints version.")

	flag.BoolVar(&help, "h", false, "Prints help menu.")
	flag.BoolVar(&help, "help", false, "Prints help menu.")

	// Check for version or help flags.
	if version {
		fmt.Println(VERSION)

		os.Exit(0)
	}

	if help {
		fmt.Print(HELPMENU)

		os.Exit(0)
	}

	var cfgPath string

	flag.StringVar(&cfgPath, "cfgpath", "/etc/wdal/conf.json", "The path to the config file.")
	flag.StringVar(&cfgPath, "c", "/etc/wdal/conf.json", "The path to the config file.")

	flag.Parse()

	// Load config.
	var cfg config.Config

	// Set defaults.
	cfg.SetDefaults()

	err := cfg.LoadFromFs(cfgPath)

	if err != nil {
		fmt.Println("Failed to load config file.")
		fmt.Println(err)

		os.Exit(1)
	}

	// Setup and load web server.
	err = web.SetupServer(&cfg)

	if err != nil {
		fmt.Println("Failed to setup and load web server.")
		fmt.Println(err)

		os.Exit(1)
	}
}
