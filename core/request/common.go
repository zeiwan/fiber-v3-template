package request

// PageReq 分页请求参数
type PageReq struct {
	PageNo   int `form:"page_no,default=1" json:"page_no,default=1" binding:"omitempty,gte=1"`             // 页码
	PageSize int `form:"page_size,default=10" json:"page_size,default=10" binding:"omitempty,gt=0,lte=40"` // 每页大小
}

// CommonId 通用ID
type CommonId struct {
	Id int32 `form:"id" json:"id" binding:"required"`
}
