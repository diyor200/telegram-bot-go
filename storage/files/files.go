package files

import (
	"gobot/gobot/lib/e"
	"gobot/gobot/storage"
	"os"
	"path/filepath"
)

type Storage struct {
	basePath string
}

const (
	defaulPerm = 0774 // read write permission code
)

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can't save", err) }()

	filePath := filepath.Join(s.basePath, page.UserName)
	if err := os.MkdirAll(filePath, defaulPerm); err != nil {
		return err
	}
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
