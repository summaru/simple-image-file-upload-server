package customutil

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type IndexHandler struct {
	path string
}

func (ih IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	dir, _ := os.Getwd()
	docpath := dir + "\\htdoc"
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		ih.path = r.URL.Path
		http.ServeFile(w, r, docpath+strings.Replace(ih.path, "/", "\\", -1))
		fmt.Println("serve to Client that it: " + ih.path)
	case "POST":
		r.ParseForm()
		f, fh, _ := r.FormFile("imagefile")
		defer f.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(f)
		size := buf.Len()
		filebyte := make([]byte, size)

		buf.Read(filebyte)
		_, errsave := os.Stat("saveImage")
		if errsave != nil {
			log.Fatal(errsave)
		}
		file, err := os.OpenFile(
			"saveImage\\"+fh.Filename,
			os.O_CREATE,
			os.FileMode(0644))

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		file.Write(filebyte)
		fmt.Println("save image file")
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}
