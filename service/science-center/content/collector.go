package content

import (
	"PolicySearchEngine/dao/es"
	"PolicySearchEngine/service"
	"PolicySearchEngine/utils"
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

func (s *ScienceContentColly) getRules() []*service.Rule {
	return []*service.Rule{
		s.xxgkCollector(),
		s.kjzcCollector(),
		s.kjbgzCollector(),
		s.zhengceCollector(),
		s.gongbaoCollector(),
		s.xinwenCollector(),
		s.chinataxCollector(),
	}
}

func (s *ScienceContentColly) updateTitle(e *colly.HTMLElement) {
	title := utils.TidyString(e.Text)
	s.metaDal.UpdateMetaTitle(title, e.Request.URL.String())
}

func (s *ScienceContentColly) updateContent(e *colly.HTMLElement) {
	var text []byte
	e.ForEach("*", func(_ int, child *colly.HTMLElement) {
		label := strings.ToLower(child.Name)
		if label == "style" || label == "table" || label == "script" {
			return
		}
		text = append(text, []byte(child.Text)...)
	})
	s.contentDal.InsertContent(e.Request.URL.String(), string(text))

	meta := s.metaDal.GetMetaByUrl(e.Request.URL.String())
	if meta == nil {
		meta = s.metaDal.GetMetaByUrl(e.Request.Headers.Get("Referer"))
	}
	if meta != nil {
		sdIDs := s.dMapDal.GetDepartmentIDsByMetaID(meta.ID)
		es.IndexDoc(meta.Date, meta.DepartmentID, meta.ProvinceID, meta.Title, meta.Url, string(text), sdIDs)
	}
}

func (s *ScienceContentColly) xxgkCollector() *service.Rule {
	rule := regexp.MustCompile("https?://www\\.most\\.gov\\.cn/xxgk/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: ".xxgk_title",
		F:           s.updateTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: "#Zoom",
		F:           s.updateContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}

func (s *ScienceContentColly) kjzcCollector() *service.Rule {

	rule1 := regexp.MustCompile("https?://www\\.most\\.gov\\.cn/satp/kjzc/zh/.*\\.html?")
	rule2 := regexp.MustCompile("https?://www\\.most\\.gov\\.cn/tztg/.*\\.html?")

	combinedRule := regexp.MustCompile(fmt.Sprintf(
		"(%s|%s)",
		rule1.String(),
		rule2.String(),
	))

	hfTitle := &service.HtmlFunc{
		QuerySelect: "#Title",
		F:           s.updateTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: "#Zoom",
		F:           s.updateContent,
	}

	return service.NormalRule(combinedRule, hfTitle, hfContent)
}

func (s *ScienceContentColly) kjbgzCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.most\\.gov\\.cn/kjbgz/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: "#Title",
		F:           s.updateTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: "#Zoom",
		F:           s.updateContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}

func (s *ScienceContentColly) zhengceCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.gov\\.cn/zhengce/content/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: "td[colspan='3']",
		F:           s.updateTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: "#UCAP-CONTENT",
		F:           s.updateContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}

func (s *ScienceContentColly) gongbaoCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.gov\\.cn/gongbao/content/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: ".share-title",
		F:           s.updateTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: ".pages_content",
		F:           s.updateContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}

func (s *ScienceContentColly) xinwenCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.gov\\.cn/xinwen/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: "#ti",
		F:           s.updateTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: ".pages_content",
		F:           s.updateContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}

func (s *ScienceContentColly) chinataxCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.chinatax\\.gov\\.cn/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: "title",
		F:           s.updateTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: "#fontzoom",
		F:           s.updateContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}
