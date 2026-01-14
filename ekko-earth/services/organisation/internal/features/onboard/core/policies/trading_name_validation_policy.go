package policies

import (
	"errors"

	"github.com/ekko-earth/organisation/internal/features/onboard/core"
)

type TradingNameValidationPolicy struct{}

func (policy TradingNameValidationPolicy) Apply(value any) error {
	organisation, ok := value.(core.Organisation)

	if !ok {
		return errors.New("value is not an organisation")
	}

	if organisation.TradingName == "" {
		return errors.New("organisation trading name is required")
	}

	return nil
}
