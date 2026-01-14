package policies_test

import (
	"testing"

	"github.com/ekko-earth/organisation/internal/features/onboard/core"
	"github.com/ekko-earth/organisation/internal/features/onboard/core/policies"
)

// Unit tests should be focused on predictable behaviour (business logic) and not on implementation details, and should be focused
// on quality of tests, rather than quantity.

func TestLegalNameValidationPolicy_Apply(t *testing.T) {
	policy := policies.LegalNameValidationPolicy{}

	t.Run("should return an error if the organisation legal name is empty", func(t *testing.T) {
		organisation := core.Organisation{
			LegalName: "",
		}

		err := policy.Apply(organisation)

		if err == nil {
			t.Errorf("expected an error")
		}
	})

	t.Run(
		"should not return an error if the organisation legal name is not empty",
		func(t *testing.T) {
			organisation := core.Organisation{
				LegalName: "Test Organisation",
			}

			err := policy.Apply(organisation)

			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		},
	)
}
