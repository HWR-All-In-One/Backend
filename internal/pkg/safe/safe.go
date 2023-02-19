package safe

import "os"

type Safe struct{}

func New() *Safe {
	return &Safe{}
}

func (s Safe) Get() string {
	return os.Getenv("HWR_AIO_KEY")
}

type Getter interface {
	Get() string
}
