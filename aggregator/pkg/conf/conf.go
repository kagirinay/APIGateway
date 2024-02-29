package conf

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// ConJson Конфигурация приложения.
type ConJson struct {
	Period  int      `json:"request_period"`
	LinkArr []string `json:"rss"`
}

func NewConfig() []ConJson {
	filename := "config.json"
	ext := filepath.Join(filename)
	// loadConfiguration Чтение и раскодированние файла конфигурации.
	bytes, err := os.ReadFile(ext)
	if err != nil {
		panic(err.Error())
	}
	var conf []ConJson
	json.Unmarshal(bytes, &conf)
	if err != nil {
		panic(err.Error())
	}

	return conf
}
