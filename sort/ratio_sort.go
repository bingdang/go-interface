package sort

import (
	"E-commerceRecommendation/common"
	"sort"
)

type RatioSorter struct {
	Tag string
}

func (self RatioSorter) Name() string {
	return self.Tag
}

// 按照好评率从大到小排序
func (self RatioSorter) Sort(products []*common.Product) []*common.Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].PositiveRatio > products[j].PositiveRatio
	})
	return products
}
