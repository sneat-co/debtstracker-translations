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
	if locale, ok = SupportedLocalesByCode5[code5]; !ok {
		err = fmt.Errorf("%w: %s", ErrUnsupportedLocale, code5)
	}
	return locale, err
}

// SupportedLocalesByCode5 - supported locales by code 5
var SupportedLocales []i18n.Locale = []i18n.Locale{
	i18n.LocaleDeDe,
	i18n.LocaleEnUS,
	i18n.LocaleEsEs,
	i18n.LocaleFaIr,
	i18n.LocaleItIt,
	i18n.LocaleRuRu,
	//i18n.LocaleIdId,
	//LocaleEsEs,
	//LocaleFrFr,
	//LocalePtPt,
	//LocalePtBr,
}

// SupportedLocales  - supported locales
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
