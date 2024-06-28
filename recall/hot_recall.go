package recall

import (
	"E-commerceRecommendation/common"
	"sort"
)

// 测试商品,假装从数据库中全部取出了商品
var allProducts = []*common.Product{
	{Id: 1, Sale: 39, Name: "p1", RatioCount: 99, PositiveRatio: 0.9},
	{Id: 2, Sale: 399, Name: "p2", RatioCount: 992, PositiveRatio: 0.9},
	{Id: 3, Sale: 1239, Name: "p3", RatioCount: 9922, PositiveRatio: 0.9},
	{Id: 4, Sale: 3932, Name: "p4", RatioCount: 99222, PositiveRatio: 0.7},
	{Id: 5, Sale: 3349, Name: "p5", RatioCount: 9922, PositiveRatio: 0.5},
	{Id: 6, Sale: 6739, Name: "p6", RatioCount: 99223422, PositiveRatio: 0.5},
	{Id: 7, Sale: 36459, Name: "p7", RatioCount: 5623, PositiveRatio: 1},
	{Id: 8, Sale: 3239, Name: "p8", RatioCount: 564923, PositiveRatio: 0.8},
}

// 4.1召回的具体实现
type HotRecall struct {
	Tag string
}

func (r HotRecall) Name() string {
	return r.Tag
}

func (r HotRecall) Recall(n int) []*common.Product {
	//对allProducts进行了原地排序
	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].Sale > allProducts[j].Sale
		// 将切片中元素的销量属性从大到小进行排序
	})

	rect := make([]*common.Product, 0, n)

	for _, product := range allProducts {
		rect = append(rect, product)

		//只需召回前5
		if len(rect) >= 5 {
			break
		}
	}

	return rect
}
