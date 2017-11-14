package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"math/rand"
	"strconv"
)

func main()  {

	http.HandleFunc("/service/random/word/generate",GenerateRandomWord)
	http.ListenAndServe(":8000",nil)

}

func GenerateRandomWord(r http.ResponseWriter,w *http.Request)  {

	urlInput := w.URL.Query().Get("number")
	limit, _ := strconv.Atoi(urlInput)

	dat, err := ioutil.ReadFile("words_alpha.txt")
	if err != nil {
		log.Fatal(err)
	}
	aa:=string(dat)
	pp:=strings.Split(aa,"\n")
	if limit==0{
		ranDomWord := fmt.Sprint("Random Word ...", pp[rand.Intn(len(pp))])
		fmt.Println(ranDomWord)
	}else {
		var ranDomWord string
		for i:=0;i<limit;i++ {
			ranDomWord = fmt.Sprint("Random Word ...", pp[rand.Intn(len(pp))])
			fmt.Println(ranDomWord)

		}

	}
}