package actions

import (
	"scrapy/core"
	"scrapy/settings"
)

type ScrapyAction struct {
	Name            string
	RequiresProject bool
	CrawlerProcess  *core.CrawlerProcess
	Settings        *settings.Settings
	crawler         *core.Crawler
}

func (c *ScrapyAction) SetCrawler(crawler *core.Crawler) {
	c.crawler = crawler
}

func (c *ScrapyAction) Syntax() string {
	return ""
}

func (c *ScrapyAction) ShortDesc() string {
	return ""
}

func (c *ScrapyAction) LongDesc() string {
	return c.ShortDesc()
}

func (c *ScrapyAction) Help() string {
	return c.LongDesc()
}

// Populate option parse with options available for this command
func (c *ScrapyAction) AddOptions() {

}

func (c *ScrapyAction) ProcessOptions() string {
	return ""
}

func (c *ScrapyAction) Run(args []string, opts []string) {
	c.CrawlerProcess.Crawl(nil, args, opts)
	c.CrawlerProcess.Start(true)
}
