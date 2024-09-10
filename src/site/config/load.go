package config

import (
	"encoding/json"
	"os"
)

func (cfg *Config) LoadFromFs(path string) error {
	var err error

	f, err := os.Open(path)

	if err != nil {
		return err
	}

	defer f.Close()

	stat, err := f.Stat()

	if err != nil {
		return err
	}

	data := make([]byte, stat.Size())

	_, err = f.Read(data)

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), cfg)

	return err
}
