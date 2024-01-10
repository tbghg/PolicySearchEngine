package scienceAndTechnology

import (
	"PolicySearchEngine/service"
	"PolicySearchEngine/service/scienceAndTechnology/content"
	"PolicySearchEngine/service/scienceAndTechnology/meta"
)

type ScienceColly struct {
	content *content.ScienceContentColly
	meta    *meta.ScienceMetaColly
}

func (s *ScienceColly) Meta() service.MetaCrawler {
	return (service.MetaCrawler)(s.meta)
}

func (s *ScienceColly) Content() service.ContentCrawler {
	return (service.ContentCrawler)(s.content)
}

func (s *ScienceColly) Register(crawler *service.Crawler) {
	s.content = new(content.ScienceContentColly)
	s.meta = new(meta.ScienceMetaColly)

	crawler.Register(s)
}

var _ service.ServiceCrawler = (*ScienceColly)(nil)
