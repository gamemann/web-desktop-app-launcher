package config

type Web struct {
	Host string `json:"host"`
	Port int    `json:"port"`

	LogToFile    bool   `json:"log_to_file"`
	LogDirectory string `json:"log_directory"`

	Env map[string]string `json:"env"`
}

type App struct {
	Name   string            `json:"name"`
	Start  string            `json:"start"`
	Stop   string            `json:"stop"`
	Banner string            `json:"banner"`
	Env    map[string]string `json:"env"`
}

type Config struct {
	Web  Web   `json:"web"`
	Apps []App `json:"applications"`
}
