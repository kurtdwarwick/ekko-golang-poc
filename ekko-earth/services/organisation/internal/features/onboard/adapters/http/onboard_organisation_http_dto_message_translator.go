package http

import (
	"github.com/ekko-earth/organisation/internal/features/onboard/core/commands"
	"github.com/ekko-earth/shared/messaging"
)

type OnboardOrganisationHttpDtoMessageTranslator struct {
	MessageTranslator messaging.MessageTranslator[OnboardOrganisationHttpDto, commands.OnboardOrganisationCommand]
}

func (translator *OnboardOrganisationHttpDtoMessageTranslator) Translate(
	message OnboardOrganisationHttpDto,
) (commands.OnboardOrganisationCommand, error) {
	return commands.OnboardOrganisationCommand{
		LegalName:   message.LegalName,
		TradingName: message.TradingName,
		Website:     &message.Website,
	}, nil
}
