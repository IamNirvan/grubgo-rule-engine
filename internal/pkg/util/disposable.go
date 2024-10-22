package util

import "context"

type Disposable interface {
	Dispose(ctx context.Context) error
}
