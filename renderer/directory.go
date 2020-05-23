package renderer

import (
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"os"
)

type RenderContext struct {
	BaseDir string
}

func NewRenderContext() RenderContext {
	path, err := ioutil.TempDir("", "render-context")
	if err != nil {
		log.Fatal(err)
	}

	copyStyle(path)

	path = filepath.ToSlash(path)

	rc := RenderContext{path}

	return rc
}

func (rc RenderContext) NewCommentDir(id string) {

	// Create Directory
	path := filepath.Join(rc.BaseDir, id)
	err := os.Mkdir(path, 0700)
	if err != nil {
		log.Fatal(err)
	}

	//Symlink Style
	rc.symlinkStyle(path)
}

func (rc RenderContext) symlinkStyle(path string) {
	err := os.Symlink(filepath.Join(rc.BaseDir, "reddit.css"), filepath.Join(path, "reddit.css"))
	if err != nil {
		log.Fatal(err)
	}
}

func copyStyle(dir string) {
	dir = filepath.FromSlash(dir)

	styleFile, err := os.Open(filepath.FromSlash("./styles/reddit.css"))
	if err != nil {
		log.Fatal(err)
	}
	defer styleFile.Close()

	styleCopy, err := os.Create(filepath.Join(dir, "reddit.css"))
	if err != nil {
		log.Fatal(err)
	}
	defer styleCopy.Close()

	_, err = io.Copy(styleCopy, styleFile)
	if err != nil {
		log.Fatal(err)
	}
	err = styleCopy.Sync()
	if err != nil {
		log.Fatal(err)
	}

}
