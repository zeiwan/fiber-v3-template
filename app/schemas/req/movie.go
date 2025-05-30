package req

type MovieHotListReq struct {
	CityId string `query:"cityId" default:"859" ` // Ensures it's never empty
}

//validate:"required"
