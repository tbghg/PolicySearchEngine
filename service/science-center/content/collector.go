package content

import (
	"PolicySearchEngine/service"
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

func (s *ScienceContentColly) xxgkCollector() *service.Rule {
	rule := regexp.MustCompile("https?://www\\.most\\.gov\\.cn/xxgk/.*\\.html?")

	hf := &service.HtmlFunc{
		QuerySelect: ".xxgk_detail_content",
		F: func(e *colly.HTMLElement) {
			//fmt.Println(e.DOM.Html())

			title := e.ChildText(".xxgk_title")
			regex := regexp.MustCompile(`[\n\t]`)
			cleanedTitle := regex.ReplaceAllString(title, "")
			fmt.Println(cleanedTitle)

			//content := e.ChildText("#Zoom")
			//fmt.Println(content)
		},
	}

	return service.NormalRule(rule, hf)
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
		F:           service.NormalTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: ".Custom_UnionStyle",
		F:           service.NormalContent,
	}

	return service.NormalRule(combinedRule, hfTitle, hfContent)
}

func (s *ScienceContentColly) kjbgzCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.most\\.gov\\.cn/kjbgz/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: "#Title",
		F:           service.NormalTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: ".Zoom",
		F:           service.NormalContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}

func (s *ScienceContentColly) zhengceCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.gov\\.cn/zhengce/content/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: "td[colspan='3']",
		F:           service.NormalTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: "#UCAP-CONTENT",
		F:           service.NormalContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}

func (s *ScienceContentColly) gongbaoCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.gov\\.cn/gongbao/content/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: ".share-title",
		F:           service.NormalTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: ".pages_content",
		F:           service.NormalContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}

func (s *ScienceContentColly) xinwenCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.gov\\.cn/xinwen/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: "#ti",
		F:           service.NormalTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: ".pages_content",
		F:           service.NormalContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}

func (s *ScienceContentColly) chinataxCollector() *service.Rule {

	rule := regexp.MustCompile("https?://www\\.chinatax\\.gov\\.cn/.*\\.html?")

	hfTitle := &service.HtmlFunc{
		QuerySelect: "title",
		F:           service.NormalTitle,
	}

	hfContent := &service.HtmlFunc{
		QuerySelect: "#fontzoom",
		F:           service.NormalContent,
	}

	return service.NormalRule(rule, hfTitle, hfContent)
}
