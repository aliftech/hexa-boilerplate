package i18n

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Bundle map[string]interface{}

var bundles = make(map[string]Bundle)

func LoadLocales(localesDir string) {
	entries, err := os.ReadDir(localesDir)
	if err != nil {
		log.Fatal("❌ Failed to read locales dir:", err)
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".toml" {
			lang := filepath.Base(entry.Name()[:len(entry.Name())-5]) // remove .toml
			data, err := os.ReadFile(filepath.Join(localesDir, entry.Name()))
			if err != nil {
				log.Printf("⚠️ Failed to read %s: %v", entry.Name(), err)
				continue
			}

			var bundle Bundle
			if err := toml.Unmarshal(data, &bundle); err != nil {
				log.Printf("⚠️ Failed to parse %s: %v", entry.Name(), err)
				continue
			}

			bundles[lang] = bundle
			log.Printf("✅ Loaded locale: %s", lang)
		}
	}
}

// Get returns localized string, e.g., Get("en", "user.created")
func Get(lang, key string) string {
	bundle, ok := bundles[lang]
	if !ok {
		lang = "en" // fallback
		bundle = bundles["en"]
	}

	keys := splitKey(key)
	value := traverse(bundle, keys)
	if value == "" {
		return key // fallback to key
	}
	return value
}

func splitKey(key string) []string {
	// "user.created" → ["user", "created"]
	return strings.Split(key, ".")
}

func traverse(bundle Bundle, keys []string) string {
	current := interface{}(bundle)
	for _, k := range keys {
		if m, ok := current.(map[string]interface{}); ok {
			if v, exists := m[k]; exists {
				if str, isStr := v.(string); isStr {
					if len(keys) == 1 {
						return str
					}
				} else {
					current = v
					continue
				}
			}
		}
		return ""
	}
	if str, ok := current.(string); ok {
		return str
	}
	return ""
}
