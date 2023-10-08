package assets

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"os"
)

type MemoryStore struct{}

func NewMemoryAssetStore(symmetricKey string) (ImageStorer, error) {

	maker := &MemoryStore{}

	return maker, nil
}

func (ms *MemoryStore) UploadImage(ctx context.Context, file []byte, path string, name string) (string, error) {
	// Save the file to specific dst.
	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		return "", err
	}
	directory := path + name
	out, err := os.Create(directory)
	if err != nil {
		return "", err
	}
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 1

	err = jpeg.Encode(out, img, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		return "", err
	}

	return directory, nil
}

func (ms *MemoryStore) DeleteImage(ctx context.Context, path string, name string) error {
	return nil
}
