package queries

type GetOrganisationsQuery struct {
	Page *int32
	Size *int32
}

func NewGetOrganisationsQuery(page *int32, size *int32) *GetOrganisationsQuery {
	return &GetOrganisationsQuery{
		Page: page,
		Size: size,
	}
}
