package main

import (
	gove "gmr/gmr-web/gove"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
	
}

func main() {
	r := gove.New()
	r.Use(gove.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gove.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gove.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gove.H{
			"title":  "gove",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gove.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gove.H{
			"title": "gove",
			"now":   time.Date(2022, 6, 19, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}

// day-07
// func main() {
// 	r := gove.Default()
// 	r.GET("/", func(c *gove.Context) {
// 		c.String(http.StatusOK, "Hello Geektutu\n")
// 	})
// 	// index out of range for testing Recovery()
// 	r.GET("/panic", func(c *gove.Context) {
// 		names := []string{"geektutu"}
// 		c.String(http.StatusOK, names[100])
// 	})
// 	r.Run(":9999")
// }