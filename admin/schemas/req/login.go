package req

type LoginReq struct {
	Account  string `json:"account"  validate:"required"`
	Password string `json:"password" validate:"required"`
	Terminal uint8  `json:"terminal" validate:"oneof=1 2 3"`
}
