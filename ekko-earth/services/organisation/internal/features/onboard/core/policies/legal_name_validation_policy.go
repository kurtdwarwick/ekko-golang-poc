package policies

import (
	"errors"

	"github.com/ekko-earth/organisation/internal/features/onboard/core"
)

type LegalNameValidationPolicy struct{}

func (policy LegalNameValidationPolicy) Apply(value any) error {
	organisation, ok := value.(core.Organisation)

	if !ok {
		return errors.New("value is not an organisation")
	}

	if organisation.LegalName == "" {
		return errors.New("organisation legal name is required")
	}

	return nil
}
