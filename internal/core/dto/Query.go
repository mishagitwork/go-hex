package dto

// QueryContract ...
type Query struct {
	Page     int
	Limit    int
	OrderBy  string
	OrderDir string
	Search   string
}

type MetaData struct {
	Count int64
}
type List[T any] struct {
	Data []T
	Meta MetaData
}
