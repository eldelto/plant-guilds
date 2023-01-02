package api

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"

	"github.com/eldelto/plant-guilds/internal/plants"
)

//go:embed assets
var AssetsFS embed.FS

//go:embed templates
var TemplatesFS embed.FS

type TemplateData struct{}

func (td *TemplateData) PlantGuildsJson() (string, error) {
	buffer := bytes.Buffer{}
	if err := json.NewEncoder(&buffer).Encode(plants.Guilds()); err != nil {
		return "", fmt.Errorf("failed to encode plant guilds: %w", err)
	}

	return buffer.String(), nil
}
