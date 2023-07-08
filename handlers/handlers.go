package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Nattchi/blog-api/models"
	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "Fail to decode json\n", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	io.WriteString(w, "Post article!...\n")

	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	resStr := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resStr)

	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resStr := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resStr)

	article1 := models.Article1
	json.NewEncoder(w).Encode(article1)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "Fail to decode json\n", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	io.WriteString(w, "Post Nice...\n")

	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "Fail to decode json\n", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	io.WriteString(w, "Post Comment...\n")

	comment := reqComment
	json.NewEncoder(w).Encode(comment)
}
