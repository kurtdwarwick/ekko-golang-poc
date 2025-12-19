package policies

import (
	"errors"
	"organisation/internal/data/entities"
)

type LegalNameValidationPolicy struct{}

func (policy LegalNameValidationPolicy) Apply(value any) error {
	organisation, ok := value.(entities.Organisation)

	if !ok {
		return errors.New("value is not an organisation")
	}

	if organisation.LegalName == "" {
		return errors.New("organisation legal name is required")
	}

	return nil
}
