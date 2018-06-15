package actions

import "scrapy/settings"
import "scrapy/crawler"

type ScrapyAction struct {
	Name            string
	RequiresProject bool
	settings        settings.Settings
	crawler         crawler.Crawler
}

func (c *ScrapyAction) SetCrawler(crawler crawler.Crawler) {
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
}
