package i18n_test

import (
	"errors"
	"testing"

	"github.com/zeroxsolutions/go-i18n"
)

func TestParse_FileType(t *testing.T) {

	r0, err := i18n.ParseFileType("JSON")

	if err != nil {
		t.FailNow()
	}

	if r0 != i18n.FileTypeJSON {
		t.FailNow()
	}

	r1, err := i18n.ParseFileType("yAml")

	if err != nil {
		t.FailNow()
	}

	if r1 != i18n.FileTypeYAML {
		t.FailNow()
	}

	r2, err := i18n.ParseFileType("yml")

	if err != nil {
		t.FailNow()
	}

	if r2 != i18n.FileTypeYAML {
		t.FailNow()
	}

	r3, err := i18n.ParseFileType("*")

	if err == nil {
		t.FailNow()
	}

	if !errors.Is(err, i18n.ErrInvalidFileType) {
		t.FailNow()
	}

	if r3 != "" {
		t.FailNow()
	}

}
