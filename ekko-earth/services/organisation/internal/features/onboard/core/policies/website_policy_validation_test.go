package policies_test

import (
	"testing"

	"github.com/ekko-earth/organisation/internal/features/onboard/core"
	"github.com/ekko-earth/organisation/internal/features/onboard/core/policies"
)

// Unit tests should be focused on predictable behaviour (business logic) and not on implementation details, and should be focused
// on quality of tests, rather than quantity.

func TestWebsiteValidationPolicy_Apply(t *testing.T) {
	policy := policies.WebsiteValidationPolicy{}

	t.Run(
		"should return an error if the organisation website is not a valid URL",
		func(t *testing.T) {
			website := "some not valid URL"

			organisation := core.Organisation{
				Website: &website,
			}

			err := policy.Apply(organisation)

			if err == nil {
				t.Errorf("expected an error")
			}
		},
	)

	t.Run(
		"should not return an error if the organisation website is a valid URL",
		func(t *testing.T) {
			website := "https://www.example.com"

			organisation := core.Organisation{
				Website: &website,
			}

			err := policy.Apply(organisation)

			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		},
	)

	t.Run("should not return an error if the organisation website is nil", func(t *testing.T) {
		organisation := core.Organisation{
			Website: nil,
		}

		err := policy.Apply(organisation)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}
