package main

import (
	"E-commerceRecommendation/common"
	"E-commerceRecommendation/filter"
	"E-commerceRecommendation/recall"
	"E-commerceRecommendation/sort"
	"fmt"
	"log"
	"time"
)

/*
//电商推荐引擎流程
整体流程：
召回商品 -> 排序 -> 过滤

1.对召回、排序、过滤分别定义接口

2.将三个步骤的接口封装到推荐引擎结构体中

3.实现推荐框架

4.对召回、排序、过滤进行具体的实现

5.使用具体的实现，将具体的实现赋值给Recommender结构体
*/

// 推荐引擎结构体
type Recommender struct {
	//可能有多路召回，这里用切片
	Recallers []recall.Recaller
	//排序，根据好评率等进行排序
	Sorter sort.Sorter
	//过滤
	Filters []filter.Filter
}

// 3.推荐框架的实现
func (rec *Recommender) Rec() []*common.Product {
	//遍历多个召回，根据不同召回规则将所有召回的商品进行集合。用map去重
	RecallMap := make(map[int]*common.Product, 100)
	//1.顺序遍历每一路召回
	for _, recaller := range rec.Recallers {
		begin := time.Now()
		products := recaller.Recall(10)
		log.Printf("召回%s耗时%dms，召回了%d个商品\n", recaller.Name(), time.Since(begin).Milliseconds(), len(products))

		for _, product := range products {
			RecallMap[product.Id] = product
		}
	}
	log.Printf("去重之后一共召回了%d个商品\n", len(RecallMap))

	//1.1将召回的商品从map转切片
	RecallSlice := make([]*common.Product, 0, len(RecallMap))

	for _, product := range RecallMap {
		RecallSlice = append(RecallSlice, product)
	}

	//2.对召回结果进行排序
	begin := time.Now()
	SortedResult := rec.Sorter.Sort(RecallSlice)
	log.Printf("排序耗时%dms", time.Since(begin).Milliseconds())

	//3.顺序执行多个过滤规则
	FilteredResult := SortedResult

	//3.1统计过滤了多少个商品
	prevCount := len(FilteredResult)
	for _, filter := range rec.Filters {
		begin := time.Now()
		FilteredResult = filter.Filter(FilteredResult)
		log.Printf("过滤规则%s耗时%dms，过滤掉了%d个商品", filter.Name(), time.Since(begin).Milliseconds(), prevCount-len(FilteredResult))
		prevCount = len(FilteredResult)
	}

	return FilteredResult
}

func main() {
	//使用三步的具体的实现
	rec := Recommender{
		Recallers: []recall.Recaller{
			recall.HotRecall{Tag: "hot"},
			recall.SizeRecall{Tag: "szie"},
		},
		Sorter: sort.RatioSorter{Tag: "ratio"},
		Filters: []filter.Filter{
			filter.RatioFilter{Tag: "ratio"},
		},
	}

	products := rec.Rec()
	for i, product := range products {
		fmt.Printf("No.%d,Id:%d,Name:%s\n", i, product.Id, product.Name)
	}
}
