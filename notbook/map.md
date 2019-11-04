# golang源码阅读之map

## 1、map的结构
```
type hmap struct {
	count     int // 元素个数
	flags     uint8
	B         uint8  // 2^B表示buckets这个数组的大小
	noverflow uint16 // 
	hash0     uint32 // 哈希种子
	// buckets是一个bucket的数组，每一个bucket中存放的是一个bmap，其中买个bmap中可以存8个key-value
	buckets unsafe.Pointer 
	// 扩容的时候使用，表示的是老的那个buckets
	oldbuckets unsafe.Pointer 
	nevacuate  uintptr        
	// 将包含指针的信息保存起来,使得gc在扫描的时候不会去扫描整个map
	extra *mapextra // optional fields
}

```
结构中 buckets是一个数组