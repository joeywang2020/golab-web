package util

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	i18n_resource_folder = "../../resources/i18n/"
	i18n_bundle_format   = "toml"
)

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc(i18n_bundle_format, toml.Unmarshal)
	// No need to load active.en.toml since we are providing default translations.
	// bundle.MustLoadMessageFile("active.en.toml")
	bundle.LoadMessageFile(i18n_resource_folder + "active.en.toml")
	bundle.LoadMessageFile(i18n_resource_folder + "active.es.toml")
	// bundle.MustLoadMessageFile("../../resources/i18n/active.zh-CN.toml")
	bundle.LoadMessageFile(i18n_resource_folder + "active.zh.toml")
}

func SimpleLocalize(lang, id string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	localizedTxt, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: id,
		},
	})
	_ = fmt.Errorf("Localize lang %v, id:%v, err: %+v", lang, id, err)
	return localizedTxt
}

func Localize(lang string, lc *i18n.LocalizeConfig) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	localizedTxt, err := localizer.Localize(lc)

	_ = fmt.Errorf("Localize lang %v, id:%v, err: %+v", lang, lc.MessageID, err)
	return localizedTxt
}

func MustLocalize(lang, id string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	localizedTxt := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: id,
		},
	})
	return localizedTxt
}
