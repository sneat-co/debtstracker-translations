package trans

import (
	"errors"
	"fmt"
	"github.com/strongo/i18n"
	"strings"
)

// DebtsTrackerLocales - defines locales specific to DebtsTracker app
type DebtsTrackerLocales struct {
}

var ErrUnsupportedLocale = errors.New("unsupported locale")

// GetLocaleByCode5 - get locale by code
func (DebtsTrackerLocales) GetLocaleByCode5(code5 string) (locale i18n.Locale, err error) {
	var ok bool
	if locale, ok = SupportedLocalesByCode5[code5]; ok {
		return locale, nil
	}
	code2 := code5[:2]
	for _, l := range SupportedLocales {
		if strings.HasPrefix(l.Code5, code2) {
			return l, nil
		}
	}
	return locale, fmt.Errorf("%w: %s", ErrUnsupportedLocale, code5)
}

// SupportedLocales - supported locales
var SupportedLocales []i18n.Locale = []i18n.Locale{
	i18n.LocaleEnUK,
	i18n.LocaleDeDe,
	i18n.LocaleEsEs,
	i18n.LocaleFaIr,
	i18n.LocaleFrFr,
	i18n.LocaleJaJp,
	i18n.LocaleItIt,
	i18n.LocaleRuRu,
	i18n.LocaleUaUa,
	i18n.LocaleZhCn,
	i18n.LocalePtPt,
	i18n.LocalePtBr,
}

// SupportedLocalesByCode5 - supported locales by code 5
var SupportedLocalesByCode5 = make(map[string]i18n.Locale, len(SupportedLocales))

func init() {
	for _, locale := range SupportedLocales {
		SupportedLocalesByCode5[locale.Code5] = locale
	}
}

// ChooseLocaleIcon - locale icons
var ChooseLocaleIcon string = strings.Join([]string{
	i18n.LocaleEnUS.FlagIcon,
	i18n.LocaleRuRu.FlagIcon,
	i18n.LocaleItIt.FlagIcon,
	//i18n.LocaleIdId.FlagIcon,
	i18n.LocaleFaIr.FlagIcon,
}, " ")
