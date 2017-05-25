package trans

import (
	"fmt"
	"github.com/strongo/app"
	"strings"
)

// DebtsTrackerLocales - defines locales specific to DebtsTracker app
type DebtsTrackerLocales struct {
}

// GetLocaleByCode5 - get locale by code
func (DebtsTrackerLocales) GetLocaleByCode5(code5 string) (locale strongo.Locale, err error) {
	var ok bool
	if locale, ok = SupportedLocalesByCode5[code5]; !ok {
		err = fmt.Errorf("Locale not found by code5: %v", code5)
	}
	return locale, err
}

// SupportedLocalesByCode5 - supported locales by code 5
var SupportedLocales []strongo.Locale = []strongo.Locale{
	strongo.LocaleEnUS,
	strongo.LocaleRuRu,
	//strongo.LocaleIdId,
	//strongo.LocaleFaIr,
	strongo.LocaleItIt,
	//LocaleDeDe,
	//LocaleEsEs,
	//LocaleFrFr,
	//LocalePtPt,
	//LocalePtBr,
}

// SupportedLocales  - supported locales
var SupportedLocalesByCode5 = make(map[string]strongo.Locale, len(SupportedLocales))

func init() {
	for _, locale := range SupportedLocales {
		SupportedLocalesByCode5[locale.Code5] = locale
	}
}

// ChooseLocaleIcon - locale icons
var ChooseLocaleIcon = strings.Join([]string{
	strongo.LocaleEnUS.FlagIcon,
	strongo.LocaleRuRu.FlagIcon,
	strongo.LocaleItIt.FlagIcon,
	//strongo.LocaleIdId.FlagIcon,
	//strongo.LocaleFaIr.FlagIcon,
}, " ")
