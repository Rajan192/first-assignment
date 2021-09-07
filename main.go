package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
<<<<<<< HEAD
	"time"
=======
>>>>>>> 498931a8c7b74d58433ea2fc30441893788d7506

	"golang.org/x/net/html"
)

var links []string
var text []string

func main() {

<<<<<<< HEAD
	//i have added here a client time out suppose get request not complet their work within
	//10 second it will return erroor

	netclient:=&http.Client{
		Timeout: time.Second*10,
	}
	resp, err := netclient.Get("https://github.com/gophercises/link/blob/master/ex2.html")
=======
	resp, err := http.Get("https://github.com/gophercises/link/blob/master/ex2.html")
>>>>>>> 498931a8c7b74d58433ea2fc30441893788d7506
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
<<<<<<< HEAD
				//for make sure that, it is text only i am ussing if condition
=======
				//for make sure that, it is text only i am using if condition
>>>>>>> 498931a8c7b74d58433ea2fc30441893788d7506
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
