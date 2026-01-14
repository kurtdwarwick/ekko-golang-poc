package policies_test

import (
	"testing"

	"github.com/ekko-earth/organisation/internal/features/onboard/core"
	"github.com/ekko-earth/organisation/internal/features/onboard/core/policies"
)

// Unit tests should be focused on predictable behaviour (business logic) and not on implementation details, and should be focused
// on quality of tests, rather than quantity.

func TestTradingNameValidationPolicy_Apply(t *testing.T) {
	policy := policies.TradingNameValidationPolicy{}

	t.Run("should return an error if the organisation trading name is empty", func(t *testing.T) {
		organisation := core.Organisation{
			TradingName: "",
		}

		err := policy.Apply(organisation)

		if err == nil {
			t.Errorf("expected an error")
		}
	})

	t.Run(
		"should not return an error if the organisation trading name is not empty",
		func(t *testing.T) {
			organisation := core.Organisation{
				TradingName: "Test Organisation",
			}

			err := policy.Apply(organisation)

			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		},
	)
}
