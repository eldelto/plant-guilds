package plant_test

import (
	"os"
	"sort"
	"testing"

	"github.com/eldelto/plant-guilds/internal/plant"
	"gopkg.in/yaml.v3"
)

func TestLexiconConsistency(t *testing.T) {
	if sanitizeLexicon(plant.EmbeddedLexicon) {
		file, err := os.Create("plants-corrected.yml")
		if err != nil {
			t.Fatalf("failed to create plants-corrected.yml: %s", err)
		}
		defer file.Close()

		listing := plant.EmbeddedLexicon.Listing()
		entries := listing.Entries
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Name < entries[j].Name
		})
		if err := yaml.NewEncoder(file).Encode(listing); err != nil {
			t.Fatalf("failed to serialize to plants-corrected.yml: %s", err)
		}
		t.Fatal("lexicon is not consistent - compare content of plants-corrected.yml with original content")
	}
}

func hasEntry(l *plant.Lexicon, plantName string) bool {
	_, ok := l.Entries[plantName]
	return ok
}

func addMissingEntries(l *plant.Lexicon, plantNames []string) (modified bool) {
	for _, plantName := range plantNames {
		if !hasEntry(l, plantName) {
			l.Entries[plantName] = plant.Info{Name: plantName}
			modified = true
		}
	}

	return modified
}

func sanitizeEntries(l *plant.Lexicon) (modified bool) {
	for _, entry := range l.Entries {
		modified = addMissingEntries(l, entry.GoodCompanions) || modified
		modified = addMissingEntries(l, entry.BadCompanions) || modified
	}

	return modified
}

func appendCompanion(companions *[]string, plantName string) bool {
	for _, companion := range *companions {
		if companion == plantName {
			return false
		}
	}
	*companions = append(*companions, plantName)
	return true
}

func addMissingCompanions(l *plant.Lexicon, entry *plant.Info) (modified bool) {
	for _, companion := range entry.GoodCompanions {
		companionEntry := l.Entries[companion]
		modified = appendCompanion(&companionEntry.GoodCompanions, entry.Name) || modified
		l.Entries[companion] = companionEntry
	}

	for _, companion := range entry.BadCompanions {
		companionEntry := l.Entries[companion]
		modified = appendCompanion(&companionEntry.BadCompanions, entry.Name) || modified
		l.Entries[companion] = companionEntry
	}

	return modified
}

func sanitizeCompanions(l *plant.Lexicon) (modified bool) {
	for _, entry := range l.Entries {
		modified = addMissingCompanions(l, &entry) || modified
	}

	return modified
}

func sanitizeLexicon(l *plant.Lexicon) (modified bool) {
	modified = sanitizeEntries(l) || modified
	modified = sanitizeCompanions(l) || modified

	return modified
}
