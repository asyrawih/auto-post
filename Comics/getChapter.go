package Comics

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/hananloser/auto-post/types"
)

func GetChapter(c *colly.Collector, url string) (chapterList []types.Chapters) {

	c.OnHTML("span.leftoff", func(e *colly.HTMLElement) {
		chapter := e.ChildAttr("a" , "href")
		chapterLink := e.Text
		tmp := types.Chapters{}
		tmp.Chapter = chapter
		tmp.Link = chapterLink
		chapterList = append(chapterList , tmp)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting Manga ...", request.URL.String())
	})

	_ = c.Visit(url)

	c.Wait()
	return chapterList
}