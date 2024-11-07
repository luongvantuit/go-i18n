package i18n

import (
	"reflect"

	"golang.org/x/text/language"
)

type I18nProvider interface {
	T(key string, l language.Tag) string
}

type I18n struct {
	Dir              string
	FileType         FileType
	Fallback         language.Tag
	UseNamespace     bool
	DefaultNamespace string
}

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

	if args.Dir != "" {
		i.Dir = args.Dir
	}

	i.FileType = FileTypeJSON
	if !reflect.ValueOf(args.FileType).IsNil() {
		i.FileType = args.FileType
	}

	if args.UseNamespace {
		i.UseNamespace = true
	}

	if args.DefaultNamespace != "" {
		i.DefaultNamespace = args.DefaultNamespace
	}

	return i, nil
}
