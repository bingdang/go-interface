package sort

import "E-commerceRecommendation/common"

// 2.定义排序接口
type Sorter interface {
	Sort([]*common.Product) []*common.Product
	Name() string
}
