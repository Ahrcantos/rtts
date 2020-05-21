package renderer


import (
	"path/filepath"
	"log"
	"io"
	"html/template"
	"github.com/Ahrcantos/rtts/reddit"
)

func GenerateHtml(c *reddit.Comment, w io.Writer) {
	temp, err := template.ParseFiles(filepath.FromSlash("styles/comment.html"))
	if err != nil {
		log.Fatal(err)
	}
	err = temp.Execute(w, c)

	if err != nil {
		log.Fatal(err)
	}

}
