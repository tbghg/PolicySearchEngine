# PolicySearchEngine 

毕设题目：面向公共政策文本检索的搜索引擎设计与实现

目前写的是爬虫部分，等爬虫写好后把整个项目结构改改，然后再做web部分

目前爬虫部分遇到的问题：

1. 爬取标题时，列表页面爬取的不全，字比较多的话是省略号，只有打开文章内部才能看到真实标题，例如：科技部 住房城乡建设部关于印发《“十四五”城镇化与城市发展科技创新专项规划》的...
2. 需要取监听，其实是meta监听，把要爬的文章给content，但目前两者没有任何关系
   1. 改进的话可以直接 meta -> MQ -> content 但又感觉没啥必要
   2. meta每天扫一次就行，content的话也是每天启动一次吧，到时候看DB里面的更新时间，更新时间再一天之内说明是新发布的文章，就把它扫了

12.29问题记录：
接入redis维护状态的持久化存储后发现，它是以colly对象为单位，当前的结构安排会把初始目录页也存入其中，导致下次读取时统一跳过。
因为内容是要不断更新的，增量的内容只看增量，如果重新全量读取一遍，那得重新建索引等等吧，有些支撑不住，这个似乎是必须得解决的
解决方案：
1. 单独拆个colly对象 - 整体结构都写好了，不想改了。。。
2. 自己把`github.com/gocolly/redisstorage`改一下，做些定制化处理，看了下代码，只能单独配置不做处理的url去匹配

权衡之后打算用方法二，源代码是把url做个处理得到requestID，然后加前缀作为redis的key，这个源代码来自colly库，我肯定是改不动的，但是redisstorage库就好改很多，对IsVisited或Visited做一下定制化处理即可

```go
	if checkRevisit && !c.AllowURLRevisit && method == "GET" {
		h := fnv.New64a()
		h.Write([]byte(u))
		uHash := h.Sum64()
		visited, err := c.store.IsVisited(uHash)
		if err != nil {
			return err
		}
		if visited {
			return ErrAlreadyVisited
		}
		return c.store.Visited(uHash)
	}
```

12.29记录：目前该写DB部分了，把DB写了，然后开始搞content部分，搞成后算是结束了爬虫的一个模块，之后开发就会快很多了
