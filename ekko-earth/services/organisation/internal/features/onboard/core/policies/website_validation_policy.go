package policies

import (
	"errors"

	"regexp"

	"github.com/ekko-earth/organisation/internal/features/onboard/core/data/entities"
)

type WebsiteValidationPolicy struct{}

func (policy WebsiteValidationPolicy) Apply(value any) error {
	organisation, ok := value.(entities.Organisation)

	if !ok {
		return errors.New("value is not an organisation")
	}

	if organisation.Website != nil {
		regex, _ := regexp.Compile(
			`^(https?://)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)$`,
		)

		if !regex.MatchString(*organisation.Website) {
			return errors.New("organisation website is not a valid URL")
		}
	}

	return nil
}
