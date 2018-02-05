package brands

import (
	"github.com/whosonfirst/go-whosonfirst-flags"
)

type Brand interface {
	Id() int64
	Name() string
	Size() string
	String() string
	IsCurrent() (flags.ExistentialFlag, error)
	IsCeased() (flags.ExistentialFlag, error)
	IsDeprecated() (flags.ExistentialFlag, error)
	IsSuperseded() (flags.ExistentialFlag, error)
	IsSuperseding() (flags.ExistentialFlag, error)
	SupersededBy() []int64
	Supersedes() []int64
}
