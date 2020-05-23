package renderer


import (
	"path/filepath"
	"log"
	"io"
	"os"
	"html/template"
	"github.com/Ahrcantos/rtts/reddit"
)

func GenerateCommentHtml(c *reddit.Comment, w io.Writer) {
	temp, err := template.ParseFiles(filepath.FromSlash("styles/comment.html"))
	if err != nil {
		log.Fatal(err)
	}
	err = temp.Execute(w, c)

	if err != nil {
		log.Fatal(err)
	}
}

func (rc RenderContext) WriteCommentHtml(c *reddit.Comment) {
	path := filepath.Join(rc.BaseDir, c.Id, c.Id + ".html")
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	GenerateCommentHtml(c, file)
}

