package main

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

var g errgroup.Group

func main() {
	mainServer := &http.Server{
		Addr:    ":8000",
		Handler: mainRouter(),
	}
	fileServer := &http.Server{
		Addr:    ":8100",
		Handler: fsRouter(),
	}
	g.Go(func() error {
		return mainServer.ListenAndServe()
	})
	g.Go(func() error {
		return fileServer.ListenAndServe()
	})
	log.Println("Service Running..")
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
func fsRouter() http.Handler {
	r := echo.New()
	r.GET("/vid", func(c echo.Context) error {
		return c.File("vid.mp4")
	})
	return r
}
func mainRouter() http.Handler {
	r := echo.New()
	r.GET("/vid", func(c echo.Context) error {
		req, _ := http.NewRequest("GET", "http://localhost:8100/vid", nil)
		client := &http.Client{}
		req.Header.Set("Range", c.Request().Header.Get("Range"))
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}
		c.Response().Header().Set(echo.HeaderContentLength, resp.Header.Get("Content-Length"))
		contentRange := resp.Header.Get("Content-Range")
		if contentRange != "" {
			c.Response().Header().Set("Content-Range", contentRange)
		}
		defer resp.Body.Close()
		return c.Stream(resp.StatusCode, "video/mp4", resp.Body)
	})
	r.GET("", func(c echo.Context) error {
		return c.File("./static/index.html")
	})
	return r
}
