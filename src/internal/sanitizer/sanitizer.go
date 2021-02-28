package sanitizer

import "github.com/microcosm-cc/bluemonday"

func Sanitize(data string) string {
	sanitizer := bluemonday.UGCPolicy()
	result := sanitizer.Sanitize(data)

	return result
}
