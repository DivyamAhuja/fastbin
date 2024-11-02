package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"fastbin/internal/pkg/env"
	"fastbin/internal/web/views"
)

type api_write_response struct {
	Key string `json:"key"`
}

type api_read_response struct {
	Error string `json:"error"`
	Text  string `json:"text"`
}

func NewServer(port int) *http.Server {
	API_URL := env.GetEnv("API_URL", "localhost:8080")

	r := gin.Default()

	engineHTMLRenderer := r.HTMLRender
	r.HTMLRender = &HTMLTemplRenderer{FallbackHtmlRenderer: engineHTMLRenderer}

	r.StaticFS("/assets", http.FS(Files))

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", views.Write())
	})

	r.POST("/", func(ctx *gin.Context) {

		text := ctx.Request.FormValue("text")
		postBody, _ := json.Marshal(map[string]string{
			"text": text,
		})

		url := API_URL + "/write"
		response, err := http.Post(url, "application/json", bytes.NewBuffer(postBody))
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/500")
		}

		var api_res api_write_response
		decoder := json.NewDecoder(response.Body)

		err = decoder.Decode(&api_res)
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/500")
		}

		if response.StatusCode == http.StatusOK {
			ctx.Writer.Header().Add("Hx-Redirect", api_res.Key)
		} else {
			ctx.Redirect(http.StatusTemporaryRedirect, "/500")
		}

	})

	r.GET("/404", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", views.NotFound())
	})
	r.GET("/500", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", views.ServerError())
	})

	r.GET("/:key", func(ctx *gin.Context) {
		url := API_URL + "/read/" + ctx.Param("key")

		response, err := http.Get(url)
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/500")
		}

		var api_res api_read_response
		decoder := json.NewDecoder(response.Body)

		err = decoder.Decode(&api_res)
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/500")
		}

		if response.StatusCode == http.StatusOK {
			ctx.HTML(http.StatusOK, "", views.Read(api_res.Text))
		} else if response.StatusCode == http.StatusNotFound {
			ctx.Redirect(http.StatusTemporaryRedirect, "/404")
		} else {
			ctx.Redirect(http.StatusTemporaryRedirect, "/500")
		}

	})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
