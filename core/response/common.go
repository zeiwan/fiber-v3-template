package response

import "github.com/spf13/cast"

// PageResp 分页返回值
type PageResp struct {
	Count    int64 `json:"count"`            // 总数
	PageNo   int   `json:"page_no"`          // 页No
	PageSize int   `json:"page_size"`        // 每页Size
	Lists    any   `json:"lists"`            // 数据
	Extend   any   `json:"extend,omitempty"` // 扩展字段
}

// Pages 计算总页数
func Pages(count int64, limit int) int64 {
	// 计算商
	quotient := count / cast.ToInt64(limit)

	// 如果count不能被limit整除，则总页数应该是商加1
	if count%cast.ToInt64(limit) != 0 {
		quotient++
	}
	return quotient
}

// ExportPage 导出分页数据
type ExportPage struct {
	Count      int64  `json:"count"`
	PageSize   int    `json:"page_size"`
	SumPage    int64  `json:"sum_page"`
	MaxPage    int    `json:"max_page"`
	AllMaxSize int    `json:"all_max_size"`
	PageStart  int    `json:"page_start"`
	PageEnd    int    `json:"page_end"`
	FileName   string `json:"file_name"`
}
