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

func (s *ScienceColly) Register(crawlers *service.Crawlers) {
	s.content = new(content.ScienceContentColly)
	s.meta = new(meta.ScienceMetaColly)

	crawlers.Register(s)
}

var _ service.Crawler = (*ScienceColly)(nil)
