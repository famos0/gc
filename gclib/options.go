package gclib

// Options struct centralize options in the tool
type Options struct {
	Quiet    bool
	Testless bool
	Stdin    bool
}

// NewOptions init Options struct
func NewOptions() *Options {
	return &Options{}
}
