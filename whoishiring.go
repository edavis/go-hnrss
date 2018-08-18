package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func HiringCommon(c *gin.Context, query string) {
	params := make(url.Values)
	params.Set("query", fmt.Sprintf("\"%s\"", query))
	params.Set("tags", "story,author_whoishiring")
	params.Set("hitsPerPage", "1")

	results, err := GetResults(params)
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}

	if len(results.Hits) < 1 {
		c.String(http.StatusBadGateway, "No results found")
		return
	}

	sp, op := ParseRequest(c)

	sp.Tags = "comment,story_" + results.Hits[0].ObjectID
	sp.SearchAttributes = "default"
	op.Title = results.Hits[0].Title
	op.Link = "https://news.ycombinator.com/item?id=" + results.Hits[0].ObjectID
	op.TopLevel = true

	Generate(c, sp, op)
}

func SeekingEmployees(c *gin.Context) {
	HiringCommon(c, "Ask HN: Who is hiring?")
}

func SeekingEmployers(c *gin.Context) {
	HiringCommon(c, "Ask HN: Who wants to be hired?")
}

func SeekingFreelance(c *gin.Context) {
	HiringCommon(c, "Ask HN: Freelancer? Seeking freelancer?")
}