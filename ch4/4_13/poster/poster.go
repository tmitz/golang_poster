package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/tmitz/poster/lib/omdbapi"
)

const ImageDir = "posterimg"

func main() {
	title := flag.String("title", "", "映画のタイトルを英語で")
	year := flag.Int("year", 0, "映画の公開年")
	flag.Parse()

	result, err := omdbapi.SearchMovies(*title, *year)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Title:%s, Year:%s, Poster:%s", result.Title, result.Year, result.Poster)

	if _, err := os.Stat(ImageDir); os.IsNotExist(err) {
		os.Mkdir(ImageDir, 0777)
	}
	_, filename := path.Split(result.Poster)
	filename = ImageDir + "/" + filename

	resp, err := http.Get(result.Poster)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(body)
}
