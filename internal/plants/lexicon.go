package plants

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Info struct {
	Name           string   `yaml:"name"`
	Height         uint     `yaml:"height"`
	Width          uint     `yaml:"width"`
	Family         string   `yaml:"family"`
	FeederType     string   `yaml:"feederType"`
	RootType       string   `yaml:"rootType"`
	GoodCompanions []string `yaml:"goodCompanions"`
	BadCompanions  []string `yaml:"badCompanions"`
}

type Lexicon struct {
	Entries []Info `yaml:"plants"`
}

//go:embed plants.yml
var rawPlantInfo string
var EmbeddedLexicon *Lexicon = &Lexicon{}

func init() {
	if err := yaml.Unmarshal([]byte(rawPlantInfo), EmbeddedLexicon); err != nil {
		panic(fmt.Errorf("failed to parse plant infos: %w", err))
	}

	// Make the raw data reclaimable by the GC.
	rawPlantInfo = ""
}
