package common

// 商品对象
type Product struct {
	Id            int
	Name          string
	Size          int    //尺寸大小
	Sale          int    //销量
	ShipAddress   string //发货地
	Price         float64
	PositiveRatio float64 //好评率
	RatioCount    int     //评论量
}
