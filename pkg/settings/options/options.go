package options

type Options struct {
}

type CreateOptions struct {
}

type UpdateOptions struct {
}

type DeleteOptions struct {
}

type PageOptions struct {
	PageNum  int
	PageSize int
}

func (opts PageOptions) Limit() int {
	return opts.PageNum
}

func (opts PageOptions) Start() int {
	return (opts.PageNum - 1) * opts.PageSize
}
