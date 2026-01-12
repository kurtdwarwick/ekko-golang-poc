package http

type GetOrganisationByIdHttpOutgoingDto struct {
	Id          string  `json:"id"`
	LegalName   string  `json:"legalName"`
	TradingName string  `json:"tradingName"`
	Website     *string `json:"website"`
}
