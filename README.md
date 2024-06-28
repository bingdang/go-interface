## 电商推荐引擎流程
### 整体流程
![rec](https://cakepanit.com/images/pasted-333.png)

1. 对召回、排序、过滤分别定义接口
   - 召回接口：`recall/recaller.go`
   - 排序接口：`sort/sorter.go`
   - 过滤接口：`filter/filter.go`
2. 将三个步骤的接口封装到推荐引擎结构体中
   - `rec.go/Recommender 结构体`
3. 实现纯接口的推荐框架
   - `rec.go/Rec() Recommender的方法`
4. 对召回、排序、过滤进行具体的实现
   - 召回实现：
     - 按照热度召回：`recall/hot_recall.go`
     - 按照size召回：`recall/size_recall.go`
   - 排序实现：
     - 按照好评率排序：`sort/ratio_sort.go`
     - 按照size排序：`sort/szie_sort.go`
   - 过滤的具体实现：
     - 按照评价进行过滤：`filter/ratio_filter.go`
5. 使用具体的实现，将具体的实现赋值给Recommender结构体
   - `rec.go/main()`

```go
~/go/src/go-interface main !1 ?1 ❯ go run ./
2024/06/28 18:03:05 召回hot耗时0ms，召回了5个商品
2024/06/28 18:03:05 召回szie耗时0ms，召回了8个商品
2024/06/28 18:03:05 去重之后一共召回了8个商品
2024/06/28 18:03:05 排序耗时0ms
2024/06/28 18:03:05 过滤规则ratio耗时0ms，过滤掉了4个商品
No.0,Id:7,Name:p7
No.1,Id:1,Name:p1
No.2,Id:2,Name:p2
No.3,Id:3,Name:p3
```