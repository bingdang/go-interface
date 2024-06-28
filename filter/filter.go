package filter

import "E-commerceRecommendation/common"

// 3.定义过滤接口
type Filter interface {
	Filter([]*common.Product) []*common.Product
	Name() string
}
