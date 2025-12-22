package http

type OnboardOrganisationHttpDto struct {
	LegalName   string `json:"legal_name"`
	TradingName string `json:"trading_name"`
	Website     string `json:"website"`
}
