package test

type IConfigService interface {
	GetConfig() (any, error)
}

func NewConfigService() IConfigService {
	return &configService{}
}

type configService struct {
	//dao *dao.Query
}

func (c configService) GetConfig() (any, error) {
	//return "nil", response.QueryError
	return "nil", nil
}
