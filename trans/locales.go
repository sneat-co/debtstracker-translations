package trans

import (
	"errors"
	"fmt"
	"github.com/strongo/bots-framework/core"
	"strings"
)

type DebtsTrackerLocales struct {
}

func (DebtsTrackerLocales) GetLocaleByCode5(code5 string) (locale bots.Locale, err error) {
	var ok bool
	if locale, ok = SupportedLocalesByCode5[code5]; !ok {
		err = errors.New(fmt.Sprintf("Locale not found by code5: %v", code5))
	}
	return locale, err
}

var SupportedLocalesByCode5 = map[string]bots.Locale{
	bots.LOCALE_EN_US: bots.LocaleEnUs,
	bots.LOCALE_RU_RU: bots.LocaleRuRu,
	bots.LOCALE_IT_IT: bots.LocaleItIt,
	bots.LOCALE_ID_ID: bots.LocaleIdId,
	bots.LOCALE_DE_DE: bots.LocaleDeDe,
	bots.LOCALE_ES_ES: bots.LocaleEsEs,
	bots.LOCALE_FR_FR: bots.LocaleFrFr,
	bots.LOCALE_PT_PT: bots.LocalePtPt,
	bots.LOCALE_PT_BR: bots.LocalePtBr,
	bots.LOCALE_FA_IR: bots.LocaleFaIr,
}

var SupportedLocales []bots.Locale = []bots.Locale{
	bots.LocaleEnUs,
	bots.LocaleRuRu,
	bots.LocaleIdId,
	bots.LocaleFaIr,
	bots.LocaleItIt,
	//LocaleDeDe,
	//LocaleEsEs,
	//LocaleFrFr,
	//LocalePtPt,
	//LocalePtBr,
}

var ChooseLocaleIcon = strings.Join([]string{
	bots.LocaleEnUs.FlagIcon,
	bots.LocaleRuRu.FlagIcon,
	bots.LocaleIdId.FlagIcon,
	bots.LocaleItIt.FlagIcon,
	bots.LocaleFaIr.FlagIcon,
}, " ")
