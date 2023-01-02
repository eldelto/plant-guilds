package plants

type Relations struct {
	Good []string
	Bad  []string
}

type GuildMapping map[string]Relations

var plantGuilds GuildMapping = GuildMapping{
	"Knoblauch": {
		Good: []string{"Erbsen", "Erdbeere"},
		Bad:  []string{"Fire", "Storm"},
	},
	"Erbsen": {
		Good: []string{"Knoblauch", "Erdbeere"},
		Bad:  []string{"Fire"},
	},
	"Erdbeere": {
		Good: []string{""},
		Bad:  []string{"Fire"},
	},
	"Fire": {
		Good: []string{""},
		Bad:  []string{""},
	},
}

func Guilds() GuildMapping {
	// TODO: Validate?
	return plantGuilds
}
