package response

type WithStatusCode interface {
	StatusCode() int
}

type EncapsulatedError interface {
	EncapsulateError() map[string]any
}
