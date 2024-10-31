package i18n

import (
	"os"
	"path/filepath"

	"golang.org/x/text/language"
)

type I18nConfigOptions struct {
	Dir       string
	Fallback  language.Tag
	Namespace string
}

type I18nConfigOptionsBuilder struct {
	Opts []func(*I18nConfigOptions) error
}

func I18nConfig() *I18nConfigOptionsBuilder {

	i := &I18nConfigOptionsBuilder{}

	i.Opts = append(i.Opts, func(args *I18nConfigOptions) error {
		execFilename, err := os.Executable()
		if err != nil {
			return err
		}
		// Set default dir to translation folder
		defaultDir := filepath.Dir(execFilename)

		args.Dir = filepath.Join(defaultDir, "i18n")

		return nil
	})

	return i

}

func (i *I18nConfigOptionsBuilder) SetDir(dir string) *I18nConfigOptionsBuilder {

	i.Opts = append(i.Opts, func(args *I18nConfigOptions) error {
		args.Dir = dir
		return nil
	})

	return i
}
