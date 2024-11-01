package i18n

import "golang.org/x/text/language"

type I18nProvider interface {
	T(key string, l language.Tag) string
}

type I18n struct{}

func (i *I18n) T(key string, lang language.Tag) string {
	// Default return key
	v := key

	return v
}

func NewI18n(opts ...*I18nConfigOptionsBuilder) (I18nProvider, error) {
	i := &I18n{}
	args, err := NewOptions(opts...)
	if err != nil {
		return nil, err
	}

	return i, nil
}
