package main

import (
	"fmt"
	"os"
	//"time"
)

func main(){
	if len(os.Args) < 2 || len(os.Args) > 3{
		fmt.Println("Error usage")
		return
	}

	fontFile := "standard.txt"

	if len(os.Args) == 3{
		fontFile = os.Args[2] + ".txt"
	}

	input := os.Args[1]

	data, err := loadBanner(fontFile)
	if err != nil{
		fmt.Println("Error loading banner")
		return
	}

	result := renderBanner(input, data)


	for _, line := range result{
		fmt.Println(line)
		//time.Sleep(400 * time.Millisecond)
	}

}