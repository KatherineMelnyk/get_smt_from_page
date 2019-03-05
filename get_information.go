package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func findInf(n int, s, c string) string {
	res := ""
	for i := n; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			res += string(s[i])
		}
		if string(s[i]) == c {
			break
		}
	}
	return res
}

func main() {
	information, err := http.Get("https://www.instagram.com/kateymelnyk/")
	if err != nil {
		log.Fatal(err)
	}
	defer information.Body.Close()
	dataInBytes, err := ioutil.ReadAll(information.Body)
	PageData := string(dataInBytes)

	Index := strings.Index(PageData, "\"edge_followed_by\"")
	if Index == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}
	fmt.Printf("Followers: %s\n", findInf(Index, PageData, "}"))

	In := strings.Index(PageData, "\"edge_follow\"")
	if In == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}
	fmt.Printf("Following: %s\n", findInf(In, PageData, "}"))

	I := strings.Index(PageData, "\"edge_owner_to_timeline_media\"")
	if I == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}
	fmt.Printf("Posts: %s\n", findInf(I, PageData, ","))
}
