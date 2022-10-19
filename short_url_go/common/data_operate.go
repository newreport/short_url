package common

type MapKY interface {
	string | int
}

// @Title GetMapAllKeys
func GetMapAllKeys[T MapKY](m map[T]T) []T {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func GetMapAllValues[T MapKY](m map[T]T) []T {
	values := make([]T, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
