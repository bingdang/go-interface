package recall

import "E-commerceRecommendation/common"

// 1.定义召回接口
type Recaller interface {
	Recall(n int) []*common.Product
	Name() string
}
