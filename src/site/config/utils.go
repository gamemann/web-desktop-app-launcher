package config

import "fmt"

func (cfg *Config) Print() {
	fmt.Println("Web Settings")
	fmt.Printf("\tHost => %s\n", cfg.Web.Host)
	fmt.Printf("\tPort => %d\n", cfg.Web.Port)

	if len(cfg.Apps) > 0 {
		fmt.Println("Applications")
		for i, app := range cfg.Apps {
			fmt.Printf("\tApplication #%d\n", i+1)
			fmt.Printf("\t\tName => %s\n", app.Name)
			fmt.Printf("\t\tStart Command => %s\n", app.Start)
			fmt.Printf("\t\tStop Command => %s\n", app.Stop)
			fmt.Printf("\t\tBanner => %s\n", app.Banner)
		}
	}
}
