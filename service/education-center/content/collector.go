package content

import (
	"PolicySearchEngine/service"
	"PolicySearchEngine/utils"
	"github.com/gocolly/colly"
	"regexp"
)

func (s *EducationContentColly) getRules() []*service.Rule {
	return []*service.Rule{
		s.zcfgCollector(),
		s.srcsiteCollector(),
	}
}

func (s *EducationContentColly) updateTitle(e *colly.HTMLElement) {
	title := utils.TidyString(e.Text)
	s.metaDal.UpdateMetaTitle(title, e.Request.URL.String())
}

func (s *EducationContentColly) zcfgCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.moe\\.gov\\.cn/jyb_sjzl/sjzl_zcfg/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: ".moe-detail-box h1",
		F:           s.updateTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: ".TRS_Editor",
		F:           service.NormalContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}

func (s *EducationContentColly) srcsiteCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.moe\\.gov\\.cn/srcsite/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: ".details-policy-box h1",
		F:           s.updateTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: ".details-policy-box",
		F:           service.NormalContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}
