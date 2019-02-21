package foo

const (
	MAX          = 100 //公開
	internal_max = 1   //非公開
)

var (
	MIN          = -100 //公開
	internal_min = -1   //非公開
)

func FooFunc(n int) int {
	// 公開
	return internalFunc(n)
}

func internalFunc(n int) int {
	// 非公開
	return n + 1
}
