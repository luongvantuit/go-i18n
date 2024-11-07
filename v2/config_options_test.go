package i18n_test

import (
	"testing"

	"github.com/zeroxsolutions/go-i18n/v2"
)

func TestI18nConfig_SetDir(t *testing.T) {

	dir := "./translations"

	options := i18n.I18nConfig().SetDir(dir)

	args, err := i18n.NewOptions(options)

	if err != nil {
		t.FailNow()
	}

	if args.Dir != dir {
		t.FailNow()
	}

}

func TestI18nConfig_SetFileType(t *testing.T) {

	options := i18n.I18nConfig().SetFileType(i18n.FileTypeYAML)

	args, err := i18n.NewOptions(options)

	if err != nil {
		t.FailNow()
	}

	if args.FileType != i18n.FileTypeYAML {
		t.FailNow()
	}

}
