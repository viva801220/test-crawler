package service

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/gocolly/colly"
)

func VisitWeb(url string) {

	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("href")
		fmt.Println(link)
		e.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request) {
		logs.Notice("Visiting", r.URL)
	})

	c.Visit(url)
}
