package sort

import (
	"E-commerceRecommendation/common"
	"sort"
)

type SizeSorter struct {
	Tag string
}

func (self SizeSorter) Name() string {
	return self.Tag
}

// 按照size从大到小排序
func (self SizeSorter) Sort(products []*common.Product) []*common.Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Size > products[j].Size
	})
	return products
}
