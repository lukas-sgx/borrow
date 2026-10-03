package utils

type Runner struct {
	err error
}

func NewRunner() *Runner {
	return &Runner{
		err: nil,
	}
}

func (r *Runner) Run(task func() error) *Runner {
	if r.Err() == nil {
		r.err = task()
	}
	return r
}

func (r *Runner) Err() error {
	return r.err
}
