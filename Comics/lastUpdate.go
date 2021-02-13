package Comics

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/hananloser/auto-post/types"
)

func LastUpdate(c *colly.Collector)(listManga []types.MangaList )  {

	c.OnHTML("div.imgu", func(e *colly.HTMLElement) {
		manga := e.ChildAttr("a[href]", "title")
		link := e.ChildAttr("a[href]", "href")
		tmp := types.MangaList{}
		tmp.Name = manga
		tmp.Link = link
		listManga = append(listManga, tmp)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Get Latest Updates" , request.URL.String())
	})

	_ = c.Visit("https://komikcast.com")

	c.Wait()

	return listManga
}