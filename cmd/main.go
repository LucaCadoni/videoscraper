package main

import (
	"fmt"
	"time"
	"github.com/LucaCadoni/videoscraper/prototype"
)

func main(){
	// fmt.Println("Hello World")
	fmt.Println("Char Method")
	start := time.Now()
	prototype.Charescape()
	charelapse := time.Since(start)
	fmt.Printf("Time Elapsed: %d\n", charelapse)

	fmt.Println("Regex Method")
	start = time.Now()
	prototype.Regex()
	regexelapse := time.Since(start)
	fmt.Printf("Time Elapsed: %d\n", regexelapse)

	if charelapse <= regexelapse{
		fmt.Println("Char Method has won")
	}else{
		fmt.Println("Regex Method has won")
	}
}


