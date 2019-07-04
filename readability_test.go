// Package readability is a Go package that find the main readable
// content from a HTML page. It works by removing clutter like buttons,
// ads, background images, script, etc.
//
// This package is based from Readability.js by Mozilla, and written line
// by line to make sure it looks and works as similar as possible. This
// way, hopefully all web page that can be parsed by Readability.js
// are parse-able by go-readability as well.
package readability_test

//go tool pprof -web cpu.pprof

import (
	"fmt"
	"github.com/pubgo/errors"
	"github.com/pubgo/go-readability"
	"github.com/pubgo/gotask"
	"github.com/rs/zerolog"
	"strings"
	"testing"
	"time"
)

// FromReader parses input from an `io.Reader` and returns the
// readable content. It's the wrapper for `Parser.Parse()` and useful
// if you only want to use the default parser.
func TestFromURL(t *testing.T) {
	defer errors.Debug()

	/*
		<meta name="keywords" content="研究观点,艾瑞网,艾瑞咨询,新经济门户,互联网数据资讯">
		<meta name="description" content="研究观点,艾瑞网聚合互联网数据资讯,融合互联网行业资源">
	*/

	r := readability.FromURL("https://blog.csdn.net/alvine008/article/details/51282868", time.Second*3)
	//r := readability.FromURL("https://mp.weixin.qq.com/s?src=11&timestamp=1561273013&ver=1685&signature=wLwncgbNaHOci0d*RiI1L49pPp1aw-wXgYoCNQjLG2tyDyK92tGy-YRc1mNeZz1LstaGoWlF9exPzFapdqVkD*r-6HAEhCZ2QW7uj*Va79CaF6TzHPy5VgQhlfAeFqUa&new=1", time.Second*3)
	//r := readability.FromURL("https://www.baidu.com/link?url=sxfIBe7sLxoSFhGklac_DwnHsmjVfJbeq_UpJdA59JDOP2s2Svr5znv-K7m-Prwd&wd=&eqid=b42f79ff00072352000000045d0e3f9c",time.Second*3)
	errors.P(r.Copy())

}

func TestFromReader(t *testing.T) {
	defer errors.Debug()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	gotask.TaskRegistry("fn", func() {
		//defer errors.Resp(func(err *errors.Err) {
			//err.Log()
		//})

		readability.FromReader(strings.NewReader(_html), "https://www.cnblogs.com/52php/p/7076484.html")
	})

	task := gotask.NewTask(4, time.Second*3)

	for i := 0; i < 10; i++ {
		task.Do("fn")
	}
	task.Wait()
	fmt.Println(task.Stat())
}

