package main

import (
	"log"
	"testing"

	"github.com/ivpusic/toml"
	"github.com/nerfmiester/sourceGo/Config"
)

func TestName(t *testing.T) {

	name := getName()

	if name != "World!" {
		t.Error("Response form getName is an unexpected value")
	}
}

func TestToml(t *testing.T) {

	var tomlConfig Config.TomlConfig

	if _, err := toml.DecodeFile("definition.toml", &tomlConfig); err != nil {
		log.Fatal(err)
		// return
	}

	owner := tomlConfig.Owner.Name
	if owner != "Tom Preston-Werner" {
		t.Error("Owner name not as definition")
	}
}
