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
var SupportedLocalesByCode5 = map[string]strongo.Locale{
	strongo.LOCALE_EN_US: strongo.LocaleEnUS,
	//strongo.LOCALE_EN_UK: strongo.LocaleEnUK,
	strongo.LOCALE_RU_RU: strongo.LocaleRuRu,
	//strongo.LOCALE_IT_IT: strongo.LocaleItIt,
	//strongo.LOCALE_ID_ID: strongo.LocaleIdId,
	//strongo.LOCALE_DE_DE: strongo.LocaleDeDe,
	//strongo.LOCALE_ES_ES: strongo.LocaleEsEs,
	//strongo.LOCALE_FR_FR: strongo.LocaleFrFr,
	//strongo.LOCALE_PT_PT: strongo.LocalePtPt,
	//strongo.LOCALE_PT_BR: strongo.LocalePtBr,
	//strongo.LOCALE_FA_IR: strongo.LocaleFaIr,
}

// SupportedLocales  - supported locales
var SupportedLocales []strongo.Locale = []strongo.Locale{
	strongo.LocaleEnUS,
	strongo.LocaleRuRu,
	//strongo.LocaleIdId,
	//strongo.LocaleFaIr,
	//strongo.LocaleItIt,
	//LocaleDeDe,
	//LocaleEsEs,
	//LocaleFrFr,
	//LocalePtPt,
	//LocalePtBr,
}

// ChooseLocaleIcon - locale icons
var ChooseLocaleIcon = strings.Join([]string{
	strongo.LocaleEnUS.FlagIcon,
	strongo.LocaleRuRu.FlagIcon,
	//strongo.LocaleIdId.FlagIcon,
	//strongo.LocaleItIt.FlagIcon,
	//strongo.LocaleFaIr.FlagIcon,
}, " ")
