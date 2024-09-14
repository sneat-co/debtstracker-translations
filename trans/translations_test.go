package trans

import (
	"fmt"
	"github.com/strongo/i18n"
	"regexp"
	"strings"
	"testing"
)

var (
	requiredLocales = []string{
		i18n.LocaleCodeEnUK,
		i18n.LocaleCodeRuRU,
	}
	supportedLocales = append(requiredLocales,
		i18n.LocaleCodeEsES,
		i18n.LocaleCodeItIT,
		i18n.LocaleCodeFaIR,
		i18n.LocaleCodeEnUS,
		i18n.LocaleCodePlPL,
		i18n.LocaleCodePtPT,
		i18n.LocaleCodeFrFR,
		i18n.LocaleCodeJaJP,
		i18n.LocaleCodeZhCN,
		i18n.LocaleCodeKoKO,
		i18n.LocaleCodeDeDE,
		i18n.LocaleCodeTrTR,
		i18n.LocaleCodeIdID,
		i18n.LocaleCodeUaUA,
		i18n.LocaleCodePtBR,
	)
	//desiredLocales   = []string{"it-IT", "fa-IR", "es-ES"}
	reVars  = regexp.MustCompile(`%[vd]|\{\{\..+?}}`)
	reWords = regexp.MustCompile(`\w+|%[vd]`)
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
		enCounts, ok := countsByLang[i18n.LocaleCodeEnUK]
		if !ok {
			t.Errorf("Key %v missing en-UK trnaslation", key)
			continue
		}
		if len(missingLocales) > 0 {
			//desiredMisses := make([]string, 0, len(missingLocales))
			requiredMisses := make([]string, 0, len(missingLocales))
			if _, ok := requiredMissCount[key]; ok {
				requiredMisses = append(requiredMisses, key)
				requiredMissCount[key] += 1
				//} else {
				//desiredMisses = append(desiredMisses, key)
			}
			//if len(desiredMisses) > 0 {
			//	t.Logf("Key `%v` is missing optional translations for: %v", key, missingLocales)
			//}
			if len(requiredMisses) > 0 {
				t.Errorf("Key `%v` is missing required translations for: %v", key, requiredMisses)
			}
		}
		wordsCount += len(reWords.FindAllString(vals[i18n.LocaleCodeEnUK], -1))
		reported := make(map[string]int)
		for lang, counts := range countsByLang {
			if lang == i18n.LocaleCodeEnUK {
				continue
			}
			for enV, enCount := range enCounts {
				if c := counts[enV]; c != enCount {
					if !(c == 0 && vals[lang] == "") {
						t.Errorf("%v:%v has %d of '%v' while en-US has %d", key, lang, counts[enV], enV, enCount)
						reported[enV] = enCount
					}
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

func TestHtmlTags(t *testing.T) {
	for code, vals := range TRANS {
		for lang, val := range vals {
			for _, tag := range []string{"b", "i"} {
				openTag, closeTag := fmt.Sprintf("<%v>", tag), fmt.Sprintf("</%v>", tag)
				if openCount, closeCount := strings.Count(val, openTag), strings.Count(val, closeTag); openCount != closeCount {
					t.Errorf("%v[%v]: %v != %v: %v != %v: %v", code, lang, openTag, closeTag, openCount, closeCount, val)
				}
			}
			if openCount, closeCount := strings.Count(val, "<a"), strings.Count(val, "</a>"); openCount != closeCount {
				t.Errorf("%v[%v]: %v != %v: %v != %v: %v", code, lang, "<a>", "</a>", openCount, closeCount, val)
			}
		}
	}
}

func isSupportedLang(l string) bool {
	for _, supportedLang := range supportedLocales {
		if l == supportedLang {
			return true
		}
	}
	return false
}
