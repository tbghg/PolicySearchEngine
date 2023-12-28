package scienceAndTechnology

import (
	"PolicySearchEngine/service"
	"fmt"
	"regexp"
)

type ScienceContentColly struct {
	// 单纯使用map匹配规则，自定义优先级较难实现
	rulesMatch map[*regexp.Regexp]func()
	// 使用数组记录规则顺序，实现优先级的排序
	rules []*regexp.Regexp
	// 等待处理的url队列
	waitQueue []string
}

func (s ScienceContentColly) Init() {
	s.register(
		// todo 示例，待完善
		regexp.MustCompile("https?://www\\.most\\.gov\\.cn/xxgk/.*\\.html?"),
		ruleExample,
	)
	// 兜底规则
	s.register(
		// todo 查下资料，看看正则兜底有什么好的写法
		nil,
		// todo 处理的话可以考虑统一报个日志，人工修一下
		nil,
	)
}

// 注册的规则数量可能会较多，单独拎个函数出来
func (s ScienceContentColly) register(rule *regexp.Regexp, f func()) {
	s.rulesMatch[rule] = f
	s.rules = append(s.rules, rule)
}

func (s ScienceContentColly) Import() (status int) {
	//TODO 从DB里面读取，目前可以先不接DB
	panic("implement me")
}

func (s ScienceContentColly) Run() {
	// todo waitQueue中的url均处理完即可
	for _, url := range s.waitQueue {
		fmt.Printf("I'm dealing %s...\n", url)
	}
}

func ruleExample() {
	// 示例规则
}

var _ service.ContentCrawler = (*ScienceContentColly)(nil)
