package i18n

import (
	"os"
	"path/filepath"

	"golang.org/x/text/language"
)

type I18nConfigOptions struct {
	Dir              string
	FileType         FileType
	Fallback         language.Tag
	UseNamespace     bool
	DefaultNamespace string
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

		// Set default file type
		args.FileType = FileTypeJSON

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

func (i *I18nConfigOptionsBuilder) SetFileType(fileType FileType) *I18nConfigOptionsBuilder {

	i.Opts = append(i.Opts, func(args *I18nConfigOptions) error {
		args.FileType = fileType
		return nil
	})

	return i
}

func (i *I18nConfigOptionsBuilder) SetFallback(fallback language.Tag) *I18nConfigOptionsBuilder {

	i.Opts = append(i.Opts, func(args *I18nConfigOptions) error {
		args.Fallback = fallback
		return nil
	})

	return i
}

func (i *I18nConfigOptionsBuilder) SetUseNamespace(useNamespace bool) *I18nConfigOptionsBuilder {

	i.Opts = append(i.Opts, func(args *I18nConfigOptions) error {
		args.UseNamespace = useNamespace
		return nil
	})

	return i
}

func (i *I18nConfigOptionsBuilder) SetDefaultNamespace(defaultNamespace string) *I18nConfigOptionsBuilder {

	i.Opts = append(i.Opts, func(args *I18nConfigOptions) error {
		args.DefaultNamespace = defaultNamespace
		return nil
	})

	return i
}

func (i *I18nConfigOptionsBuilder) SetUseNamespaceWithDefaultNamespace(useNamespace bool, defaultNamespace string) *I18nConfigOptionsBuilder {

	i.Opts = append(i.Opts, func(args *I18nConfigOptions) error {
		args.UseNamespace = useNamespace
		args.DefaultNamespace = defaultNamespace
		return nil
	})

	return i
}

func (i *I18nConfigOptionsBuilder) List() []func(*I18nConfigOptions) error {
	return i.Opts
}
