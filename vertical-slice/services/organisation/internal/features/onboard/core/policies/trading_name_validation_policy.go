package policies

import (
	"errors"
	"organisation/internal/data/entities"
)

type TradingNameValidationPolicy struct{}

func (policy TradingNameValidationPolicy) Apply(value any) error {
	organisation, ok := value.(entities.Organisation)

	if !ok {
		return errors.New("value is not an organisation")
	}

	if organisation.TradingName == "" {
		return errors.New("organisation trading name is required")
	}

	return nil
}
