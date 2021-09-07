package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

var links []string
var text []string

func main() {

	resp, err := http.Get("https://github.com/gophercises/link/blob/master/ex2.html")
	if err != nil {
		log.Fatal("Error in Fetching URL ", err)
	}

	tokenizer := html.NewTokenizer(resp.Body)
	//make sure to close the body
	defer resp.Body.Close()

	//getting text and link here

	for {
		tokenType := tokenizer.Next()
		//check for the error
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			log.Fatalf("Some error %v", tokenizer.Err())
		}

		token := tokenizer.Token()
		if tokenType == html.StartTagToken {
			if token.Data == "a" {
				for _, a := range token.Attr {
					if a.Key == "href" {
						//fmt.Println(a.Val)
						//apned link in slice
						links = append(links, a.Val)
					}
				}
				//if you are in <a> tag so next tocken will be the text
				nextTockenTypeforText := tokenizer.Next()
				//for make sure that, it is text only i am ussing if condition
				if nextTockenTypeforText == html.TextToken {
					linkText := tokenizer.Token().Data
					text = append(text, linkText)
				}

			}

		}
	}
	fmt.Println("all text over <a> tag ", text)
	fmt.Println("all link under <a> tag ", links)
}
