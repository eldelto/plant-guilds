package plants_test

import (
	"os"
	"sort"
	"testing"

	"github.com/eldelto/plant-guilds/internal/plants"
	"gopkg.in/yaml.v3"
)

func TestLexiconConsistency(t *testing.T) {
	if sanitizeLexicon(plants.EmbeddedLexicon) {
		file, err := os.Create("plants-corrected.yml")
		if err != nil {
			t.Fatalf("failed to create plants-corrected.yml: %s", err)
		}
		defer file.Close()

		entries := plants.EmbeddedLexicon.Entries
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Name < entries[j].Name
		})
		if err := yaml.NewEncoder(file).Encode(plants.EmbeddedLexicon); err != nil {
			t.Fatalf("failed to serialize to plants-corrected.yml: %s", err)
		}
		t.Fatal("lexicon is not consistent - compare content of plants-corrected.yml with original content")
	}
}

func hasEntry(l *plants.Lexicon, plantName string) bool {
	for _, entry := range l.Entries {
		if entry.Name == plantName {
			return true
		}
	}

	return false
}

func addMissingEntries(l *plants.Lexicon, plantNames []string) (modified bool) {
	for _, plantName := range plantNames {
		if !hasEntry(l, plantName) {
			l.Entries = append(l.Entries, plants.Info{Name: plantName})
			modified = true
		}
	}

	return modified
}

func sanitizeLexicon(l *plants.Lexicon) (modified bool) {
	for _, entry := range l.Entries {
		modified = addMissingEntries(l, entry.GoodCompanions) || modified
		modified = addMissingEntries(l, entry.BadCompanions) || modified
	}

	return modified
}
