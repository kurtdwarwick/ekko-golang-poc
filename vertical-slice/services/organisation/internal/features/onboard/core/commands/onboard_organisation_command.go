package commands

import "commands"

type OnboardOrganisationCommand struct {
	commands.Command

	LegalName   string
	TradingName string
	Website     *string
}
