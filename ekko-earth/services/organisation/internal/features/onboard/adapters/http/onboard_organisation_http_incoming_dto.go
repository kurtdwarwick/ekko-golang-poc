package http

type OnboardOrganisationHttpIncomingDto struct {
	LegalName   string `json:"legalName"`
	TradingName string `json:"tradingName"`
	Website     string `json:"website"`
}
