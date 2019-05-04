package tools

import (
	"errors"
	"net/url"
	"strings"

	"github.com/go-ozzo/ozzo-validation"
)

type requiredRule struct {
	message string
	skipNil bool
}

var Required = &requiredRule{message: "cannot be blank", skipNil: false}

type URL struct {
	URL      *url.URL
	Str      string
	basePath string
}

const slash = "/"

func (j *URL) SetBasePath(path string) {
	j.basePath = path
}

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
	return validation.Validate(j.Str, validation.Required)
}

// Validate checks if the given value is valid or not.
func (v *requiredRule) Validate(value interface{}) error {
	j, ok := value.(URL)
	if !ok {
		return errors.New("invalid type")
	}
	return j.Validate()
}

// Error sets the error message for the rule.
func (v *requiredRule) Error(message string) *requiredRule {
	return &requiredRule{
		message: message,
		skipNil: v.skipNil,
	}
}
