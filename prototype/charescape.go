package prototype

import (
	//"fmt"
	"io"
	"log"
	"net/http"
)

func Charescape(){
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

	start := -2
	end := -1
	currentInString := 0
	for pos, char := range(strbody){
		if(char == 60 && currentInString == 0){
			//fmt.Printf("Element start at %d (%c)", pos, char)
			start = pos
		}

		if(char == 47 && start != -2 && end == -1 && currentInString == 0){
			// ignore close
			continue
		}

		if(char == 22){
			switch currentInString{
				case 0:
					currentInString = 1
				case 1:
					currentInString = 0
			}
		}

		if(char == 27){
			switch currentInString{
				case 0:
					currentInString = 1
				case 1:
					currentInString = 0
			}
		}

		if(char == 62 && start != -2 && currentInString == 0){
			end = pos
			//fmt.Print(string(body[start:end + 1]), "\n")
			start = -2
			end = -1
		}
	}

	//regex for html <[^/<>](([^<>]*".*"[^<>]*>)|([^<>]*>))
}

