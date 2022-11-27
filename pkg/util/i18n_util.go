package util

import "github.com/nicksnyder/go-i18n/v2/i18n"

func Localize(localizer *i18n.Localizer, id string) string {
	localizedTxt := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: id,
		},
	})
	return localizedTxt
}
