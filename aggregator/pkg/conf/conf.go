package conf

import (
	"encoding/json"
	"os"
)

// ConJson Конфигурация приложения.
type ConJson struct {
	Period  int      `json:"request_period"`
	LinkArr []string `json:"rss"`
}

func NewConfig() []ConJson {
	// loadConfiguration Чтение и раскодированние файла конфигурации.
	bytes, err := os.ReadFile("C:/Users/kagir/GolandProjects/Task36a.4.1Aggregator/cmd/server/config.json")
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
