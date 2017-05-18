package trans

import (
	"regexp"
	"testing"
)

var (
	supportedLocales = []string{"ru-RU", "en-US", "fa-IR", "pl-PL", "pt-PT", "es-ES", "fr-FR", "it-IT", "ja-JP", "zh-CN", "de-DE", "ko-KO"}
	reVars = regexp.MustCompile(`%(v|d)|\{\{\..+?}}`)
	reWords = regexp.MustCompile(`\w+|%(?:v|d)`)
)

func TestTRANS(t *testing.T) {
	var wordsCount int
	for key, vals := range TRANS {
		countsByLang := make(map[string]map[string]int)
		for lang, val := range vals {
			if !isSupportedLang(lang) {
				t.Errorf("Key %v has unsupported language: %v", key, lang)
				continue
			}
			vars := reVars.FindAllString(val, -1)
			counts := make(map[string]int, len(vars))
			for _, v := range vars {
				counts[v] += 1
			}
			countsByLang[lang] = counts
		}
		enCounts, ok := countsByLang["en-US"]
		if !ok {
			t.Errorf("Key %v missing en-US trnaslation", key)
			continue
		}
		wordsCount += len(reWords.FindAllString(vals["en-US"], -1))
		reported := make(map[string]int)
		for lang, counts := range countsByLang {
			if lang == "en-US" {
				continue
			}
			for enV, enCount := range enCounts {
				if counts[enV] != enCount {
					t.Errorf("%v:%v has %d of '%v' while en-US has %d", key, lang, counts[enV], enV, enCount)
					reported[enV] = enCount
				}
			}

			for langV, langCount := range counts {
				if enCounts[langV] != langCount {
					if _, ok := reported[langV]; !ok {
						t.Errorf("%v:%v has %d of '%v' while en-US has %d", key, lang, langCount, langV, enCounts[langV])
					}
				}
			}
		}
		if _, ok := countsByLang["ru-RU"]; !ok {
			t.Errorf("%v: missing translation for ru-RU", key)
		}
	}
	t.Logf("English words count: %d", wordsCount)
}

func isSupportedLang(l string) bool {
	for _, supportedLang := range supportedLocales {
		if l == supportedLang {
			return true
		}
	}
	return false
}
