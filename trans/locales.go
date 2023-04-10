package trans

import (
	"errors"
	"fmt"
	strongo "github.com/strongo/app"
	"strings"
)

// DebtsTrackerLocales - defines locales specific to DebtsTracker app
type DebtsTrackerLocales struct {
}

var ErrUnsupportedLocale = errors.New("unsupported locale")

// GetLocaleByCode5 - get locale by code
func (DebtsTrackerLocales) GetLocaleByCode5(code5 string) (locale strongo.Locale, err error) {
	var ok bool
	if locale, ok = SupportedLocalesByCode5[code5]; !ok {
		err = fmt.Errorf("%w: %s", ErrUnsupportedLocale, code5)
	}
	return locale, err
}

// SupportedLocalesByCode5 - supported locales by code 5
var SupportedLocales []strongo.Locale = []strongo.Locale{
	strongo.LocaleDeDe,
	strongo.LocaleEnUS,
	strongo.LocaleEsEs,
	strongo.LocaleFaIr,
	strongo.LocaleItIt,
	strongo.LocaleRuRu,
	//strongo.LocaleIdId,
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
var ChooseLocaleIcon string = strings.Join([]string{
	strongo.LocaleEnUS.FlagIcon,
	strongo.LocaleRuRu.FlagIcon,
	strongo.LocaleItIt.FlagIcon,
	//strongo.LocaleIdId.FlagIcon,
	strongo.LocaleFaIr.FlagIcon,
}, " ")
