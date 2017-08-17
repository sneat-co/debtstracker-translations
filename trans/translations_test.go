package trans

import (
	"regexp"
	"testing"
	"strings"
)

var (
	supportedLocales = []string{"ru-RU", "en-US", "fa-IR", "pl-PL", "pt-PT", "es-ES", "fr-FR", "it-IT", "ja-JP", "zh-CN", "de-DE", "ko-KO"}
	requiredLocales = []string{"en-US", "es-ES", "ru-RU", "it-IT", "fa-IR"}
	desiredLocales = []string{"it-IT", "fa-IR", "es-ES"}
	reVars = regexp.MustCompile(`%(v|d)|\{\{\..+?}}`)
	reWords = regexp.MustCompile(`\w+|%(?:v|d)`)
)

func TestTRANS(t *testing.T) {
	var wordsCount int

	requiredMissCount := make(map[string]int, len(requiredLocales))

	for _, requiredLocale := range requiredLocales {
		requiredMissCount[requiredLocale] = 0
	}

	for key, vals := range TRANS {
		countsByLang := make(map[string]map[string]int)
		missingLocales := append([]string{}, requiredLocales...)
		for lang, val := range vals {
			if !isSupportedLang(lang) {
				t.Errorf("Key %v has unsupported language: %v", key, lang)
				continue
			}
			for i, ml := range missingLocales {
				if ml == lang {
					// Delete without preserving order
					// https://github.com/golang/go/wiki/SliceTricks#delete-without-preserving-order
					l := len(missingLocales)
					missingLocales[i] = missingLocales[l-1]
					missingLocales = missingLocales[:l-1]
				}
			}
			if strings.Contains(val, "https: ") || strings.Contains(val, "http: ") {
				t.Logf("Invalid http(s): link: %v=%v", key, val)
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
		if len(missingLocales) > 0 {
			desiredMisses := make([]string, 0, len(missingLocales))
			requiredMisses := make([]string, 0, len(missingLocales))
			if _, ok := requiredMissCount[key]; ok {
				requiredMisses = append(requiredMisses, key)
				requiredMissCount[key] += 1
			} else {
				desiredMisses = append(desiredMisses, key)
			}
			if len(desiredMisses) > 0 {
				t.Logf("Key `%v` is missing optional translations for: %v", key, missingLocales)
			}
			if len(requiredMisses) > 0 {
				t.Errorf("Key `%v` is missing required translations for: %v", key, requiredMisses)
			}
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
