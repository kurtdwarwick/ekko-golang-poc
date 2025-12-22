package commands

import "github.com/ekko-earth/shared/messaging"

type OnboardOrganisationCommand struct {
	messaging.Command

	LegalName   string
	TradingName string
	Website     *string
}
