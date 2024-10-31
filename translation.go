package i18n

type I18nProvider interface {
}

type I18n struct{}

func NewI18n(opts ...*I18nConfigOptionsBuilder) I18nProvider {
	i := &I18n{}

	for _, otp := range opts {

	}

	return i
}
