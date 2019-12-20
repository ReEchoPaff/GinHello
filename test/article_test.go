package test

import (
	"GinHello/initRouter"
	"GinHello/model"
	"GinHello/param"
	"bytes"
	"encoding/json"
	"gopkg.in/go-playground/assert.v1"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	router = initRouter.SetupRouter()
}

// 文章添加
func TestInsertArticle(t *testing.T) {
	article := model.Article{
		Type:    "go",
		Content: "hello~~",
	}
	marshal, _ := json.Marshal(article)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/article", bytes.NewBufferString(string(marshal)))
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)
	var restMessage param.RestMessage

	if err := json.NewDecoder(w.Body).Decode(&restMessage); err != nil {
		log.Panicln("解析w.code出错: ", err.Error())
	}

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, 200, restMessage.Code)

}
