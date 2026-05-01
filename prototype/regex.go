package prototype

import (
	//"fmt"
	"io"
	"log"
	"net/http"
	//"os"
	//"path/filepath"
	"regexp"
	//"strings"
)

func Regex(){
	// fmt.Println("Hello World")

	url := "https://www.animeworld.ac/"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	// fmt.Println(string(body))
	strbody := string(body)

	tag, _ := regexp.Compile("<[^/<>](([^<>]*\".*\"[^<>]*>)|([^<>]*>))")

	tag.FindAllString(strbody, -1)
	

	// write the output
	/* err1 := os.WriteFile("./file.txt", []byte(strings.Join(tagfound, "\n")), 0644)
	if err1 != nil {
		log.Fatal("Errore nel scrivere il file")
	} */

	//regex for html <[^/<>](([^<>]*".*"[^<>]*>)|([^<>]*>))
}


