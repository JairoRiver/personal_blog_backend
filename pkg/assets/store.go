package assets

import (
	"context"
)

type ImageStorer interface {
	UploadImage(ctx context.Context, file []byte, path string, name string) (string, error)
	DeleteImage(ctx context.Context, path string, name string) error
}
