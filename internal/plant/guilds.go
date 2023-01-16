package plant

type Relations struct {
	Good []string
	Bad  []string
}

type GuildMapping map[string]Relations

var plantGuilds GuildMapping = GuildMapping{
	"Aubergine": {
		Good: []string{"Weiße Bohnen"},
		Bad:  []string{"Paprika", "Tomaten"},
	},
	"Blaukraut": {
		Good: []string{"Boretsch", "Buschbohnen", "Erbsen", "Möhren", "Phacelia", "Salat", "Sellerie", "Spinat"},
		Bad:  []string{"Knoblauch", "Kohl", "Tomaten", "Zwiebel"},
	},
	"Blumenkohl": {
		Good: []string{"Buschbohnen", "Erbsen", "Phacelia", "Sellerie"},
		Bad:  []string{"Knoblauch", "Zwiebel"},
	},
	"Bohnenkraut": {
		Good: []string{"Bohnene", "Rote Bete", "Salat"},
		Bad:  []string{""},
	},
	"Borretsch": {
		Good: []string{"Blaukraut", "Bohnen", "Erdbeeren", "Erbsen", "Kohlrabi", "Kohlarten"},
		Bad:  []string{""},
	},
	"Buschbohnen": {
		Good: []string{"Bohnenkraut", "Borretsch", "Chinakohl", "Dill", "Erdbeeren", "Gurken", "Kapuzinerkresse", "Kartoffeln", "Kohlarten", "Kohlrabi", "Radieschen", "Rettich", "Rote Bete", "Salat", "Salbei", "Sellerie", "Spinat", "Tomaten"},
		Bad:  []string{"Erbsen", "Fenchel", "Knoblauch", "Paprika", "Porree", "Schnittlauch", "Stangenbohnen", "Zwiebeln"},
	},
	"Chinakohl": {
		Good: []string{"Bohnen", "Erbsen", "Spinat", "Salat"},
		Bad:  []string{"Radieschen", "Rettich"},
	},
	"Dicke Bohnen": {
		Good: []string{"Kartoffeln", "Schwarzwurzel", "Spinat", "Kapuzinerkresse"},
		Bad:  []string{""},
	},
	"Dill": {
		Good: []string{"Erbsen", "Möhren", "Gurken", "Kohlarten", "Rote Bete", "Salat", "Zwiebel"},
		Bad:  []string{""},
	},
	"Endivien": {
		Good: []string{"Fenchel", "Kohlarten", "Porree", "Stangenbohnen"},
		Bad:  []string{""},
	},
	"Erbsen": {
		Good: []string{"Borretsch", "Dill", "Fenchel", "Gurken", "Kohlarten", "Kohlrabi", "Kopfsalat", "Mais", "Möhren", "Radieschen", "Rettich", "Sellerie", "Spinat", "Zucchini"},
		Bad:  []string{"Bohnen", "Kartoffeln", "Knoblauch", "Porree", "Tomaten", "Zwiebel"},
	},
	"Erdbeeren": {
		Good: []string{"Borretsch", "Buschbohnen", "Knoblauch", "Kopfsalat", "Porree", "Radieschen", "Ringelblume", "Spinat"},
		Bad:  []string{"Kohlarten"},
	},
	"Feldsalat": {
		Good: []string{"Erdbeeren", "Radieschen"},
		Bad:  []string{""},
	},
	"Fenchel": {
		Good: []string{"Endivie", "Erbse", "Feldsalat", "Gurke", "Salat", "Sellerie"},
		Bad:  []string{"Bohnen", "Tomate", "Kohlrabi"},
	},
	"Gurken": {
		Good: []string{"Basilikum", "Bohnen", "Dill", "Erbsen", "Fenchel", "Kohlarten", "Kopfsalat", "Kümmel", "Mais", "Porree", "Rote Bete", "Sellerie", "Zwiebeln"},
		Bad:  []string{"Tomaten", "Kartoffeln", "Radieschen", "Rettich"},
	},
	"Kartoffeln": {
		Good: []string{"Buschbohnen", "Dicke Bohnen", "Kapuzinerkresse", "Kohlrabi", "Kümmel", "Mais", "Meerrettich", "Pfefferminze", "Spinat", "Tagetes"},
		Bad:  []string{"Erbsen", "Gurken", "Kürbis", "Melone", "Rote Bete", "Sellerie", "Sonnenblume", "Tomaten", "Zwiebel"},
	},
	"Knoblauch": {
		Good: []string{"Erdbeeren", "Gurken", "Himbeeren", "Lilien", "Möhren", "Rosen", "Rote Bete", "Tomaten"},
		Bad:  []string{"Erbsen", "Buschbohnen", "Kohlarten", "Stangenbohne"},
	},
	"Knollensellerie": {
		Good: []string{"Bohnen", "Erbsen", "Gurken", "Kohl", "Kohlrabi", "Porree", "Spinat", "Tomate"},
		Bad:  []string{"Kartoffeln", "Mais", "Salat"},
	},
	"Kohl": {
		Good: []string{"Bohnen", "Borretsch", "Dill", "Endivien", "Erbsen", "Gurken", "Rote Bete", "Salat", "Sellerie", "Spinat", "Tagetes", "Tomaten"},
		Bad:  []string{"Kohlarten", "Kartoffeln", "Knoblauch", "Kohlrabi", "Rhabarber", "Schnittlauch", "Zwiebel"},
	},
	"Kohlrabi": {
		Good: []string{"Bohnen", "Borretsch", "Dill", "Erbsen", "Erdbeeren", "Gurken", "Kartoffeln", "Porree", "Radieschen", "Rote Bete", "Salat", "Schwarzwurzel", "Sellerie", "Spargel", "Spinat"},
		Bad:  []string{"Kohl"},
	},
	"Kopfsalat": {
		Good: []string{"Bohnen", "Chicorée", "Erbsen", "Fenchel", "Gurken", "Kohlarten", "Kohlrabi", "Möhren", "Porree", "Radieschen", "Rettich", "Schwarzwurzel", "Tomaten", "Zwiebel"},
		Bad:  []string{"Kresse", "Petersilie", "Sellerie"},
	},
	"Kürbis": {
		Good: []string{"Basilikum", "Bohnen", "Dill", "Erbsen", "Fenchel", "Kohlarten", "Kopfsalat", "Kümmel", "Mais", "Porree", "Rote Bete", "Sellerie", "Zwiebeln"},
		Bad:  []string{"Tomaten", "Kartoffeln", "Radieschen", "Rettich"},
	},
	"Mairübchen": {
		Good: []string{"Mangold", "Salat", "Spinat", "Dill", "Erbsen", "Bohnen"},
		Bad:  []string{"Weiß- und Rotkohl", "Steckrüben", "Rettich", "Senf", "Raps"},
	},
	"Mais": {
		Good: []string{"Bohnen", "Gurken", "Kartoffeln", "Kopfsalat", "Kürbis", "Melonen", "Tomaten", "Zucchini"},
		Bad:  []string{"Rote Bete", "Sellerie"},
	},
	"Mangold": {
		Good: []string{"Buschbohnen", "Kohlarten", "Möhren", "Radieschen", "Rettich", "Salat"},
		Bad:  []string{"Rote Beete"},
	},
	"Merrettich": {
		Good: []string{"Kartoffeln"},
		Bad:  []string{""},
	},
	"Melonen": {
		Good: []string{"Basilikum", "Bohnen", "Dill", "Erbsen", "Fenchel", "Kohlarten", "Kopfsalat", "Kümmel", "Mais", "Porree", "Rote Bete", "Sellerie", "Zwiebeln"},
		Bad:  []string{"Tomaten", "Kartoffeln", "Radieschen", "Rettich"},
	},
	"Möhren": {
		Good: []string{"Chicorée", "Dill", "Erbsen", "Knoblauch", "Mangold", "Porree", "Radieschen", "Rettich", "Salat", "Schwarzwurzeln", "Spinat", "Tomaten", "Zwiebeln"},
		Bad:  []string{"Rote Bete", "Pfefferminze"},
	},
	"Paprika": {
		Good: []string{"Kohlarten", "Möhren", "Tomaten"},
		Bad:  []string{"Erbsen", "Fenchel", "Rote Bete"},
	},
	"Pastinaken": {
		Good: []string{"Möhren", "Kartoffeln", "Kopfsalat", "Pflücksalat", "Radieschen", "Rote Bete", "Sellerie", "Spinat"},
		Bad:  []string{""},
	},
	"Petersilie": {
		Good: []string{"Gurken", "Radieschen", "Tomaten", "Zwiebeln"},
		Bad:  []string{"Salate"},
	},
	"Pflücksalat": {
		Good: []string{"Buschbohnen", "Fenchel", "Kohlarten", "Radieschen", "Rote Bete", "Stangen Bohnen", "Pastinake"},
		Bad:  []string{""},
	},
	"Porree": {
		Good: []string{"Endivien", "Erdbeeren", "Kohlarten", "Knoblauch", "Möhren", "Petersilie", "Salat", "Schwarzwurzeln", "Sellerie", "Spinat", "Tomate"},
		Bad:  []string{"Bohnen", "Erbsen", "Rote Bete", "Stangenbohnen"},
	},
	"Radieschen": {
		Good: []string{"Bohnen", "Erbsen", "Kapuzinnerkresse", "Kohl", "Mangold", "Möhren", "Petersilie", "Salat", "Spinat", "Tomate"},
		Bad:  []string{"Gurken", "Chinakohl", "Kürbis", "Melone"},
	},
	"Rettich": {
		Good: []string{"Bohnen", "Erbsen", "Kapuzinnerkresse", "Kohl", "Mangold", "Möhren", "Petersilie", "Salat", "Spinat", "Tomate"},
		Bad:  []string{"Gurken", "Chinakohl", "Kürbis", "Melone"},
	},
	"Rhabarber": {
		Good: []string{"Buschbohnen", "Erbsen", "Kohlarten", "Salat", "Spinat "},
		Bad:  []string{""},
	},
	"Ringelblume": {
		Good: []string{"Erdbeeren", "Gurken", "Kohlarten", "Salate", "Tomaten"},
		Bad:  []string{""},
	},
	"Rote Bete": {
		Good: []string{"Bohnen", "Dill", "Gurken", "Kohl", "Kohlrabi", "Salat", "Zucchini", "Zwiebeln"},
		Bad:  []string{"Kartoffeln", "Mangold", "Porree", "Spinat"},
	},
	"Schwarzwurzel": {
		Good: []string{"Bohnen", "Kohlrabi", "Porree", "Salat"},
		Bad:  []string{""},
	},
	"Sellerie": {
		Good: []string{"Buschbohnen", "Chinakohl", "Fenchel", "Gurken", "Kamille", "Kohl", "Kohlrabi", "Pastinake", "Porree", "Salat", "Spinat", "Tomaten"},
		Bad:  []string{"Erbsen", "Kartoffeln"},
	},
	"Sonnenblumen": {
		Good: []string{"Gurken"},
		Bad:  []string{"Kartoffeln"},
	},
	"Spargel": {
		Good: []string{"Dill", "Gurken", "Petersilie", "Kohlrabi", "Salat", "Tomaten"},
		Bad:  []string{""},
	},
	"Spinat": {
		Good: []string{"Erdbeeren", "Kartoffeln", "Kohlarten", "Kohlrabi", "Radieschen", "Rettich", "Rhabarber", "Stangenbohnen", "Tomate"},
		Bad:  []string{"Rote Bete", "Mangold"},
	},
	"Stangenbohnen": {
		Good: []string{"Gurken", "Kapuzinerkresse", "Kartoffeln", "Kohlarten", "Kohlrabi", "Radieschen", "Rettich", "Rote Bete", "Salat", "Salbei", "Sellerie", "Spinat"},
		Bad:  []string{"Buschbohnen", "Erbsen", "Fenchel", "Knoblauch", "Paprika", "Porree", "Schnittlauch", "Zwiebeln"},
	},
	"Tomaten": {
		Good: []string{"Buschbohnen", "Chicorée", "Knoblauch", "Kohlrabi", "Möhren", "Pastinake", "Petersilie", "Porree", "Radieschen", "Ringelblumen", "Salat", "Sellerie", "Spinat", "Zwiebeln"},
		Bad:  []string{"Blaukraut", "Erbsen", "Fenchel", "Gurken", "Kartoffeln", "Rote Bete", "Kürbis", "Melone"},
	},
	"Zucchini": {
		Good: []string{"Basilikum", "Kapuzinerkresse", "Stangenbohnen", "Zwiebeln"},
		Bad:  []string{"Gurken", "Kürbis", "Melone"},
	},
	"Zwiebeln": {
		Good: []string{"Dill", "Bohnenkraut", "Gurken", "Kamille", "Knoblauch", "Möhren", "Pastinake", "Rote Bete", "Salat", "Tomaten", "Zucchini"},
		Bad:  []string{"Bohnen", "Erbsen", "Kartoffeln", "Kohlarten", "Porree"},
	},
}

func Guilds() GuildMapping {
	// TODO: Validate?
	return plantGuilds
}
