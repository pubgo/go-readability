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
	"github.com/pkg/profile"
	"github.com/pubgo/errors"
	"github.com/pubgo/go-readability"
	"strings"
	"sync"
	"testing"
	"time"
)

// FromReader parses input from an `io.Reader` and returns the
// readable content. It's the wrapper for `Parser.Parse()` and useful
// if you only want to use the default parser.
func TestFromURL(t *testing.T) {
	defer errors.Debug()

	r := readability.FromURL("https://blog.csdn.net/alvine008/article/details/51282868", time.Second*3)
	//r := readability.FromURL("https://mp.weixin.qq.com/s?src=11&timestamp=1561273013&ver=1685&signature=wLwncgbNaHOci0d*RiI1L49pPp1aw-wXgYoCNQjLG2tyDyK92tGy-YRc1mNeZz1LstaGoWlF9exPzFapdqVkD*r-6HAEhCZ2QW7uj*Va79CaF6TzHPy5VgQhlfAeFqUa&new=1", time.Second*3)
	//r := readability.FromURL("https://www.baidu.com/link?url=sxfIBe7sLxoSFhGklac_DwnHsmjVfJbeq_UpJdA59JDOP2s2Svr5znv-K7m-Prwd&wd=&eqid=b42f79ff00072352000000045d0e3f9c",time.Second*3)
	errors.P(r.Copy())

}

func TestFromReader(t *testing.T) {
	defer errors.Debug()

	readability.Cfg.Debug = false

	// 开始性能分析, 返回一个停止接口
	stopper := profile.Start(profile.MemProfile, profile.ProfilePath("."))
	// 在main()结束时停止性能分析
	defer stopper.Stop()

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		//go func() {
		//	wg.Add(1)
			readability.FromReader(strings.NewReader(_html), "http://news.mydrivers.com/blog/20140212.htm")
			//wg.Done()
		//}()
		//errors.P(r.Copy())
	}
	wg.Wait()

}

const _html = `

<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=gb2312">
<title>新闻中心 ——驱动之家：您身边的电脑专家</title>
<link href="/Styles/newcss.css" rel="stylesheet" type="text/css">
<link href="/Styles/main1.css" rel=stylesheet>
<script src="/js/function.js" type=text/javascript></script>
</head>
<body bgcolor=#ffffff leftmargin=0 topmargin=0 marginwidth=0 marginheight=0>
<table width="923" height="6" border="0" align="center" cellpadding="0" cellspacing="0">
  <tr>
    <td></td>
  </tr>
</table>
<table width="923" border="0" align="center" cellpadding="0" cellspacing="0">
  <tr>
    <td width="107" height="46" valign="top"><a href="http://www.mydrivers.com"><img src="http://icons.mydrivers.com/www/mydrivers-logo.gif" width="107" height="46" border="0"></a></td>
    <td valign="top"><table width=808 border=0 cellpadding=0 cellspacing=0>
      <tr>

        <td width="85" height="24" align="center" background="http://icons.mydrivers.com/news/qjtxz_01.gif" class="f12_white">驱家通行证</td>
        <td width="705" height="24" background="http://icons.mydrivers.com/news/qjtxz_02.gif">
		<table width="100%" border="0" cellspacing="0" cellpadding="0">
          <tr><td>
		  <iframe src="http://passport.mydrivers.com/newslogin.htm" allowtransparency="true" style="background-color=transparent" scrolling="no" frameborder="0" align="left" height="24" width="320"></iframe>
          </td>
			<FORM name="myform" id="myform"  action="http://so.mydrivers.com/go.aspx" style="margin:0px; margin-top:2px;" onSubmit="return go_search_keyV2();" target="_blank">
			
			
              <td valign="middle" width="4%"><img align=\"middle\" src="http://icons.mydrivers.com/news/sousou.gif" width="26" height="24"></td>
			  <td width="15%"><input name="q" class="f12_black" id="q" value="关键字"  onfocus="this.value='';"  size="15"></td>

			  <td width="4%">
                <select name="s_class" style="width:85px">
                  <option value="drivers" selected="selected">驱动搜索</option>
                  <option value="google_news">新闻评测</option>
                  <option value="soft">软件</option>
                  <option value="tools">硬件工具</option>
                </select>

				</td>
			  <td valign="middle" width="5%">
			  <input type="submit" value="" style="background:url(http://icons.mydrivers.com/news/sousou11.gif) repeat;width:24px; height:24px; border:0px;cursor:pointer"></td>
            </form>
            <td width="21%" align="center" class="f12_white"><a href="http://www.mydrivers.com/contact/contactus.html" target="_blank"><font color="#FFFFFF">联系我们</font></a> | <a href="http://www.mydrivers.com/contact/cooperate.html" target="_blank"><font color="#FFFFFF">镜像合作</font></a></td>
          </tr>
        </table></td>

        <td width="18" height="24" background="http://icons.mydrivers.com/news/qjtxz_03.gif">&nbsp;</td>
      </tr>
    </table>
        <table width="100%" height="25" border="0" cellpadding="0" cellspacing="0">
          <tr>
            <td align="center"><span><a href="http://www.mydrivers.com/" class="newstiao5">首页</a></span> <span class="f12_red">|</span> <span><a href="http://news.mydrivers.com/" class="newstiao5">资讯</a> <a href="http://viewpoint.mydrivers.com/" class="newstiao5">视点</a> <a href="http://hardware.mydrivers.com/" class="newstiao5">评测</a></span> <span class="f12_red">|</span> <span><a href="http://product.mydrivers.com/information" class="newstiao5"></a> <a href="http://product.mydrivers.com/experience" class="newstiao5">产品体验</a></span> <span class="f12_red">|</span> <span><a href="http://drivers.mydrivers.com/" class="newstiao5">驱动中心</a> <a href="http://drivers.mydrivers.com/update" class="newstiao5">更新</a> <a href="http://so.mydrivers.com/" class="newstiao5">搜索</a> <a href="http://drivers.mydrivers.com/drivers" class="newstiao5">分类</a> <a href="http://drivers.mydrivers.com/search/" class="newstiao5">查询向导</a> <a href="http://drivers.mydrivers.com/subscribe" class="newstiao5">订阅向导</a> <a href="http://www.drivergenius.com/" target="_blank" class="newstiao5">驱动精灵</a><a href="http://drivers.mydrivers.com/install" class="newstiao5"></a></span> <span class="f12_red">|</span> <span><a href="http://www.myfiles.com.cn/" target="_blank" class="newstiao5">软件之家</a><a href="http://www.mypocket.com.cn/" class="newstiao5"></a> <a href="http://tools.mydrivers.com/" class="newstiao5">硬件工具</a></span> <span class="f12_red">|</span> <span class="style14"><a href="http://bbs.mydrivers.com" target="_blank" class="newstiao5">社区</a> <span class="f12_red">|</span> <a href="http://www.mydrivers.com/wap" target="_blank" class="newstiao5">手机版</a></span></td>

          </tr>
      </table></td>
  </tr>
</table>
<table width="100%" height="70" border="0" cellpadding="0" cellspacing="0" background="http://icons.mydrivers.com/news/news-lanmu_bg.gif">
  <tr>
    <td><table width=923 border=0 align="center" cellpadding=0 cellspacing=0>
        <tr>
          <td colspan=23><img src="http://icons.mydrivers.com/news/news-lanmu_01.gif" width=870 height=1></td>
        </tr>

		<TR>
        <TD><a href="/"><img src="http://icons.mydrivers.com/news/news-lanmu_02-02.gif" width="153" height="42" border="0" /></a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split01.gif" WIDTH=1 HEIGHT=42></TD>
        <TD align="center"><a href="http://news.mydrivers.com/class/1/"><img src="http://icons.mydrivers.com/news/t-t-9-1_01.gif" width="69" height="42" border="0" /></a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split01.gif" WIDTH=1 HEIGHT=42></TD>
        <TD align="center"><a href="http://news.mydrivers.com/class/2/"><img src="http://icons.mydrivers.com/news/t-t-9-1_03.gif" width="69" height="42" border="0" /></a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split01.gif" WIDTH=1 HEIGHT=42></TD>
        <TD align="center"><a href="http://news.mydrivers.com/class/3/"><img src="http://icons.mydrivers.com/news/t-t-9-1_05.gif" width="69" height="42" border="0" /></a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split01.gif" WIDTH=1 HEIGHT=42></TD>

		<TD align="center"><a href="http://news.mydrivers.com/class/4/"><img src="http://icons.mydrivers.com/news/t-t-9-1_07.gif" width="69" height="42" border="0"/></a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split01.gif" WIDTH=1 HEIGHT=42></TD>
        <TD align="center"><a href="http://news.mydrivers.com/class/5/"><img src="http://icons.mydrivers.com/news/t-t-9-1_09.gif" width="69" height="42" border="0" /></a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split01.gif" WIDTH=1 HEIGHT=42></TD>
        <TD align="center"><a href="http://news.mydrivers.com/class/6/"><img src="http://icons.mydrivers.com/news/t-t-9-1_11.gif" width="69" height="42" border="0" /></a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split01.gif" WIDTH=1 HEIGHT=42></TD>
        <TD align="center"><a href="http://news.mydrivers.com/class/7/"><img src="http://icons.mydrivers.com/news/t-t-9-1_13.gif" width="69" height="42" border="0" /></a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split01.gif" WIDTH=1 HEIGHT=42></TD>
        <TD align="center"><a href="http://news.mydrivers.com/class/8/"><img src="http://icons.mydrivers.com/news/t-t-9-1_15.gif" width="69" height="42" border="0" /></a></TD>

        <TD><IMG SRC="http://icons.mydrivers.com/news/split01.gif" WIDTH=1 HEIGHT=42></TD>
        <TD align="center"><a href="http://news.mydrivers.com/class/9/"><img src="http://icons.mydrivers.com/news/t-t-9-1_17.gif" width="69" height="42" border="0" /></a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split01.gif" WIDTH=1 HEIGHT=42></TD>
        <TD align="center"><a href="http://news.mydrivers.com/class/10/"><img src="http://icons.mydrivers.com/news/t-t-9-1_19.gif" width="69" height="42" border="0" /></a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split01.gif" WIDTH=1 HEIGHT=42></TD>
        <TD width="60" align="center" valign="bottom"><table width="100%" border="0" cellspacing="0" cellpadding="0">
          <tr>
            <td height="18" align="center"><a href="http://news.mydrivers.com/" target="_blank" class="newstiao6">Blog版本</a></td>

          </tr>
          <tr>
            <td height="18" align="center"><a href="http://news.mydrivers.com/doc/" target="_blank" class="newstiao2">文章索引</a></td>
          </tr>
        </table>          </TD>
      </TR>
      <TR>
        <TD><a href="/"><IMG SRC="http://icons.mydrivers.com/news/news-lanmu_23-02.gif" WIDTH=153 HEIGHT=27 border="0"></a></TD>

        <TD><IMG SRC="http://icons.mydrivers.com/news/split02.gif" WIDTH=1 HEIGHT=27></TD>
        <TD width="69" height="27" align="center"><a href="http://news.mydrivers.com/class/1/" class="daohang">核心硬件</a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split02.gif" WIDTH=1 HEIGHT=27></TD>
        <TD width="69" height="27" align="center"><a href="http://news.mydrivers.com/class/2/" class="daohang">常用配件</a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split02.gif" WIDTH=1 HEIGHT=27></TD>
        <TD width="69" height="27" align="center"><a href="http://news.mydrivers.com/class/3/" class="daohang">掌上设备</a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split02.gif" WIDTH=1 HEIGHT=27></TD>

        <TD width="69" height="27" align="center"><a href="http://news.mydrivers.com/class/4/" class="daohang">数码影音</a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split02.gif" WIDTH=1 HEIGHT=27></TD>
        <TD width="69" height="27" align="center"><a href="http://news.mydrivers.com/class/5/" class="daohang">网络设备</a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split02.gif" WIDTH=1 HEIGHT=27></TD>
        <TD width="69" height="27" align="center"><a href="http://news.mydrivers.com/class/6/" class="daohang">办公外设</a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split02.gif" WIDTH=1 HEIGHT=27></TD>
        <TD width="69" height="27" align="center"><a href="http://news.mydrivers.com/class/7/" class="daohang">游戏相关</a></TD>

        <TD><IMG SRC="http://icons.mydrivers.com/news/split02.gif" WIDTH=1 HEIGHT=27></TD>
        <TD width="69" height="27" align="center"><a href="http://news.mydrivers.com/class/8/" class="daohang">OS/软件</a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split02.gif" WIDTH=1 HEIGHT=27></TD>
        <TD width="69" height="27" align="center"><a href="http://news.mydrivers.com/class/9/" class="daohang">业界科学</a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split02.gif" WIDTH=1 HEIGHT=27></TD>
        <TD width="69" align="center"><a href="http://news.mydrivers.com/class/10/" class="daohang">品牌整机</a></TD>
        <TD><IMG SRC="http://icons.mydrivers.com/news/split02.gif" WIDTH=1 HEIGHT=27></TD>

        <TD width="75" align="center"><a href="http://rss.mydrivers.com/"><img src="http://icons.mydrivers.com/news/rss-logo_03.gif" width="30" height="11" border="0"></a></TD>
      </TR>
      </table></td>
  </tr>
</table>
<table width="923" height="3" border="0" align="center" cellpadding="0" cellspacing="0">
  <tr>
    <td></td>
  </tr>
</table>
<table width="923" border="0" align="center" cellpadding="0" cellspacing="1" bgcolor="#969696">
  <tr>
    <td height="25" align="center" bgcolor="#E8E8E8"><table width="98%" border="0" align="center" cellpadding="0" cellspacing="0">
        <tr>
          <td width="68%" align="left"><span class="f12_blue1">当前位置：<a href="/">新闻中心</a> &gt; 文章索引</td>
          <td width="32%" align="right"><table width="100%" border="0" cellspacing="0" cellpadding="0"><form name="findnews"  action="http://so.mydrivers.com/news.aspx" target="_blank">
                    <tr>
                      <td align="center"><span class="f12_red1">关键字</span>
                        <input name="q" type="text" class="f14black" size="20">
	        <span class="f14"><a href="#" class="f12_black" onClick="go_search()">搜索</a></span></td>
                    </tr></form>
                  </table>
</td>
        </tr>
      </table></td>
  </tr>
</table>
<table width="923" height="5" border="0" align="center" cellpadding="0" cellspacing="0">
  <tr>
    <td><img src="/images/spacer.gif" width="1" height="1"></td>
  </tr>
</table>
<!--新闻索引(视点文章)开始-->
<table width="923" border="0" align="center" cellpadding="0" cellspacing="1" bgcolor="#CCCCCC">
<tr>

<td height="140" width="154" align="center" bgcolor="#F5F5F5">
<table width="98%" height="70" border="0" cellpadding="0" cellspacing="0">
          <tr>
            <td align="center" valign="middle"><A href="/1/525/525713.htm" target=_blank>
		<IMG height=90 src="http://news.mydrivers.com/Img/20170330/title_004625622.jpg" width=120 border="0">
	</A></td>
          </tr>
          <tr>
            <td align="center"><span class="STYLE3"><A href="/1/525/525713.htm" target=_blank class="newstiao3">三星S8/S8+官方价格公布！顿时掩面而走</A></span></td>
          </tr>
</table>
</td>

<td height="140" width="154" align="center" bgcolor="#F5F5F5">
<table width="98%" height="70" border="0" cellpadding="0" cellspacing="0">
          <tr>
            <td align="center" valign="middle"><A href="/1/525/525623.htm" target=_blank>
		<IMG height=90 src="http://news.mydrivers.com/Img/20170329/title_184824413.jpg" width=120 border="0">
	</A></td>
          </tr>
          <tr>
            <td align="center"><span class="STYLE3"><A href="/1/525/525623.htm" target=_blank class="newstiao3">华为P10屏幕上没有的疏油层到底是什么？</A></span></td>
          </tr>
</table>
</td>

<td height="140" width="154" align="center" bgcolor="#F5F5F5">
<table width="98%" height="70" border="0" cellpadding="0" cellspacing="0">
          <tr>
            <td align="center" valign="middle"><A href="/1/525/525616.htm" target=_blank>
		<IMG height=90 src="http://news.mydrivers.com/Img/20170329/title_184446528.jpg" width=120 border="0">
	</A></td>
          </tr>
          <tr>
            <td align="center"><span class="STYLE3"><A href="/1/525/525616.htm" target=_blank class="newstiao3">独享6GB内存！三星S8国行亮相：全球最强</A></span></td>
          </tr>
</table>
</td>

<td height="140" width="154" align="center" bgcolor="#F5F5F5">
<table width="98%" height="70" border="0" cellpadding="0" cellspacing="0">
          <tr>
            <td align="center" valign="middle"><A href="/1/525/525610.htm" target=_blank>
		<IMG height=90 src="http://news.mydrivers.com/Img/20170329/title_184540613.jpg" width=120 border="0">
	</A></td>
          </tr>
          <tr>
            <td align="center"><span class="STYLE3"><A href="/1/525/525610.htm" target=_blank class="newstiao3">买吗？三星翻新版Note 7售价曝光：超便宜</A></span></td>
          </tr>
</table>
</td>

<td height="140" width="154" align="center" bgcolor="#F5F5F5">
<table width="98%" height="70" border="0" cellpadding="0" cellspacing="0">
          <tr>
            <td align="center" valign="middle"><A href="/1/525/525576.htm" target=_blank>
		<IMG height=90 src="http://news.mydrivers.com/Img/20170329/title_184735492.jpg" width=120 border="0">
	</A></td>
          </tr>
          <tr>
            <td align="center"><span class="STYLE3"><A href="/1/525/525576.htm" target=_blank class="newstiao3">Excel重磅新功能上线：超赞！</A></span></td>
          </tr>
</table>
</td>

<td height="140" width="154" align="center" bgcolor="#F5F5F5">
<table width="98%" height="70" border="0" cellpadding="0" cellspacing="0">
          <tr>
            <td align="center" valign="middle"><A href="/1/525/525567.htm" target=_blank>
		<IMG height=90 src="http://news.mydrivers.com/Img/20170329/title_184644620.jpg" width=120 border="0">
	</A></td>
          </tr>
          <tr>
            <td align="center"><span class="STYLE3"><A href="/1/525/525567.htm" target=_blank class="newstiao3">提速降功耗：Intel神级CPU归来！只可惜…</A></span></td>
          </tr>
</table>
</td>

</tr>
</table>
<!--新闻索引(视点文章)结束-->
<table width="923" height="5" border="0" align="center" cellpadding="0" cellspacing="0">
  <tr>
    <td><img src="/images/spacer.gif" width="1" height="1"></td>
  </tr>
</table>
<table width="923" border="0" align="center" cellpadding="0" cellspacing="0" bordercolor="#111111" style="border-collapse: collapse">
  <tr>
    <td width="100%"><table border="1" cellpadding="0" cellspacing="0" style="border-collapse: collapse" bordercolor="#CCCCCC" width="100%" id="AutoNumber15">
        <tr>
          <td width="50%" valign="top"><table border="0" cellpadding="0" cellspacing="0" style="border-collapse: collapse" bordercolor="#111111" width="100%" id="AutoNumber36" height="26">
              <tr>
                <td width="3%" align="left" background="/images/l-m-t-bg.gif"><span class="left1"><img src="/images/l-m-t-01_temp.gif" width="18" height="22" border="0" /></span></td>
                <td width="51%" align="left" background="/images/l-m-t-bg.gif"><img border="0" src="http://www.mydrivers.com/images/tiao-123.gif" width="141" height="16" /></td>
                <td width="46%" align="left" background="/images/l-m-t-bg.gif">&nbsp;</td>
              </tr>
            </table>
            <table border="0" width="98%" cellpadding="2" cellspacing="0">
  <tr>
    <td align="center" width="26%"><table width="100%" border="0" align="center" cellpadding="3" cellspacing="0">
        
        <tr>
          <td><a href="/1/521/521372.htm" target=_blank><img height=70 src="/Img/20170228/title_202825735.jpg" width=90 border="0"></a></td>
        </tr>
        
        <tr>
          <td><a href="/1/523/523508.htm" target=_blank><img height=70 src="/Img/20170315/title_194016462.jpg" width=90 border="0"></a></td>
        </tr>
        
      </table></td>
    <td width="74%"><table width="100%" border="0" align="center" cellpadding="3" cellspacing="0">
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/521/521800.htm" class="newstiao4">完美逆袭Intel！AMD Ryzen全球首发评测</a></td>
          <td width="40"><span class="f12_date">(03-02)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/524/524482.htm" class="newstiao4">华为新旗舰P10发布会直播！评论送手机</a></td>
          <td width="40"><span class="f12_date">(03-21)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/522/522808.htm" class="newstiao4">地球第一神卡！GTX 1080 Ti首发评测：..</a></td>
          <td width="40"><span class="f12_date">(03-09)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/522/522460.htm" class="newstiao4">徕卡双摄拍照逆天 华为P10深度评测</a></td>
          <td width="40"><span class="f12_date">(03-07)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/521/521372.htm" class="newstiao4">AMD Zen反复测试后彻底信了！岂止翻身..</a></td>
          <td width="40"><span class="f12_date">(02-28)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/523/523508.htm" class="newstiao4">全球最轻15寸笔记本评测：续航彻底服！</a></td>
          <td width="40"><span class="f12_date">(03-15)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/525/525090.htm" class="newstiao4">15岁男孩私处异常 医生检查后大吃一惊...</a></td>
          <td width="40"><span class="f12_date">(03-25)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/521/521964.htm" class="newstiao4">AMD Ryzen日本首卖：Intel被这一幕吓..</a></td>
          <td width="40"><span class="f12_date">(03-04)</span></td>
        </tr>
        
      </table></td>
  </tr>
</table>
          </td>
          <td width="50%" valign="top"><table border="0" cellpadding="0" cellspacing="0" style="border-collapse: collapse" bordercolor="#111111" width="100%" id="AutoNumber37" height="26">
              <tr>
                <td width="4%" align="left" background="/images/l-m-t-bg.gif" style="font-size: 9pt"><span class="left1"><img src="/images/l-m-t-01_temp.gif" width="18" height="22" border="0" /></span></td>
                <td width="54%" align="left" background="/images/l-m-t-bg.gif" style="font-size: 9pt"><img border="0" src="http://www.mydrivers.com/images/tiao-124.gif" width="170" height="16" /></td>
                <td width="42%" align="left" background="/images/l-m-t-bg.gif" style="font-size: 9pt">&nbsp;</td>
              </tr>
            </table>
            <table border="0" width="98%" cellpadding="2" cellspacing="0">
  <tr>
    <td align="center" width="26%"><table width="100%" border="0" align="center" cellpadding="3" cellspacing="0">
        
        <tr>
          <td><a href="/1/523/523065.htm" target=_blank><img height=70 src="/Img/20170312/title_225048365.jpg" width=90 border="0"></a></td>
        </tr>
        
        <tr>
          <td><a href="/1/521/521372.htm" target=_blank><img height=70 src="/Img/20170228/title_202825735.jpg" width=90 border="0"></a></td>
        </tr>
        
      </table></td>
    <td width="74%"><table width="100%" border="0" align="center" cellpadding="3" cellspacing="0">
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/524/524482.htm" class="newstiao4">华为新旗舰P10发布会直播！评论送手机</a></td>
          <td width="40"><span class="f12_date">(03-21)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/521/521382.htm" class="newstiao4">重磅 小米松果处理器发布直播 送199好..</a></td>
          <td width="40"><span class="f12_date">(02-28)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/523/523043.htm" class="newstiao4">终于不用输入该死的验证码了！喜大普奔</a></td>
          <td width="40"><span class="f12_date">(03-11)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/521/521694.htm" class="newstiao4">曾炒到2万！华为Mate 9保时捷版售价暴..</a></td>
          <td width="40"><span class="f12_date">(03-02)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/521/521800.htm" class="newstiao4">完美逆袭Intel！AMD Ryzen全球首发评测</a></td>
          <td width="40"><span class="f12_date">(03-02)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/522/522958.htm" class="newstiao4">迪士尼最经典歌曲：《Let it go》惨输</a></td>
          <td width="40"><span class="f12_date">(03-10)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/525/525123.htm" class="newstiao4">AMD Ryzen 7战Intel酷睿i3：结局竟这..</a></td>
          <td width="40"><span class="f12_date">(03-26)</span></td>
        </tr>
        
        <tr>
          <td width="10" height="16" align="center">·</td>
          <td><a href="/1/523/523065.htm" class="newstiao4">同样四核4GHz Ryzen大战i7：竟是这样！</a></td>
          <td width="40"><span class="f12_date">(03-12)</span></td>
        </tr>
        
      </table></td>
  </tr>
</table>
          </td>
        </tr>
      </table></td>
  </tr>
</table>
<table width="923" height="5" border="0" align="center" cellpadding="0" cellspacing="0">
  <tr>
    <td><img src="/images/spacer.gif" width="1" height="1"></td>
  </tr>
</table>
<table width="923" border="0" align="center" cellpadding="0" cellspacing="0" bordercolor="#888888" id="AutoNumber3" style="border-collapse: collapse">
  <tr>
    <td width="100%"><table width="100%" border="0" cellpadding="0" cellspacing="0" id="AutoNumber4" style="border-collapse: collapse">
        <tr>
          <td width="2%" height="25" align="left" valign="middle" background="/images/l-m-t-bg.gif"><span class="left1"><img src="/images/l-m-t-01_temp.gif" width="18" height="22" border="0" /></span></td>
          <td width="48%" align="left" valign="middle" background="/images/l-m-t-bg.gif"><img border="0" src="/images/tiao-124-00000.gif" width="170" height="16" /></td>
          <td width="50%" align="left" valign="middle" background="/images/l-m-t-bg.gif">&nbsp;</td>
        </tr>
      </table>
      <table border="1" cellpadding="0" cellspacing="0" width="923" style="border-collapse: collapse" bordercolor="#BBBBBB">
        <tr>
          <td bgcolor="#F2F2F2" height="8" width="100%" style="font-size: 9pt">　
       
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292523.htm" target="_blank" title="囧 买个鸡肉卷也要上众筹网站？"><span class="title1">囧 买个鸡肉卷也要上众筹网站？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">23:17:14</span></td>
  </tr>
  <tr>
    <td align="left"><p>说起Kickstarter很多人都不陌生，一大批有才华的创意都曾通过这个众筹网站变成了产品，而随着众筹概念的流行越来越多的人会把他们的大小创意放上去集资来实现自己的想法。</p>
<p>不过你听说过想吃个鸡肉卷也要写一堆介绍并且制作出视频和图片放Kickstarter上筹钱吗？</p>
<p>前一阵一个叫做Noboru Bitoy的囧人为了买一个墨西哥鸡肉卷吃就在Kickstarter上发起了筹资活动，至于他为什这样做已经无从考证了，不过发起筹资以后他立刻就收到了别人的出资，并达到了8美元的筹资目标。</p>
<p>然而达到目标以后居然还有人不断给他捐助，于是他修改了筹资额度上限到193美元，并且延长了期限，规则方面也做了调整。</p>
<p><strong>当筹资额达到50美元时他承诺会买下4个墨西哥鸡肉卷然后接下来4天里每天吃下一个来验证时间变化对鸡肉卷味道的影响。</strong></p>
<p><strong>当达到175美元时他会品尝用24种不同材料面皮做出的鸡肉卷的味道，并把感受做成图表&hellip;&hellip;</strong></p>
<p>看来吃货的世界，别人真不懂。</p>
<p align="center"><a href="/img/20140212/28bc9535416a42aebf10521ccf84dd94.jpg" target="_blank"><img alt="囧 买个鸡肉卷也要上众筹网站？" src="/img/20140212/s_28bc9535416a42aebf10521ccf84dd94.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/50a6ea35403e461ab9abdfb3a0df27f8.jpg" target="_blank"><img alt="囧 买个鸡肉卷也要上众筹网站？" src="/img/20140212/s_50a6ea35403e461ab9abdfb3a0df27f8.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/7ea2a774463b49a39988efc232e1c2c1.jpg" target="_blank"><img alt="囧 买个鸡肉卷也要上众筹网站？" src="/img/20140212/s_7ea2a774463b49a39988efc232e1c2c1.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292522.htm" target="_blank" title="电竞比赛真的需要那男女分开吗？"><span class="title1">电竞比赛真的需要那男女分开吗？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">22:29:56</span></td>
  </tr>
  <tr>
    <td align="left"><p>女性向来是电子竞技甚至游戏领域的稀有动物，无论是选手、玩家还是游戏开发者，女性的比例都非常的稀少。在可以预见的将来，女性会更多的参与到电竞热潮中来，那么到时会简单的出现专属女性的比赛，还是女性也能和男性同台竞技呢？</p>
<p>一名Reddit在电竞版块提问说，有哪个项目是男女可以混合组队的？那么今天我们就来讨论一下性别的问题。</p>
<p><strong>体育因身体的差异而分男女，电竞或许不用如此？</strong></p>
<p>传统的体育运动，由于男女身体结构的不同，特别是男性天生就在体力方面占有优势，因此在比赛项目中区分性别尤为重要，否则比赛中你就看不到女性选手的身影。</p>
<p>至于在传统的智力运动中，尽管围棋和国际象棋会专门设有女子比赛，但除此之外其它的比赛并没有性别限制，男性能参加的比赛，女性也能参加。至于为什么很少在世界级的围棋或国际象棋比赛中看到女性的身影，并不是因为有性别的限制，而是总体来看女性选手的实力弱于男性。不过，在我们看来，<strong>在这些智力的比拼中女性弱于男性并不是性别的原因，而是本身女性选手的人群基数就远少于男性，因此出现顶尖选手的概率相对更小。</strong></p>
<p>值得一提的是，近年来颇为流行的德州扑克，女性选手与男性相比毫不逊色。</p>
<p>如果说，当前电子竞技阴衰阳盛是因为女性用户基数太小，那么未来呢？倘若有一天女性选手的数量增多并且实力不逊于男性选手，你愿意看到专门的女子比赛还是愿意看到男性与女性同台竞技呢？</p>
<p><strong>《星际2》现在还有一位女性选手Scarlett（严格说是变性人），</strong>曾经在韩国也有真正的女性选手不过成绩都较一般，可能为中国玩家所熟知的TossGirl算是最厉害的一个，曾经击败过中国著名虫族选手F91。但这些远远不够，<strong>像《英雄联盟》和《DOTA2》这样的团队竞技游戏从来没有女性独挡一面的先例。</strong></p>
<p>客观地看，女性在玩游戏这个问题上与男性相比并没有先天的劣势，现在的问题是<strong>根本就没多少女性像男性那样热衷于游戏，更不要说付出大量的时间和精力去成为职业选手</strong>。那么可以相信未来有可能看到许多非常厉害的女性选手。</p>
<p>未来最可能的是像围棋那样，有专门的女子比赛但无专门的男子比赛。不过，你最希望看到什么，男女同台还是男女分离？</p>
<p align="center"><a href="/img/20140212/fdfdf56096394fe7b54b29ad8b66ef42.jpg" target="_blank"><img alt="电竞比赛真的需要那男女分开吗？" src="/img/20140212/s_fdfdf56096394fe7b54b29ad8b66ef42.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292521.htm" target="_blank" title="为何说韩国仍走在互联网的末端"><span class="title1">为何说韩国仍走在互联网的末端</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">22:24:37</span></td>
  </tr>
  <tr>
    <td align="left"><p>韩国喜欢视自身为世界互联网领袖，它拥有世界上最快的平均宽带速度（每秒约22M），上月政府宣布到2020年将该国的无线网络升级至5G，使得下载速度比当前快1000倍。它的互联网普及率也属于世界最高的行列，新兴社区欣欣向荣（Cyworld，推出于扎克伯格创办Facebook前五年，蝉联十年国内最受欢迎的社交网络），此外，韩国还是世界上作为观赏性体育的电子竞技最为发达的国家。然而，在其它方面，这一未来主义的国家（带有未来主义色彩的国家）依然陷在黑暗的时代里。</p>
<p>去年，<strong>一个美国非政府组织（NGO）自由之家（Freedom House）将韩国的互联网评为&ldquo;部分自由&rdquo;。</strong>无国界记者组织（Reporters without Borders）在其《互联网之敌》的报告中，将它与埃及，泰国和俄罗斯一道列入&ldquo;被监督&rdquo;的国家名单。前瞻性的韩国是否真的这样落后？</p>
<p>每一周，政府审查员会删除互联网上大量的内容。去年，在韩国通信标准委员会（KCSC）&mdash;&mdash;一家名义上独立，但主要由政府任命的公共机构&mdash;的要求之下，约23000个韩国网站被删除，另外63000个被屏蔽。</p>
<p><strong>2009年KCSC仅发出4500条删除要求，过滤的主要对象是对定为非法的色情作品，卖淫活动和赌博。但是更多健康的追求也受到限制</strong>：16岁以下的未成年人不得在午夜到6点之间玩网络游戏（用户必须要输入由政府签发的身份证号码来证明自己的年龄）；来自朝鲜的网站，包括它的国家报纸，新闻机构和Twitter状态以及同情朝鲜者的网站都被屏蔽。韩国在朝鲜战争时颁布一部法律，禁止将本国地图带出境，由于严格来说，朝韩两国尚未止争，这项法律目前范围扩大，电子地图数据也被纳入其中&mdash;这意味着<strong>Google不能在服务器上处理韩国的地图数据，因此无法在该国提供驾驶指南。</strong></p>
<p>2010年联合国认定，本质上而言，KSCS是一个审查机构。一些韩国人正在进行抵抗。2011年，一位持异见的委员Park Kyung-sin在博客上发布了一张法国画家居斯塔夫&middot;库尔贝的画：《世界起源》，以抗议KSCS下令屏蔽一张他此前在同一博客上放的男性生殖器图片&mdash;就是那种在科学教材上找到的图片--的行为。他因此获罪并被罚款，尽管指控后被撤销。</p>
<p>2012年一位15岁的<strong>韩国网络游戏冠军在半夜之后参加星际争霸二的比赛（比赛于法国举行，当时正是白天）时被强行弹出游戏，当他用父亲的身份证号再次连接时，比赛已经输了</strong>。监管机构也没有幽默感：2012年，一位摄影师转发了一系列朝鲜宣传帖，并把自己继承父亲的摄影棚与朝鲜的领导人换届进行比较，他因而被判处十个月的有期徒刑，缓期执行。另一位于2008年发布雷曼破产，韩元贬值消息的博主，Park Dae-sung因散播错误谣言而在监督里呆了104天。</p>
<p><strong>2004年大选前夕，在网上发表政治评论的用户需要输入名字和身份证号码</strong>。2009年，在每日游客超过100000的网站上发表评论的人也要遵守同样的要求，这一法律已经被撤销，但是尽管政府开始放松一些限制，它加强了对社交媒体的监督。</p>
<p>2011年，KCSC成立了一个社交媒体特别委员会，2012年，它要求Twitter、Facebook等社交媒体删除4500条评论&mdash;是10年的14倍。去年，被删除的评论数量再度增长，达到了6400条。一些官员似乎享受发布假评论，删除真评论。</p>
<p>目前，一些情报人员因在2012年大选前夕，使用假身份支持现任总统Park Geun-hye而接受调查，没有证据表明这一命令出自朴槿惠。12月，朴表示，政府需要&ldquo;纠正大量通过社交网络散播的谣言&rdquo;，暗指公众对于铁路和医保私有化的强烈抗议。韩国人或许可以享受不同寻常的快速网络连接，但是却不能自由使用。</p>
<p align="center"><img alt="为何说韩国仍走在互联网的末端" src="/img/20140212/38665eaa6aab4783997971133887be84.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292520.htm" target="_blank" title="迪拜警察又买新超跑了 布加迪威龙"><span class="title1">迪拜警察又买新超跑了 布加迪威龙</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">22:18:53</span></td>
  </tr>
  <tr>
    <td align="left"><p>据中新网报道，本来就豪车满仓的迪拜警察最近又有新超跑入手了，那就是布加迪威龙。</p>
<p>据了解这辆布加迪威龙是迪拜警方定制款，搭载8.0L W16四涡轮增压发动机，极速可达四百公里以上。</p>
<p>在此之前迪拜警察就以豪车云集而闻名，奥迪R8、阿斯顿&middot;马丁One-77、兰博基尼Abentador、法拉利FF、宾利欧陆GT、奔驰SLS AMG等豪华超跑都在其名下，而本次加入布加迪威龙为其豪车名单又添一员。</p>
<p>不过布加迪威龙这么贵的车就连迪拜警察也舍不得用来抓人，据迪拜警方负责人介绍，这辆车并不用来逮捕罪犯，而是主要在旅游区进行巡逻供游客拍照。</p>
<p align="center"><a href="/img/20140212/be831fc1bc0c401d89569f9cdc857dad.jpg" target="_blank"><img alt="迪拜警察又买新超跑了 布加迪威龙" src="/img/20140212/s_be831fc1bc0c401d89569f9cdc857dad.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/c7e570a3c5bc4ea2a47b779e97ec7878.jpg" target="_blank"><img alt="迪拜警察又买新超跑了 布加迪威龙" src="/img/20140212/s_c7e570a3c5bc4ea2a47b779e97ec7878.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/941c1addf8e74fc5a269a4388d183c22.jpg" target="_blank"><img alt="迪拜警察又买新超跑了 布加迪威龙" src="/img/20140212/s_941c1addf8e74fc5a269a4388d183c22.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/50195f7af15e46dfa15062bbaaaccafd.jpg" target="_blank"><img alt="迪拜警察又买新超跑了 布加迪威龙" src="/img/20140212/s_50195f7af15e46dfa15062bbaaaccafd.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/8c08e23cd6e94787ba8a39c6d56937a3.jpg" target="_blank"><img alt="迪拜警察又买新超跑了 布加迪威龙" src="/img/20140212/s_8c08e23cd6e94787ba8a39c6d56937a3.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/b8c88be2ac15490f9652bb677883a210.jpg" target="_blank"><img alt="迪拜警察又买新超跑了 布加迪威龙" src="/img/20140212/s_b8c88be2ac15490f9652bb677883a210.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/205edf7a9bc640dbae3ed1da9d445a03.jpg" target="_blank"><img alt="迪拜警察又买新超跑了 布加迪威龙" src="/img/20140212/s_205edf7a9bc640dbae3ed1da9d445a03.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>&nbsp;</p>
<p>&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292517.htm" target="_blank" title="爆料：《使命召唤11》画质将大幅超越前作"><span class="title1">爆料：《使命召唤11》画质将大幅超越前作</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">21:33:53</span></td>
  </tr>
  <tr>
    <td align="left"><p>官方已经确认《使命召唤11》将由&ldquo;大锤&rdquo;Sledgehammer Games打造，<strong>日前业界线人爆料称，该作的画质水平将超越《使命召唤10：幽灵》，同时也是系列第一款以&ldquo;次世代&rdquo;为标准打造的作品。</strong></p>
<p>《使命召唤10》的画面不是很棒，主要是和其他厂商一样，还要顾及到现役主机，次世代非优先平台，本次的COD11是真正的次世代优先作品。</p>
<p>据悉，&ldquo;<strong>大锤&rdquo;Sledgehammer Games，Treyarch，IW将轮流为COD系列打造作品，据称，《使命召唤11》的剧情采用了现代战争的世界观，其中一些角色可能加盟。</strong></p>
<p>根据某不愿透露姓名的内部人士的爆料来看，正在开发中的《使命召唤11》是&ldquo;现代战争&rdquo;系列的一部续作，并且会照例在今年发售。</p>
<p align="center"><a href="/img/20140212/c0b8132780e14f8cb6b63a8b61ac3b1a.jpg" target="_blank"><img alt="爆料：《使命召唤11》画质将大幅超越前作" src="/img/20140212/s_c0b8132780e14f8cb6b63a8b61ac3b1a.jpg" style="border: 1px solid black;" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292516.htm" target="_blank" title="中文版Galaxy S5截图曝光：彻底扁平化"><span class="title1">中文版Galaxy S5截图曝光：彻底扁平化</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">21:21:37</span></td>
  </tr>
  <tr>
    <td align="left"><p>三星邀请函显示，Galaxy S5将会在本月24日亮相巴塞罗那的MWC 2014会场，而随着发布之日临近，法国媒体又曝出了一张中文版的Galaxy S5 UI截图。</p>
<p>这张新谍照为中文版的设置界面截图，其中的&ldquo;附近的设备&rdquo;以及&ldquo;位置&rdquo;图标都与之前<a class="f14_link" href="http://news.mydrivers.com/1/292/292347.htm" target="_blank">英文版预告</a>上出现的相同。从这些图标和界面风格可以看到，<strong>Galaxy S5采用的全新版TouchWiz UI是确定要将扁平化风格进行到底了。</strong></p>
<p>除了UI，韩国媒体The Korea Herald还透露，<strong>三星将在Galaxy S5中使用接近于&ldquo;零边框&rdquo;的超窄边设计，并引入指纹识别功能</strong>。</p>
<p>比较令人惊讶的是，<strong>该网站还爆料称S5将去掉盖世系列标志性的Home按键，改为虚拟按键</strong>。不过从目前来看，三星对实体按键极为热衷，从Galaxy S开始，三星众多Android甚至是Windows Phone手机都无一例外采用了标志性的椭圆形Home按键，而且目前曝光的多数谍照显示，<a class="f14_link" href="http://news.mydrivers.com/1/292/292362.htm" target="_blank">S5的三星风格依旧</a>，因此贸然取消的可能性不大。</p>
<p>Galaxy S5发布在即，据<a class="f14_link" href="http://news.mydrivers.com/1/292/292352.htm" target="_blank">昨天曝光的包装盒</a>显示信息，它配置十分出色，<strong>采用5.25寸Super AMOLED材质的2K屏幕（560ppi），搭载2.5GHz骁龙MSM8974AC处理器，内置3GB内存和3000mAh容量电池，配备2000万像素主摄像头，支持4G网络</strong>，运行最新的Android 4.4以及TouchWiz UI。当然，按照三星一贯风格，该机必然还会有多个变种版本。</p>
<p align="center"><img alt="中文版Galaxy S5截图曝光：彻底扁平化" src="/img/20140212/ff7b8835899a45c880523447a68ef8e7.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292515.htm" target="_blank" title="奇葩物 装在摩托车上的空调"><span class="title1">奇葩物 装在摩托车上的空调</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">21:21:32</span></td>
  </tr>
  <tr>
    <td align="left"><p>EntroSys家的BikeAir是一款摩托车空调，给摩托车装空调已经够奇葩了，现在我们想看看它是否真的有用。</p>
<p>2010年时这款产品还在开发当中，现在它已经在欧美日澳上市了。BikeAir由紧凑型空调组件和穿在摩托车夹克下面的冷却背心组成。电力由摩托车的电池提供，有三种冷热模式及三档风速，都可以通过一个可固定在车把上的小装置来无线控制。</p>
<p>作为空调，整个装置很小很轻，只有4.56千克。由于使用了固态热电系统，耗电量也很低，转动的部件只有风扇。这种方式使得它可以持久地在摩托车或是敞篷车上运行。</p>
<p><strong>安装</strong></p>
<p><strong>整个装置包含一个遥控器、背心和空调，以及用于连接到摩托车电池的电线，用于连接背心和空调的空气管，还有橡皮管帮助空气更好地散布，以及一个底座。</strong></p>
<p><strong>这个产品是为重型摩托车设计的，因为只有它们才能足够的电量和安装空间</strong>。安装起来并不难，把装置固定好，再把线和管子连接好就行了。Entrosys推荐让经销商来安装。</p>
<p><strong>上路</strong></p>
<p><strong>在室外温度高达35度时，在皮夹克下面穿上空调背心，接上空气管，骑着摩托车在路上，直接把空调开到最高档，很快就清凉了，夏日穿皮夹克的难受感一扫而光</strong>。这是gizmag作者在澳大利亚的体验，那里现在正值夏天。</p>
<p>空调背心和橡皮管的设计让冷气可以很好地输送到身体表面，背心并不重，穿起来很舒服。Entrosys的建议是穿件T恤，这大热天穿夹克骑摩托吹空调，估计也会被人认为是神经病吧。</p>
<p>装在摩托车上的整个设备不会让你觉得载了一个空调这么夸张，空气管也够长，骑车时偶尔站起来不成问题。</p>
<p>有空调的一个好处是，你可以把整个保护装备都穿起来，而不会觉得热。</p>
<p>遥控器很容易使用，有红蓝两色的灯指示冷热，及三档风扇设置，把它绑在车把上用起来最方便。</p>
<p>作为空调，除了制冷，它也能制热，在冬季或是晚上比较凉的时候很有效。</p>
<p><strong>价格</strong></p>
<p>BikeAir摩托车空调售价高达1500美元，对一般人不会有什么吸引力，但对很多不差钱的摩托车爱好者来说，还是会心动的。</p>
<p align="center"><a href="/img/20140212/2732542c2541400f82a959e3d7de1a3d.jpg" target="_blank"><img alt="奇葩物 装在摩托车上的空调" src="/img/20140212/s_2732542c2541400f82a959e3d7de1a3d.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/c278942a11404d299f08ed8fbed7c20e.jpg" target="_blank"><img alt="奇葩物 装在摩托车上的空调" src="/img/20140212/s_c278942a11404d299f08ed8fbed7c20e.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/60b4be9b67514d17bea87f7b59f92f3d.jpg" target="_blank"><img alt="奇葩物 装在摩托车上的空调" src="/img/20140212/s_60b4be9b67514d17bea87f7b59f92f3d.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292513.htm" target="_blank" title="我们需要什么样的千元智能手机"><span class="title1">我们需要什么样的千元智能手机</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">20:43:50</span></td>
  </tr>
  <tr>
    <td align="left"><p>2014年一季度，HTC亏损预警，CFO称将会做150美元的手机。无独有偶，魅族在Jwong回归后，除了MX3降价到1999以外，低价手机也在策划之中。</p>
<p>当硬件过剩之后，性能出色的旗舰产品往往销量会下降，而性能够用的低价产品会逐渐流行，这是PC上发生过的故事。智能手机会如何演绎这个过程呢？我们需要什么样的千元智能手机呢？</p>
<p><strong>一、用户的需求是什么？</strong></p>
<p>智能手机的用户如果不跑游戏的话，对手机用什么处理器是没有概念的。而手机上跑大型3D游戏本来就不是主流应用，用电视的大屏幕显然更爽，而手机GPU的3D能力也远不能和PC的显卡相比。</p>
<p>所以，大型3D游戏并不是手机的主流应用。那么用户的需求是什么呢？</p>
<p><strong>（一）手机品牌知名，或者是个国际知名品牌，或者是新品牌</strong></p>
<p>但是各种广告做的大家都知道，这个品牌有实力做广告，有实力请明星代言，不是个小厂的山寨货。因为有这种需求，三星配置极低的千元机销量依然不错。</p>
<p><strong>（二）手机体验够用</strong></p>
<p>体验够用是个很宽泛的概念。简单说，就是达到了用户的感知极限，再提高用户不会有明显的感知。</p>
<p>譬如电脑点开一个常用程序，最快的电脑用0.01秒，普通电脑用0.05秒，差异是5倍，但是超过了用户感知的能力，这就是够用。</p>
<p><strong>二、千元手机的硬件选择</strong></p>
<p>对手机来说，各个硬件规格的选择就是个学问。我们一个一个来看。</p>
<p><strong>（一）屏幕</strong></p>
<p>手机屏幕的选择其实是个学问，人手单手把持操作，手机宽度在70mm左右，用窄边框技术的话可以16：9的比例可以做到5寸，宽边框就只有4.5寸了。屏幕越大分辨率越高越贵。</p>
<p>屏幕尺寸不妥协的话，分辨率就不能太高。人眼在正常手机使用距离，300PPI足够了，贴到鼻子上看才需要400PPI，眼神差点200PPI也能用，所以720P是够用的。4.5寸的540P也不是不能接受。</p>
<p><strong>（二）摄像头</strong></p>
<p>手机摄像头这个东西到了800万BSI的时代，解析力已经超过了早期数码相机了，记录日常生活完全够用，这是苹果坚持800万至今的理由。硬件不是问题，CMOS和镜头都已经很成熟了，拍不好是软件优化的问题。</p>
<p>需要说一下的是，800万BSI市面上有1.1um和1.4um的，1.4um的老一点，但是技术更成熟。1.1um单位像素感光面积较小，靠算法弥补目前除了三星以外还没有解决的比较好的，1.4um的更靠谱一点。</p>
<p><strong>（三）处理器</strong></p>
<p>手机的SOC分集成基带和不集成基带的，一般来说集成基带的因为芯片厂已经做了很多工作，调试起来要简单一些，开发也容易一些，到了用户手里也就稳定一些。</p>
<p>所以MTK和高通大赚特赚，三星、NVIDIA的处理器性能很好，但是用的不多。</p>
<p>对千元机来说，集成基带的高通或者MTK方案也是首选。</p>
<p>手机处理器有个骗人的地方是多核心，四核心八核心&hellip;&hellip;，其实除了浏览器、软解压软件、测试软件，以及一些特别优化过的软件。绝大多数软件只支持单核心或者双核心。</p>
<p>就是说多出来核心，在大多数情况下是没有用的。核心强大、频率高才是王道。A7核心功耗小，高通的环蛇核心性能强。高频双核心的高通体验会更好的一些。</p>
<p><strong>（四）内存和闪存</strong></p>
<p>内存当然是越大越好，闪存当然是越大越好，但是越大意味着越贵。1G内存、8G闪存，再加上TF扩展是个相对平衡的容量。720P屏幕吃内存相对小一点，做好优化1G内存够用。8G闪存在可扩展的情况下也足够使用。大程序、视频之类的装到扩展卡上就是了。</p>
<p><strong>（五）厚度和电池</strong></p>
<p>厚度和电池是矛盾的关系，但是电池可更换就不是问题。不可更换的3000mah电池，不如可更换的2000mah电池靠谱，带三块电池，都比带移动电源&ldquo;打吊瓶&rdquo;方便。</p>
<p>所以，电池容量可以做小一点，降低整体的厚度，提升手感，续航通过多送两块原装电池来解决。</p>
<p><strong>三、软件UI和外观ID问题</strong></p>
<p>其实，智能手机发展到现在，各个厂商的UI在互相学习，互相进步中都做的不错，交互逻辑上基本趋同，差别在审美和风格上，外观ID和做工水平差距很大。</p>
<p>对用户来说，希望厂家竭尽所能来做，软件UI没有成本，用户不喜欢厂家高端用一套漂亮的，低端用一套难看的，而希望一视同仁。</p>
<p>用料因为成本原因，低价机可能只能用塑料，但是做工是不能妥协的。高价机低价机应该有同样的做工标准。</p>
<p>在外观ID设计上，用户希望低价机和高价机趋同，三星就总是喜欢出一些造型类似的mini版。</p>
<p><strong>四、总结</strong></p>
<p>按照上面的描述，我们可以勾勒出来受欢迎千元机的模板。</p>
<p>如果HTC来做，那么就用HTC的品牌，外形用HTCone或者蝴蝶的类似造型，塑料材质。屏幕采用5寸或者4.5寸的720P屏幕。处理器采用1.7ghz环蛇核心的高通8x30AB处理器，摄像头采用800万像素1/3.2寸的BSI摄像头，1G内存，8G存储，2000mAh可更换电池，支持TF卡扩展。UI同样使用SENSE 5.5。</p>
<p>如果HTC能在低价位做出这种手机，是会有竞争力的，这才是我们需要的千元手机。同理，如果魅族或者其他手机厂商来做的话，大家来猜想一下配置会怎么样的吧！</p>
<p align="center"><a href="/img/20140212/38b8a34195114cf182d84ec8dc030c1d.jpg" target="_blank"><img alt="我们需要什么样的千元智能手机" src="/img/20140212/s_38b8a34195114cf182d84ec8dc030c1d.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292512.htm" target="_blank" title="2013年中国手机市场：三星第一 联想第二"><span class="title1">2013年中国手机市场：三星第一 联想第二</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">20:41:08</span></td>
  </tr>
  <tr>
    <td align="left"><p>根据Digitimes Research公布的研究数据，<strong>2013年第四季度中国智能手机销售量环比增长了21.9%，但2013年全年环比仅增长15.9%，市场已经出现饱和趋势</strong>。</p>
<p>品牌方面，<strong>2013年全年三星仍是中国市场销量最高的品牌，联想、华为、酷派分列第二、第三和第四，苹果因为iPhone 5c销售不如预期，加上与中国移动的合作延误，所以仅排名第五</strong>。</p>
<p>华为和联想都在2013年第四季度推出了附属的互联网品牌产品，影响了小米科技的销量，同时也进一步加剧了国内智能手机市场的竞争。</p>
<p>Digitimes Research认为，尽管中国智能手机市场渐趋饱和，但预计2014年第一季度中国智能手机销量仍将达到1亿部。</p>
<p align="center"><a href="/img/20140212/e05de20d7587450f8c90092357a75c4d.jpg" target="_blank"><img alt="2013年中国手机市场：三星第一 联想第二" src="/img/20140212/s_e05de20d7587450f8c90092357a75c4d.jpg" style="border: 1px solid black;" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292511.htm" target="_blank" title="福利没了：魔兽复活卷轴即将成为历史"><span class="title1">福利没了：魔兽复活卷轴即将成为历史</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小路</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">20:40:32</span></td>
  </tr>
  <tr>
    <td align="left"><p>《魔兽世界》国服曾在2012年上线了全新的复活卷轴系统，该系统将允许玩家邀请以往在游戏里的好友或公会成员重回艾泽拉斯。</p>
<p>作为奖励，被邀请者将获得7天1800分钟的免费游戏时间，而邀请者和受邀者还都有机会获得史诗级坐骑奖励。</p>
<p>不过，《魔兽世界》中文官网最新发布的公告却显示，这种超级给力的福利即将离玩家而去。</p>
<p>魔兽运营团队在公告中称：&ldquo;复活卷轴已顺利完成其神圣使命，即将与我们告别了。正如其他大多数卷轴一样，复活卷轴的使用次数随着魔力的逐步消失而香消玉殒、化为凡尘。&rdquo;</p>
<p>按计划，<strong>复活卷轴系统将在2月20日停服维护后关闭邀请功能，但那些已发出的邀请仍有30天的有效期，如果在这段时间内被邀请者没有接受，那么到期将自动作废。</strong></p>
<p align="center"><a href="/img/20140212/9060983229264a27905ff63dd4451edf.jpg" target="_blank"><img alt="福利没了：魔兽复活卷轴即将成为历史" src="/img/20140212/s_9060983229264a27905ff63dd4451edf.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292510.htm" target="_blank" title="叫你丫抄袭 国内山寨手游将遇大规模诉讼"><span class="title1">叫你丫抄袭 国内山寨手游将遇大规模诉讼</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">20:37:25</span></td>
  </tr>
  <tr>
    <td align="left"><p>爆红的手游行业在一片繁荣背后隐藏着一触即发的版权风险。</p>
<p>春节前举行的一场发布会上，触控科技创始人陈昊芝透露，<strong>同样水平的游戏，直接拿&ldquo;七龙珠&rdquo;或&ldquo;进击的巨人&rdquo;的日本动漫素材来包装，投入同样的推广位，与普通产品的转化率相差高达7倍，这就是手游行业侵权事件频发的最大原因。</strong></p>
<p>戏剧性的是，<strong>3天之后，一款名为《龙珠OL》的手游登陆了Appstore，刚刚上线就冲进了Appstore免费榜的前10名，并于2天后冲进Appstore畅销榜的前20名。</strong>一位曾经试图从日本相关出版社购买《龙珠》版权的手游公司负责人向腾讯科技透露，<strong>目前龙珠的IP并未授权给任何中国手游公司，这是一款标准的盗版手游。</strong></p>
<p>尽管在Appstore上这款手游的开发商为&ldquo;xiaonan li&rdquo;，仿若一个普通的个人开发者，但是在这款游戏的详细介绍中留下的服务QQ号显示，这款游戏的代理商为上海沃势文化传播有限公司旗下的Play800，而据服务QQ提供的信息，这款游戏的开发商为北京无忧互通网络科技有限公司。</p>
<p>显然<strong>，&ldquo;xiaonan li&rdquo;只是一种对侵权行为的掩护</strong>，有业内人士透露，尽管日本动漫产业对中国手游市场的侵权动作一向缓慢，但并不代表每次侵权都能全身而退，曾经热门的侵权游戏《梦想海贼王》自1月19日起就被苹果下架，尽管目前无法查实是否与侵权有关，但这款游戏至今未能再次上线。</p>
<p><strong>侵权已成手游行业阶段性问题</strong></p>
<p>有侵权嫌疑手游并非只有《龙珠OL》。<strong>去年有一家名为有爱互动的公司依靠推出海贼王、火影忍者和圣斗士星矢三部日本动漫成为行业知名企业，但有业内人士透露，这家公司并未获得相关动漫的授权。</strong></p>
<p>据悉，尽管有爱互动从SNS游戏起家，但是并未在SNS领域获得成功，反而在手机游戏上切实的挖到了第一桶金。在推出《梦想海贼王》后，这家公司迅速完成了从三线游戏公司到一线手游公司的蜕变。据胡冰在去年移动开发者大会上透露，曾有A股上市公司对有爱互动报出15亿的高价洽谈收购。</p>
<p>除了资本市场的认可，胡冰及有爱互动还数次以成功创业者的身份登上多家创业杂志讲述自身的创业历程。胡冰也曾在多个场合表示将成立一个纯原创的工作室，不依靠外部IP来生存下去。这家公司本可以和多数同行一样在解决温饱问题后趁势洗白，但出乎意料的是该公司在《梦想海贼王》之后继续选择了火影和圣斗士两款日本动漫作品进行了开发。</p>
<p>这背后的逻辑并不复杂，触控科技创始人陈昊芝曾透露，同样水平的游戏，直接拿&ldquo;七龙珠&rdquo;或&ldquo;进击的巨人&rdquo;的日本动漫素材来包装，投入同样的推广位，与普通产品的转化率相差高达7倍。</p>
<p>有爱互动总裁任伟就曾在今年1月初的某媒体活动上公开承认，该公司《黄金圣斗士》游戏一天收入27万。有业内人士透露，这一数字远远高于普通的原创游戏所能带来的收益。且目前该游戏尚未登录苹果官方Appstore，未来收入可能会继续增长。</p>
<p><strong>手游版权诉讼高发期或将到来</strong></p>
<p>沃势文化传播或者有爱互动并非个别，<strong>有关火影、海贼王等热门动漫版权问题早已是业内公开的秘密，也并不仅仅是研发商单方面的犯规，包括手游发行商和渠道商在内，均对这种情况熟视无睹。</strong></p>
<p>有业内人士透露，在高额的流水分成面前，一线手游发行商反而对这类侵权产品青睐有加，因为这类产品转化率高，成功机会大，甚至更愿意投入精力推广。<strong>包括91、PP助手等手游渠道商亦出于同样的理由，对这类游戏的版权问题并未深究，甚至提供专属服务器与研发商联运。</strong></p>
<p>陈昊芝认为，手游行业的侵权问题是一个阶段性问题，至于这个阶段性的问题能不能被遏制，要看是不是有足够多的企业来维权。</p>
<p>事实上，针对手游领域的维权已经在开始，此前搜狐畅游和完美世界曾经联手针对国内武侠手游侵犯金庸版权进行过一轮诉讼，涉及多家即将上市或并购的手游公司，最终这些侵权手游选择了赔偿和道歉。</p>
<p>最新的传言是，已有公司针对《黄金圣斗士》以侵权为由向几大游戏渠道通发了律师函，要求下架有爱互动的相关产品，但这一消息并未得到官方证实。</p>
<p>此外，国内已有多家发行商在与日本相关出版社及版权关联方联系，几部热门动漫的手游版权将于今年年中正式公布。</p>
<p>这意味着，<strong>针对日本动漫作品侵权的手游可能将在年内遭遇大规模诉讼，这将促使中国手游行业更快的告别这一&ldquo;阶段性&rdquo;问题。</strong></p>
<p align="center"><a href="/img/20140212/ac13b7a4179c460189b03ce0dc1dbadc.jpg" target="_blank"><img alt="叫你丫抄袭 国内山寨手游将遇大规模诉讼" src="/img/20140212/s_ac13b7a4179c460189b03ce0dc1dbadc.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292509.htm" target="_blank" title="新一代“收购狂”：谷歌"><span class="title1">新一代“收购狂”：谷歌</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">20:32:10</span></td>
  </tr>
  <tr>
    <td align="left"><p><strong>以前，达成收购交易数最多的硅谷公司是Chipzilla，但是据彭博社统计，现在这一桂冠已经被谷歌抢走了；此外，WPP第二，英特尔排名第三</strong>。在过去的3年里，谷歌达成的收购交易数量比世界上的任何公司都要多。</p>
<p>自2011年出任公司首席执行官之后，谷歌联合创始人拉里&middot;佩奇带领公司向网络广告以外的领域拓展业务；在刚刚结束的一个财季里，谷歌在连网设备、商业服务和移动应用等领域投资了587亿美元的资金。</p>
<p>由Don Harrison领导的并购部门的规模在过去两年至少扩大了50%。与此同时，谷歌旗下的风投部门Google Ventures大量投资初创公司，另外它还设立了一个名为Google Capital的新部门，对创业中后期的公司进行投资。</p>
<p><strong>谷歌在过去3年里达成了127项收购交易，比它在2008年至2011年期间达成的收购交易数增加了一倍多</strong>。在2008年至2011年期间，达成收购交易数最多的Chipzilla一共收购了104家公司。它在2011年至2014年间仅达成了121项收购交易，排名降至第三位。苹果在过去3年里仅达成了12项收购交易。它现在拥有大量的现金，不定期地返还给投资者。</p>
<p>据Google Capital在LinkedIn上公开的信息显示，该部门目前至少拥有6名员工，其中包括去年从TPG Capital跳槽过来的合伙人Gene Frantz。该部门现由谷歌前并购业务主管David Lawee领导。</p>
<p>去年，Google Capital参与了对旧金山在线端到端借贷服务供应商LendingClub的1.25亿美元的投资。去年早些时候，它还跟许多私募公司一起投资了SurveyMonkey。那几家公司都瞄准着小企业用户，谷歌也希望在小企业用户市场推广其云软件。</p>
<p>SurveyMonkey的首席执行官Dave Goldberg称：&ldquo;谷歌一直就是一位慷慨的收购者，它非常擅长整合业务并从中受益。&rdquo;</p>
<p>Google Ventures设立于2009年，谷歌2012年向该部门提供了2亿美元的资金，去年又提高到3亿美元。Google Ventures去年参与了75项投资交易，结束了10项投资交易，其中有3项交易是因为投资目标实现了上市。该部门拥有50多名员工，其中包括9名一般合伙人和1名执行合伙人。</p>
<p>Google Ventures的执行合伙人Bill Maris在一封电子邮件中写道：&ldquo;我们创设Google Ventures的时候，我们对资金结构进行了专门的设计，以期建立一种前所未有的自治。我们从未想过让初创公司为谷歌服务，我们一直考虑的是让谷歌为初创公司服务。&rdquo;</p>
<p>Nextbit Systems的联合创始人Tom Moss接受了Google Ventures和Accel Partners的一笔联合投资。Moss称：&ldquo;谷歌在技术社区依然很受尊重，它是一家热爱技术并志愿推动技术进步的公司。&rdquo;<strong>由于公司去年的营收增长率已由2012年的37%下降到19%，谷歌需要寻找新的市场。除了从并购交易中获得更多的收入之外，投资也可能帮助它了解新的市场和新技术</strong>。</p>
<p>门罗公园的投资公司Bessemer Venture Partners的合伙人Ethan Kurzweil称：&ldquo;他们真的无处不在，你能想到的任何一种战略，都能在谷歌里找到类似的战略。&rdquo;</p>
<p align="center"><a href="/img/20140212/b753ba60a7c749a3942cc7c6981a62a4.jpg" target="_blank"><img alt="新一代“收购狂”：谷歌" src="/img/20140212/s_b753ba60a7c749a3942cc7c6981a62a4.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292508.htm" target="_blank" title="余额宝收益延迟7小时吓坏网友 部分用户赎回资金"><span class="title1">余额宝收益延迟7小时吓坏网友 部分用户赎回资金</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">20:25:23</span></td>
  </tr>
  <tr>
    <td align="left"><p>起床前先查看余额宝的收益，已经成为一部分人的生活习惯。<a class="f14_link" href="http://news.mydrivers.com/1/292/292398.htm" target="_blank">然而今天早上，余额宝却显示&ldquo;暂无收益&rdquo;4个字</a>，这让不少用户开始担忧自己的资金安全。支付宝方面回应称，由于系统升级，收益稍后发放。但一直到今天上午10点左右，才有部分用户的收益陆续到账。原本应该凌晨3点到的收益，迟到了7个小时。</p>
<p><strong>事件 吓坏用户 余额宝今晨无收益</strong></p>
<p>&ldquo;余额宝今天暂无收益，是怎么回事儿啊？&rdquo;早上8点半，在四惠上班的孙小姐就在微信群里发问。随即她又在微博上发现了不少与她有相似遭遇的人。一些余额宝用户已经在晒截图并吐槽了，这让她开始担心起来。</p>
<p>&ldquo;平日里都是一起床就能看到前一天的收益，如果前一天没有收益显示，也是显示再前一天的，没有见过这四个字儿啊？&rdquo;孙小姐说。</p>
<p>一直到余额宝官方微博发消息称，是&ldquo;由于系统升级所致，收益稍后发放&rdquo;后，孙小姐的心才踏实下来，但她同时也表示气愤，&ldquo;系统升级暂无收益可以理解，但系统升级为什么不能提前告知呢？多吓人啊。&rdquo;</p>
<p>一直到10点左右，孙小姐才收到昨日收益，比平常晚了7个小时。</p>
<p><strong>影响 引起恐慌 部分用户赎回资金</strong></p>
<p>虽然有部分余额宝用户在10点时左右陆续收到了收益，但今天余额宝的&ldquo;迟到&rdquo;还是牵动了不少用户的敏感神经。</p>
<p>今天上午9点09分，余额宝发布官方微博，称&ldquo;由于系统升级，收益稍后发放。粉儿们别急，一分也不会少。&rdquo;</p>
<p>这简短的两句话，在40分钟内引发近千条评论，有网友调侃称&ldquo;马云睡着了&rdquo;，但有更多的用户则表现出了担忧和气愤，有用户开始赎回自己余额宝账户的钱。</p>
<p>有用户表示，经过这次&ldquo;心理恐慌&rdquo;，自己以后购买互联网理财产品会更加慎重，&ldquo;本来想把闲钱都放进去的，现在想想还是算了。&rdquo;在丰台上班的刘先生说道。</p>
<p><strong>部分网友吐槽：</strong></p>
<p><strong>以后遇到这样的事要提前通知，不然会让人恐慌。</strong></p>
<p>马云，你是不是睡着了忘了发钱了！</p>
<p>余额宝你应该学会防火防盗，我们这么多血汗钱放在你那儿，不放心啊！</p>
<p>回应 下午3点前到账 都属正常</p>
<p>今天上午，支付宝相关负责人表示，余额宝前一日的收益本来就是第二天15点前发放到账的，&ldquo;平常我们是凌晨三四点发放，今日只是比以前晚了，但并没有超过下午3点。&rdquo;对方说道。</p>
<p>记者查询发现，余额宝官方微博昨日通知称&ldquo;今日收益将在下一日15点前发放至余额宝账户&rdquo;，并且此前每一天都会有该通知。但不少用户依然为这&ldquo;迟到&rdquo;的收益感到了恐慌。</p>
<p><strong>分析 网络理财很&ldquo;猛&rdquo; 安全成软肋</strong></p>
<p>从2013年下半年开始，余额宝、零钱宝、理财通就相继面世，且发展迅速。</p>
<p>支付宝提供的数据显示，春节过后余额宝用户达6100万，资金为2500亿人民币。理财通数据显示，截至1月28日，其规模已经突破100亿。</p>
<p>&ldquo;互联网理财产品虽然发展迅速，但资金安全问题却一直是其软肋。&rdquo;今天上午，某业内人士表示，每次出现关乎资金安全的小问题都会引起用户的担忧，&ldquo;互联网理财产品如果想继续发展，首要的问题还是要确保资金安全。&rdquo;</p>
<p><strong>追访 监管互联网金融央行正讨论</strong></p>
<p>今晨有消息称，央行将牵头对互联网金融实施监管，旨在防范快速发展的中国互联网金融领域风险上升。</p>
<p>据《华尔街日报》报道，消息人士透露，央行目前正同银监会、证监会和保监会协作，制定措施保护消费者信息不被盗窃或滥用，确保互联网投资产品进行充分的风险披露，抑制非法集资行为。</p>
<p>央行将牵头监管互联网金融，主要是因为近期大量科技公司进入了长期被传统银行所主宰的业务。这些业务包括贷款、消费者投资产品等。此外，一些互联网信贷公司的崩盘，也让官方对此作出回应。这些崩盘的公司主要为P2P借贷平台。</p>
<p>今天上午，记者从央行及银监会内部多名人士处证实，目前央行已经将相关方案上报并等待批示，但还需要各部门讨论和提意见。</p>
<p>&ldquo;此事影响不小，方案具体还没确定，不能对外公布，否则会影响很多股票停盘。&rdquo;该消息人士称。</p>
<p><strong><a class="f14_link" href="http://news.mydrivers.com/1/292/292504.htm" target="_blank">支付宝对收益延迟的解释</a>：</strong></p>
<p align="center"><a href="/img/20140212/c99ea6b1eb0043efba22d2b1f940d0fe.jpg" target="_blank"><img alt="余额宝收益延迟7小时吓坏网友 部分用户赎回资金" src="/img/20140212/s_c99ea6b1eb0043efba22d2b1f940d0fe.jpg" style="border: 1px solid black;" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292507.htm" target="_blank" title="冷饭也经典 《战神》合集将登陆PSV "><span class="title1">冷饭也经典 《战神》合集将登陆PSV </span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">20:19:08</span></td>
  </tr>
  <tr>
    <td align="left"><p>索尼的游戏掌机PSV发布以来一直被诟病缺少经典大作，索尼自己也意识到了问题，不管新作旧作，先推出来让玩家有得玩就行。</p>
<p>最近索尼官方博客宣布，<strong>包括《战神》和《战神2》在内的《战神1+2典藏版》合集将登陆PSV，</strong>虽然不是新作，但毕竟是经典游戏，PSV放在角落吃灰的玩家可以拿出来怀旧了。</p>
<p>虽然PSP上也有《战神》，但毕竟PSP的分辨率已经完全跟不上时代了，据了解<strong>PSV《战神1+2典藏版》由索尼圣莫尼卡工作室进行重制，官方保证游戏移植将会做到高清画面。</strong></p>
<p><strong>《战神1+2典藏版》将于5月6日登陆PS Vita，价格为29.99美元</strong>，同时索尼还承诺今年将有超过100款作品登陆PS Vita。</p>
<p>视频地址：<a href="http://v.youku.com/v_show/id_XNjcxNzQ3MDY0.html">http://v.youku.com/v_show/id_XNjcxNzQ3MDY0.html</a></p>
<p align="center"><a href="/img/20140212/d268fe40083e46da9de0a1108f05fc80.png" target="_blank"><img alt="冷饭也经典 《战神》合集将登陆PSV " src="/img/20140212/s_d268fe40083e46da9de0a1108f05fc80.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/d95f4c408a3f4e8f8a165d83b68980eb.png" target="_blank"><img alt="冷饭也经典 《战神》合集将登陆PSV " src="/img/20140212/s_d95f4c408a3f4e8f8a165d83b68980eb.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/a4a5b353457b475cbbccbbe3ff3492e9.png" target="_blank"><img alt="冷饭也经典 《战神》合集将登陆PSV " src="/img/20140212/s_a4a5b353457b475cbbccbbe3ff3492e9.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292506.htm" target="_blank" title="诺基亚：“重磅炸弹”24日震撼登场"><span class="title1">诺基亚：“重磅炸弹”24日震撼登场</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">20:17:27</span></td>
  </tr>
  <tr>
    <td align="left"><p>之前外媒曝光的邀请函显示，诺基亚即将在MWC 2014上举行发布会，亮出全新的设备。就在刚刚，@<a class="f14_link" href="http://weibo.com/1660811367/AwpgeqXUk?mod=weibotime" target="_blank">诺基亚官方微博</a>又对外宣布，<strong>2月24日一款&ldquo;重磅炸弹&rdquo;将在巴塞罗那震撼登场</strong>。</p>
<p>诺基亚的中文邀请函颇为煽情，&ldquo;从2012年秒杀相机的诺基亚808，到2013年的百元神机诺基亚1050，我们总会在MWC上给您带来惊喜&rdquo;，最后还表示，&ldquo;2月24号，巴塞罗那，我们约定树下见&rdquo;。</p>
<p>虽然发布会实在巴塞罗那举行，不过诺基亚也表现出了对中国市场的重视。下面附注显示，<strong>当天还有与中国相关的神秘环节，至于是什么，还有待发布会上揭晓了。</strong></p>
<p>至于当天的主角，根据目前多方爆料称，<strong>应该是一款基于Android系统定制的诺基亚手机</strong>，即<a class="f14_link" href="http://news.mydrivers.com/1/292/292231.htm" target="_blank">&ldquo;诺曼底&rdquo;</a>（准确说是Nokia X）。此外，还有消息称诺基亚还会在此次MWC 2014上亮相一款<a class="f14_link" href="http://news.mydrivers.com/1/292/292138.htm" target="_blank">由旗下Here部门打造的智能手表</a>。</p>
<p>另外，微软与诺基亚的交易基本已经完成，届时会推出全新的手机品牌也未可知。</p>
<p align="center"><a href="/img/20140212/d91d6b2e260f432992f80c840695e2b2.jpg" target="_blank"><img alt="诺基亚：“重磅炸弹”将震撼登场" src="/img/20140212/s_d91d6b2e260f432992f80c840695e2b2.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/a9a6c87731e34dfc889761602babf9ee.jpg" target="_blank"><img alt="诺基亚：“重磅炸弹”将震撼登场" src="/img/20140212/s_a9a6c87731e34dfc889761602babf9ee.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292505.htm" target="_blank" title="罕见股权交易暴露阿里巴巴估值：1280亿美元"><span class="title1">罕见股权交易暴露阿里巴巴估值：1280亿美元</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">20:16:54</span></td>
  </tr>
  <tr>
    <td align="left"><p>根据路透计算，<strong>一笔罕见的阿里巴巴集团股份出售价格，对阿里巴巴的整体估值约达1280亿美元</strong>。</p>
<p>阿里巴巴预计今年稍晚进行首次公开发行(IPO)，可望成为2012年Facebook 上市案以来全球最大规模IPO，阿里巴巴的潜在市值以及其IPO进展倍受关注。</p>
<p>根据路透日前对八位分析师的调查，<strong>估计阿里巴巴市值约1400亿美元，IPO规模预计达150亿美元</strong>。</p>
<p>周二，巨人网络表示，其在阿里巴巴的一批持股，以约1.99亿美元的价格出售给老虎环球基金(Tiger Global)，但并未透露进一步细节。</p>
<p>巨人网络在2011年时，透过中国私募股权公司云锋基金买下一批5000万美元的阿里巴巴股份；阿里巴巴董事局主席马云是云锋基金的共同创办人。</p>
<p>当时阿里巴巴的估值约为320亿美元，因为一个私募股权团队以16亿美元买下5%股权。</p>
<p><strong>假设巨人网络的持股从2011年以来没有改变，则目前这批持股的价值是原始金额的约四倍，而阿里巴巴的身价则相应提高到1280亿美元</strong>。</p>
<p>老虎环球基金还持有阿里巴巴主要对手京东商城的22.1%股权。</p>
<p align="center"><a href="/img/20140212/ab4002b5b57b41d39b073622e1abf46f.jpg" target="_blank"><img alt="罕见股权交易暴露阿里巴巴估值：1280亿美元" src="/img/20140212/s_ab4002b5b57b41d39b073622e1abf46f.jpg" style="border: 1px solid black;" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292504.htm" target="_blank" title="收益迟到 余额宝今天升的什么级？"><span class="title1">收益迟到 余额宝今天升的什么级？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">随心</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">18:53:18</span></td>
  </tr>
  <tr>
    <td align="left"><p>有个笑话，很多人会看余额宝当天的收益，来决定早餐吃什么。作为理财产品，余额宝的收益是大家很关心的。</p>
<p>今天早晨，很多用户发现，自己的手机余额宝里面显示为&ldquo;<a class="f14_link" href="http://news.mydrivers.com/1/292/292398.htm" target="_blank">暂无收益</a>&rdquo;，引起不小的讨论。余额宝迅速在微博发布公告称是&ldquo;系统升级&rdquo;的原因，但具体是什么原因导致升级呢？</p>
<p>一个小时前，支付宝官方微博贴出了长微博，详细解释了这次突然升级的原因，并对没有预先告知用户表达了歉意。</p>
<p>文中提到，<strong>由于昨天用户规模增长太快，使用用户过多，导致今天派发的收益超出了系统阀值，因此才进行了系统升级</strong>，这也是导致今早收益迟到的原因。</p>
<p>根据最近数据，余额宝的用户数量已经冲进了6000万，并且近几个月<a class="f14_link" href="http://news.mydrivers.com/1/291/291928.htm" target="_blank">每月的用户增长都在千万之上</a>，余额宝这是要发呀！</p>
<p align="center"><a href="/img/20140212/1525423cb01448aebf93780ed8dbc4e8.jpg" target="_blank"><img alt="收益迟到 余额宝今天升的什么级？" src="/img/20140212/s_1525423cb01448aebf93780ed8dbc4e8.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><img alt="收益迟到 余额宝今天升的什么级？" src="/img/20140212/c9fc9bfd4c244c92bccc0f575dfecf51.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292502.htm" target="_blank" title="给TA一个抓狂的情人节留言"><span class="title1">给TA一个抓狂的情人节留言</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">娄斌</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">18:17:26</span></td>
  </tr>
  <tr>
    <td align="left"><p>近期Flappy Bird可谓是炙手可热，模仿者一波接着一波，在其中有这样一款名叫《<a class="f14_link" href="http://www.yingyong.so/app/28/14143.htm" target="_blank">神奇的丘比特</a>（<a class="f14_link" href="http://www.yingyong.so/app/28/14143.htm" target="_blank">Amazing Cupid</a>）》的作品，它很有创意的将Flappy Bird同社交网站联系了起来，企图让玩游戏那种虐心的感觉可以与他人&ldquo;分享&rdquo;。</p>
<p>Amazing Cupid这款游戏可以绑定你的社交网站账号，<strong>通过Amazing Cupid可以给你社交网站上的好友发送信息，对方如果想看到，就必须在Amazing Cupid这款游戏中达到你设定的分数，否则就无法看到信息</strong>。试想如果是你的女神给你发送信息，而条件却是一个难以企及的分数，是不是很让人抓狂呢。</p>
<p align="center"><a href="/img/20140212/d275d9c088fc4f4aa1ba35f1c87de1cb.png" target="_blank"><img alt="给TA一个抓狂的情人节留言" src="/img/20140212/s_d275d9c088fc4f4aa1ba35f1c87de1cb.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/be7717ab85064d8cb6fffaafa724e120.png" target="_blank"><img alt="给TA一个抓狂的情人节留言" src="/img/20140212/s_be7717ab85064d8cb6fffaafa724e120.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>这款游戏的主角并不是常见的小鸟，而是一个长着蓝头发的爱神丘比特，玩法和Flappy Bird一样，确保丘比特不会撞到柱子上。如果失败，游戏会显示各种嘲笑的话语。如果成功，那么你将可以看到对方发送的消息。</p>
<p align="center"><a href="/img/20140212/10656402613445c9a57f189214b9d745.png" target="_blank"><img alt="给TA一个抓狂的情人节留言" src="/img/20140212/s_10656402613445c9a57f189214b9d745.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/216e5e73e7384b4aaeadf104a16f8fd3.png" target="_blank"><img alt="给TA一个抓狂的情人节留言" src="/img/20140212/s_216e5e73e7384b4aaeadf104a16f8fd3.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/3929fad2a8b84644b3130ee354ac0a89.png" target="_blank"><img alt="给TA一个抓狂的情人节留言" src="/img/20140212/s_3929fad2a8b84644b3130ee354ac0a89.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>如果你不添加社交网站，也提供了纯游戏模式，并设定了3种难度：普通、困难和极难。据说Amazing Cupid的开发者是一位名叫Anton Soeharyo的印尼开发者，而且他竟然还获得了Flappy Bird开发者的授权。</p>
<p><strong>下载地址</strong>：<a class="f14_link" href="http://www.yingyong.so/app/28/14143.htm" target="_blank">http://www.yingyong.so/app/28/14143.htm</a>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292501.htm" target="_blank" title="赶去投胎？几百人一起玩Flappy bird"><span class="title1">赶去投胎？几百人一起玩Flappy bird</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">随心</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">18:00:17</span></td>
  </tr>
  <tr>
    <td align="left"><p>不管是出于官司原因、还是作者避免其影响自己的生活，极品手游Flappy Bird最终还是<a class="f14_link" href="http://news.mydrivers.com/1/292/292088.htm" target="_blank">下架</a>了。</p>
<p>不过人们对它的热情不减，不但制作出其浏览器插件版本，还制作了多张无厘头的虐鸟动态图，就连安装了该应用的iPhone都炒到了12000元。</p>
<p>现在，国外技术宅又建起了网站<a class="f14_link" href="http://flapmmo.com/" target="_blank">Flapmmo.com</a>，就是依照Flappy Bird原型打造。<strong>在里面，你可以和世界上成百上千的玩家一起</strong>，尽享团队混战角逐的快感。当上百只小鸟向前冲去时，真有种急着去投胎的赶脚，几秒过后，地下也是哀鸿一片。</p>
<p>不过不幸的是，由于玩家过于热情，网站已经崩溃了，目前提示大家明天再来。还等什么？先把这个网站收藏起来，明天一决高下呀。</p>
<p style="text-align: center"><img src="http://ww4.sinaimg.cn/bmiddle/61e61e8cjw1edgq5tucetg208v04s1ky.gif" /></p>
<p>&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292500.htm" target="_blank" title="当摇滚邂逅APP 催生妖孽级应用PolyFauna"><span class="title1">当摇滚邂逅APP 催生妖孽级应用PolyFauna</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">娄斌</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">17:47:25</span></td>
  </tr>
  <tr>
    <td align="left"><p>英国摇滚乐队Radiohead（电台司令）即将发行他们的新专辑，也许是感觉单纯依靠音乐已经无法传达它们所要表达的东西了，于是Radiohead和一家软件开发商共同打造了一款概念型应用&mdash;&mdash;PolyFauna。</p>
<p align="center"><a href="/img/20140212/23e9d0f96b5048d1a44133bb619de08b.jpg" target="_blank"><img alt="当摇滚邂逅APP 催生妖孽级应用PolyFauna" src="/img/20140212/s_23e9d0f96b5048d1a44133bb619de08b.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>你很难用一个确切的标准来界定<a class="f14_link" href="http://www.yingyong.so/app/28/14157.htm" target="_blank">PolyFauna</a>究竟是一款游戏还是应用，它很抽象，从一开始上手便让你摸不着头脑，玩家需要通过手指触摸屏幕来生成各种形态的奇怪图案，而诡异的背景音效，犹如Radiohead乐队主唱的嗓音在低吟。</p>
<p>据Radiohead介绍，应用中这些随机生成的图形，灵感来源于他们的第八张专辑《The King of Limbs》和大脑中潜意识里的动物。</p>
<p>乐队和开发者的目的就是让用户的屏幕上能有一个不断变化的世界出现，推荐带上耳机来体验它，屏幕中会不时的出现一个红点，跟着它就能带领你到另外一个世界。</p>
<p>你需要的就是随着你的意识任意创造各种神奇的生物，充分发挥你的想象，体验这个意识流的世界。该应用需要Android 4.0及以上系统。</p>
<p><strong>下载地址</strong>：<a class="f14_link" href="http://www.yingyong.so/app/28/14157.htm" target="_blank">http://www.yingyong.so/app/28/14157.htm</a>
<p align="center"><a href="/img/20140212/ac813029ff134994952b0c42c066bfc8.png" target="_blank"><img alt="当摇滚邂逅APP 催生妖孽级应用PolyFauna" src="/img/20140212/s_ac813029ff134994952b0c42c066bfc8.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/4d897e33f8ca4774be1a7d9ba8e196ca.png" target="_blank"><img alt="当摇滚邂逅APP 催生妖孽级应用PolyFauna" src="/img/20140212/s_4d897e33f8ca4774be1a7d9ba8e196ca.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/b6155dd9049b441c97e73d4146792fff.png" target="_blank"><img alt="当摇滚邂逅APP 催生妖孽级应用PolyFauna" src="/img/20140212/s_b6155dd9049b441c97e73d4146792fff.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/823d3c77156641a3975fe2c22dda387f.png" target="_blank"><img alt="当摇滚邂逅APP 催生妖孽级应用PolyFauna" src="/img/20140212/s_823d3c77156641a3975fe2c22dda387f.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/409e9df67b3943d1bf555115377026aa.png" target="_blank"><img alt="当摇滚邂逅APP 催生妖孽级应用PolyFauna" src="/img/20140212/s_409e9df67b3943d1bf555115377026aa.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/1fbe0760fd6b451782bb5be5d3196420.jpg" target="_blank"><img alt="当摇滚邂逅APP 催生妖孽级应用PolyFauna" src="/img/20140212/s_1fbe0760fd6b451782bb5be5d3196420.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292499.htm" target="_blank" title="要听话 谷歌两次利用摩托罗拉教训三星 "><span class="title1">要听话 谷歌两次利用摩托罗拉教训三星 </span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">17:45:17</span></td>
  </tr>
  <tr>
    <td align="left"><p>三星是一家巨头公司，它拥有427,000名员工，年营业额超过2,700亿美元，在超过80个业务部门中拥有逾6,000亿美元资产。而谷歌却利用摩托罗拉这根球棒&ldquo;双重打击&rdquo;了三星。</p>
<p><strong>为什么？</strong></p>
<p>从表面上看，联合占有安卓市场81%的份额似乎表明，谷歌和三星是好伙伴关系。三星一直以来都是安卓迅猛增长背后的驱动力，并且令谷歌移动设备处于领先地位。</p>
<p>问题在于，<strong>三星太过居功自傲了</strong>。对三星来说，<strong>制造最受欢迎的安卓手机和平板电脑还不足够，它还想要隐藏安卓系统&mdash;&mdash;最终还想要隐藏谷歌在这份成就中所扮演的角色</strong>。它试图通过&ldquo;TouchWiz&rdquo;来实现这一目标，这是该公司的一款专利皮肤，<strong>它完全掩盖了安卓的各方面特点，使后者变得难以辨认</strong>。对那些购买了&ldquo;三星手机&rdquo;的粗心用户来说，谷歌在设备中的作用基本上已经难以辨识了。</p>
<p>事情还变得更加糟糕了。<strong>三星开始贬低安卓的性能，它将大部分软件&mdash;&mdash;拨号器、日历、邮件客户端、联系人、通知中心、音乐机视频播放器、语音控制等等&mdash;&mdash;全都换成了自己的应用程序</strong>。大部分评论对TouchWiz的评价是负面的，而且其臃肿的软件拖慢了安卓系统的速度、浪费了存储空间，大家都认为，替换后的应用程序都不如原版程序&mdash;&mdash;甚至更糟糕的是，只是些华而不实的垃圾。</p>
<p>然而三星并没有停下自己的脚步。<strong>它将TouchWiz放到了自己的智能电视机上&mdash;&mdash;这是它占据主导地位的又一个市场，并开始打造属于自己的安卓系统竞争对手Tizen&mdash;&mdash;感谢TouchWiz界面的&ldquo;功劳&rdquo;，这个系统在外行人看来似乎没什么两样。它的长期战略已经很清楚了：转而使用Tizen系统，然后用其占领绝大部分的手持设备市场</strong>。谷歌必须进行反击。</p>
<p><strong>如何反击？</strong></p>
<p>武器就是摩托罗拉。2011年8月15日，谷歌宣布以125亿美元现金收购摩托罗拉移动（Motorola Mobility），并由此获得了2万余项移动专利；<strong>谷歌公开宣称，收购这家手机制造商不会以任何方式损害其与手持设备合作伙伴之间的关系</strong>&hellip;&hellip;真的不会，不信我们拉勾。</p>
<p>当然了，谷歌并不指望手持设备合作伙伴完全相信这一说辞，其就该项交易所发布的官方声明也证实了这一点。如果谷歌利用摩托罗拉来打造自己的大型手持设备业务，那么这个市场将会是他们的。谷歌生产的手机将会装载原生安卓系统，而且没有哪家公司&mdash;&mdash;甚至连三星也是一样&mdash;&mdash;能够负担得起补贴手机的成本，因为谷歌可以动用其巨额广告营收。</p>
<p>诱饵已经投下：<strong>除非制造商（即三星）不再任意更改安卓系统，不然就用装载谷歌原生安卓系统的手持设备消灭你。</strong>谷歌不动声色地显示了自己说到做到的能力，同时它加大了Nexus手机的生产力度，并推出了广受欢迎的Motorola X和Motorola G手机&mdash;&mdash;这两款机器几乎完全剔除了对原生安卓系统的定制和美化。</p>
<p>三星上钩了。<strong>2014年1月27日，谷歌和三星签署了一份为期十年、涵盖范围广泛的全球专利协议。隐藏在这份协议之中的一个协定是，三星将减少TouchWiz的存在，重新专注于核心安卓应用程序，而不是其自己的定制及美化，并撤销诸如&ldquo;Magazine UX&rdquo;界面等更为激进的定制行为</strong>。两天后，谷歌宣布将摩托罗拉移动出售给联想，这表明<strong>两项协议的磋商是同时进行的</strong>。</p>
<p><strong>结局</strong></p>
<p>对三星的打击无疑是双倍的。</p>
<p>首先，尽管其在安卓市场的规模和主导地位毋庸置疑，但<strong>三星仍然被&ldquo;乖乖&rdquo;带回了队伍之中。该公司将不再能够随意更改安卓的设计，剔除谷歌的应用程序换上自己的应用，并将谷歌的辛苦努力深深埋藏起来</strong>。有迹象显示，Galaxy S5的发布将低调进行，这表明三星将遵守诺言。</p>
<p>其次，三星放弃安卓转而采用Tizen系统的时间节点不再明朗。随着未来三星手持设备中安卓的愈加闪耀，从安卓转换到Tizen不再会是无缝的。如果三星想要Tizen取得成功，它就必须实打实地努力去打拼，而不是偷偷摸摸地潜行在雷达之下。</p>
<p>所有这一切对安卓用户来说将会是好消息：用户会发现，在升级手机时，在不同的手持设备制造商之间进行转换变得更容易了，因为原生安卓体验（特别是搭载安卓4.4 KitKat优化的安卓系统）使得速度更快、响应更为迅速的经济型设备成为可能。不过，其是否能赋予较小规模的手持设备制造商以更大的机会和所向披靡的三星进行竞争，目前还尚不得知。</p>
<p>那么谷歌那100亿美元的亏损又如何呢？这是一种误解，人们只是用摩托罗拉移动125亿美元的收购价格简单减去了其29.1亿美元的出售价格。这种计算方法忽略了摩托罗拉32亿美元现金、24亿美元的递延所得税资产，以及2013年两个独立摩托罗拉部门总计25亿美元的出售。将联想的收购计算在内之后，谷歌只为留下的这些资产支付了10亿美元：价值55亿美元的摩托罗拉专利，以及后者处于前沿位置的研究实验室。</p>
<p>干得真漂亮啊，谷歌，真心漂亮！</p>
<p align="center"><a href="/img/20140212/f5057f235353408996a0cd2ae267b94e.jpg" target="_blank"><img alt="要听话 谷歌两次利用摩托罗拉教训三星 " src="/img/20140212/s_f5057f235353408996a0cd2ae267b94e.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292498.htm" target="_blank" title="国内首届游戏机展敲定！PS4/XB1或现身"><span class="title1">国内首届游戏机展敲定！PS4/XB1或现身</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小路</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">17:43:35</span></td>
  </tr>
  <tr>
    <td align="left"><p>随着中国正式解除上海自贸区的游戏机生产、销售禁令，封闭已久的国内游戏机市场终于看到了重新开放的希望，而与此同时相关厂商也早已对这块&ldquo;大蛋糕&rdquo;跃跃欲试，市场活动逐渐热络起来。</p>
<p>ChinaJoy官方已经正式确认将与今年的游戏展同期召开国内首届家用游戏机展会，该展会已经<strong>定名&ldquo;次世代游戏机及家庭数字娱乐产品展览会&rdquo;，简称ACH</strong>，由ChinaJoy主办方北京汉威信恒展览有限公司打造。</p>
<p>据悉，<strong>本届展会已经吸引了包括微软、索尼、联想、百事通、华为、Intel以及金士顿在内的大量业内企业关注与参与</strong>，外加之前曾传出的新一代游戏机入华的消息，因此国内玩家将很有可能在本届ACH展会上近距离感受到PS4和Xbox One的魅力。</p>
<p>目前，除了国外厂商之外，国内厂商也对游戏业务显露出极强的欲望。</p>
<p>曾在90年代引领国内游戏机市场<a class="f14_link" href="http://news.mydrivers.com/1/288/288683.htm" target="_blank">小霸王于今年早些时候宣布</a>将联手阿里巴巴推出基于&ldquo;阿里云&rdquo;平台的智能体感游戏机，而在1月份的CES展会上，华为也亮出了名为<a class="f14_link" href="http://news.mydrivers.com/1/289/289111.htm" target="_blank">TORN</a>的Android游戏机。同时，就在本周，华硕代号<a class="f14_link" href="http://news.mydrivers.com/1/292/292228.htm" target="_blank">GameBox</a>的神秘游戏掌机也被曝光。由此看来，2014年下半年的中国游戏市场或将迎来游戏机产品的集体井喷。</p>
<p>按计划，<strong>今年的ChinaJoy定在7月31日至8月3日进行，其中ACH展会将在场地的E7号馆进行。</strong></p>
<p><strong>推荐阅读：</strong></p>
<p><a class="f14_link" href="http://news.mydrivers.com/1/285/285448.htm" target="_blank">Xbox One行货板上钉钉！游戏大作同步抵达</a></p>
<p><a class="f14_link" href="http://news.mydrivers.com/1/289/289674.htm" target="_blank">PS4引进露端倪 国内公司已与索尼接触</a></p>
<p align="center"><a href="/img/20140212/a6112081fd684765a0cc755b3ddbc2d2.jpg" target="_blank"><img alt="国内首届游戏机展敲定：微软/索尼将至" src="/img/20140212/s_a6112081fd684765a0cc755b3ddbc2d2.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
ChinaJoy官方正式确认国内首届游戏机展（ACH）</p>
<p align="center"><a href="/img/20140212/47198ba85b114ae0a294afd881ea8bd3.jpg" target="_blank"><img alt="国内首届游戏机展敲定：微软/索尼将至" src="/img/20140212/s_47198ba85b114ae0a294afd881ea8bd3.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
ChinaJoy展馆平面图</p>
<p align="center"><a href="/img/20140212/61714a3fd87b46b18f45e739af35ff9e.jpg" target="_blank"><img alt="国内首届游戏机展敲定：微软/索尼将至" src="/img/20140212/s_61714a3fd87b46b18f45e739af35ff9e.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
游戏机展会位于E7号馆</p>
<p align="center"><a href="/img/20140212/98b69fe2d2254781b4d6d6e73c3087a5.jpg" target="_blank"><img alt="国内首届游戏机展敲定：微软/索尼将至" src="/img/20140212/s_98b69fe2d2254781b4d6d6e73c3087a5.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
ChinaJoy展馆场地图</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292497.htm" target="_blank" title="倔强的苹果：受伤的iPhone 5C"><span class="title1">倔强的苹果：受伤的iPhone 5C</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">17:41:03</span></td>
  </tr>
  <tr>
    <td align="left"><p>已经过去的2013年Q4，中国市场的智能手机销量环比增长21.9％，不过增速降至至15.9％，在这份新的报告中，《台湾电子》时报给出的结论是，中国智能手机市场越来越饱和。</p>
<p>报告中指出，三星继续是中国智能手机市场份额的领头羊，排在后面的分别是联想、华为以及酷派。</p>
<p>苹果在中国市场的份额为第五，<strong>虽然iPhone 5S卖的不错，但iPhone 5C过高的售价阻碍的人们的选择，所以它并没有为苹果帮什么忙。</strong>PS：衬托iPhone 5S性价比高算吗？</p>
<p>从目前的表态看，<a class="f14_link" href="http://news.mydrivers.com/1/291/291000.htm" target="_blank">苹果即将要放弃iPhone 5C，既然坚持不降价，果断放弃还是很明智的</a>。</p>
<p align="center"><a href="/img/20140212/98bb34e028c54ad7a1d29b75d9436108.jpg" target="_blank"><img alt="倔强的苹果：受伤的iPhone 5C" src="/img/20140212/s_98bb34e028c54ad7a1d29b75d9436108.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292496.htm" target="_blank" title="微软发飙：是朋友就不要让他/她用XP！"><span class="title1">微软发飙：是朋友就不要让他/她用XP！</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">刘艺</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">17:35:53</span></td>
  </tr>
  <tr>
    <td align="left"><p>几天前，微软在<a class="f14_link" href="http://news.mydrivers.com/1/291/291945.htm" target="_blank">官方博客中发帖提醒用户XP生命周期将尽</a>，并请求看到这篇文章的用户帮助宣传，让更多的人尽早摆脱XP以及古旧的PC以确保安全。</p>
<p>今天，该公司又在其官方Twitter上发帖称：&ldquo;<strong>是朋友，就不要让他/她用XP，同时也提醒用户，Windows XP在4月8日就将停止支持服务</strong>（Friends don&rsquo;t let friends use #WindowsXP. Support is ending on April 8! ）&rdquo;。</p>
<p>显然，并不是所有用户都对微软的提醒感兴趣&mdash;&mdash;目前仍有29%的用户在使用XP系统。</p>
<p>微软计划在4月份之前将XP的市场份额将至13%，不过按照目前情况来看，尽管微软用尽了浑身解数，但情况仍不容乐观。</p>
<p align="center"><a href="/img/20140212/35ba511a0bae459cbe366ee52a47650a.jpg" target="_blank"><img alt="微软发飙：是朋友就不要让他/她用XP！" src="/img/20140212/s_35ba511a0bae459cbe366ee52a47650a.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292495.htm" target="_blank" title="无线充电走向主流：两大阵营宣布互兼容"><span class="title1">无线充电走向主流：两大阵营宣布互兼容</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">17:27:19</span></td>
  </tr>
  <tr>
    <td align="left"><p>和早期的手机充电器一样，无线充电技术，遭遇了标准之争和兼容之困，影响了其快速普及。不过，近日传来好消息，<strong>三大标准阵营中的两家，宣布将相互兼容，这意味着标准之争的解决，距离用户更近了一步。</strong></p>
<p>据报道，目前争夺无线充电标准的三大阵营分别是：&ldquo;无线充电联盟&rdquo;（A4WP），PMA，以及拥有Qi标准的WPC阵营。三大阵营目前正争夺得不可开交。</p>
<p>不过日前，<strong>A4WP和PMA两大阵营宣布，将相互兼容对方的无线充电技术标准，这意味着基于两大标准的无线充电设备，均可兼容对方标准的手机，即无线充电设备的手机适配范围将扩大。</strong>不过，在涉及到两个标准的（旧产品）向后兼容上，目前尚未有详情。</p>
<p>美国科技新闻网站读写网引述博通公司（A4WP阵营的投票权会员）首席技术官Henry-Samueli的话说，无线充电技术目前最大的障碍，就是多种标准并存，他表示，希望三大阵营能够协商，找到统一的解决方案。</p>
<p>两大标准兼容之后，还剩下第三个标准&mdash;&mdash;来自WPC的Qi标准。这一标准的支持厂商包括Verizon、摩托罗拉、诺基亚、Beikin、Energizer，以及三星电子、HTC。</p>
<p>需要指出的是，三星和HTC这两家厂商，同时属于WPC和PMA阵营。之前，在GalaxyS4旗舰手机中，三星曾推出了基于Qi标准的无线充电功能（可选配）。</p>
<p>对于未来和Qi标准的兼容性，AW4P的营销总监Geoff-Gordon表示：&ldquo;我们对于和其他组织的对话保持开放态度。&rdquo;</p>
<p>值得一提的是，智能手机市场的另外一家主导厂商苹果，目前尚未加入上述任何三家阵营，苹果向来实行&ldquo;自我封闭的软硬合一&rdquo;，有消息称苹果会使用自家开发的无线充电技术，从而重演在手机&ldquo;有线&rdquo;充电器上与外界格格不入的一幕。</p>
<p align="center"><a href="/img/20140212/12a71973754e4d81a8c94245d2e78387.jpg" target="_blank"><img alt="无线充电走向主流：两大阵营宣布互兼容" src="/img/20140212/s_12a71973754e4d81a8c94245d2e78387.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292494.htm" target="_blank" title="微软大一统不远：WP8.1支持上2K屏！"><span class="title1">微软大一统不远：WP8.1支持上2K屏！</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">17:26:29</span></td>
  </tr>
  <tr>
    <td align="left"><p>由于微软已大范围向开发者开放Windows Phone 8.1的SDK，因此连日来我们看到了不少有关WP8.1的爆料。继<a class="f14_link" href="http://news.mydrivers.com/1/292/292203.htm" target="_blank">国行版通知中心</a>、<a class="f14_link" href="http://news.mydrivers.com/1/292/292384.htm" target="_blank">VPN设置</a>等界面之后，开发者Roman L又曝出了新的消息。</p>
<p>从GDR3开始，WP8开始支持上1080p分辨率以及大尺寸的屏幕，而据Roman L向WMPoweruser独家爆料的导航条信息显示，<strong>WP8.1又大大拓展了分辨率的支持范围，最大可支持2560&times;1440（即2K屏）级别</strong>。此外，WP8.1还新增540&times;960的支持。</p>
<p>与此同时，微软还向开发者表示，<strong>WP目前的逻辑分辨率将基于DPI计算（之前是按物理分辨率计算），因此建议开发者在开发应用的时候，需要使用的最小逻辑分辨率为384&times;640</strong>（之前是480&times;800），从而可以适配不同屏幕尺寸的WP设备。</p>
<p>微软拼命地扩展WP的分辨率支持范围，一方面是为追赶Android的硬件军备竞赛做准备，但其实还有更深一层的含义。</p>
<p>前些时候，微软曾表示将实现Windows系统的融合（Windows Phone、Windows RT以及Windows），虽然近来风声小了很多，但<strong>从WP拼命扩展屏幕分辨率来看，微软的确有大一统Windows的雄心壮志，等到Windows Phone完美支持2K屏幕甚至更高分辨率的时候，三个Windows系统的大融合就真不远了</strong>。</p>
<p align="center"><img alt="微软大一统不远：WP8.1支持上2K屏！" src="/img/20140212/a6a8c9dd3cb24eaabbea78e7b3f92f0a.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p align="center"><img alt="微软大一统不远：WP8.1支持上2K屏！" src="/img/20140212/2132ef70899f46628e0982da06a34e9f.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292493.htm" target="_blank" title="巴掌大的主板也能上Haswell i7"><span class="title1">巴掌大的主板也能上Haswell i7</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">17:21:27</span></td>
  </tr>
  <tr>
    <td align="left"><p>台湾工业与嵌入式厂商艾讯科技(Axiomtek)今天发布了一款奇特的主板产品&ldquo;PICO880&rdquo;，板型规格采用100&times;72毫米巴掌一般大小的Pico-ITX，而处理器却最高可以上Haswell Core i7，是同类产品中有史以来最强的，比艾讯之前用的Atom Z500、G APU强太多了。</p>
<p>这种<strong>&ldquo;单板电脑&rdquo;(single board computer)</strong>面向各种需要紧凑、高性能、多功能、高可靠性计算平台的工业和嵌入式领域，能够适应-20～+70℃的温度范围，</p>
<p>处理器可以在移动版的Haswell Core i7、Core i5、Core i3、Celeron之间选择，但如果是中高端产品，就需要主动风扇的辅助，而且<strong>看上去应该仅限于单芯片封装的U、Y系列，即那些把芯片组也整合在一起的超低功耗版本。</strong></p>
<p>因为板子实在太小，处理器自己就占据了很大一块儿，看上去怪怪的。</p>
<p>主板还提供了一条DDR3 SO-DIMM内存插槽，最大容量8GB；一条全尺寸mini PCI-E扩展插槽，兼容mSATA 6Gbps；百兆以太网卡，支持局域网唤醒、RPL/PXE；输出接口LVDS、DisplayPort，支持Intel AMT 9.5；+12V DC供电。</p>
<p>如果需要更多扩展，还有两块扩展子卡，能增加音频输入输出、四个USB 3.0、四个USB 2.0、一条PCI-E x1、一个DisplayPort、两个UART、LED、开关机等等。</p>
<p>这个小玩意儿会在三月份上市。</p>
<p align="center"><a href="/img/20140212/57969e02c6c242c4865ebc26d412c6bd.jpg" target="_blank"><img alt="巴掌大的主板也能上Haswell i7" src="/img/20140212/s_57969e02c6c242c4865ebc26d412c6bd.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/8a17be704caa444d87e9bf97e6212a95.jpg" target="_blank"><img alt="巴掌大的主板也能上Haswell i7" src="/img/20140212/s_8a17be704caa444d87e9bf97e6212a95.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/878586afcbb04437b14a10f799829828.jpg" target="_blank"><img alt="巴掌大的主板也能上Haswell i7" src="/img/20140212/s_878586afcbb04437b14a10f799829828.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/9caf4a0378384a908675cc5169c8a550.jpg" target="_blank"><img alt="巴掌大的主板也能上Haswell i7" src="/img/20140212/s_9caf4a0378384a908675cc5169c8a550.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292492.htm" target="_blank" title="一代极路由又开卖了 竟是这个价"><span class="title1">一代极路由又开卖了 竟是这个价</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">随心</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">17:14:04</span></td>
  </tr>
  <tr>
    <td align="left"><p>临近情人节，极路由<a class="f14_link" href="http://www.hiwifi.com/" target="_blank">官网</a>也做了活动，一代极路由&mdash;&mdash;极壹开卖了，<strong>售价199元</strong>，这个价格么，不多说了。</p>
<p>最初，极壹推出时就标价199元，主打合金外壳和可装应用（比如屏蔽广告），性价比较高。不过自去年11月，极壹S和极贰推出后，它就基本隐匿江湖了。</p>
<p>极壹S和极贰在主频、无线速率等方面均有提升，售价分别为99元和169元，用户需要加30元买个8GB的SD卡用于安放应用。</p>
<p>不过目前极壹S和极贰在官网都处于断货状态，而这次的极壹是敞开卖的，并且有9种颜色可选，大家会买么？</p>
<p><strong>相关评测</strong>：<a class="f14_link" href="http://news.mydrivers.com/1/279/279305.htm" target="_blank">极壹</a>，<a class="f14_link" href="http://news.mydrivers.com/1/288/288174.htm" target="_blank">极壹S</a></p>
<p align="center"><a href="/img/20140212/dc3dc643fde141cdbd9ca48c44cecbe4.jpg" target="_blank"><img alt="一代极路由又开卖了 竟是这个售价" src="/img/20140212/s_dc3dc643fde141cdbd9ca48c44cecbe4.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/e6a715dd41354914a50c76fd44940ab6.jpg" target="_blank"><img alt="一代极路由又开卖了 竟是这个售价" src="/img/20140212/s_e6a715dd41354914a50c76fd44940ab6.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292491.htm" target="_blank" title="毒物现身 适马DP新机实物曝光"><span class="title1">毒物现身 适马DP新机实物曝光</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">17:10:57</span></td>
  </tr>
  <tr>
    <td align="left"><p>本月13号日本国际摄影器材与影像展览会（CP+）就要开幕了，一大波新机相机都将会在展会上亮相。</p>
<p>前两天适马就亮相了旗下DP系列的新机DP Quattro系列，原本DP M系列就以&ldquo;毒&rdquo;著称，而在新机方面适马不论是外观还是配置都再次玩起了黑科技，其中造型上<strong>相机机身被大大拉长，手柄部分则向内弯曲</strong>，这样的设计在其他便携相机上还从来没有出现过，看起来相当怪异，甚至有人用&ldquo;奇葩&rdquo;来形容。</p>
<p>不过之前适马放出的DP Quattro图片都是渲染图，并没有真机图片，所以真正的相机外观比例还不好判断，今天适马的官方微博放出了一张 DP2 Quattro的真机图，除了硕大的遮光罩外，<strong>相机的机身看起来并没有那么长，便携性应该还是有保证的。</strong></p>
<p>DP Quattro使用的是定焦镜头，仍然<strong>有不同焦距的三款，包括DP1（28mm）、DP2(45mm）和DP3（75mm）三个不同焦段，其处理器从Ture II进化到了Ture III。</strong></p>
<p>另外看到Quattro这样的命名很多人可能会联想到奥迪的全时四驱，其实这样命名应该是跟新的感光元件有关系，<strong>它采用了全新的Foveon多层感光传感器，最上层为蓝色感光层，它被分为四个区块，拥有1960万像素分辨率，而下面两层的R、G感光层则分别可获取490万像素的信息</strong>，适马称新款传感器在画质上相较于上一代产品将有30%的性能提升，它的分辨率表现会达到普通马赛克传感器3900万像素型号的级别。</p>
<p>目前DP Quattro系列的价格还没有公布，能不能成为新的毒物还要看发布以后的更多测试。</p>
<p align="center"><a href="/img/20140212/054695e5f125407a99ffd225212be050.jpg" target="_blank"><img alt="毒物现身 适马DP新机实物曝光" src="/img/20140212/s_054695e5f125407a99ffd225212be050.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/b9ebacc684f342ca9e2a5896a4f6ad0d.jpg" target="_blank"><img alt="毒物现身 适马DP新机实物曝光" src="/img/20140212/s_b9ebacc684f342ca9e2a5896a4f6ad0d.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/5d0745f924824c5ab0bf5bb3d1d213ca.jpg" target="_blank"><img alt="毒物现身 适马DP新机实物曝光" src="/img/20140212/s_5d0745f924824c5ab0bf5bb3d1d213ca.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/c1d79b3894c2487792fe4f6396075b0c.jpg" target="_blank"><img alt="毒物现身 适马DP新机实物曝光" src="/img/20140212/s_c1d79b3894c2487792fe4f6396075b0c.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p>&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292488.htm" target="_blank" title="J.Wong黄章发的第一条微博：的确变了！"><span class="title1">J.Wong黄章发的第一条微博：的确变了！</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">16:51:43</span></td>
  </tr>
  <tr>
    <td align="left"><p><a class="f14_link" href="http://news.mydrivers.com/1/292/292439.htm" target="_blank">开通新浪微博</a>之后，魅族CEO黄章终于发话了。</p>
<p>&ldquo;<strong>大家好！我刚从火星回到地球，我将以最开放的心去包容，去接纳这个世界。大家都知道我们营销做的烂，从今天开始，我就要告诉更多人知道，除了小米手机之外，还有更好的魅族手机可以选择</strong>。&rdquo;</p>
<p>从这条微博不难看出，黄章的确变了：</p>
<p>1、未来将变的更加宽容，估计不会再说&ldquo;不喜欢就滚&rdquo;之类的话了；</p>
<p>2、努力做营销；</p>
<p>3、继续和小米竞争。</p>
<p>煤油们，你们流泪了吗？</p>
<p align="center"><a href="/img/20140212/3f77f226124a4ee89cf58f0416bc2cff.jpg" target="_blank"><img alt="魅族J.Wong黄章的第一条微博：的确变了" src="/img/20140212/s_3f77f226124a4ee89cf58f0416bc2cff.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292487.htm" target="_blank" title="只有5.8mm：最薄八核手机要来了"><span class="title1">只有5.8mm：最薄八核手机要来了</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">16:24:47</span></td>
  </tr>
  <tr>
    <td align="left"><p>金立集团总裁卢伟冰之前曾在微博上表示，<a class="f14_link" href="http://news.mydrivers.com/1/291/291931.htm" target="_blank">称金立将于2月19日发布新旗舰手机</a>：ELIFE S，并暗示该机将以超薄最为卖点。目前，金立已经发出邀请函，<strong>根据邀请函显示，该机将于2月19日晚上18:00在深圳市OCT创意展示中心正式发布</strong>。</p>
<p>根据之前的消息，<a class="f14_link" href="http://news.mydrivers.com/1/292/292195.htm" target="_blank">该机很有可能就是刚刚通过工信部入网许可的GN9000</a>。<strong>其机身厚度只有5.8mm，采用4.99寸1080p屏幕，搭载主频1.7GHz MT6592八核处理器和2GB RAM，运行Android 4.2.2系统</strong>，支持移动TD-SCDMA/WCDMA双3G网络。提供1300万后置+500万前置摄像头，有黑白两色可选。</p>
<p>金立表示，ELIFE S手机从原料选材到加工成品经历了非常繁杂的工序，克服了很多困难才得以最终呈现最完美的作品，该机的宣传口号也打出了&ldquo;不止最薄&rdquo;。</p>
<p align="center"><a href="/img/20140212/3a78a89684254195b2fb73af5b6084ae.jpg" target="_blank"><img alt="只有5.8mm：金立超薄新机发布时间确定" src="/img/20140212/s_3a78a89684254195b2fb73af5b6084ae.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/48d310c7177f4fdf85ab68bde2012180.jpg" target="_blank"><img alt="只有5.8mm：金立超薄新机发布时间确定" src="/img/20140212/s_48d310c7177f4fdf85ab68bde2012180.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><img alt="5.8mm：最薄八核手机曝光" src="/img/20140210/014676f26077439eaef5703ebcf796df.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
<p align="center"><img alt="5.8mm：最薄八核手机曝光" src="/img/20140210/fa07480d74dc4d07a1ddb79320095e8e.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
<p align="center"><img alt="5.8mm：最薄八核手机曝光" src="/img/20140210/f655cd2a0af144aebc4d731b7c3d7583.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
<p align="center"><img alt="5.8mm：最薄八核手机曝光" src="/img/20140210/3c5032eb7ff947c0a7ddee6a35cb7601.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292486.htm" target="_blank" title="联发科64位处理器曝光：主攻廉价机"><span class="title1">联发科64位处理器曝光：主攻廉价机</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">16:18:27</span></td>
  </tr>
  <tr>
    <td align="left"><p>iPhone 5S发布时，其搭载的64位A7处理器引起了众多用户的关注，而这也加速了高通等厂商推进64位处理器的速度。</p>
<p>去年12月初，<a class="f14_link" href="http://news.mydrivers.com/1/285/285854.htm" target="_blank">高通发布了一款新的64位处理器骁龙410，其主攻中国的廉价智能机市场，并且支持LTE 4G网络</a>，很快联发科也要推出与之相PK的产品。</p>
<p>现在有消息称，<strong>联发科也准备了两款64位处理器，其会被冠以MT6752/MT6732的名称，定位于骁龙410类似，支持LTE 4G网络，主攻廉价智能机市场</strong>。</p>
<p>目前还不清楚MT6752/MT6732具体参数，据悉两者最快会在今年中旬跟大家见面。</p>
<p align="center"><a href="/img/20140212/4e2823934e7046dd9a9a7aa3989184a4.jpg" target="_blank"><img alt="联发科64位处理器曝光：主攻廉价机" src="/img/20140212/s_4e2823934e7046dd9a9a7aa3989184a4.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
<a class="f14_link" href="http://news.mydrivers.com/1/292/292396.htm" target="_blank">联发科刚发布的新八核神器MT6595</a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292485.htm" target="_blank" title="Win8.1 Update 1还有重大更新？"><span class="title1">Win8.1 Update 1还有重大更新？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">刘艺</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">16:15:10</span></td>
  </tr>
  <tr>
    <td align="left"><p>进入2月份之后，微软Windows 8.1 Update 1已经正式迈入RTM Escrow阶段，目前网上也已经泄泄漏出了该版本的下载文件。</p>
<p>尽管RTM Escrow版本只是一些bug修复及性能提升，并没有带来重大更新，<strong>但现在据消息人士透露，微软依然在致力于新性能的开发，直至Windows 8.1 Update 1正式投入市场。</strong></p>
<p>目前还没有确切消息表明会带来何种重大变化，不过已经非常确定的是，在Windows 8.1 Update 1中用户可将Metro应用锁定至任务栏，并在单独的窗口中打开。</p>
<p>此外，目前已经十分确定的是，在Windows 8.1 Update 1依然不会出现开始菜单，尽管之前一些报道的观点与之相反。</p>
<p align="center"><a href="/img/20140212/51b01e6db1554fdc8c6205ecd8fc0eb2.jpg" target="_blank"><img alt="Win8.1 Update 1还有重大更新？" src="/img/20140212/s_51b01e6db1554fdc8c6205ecd8fc0eb2.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292484.htm" target="_blank" title="美国通过草案：飞机上不许打电话"><span class="title1">美国通过草案：飞机上不许打电话</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">16:09:36</span></td>
  </tr>
  <tr>
    <td align="left"><p><strong>11日，美国众议院通过了禁止航空乘客在飞行中打电话的草案，将要求美国运输部对此采取行动。</strong></p>
<p>这一草案要求美国运输部颁布管理条例，来禁止乘客在飞机飞行过程中通话。此前，美国运输部已经表示过，它正在考虑引入这样的禁令，以对消费者进行保护。</p>
<p>众议院交通和基础设施委员会没有对此草案表示反对。<strong>众议院共和党与民主党的成员均表示，在飞机上打电话，将会产生噪音，从而影响其他乘客。</strong></p>
<p><strong>报道称，这一草案不会影响去年美国联邦航空管理局(FAA)关于乘客行为的决定。</strong>去年，<a class="f14_link" href="http://news.mydrivers.com/1/280/280999.htm" target="_blank">FAA允许乘客在飞机起降过程中使用智能手机和其他个人电子设备发送电子邮件、短信、上网以及下载数据</a>。</p>
<p>在美国，出于安全考虑，飞行过程中的通话已经被禁止了22年，但美国媒体称，现在，技术的进步解决了这一忧虑。去年底，<a class="f14_link" href="http://news.mydrivers.com/1/283/283758.htm" target="_blank">联邦通信委员会(FCC)</a>曾表示，将考虑允许乘客在飞机上用手机打电话、上网，但<a class="f14_link" href="http://news.mydrivers.com/1/283/283926.htm" target="_blank">引发了不小的争议</a>。</p>
<p align="center"><a href="/img/20140212/4837d0aacef14507b15ad7a745cbc7db.jpg" target="_blank"><img alt="美国通过草案：飞机上不许打电话" src="/img/20140212/s_4837d0aacef14507b15ad7a745cbc7db.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292482.htm" target="_blank" title="真狠 情人节真有人“买光单号电影票”"><span class="title1">真狠 情人节真有人“买光单号电影票”</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">16:05:13</span></td>
  </tr>
  <tr>
    <td align="left"><p>&nbsp;&ldquo;2&middot;14情人节&rdquo;即将来临，情侣们甜蜜备节，单身男女捣乱来真的了。日前，有单身网友发起&ldquo;买光单号座位电影票&rdquo;活动，引来热捧。昨日，位于上海新天地的影院情人节当晚黄金档的一场电影被&ldquo;攻占&rdquo;。不敢暴露真名的发起人接受晨报记者采访，表示自己去年刚和女友分手，这次是开个玩笑，&ldquo;希望情侣们理解&rdquo;。而不少参与网友则是抱着&ldquo;认识一下朋友&rdquo;的目的，希望将一场&ldquo;恶搞&rdquo;变成一场&ldquo;相亲会&rdquo;。</p>
<p><strong>&ldquo;馊招&rdquo;捣乱情人节</strong></p>
<p>据悉，早年日本有源自于漫画的&ldquo;情侣去死去死团&rdquo;，后来在网络上流行了起来，单身网友在论坛上自发组成&ldquo;去死去死团&rdquo;，仇视反对恋人。每到情人节，网上都会传播一些对付情人的办法，其中就有买光所有电影院单号座位的恶搞行为。</p>
<p>孰料，原本一句玩笑话，如今竟然要成真。记者看到，网友&ldquo;UP&rdquo;通过名为&ldquo;梦立方&rdquo;的众筹网站发起募集，以每位49元的票价，号召网友们组成&ldquo;情侣去死去死团&rdquo;，一起买下2月14日晚黄金时段&ldquo;人气特别旺，情侣特别多电影院&rdquo;的单号座位票，阻止情侣约会，&ldquo;想在情人节看场电影？对不住，劳驾两位要分开坐啦，小别胜新婚嘛，给我们这些单身小孩一个机会啦。&rdquo;</p>
<p><strong>单身网友大力热捧</strong></p>
<p>活动上线不到一小时，便获得了8位网友的现金支持。不少网友大力热捧：&ldquo;UP好坏啊，竟然想到这样的损招，我一定会拉我的朋友去帮助你的&rdquo;&hellip;&hellip;此类议论不一而足。</p>
<p>昨日，该活动果然拿下了情人节当天上海新天地UME影城19点29分场次的电影《北京爱情故事》。格瓦拉网站显示，位于上海人气地段的该影城6号厅，只开放了19点29分场次《北京爱情故事》的预售，全场百来个座位，红（已售）、白（可选）相间格外整齐。</p>
<p>本不支持间隔选座，他怎么做到了？</p>
<p><strong>格瓦拉：</strong><strong>买光单号，再卖双号</strong></p>
<p>按惯例，电影票网络预售往往提前两天，且不支持间隔选座。网友自发的这场活动是如何做到的？</p>
<p>格瓦拉电影营销经理吴瑕接受晨报记者采访时，确认了该场电影票是由&ldquo;梦立方&rdquo;网友联系该网站，表达实践&ldquo;买光单号座位电影票&rdquo;愿望，该网站采取&ldquo;包场再分售&rdquo;的方式特例操作，&ldquo;我们把电影院的这一个场次包下来了，把单号先全部买走，2月13日中午开始做双号的抢票活动，单独拿去卖，应该可以吸引不少单身人士。&rdquo;</p>
<p>对于单身网友这一行为艺术，吴瑕表示，从该网站的角度，是为用户提供贴心服务，&ldquo;这也算是情人节的一个活动。情人节的观众群，不仅有情侣，还要照顾到单身男女。&rdquo;</p>
<p><strong>发起者：</strong><strong>和网友一起筹钱开个玩笑</strong></p>
<p>晨报记者与活动发起网友&ldquo;UP&rdquo;取得联系。&ldquo;UP&rdquo;首先表示，不想被情侣骂，要求匿名。</p>
<p>&ldquo;UP&rdquo;告诉记者，自己是一个宅男，去年刚刚和女朋友分手，这次的情人节是分手后的第一个情人节，本来计划和身边一些单身的朋友聚一聚的，但是无意中又看到&ldquo;买光单号座位电影票&rdquo;的段子，就萌发了做这次活动的想法，&ldquo;其实这个段子之前就看到过，不过当时有女朋友，也就一笑了之。&rdquo;</p>
<p>&ldquo;UP&rdquo;介绍，活动进展其实并不顺利，&ldquo;一开始是想自己去买单号票。但是发现网上的卖票软件不能这么买。此后又去找了一些影院想直接买，但是因为情人节当天的票都很火，尤其是爱情片，影院也不想这么卖。如果直接包场买下所有单号票，我又没有那么多钱。最后朋友给我推荐了众筹网站，在梦立方和网友一起筹钱做这个事情。哈哈，希望情侣们理解，只是跟大家开一个小玩笑而已。&rdquo;</p>
<p><strong>参与者：</strong><strong>希望&ldquo;恶搞&rdquo;变&ldquo;相亲会&rdquo;</strong></p>
<p>除了恶搞之外，不少网友是抱着&ldquo;相亲&rdquo;目的积极参与的。参与者中，25岁的金融人士刘先生表示，自己是从微信上看到这个众筹项目的，&ldquo;我本来就是单身，以往情人节过得冷冷清清，看到情侣们出双入对也会感觉很羡慕。今年的情人节我本来也没有什么安排，如今有这样好玩的活动自然想去看看。&rdquo;</p>
<p>刘先生说，&ldquo;我觉得这样无伤大雅的小捣乱其实挺好玩的，而且那时候到场的应该也有很多单身的朋友，也希望借此机会和他们认识一下，解决一下个人问题。&rdquo;</p>
<p align="center"><img alt="真狠 情人节真有人“买光单号电影票”" src="/img/20140212/32dc21ef8a8546cdb70848b74fa2f1a4.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292480.htm" target="_blank" title="手机你羞愧么？平板都能换电池"><span class="title1">手机你羞愧么？平板都能换电池</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">16:01:30</span></td>
  </tr>
  <tr>
    <td align="left"><p>这年头很多手机都采用了不可拆卸电池设计，还自以为时尚，而说起平板机，换电池似乎和它不会有什么交集，但是还真有人做到了。</p>
<p><strong>其实在去年，戴尔就推出了全球第一款可拆卸电池的商务平板机Latitude 10，而最新款的消费级Venue 11 Pro继承了这一突出特点。</strong></p>
<p>这是一款<strong>10.8寸机型</strong>，屏幕分辨率1080p，处理器最高可选Haswell Core i5，也可以选Bay Trail Atom，操作系统是Windows 8.1 Pro，内有2GB内存、32/64GB存储、Wi-Fi、蓝牙、NFC，提供800/200万像素双摄像头，microSD、miniHDMI、USB 3.0、microUSB等扩展接口。</p>
<p><strong>电池的拆卸很简单，翻到背面就可以徒手取下后壳，无需动用螺丝，然后搬动两个开关即可把电池拿下来。</strong>原装的容量32Whr，可续航8-10个小时。</p>
<p>如果你需要长时间拿着它在外使用，充电又不太方便，能换电池确实会很有用。</p>
<p><strong>此外还可以选购一个键盘底座，同样带有电池</strong>，那就相当于可以有最多三块了。</p>
<p>Venue 11 Pro根据配置不同售价<strong>500-850美元不等</strong>。</p>
<p align="center"><a href="/img/20140212/5e3f012c90324fe79849d803dd48bc4c.jpg" target="_blank"><img alt="手机你羞愧么？平板都能换电池" src="/img/20140212/s_5e3f012c90324fe79849d803dd48bc4c.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/63dbc36e2fd54c639cda39e6bb131799.jpg" target="_blank"><img alt="手机你羞愧么？平板都能换电池" src="/img/20140212/s_63dbc36e2fd54c639cda39e6bb131799.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/7f130b3cc5f844ce85fcf4fe59fb6ee2.jpg" target="_blank"><img alt="手机你羞愧么？平板都能换电池" src="/img/20140212/s_7f130b3cc5f844ce85fcf4fe59fb6ee2.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/2efec8ef2891463ebb12c2daeb42114e.jpg" target="_blank"><img alt="手机你羞愧么？平板都能换电池" src="/img/20140212/s_2efec8ef2891463ebb12c2daeb42114e.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/c9b3d6455fc34cbfaadaf1afe591cc19.jpg" target="_blank"><img alt="手机你羞愧么？平板都能换电池" src="/img/20140212/s_c9b3d6455fc34cbfaadaf1afe591cc19.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/9f974fbd16634a60885660f0b530f5e6.jpg" target="_blank"><img alt="手机你羞愧么？平板都能换电池" src="/img/20140212/s_9f974fbd16634a60885660f0b530f5e6.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292479.htm" target="_blank" title="怪事：19条短信瞬间盗走95万元"><span class="title1">怪事：19条短信瞬间盗走95万元</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">16:00:04</span></td>
  </tr>
  <tr>
    <td align="left"><p>春节假期最后一天，全国人民都还沉浸在过年的喜庆之中。但家住浙江海宁的生意人王先生却在那天中午，遇到了惊魂13分钟。</p>
<p>在这段时间内，他的手机连续19次响起短信提示音，每条短信都显示他卡里的金钱正被划出，足足有95万元之多。幸亏王先生保持了镇定，及时通知了警方和银行。最终，这笔巨款被冻结。<br />
　　<br />
但银行卡明明在身边，钱却被人划走了，这样的现象还是值得我们注意。13分钟19条短信，95万被划走。2月6日当天中午12点多，王先生刚吃完中饭，忽然，他的手机响了起来。</p>
<p>第一条、第二条，他并没有注意。但之后不断响起的短信提示音，让他觉得很不寻常，看清手机上的短信内容后，他吓得额头冒出冷汗。短短13分钟，连续19条短信，每条内容都一样，显示他的卡内有5万元被划走，最终共计95万！不过，做生意的王先生还比较冷静，在短信响个不停的同时，他立马打电话给了发卡行反映情况，随后拨打电话报警。</p>
<p>海宁市公安局经侦大队接到报警后，也迅速行动弄清楚了基本情况。这个案件不仅跨省，还跨行：银行卡虽然在王先生身上，但刷卡地却在广西；王先生的卡属于建设银行，而盗刷的地方却在民生银行。</p>
<p>警方要与犯罪分子抢时间。2月7日，办案民警迅速飞往广西南宁。当天下午1点20分到达，随后立刻赶往当地的民生银行。经过一系列的程序，终于在当晚6点多将这笔95万元的巨款冻结。在接到警方的电话后，王先生非常激动，一颗悬着的心终于放了下来。警方分析：四种情况或致类似案件发生卡明明在身边，钱却被划走了。王先生这种遭遇究竟是怎样发生的呢？</p>
<p><strong>办案民警经过分析，认为主要有4种情况：</strong></p>
<p>1、犯罪团伙派年轻的团伙成员应聘到一些夜总会、商场或酒楼内当收银员或者服务员。他们随身带着复制器，当客人刷卡时偷窥消费密码。更为严重的是，他们还收买一些收银员或者正规的小店主来出售信用卡资料，成功盗取一张卡的资料的报酬一般是200元。</p>
<p>2、嫌疑人在ATM机上安装读卡器等盗码设备，持卡人在读卡器上刷卡时，即能克隆获得储户的银行卡资料，再通过ATM机上方微型摄像头偷拍获取密码，以制作伪卡盗刷。</p>
<p>3、持卡人在ATM或POS上使用银行卡取款、商场消费后会打印银行回单，通常单据会详细记录银行卡信息，嫌疑人获得银行回单并偷窥获取事主密码后，便可制作伪卡盗刷钱款。</p>
<p>4、网上购物时，部分购物网站设置&ldquo;信用卡支付&rdquo;，不需密码，只需提供身份信息、卡号、有效期和验证码等信息就可直接扣款。事主一旦不慎登录&ldquo;钓鱼&rdquo;网站或有黑客软件侵入电脑后，信用卡资料容易泄露而被盗刷。</p>
<p><strong>警方提醒：使用芯片卡更安全</strong></p>
<p>信用卡诈骗案件的作案手段非常隐蔽，危害性也很强。针对王先生遇到的这种情况，办案民警也给出了提醒。</p>
<p>市民要提高使用银行卡的警觉性，使用银行卡在ATM机上取现或刷卡消费时要提高警惕，要注意查看是否有被人加装了读卡器、摄像头等设备，并注意观察周围有无可疑人员等，并注意妥善保管办理银行业务后拿到的回执单或刷卡凭条。</p>
<p>如果在网上交易，尽量不在网吧、单位等公共网络上操作，使用电话或网络渠道预订机票、酒店时，要选择知名度高、安全性强的商家，以免信息被泄露。最好要开通手机短信通知，这样可以随时监控自己银行卡的异常交易情况。</p>
<p>得知被盗刷后要在第一时间报警，迅速到最近的银行柜台或ATM机进行一次存、取款交易，并保留凭证，证明被盗刷时银行卡在本人手中，并迅速向发卡银行提出止付要求，最大限度地减少损失。</p>
<p>加强安全措施，现在不少银行已经开始推行芯片卡，这类卡制作成本较高，但安全性也非常强，芯片卡通过了发卡行复杂的认证机制和可靠的加密算法，可实现交易过程的安全控制。市民在刷银行卡时尽量使用具有透支功能的贷记卡，并且贷记卡中不要存入现金，透支消费后在银行规定的还款日期前将透支金额还清。</p>
<p align="center"><img alt="惊魂13分钟：19条短信盗走95万元" src="/img/20140212/cfb86fec6a194e918eebda40c83b3ab3.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
图文无关</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292478.htm" target="_blank" title="国产最强2K旗舰vivo Xplay3S要开卖了"><span class="title1">国产最强2K旗舰vivo Xplay3S要开卖了</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">15:46:35</span></td>
  </tr>
  <tr>
    <td align="left"><p>去年底，vivo正式发布了首款2K屏旗舰Xplay3S，不过由于天线和4G网络适应性等问题，该机始终未能开卖。不过如今，vivo官方微博已经宣布，Xplay3S将从本月底开始正式发售。</p>
<p>随后，vivo高层@紫金城之巅还在微博上透露，之前低估了五模十四频的难度，不过目前解决，并已经进入收官阶段。</p>
<p>除了公布开售时间安排，vivo还不忘提醒，提前抢到预定码的用户购机时可以使用<a class="f14_link" href="http://news.mydrivers.com/1/288/288020.htm" target="_blank">购机基金</a>。原来，由于部分用户认为Xplay3S发货时间太晚，有期货的嫌疑，因此vivo宣布<strong>从1月15日起开始计算，发货时间每晚一天，提前订购Xplay3S的用户将获得5元的赔偿基金，一直累加到正式发货。现在算起来，提前预订用户最多可以省下200多元。</strong></p>
<p>首发采用2K屏的vivo Xplay3S曾被誉为国产最强机，售价为3498元。在它正式开卖前，我们还是再来回顾下它的规格：<strong>配备6寸2560&times;1440屏幕、骁龙800 MSM8974AB处理器、3GB内存以及1300万摄像头，支持LTE TDD及LTE FDD双4G网络，</strong>采用ES9018+OPA2604最顶级DAC和运放芯片组合的Hi-Fi技术。</p>
<p align="center"><a href="/img/20140212/f23ac84e4e0149a98d6413f43063c1ed.jpg" target="_blank"><img alt="国产最强2K旗舰vivo Xplay3S要开卖了" src="/img/20140212/s_f23ac84e4e0149a98d6413f43063c1ed.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><img alt="国产最强2K旗舰vivo Xplay3S要开卖了" src="/img/20140212/b95b0a339dba4c06b4020f0fedf148e8.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292477.htm" target="_blank" title="两大新主机撑场面 去年美国游戏销售才增1%"><span class="title1">两大新主机撑场面 去年美国游戏销售才增1%</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">15:43:44</span></td>
  </tr>
  <tr>
    <td align="left"><p>美国市场研究公司NPD昨天公布了2013年美国游戏市场的相关数据，<strong>跟2102年相比，2013年美国游戏整体销售额仅仅增加了1%。</strong></p>
<p>根据公布的数字，去年美国玩家一共在游戏上花费了153.9亿美元，其中花<strong>在实体游戏上的钱数为63.4亿美元，而18.3亿美元花在了二手游戏盒租赁游戏上，花在数字游戏上的则达到了72.2亿美元。</strong></p>
<p>虽然实体游戏销售额比2012年减少了11%，但是数字游戏的强劲增长抵消了这一部分负增长，在去年，共有36%的13岁以上美国人购买过数字游戏。</p>
<p>另外去年<strong>游戏硬件销售额比2012年增长了5%，</strong>原因自然就是被年底两大新主机的发布所带动的。</p>
<p>虽然1%的增长看起来很可怜，但是其实这样的表现已经很不错了，因为在2011年和2012年，无论硬件还是软件的销售额都是在下降的。</p>
<p align="center"><a href="/img/20140212/0476a138b92a4c6ea36ebb90a58f3849.jpg" target="_blank"><img alt="两大新主机撑场面 去年美国游戏销售才增1%" src="/img/20140212/s_0476a138b92a4c6ea36ebb90a58f3849.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292476.htm" target="_blank" title="《财富》杂志：任天堂玩完了吗？"><span class="title1">《财富》杂志：任天堂玩完了吗？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">15:36:50</span></td>
  </tr>
  <tr>
    <td align="left"><p>美国《财富》杂志网络版今天刊登题为《任天堂玩完了吗？》(Is it game over for Nintendo?)的评论文章称，任天堂近期的业绩低迷与该公司的发展理念有着密切联系，但他们不太可能在短期内转变这种模式，而是有可能继续坚持自己的发展道路。</p>
<p>以下为文章全文：</p>
<p><strong>业绩萎靡不振</strong></p>
<p>将近30年前，任天堂几乎凭一己之力给视频游戏行业带来了新生，带来了第二次机会。1985年，当第一代NES在北美国际玩具展上发布时，甚至没人愿意理会视频游戏&mdash;&mdash;彼时，该行业刚刚经历了大崩盘，整个行业的营收从1983年的32亿美元锐减至1985年的1亿美元。</p>
<p>时光荏苒，岁月如梭，很快到了2013年。整个视频游戏行业早已今非昔比，2013年的整体营收膨胀到930亿美元&mdash;&mdash;而根据美国市场研究公司Gartner的测算，2015年的这一数据将达到1110亿美元。但现在的问题却变成了：任天堂还是视频游戏行业不可或缺的一员吗？</p>
<p>这家创立于1889年的公司凭借扑克牌起家，之后历经坎坷，多次主动变革。而现在，面临Wii U游戏机的销量骤降，他们可能要再次踏上自我革命的征程。</p>
<p>上月，任天堂总裁岩田聪宣布，该公司已经调整了截至2014年3月31日的全财年合并财务预期。最新的预测并不乐观：任天堂将Wii U的年度销量预期下调三分之二，从900万台降至280万台；还将游戏销量预期缩减一半，至1900万套。2012年11月诞生的Wii U远未实现第一代Wii的轰动效应，后者2006年发布后，在全球的累计销量已经超过1亿台。</p>
<p>然而，任天堂的悲哀还在于，虽然第一代Wii U销量不俗，但并未带动软件销量实现同等规模的增长。&ldquo;硬件的主要作用就是推动软件销量。&rdquo;独立视频游戏行业分析师比利&middot;皮德根(Billy Pidgeon)说，&ldquo;任天堂的确通过硬件业务赚了一些钱，而其他企业却未能做到这一点。但归根结底，这并不重要。&rdquo;</p>
<p>第一代Wii的确拥有庞大的用户基础，但这还不够。&ldquo;他们有那么多的硬件，本应带动更多的软件销量。&rdquo;皮德根说，&ldquo;很多人买了Wii后却只能玩系统自带的一款游戏，这种事情屡见不鲜。&rdquo;</p>
<p>自从今年1月宣布调整预期后，外界便纷纷将焦点集中到一个问题上来：任天堂应该坚持现有战略，再试一次？还是应该转变思路？</p>
<p><strong>坚持核心战略</strong></p>
<p>自从岩田聪1月中旬发布声明后，有关该公司是否会把重点转向为其他设备开发游戏的讨论，已经越来越热烈。市场研究公司IHS游戏行业研究总监皮尔斯&middot;哈丁-罗尔斯(Piers Harding-Rolls)表示，把马里奥和其他任天堂独有的游戏角色引入智能手机和平板电脑，可能带来短期的营收增长。然而，该公司似乎已经作出了自己的选择：坚持核心战略。</p>
<p>&ldquo;与苹果和谷歌类似，任天堂甚至Valve都在遵循自己的产品开发路径，不太愿意追随他们只能在其中产生边际影响的市场趋势。&rdquo;哈丁-罗尔斯在研究报告中说，&ldquo;短期或边际影响并不符合它的长期战略目标，他们希望通过长期目标来围绕创新施加更大的赌注，这虽然会增加风险，但却可以带来巨大回报。&rdquo;</p>
<p>哈丁-罗尔斯表示，任天堂很清楚Wii U的糟糕现状，但该公司已经排除了降价等短期调整措施。相反，任天堂可能会把重点集中在Wii U GamePad手柄等配件上。但这对于解决问题并没有太大效果，甚至连短期成效都很难看到。</p>
<p>&ldquo;这些改善对于一个处境不佳的平台而言是有益的，但似乎太过零碎，不够全面。&rdquo;哈丁-罗尔斯补充说，&ldquo;他们没有提到通过进一步加大投资，来让消费者真正了解Wii U是什么，同时明确区分Wii U与Wii的差异。&rdquo;</p>
<p>虽然任天堂的新一代游戏机较微软和索尼的推出时间早了一年，后两家公司都是去年11月刚刚发布的新一代产品，但马里奥等游戏角色短期内不太可能出现在这些平台中。</p>
<p>&ldquo;对任天堂来说，游戏软件和硬件是天生一对。任天堂是游戏开发商，而游戏机则是这些游戏的销售渠道。&rdquo;市场研究公司DFC Intelligence视频游戏行业分析师乔治&middot;克罗尼斯(George Chronis)说，&ldquo;这种理念与索尼和微软存在根本不同，后者主要销售硬件，并通过自主开发游戏来推动游戏机销量。&rdquo;</p>
<p>自从任天堂64在1995年面市后，该公司便始终坚持不追求高性能硬件的战略。&ldquo;与索尼和微软争相开发最前沿的游戏机似乎是在浪费资本。&rdquo;克罗尼斯补充说，&ldquo;所以无论是GameCube还是Wii，任天堂投入的研发和生产成本都低于索尼和微软。Wii开发过程中就决定瞄准还不是Xbox或PlayStaiton用户的主流消费者，通过易于上手的游戏为他们提供乐趣。而体感手柄也为该产品赋予了新奇之处。结果，Wii获得了巨大的成功。&rdquo;</p>
<p><strong>或被成功误导</strong></p>
<p>但成功有时会产生误导作用。&ldquo;这要看你怎么定义成功。&rdquo;皮德根说，&ldquo;你应该关注的是活跃用户群，以及人们是否会购买多款软件。Wii并没有实现这一目标。&rdquo;</p>
<p>软件销量的匮乏甚至导致一些第三方开发商放弃了对任天堂的支持。EA曾经针对任天堂的平台高调发布了多款游戏，但却没有赢得玩家的关注，最终不得不宣布放弃对Wii U的支持。没有了游戏大作，Wii U的硬件销量只能继续萎靡不振。</p>
<p>&ldquo;这是任天堂的另外一个重大问题。&rdquo;皮德根说，&ldquo;我们没有为第三方创造有吸引力的商业机会。&rdquo;与此同时，新的《任天堂明星大乱斗》和《马里奥赛车》也都在开发之中。</p>
<p>&ldquo;这些新项目不太可能立刻带来重大的财务转变，但却会很快消耗研发预算。&rdquo;哈丁-罗尔斯说。然而，&ldquo;这种对创新和个性的依赖是任天堂的核心所在，他们并不一味关注短期利益，而是希望保护任天堂的品牌价值。&rdquo;</p>
<p align="center"><img alt="《财富》杂志：任天堂玩完了吗？" src="/img/20140212/dfae13533c524470821e187fd2cd46d4.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292474.htm" target="_blank" title="佳能发布大底便携相机G1 X Mark II"><span class="title1">佳能发布大底便携相机G1 X Mark II</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">15:33:42</span></td>
  </tr>
  <tr>
    <td align="left"><p>佳能发布大底便携数码相机G1 X Mark II，新品<strong>搭载1.5&quot; 1310万像素CMOS传感器，感光度范围ISO 100-12800</strong>。镜头等效焦距24-120mm，最大光圈f/2.0-3.9。DIGIC 6处理器，最高连拍速度5.2fps，可录制1080p30全高清视频。3&quot; 104万画点可翻折触摸显示屏，内置Wi-Fi/NFC功能，可用于照片分享和远程遥控。</p>
<p>佳能G1 X Mark II是G1 X的后继机型，依然沿用了1.5&quot;CMOS传感器。突破较大的是自动对焦点升至31点，同时搭配最新的DIGIC 6处理器，提升了自动对焦性能。镜头焦距扩曾至24-120mm，最近对焦距离5cm，9叶片光圈。虽然取消了光学取景器，但是可以选择外置电子取景器。</p>
<p>新品在拍摄时可以选择两种长宽比，使用4:3模式时实际像素为1310万像素，3:2则降到1280万像素。除此以外机身还加入了触摸屏、Wi-Fi/NFC等新一代技术，使得操作更方便。</p>
<p><strong>佳能G1 X Mark II将于4月份上市，售价为799.99美元（约合人民币4,848元</strong>）。</p>
<p align="center"><a href="/img/20140212/06bb68d1f1b34927ad6070a1a1ccfb83.jpg" target="_blank"><img alt="佳能发布大底便携相机G1 X Mark II" src="/img/20140212/s_06bb68d1f1b34927ad6070a1a1ccfb83.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/67ad43c8148a4da2a67b8e00d0bab2af.jpg" target="_blank"><img alt="佳能发布大底便携相机G1 X Mark II" src="/img/20140212/s_67ad43c8148a4da2a67b8e00d0bab2af.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/8137c2202a1a41a49acb185754cc6089.jpg" target="_blank"><img alt="佳能发布大底便携相机G1 X Mark II" src="/img/20140212/s_8137c2202a1a41a49acb185754cc6089.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/a706c4e5bf4e41a69c5b244d6bc62028.jpg" target="_blank"><img alt="佳能发布大底便携相机G1 X Mark II" src="/img/20140212/s_a706c4e5bf4e41a69c5b244d6bc62028.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/9442feaed24f42b697ddd291eeb34692.jpg" target="_blank"><img alt="佳能发布大底便携相机G1 X Mark II" src="/img/20140212/s_9442feaed24f42b697ddd291eeb34692.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/40b8ad76c1344757803247358185f87e.jpg" target="_blank"><img alt="佳能发布大底便携相机G1 X Mark II" src="/img/20140212/s_40b8ad76c1344757803247358185f87e.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/ee319c66624a4f659275db11448b4385.jpg" target="_blank"><img alt="佳能发布大底便携相机G1 X Mark II" src="/img/20140212/s_ee319c66624a4f659275db11448b4385.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292473.htm" target="_blank" title="首款由虚幻4引擎打造的手游将于明年推出"><span class="title1">首款由虚幻4引擎打造的手游将于明年推出</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">大鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">15:30:27</span></td>
  </tr>
  <tr>
    <td align="left"><p>Rodeo Games是一家来自英国的独立游戏开发商，曾推出过《<a class="f14_link" href="http://www.yingyong.so/app/17/8529.htm" target="_blank">猎人</a>》系列以及回合制战棋游戏《战锤任务（Warhammer Quest）》等多个获奖游戏。近日，其发布消息称<strong>已获得虚幻引擎4（Unreal Engine 4）的使用许可，正在开发一款由该引擎制作的回合制策略游戏，预计2015年第2季度登陆移动平台，该作很可能成为首款采用虚幻引擎4制作的移动游戏。</strong></p>
<p><strong>虚幻引擎4（Unreal Engine 4）最大特色是实时全局光照功能，让开发商的游戏制作变得更为简便，制作出的游戏在画面、光影、游戏性等方面的表现也更棒。</strong>据Rodeo Games的创始人之一Laurent Maguire介绍，他们的终极目标是为玩家带来理想中的游戏性能和画面，虚幻引擎4能够帮助他们更加高效地实现这一愿望。</p>
<p>此次，Rodeo Games和Games Workshop两家公司再度联手，这款正在制作中的游戏很可能又是一款基于桌游《战锤40K》的新作，让我们一起拭目以待吧！</p>
<p align="center"><a href="/img/20140212/011a3f59346546fc83feab5d80004736.jpg" target="_blank"><img alt="首款由虚幻4引擎打造的手游将于明年推出" src="/img/20140212/s_011a3f59346546fc83feab5d80004736.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292472.htm" target="_blank" title="北京污染指数接近不适合人居住程度"><span class="title1">北京污染指数接近不适合人居住程度</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">15:17:04</span></td>
  </tr>
  <tr>
    <td align="left"><p>上海社会科学院及社会科学文献出版社共同举办的《国际城市蓝皮书：国际城市发展报告（2014）》今日在京举行。<strong>蓝皮书指出，北京生态问题已经成为城市升级的最大短板，污染已接近不适合人类居住的程度</strong>。</p>
<p>蓝皮书指出，在40个国际城市排名中，北京的经济升级能力较好，位列第17位，金融市场指数、经济指标超过国际城市平均值，但凸显创新内涵的智力资本和创新指数却不及平均数。</p>
<p>令人瞩目的是，北京的社会升级指标位居第2，其公平指数在所有国际城市中最高。主要缘于较低的基尼系数以及城市准入门槛的优异表现，表明其具有&ldquo;包容性城市&rdquo;特质。</p>
<p>文化升级是北京的相对薄弱环节，其创新环境指标远低于平均水平、美术馆、博物馆以及国际旅客的数目也仅仅接近于平均值，表明城市的文化吸引力相对有限。<strong>北京的生态问题成为城市升级的最大短板，生态指数居倒数第2，污染指数接近极值，已接近不适合人类居住的程度。</strong></p>
<p align="center"><a href="/img/20140212/284a51471a8445538b2ace7f04254366.png" target="_blank"><img alt="北京污染指数接近不适合人居住程度" src="/img/20140212/s_284a51471a8445538b2ace7f04254366.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292471.htm" target="_blank" title="老照片上色 重温好莱坞女星之美"><span class="title1">老照片上色 重温好莱坞女星之美</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">15:13:32</span></td>
  </tr>
  <tr>
    <td align="left"><p></p>
<p>在美女辈出的好莱坞，曾经有那么一个时代，有一些绝世美女，跨越了地域和时间的界限，直到今天依然鲜活地存在于无数影迷的记忆中。</p>
<p>俄罗斯艺术家Klimbims的上色图更让我们重温旧梦，不管是奥黛丽&middot;赫本、玛丽莲&middot;梦露还是葛丽泰&middot;嘉宝、伊丽莎白&middot;泰勒&hellip;&hellip;<strong>这些几十年前的经典老照片被他重新上色之后，焕发出一种难以名状的美感，为历史抹上一瞥亮丽的色彩，是历史的沉淀与现代技术交融产生的惊人效果。</strong></p>
<p>艺术家不仅需要熟练掌握绘图软件，还要对照片中的细节和氛围有充分的体会，才能够调配出恰如其分的色彩。</p>
<p>除了好莱坞女星Klimbims还为沙俄时期和卫国战争时期的老照片上过色。</p>
<p><strong>延伸阅读&mdash;&mdash;</strong><a class="f14_link" href="http://news.mydrivers.com/1/270/270936.htm" target="_blank">复活历史：神PS给经典黑白照片上色</a></p>
<p align="center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_bd5e2998a63741dabfe530e0a524f596.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
玛丽莲&middot;梦露(Marilyn Monroe)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/66f56efcb79c4bf1ad0a9cc796e963b1.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
玛丽莲&middot;梦露(Marilyn Monroe)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/c3b68c885f6c44b49dd83f9c819de0fe.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
玛丽莲&middot;梦露(Marilyn Monroe)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/04afe2b1e2bb4da7b11777bdae8972e6.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
玛丽莲&middot;梦露(Marilyn Monroe)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/ce878e8640ef47de921c1bab2207d79a.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
玛丽莲&middot;梦露(Marilyn Monroe)</p>
<p align="center"><a href="/img/20140212/5bb09a94fb7644159909a067f32372f4.jpg" target="_blank"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_5bb09a94fb7644159909a067f32372f4.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
奥黛丽&middot;赫本(Audrey Hepburn)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/d190730a707f4da79aeac58559b14522.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
奥黛丽&middot;赫本(Audrey Hepburn)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_2e81fb836fc74b4b938bc9a2b82cbaee.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
奥黛丽&middot;赫本(Audrey Hepburn)</p>
<p style="text-align: center"></p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_eea0aaaefc1146fea32d3b04e82c082f.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
伊丽莎白&middot;泰勒(Elizabeth Taylor)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/a36d0c9ca15a497a86ef9fd02e2ea3f5.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
伊丽莎白&middot;泰勒(Elizabeth Taylor)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/066970cb6876479fa58b6f3410271414.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
伊丽莎白&middot;泰勒(Elizabeth Taylor)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_33302b67852148b68d268884aabc51c1.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
费雯&middot;丽(Vivien Leigh)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/2437cb934b1c4e04a1008c55b22c1b35.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
费雯&middot;丽(Vivien Leigh)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/c23fbd7a90ef4f9dbdb5b7a5b00c4a21.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
费雯&middot;丽(Vivien Leigh)</p>
<p align="center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/c59ff6b99cd04aa88fbf8b42f78b614b.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
费雯&middot;丽(Vivien Leigh)</p>
<p align="center"></p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_5a8425916ae84ff48bd9be2373c387bd.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
格蕾丝&middot;凯利(Grace Kelly)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/5c48cd0af34f4e278bba480568d7bc80.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
格蕾丝&middot;凯利(Grace Kelly)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/8012ab717b2045c58221db6c6a634528.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
格蕾丝&middot;凯利(Grace Kelly)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_98c1e575f50e4d978bafdece44e6a5fc.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
丽塔&middot;海华丝(Rita Hayworth)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/a67ea3ac6105465cabcd708406723ac9.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
丽塔&middot;海华丝(Rita Hayworth)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/fda2381fd7154250b6a777a64f14da9a.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
丽塔&middot;海华丝(Rita Hayworth)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_4264b92ca48544779fdcfce7a2aa9cf2.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
葛丽泰&middot;嘉宝(Greta Garbo)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/b2ffff08de2f42d0b0410c83bc8d3412.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
葛丽泰&middot;嘉宝(Greta Garbo)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_cf8b39586e8b48b5baf2fdb651f3f828.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
艾娃&middot;加德纳(Ava Gardner)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/9ce7f547805b465ea4b3e7d0c6e20e18.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
艾娃&middot;加德纳(Ava Gardner)</p>
<p style="text-align: center"></p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_bf2ff9ba000a477d858ea558e427176e.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
安妮塔&middot;艾克伯格(Anita Ekberg)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/2e51e2a86ad4466da84052220c0dfa39.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
安妮塔&middot;艾克伯格(Anita Ekberg)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_2204d450870d4b77bbdae6b29a708e4d.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
爱莲诺&middot;帕克(Eleanor Parker)</p>
<p align="center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/7264df2e16894acf9dddbc9c7178f034.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
吉恩&middot;蒂尔尼(Gene Tierney)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/211747b108ec49949d63f6f25de6e8b9.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
金格尔&middot;罗杰斯(Ginger Rogers)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/3f7d6f2daeea4c7f9f41c9c345045360.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
艾达&middot;卢皮诺(Ida Lupino)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/36df5883e1484a299a6d2b5409d7d5bc.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
简&middot;曼斯菲尔德(Jayne Mansfield)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/47a1f4759a08468d90c7711ad5e23a89.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
珍&middot;哈露(Jean Harlow)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/07198035092048cbaeb4db9b0abddbeb.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
珍&middot;哈露(Jean Harlow)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/61ec6652cfae4479bb068b8d886863b1.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
珍&middot;哈露(Jean Harlow)</p>
<p style="text-align: center"></p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/90d358757eb04ff7aca61c39b0e27c5f.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
珍妮特&middot;麦克唐(Jeanette MacDonald)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/4183aa1526d046fa8c92deed3f7c8d5e.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
露茜&middot;鲍尔(Lucille Ball)</p>
<p align="center"><a href="/img/20140212/1f8734ca42d24675aabc42d65a0ea60f.jpg" target="_blank"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/s_1f8734ca42d24675aabc42d65a0ea60f.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
露茜&middot;鲍尔(Lucille Ball)</p>
<p align="center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/b32914ae6ea14ef4bb75de82e3969cc3.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
露茜&middot;鲍尔(Lucille Ball)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/07cb001aaa264da698c50fc80585ca99.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
露茜&middot;鲍尔(Lucille Ball)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/e704b2879d6b478d82fcdd76cdd02113.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
卡普西尼(Capucine)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/e9f7cab9c5d749c59880f8d733475bf0.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
赛德&middot;查里斯(Cyd Charisse)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/2526bc6811514f5a89dfdb22d292eed9.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
琳达&middot;达内尔(Linda Darnell)</p>
<p align="center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/5842c5c9082c410c82ee18321906a14a.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
玛丽&middot;麦当劳(Marie McDonald)</p>
<p style="text-align: center"><img alt="老照片上色 重温好莱坞女星之美" src="/img/20140212/75eb340c1bba4e1799ff15f1f6819c11.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
玛琳&middot;黛德丽(Marlene Dietrich)</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292470.htm" target="_blank" title="资深媒体人士：苹果握有非常眩目的新玩意"><span class="title1">资深媒体人士：苹果握有非常眩目的新玩意</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">15:10:31</span></td>
  </tr>
  <tr>
    <td align="left"><p>苹果CEO蒂姆&middot;库克上周在公司总部接受了《华尔街日报》的采访，他在采访过程中透露，苹果在过去两周内一共完成了140亿美元的股票回购。CNBC资深媒体人Jim Cramer对此表示，苹果之所以在短短的两周内购买140亿美元股票，除了有来自投资大亨卡尔&middot;伊坎（CarlIcahn）的压力之外，同时也显示出苹果对于新产品的强大自信。</p>
<p>Cramer认为，<strong>今年将是库克的&ldquo;ShowTime&rdquo;，因为他将会发布一款震惊业界的重量级新产品，就像当年iPhone和iPad的面世一样</strong>。&ldquo;库克先生将会以他的行动告诉那些看衰苹果的投资人：看，你说我们的股票不给力，我又让它涨回来了&rdquo;，Cramer在今天的一篇报导中如是说道。</p>
<p><strong>Cramer确信苹果现在握有一款还未发布、且非常炫目的新玩意</strong>，不过他并未指出这是一款什么样的产品。根据近段时间的多方报导及爆料，能够成为&ldquo;炫目新玩意&rdquo;的无外乎两款产品：<strong>iWatch和重大更新的Apple TV</strong>。此前曾有消息称，苹果会在3月份举行一场发布会，也许届时一切将揭晓。</p>
<p align="center"><a href="/img/20140212/033f7cbe648b439389c2f72f85767e88.png" target="_blank"><img alt="资深媒体人士：苹果握有非常眩目的新玩意" src="/img/20140212/s_033f7cbe648b439389c2f72f85767e88.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292469.htm" target="_blank" title="苹果iTunes部门营收约等于半个Google"><span class="title1">苹果iTunes部门营收约等于半个Google</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">15:01:46</span></td>
  </tr>
  <tr>
    <td align="left"><p>&ldquo;分析师之王&rdquo;Horace Dediu通过分析苹果递交给SEC的10-Q文件（相当于公司的季报）得到这样的结论：苹果iTunes部门的收入，堪比美国财富500强当中的第130名的公司。</p>
<p>在10-Q文件中，苹果特别提到由于中国、日本地区的强劲拉动下，<strong>iTunes部门净销售额有显著提高。2014年第一季度iTunes 、软件以及服务业务的收入为21亿美元，同比去年增长19%。</strong></p>
<p>苹果的iTunes部门的收入来自iTunes Store、App Store、Mac App Store、iBooks Store、AppleCare、授权许可以及其它方面。</p>
<p>Dediu分析，<strong>iTunes部门的季度收入为70亿美元，2013年的总收入为235 亿美元。</strong>以下是Dediu绘制的数据图，可以看到iTunes部门各个部分的收入，以及历年来的增长情况：</p>
<p align="center"><a href="/img/20140212/206eef07282040ea8cfc78da02f49b90.png" target="_blank"><img alt="苹果iTunes部门营收约等于半个Google" src="/img/20140212/s_206eef07282040ea8cfc78da02f49b90.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>iTunes、软件以及服务方面的收入与去年同期相比，具体为数字：第三方内容收入增长46.6%；音乐下载收入减少14%；视频内容收入增长19%；应用下载收入增长105%；书籍下载收入减少9%；服务收入收入增长37%；应用内支付收入增长29%；OS X收入减少8%；iWork and iLife收入减少8%。</p>
<p>Dediu说，苹果的iTunes部门往往不被算进苹果的企业价值当中。但 iTunes 部门的收入、规模难以忽视。以一张图作为说明，那就是<strong>iTunes部门的收入相当于Google收入的一半</strong>。</p>
<p align="center"><a href="/img/20140212/a47d0e242f45489dad6c502a2090e498.png" target="_blank"><img alt="苹果iTunes部门营收约等于半个Google" src="/img/20140212/s_a47d0e242f45489dad6c502a2090e498.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292468.htm" target="_blank" title="男子蹭Wi-Fi丢近6万元"><span class="title1">男子蹭Wi-Fi丢近6万元</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:58:18</span></td>
  </tr>
  <tr>
    <td align="left"><p>近日，@央视新闻、@人民日报等多家权威媒体官方微博曝光，称一种利用假冒商场、快餐店的免费WiFi的骗术正在流行，通过这种方法，用户手机中的网银可以很轻易地被盗用。而扬城一市民最近就遭遇了银行卡上的钱不翼而飞的怪事，警方怀疑与其在公共场所使用免费WiFi有关。警方提醒，不熟悉的免费WiFi，还是应量不要&ldquo;蹭&rdquo;。</p>
<p><strong>网银被盗用，6万剩500块</strong></p>
<p>U盾、银行卡均在手，账户还绑定了手机短信，密码也没丢失&hellip;&hellip;银行卡上的钱却不翼而飞了。近日市民小明(化名)就遭遇了这么一件蹊跷事。69笔&ldquo;隐形&rdquo;交易记录，让他账户上的6万多元，在两天内仅剩500块。</p>
<p>由于平时做生意需要，小明的手机上装有某银行的手机终端软件，也开通了相关的短信提示服务。1月27日当天，小明去银行取钱，一查余额却大吃一惊：自己卡上原本有6万多元，此时却只剩下了500多块。</p>
<p>小明当即报了警。警方调取的银行资料显示，从25日到26日下午，小明的这笔钱分69笔被人盗刷。</p>
<p>经调查，小明的钱是通过3个第三方交易平台&ldquo;溜走&rdquo;的，其中最大的一笔消费记录为4000多元。银行的工作人员发现，小明的手机还被骗子做了手脚，收短信的功能被屏蔽，因此这两天里的所有交易记录，他压根都没有收到短信提示。</p>
<p>在与第三方交易平台客服联系后，警方得知，小明的钱大多被用来充话费、购买游戏点卡等虚拟消费。</p>
<p>目前警方已对此立案调查，初步怀疑可能与小明使用公共场所的免费WiFi有关。</p>
<p><strong>假WiFi 15分钟可盗取信息</strong></p>
<p>昨日，记者就此采访了扬州警方专业人士。据介绍，这种骗术的成本并不高，一台WIN7系统电脑、一套无线网络及一个网络包分析软件，设置一个无线热点AP，就可以轻松地搭建一个不设密码的WiFi。</p>
<p>&ldquo;现如今公共场所免费WiFi铺天盖地，市民很难分辨哪个账号是不法分子搭建的。如果不慎使用了，不法分子只需15分钟就可窃取手机上的个人信息和密码，包括网银密码、炒股账户密码、信用卡账户密码等。&rdquo;警方介绍。</p>
<p>据介绍，利用WiFi窃取他人信息，主要有三种手段：</p>
<p>手段1 以假当真：把正当网站的地址&ldquo;绑架&rdquo;到自己的非法网站上。使用者输入的网址虽然没错，但登录的却是一个仿真度极高的假网站。</p>
<p>手段2 半路拦截：不法分子还可以依靠软件截获网络数据，再通过分析软件，对数据进行分析破解来获得账号和密码。</p>
<p>手段3 偷梁换柱：有些&ldquo;菜鸟&rdquo;黑客破解不了用户的账户密码，却可以通过陷阱WiFi截获受害人转账的接收账户，改成自己的账户，就能收到受害人的转账。</p>
<p><strong>陌生的免费WiFi别&ldquo;蹭&rdquo;</strong></p>
<p>对于这些新出现的骗术，该如何防范？专家建议，市民最好不要在任何陌生的网络中使用账户和密码，另外要给手机安装好防火墙、杀毒软件，设置密钥、数字证书等。此外，对于不确定的无线网络，不要老是&ldquo;蹭为上策&rdquo;，可以先咨询工作人员，确保安全后，再进行连接。</p>
<p align="center"><img alt="男子蹭Wi-Fi丢近6万元" src="/img/20140212/4aeec0e3a406468d8b229dc26eb1c6f4.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292467.htm" target="_blank" title="2K屏+3GB内存 OPPO新旗舰发布时间确认"><span class="title1">2K屏+3GB内存 OPPO新旗舰发布时间确认</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:55:01</span></td>
  </tr>
  <tr>
    <td align="left"><p>昨天我们已经推测<a class="f14_link" href="http://news.mydrivers.com/1/292/292320.htm" target="_blank">OPPO新旗舰Find 7将于3月19日正式发布</a>，今天OPPO官方则正式确认了这一消息。</p>
<p>今天下午，OPPO通过其官方微博宣布了这一消息。按照OPPO的说法，<strong>Find 7的发布时间确实是3月19日，地点则是北京798艺术中心。</strong>另外OPPO还特地把&ldquo;Are&rdquo;大写，不知道其中是否蕴含着其它的意思（有消息称Are代指的就是今天曝光的<a class="f14_link" href="http://news.mydrivers.com/1/292/292391.htm" target="_blank">配备1080p显示屏版的Find 7变种机</a>，难道OPPO本次准备双机齐发？）。</p>
<p>从此前的消息来看，OPPO Find 7将会是继Xplay 3S之后又一款<strong>配备2560&times;1440分辨率显示屏的国产手机，屏幕尺寸应该是5.5英寸，搭载主频2.5GHz的高通MSM8974AC处理器，内置3GB内存</strong>和16GB以上的机身存储空间，并支持TD-LTE。</p>
<p>此外还有消息人士透露Find 7还有可能会搭载F/1.8光圈的摄像头，并且还会采用金属材质的机身设计，但并没有得到证实。</p>
<p>一切还是留到发布会上揭晓吧。</p>
<p align="center"><a href="/img/20140212/f970e2982a2041089ce7e1cbce64c888.jpg" target="_blank"><img alt="2K屏+3GB内存 OPPO新旗舰发布时间确认" src="/img/20140212/s_f970e2982a2041089ce7e1cbce64c888.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/7f8b1e65ddc2423c8b8c2f23cb0abbf6.jpg" target="_blank"><img alt="2K屏+3GB内存 OPPO新旗舰发布时间确认" src="/img/20140212/s_7f8b1e65ddc2423c8b8c2f23cb0abbf6.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292466.htm" target="_blank" title="微软：“网络身份”被盗每年导致50亿美元损失"><span class="title1">微软：“网络身份”被盗每年导致50亿美元损失</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:51:29</span></td>
  </tr>
  <tr>
    <td align="left"><p>微软发布的一份调查报告显示，<strong>全球每年因身份失窃而产生的成本已攀升到了30亿英镑（约50亿美元</strong>）。</p>
<p>对1万名受访者进行了问卷调查后，微软消费者安全指数（Microsoft Consumer Safety Index）报告指出，5%的英国人均遭遇过网络钓鱼欺诈，平均每次损失额约为100英镑。此外，3%的英国人表示他们曾遇到过身份信息失窃的情况，2%的人声称&ldquo;他们的职业声誉受到了影响&rdquo;。在这两种情况中，解决问题的成本平均约为100英镑。</p>
<p>尽管遭遇了这样的损失，但是仅有34%的英国人表示，他们限制了哪些人可以阅读他们共享在社交网络上的信息以及控制了自己发布在网络上的个人信息数量。仅有30%的人声称，他们调整了社交网络隐私设置。</p>
<p>此外，仅有33%的人使用PIN（个人身份号码）或密码对他们的移动设备进行锁屏，尽管其中包含有大量个人和财务信息。</p>
<p>微软首席网络安全官杰奎琳-比尤切勒（Jacqueline Beauchere）说，&ldquo;互联网涉及到了我们方方面面的生活。我们通过互联网相互发送电子邮件，分享照片和视频，支付账单以及购物。但<strong>有时候，我们热爱的这些网络活动往往会将我们置于危险的境地</strong>。&rdquo;</p>
<p>微软警告说，<strong>用户应该使用PIN来保护移动设备，利用安全级别较高的密码来保护他们的所有网络账户</strong>。类似网络银行、账单支付以及网络购物的交易活动应该在安全性较高的网络上进行，而不该通过公共的WiFi进行。它还建议人们定期审查自己发布在网络上的个人信息，及时删除过期的信息。</p>
<p align="center"><a href="/img/20140212/6339797748ec4fe7ae5f6429cb283c24.jpg" target="_blank"><img alt="微软：“网络身份”被盗每年导致50亿美元损失" src="/img/20140212/s_6339797748ec4fe7ae5f6429cb283c24.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292465.htm" target="_blank" title="诺基亚彻底去谷歌化能走多远？"><span class="title1">诺基亚彻底去谷歌化能走多远？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:51:23</span></td>
  </tr>
  <tr>
    <td align="left"><p>根据美国权威媒体报道，诺基亚传言中的第一款安卓手机，将变成事实，在西班牙移动世界大会上发布，这款安卓手机，面临重重疑问。其中外媒质疑称，诺基亚安卓手机，实现了彻底的&ldquo;去谷歌化&rdquo;，在谷歌主导的安卓生态中，这款手机能走多远，将是个疑问。</p>
<p>历史上，在Meego操作系统被枪毙之后，诺基亚又开发了唯一一款Meego手机N9，业界称匪夷所思。而此次推出安卓手机，外界普遍质疑&mdash;&mdash;不是被微软收购了吗？怎么会开发和微软WP竞争的安卓手机？</p>
<p>据媒体报道，诺基亚的安卓手机项目立项之时，诺基亚并未和微软进行收购谈判。今年一季度，微软收购诺基亚的交易将全部结束，此时此刻，微软为何又默许诺基亚开发安卓手机，不得而知。</p>
<p>据报道，<strong>诺基亚这款安卓手机，将做到&ldquo;干干净净&rdquo;的去谷歌化，用户将无法访问谷歌旗下的Play数字内容商城（这也是谷歌从安卓生态唯一的收入来源），手机中也将删除网民喜爱的谷歌产品，比如Gmail，此外还将删除一部分谷歌底层开发接口，这款手机甚至不支持一些高级游戏功能。</strong></p>
<p>诺基亚将会独立提供安卓软件商店，另外，未来不会转让给微软的Here地图产品，也将成为这款安卓手机的推广重点。国外科技媒体Greenbot质疑称，没有了谷歌的互联网服务和软件，诺基亚这款安卓机到底能走多远，值得怀疑。</p>
<p>在塞班、WP时代，诺基亚手机的硬件配置、做工一直为人称道，但是对于那些期待高端配置和奢华摄像头等用户来说，诺基亚安卓机将是一个失望的消息。这款手机配置低端，价格低廉，将主要瞄准发展中国家。据称定价可能在100美元左右，价位接近于诺基亚推出的经典WP手机Lumia520。</p>
<p>外媒评论指出：&ldquo;既然微软已经决定收购诺基亚手机业务，一些人会认为诺基亚在玩火，不过，这款安卓手机，仍然会给诺基亚带来一些好处。&rdquo;比如安卓系统在全球家喻户晓，诺基亚可以借安卓机，重新找回昔日的影响力。</p>
<p>不过，诺基亚这款安卓机，是否又会成为N9一样的昙花一现和&ldquo;绝唱&rdquo;，将值得关注。</p>
<p align="center"><a href="/img/20140212/8217d108a8a74063b401aafc6536218d.jpg" target="_blank"><img alt="诺基亚彻底去谷歌化能走多远？" src="/img/20140212/s_8217d108a8a74063b401aafc6536218d.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292464.htm" target="_blank" title="IBM员工被发现编辑维基百科条目"><span class="title1">IBM员工被发现编辑维基百科条目</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:32:23</span></td>
  </tr>
  <tr>
    <td align="left"><p>作为一家百年企业，IBM对于科技行业具有非常深远的影响力，维基百科上与IBM相关的条目也是数以千计。<a class="f14_link" href="http://wikipediocracy.com/2014/02/09/elementary-my-dear-watson/" target="_blank">Wikipediocracy</a>的一篇文章指出，<strong>IBM雇员被发现编辑了与他们自己或公司相关的维基百科条目，违反了维基百科反对付费编辑和利益冲突的规定</strong>。</p>
<p>文章举例说，一位名叫Tagthestar14的用户重写了IBM计算机科学家Barry Leiba的条目，而这位用户正是Barry Leiba本人；IBM杰出工程师条目完全是由一名叫IBMDE的用户编写的；IBM资深程序员John Paul Morrison编写了他自己和另一位硬件架构师Nate Edwards的条目；David Mertz是IBM开发者园地的专栏作家，他编写了自己的条目，被删除了三次；最离奇的一个故事是IBM Watson项目雇员Karen&ldquo;Fluffernutter&rdquo;Ingraffea，在工作时间编写维基百科条目如Watson，她在维基百科IRC频道非常活跃，2011年被选为维基百科管理员，2013年与另一名维基百科管理员/仲裁人结婚。然而，尽管IBM的雇员被发现编辑了与IBM相关的条目，其总体质量仍然欠佳。</p>
<p align="center"><img alt="IBM员工被发现编辑维基百科条目" src="/img/20140212/ab194951c7454b0d88bfd0f760123f1f.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292463.htm" target="_blank" title="iPhone新用途：测量罩杯买文胸"><span class="title1">iPhone新用途：测量罩杯买文胸</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:30:09</span></td>
  </tr>
  <tr>
    <td align="left"><p>对于大多数女性来说，购买文胸的过程既不舒服，又很无趣，整个购物过程差不多是介于&ldquo;试穿比基尼&rdquo;与&ldquo;去产科医生&rdquo;之间的一种体验。</p>
<p>如果你的文胸是定制的，你还能期待与内衣销售小姐共享试衣间，接受她那冰凉的小手为你提供的服务。如果你的文胸不是定制的，可能出现的情况是，你身上所穿的胸罩尺寸不合适。</p>
<p align="center"><a href="/img/20140212/ad9a814803054dfc9cd278f1eaef5e1b.jpg" target="_blank"><img alt="iPhone新用途：测量文胸尺码" src="/img/20140212/s_ad9a814803054dfc9cd278f1eaef5e1b.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>在2014年情人节悄然来临之际，<strong>一位曾在Google担任高管的麻省理工毕业生为这个令人沮丧的零售难题找到了解决方案。</strong>此时，大部分女性购物者都急切得穿梭于维多利亚的秘密内衣(Victoria&#39;s Secret)间，设法在不受冒犯的前提下，猜对内衣的尺码。</p>
<p>为了确保胸罩与女性的身体完美吻合，Heidi Zak创办的电商初创公司<strong>ThirdLove</strong>使用了智能手机的技术，<strong>让iPhone相机充当虚拟皮尺。</strong></p>
<p align="center"><img alt="iPhone新用途：测量文胸尺码" src="/img/20140212/98a956dd7f65401b943d686c48b1e0d6.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p>ThirdLove公司最近重新设计的应用程序刚刚落户苹果的iTunes应用程序商店的特色区&mdash;&mdash;对于一家小公司来说，这是设法除却喧嚣(尤其是2月14日情人节即将来临)的聪明做法。</p>
<p>2012年，Zak与自己的丈夫、红杉资本前合伙人David Spector一起创办了ThirdLove。David Spector曾任职于Google，在消费者互联网方面拥有特长。</p>
<p>2013年，该公司宣布完成560万美元的种子期融资。该轮融资的特点是，投资人名单中有一些广受追捧的人物：社交媒体界的富豪Yuri Milner、投资人Zachary Bogue(雅虎美女CEO Marissa Mayer的丈夫)以及硅谷风投大佬Andreessen Horowitz。</p>
<p><strong>ThirdLove自主开发了胸罩罩杯完美匹配女性身体的技术。为了让这项技术真正发挥作用，这家公司申请了7项专利(另外4项正在申请之中)。该公司的内部团队成员里有4名计算机视觉与图像识别工程师，他们是从像NVIDIA这样的知名公司里招募而来的。</strong></p>
<p>购物者必须要做的事情就是穿着合身背心，站在镜子前，来2张&rdquo;自拍&rdquo;，不过快拍时选择2个略有不同的角度。iPhone手机就成了一种校准对象或测量的标准单位。</p>
<p align="center"><img alt="iPhone新用途：测量文胸尺码" src="/img/20140212/ceccbf42c4f54b9cbb599fef4ac3c75d.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p>&ldquo;<strong>当女性顾客拍照时，我们的计算器能测量出她的体型与她的手机尺寸的匹配比例。</strong>&rdquo; Zak解释道。</p>
<p>ThirdLove开发的此款应用程序立即记录下购物者的胸围等数据，随后该公司将购物者的确切尺寸(型号)通知给定制店。胸罩起始售价为39美元，而内裤的起始售价仅为9美元。</p>
<p>虽然该公司欢迎女性选购行业标准大小的胸罩和内衣(如34B或38DD等尺寸)，大约85%的顾客还倾向于使用ThirdLove自定的尺寸范围，包括半圆罩杯的尺寸。</p>
<p align="center"><img alt="iPhone新用途：测量文胸尺码" src="/img/20140212/5e67da3a2e9d40b6926d0726e0c5a1fa.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p>Zak不愿透露确切的销售数据，不过她表示，自从发布此款应用程序以来，该公司的营收环比增长率高达100%。</p>
<p>该初创公司报道，它的退货率大约为13%，低于很多知名度更大的电子商务运营店。根据《华尔街日报》最近发表的一篇文章报道，QVC公司的退货率略低于20%，但鞋业巨头Zappos的顾客退货的比例更是达到了35%左右。</p>
<p>在Zak眼里，ThirdLove才是一家真正意义上的高科技时尚公司&mdash;&mdash;这个术语被炒作的也过于频繁了。她说：&ldquo;网上销售产品不是真正意义上的&lsquo;技术&rsquo;。我们招聘的前两名员工，一名是内衣设计师，另一名是计算机视觉工程师。我们正利用技术，解决现实问题。&rdquo;</p>
<p>看看使用ThirdLove的步骤吧：<strong>首先，使用你的iPhone下载<a class="f14_link" href="https://itunes.apple.com/us/app/thirdlove/id692635364?mt=8" target="_blank">这个应用</a>，拍两张穿紧身背心的照片(不同角度)，ThirdLove会为你量身定制尺寸(精确到0.5罩杯)，并承诺确保用户隐私及照片安全。</strong>之后，你要做的就是&mdash;&mdash;烧掉其它胸罩吧(当然是开玩笑了)！接下来你就可以选择款式和颜色以及包装了。</p>
<p>目前，Android手机还不能应用此款程序，但不久的未来就能用上它。</p>
<p align="center"><img alt="iPhone新用途：测量文胸尺码" src="/img/20140212/17e8ec4537274f9abe833e77ba1557fb.jpeg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p align="center"><img alt="iPhone新用途：测量文胸尺码" src="/img/20140212/22b19f400c2649b6b60f7ad67119ed15.jpeg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p align="center"><img alt="iPhone新用途：测量文胸尺码" src="/img/20140212/1497a5e809ba4458973b22cb48fbabdb.jpeg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p align="center"><img alt="iPhone新用途：测量文胸尺码" src="/img/20140212/73edfaf42cc34439ad146bb5b7169a6c.jpeg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p align="center"><img alt="iPhone新用途：测量文胸尺码" src="/img/20140212/a7aeda69011f45d9b8ba0d7024d10044.jpeg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292462.htm" target="_blank" title="索尼A6000正式发布 世界最快对焦速度"><span class="title1">索尼A6000正式发布 世界最快对焦速度</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:30:02</span></td>
  </tr>
  <tr>
    <td align="left"><p>今天，索尼正式发布了传闻已久的新微单A6000，<strong>之前流传的E卡口相机最快对焦速度一点都不假，而更惊人的是索尼表示，A6000拥有全球最快自动对焦速度！</strong></p>
<p>先来看相机参数，<strong>A6000使用了2430万像素的APS-C CMOS传感器</strong>，重点就在对焦方面，<strong>它采用了相位/反差混合对焦系统，包括179点相位对焦点和25点反差点，对焦点共可覆盖92%的图像区域，</strong>索尼表示这套对焦系统在对焦速度和精度方面都是很大提升尤其是在连拍时的表现。索尼公布的最<strong>短合焦时间仅0.06秒。</strong></p>
<p align="center"><a href="/img/20140212/a663594cf9b74558929f6c2be9be6b99.jpg" target="_blank"><img alt="索尼A6000正式发布 世界最快对焦速度" src="/img/20140212/s_a663594cf9b74558929f6c2be9be6b99.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p><strong>A6000的处理器是跟全画幅微单A7/A7R一样的Bionz X，它的最大特点是区域降噪和衍射修正</strong>，其中区域降噪就是处理器可以自动识别暗部区域和亮部区域，然后分别对其进行不同等级的降噪，而衍射修正则能改善使用小光圈拍摄时造成的衍射效应。</p>
<p>其他方面A6000还配有<strong>两个物理拨盘、内闪和标准热靴接口，LCD屏幕为3英寸，92.1万像素，</strong>相机支持拍摄1080/60p全高清视频，并支持HDMI输出，<strong>ISO 100-51200，内置Wi-Fi和NFC，机身共7个可自定义按键。</strong></p>
<p>值得注意的是A6000的EVF并不是之前传闻中跟A7/A7R的相同规格，而是<strong>只有144万像素，比NEX-6的236万像素还缩水了</strong>，同时这也意味着<strong>A6000的定位不是之前爆料中NEX-6和NEX-7两款机型的继任者，而只是NEX-6的升级款。</strong></p>
<p align="center"><a href="/img/20140212/2d2f1d3d7b9b47f7852ec8bb5ffd4afc.jpg" target="_blank"><img alt="索尼A6000正式发布 世界最快对焦速度" src="/img/20140212/s_2d2f1d3d7b9b47f7852ec8bb5ffd4afc.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p>A6000共有黑色和银色两款，其中银色款的手柄部分是黑色，看起来有些不搭，机身售价为650美元（约合3939人民币），带有16-50mm镜头的套机售价为700美元（约合4242人民币）</p>
<p align="center"><a href="/img/20140212/c3c4d7ef29224e73bee901c7ad921b11.jpg" target="_blank"><img alt="索尼A6000正式发布 世界最快对焦速度" src="/img/20140212/s_c3c4d7ef29224e73bee901c7ad921b11.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/01b26f5ea76849ef835557f7ef87f7cd.jpg" target="_blank"><img alt="索尼A6000正式发布 世界最快对焦速度" src="/img/20140212/s_01b26f5ea76849ef835557f7ef87f7cd.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/2b4dbbae249244e893e2f541e709226e.jpg" target="_blank"><img alt="索尼A6000正式发布 世界最快对焦速度" src="/img/20140212/s_2b4dbbae249244e893e2f541e709226e.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/89187a348b1842ee9566c0421e2a7552.jpg" target="_blank"><img alt="索尼A6000正式发布 世界最快对焦速度" src="/img/20140212/s_89187a348b1842ee9566c0421e2a7552.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/9582f0cc4b2f4d58a749ea0e86ff71ea.jpg" target="_blank"><img alt="索尼A6000正式发布 世界最快对焦速度" src="/img/20140212/s_9582f0cc4b2f4d58a749ea0e86ff71ea.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/76b3f6ca84a243f2b4f2e55c32c297f7.jpg" target="_blank"><img alt="索尼A6000正式发布 世界最快对焦速度" src="/img/20140212/s_76b3f6ca84a243f2b4f2e55c32c297f7.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/ca5334c6b4f245c1be9b65e9fa62ca9b.jpg" target="_blank"><img alt="索尼A6000正式发布 世界最快对焦速度" src="/img/20140212/s_ca5334c6b4f245c1be9b65e9fa62ca9b.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/a313ce4b56c146268b88cdc0381714aa.jpg" target="_blank"><img alt="索尼A6000正式发布 世界最快对焦速度" src="/img/20140212/s_a313ce4b56c146268b88cdc0381714aa.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/aa4ff536ba6e47dfb57e4acaf3083b5b.jpg" target="_blank"><img alt="索尼A6000正式发布 世界最快对焦速度" src="/img/20140212/s_aa4ff536ba6e47dfb57e4acaf3083b5b.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p>&nbsp;</p>
<p>&nbsp;</p>
<p><br />
&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292461.htm" target="_blank" title="从“黄章的第一条微博”说开去"><span class="title1">从“黄章的第一条微博”说开去</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">随心</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:30:04</span></td>
  </tr>
  <tr>
    <td align="left"><p>新的一年，魅族要变，黄章不得不变。</p>
<p>今天中午，魅族董事长J.Wong黄章终于<a class="f14_link" href="http://news.mydrivers.com/1/292/292439.htm" target="_blank">正式入驻新浪微博</a>，短短数分钟，粉丝数逼近3万。你知道他的第一条微博是什么？答案是：换主页封面图时自动生成的。当然他立刻就删除了，不过<a class="f14_link" href="http://www.weibo.com/1738877650/Awmyj4Zbj?mod=weibotime" target="_blank">@博数码</a>抢先截了图，并附言&ldquo;黄章似乎玩不转微博呀&rdquo;。</p>
<p><strong>这本是个玩笑，但我们不仅要问，黄章真的能玩转微博（营销）这个新时代的工具么？</strong></p>
<p>魅族科技成立于2003年，距今已经有超过10年的历史了。但正如网上所说，&ldquo;魅族像是一个与世无争的隐士，除了一年一度的新品发布会外，其他时间他们都在埋头用自己的方式做着自己认为正确的事情。&rdquo;</p>
<p>不过在移动互联网时代，一切都变了，无论是产业链供给、产品生产还是终端销售。其中最大的变化就是营销模式。<strong>黄章也同意，他那种低调的管理风格，已经不适合如今移动互联网时代了。</strong></p>
<p>在国内需要打好组合拳的同时，<strong>黄章还肩负更大的担子：魅族要走向国际</strong>。今年1月的CES国际电子展上，MX3首次在美国亮相，同时魅族还宣布这款手机将于今年三季度登陆美国市场。但进军美国的门槛并不低，不仅产品要过硬，营销等手段也是必要的。</p>
<p>从某种程度上来说，微博即是营销。正如文章开头那个玩笑，黄章真的能用好微博这个工具么？他又能否脱胎换骨，带领魅族走向新的高度？就让我们拭目以待吧。</p>
<p align="center"><a href="/img/20140212/832d902f67d9446b894760bbf2dfc8a6.jpg" target="_blank"><img alt="从“黄章的第一条微博”说开去" src="/img/20140212/s_832d902f67d9446b894760bbf2dfc8a6.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/f71dfbc66c74404f93f14a1ec55ba8a3.jpg" target="_blank"><img alt="从“黄章的第一条微博”说开去" src="/img/20140212/s_f71dfbc66c74404f93f14a1ec55ba8a3.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
明显可以取消同步到微博，黄章应该是没怎么用，所以顺手没去掉，后来只好马上删掉</p>
<div id="news_toupiao"><script type="text/javascript" src="http://act.mydrivers.com/js/wforms.js" ></script><script type="text/javascript" src="http://act.mydrivers.com/js/localization-cn.js" ></script>
<style type="text/css">
LABEL.preField {	PADDING-RIGHT: 2px; DISPLAY: inline-block; PADDING-LEFT: 2px; PADDING-BOTTOM: 0px; MARGIN: 0.6em 0px 0px; ; WIDTH: expression('10em'); PADDING-TOP: 0px; min-width: 10.5em}.errMsg {	color: red;	display: inline;}.errFld {	border-color: red;}</style>
<form action="http://act.mydrivers.com/GetResult.aspx" method="post" name="VoteForm" target="_blank">
	<p>黄章变了，会给魅族带来多大改变？</p>
	<input name="SubjectId" type="hidden" value="479" /><b>1.你认为黄章在未来的一年会给我们全新的面貌么？</b><br />
	<br />
	<input id="4079" name="1029" type="radio" value="4079" /><label class="postField" for="4079">会，他已经意识到了自己过去的不恰当之处，并且说过要变</label><br />
	<input id="4080" name="1029" type="radio" value="4080" /><label class="postField" for="4080">不会，人性很难改变，魅族的文化氛围也不是一两天能改变的</label><br />
	<b>2.黄章的改变，会给魅族带来多大的变化？</b><br />
	<br />
	<input id="4081" name="1030" type="radio" value="4081" /><label class="postField" for="4081">领导者思路的改变，可以起到扭转全局的效果</label><br />
	<input id="4082" name="1030" type="radio" value="4082" /><label class="postField" for="4082">很难说，拭目以待</label><br />
	<input id="4083" name="1030" type="radio" value="4083" /><label class="postField" for="4083">应该没有多大效果</label><br />
	<br />
	<input name="Submit" style="width: 100px" type="submit" value="提 交" />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</form>
</div>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292459.htm" target="_blank" title="大夫请温柔点！恶搞手游《疯狂外科医生3》"><span class="title1">大夫请温柔点！恶搞手游《疯狂外科医生3》</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">大鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:29:00</span></td>
  </tr>
  <tr>
    <td align="left"><p>作为一款<strong>经典恶搞医疗游戏</strong>，《疯狂外科医生》系列如今已成功推出了续作《<a class="f14_link" href="http://www.yingyong.so/app/28/14151.htm" target="_blank">疯狂外科医生3</a>（<a class="f14_link" href="http://www.yingyong.so/app/28/14151.htm" target="_blank">Amateur Surgeon 3</a>）》，依然还是那么重口味。玩家将扮演一名救死扶伤的外科医生，为各种病人进行手术，不过<strong>他的医疗器械实在奇葩，居然是订书钉、打火机和电锯等</strong>，我和我的小伙伴们都惊呆了！</p>
<p>《疯狂外科医生3》将疯狂恶搞进行到底，在游戏中没有见不到，只有<strong>想不到的各种外伤，比如心脏扎了牙刷，食人鱼钻进肚子里等</strong>。玩家要根据提示，利用各种暴力器械完成手术操作，例如先用订书钉把伤口拉小，再用打火机灼烧愈合伤口。如果在现实生活中这么做，病人肯定小命不保，但是在游戏里，你就放心大胆的发挥吧！</p>
<p>游戏里一共需要经历<strong>超过20项手术</strong>，在<strong>六个充满异国情调的迥异场地</strong>完成，其中<strong>罪犯、蝙蝠、机器人和教徒等无一幸免，都将成为被你摧残的手术对象</strong>。玩家还能对手术器材进行升级，更具威力的器材会让你救死扶伤的过程变得简单粗暴。</p>
<p>《疯狂外科医生3（Amateur Surgeon 3）》比较适合喜欢恶搞重口味的玩家，<strong>小清新玩家慎入</strong>。如果你喜欢猎奇，那它绝对是你的菜！拿起你的手机，亲自操刀来一场简单粗暴的外科手术吧！</p>
<p><strong>Android版</strong>下载地址：<a class="f14_link" href="http://www.yingyong.so/app/28/14151.htm" target="_blank">http://www.yingyong.so/app/28/14151.htm</a></p>
<p><strong>iOS版</strong>下载地址：<a class="f14_link" href="https://itunes.apple.com/cn/app/amateur-surgeon-3/id601538051?mt=8" target="_blank">https://itunes.apple.com/cn/app/amateur-surgeon-3/id601538051?mt=8</a></p>
<p align="center"><a href="/img/20140212/712b05a310574462bfe4b32802208f6a.png" target="_blank"><img alt="大夫请温柔点！恶搞手游《疯狂外科医生3》" src="/img/20140212/s_712b05a310574462bfe4b32802208f6a.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/f9837de76a3149c08bb866c4dfe5cd82.png" target="_blank"><img alt="大夫请温柔点！恶搞手游《疯狂外科医生3》" src="/img/20140212/s_f9837de76a3149c08bb866c4dfe5cd82.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/f82e33ee83024fca88f7666de0430ab9.png" target="_blank"><img alt="大夫请温柔点！恶搞手游《疯狂外科医生3》" src="/img/20140212/s_f82e33ee83024fca88f7666de0430ab9.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/a296df1ee4d84b6a9660e90d55180a9f.png" target="_blank"><img alt="大夫请温柔点！恶搞手游《疯狂外科医生3》" src="/img/20140212/s_a296df1ee4d84b6a9660e90d55180a9f.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/146f82363f8744d89ad0a9655c0be7ea.png" target="_blank"><img alt="大夫请温柔点！恶搞手游《疯狂外科医生3》" src="/img/20140212/s_146f82363f8744d89ad0a9655c0be7ea.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292457.htm" target="_blank" title="独家！HTC新渴望8系时尚靓机大曝光：漂亮"><span class="title1">独家！HTC新渴望8系时尚靓机大曝光：漂亮</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:15:11</span></td>
  </tr>
  <tr>
    <td align="left"><p>本次MWC 2014上，HTC会发布新旗舰M8吗？不会！</p>
<p>就在刚刚，<strong>有消息人士向笔者独家爆料称，HTC即将发布一款时尚靓机，其名称是新渴望8系，并且会率先在中国上市</strong>。</p>
<p>此外，消息人士还奉上了一张新渴望8系的官方宣传图。从图片上来看，它的外形设计还是很漂亮的，整机看起来非常精致，机身正面搭载双扬声器。</p>
<p>这款手机将定位中端市场（售价不会太贵），屏幕铁定大于5寸，其还会有红色、黄色、橙色以及蓝绿色可选。</p>
<p>大家觉得新渴望8系如何？</p>
<p><strong>Update：</strong>消息人士再次给出的消息称，该机配备的是一块5.5寸屏，摄像头组合为1300后置+500万前置（非Ultrapixel技术），内置美颜功能，。</p>
<p align="center"><a href="/img/20140212/14f3e9d0a39147528a5e5ee275e25bdc.jpg" target="_blank"><img alt="独家！HTC新渴望8系时尚靓机大曝光：漂亮" src="/img/20140212/s_14f3e9d0a39147528a5e5ee275e25bdc.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292455.htm" target="_blank" title="老电视别扔 这样用也很赞"><span class="title1">老电视别扔 这样用也很赞</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">刘艺</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:14:12</span></td>
  </tr>
  <tr>
    <td align="left"><p>老电视不能用了，只能扔废品回收站了，可是又很喜欢老电视的外观不舍得扔怎么办？没关系，只要加一点创意和一点巧思，就有大不同。</p>
<p>比如这个鱼缸，陶冶情操，猫咪还喜欢。</p>
<p align="center"><a href="/img/20140212/06065d62706241ebaa36d34fb3adce32.jpg" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_06065d62706241ebaa36d34fb3adce32.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/56e2b06c555b4344a21c4dcee6b77792.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_56e2b06c555b4344a21c4dcee6b77792.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/47b7809300314094b173da85502c8072.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_47b7809300314094b173da85502c8072.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/21de8d8ea6e34cedb9ace9c36db07567.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_21de8d8ea6e34cedb9ace9c36db07567.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/c0d7fd7dd1de4f018cbffae884f870d7.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_c0d7fd7dd1de4f018cbffae884f870d7.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/f9a84100af3c4a49b564a4fde89326c6.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_f9a84100af3c4a49b564a4fde89326c6.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/40062ce762b345aaa99ac4f06a0c1f2d.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_40062ce762b345aaa99ac4f06a0c1f2d.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/b45005ca07284672bd83143ba73762a3.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_b45005ca07284672bd83143ba73762a3.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/16202087a9f64848879cf2c9c9611ce2.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_16202087a9f64848879cf2c9c9611ce2.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/c41cd727194d4d9588644c9245e885e2.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_c41cd727194d4d9588644c9245e885e2.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/8e5a1487da9445b19f11c7b3a76cc21c.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_8e5a1487da9445b19f11c7b3a76cc21c.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/59ec435713604a78811ad3fb98dcf3cb.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_59ec435713604a78811ad3fb98dcf3cb.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/31cf799858324b06baf9d8a074cbc7b3.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_31cf799858324b06baf9d8a074cbc7b3.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/fcfe6e3120974b378341f138fefd88e7.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_fcfe6e3120974b378341f138fefd88e7.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/498c2dfaabbe46ba930c5e94d845fa1e.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_498c2dfaabbe46ba930c5e94d845fa1e.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/08beddc2e0d24f3d9b25c1fac8548402.png" target="_blank"><img alt="老电视别扔 这样用也很赞" src="/img/20140212/s_08beddc2e0d24f3d9b25c1fac8548402.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
别让我们的老电视只能散落在这里</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292454.htm" target="_blank" title="百度起诉360劫持流量 索赔65万元"><span class="title1">百度起诉360劫持流量 索赔65万元</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:11:32</span></td>
  </tr>
  <tr>
    <td align="left"><p><strong>2月10日下午，被称为马年互联网第一案的&ldquo;百度起诉360不正当竞争案&rdquo;在北京市朝阳区人民法院一审开庭</strong>。记者从法院了解到，由于360浏览器、360网址导航站通过设置热搜词恶意拦截百度、谷歌等搜索框的用户流量，百度认为360搜索严重违反《中华人民共和国反不正当竞争法》的有关规定，构成不正当竞争行为，因此向法院起诉并索赔65万元。法院未当庭宣判。</p>
<p>百度代理律师在庭审现场出示证据显示，<strong>用户一旦安装360极速浏览器并使用百度搜索引擎进行搜索，360都会自动弹出预先设置的10个热搜词</strong>，上网用户点击其中任一热搜词，都将自动跳转至360搜索的相关搜索结果页面，从而实施直接的流量劫持。360极速浏览器还通过其标签页所捆绑的网址导航站，实施同样方式的不正当竞争行为。证据同时显示，用户安装360极速浏览器使用谷歌搜索引擎时，也会出现同样劫持流量现象。</p>
<p>百度公司认为，奇虎360上述行为违反反不正当竞争法和公平竞争的商业原则，给网民搜索造成巨大不便，也给百度品牌带来巨大损失，要求法院判令奇虎360停止不正当竞争行为并消除影响，赔偿百度方面共计65万元经济损失。</p>
<p>360律师解释称，<strong>在360浏览器和360导航站的谷歌搜索框、百度搜索框上放置导向360网站的热搜词，目的是为了方便用户了解近期热点新闻</strong>。谷歌、百度都会出现相同热词链接，切换不同搜索引擎会出现相同热词链接，360并非故意仿冒、混淆搜索框和搜索结果。</p>
<p>百度代理律师对此回应称，这恰好说明了奇虎360窃取并劫持竞品流量的本质。因360热搜词的显示均无需搜索自动弹出，搜索后自动链接到己方搜索产品，通过浏览器等终端优势，未经许可便拦截对百度、谷歌发起的搜索请求，而热词与百度、谷歌数据系统的数据更新具有高度的相似性，360热搜词更具有误导性、不正当竞争性和侵权性，严重违背反不正当竞争法，不能成为免责的借口。</p>
<p>这并非360浏览器首次因劫持百度搜索流量遭到起诉。2013年12月，360恶意篡改百度搜索页面、劫持流量一案在北京市高级人民法院做出终审判决，360不正当竞争行为成立，赔偿百度40万元。这也是奇虎360近年遭遇的第14次败诉。</p>
<p>据Hitwise发布报告显示，百度搜索的市场份额稳定得保持在75%以上，360搜索则在15%左右徘徊，其份额大幅度依赖360浏览器市场优势推广维持。</p>
<p align="center"><img alt="百度起诉360劫持流量 索赔65万元" src="/img/20140212/f5af02e1a22e4be38934b31e74ecbf59.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292452.htm" target="_blank" title="“京东白条”明日公测：借你1.5万买东西"><span class="title1">“京东白条”明日公测：借你1.5万买东西</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:08:37</span></td>
  </tr>
  <tr>
    <td align="left"><p>京东在互联网金融领域正在加速布局，继推出3分钟融资到账的&ldquo;京保贝&rdquo;业务后，京东还将推出个人消费贷款服务&ldquo;京东白条&rdquo;。</p>
<p>据知情人士透露，<strong>这项名为&ldquo;白条&rdquo;的个人消费贷款服务，让用户可以凭借&ldquo;白条&rdquo;在京东购物时，可申请最高1.5万元的个人贷款用于支付</strong>，并在随后3个月至12个月内进行分期还款。</p>
<p>京东对&ldquo;白条&rdquo;业务也有风控机制，主要是授信给那些在京东又良好信用的用户，如在京东消费是否持续稳定，是单身还是家庭用户，相对来说，家庭用户的授信信誉会更高。</p>
<p><strong>这项业务将在本月13日至14日面向用户公测</strong>，2月15日至28日，获得公测资格的京东用户将获得&ldquo;白条&rdquo;，并可以在京东进行消费。</p>
<p>分析人士认为，京东推出&ldquo;白条&rdquo;服务后，可以使那些暂时没有资金，但信誉良好的用户提供方便，并在一定程度上能增长京东的销售额。这一概念对目前正筹备上市的京东而言，也有助于提升估值。</p>
<p align="center"><img alt="“京东白条”明日公测：借你1.5万买东西" src="/img/20140212/0dcbe9547cc149b2ad8be30916364c60.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292451.htm" target="_blank" title="Intel 15核心竟然是原生的！内核照首曝"><span class="title1">Intel 15核心竟然是原生的！内核照首曝</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">14:03:15</span></td>
  </tr>
  <tr>
    <td align="left"><p>国际固态电路会议ISSCC 2014上传出消息，Intel将在下周发布<a class="f14_link" href="http://news.mydrivers.com/1/291/291820.htm" target="_blank">首款15核心处理器Xeon E7 v2系列</a>(架构代号Ivy Bridge-EX、处理器代号Brickland、内核代号Ivytown)。</p>
<p><strong>这将是Intel的第一款奇数核心处理器</strong>(当然单核不算)，之前我们一直怀疑是16核心屏蔽而来的，但<strong>现在已经得到官方确认：它是原生的15核心。</strong></p>
<p>根据Intel公布的数据，Xeon E7将采用22nm HKMG工艺制造(九个金属层)，<strong>集成晶体管43.1亿个</strong>，<strong>运行频率最低1.4GHz、加速最高3.8GHz</strong>，三级缓存最多37.5MB(2.5MB&times;15)，热设计功耗最高150W(根据泄露消息应该是155W)，集成40条PCI-E信道。</p>
<p><strong>封装形式也是LGA2011，封装面积52.5&times;51.0＝2677平方毫米</strong>，基本上和Xeon E5 v2差不多。</p>
<p>不过内存通道数量仍未最终确定，有的说是四通道，也有的怀疑是三通道。</p>
<p align="center"><a href="/img/20140212/1dada2695d4e486f9c9896fefdb48c1f.png" target="_blank"><img alt="Intel 15核心竟然是原生的！内核照首曝" src="/img/20140212/s_1dada2695d4e486f9c9896fefdb48c1f.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>Xeon E7内核照片：可以清楚地看到<strong>分成三列的15核心</strong>(中间较深的绿色部分)，以及各个核心对应的三级缓存。</p>
<p align="center"><a href="/img/20140212/b885da96258b4b63907b011094c0e897.png" target="_blank"><img alt="Intel 15核心竟然是原生的！内核照首曝" src="/img/20140212/s_b885da96258b4b63907b011094c0e897.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>内核布局简图：<strong>核心与三级缓存分成三个部分，每两部分通过环形总线相连，总共需要三条环形总线。</strong><a class="f14_link" href="http://news.mydrivers.com/1/275/275643.htm" target="_blank">这种设计和最多12核心的Xeon E5-2600 v2是类似的</a>。</p>
<p>另外，<strong>Intel只设计了15核心一个Die，其他的12、10、8、6核心型号都是由此屏蔽而来的。&nbsp;</strong></p>
<p align="center"><strong><a href="/img/20140212/afd04da2b9564a95a63751479ce366ef.png" target="_blank"><img alt="Intel 15核心竟然是原生的！内核照首曝" src="/img/20140212/s_afd04da2b9564a95a63751479ce366ef.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></strong></p>
<p>时钟域三个，电压域五个。电压域之间还有电平位移器(Level Shifter)，在非时序关键路径中使用更低漏电率的晶体管，从而使得核心、非核心面积的利用率分别达到63％、90％，总的漏电率位总功耗的大约22％。</p>
<p align="center"><a href="/img/20140212/a1cc8326766f48f38d920803c49574ff.png" target="_blank"><img alt="Intel 15核心竟然是原生的！内核照首曝" src="/img/20140212/s_a1cc8326766f48f38d920803c49574ff.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292449.htm" target="_blank" title="魅族的成熟：从J.Wong臣服开始 "><span class="title1">魅族的成熟：从J.Wong臣服开始 </span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">13:58:05</span></td>
  </tr>
  <tr>
    <td align="left"><p>2014年之前，黄章留给行业的印象似乎只有2个关键词：偏执和偏激，这2个关键词作为黄章性格的真实写照，也分别映衬了黄章不同程度上的为人特征：前者昭示了其对产品品质执着的精神，后者也显示了其在个人素质方面低劣粗暴的素质。</p>
<p>如果时间可以停留的话，那么今天的黄章将会将这2个关键词贯彻到底，继续蜗居在没人知道的南方城市里颐养天年，然后在夜深人静的时候在魅族论坛上对指出魅族缺点的用户嚷滚开。作为缔造了魅族品牌并且百分之百控股的人来说，这是黄章的个人生活选择，即便这个发迹之后的黄章，完全侵袭了传统爆发富的典型特征：自傲、不修起码的素养边幅，但也因为魅族这家公司是完全属于他的而有这个自由。</p>
<p>所以2014年之前的黄章，可以不重视产品、市场、用户、体验等等因素，只重视自我内心的陶醉，成长在自我装饰的精神国度里幸福着，对异议者说滚开，对同行友商破口大骂甚至是人身诋毁。这一切的一切都只因为魅族是黄章一个人的魅族，是属于黄章一个人的公司和玩具，因此黄章想骂人就骂人想发小脾气就发小脾气，完了之后还有人会自发地为其背书说这是&ldquo;个性&rdquo;。</p>
<p>但是随着2014年的悄然来临，黄章决定浮出了。</p>
<p>在媒体的报道材料中，<strong>黄章首先是回归到公司坐班上班了，其次是要打破天窗接受融资了，再次是要将股份期权免费给员工了，再有就是首次露面给员工讲话2个小时了，当然最重要的是将魅族论坛里的过激言论删除了</strong>&hellip;&hellip;仅仅只是一个春节的时间，黄章的前后发生了天翻地覆的变化，这起变化被人说成是黄章开始开放了，开始用互联网的方式经营和发展魅族公司了，也是魅族即将多元化发展的节奏了。</p>
<p>不管这些&ldquo;阳谋论&rdquo;是否真实，但黄章的这些举动确实在说明，以后的魅族可能确实会正常化发展了。</p>
<p>对于魅族而言，接受融资、开放股权激励、引荐人才、展开多方合作等等这些常规行为，无一都是正常商业公司的普遍现象，这对于魅族的产品和公司的正常化发展，也都是良好的积极信号。特别对于魅族的未来发展，<strong>这种开放的讯号，也有可能会让魅族从此告别过去封闭式的环境</strong>，进而借助这种开放模式迎来更大发展机会，进而拉开与其他同行友商的竞争距离。</p>
<p>对于黄章本人而言，这种开放其实是难能可贵的，一个在自我意淫的国度里呆得太久了的人，几乎不太可能会接受外在的世界，但是在当前激烈的智能手机大战中，惨烈的市场竞争想必已经让黄章开始有了焦急的心态，因此在舍弃自己的做人原则和生活方式的情况下，黄章选择了复出。<strong>这种决策对于黄章本人而言，我认为也是一个明智的选择，毕竟自己的公司再不认真打理的话，中华酷联小迟早干掉魅族，也未必不是不可能的事。</strong></p>
<p>从这点来说，黄章舍弃了个人自身身上的各种棱角，让自己完全融入正常的商业生活，进而引领魅族前进，这是黄章有别于2014年之前的黄章的地方，同时也是黄章告别个性，臣服现实，屈从商业社会的显著标志。而这种变化和举动对于一个商业人物和商业公司而言，正是最基本的正常方式，只不过这种在其他公司和领导人看来再正常不过的方式，在魅族和黄章身上居然如此费力。</p>
<p>不过好在魅族和黄章总算迷途知返。</p>
<p>对于商道，个性和规则几乎都不太可能达到兼得地步，成长历程不同于其他企业领导人的黄章，更应是在发现遵循自己做产品和梦想的路途中遭遇到鱼和熊掌的问题之后，才知晓有些商业规则需要放弃一些自身的棱角。因此今天黄章做下的这个决定，以魅族更好的发展为前景，以放弃自身的个性棱角为牺牲，这种选择所幸还并不晚。</p>
<p>但至于黄章是否有能耐将魅族搀扶上赶超中华酷联小的快速轨道，这一切让时间来说了算。</p>
<p align="center"><img alt="魅族的成熟：从J.Wong臣服开始 " src="/img/20140212/6e1b7adbb2bf4466942533026b97372c.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292446.htm" target="_blank" title="多家比特币交易所暂停提现：因遭黑客攻击"><span class="title1">多家比特币交易所暂停提现：因遭黑客攻击</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">13:46:53</span></td>
  </tr>
  <tr>
    <td align="left"><p>据《华尔街日报》报道，<strong>两家大型比特币交易所Bitstamp和BTC-e周二遭到黑客攻击，现已暂停提现服务</strong>。</p>
<p>斯洛文尼亚的比特币交易所Bitstamp在其网站上发表声明称，因遭受DDOS(拒绝服务)攻击，交易所将暂停提现服务以解决这一问题；保加利亚的比特币交易所BTC-e称遭受了同样的攻击，部分交易的授信工作将被延迟，也有一些交易可以正常进行，但提现服务已经暂停。</p>
<p>根据比特币交易追踪机构bitcoincharts.com的信息，上述两家交易所的交易额占全球总交易量的56%。</p>
<p>此次遭受攻击的比特币交易所之一Bitstamp称问题与交易可锻性(transaction malleability)有关，这与Mt.Gox上周遇到的问题相同；BTC-e没有对问题的详情进行说明。</p>
<p><strong>Mt.Gox此前对此类漏洞进行了详细说明，称比特币软件的一项漏洞可以导致交易细节被修改</strong>。具体来说，有人可以通过修改交易细节隐藏比特币钱包之间的比特币转账记录，这可能导致转账反复发生。Mt.Gox还表示，该问题会影响所有将比特币发送给第三方的交易。</p>
<p>《华尔街日报》评论称，这些攻击活动显示出比特币的脆弱。这种虚拟货币由计算机生成，没有央行的支持，一旦出现问题不会有政府介入解决。</p>
<p align="center"><a href="/img/20140212/372e51c206c242bea8f14251a2d22186.jpg" target="_blank"><img alt="多家比特币交易所暂停提现：因遭黑客攻击" src="/img/20140212/s_372e51c206c242bea8f14251a2d22186.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292445.htm" target="_blank" title="听 那是雨的声音"><span class="title1">听 那是雨的声音</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">随心</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">13:34:31</span></td>
  </tr>
  <tr>
    <td align="left"><p>吃过午饭，现在正是昏昏欲睡的时候。是不是期待窗外一场淅淅沥沥的小雨，冲散雾霾、让自己疲惫浮躁的心静一静呢？</p>
<p><a class="f14_link" href="http://www.rainymood.com/" target="_blank">Rainy Mood</a>，正如其名，是一个让你体验雨中即视感的网站，这里不仅有非常逼真的下雨声，远处还有隆隆的雷声。另外，整个页面也是一块沾满雨珠的玻璃，甚至有雨珠下滑的效果。还等什么，打开网站倾听窗外的雨声吧！</p>
<p>温馨提示：因为媒体内容丰富，可能打开有些慢。</p>
<p>如果喜欢这段音乐，可以<a class="f14_link" href="http://vdisk.weibo.com/s/DinJ2hlM5-Dv" target="_blank">点此</a>下载，大小20.4MB，长度30分20秒。</p>
<p align="center"><img alt="听 那是雨的声音" src="/img/20140212/62959d38ae2f4adf923f4d55db669be3.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292444.htm" target="_blank" title="铁杆诺粉11年心血：诺基亚拍照进化史"><span class="title1">铁杆诺粉11年心血：诺基亚拍照进化史</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">13:34:00</span></td>
  </tr>
  <tr>
    <td align="left"><p>由于现今的诺基亚Lumia手机主打是拍照功能，因此诺基亚又被戏称是卖相机的厂商。那么，诺基亚的手机拍照功能是如何逐步成为金字招牌的呢？</p>
<p>作为资深诺粉，国外用户Richard Dorman从2003年起，每年都要用当年的热门诺基亚手机来拍几张照。<strong>日前他晒出了自己历时11年拍摄的风景照，拍摄机型从2003年的诺基亚7650开始，直到2013年的Lumia 1020。</strong>我们也一同来看看诺基亚拍照手机的进化历程。</p>
<p align="center"><a href="/img/20140212/4344fa37f6bb44bfa88da1ec220bf883.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_4344fa37f6bb44bfa88da1ec220bf883.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚7650（2003年）</p>
<p align="center"><a href="/img/20140212/21dd4500f4514d75b6636c21249b2f58.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_21dd4500f4514d75b6636c21249b2f58.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚7650（2003年）</p>
<p align="center"><a href="/img/20140212/964c07f6ae34483f84284ea66bbbec73.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_964c07f6ae34483f84284ea66bbbec73.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚6600（2004年）</p>
<p align="center"><a href="/img/20140212/ea139308347645e3981cbe007b286f67.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_ea139308347645e3981cbe007b286f67.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚6600（2004年）</p>
<p align="center"><a href="/img/20140212/6e9862c0626347f0a01893f5b450b9ff.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_6e9862c0626347f0a01893f5b450b9ff.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚7610（2005年）</p>
<p align="center"><a href="/img/20140212/6ae690ca808742babbdb196a8bf27ea0.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_6ae690ca808742babbdb196a8bf27ea0.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚7610（2005年）</p>
<p align="center"><a href="/img/20140212/8addfbc5d20c453d96ece9389b96f14e.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_8addfbc5d20c453d96ece9389b96f14e.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚N70（2006年）</p>
<p align="center"><a href="/img/20140212/6e5c1210715040c4b540c85f11ed434b.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_6e5c1210715040c4b540c85f11ed434b.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚N70（2006年）</p>
<p align="center"><a href="/img/20140212/703aa6657ee24caa9278b684f9fded34.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_703aa6657ee24caa9278b684f9fded34.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚N93（2006年）</p>
<p align="center"><a href="/img/20140212/e3338d7d4c2541cb977e010e169e02ff.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_e3338d7d4c2541cb977e010e169e02ff.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚N95（2007年）</p>
<p align="center"><a href="/img/20140212/06a541c32b564ac882808cf2505b41fa.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_06a541c32b564ac882808cf2505b41fa.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚N95（2007年）</p>
<p align="center"><a href="/img/20140212/f96abb590aba4ab5b6b3a9cd402043b3.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_f96abb590aba4ab5b6b3a9cd402043b3.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚N95-8GB（2008年）</p>
<p align="center"><a href="/img/20140212/595847e8cdf347fb956e46ad35800d02.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_595847e8cdf347fb956e46ad35800d02.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚N97（2009年）</p>
<p align="center"><a href="/img/20140212/47097e79bc5b4e92b86ffabe28f9cd56.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_47097e79bc5b4e92b86ffabe28f9cd56.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚N8原型机（2010年）</p>
<p align="center"><a href="/img/20140212/3f4ffc28b5644a8baa10995f021ac70f.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_3f4ffc28b5644a8baa10995f021ac70f.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚N8（2010年）</p>
<p align="center"><a href="/img/20140212/8f57cd70ca2f4a0184e14216de4df8dc.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_8f57cd70ca2f4a0184e14216de4df8dc.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚N8（2010年）</p>
<p align="center"><a href="/img/20140212/fb5904e7e2f24d70a5fe7ce5bb51024a.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_fb5904e7e2f24d70a5fe7ce5bb51024a.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚808 PureView（2012年）</p>
<p align="center"><a href="/img/20140212/4719776582454154980edc1a397fa042.jpg" target="_blank"><img alt="铁杆诺粉11年心血：诺基亚拍照进化史" src="/img/20140212/s_4719776582454154980edc1a397fa042.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
诺基亚Lumia 1020（2013年）</p>
<p><strong>更多诺基亚手机拍照样张打包下载</strong>：<a class="f14_link" href="http://bbs.mydrivers.com/thread-353946-1-1.html" target="_blank">http://bbs.mydrivers.com/thread-353946-1-1.html</a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292443.htm" target="_blank" title="玩遍游戏放弃保送清华：斯坦福录取＋百万奖学金"><span class="title1">玩遍游戏放弃保送清华：斯坦福录取＋百万奖学金</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">13:33:47</span></td>
  </tr>
  <tr>
    <td align="left"><p>&ldquo;<strong>市面上能见到的游戏我都玩过了。</strong>&rdquo;&ldquo;因为我相信我是独特的，所以我选择再申请一次斯坦福。&rdquo;2014年2月10日晚8点过，18岁的李凌霄从工作室出来。高中毕业后，他和朋友开办工作室，开发游戏软件，每天从上午9点工作到晚上7点。</p>
<p>可能李凌霄的高中同学还不知道，这个全国信息学奥赛金牌获得者，<strong>在放弃了清华大学、加州伯克利分校之后，他两次申请斯坦福大学，并最终被斯坦福大学以每年4.2万美元(四年共计超过100万元人民币)的奖学金录取。</strong>他也由此成为近几年来，四川唯一在美国大学EA(本科提前录取)阶段获得斯坦福大学录取和奖学金的学生。</p>
<p><strong>追梦：钟情斯坦福 自主申请被拒</strong></p>
<p>&ldquo;我不太喜欢张扬。&rdquo;成都七中毕业的李凌霄介绍，虽然放弃了清华大学、加州伯克利分校，首次申请也失利了，但因从小对计算机非常热爱及执着申请，终于在2013年冬天赢得了第二次与斯坦福的美丽邂逅&hellip;&hellip;</p>
<p>&ldquo;我妈妈终于可以睡好觉了。&rdquo;10日晚上，刚从工作室出来的李凌霄开着玩笑说。他个子超过一米八，戴着黑边眼镜。高中毕业后，他和几个朋友一起创业，开办工作室，开发游戏软件，每天从上午9点工作到晚上7点。</p>
<p>&ldquo;我没去过美国，关于留学的信息都是二手的。&rdquo;初中时，斯坦福教授唐&middot;克鲁斯的一本计算机著作，让李凌霄第一次被书中技术和艺术的结合震惊。自那之后，便深藏了一个留学梦。</p>
<p>高三竞赛保送后，李凌霄才真正考虑出国的事。他选择了DIY(自主申请)这种普通高中生看来最坎坷的申请方式：&ldquo;没老师，向其他出国同学寻求帮助，同时自己上学校主页，自己准备语言考试。虽然机会小，但这种绝境激发了我冒险和挑战的本能。&rdquo;</p>
<p>斯坦福大学、麻省理工学院、加州伯克利分校等，第一次申请，他所选择的全是计算机专业世界排名前列的大牛校，同时还要求奖学金。&ldquo;我特别钟情斯坦福，当时不了解情况，很多名校明确写了不会给国际学生奖，而我需要这笔钱。&rdquo;李凌霄说，<strong>2012年第一次被加州伯克利录取却没奖学金时，他选择放弃。</strong></p>
<p><strong>圆梦：总结受挫经验 等来录取通知书</strong></p>
<p>自从决定要重新申请后，李凌霄开始重新审视自己。去年他在托福和SAT考试中分别取得了115分和2240分(拼分2310分)的好成绩。</p>
<p>李凌霄总结了自己申请受挫的经验：<strong>不清楚申请学校的奖学金政策，文书太过专业，立意范围狭窄，没展现出独特地方，而差异和个性正是国外大学想在国际学生中寻找的。</strong></p>
<p>于是第二次申请时，他把重点放在了更自我的层面上：对更自由学术环境的追求，对未知的好奇和渴望创造，以及冒险和求是精神。</p>
<p>幸运的是，这一次，<strong>斯坦福选择了他&mdash;&mdash;2013年12月14日，给他发来了录取通知，以及每年4.2万美元学费全免(折合人民币约25万元)的奖学金</strong>，成为近几年来，四川唯一在美国大学EA(本科提前录取)阶段获得斯坦福大学录取和奖学金的学生。</p>
<p>&ldquo;在所有的申请者中，我相信我的经历是独特的，放弃了保送选择出国，选择了一年的gap year(美国俗称间隔年，指高中毕业后暂时不上大学，而去游历或者体验生活的一年)。&rdquo;李凌霄说，&ldquo;<strong>成绩在申请最好的学校时几乎没什么用，独特经历才是最重要的，它定义了你是怎么样一个人</strong>。&rdquo;</p>
<p><strong>谈梦：老师开绿灯 他自学编程</strong></p>
<p>&ldquo;对计算机的兴趣，不是简单的喜欢，而是长久的坚持，坚持越久，就能在某个领域洞察到一般人看不到的东西。&rdquo;说到对计算机的兴趣，李凌霄说这还得从小学玩游戏说起，&ldquo;市面上能见到的游戏我都玩过了。&rdquo;虽然对学习没什么帮助，但却让他从小萌生了制作游戏和编程的念头。</p>
<p>在乐山读完小学后，他考上了成都外国语学校初中，&ldquo;刚开始成绩不怎么好，环境陌生，每个周末就盼望着坐两个小时校车回家。&rdquo;然而正是不适应激发了他的潜能。</p>
<p>&ldquo;我最感谢的是我的初中班主任，她教我们每天写成长日记、学会自我管理，也为我的兴趣爱好开绿灯，破例允许我每周二、周四下课后单独使用学校计算机房，在身边没任何资源的情况下自学计算机编程。&rdquo;</p>
<p>高中后，李凌霄对计算机的热情持续了下来。他是各种计算机竞赛的得奖专业户，全国信息学奥林匹克竞赛金牌获得者，因计算机编程比赛的优异成绩，获得前往俄罗斯圣彼得堡免费旅游的机会。</p>
<p>&ldquo;参加竞赛让我学会了独立思考问题，更让我在一条不同于其他人选择的路上发现自我价值，而不受所谓主流思维的影响。&rdquo;</p>
<p><strong>幕后故事：开明爸爸支持&ldquo;不上大学也无所谓&rdquo;</strong></p>
<p>李凌霄出生在乐山一个工薪家庭，爸爸对他的求学选择一直是静观其变。保送清华又放弃，申请到加州伯克利这样的牛校因为没有奖学金而放弃，高中毕业到北京实习因为工作环境与自己的求职标准不符而放弃，回到成都白手创业&hellip;&hellip;</p>
<p>这个开明的爸爸甚至对儿子大开方便之门说：&ldquo;你不上大学也无所谓。&rdquo;爸爸用自己的大学经历告诉他，<strong>相信自己的想法，并尝试着固执地走下去，这才是对自己人生负责任的选择。</strong></p>
<p>&ldquo;未来要发生的事情谁知道呢，我不擅长做长期计划，因为有太多不确定，也许我到了斯坦福会选计算机，也许会选择别的专业，毕竟只有自己体验过才知道究竟是什么感觉。&rdquo; 华西都市报记者张峥</p>
<p><strong>记者手记：以勇气，致青春</strong></p>
<p>&ldquo;不想被采访，是因为我不想打破现在平静的生活。&rdquo;虽然李凌霄一再说，自己很多想法很幼稚，但在长达两小时的对话里，始终能让人感受到他的自信与独立。</p>
<p>&ldquo;我承认我选择搞竞赛一定程度上是想回避高考，因竞赛比高考容易，至少对像我这样有相似爱好的人来说。&rdquo;</p>
<p>&ldquo;我不想去大公司实习，因为这很难在短时间内体现出个体价值。&rdquo;</p>
<p>&ldquo;申请是个双向选择的过程，因为我的经历，说明我很适合斯坦福开放自由的校风，所以我选择再次申请。&rdquo;&hellip;&hellip;</p>
<p>的确，青春就是一场肆无忌惮的碰壁，一次不计代价的赌约。从在质疑声中咬紧牙关的&ldquo;逆袭&rdquo;，到重新审视一个曾经被拒录的申请人的材料，不论对18岁的李凌霄、还是建校120余年的斯坦福来说，都是一种勇气。</p>
<p align="center"><img alt="男孩玩遍游戏放弃保送清华：斯坦福大学录取＋百万奖学金" src="/img/20140212/6504a00eb57b42169a41d305e5bb0b80.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292442.htm" target="_blank" title="外国网购为何不怕被“盗刷”？"><span class="title1">外国网购为何不怕被“盗刷”？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小路</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">13:32:02</span></td>
  </tr>
  <tr>
    <td align="left"><p>2月11日是国际互联网安全日。随着网上消费规模的不断增长，网上支付安全问题日益受到关注。在全世界网购最发达的英国、美国和日本等国，网络支付安全也不容乐观，但它们通过立法和技术等多种手段保护消费者网购时不被&ldquo;盗刷&rdquo;，其中很多内容都值得我们参考借鉴。</p>
<p><strong>英国：网上支付安全有妙招</strong></p>
<p>英国广播公司在刚刚过去的圣诞新年促销季里给消费者支招，教他们如何安全、顺利地完成网上支付。首先，要在正规的网站购买产品，要<strong>注意支付时的网址是否以&ldquo;https://&rdquo;开头，其中的字母&ldquo;s&rdquo;就代表安全，进行支付时，地址栏中还会出现一个小的挂锁图标，也代表交易安全。（国内一些金融类网站也会出现类似的提示）</strong></p>
<p>网购的一大优势就是价格往往更为低廉，不过英国广播公司也提醒消费者，网购也像传统购物方式一样，切莫贪图便宜。如果看到某个网站提供的商品&ldquo;太便宜&rdquo;，那么其所售商品就很有盗版等非法嫌疑，所提供的支付方式也可能有安全隐患。</p>
<p>除了加强媒体宣传，英国政府还专门发起成立了名为&ldquo;获得在线安全&rdquo;的服务性组织。在该组织的官方网站上，就可以找到关于各类网上支付的专家建议和注意事项。在这些建议中，专家强调最多的就是采用信誉度高的专业支付方式，而不要直接将钱款打给卖家。</p>
<p>在英国，有许多十分成熟的网上支付方式，比较知名的有苏格兰皇家银行旗下的&ldquo;世界支付系统&rdquo;、在欧洲地区使用率极高的&ldquo;Moneybookers&rdquo;等。这些支付方式往往都有比较悠久的历史和专业的团队，通过不断完善和改进相关技术，保障支付安全，并对大额支付有所限制。</p>
<p>这些支付方式在技术上有各自的安全措施，比如Moneybookers规定，用户必须在激活认证自己的注册账号后才能开始使用其服务，这就有助于遏制网络支付诈骗。</p>
<p><strong>美国：银行主要负责保障网上支付安全</strong></p>
<p>与中国不一样的是，美国网上支付是从发展成熟的线下信用卡体系延伸至互联网的，因而多沿用现有传统银行和支付管理法律来监管网上支付，由银行主要承担保障网上支付安全的责任。</p>
<p>与美国消费者支付安全直接相关的法律有两部，一部是1968年关于信用卡的《借贷真实法》。该法律规定，如果信用卡被非授权使用，持卡人的责任最多为50美元；另一部是1977年关于借记卡的《资金电子划拨法》。该法律规定，如果消费者在借记卡被非授权使用两个工作日内通知银行，其损失将被限定在50美元以内，两个工作日以外则是最多500美元。</p>
<p>另外，如果发生了非授权使用情况但借记卡本身并没有丢失或被盗窃，持卡人在收到对账单60天以内通知银行，则不用负任何责任；只有消费者收到银行对账单60天内，仍不通知银行卡被非授权使用，则损失将可能无法追回。</p>
<p><strong>日本：三管齐下保障网上支付安全</strong></p>
<p>日本主要从维护手机安全、保护个人金融信息、注重密码保护三方面来保障网上支付安全。</p>
<p>日本最大的移动运营商Docomo通信公司2011年开始提供手机病毒扫描服务，对所有的程序、文件夹、短信、压缩文件、邮件等都进行检查。如果用户发现手机丢失或被盗，可以与通信公司联系，通过远程操作，锁定包括手机钱包在内的所有功能。</p>
<p>日本手机实行实名制，一个人的名下只能在Docomo公司申请5部手机。个人支付手机费，需要去门市部找营业员刷卡缴费。如果驾照等确认本人信息的证件涉嫌伪造，公司会通知警方。如果用户不回应警方对其情况的调查，移动运营商会停止其电话使用，并且在各运营商间共享信息，在其再次申请手机时严格审查资格。</p>
<p>从世界范围来看，日本社会发生个人金融信息遭到泄露、盗用事件的频度偏低。这得益于日本政府&ldquo;立法先行&rdquo;的一贯理念。日本通过《个人信息保护法》《金融商品交易法》以及《消费者契约法》等法律，使用户的个人信息得到了金融机构的足够重视。</p>
<p>为防止信用卡被盗刷的行为，日本信用卡公司都会对用户进行详细的指导，包括对密码严格保密。对于生日、电话号码、4位相同的数字等容易被推测出的密码，信用卡公司会推荐用户更改这种密码。而更改密码时，用户要先给公司打电话，公司会发送申请变更密码的申请书和回复用的信封，不会简单地在电话中就同意变更密码。</p>
<p>网络支付安全涉及范围广、环节多。有美国专家认为，与偷窃数据的网络犯罪分子的斗争将是&ldquo;一场持久性战争&rdquo;。</p>
<p align="center"><img alt="外国网购为何不怕被“盗刷”？" src="/img/20140212/49cf015768784c929f5ee26e862a002d.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292441.htm" target="_blank" title="SanDisk发天价U盘：速度堪比SSD"><span class="title1">SanDisk发天价U盘：速度堪比SSD</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">13:25:18</span></td>
  </tr>
  <tr>
    <td align="left"><p>SanDisk在近几天动作非常多，推出了一系列U盘产品，其中就包括采用全金属外壳的高端产品Extreme Pro USB 3.0（CZ88）。</p>
<p>该U盘按此用了伸缩式设计，能够避免用户弄丢盖子这种经常性情况的发生，另外其还<strong>采用了金属外壳设计，很有质感</strong>。</p>
<p>容量方面，该U盘目前仅有128GB一款型号可供选择，我们并不排除以后推出更多容量版本的可能。它的<strong>读写速度分别达到了260/MB/s和240MB/s</strong>，已经能够和部分中低端SSD相提并论了。</p>
<p>强悍的性能带来的自然是高昂的售价，<strong>其官方零售价为199.99美元，约合人民币1212元</strong>，甚至能买一块240GB的SSD了。</p>
<p>最后笔者在通过查询之后发现该U盘居然在国内已经提前开卖了，<strong>京东上的价格为1188元</strong>，显示已优惠700元，也就是说行货价格应该是1888元。</p>
<p align="center"><a href="/img/20140212/e93c14880f5b47969cea902346c0e3ed.jpg" target="_blank"><img alt="SanDisk发天价U盘：速度堪比SSD" src="/img/20140212/s_e93c14880f5b47969cea902346c0e3ed.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/b15236ea2f104f2b959ee8b5a0fd0492.jpg" target="_blank"><img alt="SanDisk发天价U盘：速度堪比SSD" src="/img/20140212/s_b15236ea2f104f2b959ee8b5a0fd0492.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292440.htm" target="_blank" title="ARM也有自己的视频、显示处理器了"><span class="title1">ARM也有自己的视频、显示处理器了</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">13:22:47</span></td>
  </tr>
  <tr>
    <td align="left"><p>Imagination PowerVR绝对是移动GPU领域无可争议的霸主，同时还有PowerVR VXD解码器、PowerVR VXE编码器作为补充。ARM Mali则是凭借自家Cortex处理器的提携取得了不小的进步，现在也有了自己的视频处理器、显示处理器。</p>
<p>首款视频处理器是&ldquo;<strong>Mali-V500</strong>&rdquo;，<strong>支持1080p60、4k120分辨率的高清视频，编码支持H.264、VP8，解码支持H.264、H.263、MPEG4、MPEG2、VC-1/WMV、Real、VP8。</strong></p>
<p><strong>Mali-V500最多可集成八个处理核心</strong>(最少一个)，运行频率600MHz，总线支持AMBA3 AXI、AMBA4 ACE Lite，可搭配不同的Cortex处理器，并支持MMU虚拟内存。</p>
<p>该视频处理器号称是业内同类产品中面积最小的。</p>
<p align="center"><a href="/img/20140212/5bc16d4ef6ea4ba8a63e0c901d771f8a.jpg" target="_blank"><img alt="ARM也有自己的视频、显示处理器了" src="/img/20140212/s_5bc16d4ef6ea4ba8a63e0c901d771f8a.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><img alt="ARM也有自己的视频、显示处理器了" src="/img/20140212/s_65e78759f0c74435b9072718e85063fa.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p>显示处理器则是&ldquo;<strong>Mali-DP500</strong>&rdquo;，用于<strong>视频画面的合成(多个Alpha混合层)、旋转(90/180/270度)、色彩转换与增强、缩放等处理，从而让GPU从这些任务中摆脱出来，并支持立体3D、4K超高清、12-bit色彩通道、ARM TrustZone安全技术。</strong></p>
<p>该处理器支持单引擎单显示、双引擎双显示，兼容VESA、CEA、HDMI、miniDP。</p>
<p>它的核心面积同样很小，而通过支持ARM AFBC(帧缓冲压缩)技术，可以大大降低系统功耗，且安卓系统驱动原生支持。</p>
<p>Mali-V500、Mali-DP500都推荐搭配<a class="f14_link" href="http://news.mydrivers.com/1/292/292331.htm" target="_blank">刚刚发布的Cortex-A17处理器</a>，以及Mali-T720 GPU。</p>
<p align="center"><a href="/img/20140212/a1e4e377577b4154b35717a3a10b9f21.png" target="_blank"><img alt="ARM也有自己的视频、显示处理器了" src="/img/20140212/s_a1e4e377577b4154b35717a3a10b9f21.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292439.htm" target="_blank" title="J.Wong来新浪微博了！"><span class="title1">J.Wong来新浪微博了！</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">12:58:59</span></td>
  </tr>
  <tr>
    <td align="left"><p>以往总是&ldquo;神龙见首不见尾&rdquo;的J.Wong黄章终于公开露面了。<strong>在重新回到魅族上任CEO之后，黄章删除了以往论坛的过激言论（比如&ldquo;不喜欢就滚&rdquo;），并以正式身份开通新浪微博</strong>。</p>
<p><strong>黄章的微博认证为&ldquo;魅族科技董事长兼总裁 黄章&rdquo;，目前尚未发布微博，但粉丝已经瞬间破万</strong>。</p>
<p>不少网友调侃，呼吁雷军和黄章互粉一下。</p>
<p>顺便说一句，今天上午魅族科技发微博称&ldquo;今天公司有贵客造访&ldquo;，不少人都猜测是格力。<strong>就在刚刚，Flyme音乐放出了一张图片，暗示&ldquo;董小姐&rdquo;来了。（即格力总裁董明珠）</strong></p>
<p>真的像黄章说的那样，格力要投资魅族吗？</p>
<p align="center"><a href="/img/20140212/911c6793935e428486df41b67d2e5a47.jpg" target="_blank"><img alt="J.Wong黄章正式开通新浪微博" src="/img/20140212/s_911c6793935e428486df41b67d2e5a47.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/ecb6eabb05834f25b169fb4898321c7b.jpg" target="_blank"><img alt="J.Wong黄章正式开通新浪微博" src="/img/20140212/s_ecb6eabb05834f25b169fb4898321c7b.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/e9442cd2763f4237af03d2e0415c34af.jpg" target="_blank"><img alt="J.Wong黄章正式开通新浪微博" src="/img/20140212/s_e9442cd2763f4237af03d2e0415c34af.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292437.htm" target="_blank" title="中移动4G：几分钟1GB流量不翼而飞"><span class="title1">中移动4G：几分钟1GB流量不翼而飞</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">随心</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">12:31:19</span></td>
  </tr>
  <tr>
    <td align="left"><p>用手机4G不慎，会让自己一夜&ldquo;破产&rdquo;?这几天，微博和微信上都在转载一种令人有些恐慌的说法：由于手机4G上网速度非常快，睡一觉起来，发现4G没关，大量流量偷跑，自己的房子是运营商的了。&nbsp; <strong>近日，在福州发生了这样一件类似的事情，近1GB的流量在短短的几分钟内就被消耗完毕</strong>，幸亏市民及时关闭了数据开关，才&ldquo;保住了房子&rdquo;。<strong>运营商表示，将减免该市民的流量，但并不清楚流量到底去哪儿了。</strong>同时，运营商也表示，4G流量偷跑危及房子，只是网友调侃的说法，并不靠谱。</p>
<p><strong>手机已经锁屏 流量还是不翼而飞</strong></p>
<p>上周末的晚上，福州郑女士准备睡觉时，苹果手机突然密集发来多条短信。郑女士打开短信一看，原来都是运营商发来的提醒信息，不断提示自己的4G套餐流量即将用完，而最后几条短信则提示，套餐已经用完，之后使用的流量都需要进行付费。</p>
<p>郑女士很吃惊，因为在之前的一个多小时，她根本没有使用手机，而当短信密集发来的时候，她的苹果手机依然是在锁屏状态。她觉得，只要是在锁屏状态中，即便手机一些应用软件会上网并推送消息，也不至于那么多的流量会全部用完。</p>
<p>于是，郑女士向运营商查询流量使用详单，这份详单显示，当天晚上近10点开始，郑女士的手机每隔10多秒钟就产生10MB多的流量，很快，近1GB的套餐流量就已经消耗殆尽。</p>
<p><strong>消耗大量流量 运营商却无记录</strong></p>
<p>那么，郑女士的手机为何会无故产生巨额流量呢?运营商方面对记者表示，郑女士的手机确实在很短时间内就消耗了大量的流量，不过，运营商无法确定当时产生的这些流量使用在什么地方。</p>
<p>运营商方面表示，通过查询郑女士的上网记录发现，在产生这笔近1G的流量前，郑女士的网络使用记录可以查得到，一般是使用微信、网页、天气软件等服务，但是<strong>这笔短期之内出现的大流量的去向则没有具体记录，有可能是郑女士手机中的软件在自动更新。</strong></p>
<p>对于运营商的说法，郑女士质疑说，在如此短时间内出现巨额的流量费，对于已经锁屏了的苹果手机来说，并不大可能。记者在郑女士的手机上看到，手机中仅有20多个应用软件，<strong>即便所有软件全部自动更新，那也只要约100MB，而且在没有输入密码的情况下，手机软件是无法自动更新的</strong>，而如果手机是在后台运行软件，那使用的网络流量就更少了。</p>
<p>对此，运营商方面表示，可以对郑女士有疑义的部分流量进行减免。</p>
<p><strong>流量偷跑，费用相当一套房子？运营商：网友吐槽过于夸张</strong></p>
<p>郑女士告诉记者，最近网上都在流传一个说法，就是4G上网速度非常快，睡觉时如果忘关手机的数据开关，有可能醒来的时候连房子都是运营商的了，幸好自己在接到运营商发来的提醒短信后立即关闭了上网流量开关。她说，如果当时没有关闭流量开关，真有可能要付出巨额通信费。&ldquo;不过还好保住了房子。&rdquo;她开玩笑地说。</p>
<p>对于这一说法，运营商方面表示，这样的说法只是网友吐槽，现实生活中无法实现。记者从目前开通4G服务的两大运营商处了解到，<a class="f14_link" href="http://news.mydrivers.com/1/291/291460.htm" target="_blank">中国移动4G套餐目前执行500元流量封顶政策</a>，在手机使用流量额外超过500元后，流量功能将自动关停，如果用户需要再用，则需要主动打开;而中国联通的大部分套餐则执行超过500元或6GB就关闭流量功能，如果需继续上网则需主动打开，而如果超过15GB则直接关闭当月流量功能。</p>
<p>不过，虽然流量&ldquo;偷跑&rdquo;不会危及房子，但是消费者依然有可能损失话费，对此，运营商方面表示，如果暂时不使用流量，市民可以主动关闭流量开关，以免造成话费损失。</p>
<p align="center"><img alt="中移动4G：几分钟1GB流量不易而飞" src="/img/20140212/cb392708737d406ab9266bb435f0222d.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292435.htm" target="_blank" title="三星走了 WCG死了"><span class="title1">三星走了 WCG死了</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小路</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:58:08</span></td>
  </tr>
  <tr>
    <td align="left"><p>上周来自官方的消息显示，<a class="f14_link" href="http://news.mydrivers.com/1/291/291811.htm" target="_blank">2014不会再有WCG</a>。正当许多玩家为这一电子竞技界曾经的标志性赛事的谢幕而感到惋惜和怀念之时。国内媒体<a class="f14_link" href="http://games.sina.com.cn/j/n/2014-02-11/1030762539.shtml?qq-pf-to=pcqq.c2c" target="_blank">新浪游戏</a>近来从另一个角度对WCG的死因进行了分析，认为已经完成其历史使命的WCG能像这样完整的结束，是最好的结局。</p>
<p>很明显的事实在于，<strong>近几年来玩家们关注WCG更多的是一种情结，而不是对特定项目的追捧和热爱。</strong>在单独的竞技项目方面，《英雄联盟》的Season系列锦标赛、《星际争霸2》的WCS和SPL、《DOTA2》的国际邀请赛，都远比综合的WCG要更加光彩夺目。这是电子竞技发展的必然趋势，随着受众人群的不断增大，单一游戏项目就会具备足够的影响力，不再需要抱团获取玩家和选手们的支持。</p>
<p><strong>WCG的没落：从传统体育来看</strong></p>
<p>拿奥运会来类比WCG最合适不过，WCG创办之初就有意效仿奥运会，做成一个电子竞技的综合盛会，回顾过去WCG也确实做到了，在ESWC、CPL等其它赛事弄得虎头蛇尾的时候，WCG逐渐变成世界范围内唯一一个具有足够影响力的电竞盛会，就如同奥运会在体育领域那样。</p>
<p>但是，就如同足球的最高荣耀是世界杯、网球的最高荣耀是四大满贯一样，在电子竞技领域，<strong>各个游戏项目的最高荣耀逐渐不再属于WCG</strong>。曾经我们为Sky的WCG冠军十分疯狂，但是后来大家发现<strong>最高水平的对抗也已经不属于WCG</strong>。</p>
<p>对于游戏厂商们来说，自家举办的赛事如火如荼，WCG的份量自然就越来越小。</p>
<p><strong>WCG的没落：三星的离去</strong></p>
<p>三星是WCG赛事的发起者和主导者，电子竞技对韩国意义重大，三星当年发起WCG是韩国电竞产业升起的一个缩影。<strong>WCG被单一公司主导也成为日后终结的导火索，奥运会如果被单一国家或公司主导，很明显也走不到今天。</strong></p>
<p>三星内部的变动影响到的是WCG的生死存亡。据资深电竞人士BBKinG透露，WCG最初的支持者是三星的显示器部门，约2010年前后变为三星的手机部门。在三星迅速发展的那几年，其显示器部门是重要的收入增长点，那时的WCG也受益颇大，那是WCG的黄金年代。后来WCG的支持部门变成手机部门，显然电子竞技对于手机来说无明显的契合度和推动作用，结束就在情理之中了。</p>
<p>BBKinG总结称说：&ldquo;<strong>WCG的兴起是因为三星的世界性扩张，不计回报的投入，这是赞助商为主导的赛事能在一个新兴行业最初阶段可以崛起的原因。而WCG作为企业的一个市场行为，当它的发展方向与三星公司战略不再一致时，分手只是一个时间问题</strong>，我个人觉得WCG已经做的够多的了，也正因为WCG有三星基因，三星不能不管，所以才能相对比较体面的活到寿终正寝。&rdquo;</p>
<p>简而言之，曾经的三星需要WCG完成其市场推广的目标，而现在重心转向手机的三星则失去了电竞的关注。对于电子竞技来说，何时才能自己完全主导命运呢？</p>
<p align="center"><img alt="三星走了 WCG死了" src="/img/20140212/898f1967da9f42409594ebc68b43da56.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292434.htm" target="_blank" title="IUNI OS发布时间确定：号称生来纯净"><span class="title1">IUNI OS发布时间确定：号称生来纯净</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:57:36</span></td>
  </tr>
  <tr>
    <td align="left"><p>今日，IUNI OS官方发出邀请函，<strong>将于2月17日在北京举行IUNI OS公测发布会。届时，IUNI OS将面向所有公布开启公测</strong>。</p>
<p>从之前曝光的消息来看，<strong>IUNI OS系统拥有&ldquo;文艺、轻快、悦动、纯净&rdquo;4大理念，从以往曝光的界面截图来看，采用扁平化设计路线，主张在功能上&ldquo;做减法&rdquo;，并宣布拒绝一切第三方预装应用，号称&ldquo;生来纯净</strong>&rdquo;。</p>
<p>此外，IUNI首款手机将将采用骁龙800处理器，综合性能会高于小米3，并预装深度定制的IUNI OS。</p>
<p align="center"><a href="/img/20140212/6ffa0aaeb5794f04bac7fb3c3f33bef5.jpg" target="_blank"><img alt="IUNI OS发布时间确定：号称生来纯净" src="/img/20140212/s_6ffa0aaeb5794f04bac7fb3c3f33bef5.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140123/9c2a1103a3a24b2381bb96dea1542dcb.jpg" target="_blank"><img alt="哎哟不错：IUNI OS再曝光 号称零预装APP" src="/img/20140123/s_9c2a1103a3a24b2381bb96dea1542dcb.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" twffan="done" /></a></p>
<p align="center"><a href="/img/20140123/fee3c9404dca420a9e6a6a1589e6a6d7.jpg" target="_blank"><img alt="哎哟不错：IUNI OS再曝光 号称零预装APP" src="/img/20140123/s_fee3c9404dca420a9e6a6a1589e6a6d7.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" twffan="done" /></a></p>
<p align="center"><a href="/img/20140123/f934bf458eb7441e986ad2c491d6c7b0.jpg" target="_blank"><img alt="哎哟不错：IUNI OS再曝光 号称零预装APP" src="/img/20140123/s_f934bf458eb7441e986ad2c491d6c7b0.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" twffan="done" /></a></p>
<p align="center"><a href="/img/20140123/70e76033596e4ea596dc1b13eed55c11.jpg" target="_blank"><img alt="哎哟不错：IUNI OS再曝光 号称零预装APP" src="/img/20140123/s_70e76033596e4ea596dc1b13eed55c11.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" twffan="done" /></a></p>
<p align="center"><a href="/img/20140123/0d046a0dd757469bb7db0cc1a5babdac.jpg" target="_blank"><img alt="哎哟不错：IUNI OS再曝光 号称零预装APP" src="/img/20140123/s_0d046a0dd757469bb7db0cc1a5babdac.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" twffan="done" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292433.htm" target="_blank" title="金山毒霸：“东莞”相关病毒已从电脑蔓延至手机"><span class="title1">金山毒霸：“东莞”相关病毒已从电脑蔓延至手机</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">刘艺</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:56:48</span></td>
  </tr>
  <tr>
    <td align="left"><p>随着央视报道东莞色情服务的继续深入，&ldquo;东莞&rdquo;两字的搜索指数直线上升，金山毒霸安全中心昨天监测到，电脑病毒伪装成与东莞色情服务有关的视频文件在传播，每天感染上万台电脑。<strong>几乎同时，与东莞相关的安卓手机病毒同样快速增长，截至目前已经感染了近10万部手机。</strong></p>
<p align="center"><a href="/img/20140212/ce95ff45a84542bf8aac19c4137f5122.jpg" target="_blank"><img alt="金山毒霸：“东莞”相关病毒已从电脑蔓延至手机" src="/img/20140212/s_ce95ff45a84542bf8aac19c4137f5122.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
图1 金山手机毒霸拦截与东莞事件有关的手机病毒</p>
<p>金山毒霸安全专家指出，安卓病毒文件名包括&ldquo;东莞不夜城&rdquo;、&ldquo;东莞不夜城的秘密&rdquo;、&ldquo;东莞大智慧&rdquo;、&ldquo;东莞情事&rdquo;、&ldquo;东莞红灯区&rdquo;、&ldquo;乐行东莞&rdquo;、&ldquo;我在东莞的艳事&rdquo;、甚至&ldquo;东莞车辆违章查询&rdquo;等等。</p>
<p><strong>经分析，这些安卓软件均为网上拼凑的关于东莞的成人小说或新闻事件，再植入手机病毒。</strong>这些恶意程序会窃取用户隐私，后台订购服务扣费、消耗手机流量，给中毒用户造成手机资费损失。</p>
<p>目前，这些恶意安卓软件正在通过手机论坛、手机QQ等多种方式在互联网上传播。据金山手机毒霸监测统计，已发现&ldquo;东莞&rdquo;相关安卓手机病毒十多个，感染了近10万部手机。</p>
<p align="center"><img alt="金山毒霸：“东莞”相关病毒已从电脑蔓延至手机" src="/img/20140212/3b71332317eb42dd9cd22f8059e75461.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /><br />
图2 网友分享到论坛的病毒传播截图</p>
<p><strong>&ldquo;东莞&rdquo;相关的病毒程序最早出现在电脑上。</strong>2月10日，金山毒霸安全中心就已监测到大量与&ldquo;东莞&rdquo;色情服务沾边的病毒文件在多个论坛、贴吧、QQ群中传播，这些病毒多伪装成&ldquo;央视记者暗访东莞实拍曝光.avi&rdquo;、&ldquo;东莞桑拿视频&rdquo;等文件，诱导网民中毒。据统计，该类病毒平均每天感染超过1万台电脑。</p>
<p align="center">我觉得<a href="/img/20140212/c0b2d692926941cc925ef9f77a7e7ac4.jpg" target="_blank"><img alt="金山毒霸：“东莞”相关病毒已从电脑蔓延至手机" src="/img/20140212/s_c0b2d692926941cc925ef9f77a7e7ac4.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
图3 金山毒霸查杀出来的&ldquo;东莞&rdquo;相关电脑病毒</p>
<p>此外，这几天不少市民还收到了类似的欺诈短信，&ldquo;爸，我在东莞玩被抓了，速汇款5000元到x警官工行卡xxxxx，别打电话，出来再说，快！&rdquo;</p>
<p>金山毒霸安全专家建议网民，收敛一些好奇心，不要轻易在电脑和手机上安装与&ldquo;东莞色情服务&rdquo;有关的任何可疑文件，也不要轻易相信与之相关的任何短信，以免上当受骗。</p>
<p><strong>相关阅读</strong>：《<a class="f14_link" href="http://news.mydrivers.com/1/292/292349.htm" target="_blank">东莞色情短信泛滥 广东成垃圾短信重灾区》</a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292431.htm" target="_blank" title="国内独享：三星5.8英寸挡脸神器升级"><span class="title1">国内独享：三星5.8英寸挡脸神器升级</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:49:59</span></td>
  </tr>
  <tr>
    <td align="left"><p>想当初Galaxy Note刚出来的时候可以说是全球最大的Android手机了，但短短几年以后Galaxy Note一些已经被甩到了身后。尽管从综合配置上来说目前比Galaxy Note 3强的没几部，但仅论屏幕比Note 3大的话那可就多了去了，三星自家就有两款，分别是Galaxy Mega 5.8和Galaxy Mega 6.3。</p>
<p>日前，<strong>三星在国内正式发布了GT-I9152P，其代号为Galaxy Mega Plus</strong>，从代号上来看我们就知道它是Galaxy Mega 5.8升级版，其依然配备了5.8英寸960&times;540分辨率TFT材质触控屏，内置1.5GB内存和8GB机身存储空间，提供一颗190万像素前置摄像头和一颗800万像素后置摄像头，支持双卡双待，电池容量为2600mAh。</p>
<p>该机和原版Galaxy Mega 5.8最大的不同就是<strong>搭载了一颗四核1.2GHz处理器，具体型号不详，而原版搭载的则是来自高通的双核1.4GHz处理器</strong>，除此之外就再无任何变化。</p>
<p align="center"><a href="/img/20140212/1e3a48f80e4940eb8cbd33b07120b832.jpg" target="_blank"><img alt="国内独享：三星5.8英寸挡脸神器升级" src="/img/20140212/s_1e3a48f80e4940eb8cbd33b07120b832.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/4c5da19f5e0d4ca1acc00a2d3fdc0ab6.jpg" target="_blank"><img alt="国内独享：三星5.8英寸挡脸神器升级" src="/img/20140212/s_4c5da19f5e0d4ca1acc00a2d3fdc0ab6.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/327e49567f424077929791c3df8bc5f9.jpg" target="_blank"><img alt="国内独享：三星5.8英寸挡脸神器升级" src="/img/20140212/s_327e49567f424077929791c3df8bc5f9.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/3728567da68d4c9293ccce99005d9e76.jpg" target="_blank"><img alt="国内独享：三星5.8英寸挡脸神器升级" src="/img/20140212/s_3728567da68d4c9293ccce99005d9e76.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292429.htm" target="_blank" title="柔性屏版Galaxy Note 4曝光：三面都是屏"><span class="title1">柔性屏版Galaxy Note 4曝光：三面都是屏</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:46:46</span></td>
  </tr>
  <tr>
    <td align="left"><p>自从三星表示要推出全新柔性屏幕手机，不少人就期待它会在S5上开启量产。但目前，新一代的旗舰手机Galaxy S5的配置基本已经十分清楚，非常遗憾的是该机并未配备上柔性屏幕。</p>
<p>不过，据韩媒最新消息，下半年的旗舰机型Galaxy Note 4将弥补缺憾，推出柔性屏版本。</p>
<p>TheKoreaHerald报道称，<strong>三星将会在今年年底推出配备柔性OLED屏版的Galaxy Note 4</strong>。不过由于柔性屏幕产能不足，因此Note 4主要出货的还是非柔性版本，而柔性屏版Note 4是作为其变种出现，首批出货仅有百万部。</p>
<p>目前尚不清楚Note 4的外观设计是否还是经典的三星风格，不过按照去年的原型机Youm来看，<strong>三星柔性屏版本的Note 4很可能、采用&ldquo;三面屏幕&rdquo;&mdash;&mdash;呈弧形，除了正面大屏幕，两个侧边还有小屏幕，供用户查看消息提醒使用</strong>。彭博社也证实称，Note 4确有这个&ldquo;三面屏&rdquo;版本。</p>
<p>按照三星规划图，<strong>三星柔性屏将会在2015年大规模量产，而且届时还会有一款配备可折叠屏幕的设备（Note 5？）问世</strong>。</p>
<p align="center"><img alt="柔性屏版Galaxy Note 4曝光：三面都是屏" src="/img/20140212/80d180caf15e41e68829ffbfc928dd98.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
原型机Youm</p>
<p align="center"><a href="/img/20140212/741ba103d6f84ecc9f7191f17c0cd56e.jpg" target="_blank"><img alt="柔性屏版Galaxy Note 4曝光：三面都是屏" src="/img/20140212/s_741ba103d6f84ecc9f7191f17c0cd56e.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
规划图</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292427.htm" target="_blank" title="宅男创作色情小说 涉嫌传播淫秽物品被刑拘"><span class="title1">宅男创作色情小说 涉嫌传播淫秽物品被刑拘</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">刘艺</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:33:56</span></td>
  </tr>
  <tr>
    <td align="left"><p>随着网络的发展，网络小说越来越流行，网络已成为一个文学创作平台，也成为许多网友的精神寄托之地。越是这样，网络越应该带给网友积极向上的东西，而不是暴力、色情等不健康内容。今天，<strong>主持人给大家说一件这样的事：&ldquo;宅男&rdquo;宅在家中写网络小说，却不知自己已触犯法律。</strong></p>
<p>禹州市的宋某三十出头，已结婚生子。平常，他并不喜欢出去玩，而是总爱待在家中，玩玩电脑，看看小说，是一个&ldquo;宅男&rdquo;。</p>
<p>有一次，他和朋友相聚，朋友之间有人讲起了&ldquo;黄段子&rdquo;，宋某听完之后觉得十分有&ldquo;内涵&rdquo;。追问之下，朋友就给他介绍了一个论坛，还说：&ldquo;保证让你看得血脉偾张。&rdquo;</p>
<p><strong>宋某回到家后，打开电脑就进入了朋友说的论坛，里面充斥着色情电影、图片和小说。从此之后，浏览这个论坛就成了宋某的一大爱好，欲罢不能。</strong></p>
<p>2013年5月，宋某在论坛上注册了一个颇有几分&ldquo;文艺范儿&rdquo;的昵称&mdash;&mdash;&ldquo;醉枕千秋梦&rdquo;，获取邀请码后激活了账号。看完&ldquo;新手上路&rdquo;的提示后，他接连发了三个新手帖子，获取了一些免费积分。他发现，有了积分，就可以浏览更多的内容。</p>
<p>一周后，经常沉浸在&ldquo;品味&rdquo;色情小说之中的宋某，看到论坛上发布的&ldquo;文心雕龙第六届征文大赛火热进行中&rdquo;的帖子后，突然萌生了原创一部色情小说的想法。</p>
<p>2013年5月18日，<strong>宋某以网络上流行的仙侠小说为架构，发表了两个章节，共计7300余字，吸引了许多注册会员的浏览和跟帖，宋某则获得相应积分</strong>。随后，宋某接连发表了16章，共计7万余字。</p>
<p>经公安机关鉴定，这部小说每一章节都有长时间描述男女裸体、猥亵、色情和性行为的情节。因涉嫌传播淫秽物品罪，宋某被依法刑事拘留。</p>
<p><strong>此案在庭审时，宋某对勘验结果和鉴定意见没有异议，对其传播淫秽物品的犯罪事实供认不讳，当庭表示认罪。</strong></p>
<p>&ldquo;纯属无聊。主要是发表帖子越多，积分就越高，等级和权限就会提升，可以浏览更多的板块，于是就有了写一部小说的想法。&rdquo;庭审中，宋某说出自己撰写色情小说的动机。</p>
<p>同时，宋某说自己的写作过程也不是很难：&ldquo;我平时经常在网上看这样的小说，就模仿着网上的小说，再添加一些内容，就成了自己的作品。&rdquo;宋某撰写一章内容短则三天，长则一周，平均每章3500字左右。其中一次，宋某在3天内连载3章，字数长达1.4万余字。</p>
<p>从2013年5月18日至7月13日，近两个月内，这部色情小说宋某共撰写了16章，计7万余字，获取积分106个，金币1902枚，金榜18个，感谢度786，共计有50618人次浏览。</p>
<p>在法庭上，宋某供述，他写这部小说并没有和论坛的管理员取得联系，只是在有网友跟帖时，他进行回帖。<strong>网友的&ldquo;鼓励&rdquo;对宋某撰写淫秽小说无疑起到了推波助澜的作用。</strong>证据显示，在宋某发表的每一章小说后，都有许多网友跟帖，还有网友为他提供构思。不过宋某称，他自始至终没有收到任何报酬。</p>
<p><strong>宋某还表示，他不懂法，也不知道这样写写小说就触犯了法律。</strong>他当庭保证不再犯，并请求从轻处罚。</p>
<p>近日，禹州市法院一审审结了这起案件，以传播淫秽物品罪判处宋某有期徒刑一年，缓刑一年。</p>
<p>主持人：我们看到在这个案例中，当事人宋某的行为构成了传播淫秽物品罪。从道德层面上讲，我们当然不认同任何传播色情内容的行为，提倡严厉打击，但是从法律层面上讲，是不是只要在网络上发表色情小说就构成传播淫秽物品罪？什么情况才构成传播淫秽物品罪？</p>
<p>郑州市管城区法院刑庭副庭长邢艺伟：<strong>传播淫秽物品罪是指不以牟利为目的，传播淫秽的书刊、影片、音像、图片或者其他淫秽物品，情节严重的行为。</strong></p>
<p>该罪侵犯的客体是社会主义道德风尚，客观方面表现为传播淫秽的书刊、影片、音像、图片或者其他淫秽物品的行为。&ldquo;传播&rdquo;主要指出借、播放、展示、赠送、散发、交换、讲解等行为，但这种行为必须是在公众中或者公共场所实施的。传播</p>
<p>可以是公开的，也可以是秘密的。</p>
<p>当然，并不是只要在网络上发表色情小说就构成传播淫秽物品罪，按照《最高人民法院、最高人民检察院关于办理利用互联网、移动通讯终端、声讯台制作、复制、出版、贩卖、传播淫秽电子信息刑事案件具体应用法律若干问题的解释》规定，这些小说达到一定点击量，发表人就构成了犯罪。如果传播的范围较小，观看的人数较少，则不构成犯罪，可视情节轻重，给予批评教育或者治安处罚。</p>
<p>同时应当注意的是，<strong>不以牟利为目的的传播行为构成的是传播淫秽物品罪，而一旦传播者是以牟利为目的，构成的就是传播淫秽物品牟利罪，处罚会更重。</strong></p>
<p>欢迎广大读者和主持人一起来说说发生在我们身边的事儿，不管大事还是小事，只要与法有关就可。让我们从身边的事儿说起，共同来学法、知法、懂法、用法。</p>
<p align="center"><img alt="宅男创作色情小说 涉嫌传播淫秽物品被刑拘" src="/img/20140212/6be50578d56b4dd4adb101a3b77d38bd.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292426.htm" target="_blank" title="领导人新闻挨着裸女图 央媒炮轰商业网站色情营销"><span class="title1">领导人新闻挨着裸女图 央媒炮轰商业网站色情营销</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:30:19</span></td>
  </tr>
  <tr>
    <td align="left"><p>今天是国家互联网信息办公室开启打击互联网传播淫秽色情及低俗信息专项行动后的第15天，<strong>记者浏览不少影响力名列前茅的大型商业网站后惊讶地发现，曾经在行动开启当天迅速撤下的包含&ldquo;偷拍&rdquo;&ldquo;走光&rdquo;等不雅词汇的低俗图文，已全面再度&ldquo;重出江湖&rdquo;。</strong></p>
<p>记者调查发现，这些低俗图文通常与各类广告紧密相连，往往带有商业目的，是典型的&ldquo;色情营销&rdquo;。</p>
<p><strong>网络治理剑指&ldquo;色情营销&rdquo;，低俗图文暂别后&ldquo;重出江湖&rdquo;</strong></p>
<p>&ldquo;不雅照&rdquo;&ldquo;被偷拍&rdquo;&hellip;&hellip;这些不雅词汇和图文，在国内影响力名列前茅的一些大型商业门户网站中，眼下比比皆是。</p>
<p>从1月28日至2014年全国两会结束，国家互联网信息办公室开展打击互联网传播淫秽色情及低俗信息专项行动，重点打击&ldquo;色情营销&rdquo;以及低俗炒作行为，<strong>要求新闻网站和主要商业网站不得推荐所谓&ldquo;ＡＶ女星&rdquo;&ldquo;色诱&rdquo;&ldquo;偷拍&rdquo;&ldquo;走光&rdquo;等低俗图文信息</strong>，不得故意制造格调低俗、哗众取宠的标题博取眼球。</p>
<p>记者调查发现，<strong>专项行动开启当日，不雅图文和视频迅速从网站上撤离，页面上出现了难得一见的&ldquo;绿色&rdquo;和&ldquo;纯洁</strong>&rdquo;。&ldquo;这恰恰说明了这些网站对自身的违规行为心知肚明，也觉得不太光彩。&rdquo;经常在这些大型商业网站浏览新闻的任先生表示。</p>
<p>然而，记者在春节假期期间就发现，虽然专项行动尚在进行中，但在&ldquo;避风头&rdquo;似的短暂告别之后，这些不雅图文和视频再度重返一些网站的&ldquo;舞台&rdquo;，而且势头丝毫不逊色于专项行动开启之前。</p>
<p>记者11日上午浏览一些商业网站时发现，<strong>无论是在首页的醒目位置，还是在重要时政新闻的边栏，一些低俗信息很容易跃入读者的视野中。由于媒体暗访东莞色情业形成了舆论热点，不少网站不仅报道东莞事件本身，还爆发式地展示了与色情服务相关的诸多图文和视频，其中不乏带有刺激性的画面和场景。</strong></p>
<p><strong>以不雅图片吸引眼球，&ldquo;色情营销&rdquo;危害青少年</strong></p>
<p>商业网站之所以对这些低俗信息乐此不疲，主要是因为网站要利用其进行&ldquo;色情营销&rdquo;。<strong>这些不雅图文通常附着在重要新闻事件页面的边栏，其周边位置链接着各类广告。网站利用半裸、透视的美女图片和视频吸引网民眼球，从而进行产品的商业推广。</strong></p>
<p>记者点开一家大型网站的头条新闻发现，其边栏中并列着两张不雅图片。这两张图片的上方是一个服装的广告，下方是一个网络游戏的广告。该网站利用刺激性图片进行&ldquo;色情营销&rdquo;的意图十分明显。</p>
<p>&ldquo;对于什么类型的图文和视频旁边链接哪类广告，针对哪些目标群体，网站的编辑都是精心策划的。&rdquo;一位大型商业网站的工作人员告诉记者。<strong>记者发现，一些衣衫不整的美女图片旁边的产品以汽车、网络游戏领域的广告居多。很明显，这些图片是为了吸引更多男性或青少年的目光。</strong></p>
<p>&nbsp;成年人看到这些不雅图片可能会&ldquo;一笑而过&rdquo;，大多数正处于青春期的青少年看到这些不良信息难免受到影响。在不少网站中，不雅图片或视频大多与网络游戏或手机游戏的广告紧密相连，这让不少学生家长尤为气愤。</p>
<p><strong>&ldquo;网络游戏受众大多是青少年，把这些图片与网络游戏放在一起分明是为了吸引孩子们点击、浏览图片，进而推广游戏。&rdquo;</strong>一位学生家长对记者说，&ldquo;目前正处于学生的假期，孩子们上网的机会较多。青少年往往辨别力和自制力较差，这些不良信息容易导致其去寻找更加露骨的色情信息的行为，甚至走上违法犯罪的道路。&rdquo;</p>
<p><strong>低俗图文多与重要新闻&ldquo;混搭&rdquo;，重塑网络环境需长效机制</strong></p>
<p>无论是为了&ldquo;色情营销&rdquo;，还是为了增加网站人气，<strong>目前一些原本难登&ldquo;大雅之堂&rdquo;的图片和视频常与国家领导人活动，时政、军事等新闻相连在一起，呈现在同一个页面空间上，总让人觉得有些&ldquo;时空错乱&rdquo;。</strong></p>
<p>记者调查发现，在一家商业网站首页排名前15位的新闻中，页面边栏无一例外地都链接着不雅图片和视频；在另一家商业网站，<strong>一篇外媒评价中国全面深化改革和发展前景的头条新闻旁边则链接着不雅图片；还有一家商业网站在一篇国家领导人谈改革的文章结尾处右边，并列着两张半裸和透视的美女图片。</strong></p>
<p>专家表示，低俗信息的大量存在严重恶化了互联网的环境，在一些知名网站新闻报道中频繁出现更对新闻传播的严肃性产生影响。这些不雅图文死灰复燃也体现出此次专项行动的威慑力还没有完全发挥。打击治理若想取得更大成效，还应该更加具有针对性和系统性。</p>
<p>&ldquo;首先要在立法上更加完善，对于禁止的图文、视频要明确标准和边界，最大限度遏制擦边球行为，对于网站或运营商的违规行为的惩罚措施同样应给予明确。&rdquo;北京两高律师事务所合伙人律师董正伟建议，&ldquo;要建立长期、系统、稳定的执法程序，让监督和审查形成长效机制，而不是依赖运动式的突击行动。此外，还可以建立一个投诉举报平台，发动全社会的力量进行舆论监督。&rdquo;</p>
<p align="center"><a href="/img/20140212/ec736ee07b0a42a6bde9c23f4e7c33f4.jpg" target="_blank"><img alt="领导人新闻挨着裸女图 央媒炮轰商业网站色情营销" src="/img/20140212/s_ec736ee07b0a42a6bde9c23f4e7c33f4.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p>&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292425.htm" target="_blank" title="关于诺基亚Android神器 你要知道的5件事"><span class="title1">关于诺基亚Android神器 你要知道的5件事</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:29:19</span></td>
  </tr>
  <tr>
    <td align="left"><p>诺基亚疯了吗？或许没有。就在他们试图去提振Windows Phone的销量、同时完成和微软的收购案之时，有传言称他们即将在本月的MWC上推出一款Android手机。<strong>这款手机并不是像Lumia 1020那样的旗舰机型，而是一款面向新兴市场，企图创造超高销量的低端产品。</strong></p>
<p>它最初的代号是Normandy，而其正式名称将是Nokia X。对于这款手机，我们需要知道的信息都有哪些？科技网站Laptopmag日前就对此进行了汇总：</p>
<p><strong>定制版Android+诺基亚服务</strong></p>
<p>和亚马逊类似，看样子诺基亚也将会去以自己的界面对Android进行定制。这也就是说Nokia X将不会拥有Google Play商店、Gmail、YouTube、或者是任何其他的谷歌服务。ZDNet的Mary Jo Foley指出，诺基亚可能会提供他们自己的应用商店（和Kindle类似），并预装自己的服务，比如Here地图。此外，在网络上曝光的系统截图也证实诺基亚的这款手机将支持诸如BBM、Vine、Viber和微信这样的第三方应用。虽然微软对此可能并不高兴，但实际上Skype和Xbox Music这样的微软服务也会进入Nokia X。</p>
<p><strong>Windows Phone的外观和感觉</strong></p>
<p>只消看一眼evleaks曝光的系统截图，你就会恍惚觉得自己看到的是一部Windows Phone设备，其中五颜六色的磁贴状图标和Lumia手机所具备的非常相似。不过这些磁贴应该并非是动态的，所以只会在后台提供更新（和其他Android手机一样，谍照显示Nokia X将拥有自己的下拉通知菜单）。不过你至少可以随意地移动这些磁贴，并调整他们的大小。</p>
<p><strong>低配置</strong></p>
<p>当你试图去达到一个很低的价格点（可能是100美元），配置方面就难以保持在较高的水平了。报道称这款Nokia X将会配备骁龙400处理器，4GB存储空间，512MB内存，300万像素摄像头，并支持microSD卡拓展。</p>
<p><strong>为什么微软不出手干涉</strong></p>
<p>考虑到Android和Windows Phone是竞争关系，微软首先的反应可能会是将Nokia X扼杀于襁褓。但新任总裁Satya Nadella有很大的理由去支持这个项目。首先，对于那些质疑之声，微软可以说这个项目在他们收购诺基亚之前就已经处于研发阶段了。诺基亚和微软还可以说，该项目和Asha一样，瞄准的都是低端市场，前者不过只是一种替代而已。</p>
<p>此外，Nokia X可能会把数以百万计的消费者带到诺基亚和微软的服务当中。这些服务当中有的是需要订阅的，因此会创造收益；而其他的服务则可维持用户的跨平台忠诚度。</p>
<p><strong>成功还是失败？</strong></p>
<p>诺基亚在这款手机身上所做的是把Android当作是通往一款更高端Windows Phone设备的垫脚石。比起把4100万像素的摄像头塞到Lumia设备当中，这个消息会更让微软高兴。与此同时，诺基亚还能立即获得Windows Phone所缺乏的大量应用程序。那么大众会喜欢这样一部手机吗？有可能，只要诺基亚填补好谷歌服务留下的空缺。</p>
<p>如果Nokia X成功，就没有什么能够阻止诺基亚推出一款更为高端的Android手机了&mdash;&mdash;也许除了微软的自尊之外。</p>
<p align="center"><img alt="关于诺基亚Android神器 你要知道的5件事" src="/img/20140212/9219a1ad38fe4d5188e1d01f4a6765b1.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292424.htm" target="_blank" title="网友剪辑央视直播爆笑失误合集 网上热传"><span class="title1">网友剪辑央视直播爆笑失误合集 网上热传</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">刘艺</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:25:24</span></td>
  </tr>
  <tr>
    <td align="left"><p>近日，一段网友剪辑的&ldquo;央视各种直播失误合集&rdquo;视频在网络上热传，这段时长为7分48秒的视频将最近几年央视各频道节目直播时主持人的口误、导播切换镜头的失误以及嘉宾忘词等&ldquo;事故&rdquo;片段集合在一起，其中也不乏李瑞英、赵普、海霞等名嘴&ldquo;老马失蹄&rdquo;的镜头。</p>
<p>网友纷纷评论&ldquo;笑的我下巴都掉了&rdquo;，&ldquo;笑的根本停不下来&rdquo;。也有网友很理解电视直播的不易之处&ldquo;就算是精密机器也会有出错的时候，身为人类，出错更是生活中会遇上的一种情况，我们都可以理解、体谅的&rdquo;，&ldquo;习惯了央视主播平时严肃正经的样子，突然一个小失误真的好好笑&rdquo;。</p>
<p><strong>视频地址</strong>：<a href="http://video.sina.com.cn/p/news/gaoxiao/v/2014-02-10/193063487195.html">http://video.sina.com.cn/p/news/gaoxiao/v/2014-02-10/193063487195.html</a></p>
<p align="center"><img alt="网友剪辑央视直播爆笑失误合集 网上热传" src="/img/20140212/b6e2471f94214d0b96cbbe26deed387d.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /><br />
赵普播报亚洲杯新闻时忘词，脱口而出&ldquo;我看看啊&rdquo;</p>
<p align="center"><img alt="网友剪辑央视直播爆笑失误合集 网上热传" src="/img/20140212/08d74c5449f744b0bb85ebffa6130659.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /><br />
李瑞英新闻联播中口误结巴（视频截图）</p>
<p align="center"><img alt="网友剪辑央视直播爆笑失误合集 网上热传" src="/img/20140212/233c787e4e08473abb660b7d4a7f32f9.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /><br />
海霞误将&ldquo;新闻20分&rdquo;说成&ldquo;新闻20婚&rdquo;（视频截图）</p>
<p align="center"><img alt="网友剪辑央视直播爆笑失误合集 网上热传" src="/img/20140212/80e5e46a20c349d882674b3cf62c3616.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /><br />
白岩松：用五个字可以形容&ldquo;都已经准备好了&rdquo;（视频截图）</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292423.htm" target="_blank" title="Windows 9或将决定微软新CEO纳德拉的命运 "><span class="title1">Windows 9或将决定微软新CEO纳德拉的命运 </span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:24:17</span></td>
  </tr>
  <tr>
    <td align="left"><p>萨提亚&middot;纳德拉(Satya Nadella)刚刚上任微软CEO，但<strong>如果Windows 9不能纠正微软的航向，那么纳德拉或被抛弃的命运将不足为奇</strong>。</p>
<p>毫无疑问，是Windows 8的灾害导致史蒂夫&middot;鲍尔默辞职。用两个针对不同受众的界面来创建一款单一的操作系统，是不会令任何人满意的，消费者和企业都会避而远之。Windows 8本应该能在移动领域拯救微软，但Windows Phone仍深陷困境，在美国市场仅拥有4%的市场份额，在中国市场的市场份额甚至不足1%，这两大市场则是全球两个最大的智能手机市场。另外，Windows平板电脑的销量也是如此糟糕，致使微软被迫对未售出Surface RT库存减记9亿美元。</p>
<p>台式机和笔记本电脑的销量一直在萎缩，平板电脑和智能手机的销量则一路顺风，Windows 8反而推动了传统PC销量的下滑，要在移动领域帮助微软，Windows 8已无能为力。这就是市场对Windows 9寄予厚望的原因所在，Windows 9预计将于2015年年中推出，不过也有传言称该产品要在今年十月份才能推出。</p>
<p><strong>要抓住台式机和笔记本电脑的未来，并大举进军移动领域，Windows 9可能是微软的最后机会</strong>。如果Windows 9表现得不好，那么微软重回其辉煌岁月的几率将非常小。在2000年推出Windows ME后，微软于2001年推出了Windows XP，在2007年推出惨败的Windows Vista后，该公司又在2009年推出Windows 7。但成功来自于远见卓识、努力工作以及对过去经验教训的总结，尽管有先例，但一款好的操作系统的推出并不意味着前一款操作系统很糟糕。</p>
<p>尽管Windows 9的开发工作正在顺利进行中，是好是坏，纳德拉的命运都将由此判断，如果是一款伟大的操作系统，那么将巩固纳德拉未来几年在微软的地位，但如果不是这样，鉴于微软处于水深火热之中，纳德拉的命运也将未卜。&nbsp;</p>
<p align="center"><a href="/img/20140212/ee387968595245a3a93185da6c21e424.jpg" target="_blank"><img alt="Windows 9或将决定微软新CEO纳德拉的命运 " src="/img/20140212/s_ee387968595245a3a93185da6c21e424.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292422.htm" target="_blank" title="华为：令人眼前一亮的产品马上有"><span class="title1">华为：令人眼前一亮的产品马上有</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:17:54</span></td>
  </tr>
  <tr>
    <td align="left"><p>春节期间，华为终端的领导们集体在微博上&ldquo;销声匿迹&rdquo;，对此有用户在微博上询问华为副总裁余承东是否是因为华为高管集体退出微博营销。</p>
<p>另外没想到的是余承东居然转发了这条微博，<strong>并表示华为高管不会退出微博</strong>，只是在春节假期中抽时间多陪了陪家人而已，老余这张亲情牌打的确实不错。</p>
<p>此外余承东还表示在接下来的<strong>巴展（MWC2014）上华为将会给大家带来令人眼睛一亮的产品</strong>，但具体是什么产品他并没有透露。</p>
<p>我们猜测其口中的所谓令人眼前一亮的产品应该就是此前曝光的P7，<a class="f14_link" href="http://news.mydrivers.com/1/291/291657.htm" target="_blank">其代号为索菲亚</a>（sophia），将<a class="f14_link" href="http://news.mydrivers.com/1/291/291891.htm" target="_blank">采用5英寸1080p显示屏</a>，<strong>搭载海思K3V2 Pro处理器，内置2GB内存和16GB机身存储空间，提供一颗800万像素前置摄像头和一颗1300万像素后置摄像头</strong>，运行Android 4.4.2操作系统，电池容量为2460mAh，支持TD-LTE网络。</p>
<p>该机在外观方面会和P6差不多，会延续金属机身设计，但由于屏幕大了不少，因此机身尺寸必然要比P6大一些，希望它真的能够让我们&ldquo;眼前一亮&rdquo;。</p>
<p align="center"><a href="/img/20140212/11f9c998b0364ed296b561898a3f7f8f.jpg" target="_blank"><img alt="华为：令人眼前一亮的产品马上有" src="/img/20140212/s_11f9c998b0364ed296b561898a3f7f8f.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><img alt="华为：令人眼前一亮的产品马上有" src="/img/20140212/76b64e8c5b9244fa9ff575ac0ec29e19.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
此前网友曝光的P7</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292421.htm" target="_blank" title="美国6000多家网站集体抗议国安局监控"><span class="title1">美国6000多家网站集体抗议国安局监控</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:14:59</span></td>
  </tr>
  <tr>
    <td align="left"><p>11日，超过6000家网站在首页挂出标语，抗议美国国家安全局(NSA)的监控行为，并对去年自杀的美国黑客奇才亚伦&middot;斯沃茨进行纪念。</p>
<p><strong>包括Reddit、Tumblr、Mozilla在内的超过6000家网站，参加了这一名为&ldquo;反击日&rdquo;的抗议活动。</strong>网站在首页顶部挂出标语，写道：&ldquo;<strong>亲爱的互联网，我们厌倦了抱怨NSA。我们希望有新的法律能够限制在线监控。</strong>&rdquo;</p>
<p>标语链接到抗议活动网站，在网站中，美国互联网用户能够通过电邮或者互联网电话联系美国国会，反对将将强化NSA监控行为合法性的外国情报监视法案修订案，而支持美国自由法案。根据网站信息，已有超过15万的电子邮件发出，超过76万个电话打出。</p>
<p>美国境外的访问者则被呼吁签署一份请愿书，反对大规模监控。已有超过10万人签署了请愿书。</p>
<p>本次抗议活动的日期选择，是为了纪念2年前对美国国会众议院&ldquo;反网络盗版法&rdquo;(SOPA)及参议院&ldquo;知识产权保护法&rdquo;(PIPA)的抗议行动，以及1年前自杀的美国黑客奇才亚伦&middot;斯沃茨。</p>
<p>2012年1月，&ldquo;维基百科&rdquo;与数千家其他网站关闭24小时，抗议SOPA和PIPA法案，促使美国国会宣布暂停两项反网上盗版的立法辩论。</p>
<p>2011年7月，斯沃茨因被指控从麻省理工学院和期刊数据库西文刊全文库中偷窃文档被逮捕，但他拒绝认罪，后于2012年1月11日自杀，年仅26岁。</p>
<p align="center"><img alt="美国6000多家网站集体抗议国安局监控" src="/img/20140212/434a16a230c14a008f162af9fd9aaaf7.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292419.htm" target="_blank" title="Nubia新UI曝光？你们都被忽悠了"><span class="title1">Nubia新UI曝光？你们都被忽悠了</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:02:50</span></td>
  </tr>
  <tr>
    <td align="left"><p>几天前Nubia通过其官方微博暗示<a class="f14_link" href="http://news.mydrivers.com/1/291/291991.htm" target="_blank">全新扁平化的Nubia UI即将到来</a>，而随后<a class="f14_link" href="http://news.mydrivers.com/1/292/292217.htm" target="_blank">网上就出现了这款UI的真容</a>，但经过查证，其实大家都被忽悠了。</p>
<p>事情的起因是微博用户@牛肉丸曝光了一张通知栏的截图，并称&ldquo;不要问我这是什么，我也不知道！某人私信给我的，让我曝光&rdquo;，<strong>随后Nubia手机品牌总监也转发了这条微博</strong>，这让很多人都认为@牛肉丸曝光的图片是Nubia的新UI。</p>
<p>但随后有网友反馈称这并不是什么Nubia的新UI，而是<a class="f14_link" href="http://www.zcool.com.cn/work/ZMjY1OTgyMA==.html" target="_blank">康佳Muse UI大赛的参赛作品</a>，<strong>作者名字叫pulrou</strong>，和Nubia没有半毛钱的关系。</p>
<p>看来这件事情从头到尾只是一个误会而已，同时也希望Nubia的新UI能早点到来。</p>
<p align="center"><img alt="Nubia新UI曝光？你们都被忽悠了" src="/img/20140212/2fcd54a9e35946abab40c63d41b7db8c.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
此前曝光的所谓Nubia新UI</p>
<p align="center"><a href="/img/20140212/f7efe9d0798f4c28897cdb63773328af.jpg" target="_blank"><img alt="Nubia新UI曝光？你们都被忽悠了" src="/img/20140212/s_f7efe9d0798f4c28897cdb63773328af.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
实际上是网友的新作品</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292418.htm" target="_blank" title="诺基亚Android手机不止一款：最快5月上市"><span class="title1">诺基亚Android手机不止一款：最快5月上市</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:02:31</span></td>
  </tr>
  <tr>
    <td align="left"><p>根据《华尔街日报》日前的报道，<a class="f14_link" href="http://news.mydrivers.com/1/292/292231.htm" target="_blank">诺基亚将会在即将到来的MWC 2014上推出旗下首款搭载Android系统的手机诺曼底</a>（Nokia X）。该机定位低端入门智能机，隶属于Asha系列。</p>
<p>而据<a class="f14_link" href="http://tech.qq.com/a/20140212/012223.htm" target="_blank">腾讯科技</a>获得的消息，<strong>除了Nokia X之外，诺基亚还将推出多款定位不同的Android手机。目前，诺基亚还在测试其它机型</strong>。</p>
<p>据悉，<strong>诺基亚Android手机虽然运行的是Android系统，但已经将谷歌服务去除，用户无法从Google Play商店下载Android应用</strong>。诺基亚还在手机中加入了自家的特色应用，比如地图服务Here、流媒体音乐服务MixRadio和一款诺基亚应用商店。</p>
<p>根据之前的报道，诺基亚Android手机最快要等到5月-6月份上市，首款Nokia X的定价预计在1500元左右。</p>
<p align="center"><a href="/img/20140212/ad5b3c868d924aa093b872ae37cb7f8c.jpg" target="_blank"><img alt="诺基亚Android手机不止一款：最快5月上市" src="/img/20140212/s_ad5b3c868d924aa093b872ae37cb7f8c.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292417.htm" target="_blank" title="谷歌、微软：NSA 你接招吧！"><span class="title1">谷歌、微软：NSA 你接招吧！</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">随心</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">11:00:50</span></td>
  </tr>
  <tr>
    <td align="left"><p>美国NSA（国家安全局）的监控丑闻曝光后，给人们的隐私带来了威胁，这也惹怒了多家互联网公司。上个月下旬，微软就通过<a class="f14_link" href="http://news.mydrivers.com/1/291/291036.htm" target="_blank">允许用户将数据存在国外</a>，打出了迎击NSA的第一枪，随后谷歌也宣布<a class="f14_link" href="http://news.mydrivers.com/1/291/291238.htm" target="_blank">未来将推出&ldquo;加密&rdquo;服务</a>，让政府无法探秘。而现在，更大的统一战线成立，NSA，接招吧！</p>
<p>据外媒报道，今天早些时候，微软、谷歌、Facebook等多家公司结成的反监控联盟。他们在<a class="f14_link" href="http://www.reformgovernmentsurveillance.com/" target="_blank">Reform Government Surveillance</a>网站放出了&ldquo;<strong>今天,让我们反击</strong>(TODAY WE FIGHT BACK)&rdquo;的标语，直指美国政府监控项目，并且深情的写了一封《致华盛顿的公开信》。</p>
<p align="center"><a href="/img/20140212/f0b50afcd5a84ed68db2634f89aa7dfb.jpg" target="_blank"><img alt="谷歌、微软跟NSA干上了" src="/img/20140212/s_f0b50afcd5a84ed68db2634f89aa7dfb.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/795fe2d9b1fc4ffd8f2b8cc5ca1e2ace.jpg" target="_blank"><img alt="谷歌、微软跟NSA干上了" src="/img/20140212/s_795fe2d9b1fc4ffd8f2b8cc5ca1e2ace.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/b75aee46e8ce4d25942133483e16ff2b.jpg" target="_blank"><img alt="谷歌、微软跟NSA干上了" src="/img/20140212/s_b75aee46e8ce4d25942133483e16ff2b.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292416.htm" target="_blank" title="ARM Cortex-A17详细解析：比肩A15、踩死A12"><span class="title1">ARM Cortex-A17详细解析：比肩A15、踩死A12</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:58:11</span></td>
  </tr>
  <tr>
    <td align="left"><p><a class="f14_link" href="http://news.mydrivers.com/1/292/292331.htm" target="_blank">ARM昨天发布了新的内核架构，但起了个不伦不类的名字&ldquo;Cortex-A17&rdquo;。</a>看上去，它比32位架构中的顶级内核Cortex-A15还要高级，但实际上只是Cortex-A12的改进版本，定位2015年的主流市场。</p>
<p>A17仍然基于32位ARMv7-A指令集，本质架构和A12一样都是<strong>双宽度、乱序发射</strong>。尽管架构细节尚未披露，但据了解，<strong>A17、A12的前端、执行引擎都是一样的，性能、能效的改进基本都来源于内存子系统的变化。</strong></p>
<p><strong>ARM宣称，在特定频率、工艺、内存条件下，A17的性能相比于A9r4可以提升大约60％，而此前A12宣称的提升幅度是40％。照此计算，A17的性能已经可以和A15处于一个档次了！</strong></p>
<p align="center"><strong><a href="/img/20140212/ad14167e9ac648c2aa088fd2268b6fab.png" target="_blank"><img alt="ARM Cortex-A17详细解析：比肩A15、踩死A12" src="/img/20140212/s_ad14167e9ac648c2aa088fd2268b6fab.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></strong></p>
<p>在特定工艺下，<strong>A17的核心面积会比A9大出约20％</strong>，同时稍稍大于A12。</p>
<p>而运行同样的负载，<strong>A17的能效可比A9高出约20％</strong>(主要是休眠速度)，但是预计A17的峰值功耗也会更高。</p>
<p>或许，<strong>这才是ARM将它的命名拔高的原因：性能接近A15，核心面积更小，功耗更低。</strong>ARM称要以此在2015年给主流用户提供2014年的高端体验。</p>
<p>但这也不由得让我们怀疑：A15你到底是来干啥的？说定位高端，结果现在被一个中端核心打得满地找牙(至少理论上)，还有什么存在价值？&nbsp;</p>
<p>此外，A12因为缺乏必要的一致性总线，不支持big.LITTLE，A17则弥补了这一缺憾，<strong>可以作为&ldquo;大核心&rdquo;，搭配&ldquo;小核心&rdquo;的Cortex-A7。</strong></p>
<p align="center"><strong><a href="/img/20140212/e7ac911c61744da197dafa68d8e9a888.png" target="_blank"><img alt="ARM Cortex-A17详细解析：比肩A15、踩死A12" src="/img/20140212/s_e7ac911c61744da197dafa68d8e9a888.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></strong></p>
<p>A17初期会采用28nm工艺，后期进化到20nm。ARM给出了一张有趣的图示，显示了历代工艺带来的芯片面积变化，以及晶体管成本&mdash;&mdash;每1美元能买到的晶体管数量。</p>
<p>根据这张图，<strong>当前的28nm是个&ldquo;甜点&rdquo;，晶体管单位成本最低</strong>(1美元2000万个)，而下一代的20nm只能做到和它持平，19nm的成本反而会略有增加。这的确是半导体行业不得不面临的难题，更先进的工艺已经很难带来足够多的好处。</p>
<p>A17定位中端，对面积、成本自然更加敏感，20nm必须完全成熟了才会采纳。&nbsp;</p>
<p align="center"><a href="/img/20140212/be5d98273b1c4dffb9a08f0f90445123.png" target="_blank"><img alt="ARM Cortex-A17详细解析：比肩A15、踩死A12" src="/img/20140212/s_be5d98273b1c4dffb9a08f0f90445123.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>ARM的另一张图则显示了不同面积芯片对应的不同领域：4-20平方毫米的小家伙用于传感器、物联网、可穿戴设备、实时嵌入式，中端移动、消费电子、可穿戴设备的需要控制在30-70平方毫米，高端高性能移动设备可以放宽到70-100平方毫米，而那些100-300平方毫米的大家伙只能用于网络基础架构、存储和服务器。</p>
<p align="center"><strong><a href="/img/20140212/1b40f1cbf1f34b99b39b7f4270ba80e3.png" target="_blank"><img alt="ARM Cortex-A17详细解析：比肩A15、踩死A12" src="/img/20140212/s_1b40f1cbf1f34b99b39b7f4270ba80e3.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></strong></p>
<p>现在最大的问题是，A12何去何从？ARM的态度是二者将同时存在一段时间，A12产品今年下半年诞生，A17则在明年接班。</p>
<p>但是厂商们可不愿意这么看：<strong>有更高级、更好看的A17，谁还愿意要落伍的A12？</strong>瑞芯微RK3288、<a class="f14_link" href="http://news.mydrivers.com/1/292/292333.htm" target="_blank">联发科MT6595</a>都已经迫不及待地采用了A17。可以预料，A12的命运只能是直接胎死腹中。</p>
<p>架构方面，虽然在热炒64位，但是ARM很冷静地指出，<strong>32/64位的转换需要3-4年才能完成，ARMv7-A因此仍有较长的寿命</strong>，今年就会有超过10亿部这种架构的智能手机出货，况且ARMv8-A架构向下兼容它。</p>
<p align="center"><a href="/img/20140212/da3305e868f045b2a856531f33d4c84d.png" target="_blank"><img alt="ARM Cortex-A17详细解析：比肩A15、踩死A12" src="/img/20140212/s_da3305e868f045b2a856531f33d4c84d.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/3e086eb7a66445d6a5fd3ddf53c34f82.png" target="_blank"><img alt="ARM Cortex-A17详细解析：比肩A15、踩死A12" src="/img/20140212/s_3e086eb7a66445d6a5fd3ddf53c34f82.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292415.htm" target="_blank" title="索尼的新花招：Walkman泡在水里出售"><span class="title1">索尼的新花招：Walkman泡在水里出售</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小路</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:53:11</span></td>
  </tr>
  <tr>
    <td align="left"><p>你能从路边的饮料自动售卖机里买到什么？可乐还是果汁？现在，索尼为你提供了新的选择&mdash;&mdash;泡在瓶装水里的Walkman。</p>
<p>索尼近日在新西兰启动了一项针对旗下NWZ-270系列随身听的特别营销攻势&mdash;&mdash;将随身听泡在瓶装水中出售，以突显产品出色防水性能。</p>
<p>据了解，索尼的<strong>NWZ-270系列随身听主打运动用户，将耳机与机身融为一体并采用耳挂佩戴方式，具有3分钟快速充电和2米的防水功能（切勿泡在海水里），基础售价99.99美元（约合610元人民币，国内官方报价629元）。</strong></p>
<p>未来，消费者将在新西兰境内城市健身馆、地铁站等主要地点周围的Walkman主题自动售卖机中见到这些悬浮在水中的索尼随身听，但如果你想要购买的话请先确保自己准备了足够数量的硬币。</p>
<p><strong>水装Walkman官方广告：</strong><a class="f14_link" href="http://v.youku.com/v_show/id_XNjcyMDE2MDYw.html" target="_blank">http://v.youku.com/v_show/id_XNjcyMDE2MDYw.html</a></p>
<p align="center"><a href="/img/20140212/1381c04aaed94d8bb2d7899ade806ec3.jpg" target="_blank"><img alt="索尼的新花招：Walkman泡在水里出售" src="/img/20140212/s_1381c04aaed94d8bb2d7899ade806ec3.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/c42ae4aff84b43079fd64c317f70d4bc.jpg" target="_blank"><img alt="索尼的新花招：Walkman泡在水里出售" src="/img/20140212/s_c42ae4aff84b43079fd64c317f70d4bc.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/25698a4405a24057b55ffc9b9da2ecf8.jpg" target="_blank"><img alt="索尼的新花招：Walkman泡在水里出售" src="/img/20140212/s_25698a4405a24057b55ffc9b9da2ecf8.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/01edf7bb659e453f95bb056864a72368.jpg" target="_blank"><img alt="索尼的新花招：Walkman泡在水里出售" src="/img/20140212/s_01edf7bb659e453f95bb056864a72368.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/6a408ea3abf74e16a5243ce9a8797873.jpg" target="_blank"><img alt="索尼的新花招：Walkman泡在水里出售" src="/img/20140212/s_6a408ea3abf74e16a5243ce9a8797873.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/8d107b16ae2041debdbbc9ad6887a30e.jpg" target="_blank"><img alt="索尼的新花招：Walkman泡在水里出售" src="/img/20140212/s_8d107b16ae2041debdbbc9ad6887a30e.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292413.htm" target="_blank" title="极客最爱 星球大战长裙亮相纽约时装周"><span class="title1">极客最爱 星球大战长裙亮相纽约时装周</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:48:42</span></td>
  </tr>
  <tr>
    <td align="left"><p>作为四大时装周里最烧钱同时也是商业气息最浓厚的纽约时装周，虽然有时被吐槽创新不足但却是却最接地气的时装周。</p>
<p>在近日的纽约秋冬时装周上，著名女装品牌Rodarte发布了五款带有星球大战元素的长裙，包括天行者阿纳金、机器人C-3PO,尤达大师、死星以及阿纳金的故乡塔图因，虽然这些星战元素跟性感毫不沾边，但依旧被模特们演绎的风情万种。</p>
<p>可能大部分女性对星战都不感兴趣，但毕竟女为悦己者容，如果有一个极客男友那么这些款长裙肯定很有吸引力。</p>
<p align="center"><a href="/img/20140212/743926a3a20b4a0bb46969e85a741f9e.jpg" target="_blank"><img alt="极客最爱 星球大战长裙亮相纽约时装周" src="/img/20140212/s_743926a3a20b4a0bb46969e85a741f9e.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/f73ed135de204f19b9694bb77b4629d1.jpg" target="_blank"><img alt="极客最爱 星球大战长裙亮相纽约时装周" src="/img/20140212/s_f73ed135de204f19b9694bb77b4629d1.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/4ee18e4ecd4c466ab18cb5eb25469e62.jpg" target="_blank"><img alt="极客最爱 星球大战长裙亮相纽约时装周" src="/img/20140212/s_4ee18e4ecd4c466ab18cb5eb25469e62.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/2e3edc7128604f00b1fa2ad7991089e7.jpg" target="_blank"><img alt="极客最爱 星球大战长裙亮相纽约时装周" src="/img/20140212/s_2e3edc7128604f00b1fa2ad7991089e7.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/e30601e6ed934f47a1650fb766a4d70e.jpg" target="_blank"><img alt="极客最爱 星球大战长裙亮相纽约时装周" src="/img/20140212/s_e30601e6ed934f47a1650fb766a4d70e.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292412.htm" target="_blank" title="电信版红米/小米3还得等？"><span class="title1">电信版红米/小米3还得等？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:43:59</span></td>
  </tr>
  <tr>
    <td align="left"><p>春节前夕，小米科技CEO雷军在微博上表示<a class="f14_link" href="http://news.mydrivers.com/1/290/290365.htm" target="_blank">小米手机3电信版和红米电信版都将在2月中旬正式开卖</a>。日前我们也作出<a class="f14_link" href="http://news.mydrivers.com/1/291/291897.htm" target="_blank">预测称电信版红米/小米3都将在2月18日正式开卖</a>，但在此我们要非常遗憾的告诉大家：电信版红米/小米3可能还得等。</p>
<p>在昨天的开放购买活动结束之后，<strong>小米已经开始了下一轮开放购买预约，具体时间为2月18日中午12点</strong>。但在查询了开放购买产品列表之后<strong>我们遗憾的发现电信版红米和小米手机3并没有出现在列表当中</strong>，所以2月18日开卖是不可能的了。</p>
<p>我们猜测<strong>电信版红米不能开放购买的原因是还没有获得入网许可证</strong>，至少目前在工信部网站上暂时还查询不到。而电信版小米手机3还未开卖应该是要等待电信版红米一同上市。希望2月25日的本月最后一轮开放购买中会出现这两款产品。</p>
<p align="center"><a href="/img/20140212/8f506d26fde54feca823ea18eabb2d32.jpg" target="_blank"><img alt="电信版红米/小米3还得等？" src="/img/20140212/s_8f506d26fde54feca823ea18eabb2d32.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
2月18日开放购买列表</p>
<p align="center"><img alt="电信版红米/小米3还得等？" src="/img/20140212/78819fcd5b6f48628d7ff08cea5e20e7.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
电信版红米手机尚未拿到入网许可，而电信版米3已经拿到了（2013063）</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292411.htm" target="_blank" title="科技感十足：魅族智能手表Mwatch首曝光"><span class="title1">科技感十足：魅族智能手表Mwatch首曝光</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:36:45</span></td>
  </tr>
  <tr>
    <td align="left"><p>虽然从去年开始，三星、索尼以及盛大等厂商就已经开始发力智能手表，但在苹果iWatch到来之前，尚未出现一款能够引领潮流的产品出现。不过，可以肯定的是，智能手表将是继智能手机之后，未来智能设备的又一产品形态，也会有越来越多的厂商加入这一领域。</p>
<p>就在刚刚，<strong>魅族高级经理张恒在微博上曝光了魅族内部的智能手表概念产品图：Mwatch</strong>。这是暗示魅族也要做智能手表了？</p>
<p>从曝光的概念图来看，<strong>魅族智能手表采用2.4寸、480x640分辨率柔性全贴合触摸屏，配备500万像素自动变焦摄像头，左侧设有物理按键。手表采用高强度磨砂金属边框，含心率探测的高强度柔性合金表带，表带内衬为带凹凸纹理高能胶体</strong>。</p>
<p>此外，<strong>Mwatch支持超过5000个Apps应用程序，支持云服务，可通过Wi-Fi、蓝牙、NFC连接魅族智能手机使用</strong>。其它特色功能暂时未知。</p>
<p>目前，尚不清楚魅族是否已经着手Mwatch的研发，但至少表明魅族已经有这方面的打算了。</p>
<p align="center"><a href="/img/20140212/be53f330af4c4b0d9ff2e3525b9ac3e6.jpg" target="_blank"><img alt="魅族智能手表Mwatch首曝光" src="/img/20140212/s_be53f330af4c4b0d9ff2e3525b9ac3e6.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/2bb357dc61a94670bcb93a7f2da8d735.jpg" target="_blank"><img alt="魅族智能手表Mwatch首曝光" src="/img/20140212/s_2bb357dc61a94670bcb93a7f2da8d735.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/b9f80cf0803e42e1a2df962a8e2c39de.jpg" target="_blank"><img alt="魅族智能手表Mwatch首曝光" src="/img/20140212/s_b9f80cf0803e42e1a2df962a8e2c39de.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p>　　</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292410.htm" target="_blank" title="Surface通过FAA认证 占领飞机驾驶舱"><span class="title1">Surface通过FAA认证 占领飞机驾驶舱</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:34:02</span></td>
  </tr>
  <tr>
    <td align="left"><p>昨日，美国联邦航空管理局(简称FAA)正式授权微软Surface 2和Surface Pro 2平板机可以在飞机飞行的所有阶段作为一级或二级电子飞行包（EFB）使用。</p>
<p>在本周的新加坡航空展上，Jeppesen展示了他们为Windows 8.1开发的FliteDeck Pro应用程序，飞行员能更加轻松、安全地查看数据，例如实时天气、新闻。</p>
<p>尽管商务飞机完全实现无纸化驾驶舱不是一朝一夕的事，不过IT设备在飞机上的迅猛增长已然是无法阻挡的。据2013航空运输业趋势调查报告指出，2016年，87%的商务航班、所有飞机中的71%都将部署平板机。</p>
<p align="center"><a href="/img/20140212/cbe10b6cc7a2418586e25c650d259976.jpg" target="_blank"><img alt="Surface通过FAA认证 占领飞机驾驶舱" src="/img/20140212/s_cbe10b6cc7a2418586e25c650d259976.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/4a632f8a45354818b700b39ffea3635c.png" target="_blank"><img alt="Surface通过FAA认证 占领飞机驾驶舱" src="/img/20140212/s_4a632f8a45354818b700b39ffea3635c.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/56066f06483646bfa85bee77a833ab3f.png" target="_blank"><img alt="Surface通过FAA认证 占领飞机驾驶舱" src="/img/20140212/s_56066f06483646bfa85bee77a833ab3f.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p>&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292409.htm" target="_blank" title="罗永浩：只有苹果和锤子能拯救中文字体丑陋"><span class="title1">罗永浩：只有苹果和锤子能拯救中文字体丑陋</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">随心</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:32:53</span></td>
  </tr>
  <tr>
    <td align="left"><p>昨天大半夜的，罗永浩又发微博称，&ldquo;<strong>除了苹果和锤子，没有第三家智能设备厂商解决了中文字体丑陋的问题。</strong>&rdquo;看来他对自家锤子手机的字体美观还是比较自信的。</p>
<p>就桌面PC来看，Mac确实要比Windows PC的字体更富美感。iOS字体也要比Android养眼一些，这其中有哪些道道呢？我们能不能在Windows PC上实现Mac的字体效果呢？</p>
<p>从<a class="f14_link" href="http://www.zhihu.com/question/21211748/answer/17526912" target="_blank">知乎</a>我们了解到，android采用了linux的freetype字体库和chrome的2d图形渲染引擎skia来显示字体，这样的选择是为了兼顾清晰和美观，<strong>毕竟android手机从高端到低端，ppi差别很大，不可能像ios一样只追求美观，不需要考虑低ppi下字体显示是否清晰</strong>（低分辨率显示器显示mac上的字体就不如windows清晰，高分辨率下mac字体更美观）。</p>
<p>原生android系统要适用高中低所有android手机，字体设计必须考虑兼容性，不考虑这个大前提而贬低android的字体显示效果无异是耍流氓。</p>
<p><strong>文中还给出锤子手机字体更加美观的可能方案：更换字体。</strong>从先前消息，锤子手机的定位不会是入门级，屏幕分辨率不会很低，因此更换字体在高ppi下值得一试。&ldquo;但这样做却没办法根本解决字体库和字体渲染问题，如果有能耐把ios的mactype移植到android才真牛逼，毕竟换字体在任何root了的手机上都轻而易举。&rdquo;</p>
<p>到底锤子手机的字体多么美观，我们只能拭目以待了。不过现在<strong>正在使用Windows PC的童鞋可以试试<a class="f14_link" href="http://bbs.mydrivers.com/thread-353933-1-1.html" target="_blank">MacType</a>这款字体渲染软件</strong>，相信效果会比较满意的。</p>
<p align="center"><img alt="罗永浩：只有苹果和锤子能拯救中文字体丑陋" src="/img/20140212/83a0ec3bbb0145be88f10a8221fd6e80.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p align="center"><img alt="罗永浩：只有苹果和锤子能拯救中文字体丑陋" height="124" src="/img/20140212/2902e4f73dfc4999b729d64fe9ca98e5.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" width="303" /><br />
Windows下正常显示效果</p>
<p align="center"><img alt="罗永浩：只有苹果和锤子能拯救中文字体丑陋" src="/img/20140212/507dacf7501c47d5afa6a9113964a601.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
Windows PC下实现iOS风格字体</p>
<p align="center"><img alt="罗永浩：只有苹果和锤子能拯救中文字体丑陋" src="/img/20140212/19e85f9d8de74ef597be2ce2238082ec.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
正常字体效果</p>
<p align="center"><img alt="罗永浩：只有苹果和锤子能拯救中文字体丑陋" src="/img/20140212/a8e75bb3b01a4cf9bdf02b9c2388dcc0.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
渲染效果</p>
<p align="center"><img alt="罗永浩：只有苹果和锤子能拯救中文字体丑陋" src="/img/20140212/6b997825f34c426d9ae1c76a5a14585f.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
在状态栏图标处可以选择更多渲染效果</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292408.htm" target="_blank" title="影响2014中国手机市场的六大技术趋势"><span class="title1">影响2014中国手机市场的六大技术趋势</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:27:23</span></td>
  </tr>
  <tr>
    <td align="left"><p>第三方分析机构IDC周二称，2014年中国智能手机市场仍将保持较大幅度的增长，但智能手机普及率的提升将导致增速减缓。智能手机普及之后，产品必然会更新换代。</p>
<p>IDC中国负责手机市场跟踪报告的高级分析师闫占孟表示，2014年电信运营商会对定制机型的新产品硬件配置提出更高的要求，手机厂商会为满足电信运营商的需求而不断调整策略。</p>
<p>他说，由于产品的降价和新品的大量上市，厂商为了吸引最终用户和完成更高的销售目标，也会开展更加多样的市场活动。2014年中国智能手机市场的竞争将会更加激烈。因此，IDC分析了影响2014年中国手机市场六种技术趋势：</p>
<p><strong>1、芯片厂商对五模或者多模4G芯片平台的支持，会推动移动产业的革新</strong>。IDC预测在中国4G手机市场发展初期，首先会集中部署三模的4G芯片。待芯片厂商技术成熟和网络部署完善后，将会有更多的芯片供应商发布4G的五模或者多模芯片平台。融合的五模或者多模的4G芯片平台会逐渐成为移动终端的主要模式。</p>
<p><strong>2、低温多晶硅工艺的2K屏幕，会被更多的高端智能手机所采用</strong>。越来越多的消费者已经开始使用大屏智能手机观看高清视频，而高清屏幕是观看移动视频获得优良体验的基础，所以厂商对高清屏幕的需求在快速增长。最新一代低温多晶硅工艺的2K屏幕，它不但具有高分辨率、高色彩饱和度、成本低廉的特点，还可降低电力消耗。</p>
<p><strong>3、光学防抖技术在高端智能手机的应用，将成为热点</strong>。 智能手机的摄像头已经从200万像素发展到1600万像素，甚至更高，目前已经很难再有突破，而用户对拍照防抖功能的需求依然旺盛。</p>
<p><strong>4、NFC技术需要电信运营商、终端厂商和金融系统集体合作，才能得到普及</strong>。虽然2013年已有一些手机支持NFC功能，但是缺乏大型企业的推动和产业链上下游的广泛合作，导致NFC技术应用市场声音偏少，且NFC技术应用较窄，不能满足用户的更多需求。</p>
<p><strong>5、弯曲技术在手机的应用未来前景广阔</strong>。弯曲技术在手机上的应用主要体现在：屏幕可弯曲和电池可弯曲。目前这两项弯曲技术已经有所突破，但是还面临诸多难题，例如：在屏幕可弯曲的前提下，高分辨率和色彩纯度较难得到保证。同时，找到低成本量产的方法也是其发展的关键，这对于目前正在发展的弯曲技术是一个挑战，但同时也是一个机会。</p>
<p><strong>6、双智能操作系统的应用为用户带来不同的体验</strong>。双操作系统在手机上已经有过应用，但是真正的双智能操作系统并不多见。在2014年，智能手机操作系统的竞争会更加激烈，厂商出于创新和业务拓展的需要，将会在一部智能手机上开发两个操作系统。双操作系统的应用也可为用户带来不同的新鲜体验。同时，多操作系统也为跨平台的个人云服务提供了更多的发展空间。</p>
<p align="center"><a href="/img/20140212/3a30c8e8c3ef45edab8401074b7b538f.jpg" target="_blank"><img alt="影响2014中国手机市场的六大技术趋势" src="/img/20140212/s_3a30c8e8c3ef45edab8401074b7b538f.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292407.htm" target="_blank" title="干掉爱立信 华为成全球最大设备商"><span class="title1">干掉爱立信 华为成全球最大设备商</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:27:03</span></td>
  </tr>
  <tr>
    <td align="left"><p>1月31日，爱立信发布了2013年全年及第四季财报。财报显示，其全年净销售额2274亿瑞典克朗(约合353亿美元)，与2012年持平；净利润122亿瑞典克朗，同比增长105%(约合18.6亿美元)。<strong>尽管爱立信2013年实现扭亏为盈，但在全年销售上被华为首次超越。</strong></p>
<p>华为在1月15日公布的业绩称其全年销售收入为2380~2400亿元人民币，约合393~396亿美元。如此，<strong>华为理所当然地登上了全球电信设备商榜首。</strong></p>
<p>华为首席财务官孟晚舟表示，华为2013年经营性现金流和资产负债率均保持稳定；集团坚持推行管理变革，通过简化管理实现运营效率持续提升，未来五年销售收入复合增长率将保持10%左右的水平。</p>
<p><strong>华为超越爱立信成全球最大设备商</strong></p>
<p>其实，华为成为全球设备商的势头早在2012年就显山露水。<strong>在2012年上半年，华为曾以161亿美元的营收一度超越爱立信的152.5亿美元，不过在下半年又被爱立信反超。</strong>2012年年，华为营收2202亿人民币(约合353.6亿美元)，净利润154亿人民币(约合24.69亿美元)。从营收规模和行业排序规则来看，华为超越爱立信仅一步之遥。</p>
<p><strong>而此次，爱立信被华为赶超多少在情理之中。</strong>据透露，华为去年企业业务营收占总营收的7%，消费者业务占23%，成为两大增长点，运营商业务占七成比例。</p>
<p>爱立信首席执行官卫翰思在董事会会议中表示：&ldquo;公司第四季度的营收面临着一些压力。这主要是因为美国和日本两国的运营商减少了电信基础设备投入。&rdquo;</p>
<p>截至目前，全球五大电信础设施供应商中，华为、爱立信、诺基亚、阿朗4家都披露了2013年全年业绩。按2013年12月31日汇率折算，主营业务收入分别为393-396亿美元、353亿美元、152亿美元。阿朗受到移动业务减记、重组费用的拖累，2013年全年也净亏13.04亿欧元。</p>
<p><strong>加大转型步伐，提升运营效率</strong></p>
<p>过去几年，在全球经济形势持续低迷、电信基础设施投资增长乏力的情况下，电信设备商的市场空间可谓如履薄冰，而爱立信与华为这两家行业巨头却走上两条完全不同的转型道路。</p>
<p><strong>爱立信认为运营商市场远远未达到天花板</strong>，因此对于运营商市场的专注度可见一斑，先后发力电信管理服务、运营支撑系统与业务支撑系统OSS/BSS，2013年这两块业务的净销售额总计为1096亿瑞典克朗，约占整体收入的48%；并重组专利部门、货币化其知识产权资产。</p>
<p><strong>而华为从2010年开始战略转型，提出&ldquo;云算端&rdquo;的云计算方案</strong>：重塑 ICT、进军企业网、重视终端(手机、数据卡等)消费市场。2011年，华为开启了向ICT综合设备商转型的征程，试图将电信领域的竞争优势延伸到企业及消费级市场，构建了多元化的产业链，以引导公司经营领域所带来的机遇。2013年，华为在企业级市场多点开花，并初步构建了一条完善的企业级产品线。&ldquo;在数据中心，有服务器、存储、云计算等，在企业网络领域，有敏捷网络、敏捷交换机等，在UC领域有智真、统一通讯、融合会议等。</p>
<p><strong>在终端上，华为也没有落下。荣耀品牌独立，终端业务向B2C业务转型都显示了其企业基因的包容性。</strong>华为终端全球交付与服务部部长郭新心表示，：&ldquo;华为终端必须更加主动和开放，必须学会以消费者的思维来思考问题，必须进行服务转型。&rdquo;</p>
<p>此外，华为近年来还加大了对新技术的研发。据了解每年都将销售收入10%以上的资金用于创新。2004-2013年华为累计在研发创新上投入1539亿元人民币。仅在2013年，公司的研发投入就高达约330亿元人民币，占销售收入的约14%。</p>
<p>可以看出，进过这几年的转型，华为的盈利能力得到了显著地提升。孟晚舟说，通过简化管理，向一线放权，整个公司的运营效率得到提升。</p>
<p><strong>4G是重要战略机遇</strong></p>
<p>在日前工信部工作会议上工信部部长苗圩表示，到2014年底，国内4G商用城市数量超过300个，用户数超过3000万。这是更提及，4G对于2014年通信业来说是一大投资亮点，仅这一领域的投资就将达到1000亿元人民币。从2009年到2013年华为过去5年的主营业务利润分别为222亿、307亿、186亿、200亿和286~294亿，2010年获得丰厚利润的背景正是中国3G的启动，现在随着中国4G的启动，电信设备商明年的&ldquo;收成&rdquo;充满了想象空间。</p>
<p>而在最为密切的4G元年，华为的4G建设之路没有落下。<strong>在全球商用LTE网络数244张中，华为参与部署了110张商用网络，华为LTE进入全球100多个首都城市及九大金融中心</strong>。此外，<strong>华为已经在全球获得了200多个4G商用合作，位列业界第一。</strong>在终端领域，华为曾发布全球首款LTE多模数据卡以及最早支持LTE Cat4(下行速率150Mbps)的手机。</p>
<p>华为中国区终端副总裁和CMO表示，如果说1G时代是摩托罗拉的天下，2G时代是诺基亚的时代，3G时代是苹果、三星的时代，那么4G时代一定是华为的天下。&rdquo;对于华为的未来，孟晚舟认为，物联网和云计算将给电信业带来更广阔的空间。无论是传统的运营商业务，还是华为从前年开始重点发力的企业业务和消费者业务，都具有广阔的增长空间。</p>
<p>不过，在移动互联网时代，华为的营收和利润虽然令人嫉羡，但越来越多的人并不看好华为，认为&ldquo;廉颇已老，华为已死&rdquo;。</p>
<p>其实，无论是看好还是看衰，对于百花争鸣的4G元年，华为所能做的只有进一步向终端、设备、运营商多元化并举，坚持拥抱变革的态度求同存异，相信在未来，华为会给我们呈现更多的意想不到的精彩。</p>
<p align="center"><a href="/img/20140212/cf0281cf1e534b2ba11e21eb67e3398a.jpg" target="_blank"><img alt="干掉爱立信 华为成全球最大设备商" src="/img/20140212/s_cf0281cf1e534b2ba11e21eb67e3398a.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292406.htm" target="_blank" title="《泰坦降临》Beta资格申请现已开始"><span class="title1">《泰坦降临》Beta资格申请现已开始</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">Zhengogo</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:23:48</span></td>
  </tr>
  <tr>
    <td align="left"><p>Respawn Entertainment工作室及项目负责人Vince Zampella早前发推称beta测试资格将于太平洋标准时下午六点开始申请，但申请程序其实已提前启动：</p>
<p style="text-align:center"><a href="/img/20140212/d9f6d41595a948f987611184691f8254.jpg" target="_blank"><img alt="《泰坦降临》Beta资格申请现已开始" src="/img/20140212/s_d9f6d41595a948f987611184691f8254.jpg" style="border:1px solid black;" /></a></p>
<p>玩家现在就可前往该作<a class="f14_link" href="http://www.titanfall.com/beta" target="_blank">官方网站</a>申请自己的beta测试资格（需Origin帐号）：</p>
<p style="text-align:center"><a href="/img/20140212/47b21086ab0d4a91acd4f57c994c3f6c.jpg" target="_blank"><img alt="《泰坦降临》Beta资格申请现已开始" src="/img/20140212/s_47b21086ab0d4a91acd4f57c994c3f6c.jpg" style="border:1px solid black;" /></a></p>
<p>本轮beta测试仅面向PC及Xbox One用户，不涉及Xbox 360平台，无需事先预订游戏。</p>
<p>另外，Grinding Gear Games宣布将于当地时间3月5号发布《放逐之路》（Path of Exile）迷你资料片，该资料片将以1.1.0版补丁的形式提供。</p>
<p>《放逐之路》可谓RPG领域的一匹黑马，其素质全方位秒杀D3无悬念&mdash;&mdash;玩过的人知道我在说什么，PoE才是真正的Diablo 3！</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292404.htm" target="_blank" title="不做死不会死：OCZ接着贱卖自己"><span class="title1">不做死不会死：OCZ接着贱卖自己</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:18:49</span></td>
  </tr>
  <tr>
    <td align="left"><p>曾经在DIY领域叱咤风云的<a class="f14_link" href="http://news.mydrivers.com/1/284/284513.htm" target="_blank">OCZ终于在去年底把自己给玩死了</a>，宣布破产之后被东芝全额收购，<a class="f14_link" href="http://news.mydrivers.com/1/290/290791.htm" target="_blank">成为了东芝旗下的一个全资子公司</a>，继续独立运营，名字从&ldquo;OCZ Technology&rdquo;改成&ldquo;OCZ Storage Solutions&rdquo;。</p>
<p>事实上<strong>东芝只对OCZ的存储业务感兴趣</strong>，也就是OCZ旗下的SSD业务，对于其它业务并没有继续发展下去的打算，因此OCZ只能把旗下的其它业务都贱卖出去。</p>
<p>日前，来自<strong>美国特拉华州的一家电源公司FirePower宣布收购了OCZ旗下的PC Power及Cooling&reg;电源业务，令人震惊的是成交额只有95万美元</strong>，要知道OCZ旗下的电源此前在DIY圈子里是颇受好评的产品，名声并不像SSD那么差劲，这次以95万美元的价格贱卖出去，收购者FirePower恐怕做梦都要笑出来。</p>
<p>在官网的<a class="f14_link" href="http://www.pcpower.com/news/FirePower_Technology_Acquires_Power_Supply_Assets_of_OCZ_Technology_Group_and_PC_Power_and_Cooling.html" target="_blank">收购声明</a>中FirePower并没有公布究竟会如何处理OCZ的电源业务，但我们还是希望它能够存在下去，毕竟对于一些老牌DIYer来说，OCZ的电源确实是一代经典，就这么消失实在是可惜了。</p>
<p align="center"><a href="/img/20140212/3c16bf0c072445d297efb0b8721de1d5.jpg" target="_blank"><img alt="不做死不会死：OCZ接着贱卖自己" src="/img/20140212/s_3c16bf0c072445d297efb0b8721de1d5.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292403.htm" target="_blank" title="情人节要送啥？鲜花、巧克力都落伍了"><span class="title1">情人节要送啥？鲜花、巧克力都落伍了</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:13:52</span></td>
  </tr>
  <tr>
    <td align="left"><p>本周五是西方的情人节，想必不少情侣们都要考虑送另一半的礼物了。送花，送巧克力？都落伍了。</p>
<p>据国外网站Play最新的一项调查显示，<strong>对于期望收到的情人节礼物，大家最青睐的首选是iPad，第二个才是鲜花</strong>。紧跟其后的<strong>第三、第四、第五分别是珠宝、笔记本电脑以及PS4游戏机</strong>，而以往的热门礼物巧克力则被挤出前五。</p>
<p>果然是长得好、还有钱的高富帅、白富美们才有情人节，其他的人还是想想如何过元宵节吧！</p>
<p align="center"><a href="/img/20140212/14daae8e7b7e4fd1a4f822a4f001e13c.jpg" target="_blank"><img alt="情人节要送啥？鲜花、巧克力都落伍了" src="/img/20140212/s_14daae8e7b7e4fd1a4f822a4f001e13c.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292402.htm" target="_blank" title="专家：WinXP停止服务后 电脑10分钟就会出事"><span class="title1">专家：WinXP停止服务后 电脑10分钟就会出事</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">刘艺</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:13:03</span></td>
  </tr>
  <tr>
    <td align="left"><p>现在距离Windows XP生命周期结束只剩下不到60天的时间了，微软已拼尽浑身解数劝XP用户尽快升级新系统，但目前仍有29%的用户在坚守着老系统。</p>
<p>今天，一位曾为军用计算机专家及网络工程师的Michael Menor提醒广大用户，<strong>一旦微软在4月8号停止对XP的支持服务，这些用户的电脑在10分钟之内便会受到感染，网络罪犯们不会放过任何一个可攻击的漏洞。</strong></p>
<p>对此微软也表示：&ldquo;如果你在XP停止支持之后还坚持使用该系统，并联网使用你的计算机的话，你的电脑就会陷入重大危险中。&rdquo;</p>
<p align="center"><a href="/img/20140212/4bdca89314bb4eb882e3784e678de2bd.jpg" target="_blank"><img alt="专家：WinXP停止服务后 电脑10分钟就会出事" src="/img/20140212/s_4bdca89314bb4eb882e3784e678de2bd.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292401.htm" target="_blank" title="纠结的蓝宝石屏：iPhone 6向左 iWatch向右"><span class="title1">纠结的蓝宝石屏：iPhone 6向左 iWatch向右</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:11:55</span></td>
  </tr>
  <tr>
    <td align="left"><p>虽然苹果没有明说，<a class="f14_link" href="http://news.mydrivers.com/1/291/291863.htm" target="_blank"><strong>但iWatch已经可以断定是库克口中多次提到的全新产品了</strong></a>。</p>
<p>对于这款可穿戴设备来说，<a class="f14_link" href="http://news.mydrivers.com/1/292/292279.htm" target="_blank">iWatch将会以智能手表的形式出现，并且主打健康功能</a>。当然了，从苹果近期的表现和库克多次的发言我们不难看出，这款苹果全新的产品会在今年跟大家见面，至于具体时间嘛，极有可能是在WWDC 2014上，跟iOS 8一起亮相。</p>
<p>此外，iWatch即将跟大家见面的另一个征兆那就是产业链的活跃。<strong>今天早些时候，台湾产业链给出的消息称，iWatch确定配备蓝宝石屏幕，其屏幕大小在2寸左右</strong>。</p>
<p>之前还有消息称，iPhone 6会采用蓝宝石屏幕，<a class="f14_link" href="http://news.mydrivers.com/1/292/292083.htm" target="_blank">而富士康也已经组装出了100台相对应的原型机</a>，不过最新的消息显示，<strong>由于配备蓝宝石屏幕后，新一代iPhone的售价必定要上涨，所以成本提升是苹果最后放弃的主要因素</strong>。</p>
<p>希望iWatch能让我们眼前一亮，因为苹果已经好久没有带来这样的新品了。</p>
<p><strong>相关阅读：</strong></p>
<p><a class="f14_link" href="http://news.mydrivers.com/1/292/292365.htm" target="_blank">什么？iPhone 6屏幕竟然这么屌</a></p>
<p><a class="f14_link" href="http://news.mydrivers.com/1/292/292360.htm" target="_blank">iPhone 7前置摄像头要变变了</a></p>
<p align="center"><a href="/img/20140212/0d3abce3bacd4cb291879349dc21798c.jpg" target="_blank"><img alt="纠结的蓝宝石屏：iPhone 6向左 iWatch向右" src="/img/20140212/s_0d3abce3bacd4cb291879349dc21798c.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
概念iWatch设计</p>
<p align="center"><a href="/img/20140212/7f9e66c914c74873984e09b14e94724e.jpg" target="_blank"><img alt="纠结的蓝宝石屏：iPhone 6向左 iWatch向右" src="/img/20140212/s_7f9e66c914c74873984e09b14e94724e.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/9bff015a927d4e78894b9995e5b4a073.jpg" target="_blank"><img alt="纠结的蓝宝石屏：iPhone 6向左 iWatch向右" src="/img/20140212/s_9bff015a927d4e78894b9995e5b4a073.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/6ff2922254c545dd82b02916d087c72e.jpg" target="_blank"><img alt="纠结的蓝宝石屏：iPhone 6向左 iWatch向右" src="/img/20140212/s_6ff2922254c545dd82b02916d087c72e.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/2bc6e40ecd444999a41677f3a0dc2721.jpg" target="_blank"><img alt="纠结的蓝宝石屏：iPhone 6向左 iWatch向右" src="/img/20140212/s_2bc6e40ecd444999a41677f3a0dc2721.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/67a0eba1e1aa48abb8dcd6d0c9e1e59f.jpg" target="_blank"><img alt="纠结的蓝宝石屏：iPhone 6向左 iWatch向右" src="/img/20140212/s_67a0eba1e1aa48abb8dcd6d0c9e1e59f.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/6ea4b46ac83c454fb04371baed2f47d1.jpg" target="_blank"><img alt="纠结的蓝宝石屏：iPhone 6向左 iWatch向右" src="/img/20140212/s_6ea4b46ac83c454fb04371baed2f47d1.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/d62315bae17d40f9ae8c6109f7413a86.jpg" target="_blank"><img alt="纠结的蓝宝石屏：iPhone 6向左 iWatch向右" src="/img/20140212/s_d62315bae17d40f9ae8c6109f7413a86.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292400.htm" target="_blank" title="AMD还真敢再做双芯：已经能开机了"><span class="title1">AMD还真敢再做双芯：已经能开机了</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:09:40</span></td>
  </tr>
  <tr>
    <td align="left"><p>双芯显卡是实力的象征，是发烧友的终极追求，但同时也对设计能力要求甚为苛刻。近来陆续有传闻称，AMD、NVIDIA都在打造新一代的双芯显卡，都准备再&ldquo;烧&rdquo;一把。</p>
<p><a class="f14_link" href="http://news.mydrivers.com/1/282/282951.htm" target="_blank">AMD的会采用两颗Hawaii XT GPU核心</a>，也就是R9 290X二合一。最新消息称，<strong>这个怪物已经能够启动开机了，进展顺利，最快有望在六月初的台北电脑展上看到展示。</strong></p>
<p>但是具体规格不得而知，流处理器会不会全部开放，频率会降低到多少，功耗、温度、噪音能否让人承受都有待观察，不过核心既然是最高端的Hawaii XT，<strong>有可能会是流处理器开满、频率妥协</strong>。</p>
<p>至于命名，是叫做R9 290X&times;2？还是称为R9 295X？也还没有决定。</p>
<p><a class="f14_link" href="http://news.mydrivers.com/1/288/288681.htm" target="_blank">NVIDIA的双芯基本断定叫GeForce GTX 790</a>，GK110大核心二合一，4992个流处理器，10GB显存，300W功耗，最早三月份面世。</p>
<p align="center"><a href="/img/20140212/bd4207521550451b8ed01b83c75699fb.jpg" target="_blank"><img alt="AMD还真敢再做双芯：已经能开机了" src="/img/20140212/s_bd4207521550451b8ed01b83c75699fb.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
Hawaii XT</p>
<p align="center">&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292398.htm" target="_blank" title="余额宝首现“暂无收益”：回应称系统升级"><span class="title1">余额宝首现“暂无收益”：回应称系统升级</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:01:00</span></td>
  </tr>
  <tr>
    <td align="left"><p>今天早晨，不少余额宝用户发现，手机版余额宝首次出现&ldquo;暂无收益&rdquo;的情况，并未显示前一天的收益。<strong>由于是首次出现，而且官网之前并未提前公告，这一情况引起了不少人的猜测和不安。不过，经证实，系系统升级所致，不影响昨日收益和发放。</strong></p>
<p>据余额宝方面表示，<strong>这是由于余额宝昨晚系统升级，并不影响昨天收益的产生和发放。用户可在今天11点后正常查看收益情况</strong>。该人士强调，此次升级为常规升级，不会上线新功能，也不涉及改版。</p>
<p>9时许，余额宝发微博称：&ldquo;<strong>由于系统升级，收益稍后发放。粉儿们别急，一分也不会少</strong>。&rdquo;</p>
<p>支付宝方面也对此进行了回应：&ldquo;目前收益正在发放中，已经有好多用户看到了，暂时还没有看到收益的用户请稍等等，所有的收益会在今天中午12点前发放完毕。&rdquo;</p>
<p>春节过后，余额宝、财付通等多款理财产品都出现了收益下降的情况，目前余额宝7天年化收益率6.2630%。</p>
<p>此外，支付宝正在筹备推出&ldquo;元宵理财&rdquo;，预期收益率为7%，对接珠江人寿的一款万能险理财产品。</p>
<p align="center"><a href="/img/20140212/22482c312ed54479a0aeef961f07f0c3.jpg" target="_blank"><img alt="余额宝首现“暂无收益”：回应称系统升级" src="/img/20140212/s_22482c312ed54479a0aeef961f07f0c3.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/d01ad7204f444bdaafc8fbd9f8aaa0c0.jpg" target="_blank"><img alt="余额宝首现“暂无收益”：回应称系统升级" src="/img/20140212/s_d01ad7204f444bdaafc8fbd9f8aaa0c0.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/b29eeec48c24407698c122fdbf74ff8e.png" target="_blank"><img alt="余额宝首现“暂无收益”：回应称系统升级" src="/img/20140212/s_b29eeec48c24407698c122fdbf74ff8e.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
图片来自新浪科技</p>
<p>&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292397.htm" target="_blank" title="Surface Pro/Pro2、Surface 2固件更新"><span class="title1">Surface Pro/Pro2、Surface 2固件更新</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:00:04</span></td>
  </tr>
  <tr>
    <td align="left"><p>微软在昨天的周二补丁日也为Surface用户推送了固件更新，Surface Pro、Surface 2、Surface Pro 2用户都可以通过Windows Update安装最新固件（Surface RT无固件更新）。</p>
<p>以下是不同机型的固件更新内容：</p>
<p><strong><a class="f14_link" href="http://www.microsoft.com/surface/en-us/support/install-update-activate/pro-2-update-history" target="_blank">Surface Pro 2</a></strong></p>
<p>1、Surface Display Panel显示面板 (v1.0.2.0)</p>
<p>2、Surface Pro Embedded Controller Firmware嵌入式控制器固件 (v24.00.50)</p>
<p>3、Surface Pro System Aggregator Firmware系统聚合器固件 (v2.04.0150)</p>
<p>4、Surface Touch Cover 2 Firmware Update Device (v2.0.228.0)<br />
&nbsp;<br />
5、Surface Type Cover 2 Firmware Update Device (v2.0.226.0)</p>
<p><strong><a class="f14_link" href="http://www.microsoft.com/surface/en-us/support/install-update-activate/2-update-history" target="_blank">Surface 2</a></strong></p>
<p>1、Marvell AVASTAR Bluetooth Radio Adapter (v14.62.28058.10171)</p>
<p>2、Marvell AVASTAR 350N Wireless Network Controller (v14.62.28058.10171)</p>
<p><strong><a class="f14_link" href="http://www.microsoft.com/surface/en-us/support/install-update-activate/pro-update-history" target="_blank">Surface Pro (Windows 8.0/8.1更新)</a></strong></p>
<p>1、Surface Pro Embedded Controller Firmware (v9.00.50)</p>
<p>2、Surface Pro System Aggregator Firmware (v2.84.0150)</p>
<p>3、Surface Pro UEFI (v1.6.50)</p>
<p align="center"><a href="/img/20140212/e8d0dcc836514481b27bb7704665b982.jpg" target="_blank"><img alt="Surface Pro/Pro2、Surface 2固件更新" src="/img/20140212/s_e8d0dcc836514481b27bb7704665b982.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292396.htm" target="_blank" title="看 这就是联发科新八核神器MT6595"><span class="title1">看 这就是联发科新八核神器MT6595</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">10:00:04</span></td>
  </tr>
  <tr>
    <td align="left"><p>昨天ARM正式发布了<a class="f14_link" href="http://news.mydrivers.com/1/292/292331.htm" target="_blank">大马甲Cortex-A17架构</a>，<strong>仍然基于32位ARMv7指令集，架构和Cortex-A12完全相同</strong>，引入了新的一致性总线AMBA4 ACE(原来是AMBA4 AXI)，可以更快速地连接内存控制器，从而改善性能和能效。性能方面，ARM宣称Cortex-A17架构的性能可以逼近Cortex-A15，但这个命名实在是有点坑爹了。</p>
<p>在Cortex-A17架构发布不久，<a class="f14_link" href="http://news.mydrivers.com/1/292/292333.htm" target="_blank">联发科就发布了全球首款采用Cortex-A17架构的处理器MT6595</a>，号称是全球第一款真八核的4G LTE智能手机单芯片方案。这不得不让人怀疑是联发科和ARM商量好了的，选在同一个时间来发布新品。</p>
<p>目前几乎所有的消息都认为MT6595处理器不可能会采用八核Cortex-A17核心，<strong>而是采用四核心Cortex-A17搭配四核心Cortex-A7来组成big.LITTLE架构</strong>，前者最高频率可达2.5GHz，后者则可达1.7GHz。当然，八核心是可以同时运行的，否者联发科也不敢以八核来定位。</p>
<p>其图形核心为PowerVR G6200系列，和苹果A7处理器系出同门，但具体型号不明，只是表示<strong>屏幕分辨率最高支持2560&times;1600 WQXGA，摄像头最高支持2000万像素。</strong></p>
<p>基带方面是MT6595处理器的另一大两点，支持LTE Release 9 Cat.4，最高上传、下载速率分别为50Mbps、150Mbps；支持双模4G TD-LTE、FDD-LTE，也支持DC-HSPA+ 42Mbps、TD-SCDMA、EDGE/GSM。搭配联发科自己出品的射频芯片之后能够支持30多个频段，几乎做到了全球漫游（CDMA网络下就不行了）。</p>
<p>最后令人震惊的就是今天外媒居然找到了联发科MT6595处理器的&ldquo;靓照&rdquo;，要知道这在此前是根本不可能的，吝啬的联发科连规格参数都不肯说清楚，又怎么会公布照片呢？这次真不知道是抽了哪门子疯。</p>
<p align="center"><a href="/img/20140212/7862646e34a9432bb5a04eed1a0180d5.jpg" target="_blank"><img alt="看 这就是联发科新八核神器MT6595" src="/img/20140212/s_7862646e34a9432bb5a04eed1a0180d5.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292395.htm" target="_blank" title="NVIDIA不走寻常路：谁说黑盒版就是黑的？"><span class="title1">NVIDIA不走寻常路：谁说黑盒版就是黑的？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:59:22</span></td>
  </tr>
  <tr>
    <td align="left"><p>NVIDIA GK110核心的终极版本GeForce GTX Titan Black正风雨欲来，但是这个&ldquo;黑盒版&rdquo;可能会让你大跌眼镜，因为它其实并没有想象中那么黑。</p>
<p>VideoCardz独家获得了该卡的第一张照片：</p>
<p align="center"><a href="/img/20140212/2a00b8976ce546f9a397ed0fe4799abe.jpg" target="_blank"><img alt="GTX Titan Black真身首曝：其实不黑" src="/img/20140212/s_2a00b8976ce546f9a397ed0fe4799abe.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/d8795513998c4fe9a4a83ce51278a9bd.jpg" target="_blank"><img alt="GTX Titan Black真身首曝：其实不黑" src="/img/20140212/s_d8795513998c4fe9a4a83ce51278a9bd.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p><strong>相比于现在的GTX Titan，的确是黑了一些，但是仅限外壳原来的灰色部分，以及头部的TITAN几个字母，其余银色的外壳依然是银色。事实上，倒是和GTX 780 Ti几乎如出一辙。</strong></p>
<p>VideoCardz特别说：&ldquo;我知道这有点蠢，但看起来它是真的。<strong>可以向大家保证，这不是PS的。</strong>&rdquo;</p>
<p>管它呢，谁说黑盒版就一定要全黑了？</p>
<p>规格方面暂无进一步消息，<a class="f14_link" href="http://news.mydrivers.com/1/291/291245.htm" target="_blank">此前称会开满全部2880个流处理器，搭配6GB显存</a>，双精度计算性能完全释放。</p>
<p>据说会在本月18日发布，价格999美元，不过这个时间几乎可以肯定不靠谱，因为那天会发布GTX 750 Ti、GTX 750，容不下它了。</p>
<p align="center"><a href="/img/20140212/28ec155e6cbf41079dc174013cc103db.jpg" target="_blank"><img alt="GTX Titan Black真身首曝：其实不黑" src="/img/20140212/s_28ec155e6cbf41079dc174013cc103db.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
GTX Titan</p>
<p style="text-align: center"><a href="/img/20140212/ff9818b156ea478bb4d27a6804b2a031.jpg" target="_blank"><img alt="GTX Titan Black真身首曝：其实不黑" src="/img/20140212/s_ff9818b156ea478bb4d27a6804b2a031.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
GTX 780 Ti</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292394.htm" target="_blank" title="索尼A6000真机首曝 双物理拨盘"><span class="title1">索尼A6000真机首曝 双物理拨盘</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:45:26</span></td>
  </tr>
  <tr>
    <td align="left"><p>A6000马上就要发布了，虽然之前关于其配置已经有不少爆料，但是外观方面却没有任何消息，作为NEX-6和NEX-7两款相机的继任者，到底长得会像谁呢？</p>
<p>日媒最新爆料了一张据说是A6000的真机图片，根据这张图片<strong>A6000有两个物理拨盘，其中一个可以自定义功能，这点更像NEX-7。</strong></p>
<p>其他方面，内置EVF、热靴、内闪一个不少，不过快门键好像小了一些，<strong>关于EVF之前SAR猜测可能使用的是和全画幅旗舰A99以及A7/A7R一样规格的OLED电子取景器，236万像素，视野率100%。</strong></p>
<p>配置上可以确定的是<strong>A6000会采用2430万像素的Exmor-R传感器以及Bionz X影像处理器</strong>，背部显示屏为3.0英寸，92 . 1万像素，另外据说A6000的对焦将速度会是所有E卡口相机里最快的，这点很值得期待。</p>
<p align="center"><a href="/img/20140212/d0295a852ea4423ab3c646e55df15893.jpg" target="_blank"><img alt="索尼A6000真机首曝 双物理拨盘" src="/img/20140212/s_d0295a852ea4423ab3c646e55df15893.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292393.htm" target="_blank" title="紧咬PS4：Xbox One官方立体声耳机亮相"><span class="title1">紧咬PS4：Xbox One官方立体声耳机亮相</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小路</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:45:13</span></td>
  </tr>
  <tr>
    <td align="left"><p>微软Xbox Live业务主管Larry Hryb近日突然公布了两款Xbox One官方配件，都与音频有关。除了首次亮相的头戴立体声耳机之外，还有此前传言已久的第三方耳机适配器。</p>
<p>Larry Hryb在自己的博客中表示，官方的<strong>Xbox One头戴立体声耳机拥有20Hz-20000Hz的响应范围</strong>，能够提供比现有Xbox One原装单声道耳麦更好的聊天、游戏音效，该产品<strong>计划在3月份上市，售价79.99美元（约合480元人民币），并随包装附送一个第三方耳机适配器。</strong></p>
<p>至于，<strong>微软第三方耳机适配器，其售价定在24.99美元（约合150元人民币）。</strong>它的一端设计有Xbox One的新型扁平音频接口，而另一端则<strong>支持传统的标准3.5mm/2.5mm耳机插头</strong>，可连接第三方耳机产品，同时适配器也<strong>兼顾音量调节</strong>功能。</p>
<p>巧合的是，索尼也于上周才刚刚宣布了自己的<a class="f14_link" href="http://news.mydrivers.com/1/291/291996.htm" target="_blank">PS4头戴耳机套装</a>，虽然售价高达99.99美元（约合610元人民币），但其能够模拟7.1声道音效，并支持无线和有线两种连接方式，计划本月内上市。</p>
<p><strong>相关阅读：</strong></p>
<p><a class="f14_link" href="http://news.mydrivers.com/1/292/292381.htm" target="_blank">索尼：PS4接班PS3？它还太年轻</a></p>
<p><a class="f14_link" href="http://news.mydrivers.com/1/291/291996.htm" target="_blank">图赏：PS4无线头戴耳机开箱靓照</a></p>
<p align="center"><a href="/img/20140212/b3f49680429c430eae8cf92845639229.jpg" target="_blank"><img alt="紧咬PS4：Xbox One官方立体声耳机亮相" src="/img/20140212/s_b3f49680429c430eae8cf92845639229.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
Xbox One官方立体声耳机</p>
<p align="center"><a href="/img/20140212/7c541a54a80747b68eae47a2134bc02f.jpg" target="_blank"><img alt="紧咬PS4：Xbox One官方立体声耳机亮相" src="/img/20140212/s_7c541a54a80747b68eae47a2134bc02f.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
Xbox One耳机所使用的新型扁平接口</p>
<p align="center"><a href="/img/20140212/9636d2f8cbcd41d5aeb50dac096a2db3.jpg" target="_blank"><img alt="紧咬PS4：Xbox One官方立体声耳机亮相" src="/img/20140212/s_9636d2f8cbcd41d5aeb50dac096a2db3.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
PS4官方头戴无线耳机</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292392.htm" target="_blank" title="AMD还有俩“新”卡：这么闹腾挺无语的"><span class="title1">AMD还有俩“新”卡：这么闹腾挺无语的</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:44:23</span></td>
  </tr>
  <tr>
    <td align="left"><p><a class="f14_link" href="http://news.mydrivers.com/1/292/292154.htm" target="_blank">Radeon HD 7770 GHz改名而来的Radeon R7 250X</a>正式发布之后，AMD R200系列仍将不断有新成员加入，但其实并没有什么新意。</p>
<p>VR-Zone了解到，<strong>AMD首先会在2000元左右的性能级市场上新推&ldquo;Radeon R9 280&rdquo;</strong>，定位介于R9 280X、R9 270X之间，核心代号为&ldquo;<strong>Tahiti Pro 2</strong>&rdquo;。</p>
<p>也就是说，<strong>它将是Tahiti XT R9 280X的精简版本</strong>，按惯例流处理器会更少，频率也会更低，<strong>同时它又是Radeon HD 7950新的身份</strong>，可能会有频率上的不同，甚至完全相同。</p>
<p><strong>其次是主流领域的&ldquo;Radeon R7 265&rdquo;</strong>，但是核心既不同于更高的R9 270 Curacao，也不同于更低的R7 260X Bonaire，而是<strong>直接挪用原来的Pitcairn Pro，换句话说就是Radeon HD 7850的马甲。</strong></p>
<p>这两款卡的主要任务就是以最低的代价进一步丰富产品线，并让老核心继续发挥余热。</p>
<p>至于具体规格、何时发布，目前还没有任何消息，可能会和R7 250X一样突然低调登场，毕竟都没什么只得炫耀的地方。</p>
<p align="center"><a href="/img/20140212/0a5a9f91b5fe4eb5a8df198a24d4a201.jpg" target="_blank"><img alt="AMD还有俩“新”卡：这么闹腾挺无语的" src="/img/20140212/s_0a5a9f91b5fe4eb5a8df198a24d4a201.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p>&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292391.htm" target="_blank" title="OPPO Find 7新变种：1080p版首曝光"><span class="title1">OPPO Find 7新变种：1080p版首曝光</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:39:24</span></td>
  </tr>
  <tr>
    <td align="left"><p>OPPO即将推出新一代旗舰手机OPPO Find 7，不过除了2K屏版本，该公司似乎还有别的准备。</p>
<p>日前，<strong>又有一款Find 7出现在跑分测试网站GFXbench上面，其型号为X9007，令人意外的是，该机采用的屏幕并非2560&times;1440，而是1920&times;1080</strong>。</p>
<p>OPPO官方早就宣布过Find 7将采用2K级别的屏幕，因此2K版绝对会亮相，而这款屏幕缩水的X9007很大可能是Find 7的变种版本。</p>
<p>从测试信息看到，它配备MSM8974处理器。而从Manhattan Offscreen以及1080p T-Rex Offscreen等项目的得分表现来看，<strong>X9007采用的应该是与2K版本相同的MSM8974AB四核</strong>。</p>
<p>除了屏幕和处理器，1080p的Find 7其他配置暂时还未公布出来。最新消息称<a class="f14_link" href="http://news.mydrivers.com/1/292/292320.htm" target="_blank">OPPO将于3月19日在北京推出Find 7</a>，相信这款1080p的版本届时也会亮相。</p>
<p align="center"><a href="/img/20140212/f0824aa39eeb44ed952cdaa788841b6e.jpg" target="_blank"><img alt="OPPO Find 7新变种：1080p版首曝光" src="/img/20140212/s_f0824aa39eeb44ed952cdaa788841b6e.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/c9e58a3bddb24e4baf3d5e67fe481b46.jpg" target="_blank"><img alt="OPPO Find 7新变种：1080p版首曝光" src="/img/20140212/s_c9e58a3bddb24e4baf3d5e67fe481b46.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292390.htm" target="_blank" title="Win8.1靠边站：来个Windows XP SP4如何？"><span class="title1">Win8.1靠边站：来个Windows XP SP4如何？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:38:32</span></td>
  </tr>
  <tr>
    <td align="left"><p>Windows XP生命周期终止于4月8日，SP3是其最后一个补丁包。不少人猜测，也许Modern版的Windows XP能比Windows 8、8.1获得更大的成功，毕竟这款老系统满足了用户的基本需求，包括简洁的界面、开始菜单、稳定、支持旧硬件。</p>
<p>想到此处，国外网友<a class="f14_link" href="http://p0isonparadise.deviantart.com/art/Concept-Windows-XP-SP4-432449474" target="_blank">p0isonParadise</a>马上起身设计了一张Windows XP SP4概念图，在原本的Windows XP界面中添加了Modern元素，实现扁平化视觉效果，无Aero，运行IE11和Windows Media Player 12。</p>
<p align="center"><a href="/img/20140212/4112b6a2edf5435487b9d10ac1e5a541.png" target="_blank"><img alt="Windows XP SP4概念设计" src="/img/20140212/s_4112b6a2edf5435487b9d10ac1e5a541.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/46c9cd8e9aa24f68b7267d67a8cb6457.png" target="_blank"><img alt="Windows XP SP4概念设计" src="/img/20140212/s_46c9cd8e9aa24f68b7267d67a8cb6457.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center">&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292388.htm" target="_blank" title="AIDA64 4.20正式版发布：硬件神器新跨越"><span class="title1">AIDA64 4.20正式版发布：硬件神器新跨越</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:31:24</span></td>
  </tr>
  <tr>
    <td align="left"><p>全球No.1的硬件识别、诊断工具AIDA64今天发布了<strong>最新正式版4.20.2800</strong>，更新之处非常多，也都颇具重量级。强烈推荐！</p>
<p>顺便一提，AIDA64正式版的版本号都是跳跃性的，最近几个分别是3.00、3.20、4.00、4.20。</p>
<p><strong>AIDA64 4.20正式版更新日志：</strong></p>
<p>－ 支持<strong>AMD Mantle API</strong></p>
<p align="center"><strong><a href="/img/20140212/14414c85adc046bd9a879729982fe018.jpg" target="_blank"><img alt="AIDA64 4.20正式版发布：硬件神器新跨越" src="/img/20140212/s_14414c85adc046bd9a879729982fe018.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></strong></p>
<p align="center"><strong><a href="/img/20140212/bdd139ef17bb401d911e8fbbc7d41093.png" target="_blank"><img alt="AIDA64 4.20正式版发布：硬件神器新跨越" src="/img/20140212/s_bdd139ef17bb401d911e8fbbc7d41093.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></strong></p>
<p>－ <strong>全新编写的多线程内存压力测试</strong>，负载更重，支持SSE、SSE2、AVX、AVX2、FMA、BMI、BMI2指令集加速，最多支持128个逻辑处理器核心(商业版/工程师版最多640个)</p>
<p>－ 全面支持<strong>AMD下一代低功耗APU Mullins/Beema</strong>，包括CPUID面板、缓存与内存测试、系统稳定性测试、全部相关内存和处理器测试，支持其AVX/SSE指令集，同时支持其整合芯片组Yangtze</p>
<p>－ 全面支持Intel Bay Trail桌面版奔腾/赛扬J、移动版奔腾/赛扬N、平板版Atom Z3000，包括测试中可利用其SSE4、AES-NI指令集</p>
<p>－ 优化Intel Avoton/Rangeley Atom C2000 64位测试</p>
<p align="center"><strong><a href="/img/20140212/ce4d7326c8e14edfa42b622ad239941e.png" target="_blank"><img alt="AIDA64 4.20正式版发布：硬件神器新跨越" src="/img/20140212/s_ce4d7326c8e14edfa42b622ad239941e.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></strong></p>
<p>－ 全面支持即将发布的Intel Haswell Refresh处理器及配套的Wildcat Point 9系列芯片组，包括AVX2/FMA 64位指令集测试</p>
<p>－ 改进支持Intel Broadwell处理器，增加支持AES-NI硬件加速</p>
<p>－ 初步支持Intel Royston SoC低端智能手机处理器</p>
<p>－ 支持OCZ Vertex 460固态硬盘</p>
<p>－ 显卡完全支持AMD Radeon R9、R7、R5，NVIDIA GeForce GTX Titan Black、GTX 750 Ti、GTX 750、GT 710，GeForce 840M、825M，Tesla K40</p>
<p align="center"><strong><a href="/img/20140212/97492cef8d6e429e96e56893f71018f5.png" target="_blank"><img alt="AIDA64 4.20正式版发布：硬件神器新跨越" src="/img/20140212/s_97492cef8d6e429e96e56893f71018f5.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></strong></p>
<p><strong>AIDA64 4.20正式绿色版本地下载：</strong><br />
<a href="http://tools.mydrivers.com/soft/1171.htm">http://tools.mydrivers.com/soft/1171.htm</a></p>
<p align="center"><strong><a href="/img/20140212/9d05e857412f4926b7928fce1262a686.png" target="_blank"><img alt="AIDA64 4.20正式版发布：硬件神器新跨越" src="/img/20140212/s_9d05e857412f4926b7928fce1262a686.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></strong></p>
<p align="center"><a href="/img/20140212/976b7c8101af4bfa9929c4687f3e2e23.jpg" target="_blank"><img alt="AIDA64 4.20正式版发布：硬件神器新跨越" src="/img/20140212/s_976b7c8101af4bfa9929c4687f3e2e23.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/301348d7e4dd415c898d9d4bce73f625.jpg" target="_blank"><img alt="AIDA64 4.20正式版发布：硬件神器新跨越" src="/img/20140212/s_301348d7e4dd415c898d9d4bce73f625.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><strong><a href="/img/20140212/68a81ffe141f4854a1629c16eeec5176.png" target="_blank"><img alt="AIDA64 4.20正式版发布：硬件神器新跨越" src="/img/20140212/s_68a81ffe141f4854a1629c16eeec5176.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></strong></p>
<p align="center"><strong><a href="/img/20140212/6b0c0420f03c45c48b9251f4e9ba5662.png" target="_blank"><img alt="AIDA64 4.20正式版发布：硬件神器新跨越" src="/img/20140212/s_6b0c0420f03c45c48b9251f4e9ba5662.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></strong></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292387.htm" target="_blank" title="这就是HTC One 2！"><span class="title1">这就是HTC One 2！</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:30:51</span></td>
  </tr>
  <tr>
    <td align="left"><p>尽管目前HTC M8（HTC One 2）的发布日期还不清楚，但可以肯定的是它已经越来越近了。在<a class="f14_link" href="http://news.mydrivers.com/1/292/292008.htm" target="_blank">真机谍照曝光</a>之后，又有文件显示<a class="f14_link" href="http://news.mydrivers.com/1/292/292250.htm" target="_blank">该机已经通过了蓝牙联盟的认证</a>。现在，国外知名论坛Xda-developers上已经有网友放出了HTC M8的官方渲染图和壁纸。</p>
<p>从渲染图来看，<strong>HTC M8首次使用了虚拟按键设计（<a class="f14_link" href="http://news.mydrivers.com/1/291/291462.htm" target="_blank">正如我们之前的报道</a>），正面没有任何实体按键（HTC One的触摸按键也被取消），看起来更加简洁，保留了HTC的Logo</strong>，目前还不清楚该Logo是否可充当Home键使用。</p>
<p><strong>手机背部采用了双镜头设计，同时两粒LED闪灯的颜色也略有差异</strong>，有些类似于iPhone 5s。根据外媒的推测，摄像头像素可能会提升到500万（依然是Ultra pixel镜头），而且极有可能硬件上支持&ldquo;先拍照、后对焦&rdquo;。另外，双LED闪光灯左上方有一个小孔，估计是麦克风，可以提供更出色的视频录制音质效果。</p>
<p>HTC One左上方的传感器被改到了前置镜头旁边，采用上下排列的布局。<strong>手机采用超窄边框（据称边框也会采用金属材质），整体看起来更加修长</strong>，可以在有限体积下放入尺寸放大的屏幕。</p>
<p>此外，一同泄露的还有HTC M8的壁纸。根据爆料者的说法，该壁纸原始尺寸为2160&times;1080，达到了2K标准。但实际上，该壁纸为滚动壁纸，<strong>HTC M8实际上还是用的1080p屏幕</strong>。</p>
<p>至于处理器目前尚不确定，估计是高通的骁龙805处理器，当然也不排除8974AB的可能。</p>
<p align="center"><a href="/img/20140212/68230e9882494b9390a06bfec0f6df72.jpg" target="_blank"><img alt="这就是HTC One 2！" src="/img/20140212/s_68230e9882494b9390a06bfec0f6df72.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
HTC M8官方渲染图</p>
<p align="center"><a href="/img/20140212/024a63bc73d34878bc1bfa20a182d41a.jpg" target="_blank"><img alt="这就是HTC One 2！" src="/img/20140212/s_024a63bc73d34878bc1bfa20a182d41a.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/7a0a806b13f44578ac259aee98673f15.jpg" target="_blank"><img alt="这就是HTC One 2！" src="/img/20140212/s_7a0a806b13f44578ac259aee98673f15.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
HTC M8壁纸</p>
<p align="center"><a href="/img/20140212/ce61124c51ec42b4bfa3376b8b04c71e.jpg" target="_blank"><img alt="这就是HTC One 2！" src="/img/20140212/s_ce61124c51ec42b4bfa3376b8b04c71e.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
HTC M8中文界面，其中的&ldquo;立即ID&rdquo;不知道是否和指纹识别有关系<br />
<br />
&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292386.htm" target="_blank" title="滴滴打车：我们已经补贴4个亿了"><span class="title1">滴滴打车：我们已经补贴4个亿了</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:29:10</span></td>
  </tr>
  <tr>
    <td align="left"><p>嘀嘀打车对外公布接入微信支付的总成绩单，包括嘀嘀打车自身运营状况，详细情况为：</p>
<p>1、1月10日至2月9日，<strong>嘀嘀打车中平均日微信支付订单数为70万单，微信支付订单总数约为2100万单</strong>。</p>
<p>2、嘀嘀打车开通服务的58个城市均有成功使用微信支付来支付打车费的记录，其中33个城市日均微信支付订单超过1万单。</p>
<p>3、嘀嘀打车过去30天日均订单为183万单，2月7日节后第一天达到峰值262万单，微信支付订单峰值过200万单。</p>
<p>4、<strong>嘀嘀打车和腾讯总计为微信支付补贴4亿元</strong>。</p>
<p>目前<a class="f14_link" href="http://news.mydrivers.com/1/292/292179.htm" target="_blank">嘀嘀打车已经将补贴额缩至每单5元</a>，即：每单乘客奖励5元，每天3笔；每单司机奖励5元，每天最多5次。</p>
<p align="center"><img alt="滴滴打车：我们已经补贴4个亿了" src="/img/20140212/23a87ddcb21b4a9f9d582f221fc40795.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292385.htm" target="_blank" title="谷歌贴心新功能：餐厅菜单也能搜的到"><span class="title1">谷歌贴心新功能：餐厅菜单也能搜的到</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:25:01</span></td>
  </tr>
  <tr>
    <td align="left"><p>谷歌搜索引擎已推出一项新功能，<strong>当用户专门搜索菜单信息时，谷歌搜索引擎将显示完整的餐厅菜单。</strong></p>
<p>一名名为艾丽&middot;布朗(Allie Brown)的用户最先发现了这一新功能，她将关键词&ldquo;Jones Brunch Menu&rdquo;的搜索结果公布在了Twitter之上。对于这一关键词的搜索，谷歌提供了一个卡片式答案，显示了这家餐厅多种类型食品的菜单，例如餐前小菜、主菜和三明治等。</p>
<p>谷歌尚未就这一最新功能测试置评。业内人士指出，谷歌搜索结果中的菜单信息来自AllMenus.com网站。<strong>有一点可以确定，这一信息并非来自该餐厅官方网站。</strong>官方网站上的菜单为一段Flash动画，此外也没有提供谷歌给出的价格信息。</p>
<p>业内人士指出，<strong>谷歌此举可能对餐厅业主造成不利影响</strong>，因为用户将没有必要再前往餐厅官方网站去了解信息。</p>
<p>不过近期的一项研究也表明，美国的独立餐厅中只有不到一半建设了自主网站，其中也只有40%在网上提供了菜单。在许多情况下，这些网上菜单也是过时或不完整的。考虑到有研究表明，80%消费者希望在去餐厅用餐前查看菜单，谷歌推出这一功能也是合理的。</p>
<p align="center"><img alt="谷歌贴心新功能：餐厅菜单也能搜的到" src="/img/20140212/abf506c3897c489aa89da0cab60b2533.png" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292384.htm" target="_blank" title="组图：Windows Phone 8.1抢先看"><span class="title1">组图：Windows Phone 8.1抢先看</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:22:56</span></td>
  </tr>
  <tr>
    <td align="left"><p>近几日来，Windows Phone 8.1的新功能陆续曝光，例如通知中心、更大的磁贴尺寸、类Swype输入法。Reddit论坛上出现了Windows Phone 8.1的多张图片，这些图片来自微软发布的SDK。从中我们可以看到，<strong>Windows Phone 8.1带来了大量全新的设置选项和功能</strong>。</p>
<p><strong>Windows Phone 8.1提供了存储管理、屏幕投影设置、更多的定制选项</strong>，预计会在4月份的BUILD微软开发者大会上发布。尽管从这些图片中我们没能看到大幅的界面调整，不过一些细节之处确实是有变化的，例如全新的虚拟导航按钮。</p>
<p align="center"><a href="/img/20140212/561245a80f504d70acc6251b550e8f39.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_561245a80f504d70acc6251b550e8f39.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/76744d0bf3a84fc08704fcb0d6e09970.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_76744d0bf3a84fc08704fcb0d6e09970.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/0fe0de77442449dfb2abf89adb3d0098.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_0fe0de77442449dfb2abf89adb3d0098.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/6a0e908e1fca4e7fa9dd49c532310f0f.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_6a0e908e1fca4e7fa9dd49c532310f0f.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/83906ffe39304a938c48a393028c1f2a.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_83906ffe39304a938c48a393028c1f2a.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/4f8e3b9237194644abf89f0644df86b7.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_4f8e3b9237194644abf89f0644df86b7.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/19569aa400a04e0cabeec584ee6b953b.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_19569aa400a04e0cabeec584ee6b953b.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/8feb1d3d527349d399eeacf3ef4b2bc6.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_8feb1d3d527349d399eeacf3ef4b2bc6.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/a7e5ecdc16514afe9cbaef5e3a08e2c2.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_a7e5ecdc16514afe9cbaef5e3a08e2c2.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/5c8e78ccaeea46498a8b62da48cddb93.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_5c8e78ccaeea46498a8b62da48cddb93.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/276e08f2882d4a5594d20e00bad7eb1f.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_276e08f2882d4a5594d20e00bad7eb1f.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/11b1ddf3741843bebe18056f7b530906.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_11b1ddf3741843bebe18056f7b530906.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/aacafa58a054492c9a119e2d0f0b145e.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_aacafa58a054492c9a119e2d0f0b145e.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/868dc3bbc88e498b9c5ead63be4e8c3b.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_868dc3bbc88e498b9c5ead63be4e8c3b.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/4a6230ef51194af7af403fde2abf0ed4.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_4a6230ef51194af7af403fde2abf0ed4.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/5a3e5bfab37d4ceb8d8ffe019cf3db79.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_5a3e5bfab37d4ceb8d8ffe019cf3db79.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/f22e0cfacf39468dadc888fbbcdc2a72.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_f22e0cfacf39468dadc888fbbcdc2a72.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/2782fac331f749ae85601cee20a39737.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_2782fac331f749ae85601cee20a39737.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/70708413a82a4ebca0a8826d0ba7db62.png" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_70708413a82a4ebca0a8826d0ba7db62.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/789627372a624d498ff8b4264b92ef74.png" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_789627372a624d498ff8b4264b92ef74.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/5558d2620644445e86e9ab13da953b1d.jpg" target="_blank"><img alt="组图：Windows Phone 8.1抢先看" src="/img/20140212/s_5558d2620644445e86e9ab13da953b1d.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p>&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292383.htm" target="_blank" title="Flappy Bird会永远消失吗？"><span class="title1">Flappy Bird会永远消失吗？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">刘艺</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:22:01</span></td>
  </tr>
  <tr>
    <td align="left"><p>Flappy Bird开发者近日声称该游戏的走红让他感到很有压力，因而将其下架。科技网站Re/Code今天发布文章称，Flappy Bird开发者Dong Nguyen想着将该款爆红的游戏下架苹果App Store和谷歌Google Play，<strong>就能让它&ldquo;永远消失&rdquo;，但事与愿违。</strong></p>
<p>下架游戏或许能让他恢复原来的简单生活<strong>，但不管怎么样，他都无法让Flappy Bird&ldquo;永远消失&rdquo;。</strong>Flappy Bird依然活跃于数百万的手机和平板电脑上，依然有大量玩家在玩那款永远都不会再更新的游戏应用。</p>
<p>Dong Nguyen在接受《福布斯》采访时明确了一点：<strong>他的游戏不会复活重新上架</strong>。上周The Verge报道称，他一天可从Flappy Bird赚得5万美元。</p>
<p>数字稀缺性的概念通常只会在人们出于保护未来利润的目的拒绝将内容开放给合法的网络渠道的时候出现。你最喜欢的乐队的作品不在iTunes上？你最喜欢的电视节目不在Netflix上？一旦适当的协议大胜，这些都有可能发生改变。</p>
<p><strong>而人们对Flappy Bird的兴趣却仍呈加速增长之势，或许是因为Dong Nguyen催生了一种全新的数字稀缺性。</strong>通过将该应用撤出应用商店，Dong Nguyen使得拥有他的游戏的手机成了专有俱乐部的一部分，甚至是eBay上的出售商品。如今就剩下原来装有Flappy Bird的人才有那款游戏，这在如今什么都有的世界里着实奇怪。</p>
<p>当然，<strong>山寨开发者们也纷纷去填补Flappy Bird留下的空缺，但这并无意义</strong>。即便他们能够匹敌甚至胜过Dong Nguyen的简单游戏，他们也无法在狂热程度上达到Flappy Bird的高度。即使不再提供，Flappy Bird也依然是人们在社交网络上的热议话题，有着很旺的人气。</p>
<p>讽刺的是，在Flappy Bird进入鼎盛阶段之时将其扼杀，使得它的退出带有标志性色彩，令人忘怀，而几乎所有其它的游戏&mdash;&mdash;其中一部分花了无数的人力和时间去制作&mdash;&mdash;却都逐渐消失在人们的视野当中。</p>
<p align="center"><img alt="Flappy Bird会永远消失吗？" src="/img/20140212/7aeac7c69c8f4e9487af64430aa448ce.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292382.htm" target="_blank" title="微信沃卡进军广西"><span class="title1">微信沃卡进军广西</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:20:17</span></td>
  </tr>
  <tr>
    <td align="left"><p>由广东联通和腾讯打造的微信沃卡在去年下半年曾名噪一时，<strong>目前该业务合作将进一步扩军，广西联通已开始与腾讯洽谈合作。</strong></p>
<p>据了解，去年由广东联通掀起的微信沃卡曾风靡一时，成为业界焦点，当时，腾讯等OTT业务商不断与电信运营商显现不和谐，运营商的语音通话、短信、彩信业务的下滑都被认为受微信等OTT业务蚕食影响，但是，去年7月，广东联通与腾讯牵手，推出了微信沃卡，成为国内运营商与OTT深度合作的破冰者。</p>
<p><strong>微信沃卡它不仅是一张电话卡，也不局限于微信定向流量优惠，它还配备了群组、表情、支付和游戏特权。</strong>让微信用户过足了瘾，因此也受到用户热烈欢迎，为广东联通带来了大量用户。</p>
<p>不过，广东联通推出微信沃卡后迄今已过去半年，一直没有见到其他运营商乃至中国联通其它省公司与微信合作的身影，表明三大运营商仍对微信有不同看法。</p>
<p>种种迹象显示，广西联通与广东联通、上海联通、浙江联通、贵州联通等一起，是中国联通在南方改革试水的&ldquo;试验田&rdquo;，例如2012年广西联通推出沃易得,创立与银行合作的新模式；2013年，建立的&ldquo;沃易购&rdquo;平台，也开辟了运营商主动放弃代销利益、建立手机企业直销平台等新模式。这次，与腾讯微信联手，是转型互联网模式又一次试水。</p>
<p align="center"><img alt="微信沃卡进军广西" src="/img/20140212/48dc491b14294e2eab73a44531e568d2.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292381.htm" target="_blank" title="索尼：PS4接班PS3？它还太年轻"><span class="title1">索尼：PS4接班PS3？它还太年轻</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小路</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:11:29</span></td>
  </tr>
  <tr>
    <td align="left"><p>索尼的新一代游戏机PS4已经发售近3个月时间，该主机目前在全球销量上一扭PS3时代的颓势而力压对手微软的产品占据主导，因此有声音表示是时候将接力棒交给PS4，让PS3回家养老了。</p>
<p>对此，索尼PlayStation营销主管John Koller在接受外媒GameSpot采访时作出了回应，他表示：<strong>尽管PS4的开局非常不错，但其毕竟还是个&ldquo;新手&rdquo;，因此我们至少还会让PS3再坚持个3、4年时间，期间索尼将继续为老平台用户提供包括技术和游戏上支持，但另一方面工作团队也会更加努力的为PS4带来更加精彩内容，好让这个毛头小伙早日成熟。</strong></p>
<p>同时，<strong>John Koller还表示目前官方暂无计划推出PS Vita与PS4的同捆套装</strong>，原因在于索尼并不想让PS Vita彻底沦为PS4的&ldquo;配件&rdquo;，而是继续让它以独立设备的姿态来进行呈现，今后不论是PS4玩家还是PS Vita玩家都会获得更多只属于自己的独占游戏作品。近日，索尼先后宣布了最新PS Vita 2000（PS Vita Slim）登陆英国和美国的消息。</p>
<p>稍早索尼官方曾透露将在<strong>今年为PS4玩家呈上100款内容各异的游戏作品</strong>，而在今年的台北电玩展（TpGS）上，索尼则一口气公布了包括《<a class="f14_link" href="http://news.mydrivers.com/1/291/291406.htm" target="_blank">闪之轨迹1、2</a>》、《<a class="f14_link" href="http://news.mydrivers.com/1/291/291246.htm" target="_blank">刀剑神域：虚空断章</a>》在内的数款中文化PS Vita游戏，意在大力开拓中国地区玩家市场。</p>
<p align="center"><a href="/img/20140212/7925aba5b9ef4b6ab2978da32d37f291.jpg" target="_blank"><img alt="PS4何时能接过PS3手中的接力棒？" src="/img/20140212/s_7925aba5b9ef4b6ab2978da32d37f291.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292380.htm" target="_blank" title="航班管家否认遭封杀 第三方值机软件陷授权难题"><span class="title1">航班管家否认遭封杀 第三方值机软件陷授权难题</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:03:29</span></td>
  </tr>
  <tr>
    <td align="left"><p>昨日（2月11日），目前国内最大的航空类APP&ldquo;航班管家&rdquo;发布澄清公告，否认遭航空公司封杀，称其提供的手机值机服务最终依然是通过航空公司官网，用户与航空公司之间的权利与义务不会因为第三方值机而改变。</p>
<p>此前，东航、国航在其官网发布声明称，第三方手机应用软件提供所谓的值机服务并没有得到航空公司的授权，侵犯了其合法权益，不保证使用第三方手机值机服务的旅客能获得与官方渠道手机值机同等的后续服务。这被外界解读为，第三方手机自助值机或正遭遇航空公司的集体抵制与封杀。</p>
<p>有民航业内人士认为可能是第三方平台触及了相关方的利益。某第三方软件负责人对《每日经济新闻》记者表示，航空公司的声明可能更多解读为对旅客的善意提醒，因为携程旅行网以及去哪儿网等应该与航空公司都有相关的合作协议。</p>
<p>不过，有业内人士认为，航空公司抵制的背后，在于其经营理念影响，第三方手机自助值机直接冲击了旅客对航空公司线上的官方认知度以及忠诚度，进而影响航空公司的线上营销。未来双方的授权以及深度合作方式或都将围绕此展开。</p>
<p><strong>&ldquo;航班管家&rdquo;否认遭封杀</strong></p>
<p>继东航1月23日发布《关于第三方提供东航值机服务的声明》之后，国航2月10日也在其官网发布《关于第三方提供国航值机服务的声明》，二者均指第三方手机应用提供的东航、国航自助值机服务，未得到二者的授权，并指第三方手机自助值机未经企业许可，使用了企业的商标、企业名称、企业标识等，侵犯了企业的合法权益，希望有关第三方立即停止侵权行为。</p>
<p>据媒体日前报道，虽然南航尚未发布类似声明，但南航客服人员建议旅客最好通过南航官网和南航手机客户端办理网上值机。</p>
<p>海南航空相关人士昨日也对《每日经济新闻》记者表示，&ldquo;我们建议旅客最好通过海南航空官方网站、移动客户端、微信服务号等官方渠道办理网上值机等服务，这些官方渠道能为广大旅客提供更多贴心便捷、符合民航局安全管理规定的服务，最大程度的保障航空安全及保护广大旅客的权益。&rdquo;</p>
<p>记者了解到，目前中国民航信息网络股份有限公司开发的 &ldquo;航旅纵横&rdquo;，携程旅行网的&ldquo;携程旅行&rdquo;以及去哪儿网的 &ldquo;去哪儿旅行&rdquo;等较为权威的第三方软件平台皆提供手机值机服务。</p>
<p>对于航空公司的表态，航班管家回应称，已经与国航、深航等达成紧密合作，并且正在与东航、海航、南航等航空公司进行密切交流，准备进一步加深合作。不过就同国航、深航已达成的具体何种合作方式，是否已经授权，仍未做相关说明。</p>
<p>值得注意的是，东航、国航在声明中均指&ldquo;第三方提供的手机自助值机服务违反民航局关于在自助值机渠道告知危险品托运要求的规定，影响航空安全&rdquo;。</p>
<p>不过，航班管家的回应中对此并未提及。记者昨日手机登录航班管家选择国航值机，其页面已经提示，&ldquo;接下来的值机服务将由国航官方网站提供&rdquo;，随后出现&ldquo;危险品须知&rdquo;以及&ldquo;国航手机乘机等级服务旅客须知&rdquo;的提示。记者又选择其他多家航空公司进行值机，在正式值机流程之前，航班管家也出现了&ldquo;危品告知&rdquo;的页面。</p>
<p>记者昨日手机登录&ldquo;航旅纵横&rdquo;以及&ldquo;携程&rdquo;APP办理值机，仍未有明确的危险品告知提示。携程只提示&ldquo;点击&lsquo;办理值机&rsquo;，代表您已阅读并同意&lsquo;值机协议&rsquo;和&lsquo;值机须知&rsquo;&rdquo;。</p>
<p>去哪儿网相关人士昨日晚间给《每日经济新闻》记者发来回应称，去哪儿网与航空公司始终保持着密切的合作，去哪儿网无线客户端的在线值机功能，与航空公司官网数据同步，系统十分稳定并且安全。去哪儿旅行客户端自去年11月以来增加手机值机功能以来，暂时还没有接到过&ldquo;到了机场后无法正常值机&rdquo;的用户反馈。</p>
<p><strong>第三方手机值机授权待解</strong></p>
<p>事实上，在业内人士看来，&ldquo;危险品&rdquo;告知提示通过技术改进，可以轻而易举实现。航空公司抵制第三方值机软件的背后，重点还在于其经营理念受到影响。</p>
<p>业内人士对 《每日经济新闻》记者分析，正如东航、国航在其声明中所强调，希望广大旅客登录公司门户网站、手机网站和手机客户端等官方渠道办理自助值机手续，因为其官方渠道往往包含了多种营销策划和推广，而第三方提供的手机值机的便捷性，使得旅客跳开航空公司的官方渠道，直接影响旅客对航空公司的线上认知度以及忠诚度。</p>
<p>但毋庸置疑的是，目前市场上的第三方手机值机服务利用移动互联网技术，可以整合各大航空公司的值机入口，解决用户以往需要登录十数家航空公司网站和手机客户端进行多种体验操作的痛苦，已经产生了巨大用户需求，被认为是大势所趋。</p>
<p>航班管家昨日的数据显示，2013年9月中旬推出的支持十多家航空公司的手机值机选座服务，推出2周内，日值机量已经接近国内在线值机总量的10%；2014年春运期间，日值机量翻倍。</p>
<p>对于未来航空公司与第三方值机软件之间的 &ldquo;授权&rdquo;如何进行，业内人士认为，或将加入围绕航空公司线上营销的各种 &ldquo;加深合作&rdquo;的方式，甚至不排除相关费用的产生，最主要的是第三方值机软件的&ldquo;资质&rdquo;要得到航空公司的认可。</p>
<p align="center"><img alt="航班管家否认遭封杀 第三方值机软件陷授权难题" src="/img/20140212/5f1318c0813f4708bea82a0d3a648ab5.png" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292379.htm" target="_blank" title="小米新“玩具”曝光"><span class="title1">小米新“玩具”曝光</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">09:02:34</span></td>
  </tr>
  <tr>
    <td align="left"><p>小米在马年的第一份产品是什么呢？昨天晚上小米社区运营经理@小米浩子同学在微博上给我们揭晓了答案。</p>
<p>按照他的说法，<strong>小米即将推出一款NFC标签碰碰贴</strong>，号称极限超薄，可以贴在任意平面甚至曲面上使用。NFC线圈全铝制，中间搭配一个金属&ldquo;MI&rdquo;logo，通过搭配的APP可以快速切换手机模式、并设置快捷方式等等。当然，<strong>前提是你的手机必须支持NFC，比如小米手机2A和小米手机3</strong>。</p>
<p>事实上这个功能并不新鲜，此前多款支持NFC的国产手机包装盒中均附赠有这一配件，功能上可能没有小米做的那么丰富，但主要功能都实现了。</p>
<p align="center"><a href="/img/20140212/e27eadc14e9f45acb1d3a7b9e9ca3984.jpg" target="_blank"><img alt="小米新“玩具”曝光" src="/img/20140212/s_e27eadc14e9f45acb1d3a7b9e9ca3984.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/e50d147aac374d5dba53333726041f63.jpg" target="_blank"><img alt="小米新“玩具”曝光" src="/img/20140212/s_e50d147aac374d5dba53333726041f63.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/f9d1f3954f504ef488df8676c8df1cdb.jpg" target="_blank"><img alt="小米新“玩具”曝光" src="/img/20140212/s_f9d1f3954f504ef488df8676c8df1cdb.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/334c587e395e4147a53086b3a1fc3afd.jpg" target="_blank"><img alt="小米新“玩具”曝光" src="/img/20140212/s_334c587e395e4147a53086b3a1fc3afd.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/8a1b7ed42c824ae8808cf8b638f78b0c.jpg" target="_blank"><img alt="小米新“玩具”曝光" src="/img/20140212/s_8a1b7ed42c824ae8808cf8b638f78b0c.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/6b060dba97a2420d94a995d3d11f1161.jpg" target="_blank"><img alt="小米新“玩具”曝光" src="/img/20140212/s_6b060dba97a2420d94a995d3d11f1161.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/05631fe998ce495c9376f8d5c33a1585.jpg" target="_blank"><img alt="小米新“玩具”曝光" src="/img/20140212/s_05631fe998ce495c9376f8d5c33a1585.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292378.htm" target="_blank" title="冯氏春晚其实是被微博吐槽扶起来的"><span class="title1">冯氏春晚其实是被微博吐槽扶起来的</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">黄栋</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:58:35</span></td>
  </tr>
  <tr>
    <td align="left"><p>人倒了还能扶起来，要微博倒了，春晚可就真扶不起来了！</p>
<p>刚刚过去的春节，这条段子在微博和微信上被大量转发。<strong>即使春晚请来了顶级名导冯小刚，依然没能挽回被吐槽的命运，不可否认的是，看春晚最大的乐趣已经变成在新浪微博看大家吐槽春晚。</strong></p>
<p>面对可预料的欢乐吐槽，新浪微博却早早宣布成为独家二维码互动合作伙伴，推出独特的玩法。早就听说新浪微博要狠搞电视，从与春晚的合作效果来看，微博成功打造出一个合作案例，对电视台和广告主们具备足够的吸引力。</p>
<p><strong>强大的热点汇聚效应</strong></p>
<p>大年三十的晚上，即使春晚再无聊，绝大多数的国人还是会跟春晚扯上关系，毕竟这是三十年来养成的家庭式习惯。</p>
<p><strong>这个时候，微博的价值就体现出来了，各种热闹欢乐全冒出来，要比春晚好看很多，而这种赶脚在微信永远找不到，即使朋友圈里你看到的段子，百分百来自微博。</strong>就连冯小刚也在春晚开场短片里来了一句，&ldquo;春晚最大的乐趣是什么，吐槽啊！&rdquo;</p>
<p>实际上，在微博围观和吐槽春晚早已成为网友除夕夜的一项欢乐的活动。欢乐的活动总会吸引大量网友的参与，官方统计显示，直播期间，3,447万微博用户参与春晚互动，互动量(原创、转发、评论、赞)高达6,895万，春晚的微博提及量最终达到4,541万条，同比去年1900万增长1.39倍。</p>
<p>武汉大学沈阳教授团队观察到，<strong>凡是有微博的微信用户，年三十晚上很多都已经回流到微博当中来进行转发和表态，没有微博的微信用户还是在微信中进行交流，</strong>这或者意味着微博对新增用户的开拓还有比较大的空间。</p>
<p>该团队认为，微博和微信正在出现一种双项内容的筛选和排挤机制，也就是说，<strong>微信会把公共传播性质的这种内容自动筛选排挤到微博当中去，用户在使用微信时，更多地去进行私密信息的交流。除此之外微信具有一种意见隔离的机制，没有办法形成大范围的信息交流。</strong>从这个角度而言，微博仍然有其非常强大的生命力。</p>
<p>由此再次证明，在重大热点事件中，微博依然是网友获悉信息、参与讨论的最便捷平台，热点汇聚效应非常明显。即使强大如央视春晚，都通过与微博合作提升了曝光量乃至收视率，其他电视台与电视节目自然还有很多可以发挥的空间。</p>
<p><strong>社会化营销的深层次思考</strong></p>
<p>当然，新浪微博舍得投入重金成为独家二维码互动合作伙伴，并不只是要简单证明平台的热点汇聚效应，自然还会有社会化营销的深层次思考。</p>
<p>1月6日，新浪微博连续四年发起&ldquo;让红包飞&rdquo;活动。与往年相比，今年的活动邀请了更多企业参与，除了恒大这样的土豪，还有为数不少的中小企业共同造势。在拉来春晚这样的重磅合作伙伴后，今年的让红包飞无疑具备了更强的吸引力。</p>
<p>根据微博春晚的合作约定，网友参与春晚话题讨论或者直接扫描直播页面的二维码就能参与抽奖，在开奖时将看到是谁提供的奖品。对企业而言，这种通过奖品刺激所形成品牌曝光带给网友的印象，要比单纯电视广告投放所形成印象更为深刻，品牌曝光与粉丝互动同时进行，更为重要的是，并不是每家企业都有在春晚投放广告的实力。</p>
<p>官方数据显示，<strong>让红包飞#春晚#专场累计送出2,300万个红包，参与企业粉丝累计增长950万，其中5家合作方粉丝翻倍增长，@俏十岁 @恒大冰泉 涨粉量均在300%以上。</strong></p>
<p>微博与春晚联手打造的台网联动模式，给合作方带来了更多品牌曝光机会和社会化资产增值的机会。可以说，<strong>微博已经成为重要的营销工具和平台，使原本不能在电视媒体、平面媒体讲述的品牌故事，互动机会成为可能，也使企业把微博当成了互动营销的大舞台。</strong></p>
<p>有消息透露，为了成为央视独家二维码合作伙伴，新浪微博投入了几百万，但是&ldquo;让红包飞&rdquo;的广告已经卖出几千万，即使与央视对半分成，收入规模至少达到了千万元。无论对于新浪微博还是央视，这都是个不错的战果，也为下一次合作提供了更大想象空间。</p>
<p>去年12月，新浪微博与央视索福瑞联手推出微博收视指数，新浪微博总经理王高飞喊出宣言，&ldquo;互联网公司都是想让更多人不看电视，只有微博希望更多人看电视！&rdquo;年底的冯氏春晚，新浪微博则同时狠赚人气和财气。1个月之内，新浪微博在电视行业先后点燃两把火，相信不用等到巴西世界杯，第三把火很快就会烧起来，而在做世界杯计划时，必然会有更多的电视台和广告主们重新思考微博的价值和意义。</p>
<p align="center"><a href="/img/20140212/88f46e3d4ee941459764f8e9a0a07aff.jpg" target="_blank"><img alt="冯氏春晚其实是被微博吐槽扶起来的" src="/img/20140212/s_88f46e3d4ee941459764f8e9a0a07aff.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/d310fd3429b94940b80ded0760e4dd2d.jpg" target="_blank"><img alt="冯氏春晚其实是被微博吐槽扶起来的" src="/img/20140212/s_d310fd3429b94940b80ded0760e4dd2d.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/624d3e89065e4e6b8ad1b0b1b1475ff9.jpg" target="_blank"><img alt="冯氏春晚其实是被微博吐槽扶起来的" src="/img/20140212/s_624d3e89065e4e6b8ad1b0b1b1475ff9.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292377.htm" target="_blank" title="虚拟运营商最早4月放号：“吉祥号码”成热门"><span class="title1">虚拟运营商最早4月放号：“吉祥号码”成热门</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">小呆</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:58:28</span></td>
  </tr>
  <tr>
    <td align="left"><p>继工信部确定&ldquo;170&rdquo;号段为移动通信转售业务专属号段后，北京商报记者昨日获悉，19家虚拟运营商的客服号码也已确定，<a class="f14_link" href="http://news.mydrivers.com/1/292/292280.htm" target="_blank">沿用基础运营商的&ldquo;100&times;&times;&rdquo;</a>。虚拟运营商准备工作持续推进，业界预计最早将于4月正式放号。</p>
<p>据了解，<strong>工信部确定的电信虚拟运营商客服号码在1002&times;-1004&times;之间，已获得牌照的19家企业主要是采取抽签的形式确定，在此过程中，一些诸如&ldquo;10020&rdquo;、&ldquo;10030&rdquo;、&ldquo;10028&rdquo;等所谓的&ldquo;吉祥号码&rdquo;成为热门</strong>。</p>
<p>虚拟运营商发展研究中心秘书长邹学勇指出，虚拟运营商之所以沿用三大基础运营商的&ldquo;100&times;&times;&rdquo;，主要是因为用户对中国移动10086、中国联通10010、中国电信10000已形成了根深蒂固的认知，沿续这种传统方便用户记忆和使用。另一方面，这也有利于虚拟运营商移动通信转售业务的市场营销及推广。</p>
<p>首批获牌的分享在线执行总裁康志斌表示：&ldquo;分享在线对抽到的客服号码满意，目前与中国电信在移动通信转售业务的合作对接工作接近尾声，放号运营进入倒计时，第一期的试点主要在北京、上海、广东和湖南四个省份开展，到2015年起进行第二期推广。&rdquo;</p>
<p>至于虚拟运营商的号码，除了已知的&ldquo;170&rdquo;号段外，北京商报记者还了解到，<strong>11位手机号的前四位仍是以基础运营商来区分的。&ldquo;1700&rdquo;为中国电信的转售号码标识，&ldquo;1705&rdquo;为中国移动，而&ldquo;1709&rdquo;则为中国联通， 值得注意的是，因虚拟运营商企业数量较多，在号码分配上并没有明显体现各个虚拟运营商的品牌标识。</strong>除了客服号码、手机号段确定，各个虚拟运营商的系统平台也在加紧推进。据一位接近虚拟运营商华翔联信的人士透露，该公司的系统平台近日已基本建好，不少员工一直处于加班状态，华翔联信有望成为最早上线运营转售业务的虚拟运营商。</p>
<p>另一家获牌企业天音通信也表示，该公司的人事招聘工作持续进行中，既包括高管，也涉及技术人员，目前已有多名来自基础运营商的企业加入到天音通信的转售业务部门，客服系统和计费系统已就绪。</p>
<p>不难看出，用户有望在今年二季度享受到虚拟运营商的业务。需要指出的是，在初期阶段，各个企业的运营模式主要是转售，即转售从基础运营商批发来的业务，包括手机卡、无线上网卡以及语音、短信、数据流量套餐等。这在北京商报记者对迪信通、天音通信、爱施德、乐语通讯等多家渠道商的采访中也得到了证实，另一些包括分享在线在内的企业解决方案提供商，转售业务则主要集中在企业级市场上。</p>
<p>邹学勇认为，初期简单的转售业务盈利模式单一，且利润微薄，对虚拟运营商真正的考验在于如何实现与既有业务的整合及互补。在转售之后的第二个阶段里，虚拟运营商企业很有可能会销售补贴定制手机等终端设备，除此之外，各企业更有可能在现有的企业解决方案业务、移动互联网业务等基础上，结合大数据分析，推出更多创新性、细分化的业务及服务。</p>
<p align="center"><a href="/img/20140212/9824e8c3f3874dd6abddec0f2d443a84.jpg" target="_blank"><img alt="虚拟运营商最早4月放号：“吉祥号码”成热门" src="/img/20140212/s_9824e8c3f3874dd6abddec0f2d443a84.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292375.htm" target="_blank" title="索尼没落原因：沉迷过去辉煌 目光短浅"><span class="title1">索尼没落原因：沉迷过去辉煌 目光短浅</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">刘艺</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:55:57</span></td>
  </tr>
  <tr>
    <td align="left"><p>据韩国《朝鲜日报》2月10日消息，凭借特丽珑(Trinitron)和平面显像管等尖端技术一度主导世界电视市场的索尼6日宣布了一项结构调整计划，在7月前拆分电视机部门并裁员5000多人。</p>
<p>索尼近年来一直亏损，信用评级被下调至垃圾级，因此决定整体出售个人电脑业务。据评价，索尼走向没落是因为首席执行官陶醉于过去的成就且只追求短期的成果，不断做出错误决策。</p>
<p><strong>陶醉于过去辉煌而走向没落</strong></p>
<p>索尼1968年开发特丽珑显象管后瞬间颠覆了电视市场的格局。这种技术是一个电子枪射出三支电子束，因此画质比过去的显像管更清晰。索尼利用这项技术超越了显像管电视的始祖美国RCA公司。</p>
<p>索尼1996年开发出平面显像管后，再次撼动世界电视市场。但索尼的成功阻碍了其前进的步伐。三星电子和LG电子等韩国竞争企业纷纷利用液晶电视发出挑战，而索尼依然固守平面显像管电视。直到21世纪索尼才承认显像管电视退出历史舞台并改变发展方向。</p>
<p>由于核心零配件液晶面板的相关研究和投资不到位，索尼只能从三星等竞争企业采购核心零配件。因此，索尼电视部门连续8年出现亏损。</p>
<p>索尼一向认为只要是索尼制造就是世界标准，这种傲慢态度也成为严重问题。索尼曾凭借创新性产品&ldquo;随身听&rdquo;席卷世界市场，但仍无法抵御苹果iPod等MP3播放器的登场。</p>
<p>索尼公司20世纪70年代后期在录像机(VCR)市场上坚持采用自主技术&ldquo;Betamax格式&rdquo;，结果被竞争企业的VHS格式超越。迷你光盘(Mini Disc)和数码相机内存卡也因忽视国际标准而陷入孤立局面。</p>
<p><strong>目光短浅 忘记创业之本</strong></p>
<p>管理者急于追求短期利益，推进事业多元化也是索尼没落的一个原因。1995年就任索尼首席执行官(CEO)的出井伸之宣布要&ldquo;重建索尼&rdquo;，目标是将索尼打造成硬件和软件都很强大的公司，并重点投资了娱乐领域。</p>
<p>他还高呼&ldquo;全球索尼&rdquo;的口号引入美国式外部董事制，将公司业务分割给25个公司。公司体系朝取得短期成果的方向发展后，技术人员大举离职，因为他们觉得&ldquo;索尼不再是技术型公司&rdquo;。</p>
<p>索尼在电视机市场上丧失主导权后，于2005年将并非技术人员出身的传媒、娱乐专家霍华德&bull;斯金格任命为首席执行官。索尼第一位外国首席执行官斯金格不断与技术人员发生矛盾，没有成功推出新的发展蓝图和热门商品。</p>
<p>日本庆应大学教授柳町功指出，由于专业经营者过于追求短期利益，从而丧失了所谓技术开发的索尼创业基因。创始人掌舵的三星、LG等企业反而通过长期研发和重点投资战胜了索尼。</p>
<p><strong>凭借智能手机决一胜负？</strong></p>
<p>2012年就任首席执行官的平井一夫6日表示：&ldquo;今后将重点开发智能手机和平板电脑。&rdquo;</p>
<p>但与苹果和三星电子相比，索尼智能手机业务起步晚，在日本市场上也大幅落后于苹果iPhone。索尼智能手机在世界市场上的占有率仅为3.5%，列第七位。智能手机市场上已经有三星和苹果两大巨头，而且中国企业也在迅猛追击。最重要的问题是索尼一直忙于摆脱亏损，没有针对后智能手机时代做好准备。柳町功表示：&ldquo;智能手机市场的增长势头今后可能会持续放慢，索尼用智能手机决一胜负并不明智。&rdquo;</p>
<p>索尼结构调整过于保守也是一个问题。据评价，最近出售个人电脑业务并非主动性结构调整，而是旨在确保收益的资产转让式结构调整。还有人提出悲观预测称，索尼最终将放弃制造业。索尼2013年4月至12月的结算中，只有金融和音乐领域创造巨大收益，而制造业领域大都亏损。</p>
<p align="center"><img alt="索尼没落原因：沉迷过去辉煌 目光短浅" src="/img/20140212/fa12f7e9804f44508c18ee268884d83a.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292374.htm" target="_blank" title="持股京东待上市 沙特王子个人财富直逼巴菲特"><span class="title1">持股京东待上市 沙特王子个人财富直逼巴菲特</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:55:17</span></td>
  </tr>
  <tr>
    <td align="left"><p>有着&ldquo;阿拉伯巴菲特&rdquo;之称的沙特王子阿勒瓦利德&middot;本&middot;塔拉勒&middot;阿苏德(Alwaleed bin Talal Al Saud)凭借其投资旗舰王国控股公司(Kingdom Holding Co. )过去五月的出色表现，个人总资产暂时比肩股神巴菲特。</p>
<p><strong>隐性资产丰厚</strong></p>
<p>彭博消息显示，过去5个月，阿勒瓦利德持有95%股权的王国控股公司股价持续攀升，自去年9月5日低点起，股价已上涨52%，同期标普500指数上涨了7%。目前王国控股公司的市值股价飞跃，为该王子总财富增加26%，约66亿美元，而同期巴菲特的总财富没什么变化。</p>
<p>&ldquo;我们还拥有大量的隐性财富，这些丰富资产并没有进行公开交易。&rdquo;这位现年58岁的沙特王子表示，&ldquo;他们价值超过110亿美元，而市场估值至少200多亿美元。&rdquo;</p>
<p>根据彭博亿万富翁指数，位居第四的巴菲特目前个人资产为579亿美元，沙特王子为319亿美元，计入王国控股公司为其增加的66亿美元财富后将达到385亿美元，若其隐性财富达到200亿美元，则将超过巴菲特的个人资产。</p>
<p>不过阿勒瓦利德却谦虚地表示：&ldquo;回顾2013年，巴菲特是从股市中获得最多利润的几位大佬之一，远超过我。不能只看四五个月的情形。我认为巴菲特仍领先于我。&rdquo;</p>
<p>根据福布斯2013全球亿万富豪榜，巴菲特去年以个人总资产535亿美元位居第四。而阿勒瓦利德王子则位居第26位，总资产为196亿美元。不过该亿万富翁对《福布斯》的财富估值公开表示异议，对把他的个人净资产列为仅200亿美元及排名感到非常失望。</p>
<p><strong>王子青睐电商</strong></p>
<p>近年来，沙特王子的成功投资已帮助其扭转自2008年全球金融危机以来一直拖累着上市公司的衰退局面。随着人们对网络服务需求的增多，阿勒瓦利德开始对电子商务和科技企业进行投资。</p>
<p><strong>去年2月， 王国控股公司带领一些投资者购买了价值15亿里亚尔(约合4亿美元)的京东商城股份</strong>，其中，王国控股公司投资了1.25亿美元。王国控股公司相关负责人此前曾表示，投资京东符合王国控股公司的私募股权融资战略，即选择可能在3年内于某个国际资本市场上市的高增长企业。&rdquo;</p>
<p>事实上，<strong>阿勒瓦利德一直积极推动京东上市，而京东也不负众望，在今年1月正式递交上市申请</strong>，计划融资15亿美元，或将成为2014年首个赴美上市的中国公司。而根据京东1月30日提交的招股书信息披露，沙特投资公司王国控股公司管理基金(Kingdom 5-KR-233, Ltd. managed funds)持股5%。</p>
<p>此外，阿勒瓦利德曾在2011年向推特网站投资了3亿美元。去年年底，他表示对推特的投资已赚了3倍，投资资产市值达12亿美元。不过上周，由于增加的活跃用户数不及预期，Twitter股价下挫24%。</p>
<p>值得注意的是，沙特王子的投资之路并不是一帆风顺。2000年阿勒瓦利德投资世通公司亏损2亿美元，2005年出售手中苹果股份，金融危机期间持有大量花旗银行股票。阿勒瓦利德承认，基于健全的信息和长期分析进行投资，投资结果有好有坏。</p>
<p align="center"><a href="/img/20140212/d55be0733d644a60af495269a02267e8.jpg" target="_blank"><img alt="持股京东待上市 沙特王子个人财富直逼巴菲特" src="/img/20140212/s_d55be0733d644a60af495269a02267e8.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292373.htm" target="_blank" title="团圆桌上“低头族” 有你吗？"><span class="title1">团圆桌上“低头族” 有你吗？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">上方文Q</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:53:58</span></td>
  </tr>
  <tr>
    <td align="left"><p>马年春节已经过去了。以往吃团年饭的时候，饭桌上总是要畅谈一下新年理想，但现在，我们发现年夜饭和拜年过程中，相当多的年青人埋首只顾玩手机。&ldquo;低头族&rdquo;真正成为了长辈们的&ldquo;眼中钉&rdquo;。</p>
<p>国人为何如此热衷玩手机？专家认为，&ldquo;低头族&rdquo;之所以产生，源于他们对自媒体的依赖。患上手机依赖症，都是过度使用自媒体造成的结果。专家建议，多阅读纸质媒体和书籍，增加自我理性思考的时间，并刻意地让自己离开自媒体，增加社交活动、户外活动都是&ldquo;低头族&rdquo;不错的选择。</p>
<p><strong>现象：</strong><strong>团圆时刻&ldquo;低头族&rdquo;只顾玩手机</strong></p>
<p>家住旧城的张先生夫妇都是上班族。平时，他们习惯在坐车、开会，甚至是午休的时候，掏出手机打发时间，久而久之两人都成了&ldquo;低头族&rdquo;。</p>
<p>有其父必有其子。张先生的儿子，目前就读大学三年级的小张也是个手机迷。现在他在一家服装公司实习，平时没事的时候就爱拿个手机玩几下。据小张自己说，有些时候一天到晚没人找他，为了打发时间，他渐渐迷上了手机游戏。为此他还晒出自己的&ldquo;战绩&rdquo;：一个500级的手游账号。</p>
<p>子孙成为&ldquo;低头族&rdquo;，最烦闷的要数盼着子孙回家过年的张爷爷张奶奶。儿子儿媳平时忙于工作，小孙子又在外上学，一家人很少有机会聚在一起，只有过年才能团圆。可是，全家人聚齐了，老两口却不像前些年那么快乐了。原来，<strong>儿子儿媳闲下来了，却不和老人聊聊天、谈谈心，而是坐在那儿玩手机，刷微博、发微信拜年；孙子小张更是一有空挡儿就玩手机游戏玩得不亦乐乎。</strong></p>
<p>张爷爷无奈地说：&ldquo;儿孙们都回来了，我和老伴儿别提有多高兴啦！可一家人坐在一起很少交流，他们时时刻刻手机不离手，我是想说也没机会说呀！难道他们真的就那么&lsquo;忙&rsquo;吗？&rdquo;</p>
<p><strong>危害：</strong><strong>颈椎病患者过半数是年轻&ldquo;低头族&rdquo;</strong></p>
<p>智能手机的普及，除了影响人们的正常交流，使得人与人之间变得越来越冷漠外，也在无形中为人们带来很多疾病。</p>
<p>据专业人员证实，<strong>经常低头玩手机对颈椎、眼睛的危害是最大的，特别是晚上睡觉前关灯玩手机对眼睛的伤害更是惊人。</strong>坐在公交车上看手机，公交车晃动会造成视力疲劳，长时间会导致视力下降。长时间低头玩手机，<strong>还会使面颊和下巴的肌肉因重力而下垂，时间长了可能形成双下巴。</strong>此外，手机显示屏上的文字小，人们会下意识地眯眼看，<strong>时间长了就会在眉间眼角形成皱纹。</strong></p>
<p>据专家介绍，在我们低头时，前屈极限(下巴碰到胸骨的状态)只能是45&deg;。如果前屈幅度达到30&deg;时，就可以影响到颈椎。如果颈椎长期处于极度前屈的异常稳定状态，就会对颈椎造成伤害，这种危害比看电脑还要高几十倍。</p>
<p>前段时间，<strong>苏州一名高二女生因长期低头看手机导致颈椎间盘突出8mm，压迫脊髓，患上严重的脊髓型颈椎病。</strong></p>
<p>一般来说，颈椎病会随着年龄的增长而增加，但如今开始有低龄化趋势，特别是近年来，随着智能手机的普及，呈快速发展趋势。</p>
<p>上个世纪90年代，某医院推拿科一个月的门诊，大约可以接诊两三百例颈椎病患者，而且都是以中老年为主，病因除了年纪增大外，还有躺在床上看电视、看书，或者长期坐姿不正确等。而易患人群则为教师、会计等需要俯首工作的职业人群。</p>
<p>可如今，他们每月至少要接诊近千例颈椎病患者。其中，20岁至35岁的大学生和白领就占了半壁江山。而几乎所有的年轻病人都会对医生说，他们喜欢长时间低头玩手机和iPad。</p>
<p><strong>观点：</strong><strong>对自媒体过分依赖 </strong><strong>会失去自我保护意识</strong></p>
<p>根据中国互联网信息中心最新发布的《中国互联网络发展状况统计报告》，截至去年6月底，我国即时通信网民接近5亿人，手机即时通信网民规模接近4亿人。<strong>在我国，坐车、走路、吃饭、开会&hellip;&hellip;各种生活场景，你总会看到人们对着手机刷呀刷、划呀划。在诧异的外国人眼中，简直成了奇观。</strong></p>
<p>国人为何如此热衷玩手机？如何不被微信&ldquo;绑架&rdquo;、异化，沦为&ldquo;信奴&rdquo;，成为摆在全民面前的新课题。</p>
<p>有专家认为，&ldquo;低头族&rdquo;产生的原因源于他们对自媒体的依赖症。该专家认为，媒体具有工具的属性，应该为人所用，而<strong>&ldquo;低头族&rdquo;已被工具所控制，出现心理疾病，患上手机依赖症，这都是过度使用自媒体造成的结果。</strong></p>
<p>专家表示，只要进入到社会环境中，无时无刻不存在风险。在公共空间中，人应该要有自我保护的意识，但是现在我们对自媒体过分依赖，失去了理性和自我保护意识，自我保护意识亦逐渐淡漠。</p>
<p>专家强调，&ldquo;低头族&rdquo;被媒体牵引，已经忽视了自我的存在，不再以理性的形式生活工作，取而代之的是非常消极的态度。&ldquo;过分依赖自媒体会造成社交障碍、心理障碍、情感的冷漠化等危害。公民同时应该对自己负责。&rdquo;某专家说。</p>
<p>&ldquo;低头族&rdquo;该如何控制自己，回归理性生活？专家建议，多阅读纸质媒体和书籍，增加自我理性思考的时间，并刻意地让自己离开自媒体，增加社交活动、户外活动都是不错的选择。</p>
<p>&#9664;&ldquo;低头族&rdquo;是颈椎病等疾病的高发人群。</p>
<p>▼微信微博拜年成了越来越多&ldquo;低头族&rdquo;的合理理由。</p>
<p align="center"><img alt="团圆桌上“低头族” 有你吗？" src="/img/20140212/45121bc9a4a149cf9788e5db7173f32e.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p align="center"><img alt="团圆桌上“低头族” 有你吗？" src="/img/20140212/be7144fef03c40cb86e84a0dbdfba263.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p align="center"><img alt="团圆桌上“低头族” 有你吗？" src="/img/20140212/96ae20d9c186473eb15d1b1f6a55c3cb.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292371.htm" target="_blank" title="同性恋者斥歧视 58同城诉侵权"><span class="title1">同性恋者斥歧视 58同城诉侵权</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:48:52</span></td>
  </tr>
  <tr>
    <td align="left"><p>近日，求职者赵鹏通过微博投诉称，去年他在求职节目中得到58同城CEO姚劲波的口头承诺，获得职位，至今无法入职。赵鹏认为，58同城拒绝其入职是因为他的同性恋身份。</p>
<p>随后，赵鹏将58同城告上法庭并要求赔偿。而58同城也向朝阳区法院递交了民事起诉状，反诉赵鹏侵犯其名誉权。同时，58同城表示，他们在招聘时不会考虑应聘者恋爱取向。</p>
<p><strong>求职者</strong></p>
<p><strong>受聘：电视节目求职获职位</strong></p>
<p>去年9月28日，赵鹏参加了求职节目《非你莫属》。10月21日，这期节目播出。节目最后，有两家企业向赵鹏提供了工作岗位，最后他选择58同城CEO姚劲波口头许诺的社交产品经理职位，月薪15000元。</p>
<p>赵鹏称，他没有和58同城签任何协议，&ldquo;姚劲波就给了我一个邮箱，叫我把简历发给他。&rdquo;</p>
<p><strong>质疑：称因同性恋入职遭拒</strong></p>
<p>赵鹏表示，节目播出之后，他就发邮件询问入职问题，姚劲波回应称&ldquo;11月初联系&rdquo;。此后，58同城的人力资源部门联系赵鹏索取其个人简历，此后还有过两三次电话联络，但他一直都没有得到入职的回复。</p>
<p>&ldquo;直到12月，姚劲波在我的微博回复，告诉我不能入职&rdquo;。赵鹏提供微博截图显示，姚劲波称赵鹏没有积极和HR部门联系，也没有针对性准备，与其公司文化不符。同时祝赵鹏在其他领域取得成功。</p>
<p>而赵鹏认为，节目结束后，他一直积极和58同城保持联系，58同城拒绝与其签订劳动合同是因为他的性取向。&ldquo;他们应该是看到我的认证微博，发现我是同性恋，然后决定不录用的&rdquo;。</p>
<p>12月18日，赵鹏在其微博发布了一条置顶微博称，&ldquo;58同城歧视同性恋&rdquo;，目前该微博已经删除。</p>
<p><strong>起诉：要求58同城赔偿损失</strong></p>
<p>由于一直没有得到58同城的入职回复，赵鹏将58同城告上法庭，并要求10万元赔偿。</p>
<p>前天，朝阳法院表示不立案。法院认为这是劳动争议案件，需要先申请劳动仲裁。赵鹏此案的代理律师赵占领则认为，在这个案件中原、被告之间不存在事实劳动关系。应该按照缔约过失责任纠纷，直接由法院受理。</p>
<p>赵占领表示，他们计划本周内向朝阳区劳动仲裁委员会申请劳动仲裁，在拿到不予受理的决定书后再去法院起诉。</p>
<p><strong>58同城</strong></p>
<p><strong>回应1：招聘不会考虑恋爱取向</strong></p>
<p>前天，姚劲波回应称，58同城曾经有非常活跃的同性交友板块，58同城招聘从来不会，以后也不会关心应聘者喜欢男的还是女的。</p>
<p>58同城也声明，本次事件仅为双方就招聘入职相关事宜产生的法律事件，并不涉及任何与恋爱取向、身体残疾等相关的内容，当事人赵鹏利用其转移事件核心焦点，其在公开媒体上所发布的言论已经对58同城企业声誉造成了不可挽回的损失。</p>
<p><strong>回应2：应聘者与企业文化不符</strong></p>
<p>58同城相关负责人表示，在当天《非你莫属》节目录制之后，姚劲波就安排公司HR部门相关负责人跟进，并与赵鹏沟通将简历发送给公司人力资源部。但是，在此次沟通后，赵鹏对于入职表现态度并不积极，在与人力的沟通中对于是否愿意继续入职58同城不置可否。在了解赵鹏通过网络微博发布的情况后，公司再次安排专人与其进行沟通。&ldquo;然而对方情绪激动并不愿与公司进行任何沟通，表现得较为强势。这给沟通带来了一定的困难，导致目前事件朝我们不愿看到的方向发展&rdquo;。</p>
<p>该负责人表示，58同城在人才招聘方面，除了注重个人能力之外，也非常重视对方是否符合我们的企业文化。招聘入职本是双方选择和考察的过程，但在此次事件中，他们认为赵鹏的表现与他们的企业文化不符。</p>
<p><strong>回应3：反诉求职者侵犯名誉权</strong></p>
<p>就在赵鹏对58同城提起诉讼后，58同城也向朝阳区法院递交了民事起诉状，反诉赵鹏侵犯其名誉权。请求法院判令赵鹏在其微博置顶位置公开赔礼道歉30天并赔偿其损失及合理支出50万元。2月2日，赵鹏已经收到了法院邮寄来的传票和起诉书。</p>
<p>对此，赵鹏表示他也在准备反诉，将在18日开庭时递交反诉状，&ldquo;起诉58同城因歧视同性恋拒绝与我缔约劳动关系&rdquo;。而对于这次反诉，赵鹏表示将由他自己进行辩护。</p>
<p align="center"><img alt="同性恋者斥歧视 58同城诉侵权" src="/img/20140212/f917806bb3eb43e9b08290c0d0a39996.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292369.htm" target="_blank" title="他会成为下一个比尔·盖茨？"><span class="title1">他会成为下一个比尔·盖茨？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">朝晖</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:46:39</span></td>
  </tr>
  <tr>
    <td align="left"><p>也许很难找出一个比萨蒂亚&middot;纳德拉更没有&ldquo;个性&rdquo;的微软员工了。</p>
<p><strong>在风气以自负闻名、争论以激烈著称的微软，纳德拉是出了名的谦逊、好脾气、善于与人合作</strong>：以至于他的同事确信，公司里任何认识他的人都至少能说出一条纳德拉的优点。与号称&ldquo;粗暴篮球迷&rdquo;的第二任微软首席执行官(CEO)史蒂夫&middot;鲍尔默不同，纳德拉的爱好是诗歌与被称为&ldquo;绅士游戏&rdquo;的板球。</p>
<p>他恐怕不会想到，有一天，比尔&middot;盖茨和鲍尔默的接力棒，交到了自己手中。</p>
<p>他在今年2月被推向全球媒体镜头聚焦处。虽然去年他还曾谦虚地对高中同学说自己距离CEO&ldquo;差得远了&rdquo;。2月4日，微软以一纸新闻稿宣布，这位文质彬彬的工程师成为了微软的新一任首席执行官。该公司官网这样介绍他：&ldquo;作为微软的第三任CEO，他为他的新角色带来了不断推动创新与协作的精神&rdquo;。</p>
<p><strong>而微软创始人比尔&middot;盖茨也对新任CEO表达了乐观的期许：&ldquo;在当前转型时期，没有人能比萨蒂亚&middot;纳德拉更好地带领微软。&rdquo;</strong></p>
<p>一夜之间，全世界都在问：他是谁？他是否有远见与能力带领微软完成艰难的转型？</p>
<p>上任第一天，纳德拉在发给全体员工的电子邮件中称：&ldquo;我今年46岁，已经结婚22年，有3个孩子。和所有人一样，我的处事方式和思维方式都是在家庭和生活经历的影响下形成的&hellip;&hellip;所以，家庭、好奇心和求知欲都定义了我&rdquo;。</p>
<p>也有人质疑：他只是一只绵羊，一个跟随者，不是一个领导型人才！</p>
<p>在一堆穿着T恤、牛仔裤的IT男中，来自印度南部一个婆罗门家庭的纳德拉是唯一一个经常穿着时尚的运动外衣出现在公司，还戴着设计师设计的眼镜的家伙。因为坚持锻炼，年近五旬的他一直保持着瘦削匀称的身材。去年，在接受印度老家一家媒体访问时，纳德拉想对&ldquo;有抱负的学生去追求梦想的学生&rdquo;留下的忠告是：&ldquo;永远都努力学习&rdquo;。</p>
<p>这毫无新意的叮嘱倒也不尽是场面话。<strong>纳德拉本人就是个例子，他报名了&ldquo;永远也学不完&rdquo;的在线课程，每天早上都给自己安排下&ldquo;超级雄心勃勃&rdquo;的15分钟，听听&ldquo;神经科学&rdquo;之类跟本行八竿子打不到一块儿的课程</strong>。</p>
<p>在发给员工的电子邮件中，他这样展开了对公司未来的期许：</p>
<p>&ldquo;我相信在未来10年&hellip;&hellip;前方的机遇将要求我们重新设想过去所作的事情，设想它们在一个移动以及云计算为主的世界中的模样，而后作出新的动作&rdquo;。</p>
<p>纳德拉所面对的形势与前任鲍尔默刚接手微软时截然不同。2000年，鲍尔默得到的是全世界最强大和最有影响力的科技公司，那时地球上绝大多数电脑使用的操作系统都出自这里。那年头，谷歌尚在琢磨如何赚钱，马克&middot;扎克伯格正在上中学，世间最酷的手机还是黑莓，当完全没有键盘的iPhone问世时，包括鲍尔默在内的科技大佬们对它全然不屑一顾。</p>
<p>如今，微软依然是世界上最大的公司之一，提供的服务与设备足够建立起一个小王国，但问题是，越来越少的消费者愿意待在这个小王国里了。</p>
<p>&ldquo;<strong>我们的行业不尊重传统，只尊重创新，&rdquo;这位新任CEO在自己的就职讲话中感叹说，&ldquo;这个行业以及微软都到了紧要关头。&rdquo;</strong></p>
<p>纳德拉已经在这间公司工作了22年。他一度负责必应搜索引擎业务，那时候他很容易跟女儿解释自己在干嘛，但后来，负责云计算和云操作系统业务的他只能说些让女儿犯迷糊的话：我做的是Windows，但又不是你们见到的那种Windows。</p>
<p>与纳德拉上任同时，比尔&middot;盖茨开始担任微软&ldquo;创始人兼技术顾问&rdquo;的新职位。只是，外界似乎对此并不乐观。<strong>两人上任第一天，《纽约客》在网站上登了一个绘声绘色的讽刺段子，狠狠地损了一把微软产品。段子说的是比尔&middot;盖茨刚返回办公室，就花了一整天在给电脑安装Windows 8 操作系统上，因为死活没法安装新系统，他找来新CEO纳德拉帮忙，两个人大眼瞪小眼半天，也没把系统升级成功。</strong></p>
<p>但纳德拉自身关注的，显然比这个段子更多，那就是带着这个世界上最有影响力的公司在即将到来的云计算为主的世界大显身手。</p>
<p>数年前，他问鲍尔默，微软应该如何跨越以往的辉煌？但鲍尔默的答案是，这样的问题没有意义，它的关注点不是公司的未来。</p>
<p>这席话让纳德拉悟出了一点：<strong>在这个行业最重要的不是能否长久，而是能否奠定影响力</strong>。</p>
<p>这个性格温和却颇有野心的新CEO，来自地球另一端。1967年，<strong>纳德拉出生在印度第六大城市海得拉巴一个公务员家庭，老爷子是一位退休政府高官，但想法不拘一格，被朋友描述为&ldquo;有着非政府组织的思维&rdquo;。据说纳德拉继承了这一点，思维不落俗套，而且简洁明了。</strong></p>
<p>在纳德拉成为微软CEO的消息传出后，这个印度城市为之沸腾。他高中时代的一个好友对记者说：&ldquo;这对印度和海得拉巴而言都是伟大的一天&rdquo;；另一个高中同学则断言，&ldquo;他是一个从来没忘记自己根在哪儿的男人！没准就是因为这样他才能取得现在的成功&rdquo;。</p>
<p>只是，从周围人的描述来看，纳德拉从来都不是一个引人瞩目的人物。他好学、聪明、愉快，从不与同学发生矛盾。少年时，纳德拉在当地中学读书，据说他在这里最大的收获，是迷上了板球&mdash;&mdash;正是这项运动，让他学会了与他人合作。只不过，当时中学的校长和一些同学对他印象并不深，他是一个不起眼，说话轻声细语的孩子。</p>
<p>如今，微软公司为这所学校提供技术支持，帮助感兴趣的学生制造机器人。</p>
<p>哪怕是现在，周围邻居也是从记者口中听说这家人&ldquo;出了那么一个大人物&rdquo;，一位邻居激动地说：&ldquo;下回他出现在这儿，我们都得瞪大了眼睛好好瞧瞧他。&rdquo;而他家的司机则用一连串方言对记者说，纳德拉是个好人，会讲两种印度方言而且都很流利，每次从美国回来都不会忘了给自己带点小礼物。</p>
<p>纳德拉移居美国是大学毕业以后的事。最初他在美国威斯康星-密尔沃基大学完成了计算机硕士课程的学习。1992年，纳德拉与他高中时代的恋人结婚，同年加入微软公司。除了这个工作机会，他还有机会攻读工商管理硕士，面对这个可能成为二选一的局面，纳德拉选择了一样也不放弃：平时工作，每周五晚飞去芝加哥大学读书。他在两年半后获得了第二个硕士学位。</p>
<p>那个时代，科技界的龙头公司还是IBM，他们的主力业务是大型计算机。</p>
<p>如今，面对微软更换CEO的局面，英国《金融时报》副主编约翰&middot;加普想起了20年前，能够威胁龙头老大IBM(179.7, 2.56, 1.45%)的企业，还是日立和富士通。但最终，谁还记得这些公司？长江后浪推前浪，那一轮浪花最高处，是微软公司席卷全球的个人电脑业务。</p>
<p>只是，新的浪花一年年都在翻涌上岸。而现在的人们，会用苹果公司的经验来衡量微软：<strong>新CEO是很好的合作者，这真的是好事？史蒂夫&middot;乔布斯就不是个优秀的合作者</strong>。</p>
<p>在微软，纳德拉从不吝于向上级展示自己善于合作的特长，去年8月，当鲍尔默要求高管团队去他的办公室开会时，纳德拉因为&ldquo;有更重要的关于并购的事情&rdquo;，不想出席。</p>
<p>但这一回，事情有点不一般。鲍尔默的助手对纳德拉说，他最好再考虑下。纳德拉终究出席了那次会议，鲍尔默在会上宣布了一个令众人震惊的消息：他将卸任微软的首席执行官一职。</p>
<p>那之后，公司董事会花了5个多月的时间，从100多位来自各行各业的潜在人选中寻找一位新的CEO，候选人中甚至包括福特汽车公司CEO穆拉利。<strong>今年年初，当最终人选定下时，负责挑选的董事已经精疲力竭。《华尔街日报》分析称，过程如此漫长，意味着多年来影响力持续下降的微软公司已经无法吸引一位出色的CEO了。</strong></p>
<p>而在22年前，从印度来到华盛顿州的纳德拉，心怀着改变世界的梦想。</p>
<p>&ldquo;我来到这里，因为我相信微软是世界上最棒的公司，我清晰地看到，我们的创新使得人们能做许多神奇的事情，并且将世界变得更加美好。我知道如果我想作出一点改变，来到这里是最好的选择。&rdquo;</p>
<p>他从未忘记，在微软公司早期的岁月，大家的使命是让个人电脑进入每家每户，被摆放上每一张办公桌。</p>
<p>在22年后的今天，这些梦想实现了。他就任的消息，化成不同的语言文字，通过光纤，出现在遍布五大洲的那些个人电脑的屏幕上。</p>
<p align="center"><a href="/img/20140212/99234e9bf9c04f62af00365f41ea0121.jpg" target="_blank"><img alt="他会成为下一个比尔·盖茨？" src="/img/20140212/s_99234e9bf9c04f62af00365f41ea0121.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
萨蒂亚&middot;纳德拉（中）与比尔&middot;盖茨（中左）和史蒂夫&middot;鲍尔默（中右）一起对员工发表讲话。</p>
<p align="center"><a href="/img/20140212/d9159b26c4364c3ea17e23c4f0fead0c.jpg" target="_blank"><img alt="他会成为下一个比尔·盖茨？" src="/img/20140212/s_d9159b26c4364c3ea17e23c4f0fead0c.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a><br />
纳德拉</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292368.htm" target="_blank" title="小米手机春晚备选广告曝光：《青春万岁》"><span class="title1">小米手机春晚备选广告曝光：《青春万岁》</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">鲲鹏</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:46:31</span></td>
  </tr>
  <tr>
    <td align="left"><p>马年春晚前，小米在黄金时段<a class="f14_link" href="http://news.mydrivers.com/1/291/291205.htm" target="_blank">投放了长达1分钟的广告片</a>，让众人大呼小米&ldquo;真土豪&rdquo;。</p>
<p>小米这部广告片名为《我们的时代》，<strong>整体延续了小米此前广告片那种励志的风格</strong>，总体来说就是&ldquo;不经历风雨怎能见彩虹&rdquo;。小米对该视频的介绍是：我们的名字叫年轻，在追逐梦想的路上，我们不断向前，去探索、去改变、去拼搏，我们的时代来了，小米，为发烧而生！</p>
<p>事实上小米共为本次春晚准备了多部备选广告，最终《我们的时代》成功脱颖而出。昨天下午，<strong>雷军就在微博上曝光了春晚的备选广告方案《青春万岁》，</strong>整体风格和《我们的时代》类似，但励志效果似乎不如后者。</p>
<p>视频观看地址：</p>
<p><a class="f14_link" href="http://v.youku.com/v_show/id_XNjcxNDEwNzU2.html" target="_blank">http://v.youku.com/v_show/id_XNjcxNDEwNzU2.html</a></p>
<p align="center"><a href="/img/20140212/246165f99c574196a17a01f6df7f8ed7.jpg" target="_blank"><img alt="小米手机春晚备选广告曝光：《青春万岁》" src="/img/20140212/s_246165f99c574196a17a01f6df7f8ed7.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/a3210d79622e4f6e87102d0320cf3b26.jpg" target="_blank"><img alt="小米手机春晚备选广告曝光：《青春万岁》" src="/img/20140212/s_a3210d79622e4f6e87102d0320cf3b26.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/e757d17a884d4015bfa6bac2659c1516.jpg" target="_blank"><img alt="小米手机春晚备选广告曝光：《青春万岁》" src="/img/20140212/s_e757d17a884d4015bfa6bac2659c1516.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/1baa634c411742899c98d01a4c964385.jpg" target="_blank"><img alt="小米手机春晚备选广告曝光：《青春万岁》" src="/img/20140212/s_1baa634c411742899c98d01a4c964385.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/7ee80d737c164e299b9e86efb10b024f.jpg" target="_blank"><img alt="小米手机春晚备选广告曝光：《青春万岁》" src="/img/20140212/s_7ee80d737c164e299b9e86efb10b024f.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/2b803ad89708408f9df6883b7b2a26c6.jpg" target="_blank"><img alt="小米手机春晚备选广告曝光：《青春万岁》" src="/img/20140212/s_2b803ad89708408f9df6883b7b2a26c6.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/42c2a0efc16b4051aacb397493d8c8e5.jpg" target="_blank"><img alt="小米手机春晚备选广告曝光：《青春万岁》" src="/img/20140212/s_42c2a0efc16b4051aacb397493d8c8e5.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
<p align="center"><a href="/img/20140212/fcc567527ccd4ae7a0ee48f74b5e72f3.jpg" target="_blank"><img alt="小米手机春晚备选广告曝光：《青春万岁》" src="/img/20140212/s_fcc567527ccd4ae7a0ee48f74b5e72f3.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292367.htm" target="_blank" title="谁偷走了你的手机流量？揭秘流量陷阱"><span class="title1">谁偷走了你的手机流量？揭秘流量陷阱</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:46:02</span></td>
  </tr>
  <tr>
    <td align="left"><p>近日，安全研究及漏洞报告平台乌云网站发布安全警示，指出手机软件被篡改，乱弹广告后台偷跑流量的现象日趋严重，而关于手机有可能偷跑流量的消息在众多智能手机用户当中引发了不小的反响，有超过30%的网友反馈曾经出现手机偷跑流量的现象。</p>
<p>至于流量偷跑的原因，部分有一定手机知识基础的被访者称，会采用流量监控软件进行测试，发现偷跑流量的软件有很多是在第三方软件市场下载的，它们被黑客&ldquo;二次打包&rdquo;写入了恶意代码。事实上，这只是流量偷跑的&ldquo;冰山一角&rdquo;，算是恶意而为之的部分。另外，还有些&ldquo;无心插柳&rdquo;的流量偷跑情况。专家提示，想要完全杜绝流量偷跑，除非关闭数据功能。否则，放弃你的智能手机吧。</p>
<p><strong>上网就无法避免流量偷跑</strong></p>
<p>&ldquo;手机偷跑流量&rdquo;，按照通常情况看，大都指在消费者不知道的情况下，手机因移动上网而产生不明流量甚至费用。<br />
<br />
当今信息化社会中，智能手机几乎成了人们生活的必需品，据CNNIC近日发布的《第32次中国互联网发展状况统计报告》显示，2013年上半年，手机成为新增网民的第一来源，占新增网民的70%。受其影响，作为互联网第一大应用&mdash;&mdash;即时通信(微信类产品)在手机端也有不俗的表现，其网民规模达3.97亿，使用率为85.7%。随着即时通信应用逐渐从单一的聊天工具向综合平台转变，其市场潜力也被激发。</p>
<p>和短信不一样，即时通信软件是通过流量发送信息的。也就是说在没有手机信号的情况下，只要有Wi-Fi网络，也能发信息。如今，越来越多的人开始用微信替代短信。如此一来，手机上网成了必不可少的生活配件。</p>
<p>如果在功能机年代，手机上网便上网罢了。因为都是单任务运行，基本不会出现偷跑流量的问题。智能机时代则不然，有手机专家表示，想要杜绝偷跑流量，除非关闭上网功能。但这样智能手机便没了用武之地，还不如扔了，并不现实。</p>
<p>&ldquo;然而，即便是偷跑流量，也有&lsquo;有意&rsquo;和&lsquo;无意&rsquo;两种。&rdquo;该专家称，大多数流量偷跑是无意的，电信运营商不会和应用开发商就流量获得利润进行分成，所以即便是有意而为之的偷跑流量，目的也是弹出更多的广告或实现内购等收费项目，&ldquo;醉翁之意不在酒&rdquo;。</p>
<p><strong>流量&ldquo;偷跑&rdquo;五大陷阱</strong></p>
<p>家住在新城的陈先生的手机一夜之间不仅用光了套餐中所有流量，还欠费700多元。据陈先生介绍，他睡觉前忘记关闭数据流量业务，当套餐内流量走完时，的确有收到运营商发来的提醒短信，但当时他正在睡觉，看到短信已是第二天早上了。对此，陈先生很无奈也很疑惑：我的手机流量是如何被偷走的？</p>
<p>据介绍，很多原因都可能导致&ldquo;流量偷跑&rdquo;，但最常见的有5种。推送功能的实现必须建立在网络联接的基础上，才能检测软件更新以及消息提醒。因此你即使没有主动上网，在网络联接的情况下也会在后台工作，必然会产生一定的流量。</p>
<p>一旦用户不慎将带病毒的恶意软件下载到手机中，木马程序便会开启Root权限进行后台联网，下载用于恶意推广的软件并同步消耗用户的流量，导致流量偷跑。</p>
<p>智能手机在Wi-Fi信号不稳定和较弱时，有自动搜索并切换到2G/3G网络，继续提供上网服务的功能。如果用户没有注意到，继续使用需要耗费流量的服务时，就可能产生高额的流量费了。</p>
<p>现在很多手机都有云端同步功能，可以帮助使用者将联系人或短信甚至是照片等信息自动备份到云端，但是这样会耗费许多流量，特别是开启图片同步上传功能时，流量的消耗会更厉害。</p>
<p>很多运行在用户手机中的游戏软件，功能并不像看上去那么简单，除了明面上为用户提供的游戏服务之外，往往还会在暗地里搜集用户的机型、地理位置等个人信息，此时，流量就这么悄悄地逃走了。</p>
<p><strong>安卓系统流量被偷？解决方法请看这里！</strong></p>
<p>大家都知道，安卓系统不管是系统自带的软件，还是我们后来自己安装的软件，部分软件总会在我们不知不觉的情况下&ldquo;偷&rdquo;走我们的流量。今天就来告诉你如何防止这种事情发生。</p>
<p>第一，控制好你的GPRS端口。每台安卓手机都会有GPRS端口。你可以通过&ldquo;设置&rdquo;&mdash;&ldquo;无线与网络&rdquo;&mdash;&ldquo;移动网络&rdquo;中关闭端口。当你用完上网的功能时，关闭该端口，软件就不可以在你不知不觉的情况下把你的流量偷走了。如果你觉得麻烦，你可以通过辅助软件的方法来一键关闭该端口。</p>
<p>第二，通过软件监控实时流量。</p>
<p>第三，巧用任务管理器。当你觉得你的手机流量莫名其妙增加的时候，你打开任务管理器查看有什么程序正在运行。直接结束掉无关的进程。</p>
<p align="center"><img alt="谁偷走了你的手机流量？揭秘五大流量陷阱" src="/img/20140212/4e539db093e74724a535c7719d64abe3.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292366.htm" target="_blank" title="央行正牵头监管互联网金融"><span class="title1">央行正牵头监管互联网金融</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">萧萧</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:43:56</span></td>
  </tr>
  <tr>
    <td align="left"><p>知情者向《华尔街日报》透露，中国央行正牵头部署对快速发展的互联网金融行业的监管，以防范该行业日益上升的风险。</p>
<p>监管措施还将涵盖基于互联网的P2P信贷公司，行业已有部分公司倒闭。</p>
<p>知情者称，一些官员担忧投资者不知道自己的钱到底去了哪儿，另外还存在个人信息被盗的风险。</p>
<p>知情者称，<strong>中国央行正会同银监会、证监会和保监会一道，将实施措施确保消费者信息不被盗或滥用，充足的产品风险披露，同时禁止非法集资行为</strong>。</p>
<p>一位央行官员表示，&ldquo;目前，监管目标不是打压行业，而是培育其健康发展。&rdquo;</p>
<p>1月23日，有媒体报道称，国务院已点头同意，由央行牵头成立的&ldquo;互联金融协会&rdquo;，该协会归口央行分管。</p>
<p>央行有关人士透露，&ldquo;协会的成立，预示着未来互联网金融的相关业务将纳入到监管中来，先从行业规范、自律开始。目前央行正联合几大部委起草规范互联网金融的相关文件。&rdquo;</p>
<p>据了解，&ldquo;<strong>互联网金融协会&rdquo;是全国性的一级协会，直接由央行分管</strong>。而不是隶属于某个一级协会下的二级单位，或二级协会。这与去年12月3日由央行分管的支付清算协会成立的互联网金融专业委员会，以及今年1月16日由工信部分管的中国互联网协会成立的互联网金融工作委员会，有着本质的不同。</p>
<p>据知情人士透露，成立全国性的互联网金融协会，只是互联网金融发展与监管课题内容的很小一部分。该课题是去年由央行牵头承担的国务院19大课题之一。课题组已于今年1月16日正式向国务院汇报，提出未来的监管思路，即支持和包容互联网金融发展，但绝不能踩非法吸收公众存款，和非法集资的底线等。</p>
<p>另据了解，中国支付清算协会2月10日下发通知，决定2月27日召开P2P网贷业务座谈会，以了解成员单位业务发展情况以及相关服务需求，加强交流沟通，为监管政策的制定建言献策。</p>
<p align="center"><a href="/img/20140212/4e6f9c33926f4072a9c8cb00681f185c.jpg" target="_blank"><img alt="央行正牵头监管互联网金融" src="/img/20140212/s_4e6f9c33926f4072a9c8cb00681f185c.jpg" style="border-top: black 1px solid; border-right: black 1px solid; border-bottom: black 1px solid; border-left: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292365.htm" target="_blank" title="什么？iPhone 6竟用无边框屏幕"><span class="title1">什么？iPhone 6竟用无边框屏幕</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:34:52</span></td>
  </tr>
  <tr>
    <td align="left"><p>今天早些时候，韩国媒体再度给出了让人惊叹的消息，即iPhone 6会采用无边框屏幕。</p>
<p><strong>除了提到iPhone 6会采用无边框屏幕外，报道中还声称，iPhone 6会有4.7寸以及5.5寸两个版本</strong>。</p>
<p>该韩国媒体在报道中称，<strong>三星Galaxy S5将采用无边框设计，并且增加指纹识别识别功能（标准指纹传感器），而苹果目前正在研发的iPhone 6原型机与S5极为相似，都采用了无边框屏幕设计。</strong></p>
<p>其实，相同的消息去年也曾出现过，当时彭博社曾报道，iPhone 6会配备大屏幕，并且是无边框设计，而屏幕外层有弧面玻璃（玻璃中心拱起并向两侧曲线下滑至机身边框成为统一整体），不过《华尔街日报》则不认为苹果会为iPhone 6配备弧形玻璃。</p>
<p>别的暂且不说，如果iPhone 6使用无边框屏，相信视觉冲击力一定很强。</p>
<p align="center"><a href="/img/20140212/50832fc36a6041fbb2a432ac2efbc44e.jpg" target="_blank"><img alt="什么？iPhone 6竟用无边框屏幕" src="/img/20140212/s_50832fc36a6041fbb2a432ac2efbc44e.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a><br />
概念无边框iPhone设计</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292364.htm" target="_blank" title="智能手机屏占比大比拼：iPhone垫底"><span class="title1">智能手机屏占比大比拼：iPhone垫底</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:21:53</span></td>
  </tr>
  <tr>
    <td align="left"><p>屏占比是指屏幕与手机正面的对比值。对于触屏智能手机来说，更大的屏占比才会有更强的视觉冲击力。更大的屏占比还能在不增加手机尺寸的情况下让用户看到更多的内容，这也就是为什么如此多的iPhone概念渲染图会选择这种无边框的设计。国外有用户制作了一张热门智能手机屏占比对比图，包含了来自LG、三星、HTC、诺基亚、摩托罗拉和苹果的多款流行智能手机型号。</p>
<p>结果LG G2以75.7%的屏占比获得了冠军，而很多型号的iPhone的屏占比都低于60%。今天早些时候，有传言称下一代iPhone可能采用无边框设计，这样会进一步增加屏占比。</p>
<p>此外，还有专利显示未来苹果可能会将Touch ID指纹识别传感器集成在屏幕下方，这样iPhone下巴部分也可以省掉了？</p>
<p align="center"><img alt="智能手机屏占比大比拼：iPhone垫底" src="/img/20140212/92001e864f894d8fb656047661501629.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
<p><strong>感谢CV小天的投递</strong></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292363.htm" target="_blank" title="诺基亚征专利费：国产厂商集体悲剧"><span class="title1">诺基亚征专利费：国产厂商集体悲剧</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:18:41</span></td>
  </tr>
  <tr>
    <td align="left"><p>在诺基亚专利&ldquo;大棒&rdquo;下， HTC已经低头，并承诺向诺基亚支付现金以获得专利授权。不过对于国内其他品牌手机厂商来说，并不是一个好消息。</p>
<p>多家手机厂商大佬与笔者近期的聊天中，已透露出对诺基亚未来转型&ldquo;专利公司&rdquo;的忧虑。在诺基亚、微软、爱立信、高通等公司的专利&ldquo;大棒&rdquo;下，国产手机每出口一台智能手机，就面临缴纳超过10%的专利授权费用，远高于企业自身利润。</p>
<p><strong>国产手机厂商加速出海</strong></p>
<p>2014年被视为国产手机厂商集体&ldquo;出海&rdquo;的关键一年，与十年前那轮以 GSM功能机为主的&ldquo;山寨出海&rdquo;、主打印度非洲等新兴市场不同。在智能手机市场迅速崛起的&ldquo;中华酷联&rdquo;纷纷投入重金、以自主品牌向欧美国际一线市场规模进军。</p>
<p>一位手机厂商老总对笔者表示，中国已成为全球智能手机厂商数量最多、竞争最激烈、同时也是利润最低的主要区域市场。而根据IDC数据，华为、联想、中兴、酷派、TCL等越来越多的中国厂商都进入全球出货排名前10位。在国内市场逐渐饱和、盈利压力和向高端品牌目标促使下，国产品牌厂商2014年都加大了海外布局。联想甚至出资 29亿美元收购摩托罗拉移动正式开拓海外市场。</p>
<p><strong>诺基亚开征手机专利费：大棒随时落下</strong></p>
<p>不过，该人士也表示，过去华为、中兴手机拓展欧美发达市场，每3G/4G Android手机要分别向微软、高通的等缴纳专利费用。但在微软收购诺基亚、全球手机行业进入 4G格局后，这一状况正在加剧。</p>
<p>微软收购诺基亚案中，并不包括收购诺基亚的专利。而保留了专利的诺基亚也不再进行手机生产制造，之前业内专利交叉授权的惯例今后对诺基亚并不适用，因此诺基亚正谋求其专利利益最大化。诺基亚发言人马克&middot;杜兰特之前就曾表示：&ldquo;当手机业务剥离后，我们将会尝试进行技术对外许可。&rdquo;</p>
<p>目前，诺基亚在美国拥有1.6万个通讯专利权，并在美国亦有约4500个待审通讯专利；在欧洲确定与待审专利合计更高达2万个。更重要的是，诺基亚持有4G LTE专利比重达19%，远超排名第二的高通12.5%。也就是说，任何手机厂商制造4G手机都已绕不开诺基亚。</p>
<p>同时，微软也围绕手机软件与硬件在全球持有超过34000项专利，谷歌Android系统本身开源免费，但大型Android手机厂商还必须向微软而不是谷歌缴纳专利授权费用。随着微软对诺基亚的收购，业界预期其很快就会启动新的专利策略，一方面通过打压Android厂商扩大营收，另一方面也向手机企业推销市场份额仅3%的Windows Phone系统。</p>
<p><strong>专利授权费远高于国产厂商利润</strong></p>
<p>更令手机厂商郁闷的是，爱立信也计划对在欧美市场销售的4G手机征收LTE专利费用。一位国产品牌手机厂商副总对笔者抱怨，在国内市场利润越来越低，而走向国际市场也面临同样状况，并且随时还会因专利问题被禁售。</p>
<p>具体来说，如一部售价150美元的国产智能手机，目前要缴纳的专利费用包括：</p>
<p>1、高通按照手机售价收取4%-5%；</p>
<p>2、爱立信按照手机售价收取3%；</p>
<p>3、诺基亚按照手机售价收取2%-3%；</p>
<p>4、微软以每台Android手机5美元、平板电脑10美元的价格收取专利许可费用。</p>
<p>也就是说，国产手机每部手机出海即面临超过10%的高额专利缴纳费用，远高3G时代，也高于国产手机厂商目前的自主利润率。</p>
<p><strong>监管部门需关键时期&ldquo;扶一程&rdquo;</strong></p>
<p>业内分析人士认为，全球主要手机制造业都集中在中国，同时中国还是全球最大的智能手机消费市场。不过，国产手机缺乏技术和专利积累，在核心配件、系统软件方面更是没有话语权，国产手机企业需要国家相关管理部门在关键时期&ldquo;扶一程&rdquo;。</p>
<p>如欧盟就曾警告诺基亚不要成为&ldquo;专利投机者&rdquo;，否则将启动反垄断调查。在对高通的反垄断调查中，韩国政府在2009年就曾开出了2730亿韩元(约合2.52亿美元)的罚款。而中国发改委针对高通的反垄断调查也需要有力度，同时扩大范围到诺基亚等&ldquo;专利公司&rdquo;，敦促其专利费用降至合理的范围之内。</p>
<p>手机中国联盟秘书长王彦辉认为：&ldquo;往坏的想，可能会发生DVD产业这个结果吧。中国的一些公司生产DVD可能就挣几块钱人民币，但是海外的专利拥有者，他会挣到很多美金。结果就是中国的整个产业，彻底沦为制造业。&rdquo;</p>
<p>多家手机厂商老总也同时呼吁，希望成立专业行业协会，通过建立&ldquo;专利池&rdquo;共同应对国际专利问题，同时协调国产厂商之间减少专利战、价格战，维护中国手机企业共同的利益和形象。</p>
<p align="center"><a href="/img/20140212/5e82e3b251d44ff88019674215a0757c.jpg" target="_blank"><img alt="诺基亚征专利费：国产厂商集体悲剧" src="/img/20140212/s_5e82e3b251d44ff88019674215a0757c.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292362.htm" target="_blank" title="Galaxy S5外形曝光：太三星了"><span class="title1">Galaxy S5外形曝光：太三星了</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">08:14:25</span></td>
  </tr>
  <tr>
    <td align="left"><p>昨天晚些时候我们报道了Galaxy S5的配置参数，想必大家还想知道它到底长什么样吧？</p>
<p>之前曾多次曝光苹果新品的消息人士Sonny Dickson，其带来了所谓Galaxy S5的外形设计图，虽然目前还没法证实其准确性，但很有三星的风格。</p>
<p>从设计图上来看，<strong>该机的外形与Galaxy S4很相似，机身三围是141.7&times;72.5&times;8.2mm，作为对比GS4的机身尺寸为136.6&times;69.8&times;7.9mm，而从昨天曝光的参数图来看，GS5配备的是5.25寸2K分辨率屏幕。</strong></p>
<p>据悉，Galaxy S5还将搭载2.5GHz骁龙800处理器（型号为MSM8974AC），内置3GB RAM和3000mAh容量电池，支持4G网络，提供2000万像素摄像头（有消息称1600万像素摄像头），运行Android 4.4系统。</p>
<p>如果GS5外形真的是这样，大家有爱吗？</p>
<p><strong>相关阅读：</strong></p>
<p><a class="f14_link" href="http://news.mydrivers.com/1/292/292347.htm" target="_blank">三星全新UI来了</a></p>
<p><a class="f14_link" href="http://news.mydrivers.com/1/292/292352.htm" target="_blank">超赞！三星新旗舰Galaxy S5配置一览</a></p>
<p align="center"><a href="/img/20140212/82db7d68077d4cdbb82ab5f8d3d12797.jpg" target="_blank"><img alt="Galaxy S5外形曝光：太三星了" src="/img/20140212/s_82db7d68077d4cdbb82ab5f8d3d12797.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292361.htm" target="_blank" title="诺基亚要推Android机 微软你点头了吗？"><span class="title1">诺基亚要推Android机 微软你点头了吗？</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">07:54:48</span></td>
  </tr>
  <tr>
    <td align="left"><p>诺基亚终于还是没有与Android绝缘。据路透社、《华尔街日报》等国外媒体援引可靠线人的消息称，诺基亚将在本月底在巴塞罗那举行的移动世界大会(MWC)期间推出一款廉价的Android智能手机。</p>
<p>虽然诺基亚、微软均拒绝对此置评，但据网易科技从多个渠道了解，诺基亚确实有一个很小规模的团队从去年就开始研发Android手机，不仅如此，诺基亚在中国已经开始了销售渠道的准备工作。</p>
<p>即将被微软收购的诺基亚，真的会如期推出Android手机吗？诺基亚即将推出的这款Android手机，会不会像当年的Meego手机一样成为&ldquo;独子&rdquo;？如果诺基亚的这一行动被微软默许了，Android手机能给诺基亚带来什么？Android市场的竞争已经如此惨烈，消费者还能为诺基亚的廉价Android手机&ldquo;叫好又叫座&rdquo;吗？</p>
<p><strong>与Android的未解之缘</strong></p>
<p>当年，诺基亚很坚决地拒绝了Android，认为只有微软的Windows Phone才是拯救诺基亚的唯一正确选择，于是就有了后来的合作姻缘。</p>
<p>但在双方合作一段时间之后，诺基亚高管的口风发生了微妙变化&mdash;&mdash;&ldquo;如果Windows Phone战略失败，诺基亚还有紧急备份计划&rdquo;、&ldquo;如果可以重来，诺基亚也许会改变此前做出的一些决定&rdquo;。</p>
<p>尤其是后来，诺基亚高管在接受媒体采访时直接表达了&ldquo;对Android持开放态度&rdquo;的观点。而诺基亚高管的这一措辞，被媒体解读为&ldquo;诺基亚站队的动摇&rdquo;。</p>
<p><strong>从目前的消息来看，诺基亚与Android确实不是一对绝缘体。据一位诺基亚员工介绍，诺基亚在中国有一个很小规模的团队，从去年开始就在研发Android手机。另据可靠消息人士透露，诺基亚在中国的销售渠道也在为Android手机的上市做准备。</strong></p>
<p>看起来，诺基亚的Android手机真的要来了。不过，有消息称，即将推出的诺基亚Android 手机与传统的Android手机不同，它的系统和界面都将经过深度定制，不会包含任何谷歌开发的关键功能。</p>
<p>也有业内人士分析称，诺基亚研发Android手机，很有可能是在诺基亚&ldquo;贱卖&rdquo;给微软之前做出的决定，虽然这项收购案还没有完成，但也不排除&ldquo;这一发布计划被微软叫停&rdquo;的可能。</p>
<p><strong>诺基亚用意何在？</strong></p>
<p>目前，市场上对诺基亚Android手机的命运有三种猜测：一是在发布之前被微软叫停，诺基亚与Android的缘分到此为止；二是由于收购案还没完成，诺基亚侥幸发布了这款手机，但这款手机最终成为&ldquo;独子&rdquo;，就像当年的Meego平台手机一样；三是微软默许了诺基亚的Android计划，但系统内置的必须是微软或诺基亚自己的应用程序。</p>
<p>无论结果如何，诺基亚把一部分（即便是很小一部分）精力拿去做Android手机，已经表露了它最初做这个决定的用意。</p>
<p>一方面，在押注Windows Phone的智能手机转型策略实施两年多后，诺基亚并没有如愿迎来&ldquo;扭转颓势&rdquo;的新局面，整个Windows Phone阵营也没有太多起色。据Strategy Analytics的数据显示，2013年Android手机占全球智能手机出货量的79%，iPhone占15 %，而Windows Phone仅占4%。</p>
<p>另一方面，即便是诺基亚在短期内不可能&ldquo;转投&rdquo;Android阵营，但它把一部分资源拿出来，推出面向新兴市场的Android廉价手机，也是有必要的&mdash;&mdash;诺基亚原来庞大的低价功能机市场，正在被同样低价的来自中国厂商的Android智能手机所侵蚀，而这一市场是诺基亚手机销售的重要来源，也是为它长期未见起色的Windows Phone中高端机市场&ldquo;补血&rdquo;的来源。</p>
<p><strong>消费者能接受吗？</strong></p>
<p>如果诺基亚的Android廉价手机一款接一款地上市了，你会去买一部吗？</p>
<p>不少消费者的回答是肯定的。对于大部分消费者，尤其是那些70后、80后来说，诺基亚是他们美好回忆的一部分：人生第一部手机是诺基亚，在大学宿舍里用来煲电话粥的是诺基亚，甚至用来砸核桃的也是诺基亚&hellip;&hellip;</p>
<p>但也有一部分消费者对未来的诺基亚手机表示怀疑：它还能那么结实、抗摔吗？它不兼容谷歌的一些应用了，系统还会好用吗？中国有那么多超便宜、配置又不低的Android手机，还有买诺基亚的必要吗？</p>
<p>从目前的Android市场竞争状况来看，步伐晚了N年的诺基亚想要在这个市场有所获，也并不是一件容易的事情。</p>
<p>在全球，Android手机的竞争已经非常惨烈。三星霸占了Android智能手机市场的大部分份额，其中的低端手机也销售火爆；刚刚宣布收购摩托罗拉移动的联想，正在向海外市场扩张，而联想也一直专注于中低价位的Android机型；砍掉了PC业务索尼，也开始倾其资源在Android手机等移动产品上。</p>
<p>不仅如此，中兴、华为、酷派、TCL、小米、魅族、步步高、金立等来自中国的&ldquo;军团&rdquo;，都是争抢Android市场蛋糕的强大势利，而且都是销售中低价位手机的高手。</p>
<p>在竞争如此惨烈Android手机市场，刚刚探出头的诺基亚未来会怎样？</p>
<p align="center"><a href="/img/20140212/1951701a7bf145699efa2f4630cb636a.jpg" target="_blank"><img alt="诺基亚要推Android机 微软你点头了吗？" src="/img/20140212/s_1951701a7bf145699efa2f4630cb636a.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></a></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292360.htm" target="_blank" title="苹果密谈索尼：下一代iPhone用新前置摄像头"><span class="title1">苹果密谈索尼：下一代iPhone用新前置摄像头</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">07:40:38</span></td>
  </tr>
  <tr>
    <td align="left"><p>虽然iPhone 6还没有跟我们见面，但从目前的情况来看，其前置摄像头应该不会有太多吸引人的地方。</p>
<p>一直以来，苹果都把提升后置摄像头作为重点，<strong>现在日本媒体给出的消息称，苹果已经开始与索尼进行谈判，其希望后者为明年的iPhone提供质量更高的前置摄像头</strong>（之前基本都是OV前置摄像头）。</p>
<p>报道中还提到，索尼应苹果的要求，在今年1月决定收购日本芯片厂商瑞萨电子旗下的一家工厂，为的是提高CMOS传感器产量。</p>
<p>PS：iPhone 6是别想了，iPhone 7用上索尼的前置摄像头应该没啥问题。</p>
<p align="center"><img alt="苹果密谈索尼：下一代iPhone用新前置摄像头" src="/img/20140212/fd45cae17fdb4abca47546bebcada96a.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /><br />
&nbsp;</p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
           <table width="82%" border="0" align="center" cellpadding="3" cellspacing="0">
  <tr>
    <td height="30" align="left" style="font-size:14px"><a href="/1/292/292359.htm" target="_blank" title="铁杆粉丝吐槽：Windows 8完全是一个灾难"><span class="title1">铁杆粉丝吐槽：Windows 8完全是一个灾难</span></a>　<img src="/images/c-l-g-t-01.gif" width="12" height="11" /> <span class="f12g">雪花</span> <img src="/images/c-l-g-t-02.gif" width="12" height="11" /> <span class="f12g">07:26:56</span></td>
  </tr>
  <tr>
    <td align="left"><p>虽然微软Windows 8的下一个版本要到今年4月才会发布，但这个名为Windows 8 Update1的新版本已经遭到微软的一名铁杆支持者的吐槽：Windows 8完全就是一个灾难。</p>
<p>在业内颇具影响力的Windows博客&ldquo;Windows超级网站&rdquo;的博主保罗杜罗特称，新版本只会让Windows 8变得更难使用，尤其是在平板电脑上。<strong>杜罗特在博客中写道：&ldquo;Windows 8完全就是一个灾难。这一点毋庸讨论，亦非妄言。Windows 8就是一个灾难，完毕。&rdquo;</strong></p>
<p>没有必要去重新规整用户们对Windows 8提出的不满意见，<strong>一言以概之：Windows 8设计得不好，它太不好用了。</strong></p>
<p>虽然微软原先的设想是好的，它打算通过Windows 8将PC和平板电脑合二为一，但是杜罗特说，事情并不想微软想的那样，Windows 8完全就是一团糟，因为它实际上就是由两个不同的操作系统即移动系统和桌面系统糅合在一起组成的怪物，就象弗兰肯斯坦的怪物一样。</p>
<p>他为Windows 8出现这样的情况找了个很好的原因：这完全都是拜Windows前高管史蒂文辛诺夫斯基所赐！</p>
<p>杜罗特写道：&ldquo;公平地讲，辛诺夫斯基在某些方面与乔布斯是很相似的，两人在性格上都比较好斗、偏执、很难与他人合作，拒绝接受那些不是由他领导的团队开发出来的新功能和新技术，而且完全不考虑客户反馈信息。但是他没有乔布斯最宝贵的一项特质：与生俱来的、对完美设计的理解。&rdquo;</p>
<p>自从Windows 8被发布以来，微软所做的事情就是不断地道歉。它想让所有的人都满意，今天改改这里，明天改改那里，结果反而把用户们弄得更摸不着头脑了。</p>
<p>杜罗特说，解决这个问题的方法很简单，那就是不要再试图将Windows打造成任何人都能用的&ldquo;万金油&rdquo;了。他建议微软将关注的重点放在它最重要的用户即企业用户和生产线工人上。</p>
<p align="center"><img alt="铁杆粉丝吐槽：Windows 8完全是一个灾难" src="/img/20140212/b13a00512980473e9844f6df8a5cf936.jpg" style="border-bottom: black 1px solid; border-left: black 1px solid; border-top: black 1px solid; border-right: black 1px solid" /></p>
</td>
  </tr>
  <tr>
    <td><hr align="center" color="#ffffff" size="1" /></td>
  </tr>
</table>
            
            <table width="870" border="0" align="center" cellpadding="0" cellspacing="0" bordercolor="#111111" class="table2" style="border-collapse: collapse">
              <tr>
                <td align="left"> 过去5天的新闻&nbsp;|&nbsp;<table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td height="25"><span class="STYLE8">2014年02月12日&nbsp;&nbsp;星期三</span></td></tr><tr><td height="25"><a href="/blog/20140211.htm" class="STYLE9">02月11日</a>&nbsp;|&nbsp;<a href="/blog/20140210.htm" class="STYLE9">02月10日</a>&nbsp;|&nbsp;<a href="/blog/20140209.htm" class="STYLE9">02月09日</a>&nbsp;|&nbsp;<a href="/blog/20140208.htm" class="STYLE9">02月08日</a>&nbsp;|&nbsp;<a href="/blog/20140207.htm" class="STYLE9">02月07日</a>&nbsp;|&nbsp;<a href="/blog/20140206.htm" class="STYLE9">02月06日</a>&nbsp;|&nbsp;</td></tr></table></td>
              </tr>
              <tr>
                <td> 希望给我们提供新闻？请将消息发至<a href="mailto:news@mydrivers.com" style="text-decoration: none; color: #0000FF; font-size: 9pt">news@mydrivers.com</a></td>
              </tr>
            </table></td>
        </tr>
      </table></td>
  </tr>
</table>
<table width="923" height="6" border="0" align="center" cellpadding="0" cellspacing="0">
  <tr>
    <td></td>
  </tr>
</table>
<table width="100%" height="1" border="0" align="center" cellpadding="0" cellspacing="0" bgcolor="#969696">
  <tr>
    <td></td>
  </tr>
</table>
<table width="100%" border="0" align="center" cellpadding="0" cellspacing="0">
  <tr>
    <td height="33" colspan="2" align="center" background="http://icons.mydrivers.com/www/index-banquan_01.gif" class="f12_blue1-copyright"><table width="100%" height="22" border="0" cellpadding="0" cellspacing="0">
      <tr>
        <td align="center" valign="bottom" class="f12_blue1-copyright">【<a href="http://www.mydrivers.com/contact/contactus.html" target="_blank" class="copyright-blue">联系我们</a>】   -   【<a href="http://www.mydrivers.com/contact/advertising.html" target="_blank" class="copyright-blue">广告刊例</a>】   -   【<a href="http://www.mydrivers.com/contact/privacy.html" target="_blank" class="copyright-blue">隐私权政策</a>】   -   【<a href="http://www.mydrivers.com/contact/jobs.html" target="_blank" class="copyright-blue">诚聘精英</a>】   -   【<a href="http://rss.mydrivers.com" target="_blank" class="copyright-blue">RSS订阅</a>】</td>
      </tr>
    </table></td>
  </tr>
  <tr>
    <td width="750" height="48" align="center" bgcolor="#F1F1F1"><table width="100%" border="0" align="center" cellpadding="0" cellspacing="0">
        <tr>
          <td height="18" align="center" class="f12-copyright">电信与信息服务业务经营许可证：<a href="http://www.miibeian.gov.cn/" target="_blank" class="newstiao4">京ICP备11024344号</a> 　驱动之家·版权所有</td>
        </tr>
    </table></td>
    <td width="150" align="right" bgcolor="#F1F1F1"><table width="150" border="0" cellpadding="0" cellspacing="0">
      <tr>
        <td width="48"><a href=http://www.zzwljc.com><IMG SRC="http://icons.mydrivers.com/www/jt.gif" WIDTH=48 HEIGHT=48 border="0"></a></td>
        <td width="102"><img src="http://icons.mydrivers.com/www/index-banquan_03.gif" width=102 height=48 alt="" /></td>
        </tr>
    </table></td>
  </tr>
  <tr>
    <td height="19" colspan="2" align="right" bgcolor="#D5D5D5"></td>
  </tr>
</table>
<script type="text/javascript"> 
var _bdhmProtocol = (("https:" == document.location.protocol) ? " https://" : " http://");
document.write(unescape("%3Cscript src='" + _bdhmProtocol + "hm.baidu.com/h.js%3Ffa993fdd33f32c39cbb6e7d66096c422' type='text/javascript'%3E%3C/script%3E"));
</script>
<script src="http://www.google-analytics.com/urchin.js" type="text/javascript"></script>
<script type="text/javascript">
_uacct = "UA-2034714-3";
urchinTracker();
</script>
</body>
</html>

`
