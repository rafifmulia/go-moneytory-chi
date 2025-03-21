package helper

func InlinePointer[T any](v T) *T {
	return &v
}
