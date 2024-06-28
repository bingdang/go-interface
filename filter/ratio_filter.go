package filter

import "E-commerceRecommendation/common"

type RatioFilter struct {
	Tag string
}

func (self RatioFilter) Name() string {
	return self.Tag
}

func (self RatioFilter) Filter(products []*common.Product) []*common.Product {
	rect := make([]*common.Product, 0, len(products))
	for _, product := range products {
		//评论大于10 && 好评率大于0.8
		if product.RatioCount > 10 && product.PositiveRatio > 0.8 {
			rect = append(rect, product)
		}
	}
	return rect
}
