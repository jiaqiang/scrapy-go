package crawler

import (
	"fmt"
	"scrapy/core"
	"scrapy/settings"
	"scrapy/spiders"
)

type Crawler struct {
	spider   *spiders.Spider
	Settings *settings.CrawlerSettings
	engine   *core.ExecutionEngine
	crawling bool
}

func (c *Crawler) spiders() []spiders.Spider {
	return nil
}

func (c *Crawler) Crawl(args []string, kwargs []string) {
	c.crawling = true
	c.spider = c.createSpider()
	c.engine = c.createEngine()
	startRequests := c.spider.StartRequests()
	go c.engine.OpenSpider(c.spider, startRequests, true)
}

func (c *Crawler) Stop() {

}

func (c *Crawler) createSpider() *spiders.Spider {
	return nil
}

func (c *Crawler) createEngine() *core.ExecutionEngine {
	return core.NewExecutionEngine(c)
}

type CrawlerRunner struct {
	Settings *settings.CrawlerSettings
	crawlers []*Crawler
	aa       string
}

func (cr *CrawlerRunner) Crawl(crawlerOrSpiderCls interface{}, args []string, kwargs []string) bool {
	crawler := cr.CreateCrawler(crawlerOrSpiderCls)
	return cr.crawl(crawler, args, kwargs)
}

func (cr *CrawlerRunner) crawl(crawler *Crawler, args []string, kwargs []string) bool {
	cr.crawlers = append(cr.crawlers, crawler)
	crawler.Crawl(args, kwargs)
	return true
}

func (cr *CrawlerRunner) CreateCrawler(crawlerOrSpiderCls interface{}) *Crawler {
	_, ok := crawlerOrSpiderCls.(Crawler)
	if ok {
		return crawlerOrSpiderCls.(*Crawler)
	} else {
		return cr.createCrawler(crawlerOrSpiderCls)
	}
}

func (cr *CrawlerRunner) createCrawler(spiderCls interface{}) *Crawler {
	_, ok := spiderCls.(string)
	if ok {
		spiderCls = nil
	}
	return &Crawler{spider: spiderCls.(*spiders.Spider), Settings: cr.Settings}
}

func (cr *CrawlerRunner) Stop() {

}

func (cr *CrawlerRunner) Print() string {
	return fmt.Sprintf("%s .......", cr.aa)
}

type CrawlerProcess struct {
	*CrawlerRunner
	bb string
}

func (cp *CrawlerProcess) Start(stopAfterCrawl bool) {

}

func (cp *CrawlerProcess) Print() string {
	return fmt.Sprintf("%s ....... %s ", cp.aa, cp.bb)
}

func NewCrawlerProcess() *CrawlerProcess {
	return &CrawlerProcess{&CrawlerRunner{}, "bbb"}
}
