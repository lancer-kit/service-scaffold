package tools

import (
	"net/url"
	"strings"

	"github.com/go-ozzo/ozzo-validation"
)

type URL struct {
	URL      *url.URL
	Str      string
	basePath string
}

const slash = "/"

func (j *URL) WithPath(path string) string {
	ur := *j.URL
	ur.Path = j.basePath + slash + strings.TrimPrefix(path, slash)

	return ur.String()
}

func (j *URL) WithPathURL(path string) url.URL {
	ur := *j.URL
	ur.Path = j.basePath + slash + strings.TrimPrefix(path, slash)

	return ur
}

func (j *URL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}
	j.Str = s

	u, err := url.Parse(s)
	j.URL = u
	j.basePath = strings.TrimSuffix(u.Path, slash)
	return err
}

func (j *URL) Validate() error {
	return validation.ValidateStruct(j,
		validation.Field(&j.Str, validation.Required))
}
