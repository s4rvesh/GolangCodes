package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Homepage endpoint hit")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", editArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("End point hit return allArticles")

	json.NewEncoder(w).Encode(Articles)
}

func editArticle(w http.ResponseWriter, r *http.Request) {

	fmt.Println("hit edit block")

	vars := mux.Vars(r)
	key := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)

	var artiedit Article

	json.Unmarshal(reqBody, &artiedit)

	for index, article := range Articles {

		if article.Id == key {

			/*article.Content = artiedit.Content
			article.Desc = artiedit.Desc
			article.Title = artiedit.Title*/
			Articles = append(Articles[:index], Articles[index+1:]...)
			Articles = append(Articles, artiedit)

		}
	}

	//Articles = append(Articles, article)

	fmt.Fprintf(w, "%+v", string(reqBody))

}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "key: "+key)

	for _, article := range Articles {

		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {

	fmt.Println("post reached")

	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	fmt.Fprintf(w, "%+v", string(reqBody))
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {

	fmt.Println("hit Delete block")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {

			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func main() {

	fmt.Println("Hello World!!")

	Articles = []Article{

		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequest()

}