const _html = `
<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="referrer" content="origin" />
    <meta http-equiv="Cache-Control" content="no-transform" />
    <meta http-equiv="Cache-Control" content="no-siteapp" />
    <title>[Go] sync.Pool 的实现原理 和 适用场景 - 52php - 博客园</title>
    <meta property="og:description" content="摘录一： Go 1.3 的 sync 包中加入一个新特性：Pool。 官方文档可以看这里&#160;http://golang.org/pkg/sync/#Pool 这个类设计的目的是用来保存和复用临" />
    <link type="text/css" rel="stylesheet" href="/bundles/blog-common.css?v=KOZafwuaDasEedEenI5aTy8aXH0epbm6VUJ0v3vsT_Q1"/>
<link id="MainCss" type="text/css" rel="stylesheet" href="/skins/BlueFresh/bundle-BlueFresh.css?v=b3yN5lKt4bOBKbsD14DAsAqpqqfyM4Ejm9hmWmfDgXg1"/>
<link type="text/css" rel="stylesheet" href="/blog/customcss/110339.css?v=Sp9xpng2yAx%2b5GpkCVAq3%2fU8ErM%3d"/>
<link id="mobile-style" media="only screen and (max-width: 767px)" type="text/css" rel="stylesheet" href="/skins/BlueFresh/bundle-BlueFresh-mobile.css?v=H3dUoyw06nS24wgnqYkz7A9cE_1mtRYiCQYMxIir4zs1"/>
    <link title="RSS" type="application/rss+xml" rel="alternate" href="https://www.cnblogs.com/52php/rss"/>
    <link title="RSD" type="application/rsd+xml" rel="EditURI" href="https://www.cnblogs.com/52php/rsd.xml"/>
<link type="application/wlwmanifest+xml" rel="wlwmanifest" href="https://www.cnblogs.com/52php/wlwmanifest.xml"/>
    <script src="//common.cnblogs.com/scripts/jquery-2.2.0.min.js"></script>
    <script>var currentBlogId=110339;var currentBlogApp='52php',cb_enable_mathjax=false;var isLogined=false;</script>
    <script src="/bundles/blog-common.js?v=smtcUT5dhdu_5eEO8CKHYoVc7DPLgEBGzp6zKkstlzg1" type="text/javascript"></script>
</head>
<body>
<a name="top"></a>

<!--PageBeginHtml Block Begin-->
<ul class="clearfix" id="navList_new">
    <li><a href="http://52php.cnblogs.com/">首页</a></li>
    <li><a rel="nofollow" href="http://i.cnblogs.com/EditPosts.aspx?opt=1">添加随笔</a></li>
</ul>
<!--PageBeginHtml Block End-->

<!--done-->
<div id="home">
<div id="header">
	<div id="blogTitle">
	<a id="lnkBlogLogo" href="https://www.cnblogs.com/52php/"><img id="blogLogo" src="/Skins/custom/images/logo.gif" alt="返回主页" /></a>			
		
<!--done-->
<h1><a id="Header1_HeaderTitle" class="headermaintitle" href="https://www.cnblogs.com/52php/">52PHP</a></h1>
<h2><span style="color:red">革命尚未成功，同志仍须努力。。。</h2>



		
	</div><!--end: blogTitle 博客的标题和副标题 -->
	<div id="navigator">
		
<ul id="navList">
<li></li>
<li></li>
<li></li>
<li></li>
<li>
<!----></li>
<li><a id="blog_nav_admin" class="menu" rel="nofollow" href="https://i.cnblogs.com/">管理</a></li>
</ul>
		<div class="blogStats">
			
			<div id="blog_stats">
<span id="stats_post_count">随笔 - 1026&nbsp; </span>
<span id="stats_article_count">文章 - 0&nbsp; </span>
<span id="stats-comment_count">评论 - 84</span>
</div>
			
		</div><!--end: blogStats -->
	</div><!--end: navigator 博客导航栏 -->
</div><!--end: header 头部 -->

<div id="main">
	<div id="mainContent">
	<div class="forFlow">
		
        <div id="post_detail">
<!--done-->
<div id="topics">
	<div class = "post">
		<h1 class = "postTitle">
			<a id="cb_post_title_url" class="postTitle2" href="https://www.cnblogs.com/52php/p/7076484.html">[Go] sync.Pool 的实现原理 和 适用场景</a>
		</h1>
		<div class="clear"></div>
		<div class="postBody">
			<div id="cnblogs_post_body" class="blogpost-body"><h2>摘录一：</h2>
<p>Go 1.3 的 sync 包中加入一个新特性：Pool。</p>
<p>官方文档可以看这里&nbsp;<a href="http://golang.org/pkg/sync/#Pool" target="_blank">http://golang.org/pkg/sync/#Pool</a></p>
<p>这个类设计的目的是用来保存和复用临时对象，以减少内存分配，降低CG压力。</p>
<pre class="cnblogs_Highlighter brush:go;">type Pool  
    func (p *Pool) Get() interface{}  
    func (p *Pool) Put(x interface{})  
    New func() interface{}  </pre>
<p>Get 返回 Pool 中的任意一个对象。</p>
<p>如果 Pool 为空，则调用 New 返回一个新创建的对象。</p>
<p>&nbsp;</p>
<p>如果没有设置 New，则返回 nil。</p>
<p>还有一个重要的特性是，放进 Pool 中的对象，会在说不准什么时候被回收掉。</p>
<p>所以如果事先 Put 进去 100 个对象，下次 Get 的时候发现 Pool 是空也是有可能的。</p>
<p>不过这个特性的一个好处就在于不用担心 Pool 会一直增长，因为 Go 已经帮你在 Pool 中做了回收机制。</p>
<p>这个清理过程是在每次垃圾回收之前做的。垃圾回收是固定两分钟触发一次。</p>
<p>而且每次清理会将 Pool 中的所有对象都清理掉！</p>
<p>&nbsp;</p>
<pre class="cnblogs_Highlighter brush:go;">package main

import(
    "sync"
    "log"
)

func main(){
    // 建立对象
    var pipe = &amp;sync.Pool{New:func()interface{}{return "Hello, BeiJing"}}
    
    // 准备放入的字符串
    val := "Hello,World!"
    
    // 放入
    pipe.Put(val)
    
    // 取出
    log.Println(pipe.Get())
    
    // 再取就没有了,会自动调用NEW
    log.Println(pipe.Get())
}

// 输出
2014/09/30 15:43:30 Hello, World!
2014/09/30 15:43:30 Hello, BeiJing</pre>
<p>&nbsp;</p>
<p>摘自：<a href="http://www.nljb.net/default/sync.Pool/" target="_blank">http://www.nljb.net/default/sync.Pool/</a></p>
<p>&nbsp;</p>
<h2>摘录二：</h2>
<p>众所周知，go 是自动垃圾回收的(garbage collector)，这大大减少了程序编程负担。但 gc 是一把双刃剑，带来了编程的方便但同时也增加了运行时开销，使用不当甚至会严重影响程序的性能。因此性能要求高的场景不能任意产生太多的垃圾（有gc但又不能完全依赖它挺恶心的），如何解决呢？那就是要重用对象了，我们可以简单的使用一个 chan 把这些可重用的对象缓存起来，但如果很多 goroutine 竞争一个 chan性能肯定是问题.....由于 golang 团队认识到这个问题普遍存在，为了避免大家重造车轮，因此官方统一出了一个包 Pool。但为什么放到 sync 包里面也是有的迷惑的，先不讨论这个问题。</p>
<p>先来看看如何使用一个 pool：</p>
<pre class="cnblogs_Highlighter brush:go;">package main  
  
import(  
    "fmt"  
    "sync"  
)  
  
func main() {  
    p := &amp;sync.Pool{  
        New: func() interface{} {  
            return 0  
        },  
    }  
  
    a := p.Get().(int)  
    p.Put(1)  
    b := p.Get().(int)  
    fmt.Println(a, b)  
}  </pre>
<p>上面创建了一个缓存 int 对象的一个 pool，先从池获取一个对象然后放进去一个对象再取出一个对象，程序的输出是 0 1。创建的时候可以指定一个 New 函数，获取对象的时候如何在池里面找不到缓存的对象将会使用指定的 new 函数创建一个返回，如果没有 new 函数则返回 nil。用法是不是很简单，我们这里就不多说，下面来说说我们关心的问题：</p>
<h3>1、缓存对象的数量和期限</h3>
<p>上面我们可以看到 pool 创建的时候是不能指定大小的，所有 sync.Pool 的缓存对象数量是没有限制的（只受限于内存），因此使用 sync.pool 是没办法做到控制缓存对象数量的个数的。另外 sync.pool 缓存对象的期限是很诡异的，先看一下 src/pkg/sync/pool.go 里面的一段实现代码：</p>
<pre class="cnblogs_Highlighter brush:go;">func init() {  
    runtime_registerPoolCleanup(poolCleanup)  
}  </pre>
<p>可以看到 pool 包在 init 的时候注册了一个 poolCleanup 函数，它会清除所有的 pool 里面的所有缓存的对象，该函数注册进去之后会在每次 gc 之前都会调用，因此 sync.Pool 缓存的期限只是两次 gc 之间这段时间。例如我们把上面的例子改成下面这样之后，输出的结果将是 0 0。正因 gc 的时候会清掉缓存对象，也不用担心 pool 会无限增大的问题。</p>
<pre class="cnblogs_Highlighter brush:go;">a := p.Get().(int)  
p.Put(1)  
runtime.GC()  
b := p.Get().(int)  
fmt.Println(a, b)  </pre>
<p>这是很多人错误理解的地方，正因为这样，我们是不可以使用sync.Pool去实现一个socket连接池的。</p>
<h3>2、缓存对象的开销</h3>
<p>如何在多个 goroutine 之间使用同一个 pool 做到高效呢？官方的做法就是尽量减少竞争，因为 sync.pool 为每个 P（对应 cpu，不了解的童鞋可以去看看 golang 的调度模型介绍）都分配了一个子池，如下图：</p>
<p><img src="https://images2015.cnblogs.com/blog/381128/201706/381128-20170625132332382-1488907925.png" alt="" /></p>
<p>当执行一个 pool 的 get 或者 put 操作的时候都会先把当前的 goroutine 固定到某个P的子池上面，然后再对该子池进行操作。每个子池里面有一个私有对象和共享列表对象，私有对象是只有对应的 P 能够访问，因为一个 P 同一时间只能执行一个 goroutine，因此对私有对象存取操作是不需要加锁的。共享列表是和其他 P 分享的，因此操作共享列表是需要加锁的。</p>
<p><span style="color: #0000ff;"><strong>获取对象过程是：</strong></span></p>
<p>1）固定到某个 P，尝试从私有对象获取，如果私有对象非空则返回该对象，并把私有对象置空；</p>
<p>2）如果私有对象是空的时候，就去当前子池的共享列表获取（需要加锁）；</p>
<p>3）如果当前子池的共享列表也是空的，那么就尝试去其他P的子池的共享列表偷取一个（需要加锁）；</p>
<p>4）如果其他子池都是空的，最后就用用户指定的 New 函数产生一个新的对象返回。</p>
<p>可以看到一次 get 操作最少 0 次加锁，最大 N（N 等于 MAXPROCS）次加锁。</p>
<p><span style="color: #0000ff;"><strong>归还对象的过程：</strong></span></p>
<p>1）固定到某个 P，如果私有对象为空则放到私有对象；</p>
<p>2）否则加入到该 P 子池的共享列表中（需要加锁）。</p>
<p>可以看到一次 put 操作最少 0 次加锁，最多 1 次加锁。</p>
<p>由于 goroutine 具体会分配到那个 P 执行是 golang 的协程调度系统决定的，因此在 MAXPROCS&gt;1 的情况下，多 goroutine 用同一个 sync.Pool 的话，各个 P 的子池之间缓存的对象是否平衡以及开销如何是没办法准确衡量的。但如果 goroutine 数目和缓存的对象数目远远大于 MAXPROCS 的话，概率上说应该是相对平衡的。</p>
<blockquote>
<p>总的来说，<span style="color: #ff0000;">sync.Pool</span> 的定位不是做类似连接池的东西，它的用途仅仅是增加对象重用的几率，减少 gc 的负担，而开销方面也不是很便宜的。</p>
</blockquote>
<dl id="btnDigg" onclick="btndigga();"></dl>
<p>&nbsp;</p>
<p>摘自：<a href="http://blog.csdn.net/yongjian_lian/article/details/42058893" target="_blank">http://blog.csdn.net/yongjian_lian/article/details/42058893</a></p></div><div id="MySignature"></div>
<div class="clear"></div>
<div id="blog_post_info_block">
<div id="BlogPostCategory"></div>
<div id="EntryTag"></div>
<div id="blog_post_info">
</div>
<div class="clear"></div>
<div id="post_next_prev"></div>
</div>


		</div>
		<div class = "postDesc">posted @ <span id="post-date">2017-06-25 13:22</span> <a href='https://www.cnblogs.com/52php/'>52php</a> 阅读(<span id="post_view_count">...</span>) 评论(<span id="post_comment_count">...</span>)  <a href ="https://i.cnblogs.com/EditPosts.aspx?postid=7076484" rel="nofollow">编辑</a> <a href="#" onclick="AddToWz(7076484);return false;">收藏</a></div>
	</div>
	<script type="text/javascript">var allowComments=true,cb_blogId=110339,cb_entryId=7076484,cb_blogApp=currentBlogApp,cb_blogUserGuid='384d7c4b-b863-e111-aa3f-842b2b196315',cb_entryCreatedDate='2017/6/25 13:22:00';loadViewCount(cb_entryId);var cb_postType=1;var isMarkdown=false;</script>
	
</div><!--end: topics 文章、评论容器-->
</div><a name="!comments"></a><div id="blog-comments-placeholder"></div><script type="text/javascript">var commentManager = new blogCommentManager();commentManager.renderComments(0);</script>
<div id='comment_form' class='commentform'>
<a name='commentform'></a>
<div id='divCommentShow'></div>
<div id='comment_nav'><span id='span_refresh_tips'></span><a href='javascript:void(0);' onclick='return RefreshCommentList();' id='lnk_RefreshComments' runat='server' clientidmode='Static'>刷新评论</a><a href='#' onclick='return RefreshPage();'>刷新页面</a><a href='#top'>返回顶部</a></div>
<div id='comment_form_container'></div>
<div class='ad_text_commentbox' id='ad_text_under_commentbox'></div>
<div id='ad_t2'></div>
<div id='opt_under_post'></div>
<script async='async' src='https://www.googletagservices.com/tag/js/gpt.js'></script>
<script>
  var googletag = googletag || {};
  googletag.cmd = googletag.cmd || [];
</script>
<script>
  googletag.cmd.push(function() {
        googletag.defineSlot('/1090369/C1', [300, 250], 'div-gpt-ad-1546353474406-0').addService(googletag.pubads());
        googletag.defineSlot('/1090369/C2', [468, 60], 'div-gpt-ad-1539008685004-0').addService(googletag.pubads());
        googletag.pubads().enableSingleRequest();
        googletag.enableServices();
  });
</script>
<div id='cnblogs_c1' class='c_ad_block'>
    <div id='div-gpt-ad-1546353474406-0' style='height:250px; width:300px;'></div>
</div>
<div id='under_post_news'></div>
<div id='cnblogs_c2' class='c_ad_block'>
    <div id='div-gpt-ad-1539008685004-0' style='height:60px; width:468px;'></div>
</div>
<div id='under_post_kb'></div>
<div id='HistoryToday' class='c_ad_block'></div>
<script type='text/javascript'>
 if(enablePostBottom()) {
    codeHighlight();
    fixPostBody();
    setTimeout(function () { incrementViewCount(cb_entryId); }, 50);
    deliverT2();
    deliverC1();
    deliverC2();    
    loadNewsAndKb();
    loadBlogSignature();
    LoadPostInfoBlock(cb_blogId, cb_entryId, cb_blogApp, cb_blogUserGuid);
    GetPrevNextPost(cb_entryId, cb_blogId, cb_entryCreatedDate, cb_postType);
    loadOptUnderPost();
    GetHistoryToday(cb_blogId, cb_blogApp, cb_entryCreatedDate);  
}
</script>
</div>

    
	</div><!--end: forFlow -->
	</div><!--end: mainContent 主体内容容器-->

	<div id="sideBar">
		<div id="sideBarMain">
			
<!--done-->
<div class="newsItem">
<h3 class="catListTitle">公告</h3>
	<div id="blog-news"></div><script type="text/javascript">loadBlogNews();</script>
</div>

			<div id="blog-calendar" style="display:none"></div><script type="text/javascript">loadBlogDefaultCalendar();</script>
			
			<div id="leftcontentcontainer">
				<div id="blog-sidecolumn"></div><script type="text/javascript">loadBlogSideColumn();</script>
			</div>
			
		</div><!--end: sideBarMain -->
	</div><!--end: sideBar 侧边栏容器 -->
	<div class="clear"></div>
	</div><!--end: main -->
	<div class="clear"></div>
	<div id="footer">
		
<!--done-->
Copyright &copy;2019 52php
	</div><!--end: footer -->
</div><!--end: home 自定义的最大容器 -->

<!--PageEndHtml Block Begin-->
<div class="back_to_top" onclick="window.scrollTo('0','0');" title="返回顶部"></div>
<script>
var _hmt = _hmt || [];
(function() {
  var hm = document.createElement("script");
  hm.src = "https://hm.baidu.com/hm.js?740d750ddb9677faa51e7896d9642ed7";
  var s = document.getElementsByTagName("script")[0]; 
  s.parentNode.insertBefore(hm, s);
})();
</script>
<!--PageEndHtml Block End-->
</body>
</html>
`
