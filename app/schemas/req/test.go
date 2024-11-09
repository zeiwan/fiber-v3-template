package req

type TestGetReq struct {
	Account string `form:"account"`
	Id      int    `form:"id" validate:"required"`
	Ids     int    `form:"ids" validate:"required"`
}
