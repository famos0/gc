package gclib

type Options struct {
	Quiet    bool
	Testless bool
}

func NewOptions() *Options {
	return &Options{}
}
