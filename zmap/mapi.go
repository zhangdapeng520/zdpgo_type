package zmap

// MapI Map类型的接口
type MapI[K comparable, V any] interface {
	Setter[K, V]
	Getter[K, V]
	Loader[K, V]
	Clear()
	Len() int
	Range(func(k K, v V) bool)
	Has(k K) bool
	Keys() []K
	Values() []V
	Merge(MapI[K, V])
	Equal(MapI[K, V]) bool
	Delete(k K)
}

// Setter 设置值的接口
type Setter[K comparable, V any] interface {
	Set(K, V)
}

// Getter 获取值的接口
type Getter[K comparable, V any] interface {
	Get(k K) (v V)
}

// Loader 加载值的接口
type Loader[K comparable, V any] interface {
	Load(k K) (v V, ok bool)
}
