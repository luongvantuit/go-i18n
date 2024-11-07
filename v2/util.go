package i18n

import "reflect"

func NewOptions(opts ...*I18nConfigOptionsBuilder) (*I18nConfigOptions, error) {
	args := new(I18nConfigOptions)

	for _, opt := range opts {
		if opt == nil || reflect.ValueOf(opt).IsNil() {
			continue
		}

		for _, setArgs := range opt.List() {
			if setArgs == nil {
				continue
			}
			if err := setArgs(args); err != nil {
				return nil, err
			}
		}
	}

	return args, nil
}
