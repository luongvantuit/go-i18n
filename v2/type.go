package i18n

import "strings"

type FileType string

const (
	FileTypeJSON FileType = "json"
	FileTypeYAML FileType = "yaml"
)

func ParseFileType(fileType string) (FileType, error) {

	switch strings.ToLower(strings.Trim(fileType, " ")) {
	case "yaml", "yml":
		return FileTypeYAML, nil
	case "json":
		return FileTypeJSON, nil
	}

	return "", ErrInvalidFileType
}
