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
	"github.com/pubgo/errors"
	"github.com/pubgo/go-readability"
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

	readability.FromReader(strings.NewReader(_html), "https://mp.weixin.qq.com/s/Hwp62XMBvMT0NAk3RM928w")
}

const _html = `
<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width,initial-scale=1.0,maximum-scale=1.0,user-scalable=0,viewport-fit=cover">
<link rel="shortcut icon" type="image/x-icon" href="//res.wx.qq.com/a/wx_fed/assets/res/NTI4MWU5.ico">
<link rel="mask-icon" href="//res.wx.qq.com/a/wx_fed/assets/res/MjliNWVm.svg" color="#4C4C4C">
<link rel="apple-touch-icon-precomposed" href="//res.wx.qq.com/a/wx_fed/assets/res/OTE0YTAw.png">
<meta name="apple-mobile-web-app-capable" content="yes">
<meta name="apple-mobile-web-app-status-bar-style" content="black">
<meta name="format-detection" content="telephone=no">


        <script nonce="1977935583" type="text/javascript">
            window.logs = {
                pagetime: {}
            };
            window.logs.pagetime['html_begin'] = (+new Date());
        </script>
        <title>
</title>
        
<style>.radius_avatar{display:inline-block;background-color:#fff;padding:3px;border-radius:50%;-moz-border-radius:50%;-webkit-border-radius:50%;overflow:hidden;vertical-align:middle}.radius_avatar img{display:block;width:100%;height:100%;border-radius:50%;-moz-border-radius:50%;-webkit-border-radius:50%;background-color:#eee}.rich_media_inner{word-wrap:break-word;-webkit-hyphens:auto;-ms-hyphens:auto;hyphens:auto}.rich_media_area_primary{padding:20px 16px 12px;padding:calc(20px + constant(safe-area-inset-top)) calc(16px + constant(safe-area-inset-right)) 12px calc(16px + constant(safe-area-inset-left));padding:calc(20px + env(safe-area-inset-top)) calc(16px + env(safe-area-inset-right)) 12px calc(16px + env(safe-area-inset-left));background-color:#fafafa}.rich_media_area_primary.voice{padding-top:66px}.rich_media_area_primary .weui-loadmore_line .weui-loadmore__tips{color:rgba(0,0,0,0.3);background-color:#fafafa}.rich_media_area_extra{padding:0 16px 16px;padding:0 calc(16px + constant(safe-area-inset-right)) calc(16px + constant(safe-area-inset-bottom)) calc(16px + constant(safe-area-inset-left));padding:0 calc(16px + env(safe-area-inset-right)) calc(16px + env(safe-area-inset-bottom)) calc(16px + env(safe-area-inset-left))}.rich_media_extra{padding-top:30px}.mpda_bottom_container .rich_media_extra{padding-top:24px}.mpda_bottom_container .rich_media_extra .mpad_more_list{right:-10px}html{-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%;line-height:1.6}body{-webkit-touch-callout:none;font-family:-apple-system-font,BlinkMacSystemFont,"Helvetica Neue","PingFang SC","Hiragino Sans GB","Microsoft YaHei UI","Microsoft YaHei",Arial,sans-serif;color:#333;background-color:#f2f2f2;letter-spacing:.034em}h1,h2,h3,h4,h5,h6{font-weight:400;font-size:16px}*{margin:0;padding:0}a{color:#576b95;text-decoration:none;-webkit-tap-highlight-color:rgba(0,0,0,0)}.rich_media_title{font-size:22px;line-height:1.4;margin-bottom:14px}@supports(-webkit-overflow-scrolling:touch){.rich_media_title{font-weight:700}}.rich_media_meta_list{margin-bottom:22px;line-height:20px;font-size:0;word-wrap:break-word;word-break:break-all}.rich_media_meta_list em{font-style:normal}.rich_media_meta{display:inline-block;vertical-align:middle;margin:0 10px 10px 0;font-size:15px;-webkit-tap-highlight-color:rgba(0,0,0,0)}.rich_media_meta.icon_appmsg_tag{margin-right:4px}.rich_media_meta.meta_tag_text{margin-right:0}.rich_media_meta_primary{display:block;margin-bottom:10px;font-size:15px}.meta_original_tag{padding:0 .5em;font-size:12px;line-height:1.4;background-color:#f2f2f2;color:#888}.meta_enterprise_tag img{width:30px;height:30px!important;display:block;position:relative;margin-top:-3px;border:0}.rich_media_meta_link{color:#576b95}.rich_media_meta_text{color:rgba(0,0,0,0.3)}.rich_media_meta_text.rich_media_meta_split{padding-left:10px}.rich_media_meta_text.rich_media_meta_split:before{position:absolute;top:50%;left:0;margin-top:-6px;content:' ';display:block;border-left:1px solid #888;width:200%;height:130%;box-sizing:border-box;-moz-box-sizing:border-box;-webkit-box-sizing:border-box;-webkit-transform:scale(0.5);transform:scale(0.5);-webkit-transform-origin:0 0;transform-origin:0 0}.rich_media_meta_text.article_modify_tag{position:relative}.rich_media_meta_nickname{position:relative}.rich_media_thumb_wrp{margin-bottom:6px}.rich_media_thumb_wrp .original_img_wrp{display:block}.rich_media_thumb{display:block;width:100%}.rich_media_content{overflow:hidden;color:#333;font-size:17px;word-wrap:break-word;-webkit-hyphens:auto;-ms-hyphens:auto;hyphens:auto;text-align:justify;position:relative;z-index:0}.rich_media_content *{max-width:100%!important;box-sizing:border-box!important;-webkit-box-sizing:border-box!important;word-wrap:break-word!important}.rich_media_content p{clear:both;min-height:1em}.rich_media_content em{font-style:italic}.rich_media_content fieldset{min-width:0}.rich_media_content .list-paddingleft-1,.rich_media_content .list-paddingleft-2,.rich_media_content .list-paddingleft-3{padding-left:2.2em}.rich_media_content .list-paddingleft-1 .list-paddingleft-2,.rich_media_content .list-paddingleft-2 .list-paddingleft-2,.rich_media_content .list-paddingleft-3 .list-paddingleft-2{padding-left:30px}.rich_media_content .list-paddingleft-1{padding-left:1.2em}.rich_media_content .list-paddingleft-3{padding-left:3.2em}.rich_media_content .code-snippet,.rich_media_content .code-snippet__fix{max-width:1000%!important}.rich_media_content .code-snippet *,.rich_media_content .code-snippet__fix *{max-width:1000%!important}.rich_media_content img{height:auto!important}@media screen and (min-width:1024px){.rich_media_area_primary_inner,.rich_media_area_extra_inner{max-width:677px;margin-left:auto;margin-right:auto}.rich_media_area_primary{padding-top:32px}}blockquote{padding-left:10px;border-left:3px solid #dbdbdb;color:rgba(0,0,0,0.5);font-size:15px;padding-top:4px;margin:1em 0}.blockquote_info{color:rgba(0,0,0,0.3);margin-top:1.1764705882352942em;word-wrap:break-word;-webkit-hyphens:auto;-ms-hyphens:auto;hyphens:auto}.blockquote_article{display:block}.appmsg_share_notice{font-size:16px;color:#888;position:relative;padding:1.25em 0;margin-bottom:1.75em}.appmsg_share_notice:before{content:" ";position:absolute;left:0;top:0;right:0;height:1px;border-top:1px solid #dfdfdf;-webkit-transform-origin:0 0;transform-origin:0 0;-webkit-transform:scaleY(0.5);transform:scaleY(0.5)}.appmsg_share_notice:after{content:" ";position:absolute;left:0;bottom:0;right:0;height:1px;border-bottom:1px solid #dfdfdf;-webkit-transform-origin:0 100%;transform-origin:0 100%;-webkit-transform:scaleY(0.5);transform:scaleY(0.5)}.appmsg_share_notice_hd{font-weight:700;padding-bottom:.2em}.dn{display:none}.qa__card{font-size:15px;margin:20px 0;background-color:rgba(0,0,0,0.03)}.qa__card .qa__item{padding:24px 0 16px 16px}.qa__card .qa__item-question{padding-right:16px}.qa__card .qa__item-question-info{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;color:rgba(0,0,0,0.7);margin-bottom:9px;font-size:15px}.qa__card .qa__item-info-avatar{width:24px;height:24px;border-radius:2px;overflow:hidden;margin-right:8px}.qa__card .qa__item-info-avatar img{width:100%}.qa__card .qa__item-info-account{margin-right:.5em}.qa__card .qa__item-question{margin-bottom:24px}.qa__card .qa__answers-question-title{font-size:15px;margin-bottom:8px}.qa__card .qa__detail-question-content{color:rgba(0,0,0,0.5);font-size:15px}.qa__card .qa__item-reply-content{padding-right:16px}.qa__card .qa__item-reply-head{color:rgba(0,0,0,0.6);position:relative;padding-left:10px;margin-bottom:9px}.qa__card .qa__item-reply-head::before{content:"";display:block;width:2px;height:15px;background-color:#1aad19;position:absolute;left:0;top:50%;-webkit-transform:translateY(-50%);transform:translateY(-50%)}.qa__card .qa__detail-question-imgs{margin-top:4px;color:#576b95}.qa__card .qa__show_detail{position:relative;padding:16px 0;display:block;margin-top:16px;margin-bottom:-16px;color:#576b95}.qa__card .qa__show_detail::before{content:'';position:absolute;display:block;width:100%;border-top:1px solid rgba(0,0,0,0.1);-webkit-transform:scaleY(0.5);transform:scaleY(0.5);left:0;top:0;-webkit-transform-origin:left top;transform-origin:left top}.qa__card .qa__show_detail::after{content:'';display:block;width:8px;height:8px;border-top:2px solid rgba(0,0,0,0.3);border-right:2px solid rgba(0,0,0,0.3);-webkit-transform:rotate(45deg) translateY(-50%);transform:rotate(45deg) translateY(-50%);position:absolute;right:16px;top:50%;border-radius:2px;-webkit-transform-origin:center;transform-origin:center}.qa__card-empty{line-height:115px;color:rgba(0,0,0,0.3);text-align:center}.code-snippet{margin:10px 0;display:block;overflow-x:auto;font-size:14px;padding:1em 1em 1em 3em;color:#333;position:relative;background-color:#fafafa;border:1px solid #f0f0f0;border-radius:2px;counter-reset:line;white-space:normal;-webkit-overflow-scrolling:touch}.code-snippet code{text-align:left;font-size:14px;display:block;white-space:pre-wrap;position:relative;font-family:Consolas,"Liberation Mono",Menlo,Courier,monospace}.code-snippet code::before{position:absolute;min-width:1.5em;text-align:right;left:-2.5em;counter-increment:line;content:counter(line);display:inline;margin-right:12px;color:rgba(0,0,0,0.15)}.code-snippet_nowrap code{white-space:pre;display:-webkit-box;display:-webkit-flex;display:flex}.code-snippet__fix{font-size:14px;margin:10px 0;display:block;color:#333;position:relative;background-color:rgba(0,0,0,0.03);border:1px solid #f0f0f0;border-radius:2px;display:-webkit-box;display:-webkit-flex;display:flex;line-height:26px}.code-snippet__fix pre{overflow-x:auto;padding:1em;padding-left:0;white-space:normal;-webkit-box-flex:1;-webkit-flex:1;flex:1;-webkit-overflow-scrolling:touch}.code-snippet__fix code{text-align:left;font-size:14px;display:block;white-space:pre;display:-webkit-box;display:-webkit-flex;display:flex;position:relative;font-family:Consolas,"Liberation Mono",Menlo,Courier,monospace}.code-snippet__fix .code-snippet__line-index{counter-reset:line;-webkit-flex-shrink:0;flex-shrink:0;height:100%;padding:1em;list-style-type:none}.code-snippet__fix .code-snippet__line-index li{list-style-type:none;text-align:right}.code-snippet__fix .code-snippet__line-index li::before{min-width:1.5em;text-align:right;left:-2.5em;counter-increment:line;content:counter(line);display:inline;color:rgba(0,0,0,0.15)}.code-snippet__comment,.code-snippet__quote{color:#afafaf;font-style:italic}.code-snippet__keyword,.code-snippet__selector-tag,.code-snippet__subst{color:#ca7d37}.code-snippet__number,.code-snippet__literal,.code-snippet__variable,.code-snippet__template-variable,.code-snippet__tag .code-snippet__attr{color:#0e9ce5}.code-snippet__string,.code-snippet__doctag{color:#d14}.code-snippet__title,.code-snippet__section,.code-snippet__selector-id{color:#d14}.code-snippet__subst{font-weight:normal}.code-snippet__type,.code-snippet__class .code-snippet__title{color:#0e9ce5}.code-snippet__tag,.code-snippet__name,.code-snippet__attribute{color:#0e9ce5;font-weight:normal}.code-snippet__regexp,.code-snippet__link{color:#ca7d37}.code-snippet__symbol,.code-snippet__bullet{color:#d14}.code-snippet__built_in,.code-snippet__builtin-name{color:#ca7d37}.code-snippet__meta{color:#afafaf}.code-snippet__deletion{background:#fdd}.code-snippet__addition{background:#dfd}.code-snippet__emphasis{font-style:italic}.code-snippet__strong{font-weight:bold}.cell{padding:.8em 0;display:block;position:relative}.cell_hd,.cell_bd,.cell_ft{display:table-cell;vertical-align:middle;word-wrap:break-word;word-break:break-all;white-space:nowrap}.cell_primary{width:2000px;white-space:normal}.flex_cell{padding:10px 0;display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center}.flex_cell_primary{width:100%;-webkit-box-flex:1;-webkit-flex:1;box-flex:1;flex:1}.original_tool_area{display:block;padding:.75em 1em 0;-webkit-tap-highlight-color:rgba(0,0,0,0);color:#333;border:1px solid #eaeaea;margin:20px 0}.original_tool_area .tips_global{position:relative;padding-bottom:.5em;font-size:15px}.original_tool_area .tips_global:after{content:" ";position:absolute;left:0;bottom:0;right:0;height:1px;border-bottom:1px solid #dbdbdb;-webkit-transform-origin:0 100%;transform-origin:0 100%;-webkit-transform:scaleY(0.5);transform:scaleY(0.5)}.original_tool_area .radius_avatar{width:27px;height:27px;padding:0;margin-right:.5em}.original_tool_area .radius_avatar img{height:100%!important}.original_tool_area .flex_cell_bd{width:auto;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;word-wrap:normal}.original_tool_area .flex_cell_ft{font-size:14px;color:#888;padding-left:1em;white-space:nowrap}.original_tool_area .icon_access:after{content:" ";display:inline-block;height:8px;width:8px;border-width:1px 1px 0 0;border-color:#cbcad0;border-style:solid;transform:matrix(0.71,0.71,-0.71,0.71,0,0);-ms-transform:matrix(0.71,0.71,-0.71,0.71,0,0);-webkit-transform:matrix(0.71,0.71,-0.71,0.71,0,0);position:relative;top:-2px;top:-1px}.rich_media_global_msg{position:fixed;top:0;left:0;right:0;padding:.85em 35px .85em 15px;z-index:2;background-color:#c6e0f8;color:#888;font-size:12px}.rich_media_global_msg .icon_closed{position:absolute;right:15px;top:50%;margin-top:-5px;line-height:300px;overflow:hidden;-webkit-tap-highlight-color:rgba(0,0,0,0);background:transparent url(//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/appmsg/icon_appmsg_msg_closed_sprite.2x42f400.png) no-repeat 0 0;width:11px;height:11px;vertical-align:middle;display:inline-block;background-size:100% auto}.rich_media_global_msg .icon_closed:active{background-position:0 -17px}.rich_media_global_msg.voice{color:#1aad19;background-color:#e8f6e8;padding-left:43.3px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap}.rich_media_global_msg.voice .ic_voice{position:absolute;top:50%;margin-top:-10px;left:15px;display:inline-block;width:13.3px;height:18.3px;background:url(data:image/png;base64,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) no-repeat center;background-size:contain}.rich_media_global_msg.voice .icon_more{position:absolute;right:15px;top:50%;margin-top:-6.5px;width:8px;height:13px;background:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAmCAYAAADeB1slAAAAAXNSR0IArs4c6QAAA49JREFUSA2tVk1IVFEUPveOmhQk1ibfvLFZSJtACqJIIpCIoB+imcFV2Z9ZUVnRomgRGFn2i6mFf5XVotCZMYmgIMhlq6BFENFiwjfPAosgFzXOvNu5o3fmeue98Y16F3PP+c453/fu/xCQmvZSWwwJaC0G6P8WND9IoTmbVFRycpqAdgBWkwRo94V960RsPn1aoPpN9RI6CR0WsDWcjAErtUiqTR/S18+HnNcSTv5zYrwDSatVMgLkH6aciwfj79WYW5+OT4xfsyPnBIgvwt872qBW45ZQzaMlJXCfMPJbDQgfRUoIhduVUW2TwArpaWyX+bnYU9JAAH44FaJIccqCW76Ib7NTjhOOvFPN/8q/IvE3cR+npFJgao9rkiSUXjD2GCNqzMnPbNPYjth3mqINlJAvTsk4kiJmWa3eQe8WpxwVzwjwgFFn/CorL29EkY9qovC5CFDrKq7JVoHl6zNTJCf53/lLk78SN/FcbJRx2cZCixC4ZATGXsu4as8YgQjGamN/q5atOotn4K3A1J4BUMbgsh7Wt6sx2bcV4AkjtSPJxkDjRQL0hVwg21zEglSzN+rdKeOybTtFcgK39ajWZDFWr+IZHxfGQ8mV0YA5nMGmDVcCPNcb8R5gYJ1UCWQfp6PFCI4NyZhrAV6EIiHGrPOAB0ImkW2c0la8u8ICc0wUCWrvC1dsw+3TnN6uanDaJ4zeiIfiA9wtWIAX8XspxeA6iuBlaN88hNzCNXk+JwFOuXJYW5tMQmc+EcqgzXGb2n9XFk1NstX5yKczl89JAF+6vRaBM1m5XAu//qkRGrtb8BThTtqH2/V0LmUWwbvsiREw8X0vcJHxxNbjNm3KUuVaMjmPFuWm2CPesHc/kp+yj06hFMhj/PIOOcfVFLk6xZT0G3vMTpmc27MusityZk/OBfKOwA05MPLIDJn3OJldcxTQotpBYOyEXVEWIw/NoInvuHOznSJfVDs0GzneNw9mI+eyOSPQI9phfCqPO38TFhHaFw/Eu/LliNiMESB5w6zkzD05F8mMYJr8mFC263Fa+vAadvXloj4toEcrjlgMjgrQrseHpBcfkm67WD6M6MN6lTWZeoZjyYxGLcB57MGnsEfF3fjU2G189QBpcUqmBLrnSs4504s8GjLxrwm5qYrg3dKFf6x6VbwQ3yOS/wz8+bS0ruwfELaBY2nyoNkn4gvW65GKRn4WForwP+dONHDaOHeZAAAAAElFTkSuQmCC) no-repeat center;background-size:contain}.preview_appmsg .rich_media_title{margin-top:2.3em}@media screen and (min-width:1024px){.rich_media_global_msg{position:relative;margin-bottom:32px}.preview_appmsg .rich_media_title.rich_media_title{margin-top:0}}.pages_reset{color:#333;line-height:1.6;font-size:16px;font-weight:400;font-style:normal;text-indent:0;letter-spacing:normal;text-align:left;text-decoration:none;white-space:normal}.weapp_element,.weapp_display_element,.mp-miniprogram{display:block;margin:1em 0}.share_audio_context{margin:16px 0}.weapp_text_link{font-size:17px}.weapp_text_link:before{content:'';display:inline-block;line-height:1;background-size:contain;background-repeat:no-repeat;background-image:url('data:image/png;base64,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');vertical-align:middle;font-size:11px;color:#888;margin-right:6px;margin-top:-3px;background-position:center;height:20px;width:20px}.flex_context{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center}.flex_bd{-webkit-box-flex:1;-webkit-flex:1;flex:1;word-wrap:break-word;word-break:break-all}.original_page{background-color:#fff;font-size:16px}.account_info{padding:0 0 20px}.account_info .flex_bd{padding-left:.85em}.radius_avatar.account_avatar{width:40px;height:40px;padding:0}.account_nickname{overflow:hidden;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical;-webkit-line-clamp:1;line-height:1.2;color:#576b95;font-size:14px}.account_nickname_inner{font-weight:400;vertical-align:top}.account_desc{overflow:hidden;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical;-webkit-line-clamp:1;color:#b2b2b2;font-size:13px;line-height:1.2;padding-top:.3em}.account_desc_inner{display:inline;vertical-align:top}.share_notice{margin-bottom:20px;word-wrap:break-word;-webkit-hyphens:auto;-ms-hyphens:auto;hyphens:auto}.share_media{padding-bottom:18px}.original_panel{padding:20px;background-color:#fcfcfc;word-wrap:break-word;-webkit-hyphens:auto;-ms-hyphens:auto;hyphens:auto;overflow:hidden;position:relative}.original_panel:after{content:" ";border:1px solid #d8d8d8;border-radius:4px;-moz-border-radius:4px;-webkit-border-radius:4px;position:absolute;top:0;left:0;width:200%;height:200%;-webkit-transform:scale(0.5);transform:scale(0.5);-webkit-transform-origin:0 0;transform-origin:0 0;box-sizing:border-box;-moz-box-sizing:border-box;-webkit-box-sizing:border-box}.original_panel .original_account{margin-bottom:18px;position:relative;z-index:1}.original_panel .original_account_avatar{width:28px;height:28px;padding:0}.original_panel .original_account_nickname{padding-left:.85em;font-size:15px;color:#576b95}.original_panel_title{font-size:23px;color:#000;margin:0 0 16px}.original_panel_content{color:#333}.original_panel_tool{padding-top:20px;position:relative;z-index:1}.original_tool_area.disabled{padding:0}.original_tool_area.disabled .flex_cell_bd{-webkit-box-flex:initial;-webkit-flex:initial;box-flex:initial;flex:initial}.original_tool_area.disabled .flex_cell_ft{-webkit-box-flex:1;-webkit-flex:1;box-flex:1;flex:1;padding:0}.original_tool_area.disabled .original_tool_tips:after{content:"："}.original_tool_area.disabled .original_cell_nickname{color:rgba(0,0,0,0.3);overflow:hidden;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical;-webkit-line-clamp:1;white-space:normal;max-width:none;min-height:1.6em}.original_tool_area.disabled:before,.original_tool_area.disabled:after,.original_tool_area.disabled .radius_avatar,.original_tool_area.disabled .icon_access:after{display:none}.original_cell_nickname{display:inline-block;vertical-align:middle;font-weight:400;width:auto;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;word-wrap:normal;max-width:8em;color:#000}.original_tool_area{padding:1.17em 0;border-width:0;position:relative}.original_tool_area:before{content:" ";position:absolute;left:0;top:0;right:0;height:1px;border-top:1px solid #dfdfdf;-webkit-transform-origin:0 0;transform-origin:0 0;-webkit-transform:scaleY(0.5);transform:scaleY(0.5)}.original_tool_area:after{content:" ";position:absolute;left:0;bottom:0;right:0;height:1px;border-bottom:1px solid #dfdfdf;-webkit-transform-origin:0 100%;transform-origin:0 100%;-webkit-transform:scaleY(0.5);transform:scaleY(0.5)}.original_tool_area .radius_avatar{width:20px;height:20px;margin-right:.2em}.original_tool_area .flex_cell{padding:0;font-size:14px}.original_tool_area .icon_access:after{margin-right:4px;top:0}.preview_appmsg .rich_media_title{margin-top:1.5em}.preview_appmsg .account_info{padding-top:3em}.original_page{background-color:transparent}.account_info{-webkit-tap-highlight-color:rgba(0,0,0,0);outline:0;padding-bottom:20px;font-size:14px}.account_info.appmsg_account_info{padding-bottom:32px}.account_info .radius_avatar img{background-color:transparent}.share_notice{font-size:16px;word-wrap:break-word;-webkit-hyphens:auto;-ms-hyphens:auto;hyphens:auto}.original_panel{background-color:#fafafa}.original_panel:after{border-color:#e6e6e6}.original_panel .original_account_avatar{width:30px;height:30px}.original_panel_tool{font-size:14px}.original_panel_tool a{color:#576b95}.original_panel_content{font-size:17px}.share_media{padding-bottom:30px}.icon_appmsg_tag{display:inline-block;vertical-align:middle;padding:0 .5em;font-size:12px;line-height:1.5;background-color:#c3c3c3;color:#fff;border-radius:2px;-moz-border-radius:2px;-webkit-border-radius:2px;width:auto;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;word-wrap:normal;max-width:70%}.icon_appmsg_tag.primary{color:#3bb638;padding:4px .78em;background-color:rgba(9,187,7,0.08);font-size:12px;border-top-left-radius:.95em 50%;-moz-border-radius-topleft:.95em 50%;-webkit-border-top-left-radius:.95em 50%;border-top-right-radius:.95em 50%;-moz-border-radius-topright:.95em 50%;-webkit-border-top-right-radius:.95em 50%;border-bottom-left-radius:.95em 50%;-moz-border-radius-bottomleft:.95em 50%;-webkit-border-bottom-left-radius:.95em 50%;border-bottom-right-radius:.95em 50%;-moz-border-radius-bottomright:.95em 50%;-webkit-border-bottom-right-radius:.95em 50%}.icon_appmsg_tag.default{border:1px solid rgba(0,0,0,0.1);color:rgba(0,0,0,0.3);background-color:transparent;padding:0 .54em;font-size:15px;line-height:1.3;border-top-left-radius:.67em 50%;-moz-border-radius-topleft:.67em 50%;-webkit-border-top-left-radius:.67em 50%;border-top-right-radius:.67em 50%;-moz-border-radius-topright:.67em 50%;-webkit-border-top-right-radius:.67em 50%;border-bottom-left-radius:.67em 50%;-moz-border-radius-bottomleft:.67em 50%;-webkit-border-bottom-left-radius:.67em 50%;border-bottom-right-radius:.67em 50%;-moz-border-radius-bottomright:.67em 50%;-webkit-border-bottom-right-radius:.67em 50%}.rich_media_meta_list .icon_appmsg_tag.default{margin-top:-1px}.icon_appmsg_tag.title_tag{background-color:#d04b4e}.icon_global_tag_wrp{text-align:right;padding-bottom:12px}.icon_global_tag{background-color:rgba(118,118,118,0.16);color:rgba(0,0,0,0.41);line-height:2.2;border-top-left-radius:1em 50%;-moz-border-radius-topleft:1em 50%;-webkit-border-top-left-radius:1em 50%;border-bottom-left-radius:1em 50%;-moz-border-radius-bottomleft:1em 50%;-webkit-border-bottom-left-radius:1em 50%;padding:0 1.8em 0 1.34em;font-size:12px;margin-right:-24px;display:inline-block;vertical-align:top}.global_plain_btn{display:inline-block;vertical-align:middle;padding:0 1em;line-height:2;font-size:14px;-webkit-tap-highlight-color:rgba(0,0,0,0);border-radius:5px;-moz-border-radius:5px;-webkit-border-radius:5px}.global_plain_btn.primary{color:#1aad19;border:1px solid currentColor}.global_plain_btn.primary:active{color:rgba(26,173,25,0.6)}.wx_video_context{color:#fefefe;position:relative;background-color:#000}.wx_video_thumb,.wx_video_thumb_primary{position:absolute;left:0;width:100%;height:100%!important;top:0}.wx_video_thumb_primary{background-size:cover;background-position:50% 50%;background-repeat:no-repeat}.wx_video_play_btn{position:absolute;top:50%;left:50%;-webkit-transform:translate(-50%,-50%);transform:translate(-50%,-50%);margin-top:-2px;font-size:0;border-width:0;background-color:transparent;padding:0;outline:0;z-index:2}.wx_video_play_btn:before{content:" ";display:inline-block;width:0;height:0;border-width:14px 0 14px 25px;border-color:transparent transparent transparent #fff;border-style:dotted dotted dotted solid}.wx_video_mask{position:absolute;top:0;left:0;right:0;bottom:0;z-index:1;background-color:rgba(0,0,0,0.5)}.place_audio_area{min-height:100px;background-color:#fdfdfd;display:block;margin:16px 0}.place_music_area{min-height:68px;background-color:#fdfdfd;display:block;margin:17px 0 16px}.rich_media_empty_extra{background-color:#fafafa}.appmsg_skin_default.rich_media_empty_extra{background-color:#fff}.appmsg_skin_default .rich_media_area_primary{background-color:#fff}.appmsg_skin_default .rich_media_area_primary .weui-loadmore_line .weui-loadmore__tips{background-color:#fff}.appmsg_style_default .rich_media_tool{padding-top:15px}.rich_media_title_ios{font-weight:700}.my_comment_empty_data{background-color:#fff}.read-more__area{margin:30px 0}.read-more__desc{margin-bottom:10px}.read-more__article__extra{font-size:14px;color:rgba(0,0,0,0.5)}.read-more__article__item{margin-bottom:10px}</style>

<style>
        </style>
<!--[if lt IE 9]>
<link rel="stylesheet" type="text/css" href="//res.wx.qq.com/mmbizwap/zh_CN/htmledition/style/page/appmsg_new/pc46b604.css">
<![endif]-->

    </head>
    <body id="activity-detail" class="zh_CN mm_appmsg  appmsg_skin_default appmsg_style_default ">
        
<script nonce="1977935583" type="text/javascript">
    var biz = ""||"MzU4MjQ0MTU4Ng==";
    var sn = "" || ""|| "cb013878dad4a29a9a354d12a5135e4b";
    var mid = "" || ""|| "2247484319";
    var idx = "" || "" || "1";
    var LANG= "zh_CN";
    window.__allowLoadResFromMp = true; 
    
        
</script>
<script nonce="1977935583" type="text/javascript">
var page_begintime=+new Date,is_rumor="",norumor="";
1*is_rumor&&!(1*norumor)&&biz&&mid&&(document.referrer&&-1!=document.referrer.indexOf("mp.weixin.qq.com/mp/rumor")||(location.href="http://mp.weixin.qq.com/mp/rumor?action=info&__biz="+biz+"&mid="+mid+"&idx="+idx+"&sn="+sn+"#wechat_redirect"));
</script>
<script nonce="1977935583" type="text/javascript">
var MutationObserver=window.WebKitMutationObserver||window.MutationObserver||window.MozMutationObserver,isDangerSrc=function(t){
if(t){
var e=t.match(/http(?:s)?:\/\/([^\/]+?)(\/|$)/);
if(e&&!/qq\.com(\:8080)?$/.test(e[1])&&!/weishi\.com$/.test(e[1]))return!0;
}
return!1;
},ishttp=0==location.href.indexOf("http://");
-1==location.href.indexOf("safe=0")&&ishttp&&"function"==typeof MutationObserver&&"mp.weixin.qq.com"==location.host&&(window.__observer_data={
count:0,
exec_time:0,
list:[]
},window.__observer=new MutationObserver(function(t){
window.__observer_data.count++;
var e=new Date,r=[];
t.forEach(function(t){
for(var e=t.addedNodes,o=0;o<e.length;o++){
var n=e[o];
if("SCRIPT"===n.tagName){
var i=n.src;
isDangerSrc(i)&&(window.__observer_data.list.push(i),r.push(n)),!i&&window.__nonce_str&&n.getAttribute("nonce")!=window.__nonce_str&&(window.__observer_data.list.push("inlinescript_without_nonce"),
r.push(n));
}
}
});
for(var o=0;o<r.length;o++){
var n=r[o];
n.parentNode&&n.parentNode.removeChild(n);
}
window.__observer_data.exec_time+=new Date-e;
}),window.__observer.observe(document,{
subtree:!0,
childList:!0
})),function(){
if(-1==location.href.indexOf("safe=0")&&Math.random()<.01&&ishttp&&HTMLScriptElement.prototype.__lookupSetter__&&"undefined"!=typeof Object.defineProperty){
window.__danger_src={
xmlhttprequest:[],
script_src:[],
script_setAttribute:[]
};
var t="$"+Math.random();
HTMLScriptElement.prototype.__old_method_script_src=HTMLScriptElement.prototype.__lookupSetter__("src"),
HTMLScriptElement.prototype.__defineSetter__("src",function(t){
t&&isDangerSrc(t)&&window.__danger_src.script_src.push(t),this.__old_method_script_src(t);
});
var e="element_setAttribute"+t;
Object.defineProperty(Element.prototype,e,{
value:Element.prototype.setAttribute,
enumerable:!1
}),Element.prototype.setAttribute=function(t,r){
"SCRIPT"==this.tagName&&"src"==t&&isDangerSrc(r)&&window.__danger_src.script_setAttribute.push(r),
this[e](t,r);
};
}
}();
</script>

        <link rel="dns-prefetch" href="//res.wx.qq.com">
<link rel="dns-prefetch" href="//mmbiz.qpic.cn">
<link rel="shortcut icon" type="image/x-icon" href="//res.wx.qq.com/a/wx_fed/assets/res/NTI4MWU5.ico">
<link rel="mask-icon" href="//res.wx.qq.com/a/wx_fed/assets/res/MjliNWVm.svg" color="#4C4C4C">
<link rel="apple-touch-icon-precomposed" href="//res.wx.qq.com/a/wx_fed/assets/res/OTE0YTAw.png">
<script nonce="1977935583" type="text/javascript">
    String.prototype.html = function(encode) {
        var replace =["&#39;", "'", "&quot;", '"', "&nbsp;", " ", "&gt;", ">", "&lt;", "<", "&yen;", "¥", "&amp;", "&"];
        
		
		
		
		
        
        var replaceReverse = ["&", "&amp;", "¥", "&yen;", "<", "&lt;", ">", "&gt;", " ", "&nbsp;", '"', "&quot;", "'", "&#39;"];
	    var target;
	    if (encode) {
	    	target = replaceReverse;
	    } else {
	    	target = replace;
	    }
        for (var i=0,str=this;i< target.length;i+= 2) {
             str=str.replace(new RegExp(target[i],'g'),target[i+1]);
        }
        return str;
    };

    window.isInWeixinApp = function() {
        return /MicroMessenger/.test(navigator.userAgent);
    };

    window.getQueryFromURL = function(url) {
        url = url || 'http://qq.com/s?a=b#rd'; 
        var tmp = url.split('?'),
            query = (tmp[1] || "").split('#')[0].split('&'),
            params = {};
        for (var i=0; i<query.length; i++) {
            var arg = query[i].split('=');
            params[arg[0]] = arg[1];
        }
        if (params['pass_ticket']) {
        	params['pass_ticket'] = encodeURIComponent(params['pass_ticket'].html(false).html(false).replace(/\s/g,"+"));
        }
        return params;
    };

    (function() {
	    var params = getQueryFromURL(location.href);
        window.uin = params['uin'] || "" || '';
        window.key = params['key'] || "" || '';
        window.wxtoken = params['wxtoken'] || '';
        window.pass_ticket = params['pass_ticket'] || '';
        window.appmsg_token = "";
    })();

    function wx_loaderror() {
        if (location.pathname === '/bizmall/reward') {
            new Image().src = '/mp/jsreport?key=96&content=reward_res_load_err&r=' + Math.random();
        }
    }

</script>

<script nonce="1977935583" type="text/javascript">
        window.__moon_report_uin = "0";
            window.no_moon_ls = 0;
    </script>

        
<script nonce="1977935583" type="text/javascript">
    var write_sceen_time = (+new Date());
</script>

<div id="js_article" class="rich_media">
    
    <div id="js_top_ad_area" class="top_banner"></div>
    
    <div class="rich_media_inner">

        
        
		<div id="page-content" class="rich_media_area_primary">

		  <div class="rich_media_area_primary_inner">

            
                          
                        

            <div id="img-content">

                
                <h2 class="rich_media_title" id="activity-name">
                    
                    
                    
            Istio多集群管理方案详解
                      </h2>
                <div id="meta_content" class="rich_media_meta_list">
                                                                                
                                        <span class="rich_media_meta rich_media_meta_nickname" id="profileBt">
                      <a href="javascript:void(0);" id="js_name">
                        k8s技术圈                      </a>
                      <div id="js_profile_qrcode" class="profile_container" style="display:none;">
                          <div class="profile_inner">
                              <strong class="profile_nickname">k8s技术圈</strong>
                              <img class="profile_avatar" id="js_profile_qrcode_img" src="" alt="">

                              <p class="profile_meta">
                              <label class="profile_meta_label">微信号</label>
                              <span class="profile_meta_value">kube100</span>
                              </p>

                              <p class="profile_meta">
                              <label class="profile_meta_label">功能介绍</label>
                              <span class="profile_meta_value">专注容器、专注 kubernetes 技术......</span>
                              </p>
                              
                          </div>
                          <span class="profile_arrow_wrp" id="js_profile_arrow_wrp">
                              <i class="profile_arrow arrow_out"></i>
                              <i class="profile_arrow arrow_in"></i>
                          </span>
                      </div>
                    </span>


                    <em id="publish_time" class="rich_media_meta rich_media_meta_text"></em>





                </div>
                
                
                
                
                                                
                                                                
                                
                
                <div class="rich_media_content " id="js_content">
                    

                    

                    
                    
                    <p class="p1"><span style="font-size: 14px;font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;">随着Kubernetes成为云原生领域应用编排的标准，传统企业及新型互联网企业都在逐步将应用容器化及云化。为了实现高并发和高可用，企业通常会将应用部署在多集群甚至多云、混合云等多种环境中，因此，多集群方案逐步成为企业应用部署的最佳选择。很多云厂商都推出了自己的多云、混合云方案，虽然几乎都提供了多集群管理及跨集群的服务访问能力，但是在服务治理方面都有所欠缺。</span></p><p class="p1"><span style="font-size: 14px;font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;"><br  /></span></p><p style="text-align: center;"><a href="http://mp.weixin.qq.com/s?__biz=MzU4MjQ0MTU4Ng==&amp;mid=2247484223&amp;idx=1&amp;sn=879d99b07136047e534def13656618e3&amp;chksm=fdb90c22cace853489a891c05d8c89779785bd49ddf083734a0d9e5b2a1768081d07491a59b5&amp;scene=21#wechat_redirect" target="_blank" data-itemshowtype="0" data-linktype="1"><span class="js_jump_icon h5_image_link" data-positionback="static" style="top: auto;left: auto;margin: 0px;right: auto;bottom: auto;"><img class="rich_pages" data-ratio="0.4255555555555556" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_png/z9BgVMEm7YsbvT2Aia1aQOf56e01MiaTXAl0yicmia67ooK7rIiav3tic32lnZYUKA73DrMZib39N7LNxDAox8tvLGVvA/640?wx_fmt=png" data-type="png" data-w="900" style="margin: 0px;"  /></span></a></p><p class="p1"><span style="font-size: 14px;font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;"></span><br  /></p><p class="p3"><span style="font-size: 14px;">因此，越来越多的用户对跨集群的服务治理表现出强烈的兴趣和需求。在此背景下，Istio作为Service Mesh领域的事实标准，推出了三种多集群管理方案，这里讲解其中的多控制面管理方案。</span></p><p class="p3"><br  /></p><p class="p6"><span style="font-size: 14px;">多控制面拓扑模型是在Istio 1.1中新增的一种多集群管理模型，如图7-1所示，每个Kubernetes集群都分别部署自己独立的Istio控制面，并且每个集群的控制面部署形态都是相似的，都各自管理自身的Endpoint。</span></p><p class="p7"><span style="font-size: 14px;">&nbsp;</span></p><p style="text-align: center;"><img class="rich_pages" data-ratio="0.6086956521739131" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_png/z9BgVMEm7YsbvT2Aia1aQOf56e01MiaTXAEODRXjscszALHVl2Uia6URoDE51PmsgCqAfYbDxia4rzqnMG4y2AHQ5g/640?wx_fmt=png" data-type="png" data-w="1058" style=""  /></p><p class="p8" style="text-align: center;"><span style="font-size: 12px;">多控制面拓扑模型</span></p><p class="p2"><br  /></p><p class="p2"><span style="font-size: 14px;">多控制面模型还有以下特点。</span></p><ul class="ul1 list-paddingleft-2" style=""><li><p><span style="font-size: 14px;">共享根CA。为了支持安全的跨集群通信mTLS，该模型要求每个集群控制面都使用相同的中间CA证书，供Citadel签发证书使用，以支持跨集群的TLS双向认证。</span></p></li><li><p><span style="font-size: 14px;">不要求不同集群之间共享网络，即容器网络不需要打通，跨集群的访问通过Istio Gateway转发。</span></p></li><li><p><span style="font-size: 14px;">每个Kubernetes集群的Pod地址范围与服务地址范围都可以与其他集群重叠，双方集群互不干扰，因为每个集群的Istio控制面都只管理自己集群的Endpoint。</span></p></li><li><p><span style="font-size: 14px;">该模型依赖于DNS解析，允许服务实例解析本集群或者远端集群的服务名称。它除了使用了Kubernetes默认的cluster.local和&lt;namesapce&gt;.cluster.local后缀，还扩展了Pod的DNS解析，添加了对“.global”后缀的服务支持，以支持remote集群的服务地址解析。</span></p><p><span style="font-size: 14px;"></span></p><p><br  /></p></li></ul><p class="p2"><span style="font-size: 14px;">在多控制面Gateway直连模型中，每个工作负载都可以像单集群一样使用典型的Kubernetes服务域名访问同一集群内的服务。然而对于Remote集群的服务访问，Istio扩展了CoreDNS服务器，处理“&lt;name&gt;.&lt;namespace&gt;.global”形式的服务地址解析。</span></p><p class="p2"><span style="font-size: 14px;">在多控制面Gateway直连模型中，服务间的访问方式分为以下两种。</span></p><p class="p2"><span style="font-size: 14px;"><br  /></span></p><ul class="ul1 list-paddingleft-2" style=""><li><p><span style="font-size: 14px;">同一集群内部的服务访问。这种访问方式与单集群模型没有任何区别。</span></p></li><li><p><span style="font-size: 14px;">跨集群的服务访问。这种方式需要用户创建ServiceEntry规则，将Remote集群的服务暴露在本集群的服务网格内，并且由于集群之间的网络并不互通，所以这种模型依赖Remote集群的Gateway中转流量。</span></p></li></ul><p style="text-align: left;"><span style="font-size: 14px;font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;"><br  /></span></p><p style="text-align: left;"><span style="font-size: 14px;"><strong><span style="font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;">1.1 服务DNS解析的原理</span></strong></span><span style="font-size: 15px;"><strong><span style="font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;"></span></strong><strong><span style="font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;"></span></strong></span></p><p class="p2"><span style="font-size: 14px;">在多控制面模型中，每个集群都由独立的Istio控制面管理。我们可以在宏观上认为不同集群的工作负载属于不同的网格。但是在微观上，其实是所有集群联合起来构建了一个大的网格，因为在本地集群中可以以服务域名的形式访问Remote集群的服务。例如，Remote集群的forecast服务在命名空间weather中，那么本地集群在访问forecast服务时使用的是forecast. weather.global域名。这种带“.global”后缀的域名是Istio在多控制面模型中设计的虚假域名。在本地集群中，DNS将这种带“.global”后缀的服务域名解析成一个虚假的在服务网格内没有使用的IP地址。</span></p><p class="p2"><span style="font-size: 14px;"><mpcpc js_editor_cpcad="" class="js_cpc_area res_iframe cpc_iframe" src="/cgi-bin/readtemplate?t=tmpl/cpc_tmpl#1562211069554" data-category_id_list="1|16|26|36|42|43|47|48|5|6" data-id="1562211069554"></mpcpc></span></p><p class="p2"><span style="font-size: 14px;">多控制面DNS的解析原理为：多控制面模型要求Istio能够提供Remote集群的服务地址解析，并且不影响已有的服务。典型的应用期望使用服务DNS名称解析服务地址，并通过地址访问服务。Istio的Sidecar本身在服务之间路由请求时并不使用DNS名称，但是Istio离不开DNS解析，这是因为在一般情况下服务实例是以域名形式访问其他服务的，所以在服务实例发起请求时首先进行DNS解析，才能将请求发送出去。本地集群的服务共享相同的DNS后缀（svc.cluster.local），Kubernetes集群默认的DNS服务器提供对这类服务的域名解析。</span></p><p class="p2"><span style="font-size: 14px;"><br  /></span></p><p class="p6"><span style="font-size: 14px;">对于Remote集群的服务，本地集群的DNS显得束手无策，因为Kubernetes DNS只能通过Kube-apiserver获取本集群的服务及服务实例信息。Istio为支持Remote集群的服务DNS解析，做出了如下要求。</span></p><p class="p6"><span style="font-size: 14px;"><br  /></span></p><p class="p13"><span style="font-size: 14px;">1）在每个Kubernetes集群中都额外部署一个istiocoredns组件</span></p><p class="p14"><span style="font-size: 14px;">istiocoredns与在集群中默认安装的kube-dns和CoreDNS是级联的关系，协助默认的DNS服务器解析带“.global”后缀的服务，因此需要配置kube-dns和CoreDNS的私有DNS域服务器。如下所示为kube-dns服务器配置私有DNS域（stub domains）：</span></p><p class="p15"><br  /></p><section class="code-snippet__fix code-snippet__js"><ul class="code-snippet__line-index code-snippet__js"><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li></ul><pre class="code-snippet__js" data-lang="makefile"><code><span class="code-snippet_outer" style="">apiVersion: v1</span></code><code><span class="code-snippet_outer" style="">kind: ConfigMap</span></code><code><span class="code-snippet_outer" style="">metadata:</span></code><code><span class="code-snippet_outer">  name: kube-dns</span></code><code><span class="code-snippet_outer">  namespace: kube-system</span></code><code><span class="code-snippet_outer">  data:</span></code><code><span class="code-snippet_outer" style="">stubDomains: |</span></code><code><span class="code-snippet_outer">  {<span class="code-snippet__string">"global"</span>: [<span class="code-snippet__string">"$(kubectl get svc -n istio-system istiocoredns -o jsonpath={.spec.clusterIP})"</span>]}</span></code></pre></section><p class="p15"><span style="font-size: 14px;"></span><br  /></p><p class="p14"><span style="font-size: 14px;">CoreDNS的私有DNS域设置如下：</span></p><section class="code-snippet__fix code-snippet__js"><ul class="code-snippet__line-index code-snippet__js"><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li></ul><pre class="code-snippet__js" data-lang="properties"><code><span class="code-snippet_outer"><span class="code-snippet__meta">apiVersion</span>: <span class="code-snippet__string">v1</span></span></code><code><span class="code-snippet_outer"><span class="code-snippet__attr">kind</span>: <span class="code-snippet__string">ConfigMap</span></span></code><code><span class="code-snippet_outer"><span class="code-snippet__attr">metadata</span>:<span class="code-snippet__string"></span></span></code><code><span class="code-snippet_outer">  <span class="code-snippet__attr">name</span>: <span class="code-snippet__string">coredns</span></span></code><code><span class="code-snippet_outer">  <span class="code-snippet__meta">namespace</span>: <span class="code-snippet__string">kube-system</span></span></code><code><span class="code-snippet_outer"><span class="code-snippet__attr">data</span>:<span class="code-snippet__string"></span></span></code><code><span class="code-snippet_outer">  <span class="code-snippet__attr">Corefile</span>: <span class="code-snippet__string">|</span></span></code><code><span class="code-snippet_outer">    <span class="code-snippet__meta">.</span>:<span class="code-snippet__string">53 {</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__attr">errors</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__attr">health</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__meta">kubernetes</span> <span class="code-snippet__string">cluster.local in-addr.arpa ip6.arpa {</span></span></code><code><span class="code-snippet_outer">           <span class="code-snippet__meta">pods</span> <span class="code-snippet__string">insecure</span></span></code><code><span class="code-snippet_outer">           <span class="code-snippet__attr">upstream</span></span></code><code><span class="code-snippet_outer">           <span class="code-snippet__attr">fallthrough</span> <span class="code-snippet__string">in-addr.arpa ip6.arpa</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__attr">}</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__meta">prometheus</span> :<span class="code-snippet__string">9153</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__attr">proxy</span> <span class="code-snippet__string">. /etc/resolv.conf</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__attr">cache</span> <span class="code-snippet__string">30</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__attr">loop</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__attr">reload</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__attr">loadbalance</span></span></code><code><span class="code-snippet_outer">    <span class="code-snippet__attr">}</span></span></code><code><span class="code-snippet_outer">    <span class="code-snippet__attr">global</span>:<span class="code-snippet__string">53 {</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__attr">errors</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__attr">cache</span> <span class="code-snippet__string">30</span></span></code><code><span class="code-snippet_outer">        <span class="code-snippet__attr">proxy</span> <span class="code-snippet__string">. $(kubectl get svc -n istio-system istiocoredns -o jsonpath={.spec.clusterIP})</span></span></code><code><span class="code-snippet_outer">    <span class="code-snippet__attr">}</span></span></code></pre></section><p class="p15"><span style="font-size: 14px;"></span><br  /></p><p class="p17"><span style="font-size: 14px;">这两种DNS服务器的私有DNS域服务器都指向了istiocoredns服务，这样“&lt;name&gt;.&lt;namespace&gt;.global”类型的域名解析都会被转发到istiocoredns服务器上。至于istiocoredns服务器为什么可以解析“&lt;name&gt;.&lt;namespace&gt;.global”，请继续阅读下文。</span></p><p class="p17"><span style="font-size: 14px;"><br  /></span></p><p class="p18"><span style="font-size: 14px;">2）为Remote集群服务创建ServiceEntry</span></p><p class="p14"><span style="font-size: 14px;">我们可以通过创建ServiceEntry将Remote集群的服务暴露到本地集群内，除此之外，在多控制面模型中，ServiceEntry还可以用于带“.global”后缀的服务DNS进行解析。如下所示是remote集群的forecast服务在primary集群中创建的ServiceEntry：</span></p><p class="p15"><span style="font-size: 14px;"></span></p><section class="code-snippet__fix code-snippet__js"><ul class="code-snippet__line-index code-snippet__js"><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li><li></li></ul><pre class="code-snippet__js" data-lang="properties"><code><span class="code-snippet_outer"><span class="code-snippet__meta">apiVersion</span>: <span class="code-snippet__string">v1</span></span></code><code><span class="code-snippet_outer"><span class="code-snippet__meta">apiVersion</span>: <span class="code-snippet__string">networking.istio.io/v1alpha3</span></span></code><code><span class="code-snippet_outer"><span class="code-snippet__attr">kind</span>: <span class="code-snippet__string">ServiceEntry</span></span></code><code><span class="code-snippet_outer"><span class="code-snippet__attr">metadata</span>:<span class="code-snippet__string"></span></span></code><code><span class="code-snippet_outer">  <span class="code-snippet__attr">name</span>: <span class="code-snippet__string">forecast</span></span></code><code><span class="code-snippet_outer"><span class="code-snippet__meta">spec</span>:<span class="code-snippet__string"></span></span></code><code><span class="code-snippet_outer">  <span class="code-snippet__meta">hosts</span>:<span class="code-snippet__string"></span></span></code><code><span class="code-snippet_outer" style="">  # 必须是 name.namespace.global 的格式</span></code><code><span class="code-snippet_outer">  <span class="code-snippet__meta">-</span> <span class="code-snippet__string">forecast.weather.global</span></span></code><code><span class="code-snippet_outer" style="">  # 将Remote集群的服务认为是网格内部的服务，因为所有的集群共享相同的根证书</span></code><code><span class="code-snippet_outer">  <span class="code-snippet__attr">location</span>: <span class="code-snippet__string">MESH_INTERNAL</span></span></code><code><span class="code-snippet_outer">  <span class="code-snippet__meta">ports</span>:<span class="code-snippet__string"></span></span></code><code><span class="code-snippet_outer">  <span class="code-snippet__meta">-</span> <span class="code-snippet__string">name: http1</span></span></code><code><span class="code-snippet_outer">    <span class="code-snippet__attr">number</span>: <span class="code-snippet__string">8000</span></span></code><code><span class="code-snippet_outer">    <span class="code-snippet__attr">protocol</span>: <span class="code-snippet__string">http</span></span></code><code><span class="code-snippet_outer">  <span class="code-snippet__meta">resolution</span>: <span class="code-snippet__string">DNS</span></span></code><code><span class="code-snippet_outer">  <span class="code-snippet__meta">addresses</span>:<span class="code-snippet__string"></span></span></code><code><span class="code-snippet_outer" style="">  # 虚假的地址，所有发到这个地址的流量都会被Envoy拦截</span></code><code><span class="code-snippet_outer" style="">  # 并转发到${CLUSTER2_GW_ADDR}</span></code><code><span class="code-snippet_outer">  <span class="code-snippet__meta">-</span> <span class="code-snippet__string">127.255.0.2</span></span></code><code><span class="code-snippet_outer">  <span class="code-snippet__meta">endpoints</span>:<span class="code-snippet__string"></span></span></code><code><span class="code-snippet_outer" style="">  # 集群2的入口网关地址</span></code><code><span class="code-snippet_outer">  <span class="code-snippet__meta">-</span> <span class="code-snippet__string">address: ${CLUSTER2_GW_ADDR}</span></span></code><code><span class="code-snippet_outer">    <span class="code-snippet__meta">ports</span>:<span class="code-snippet__string"></span></span></code><code><span class="code-snippet_outer">      <span class="code-snippet__attr">http1</span>: <span class="code-snippet__string">15443 # Do not change this port value</span></span></code></pre></section><p class="p15"><span style="font-size: 14px;"></span><br  /></p><p class="p14"><span style="font-size: 14px;">其中，host名称必须是“&lt;name&gt;.&lt;namespace&gt;.global”形式，address字段必须是在网格中没有使用过的唯一IP地址，“&lt;name&gt;.&lt;namespace&gt;.global”服务的地址将被istiocoredns解析成该IP地址。必须将endpoints设置为Remote集群的Gateway的地址和端口，这里的端口必须是网关服务的TLS端口，默认是15443。</span></p><blockquote class="js_blockquote_wrap" data-type="2" data-url="" data-author-name="" data-content-utf8-length="77" data-source-title=""><section class="js_blockquote_digest"><section>注意：在多控制面的多集群模式下，网关地址必须是从对端集群可以访问的地址，在一般情况下，会通过负载均衡器为网关服务分配一个从集群外部可以访问的虚拟IP地址。</section></section></blockquote><p class="p14"><span style="font-size: 14px;">继续回到forecast.weather.global域名解析的主题上来，Kubernetes集群默认部署的DNS服务器根本没有该类型服务的IP信息，但是配置了存根域DNS服务器istiocoredns。istiocoredns是Istio社区为了支持全局的服务DNS解析而专门部署的。Istio社区专门开发了一个CoreDNS插件（https://github.com/istio-ecosystem/istio-coredns-plugin/blob/master/ plugin.go）用于解析ServiceEntry获取全局服务的IP地址，然后与CoreDNS部署在同一个Pod中，通过gRPC协议提供“*.global”类型的服务的DNS解析。</span></p><p class="p14"><span style="font-size: 14px;"><br  /></span></p><p class="p2"><span style="font-size: 14px;">“*.global”类型的服务域名解析流程如下：</span></p><ul class=" list-paddingleft-2" style="list-style-type: disc;"><li><p class="p14"><span style="font-size: 14px;">为支持Remote集群的forecast服务访问，需要创建ServiceEntry。</span></p></li><li><p class="p14"><span style="font-size: 14px;">在本地集群中访问“forecast.weather.global”服务时首先会进行DNS解析，Kubernetes集群内的DNS解析首先会被发送到kube-dns服务。</span></p></li><li><p class="p14"><span style="font-size: 14px;">默认的kube-dns/coredns根据存根DNS的设置，将“*.global”类型的域名解析请求转发到私有DNS服务器istiocoredns。</span></p></li><li><p class="p14"><span style="font-size: 14px;">istiocoredns通过Kube-apiserver获取ServiceEntry进而获取“forecast.weather. global”的IP地址。</span></p></li></ul><p class="p14"><span style="font-size: 14px;"><br  /></span></p><p class="p6"><span style="font-size: 14px;">*.global类型的服务域名解析流程如下图所示。</span></p><p style="text-align: center;"><img class="rich_pages" data-ratio="0.6077586206896551" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_png/z9BgVMEm7YsbvT2Aia1aQOf56e01MiaTXAaria8ZaYAZia4Jufx0xoxU65oKf9OxMHL2HhyF2asIibERUicjNLh7clsA/640?wx_fmt=png" data-type="png" data-w="928" style=""  /></p><p class="p21" style="text-align: center;"><span style="font-size: 12px;">*.global类型的服务域名解析流程<br  /></span></p><p class="p22"><span style="font-size: 14px;"><br  /></span></p><p class="p22"><strong><span style="font-size: 14px;">1.2 Gateway连接的原理</span></strong></p><p class="p2"><span style="font-size: 14px;">在多控制面模型中，Istio对于底层网络连通性没有过多的要求，跨集群的服务访问完全通过Gateway转发，因而只需每个集群的Gateway服务都对外暴露一个虚拟IP地址，供集群外部访问。从网络拓扑来看，多控制面模型部署比较轻量化。</span></p><p class="p2"><span style="font-size: 14px;"><mpcpc js_editor_cpcad="" class="js_cpc_area res_iframe cpc_iframe" src="/cgi-bin/readtemplate?t=tmpl/cpc_tmpl#1562211085305" data-category_id_list="1|16|26|36|42|43|47|48|5|6" data-id="1562211085305"></mpcpc></span></p><p class="p14"><span style="font-size: 14px;">最简单的多控制面跨集群访问模型如图7-3所示，其中，Cluster1中的frontend服务访问Cluster2中的forecast服务，集群1中的工作负载发起到forecast.weather.global的请求，请求首先被Sidecar拦截，Sidecar再根据配置规则将请求转发到Cluster2的Gateway，Gateway再将请求转发到forecast服务实例，基本的流量转发依赖Istio的ServiceEntry、Gateway、VirtualService配置规则。</span></p><p class="p14"><span style="font-size: 14px;"><br  /></span></p><p class="p6"><span style="font-size: 14px;">随着多集群的复杂度提高，例如在级联多集群场景下，相同的应用及服务跨集群部署频见，这对Istio配置规则的设置要求非常高，并且很难自动控制不同集群中服务实例的负载均衡权重。虽然这种多控制面模型基本能够满足跨集群的服务访问，但是需要额外设置很多VirtualService、DestinationRule、Gateway、ServiceEntry等API对象，比较复杂。</span></p><p class="p7"><span style="font-size: 14px;">&nbsp;</span></p><p style="text-align: center;"><img class="rich_pages" data-ratio="0.6262230919765166" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_png/z9BgVMEm7YsbvT2Aia1aQOf56e01MiaTXAGfMdeMV6NsNeIoFVpHPs5J9Aibia6M0ppBhZqTUunXOZJUdeQMwJebFA/640?wx_fmt=png" data-type="png" data-w="1022" style=""  /></p><p class="p8" style="text-align: center;"><span style="font-size: 12px;">最简单的多控制面跨集群访问模型</span><span style="font-size: 14px;"><br  /></span></p><p><br  /></p><p><br  /></p><p class="p1"><span style="font-size: 14px;">本篇内容节选及改编自《云原生服务网格Istio：</span><span style="font-size: 14px;">原理、实战、架构与源码解析》</span></p><p class="p1"><span style="font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;font-size: 14px;"><span style="font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;">本书京东链接：</span></span><span style="font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;font-size: 14px;">https://item.jd.com/12538407.html</span><br  /></p><p class="p1"><span style="font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;font-size: 14px;"><br  /></span></p><p class="p1"><span style="font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;font-size: 14px;"></span></p><section class="p1"><mpcps frameborder="0" class="js_editor_cps" data-datakey="1562210423788_0.7982237946688111" data-uid="1562210423785" style="width:100% !important;border:0;" data-type="1" data-product="" data-templateid="list" data-pid="27889100" data-packid="" data-smartnum="" data-color="#fa7834" data-categoryid="3" data-appid="wx831660fe3ded4389" data-report="s0%3D0%26s1%3D0%26s2%3D0%26s3%3Distio%26s4%3D0%26s5%3D10%26s6%3Did_1562211155096_140298%26s7%3D%26s8%3D%26s9%3D%26s10%3D%26pid%3Dwx831660fe3ded4389_27889100%26uuid%3D3582441586728973458%26title%3D%25E4%25BA%2591%25E5%258E%259F%25E7%2594%259F%25E6%259C%258D%25E5%258A%25A1%25E7%25BD%2591%25E6%25A0%25BCIstio%25EF%25BC%259A%25E5%258E%259F%25E7%2590%2586%25E3%2580%2581%25E5%25AE%259E%25E8%25B7%25B5%25E3%2580%2581%25E6%259E%25B6%25E6%259E%2584%25E4%25B8%258E%25E6%25BA%2590%25E7%25A0%2581%25E8%25A7%25A3%25E6%259E%2590%26sid%3D3%26cid%3D3%26ratio%3D17.00%2525%26price%3D109.80%26"></mpcps></section><p><br  /></p><hr style="border-style: solid;border-width: 1px 0 0;border-color: rgba(0,0,0,0.1);-webkit-transform-origin: 0 0;-webkit-transform: scale(1, 0.5);transform-origin: 0 0;transform: scale(1, 0.5);"  /><p><br  /></p><h2 class="" style="margin-bottom: 14px;font-size: 22px;max-width: 100%;font-family: -apple-system-font, system-ui, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;letter-spacing: 0.544px;white-space: normal;line-height: 1.4;text-align: start;background-color: rgb(255, 255, 255);box-sizing: border-box !important;overflow-wrap: break-word !important;"><span style="max-width: 100%;font-size: 14px;box-sizing: border-box !important;overflow-wrap: break-word !important;"><a href="http://mp.weixin.qq.com/s?__biz=MzU4MjQ0MTU4Ng==&amp;mid=2247484223&amp;idx=1&amp;sn=879d99b07136047e534def13656618e3&amp;chksm=fdb90c22cace853489a891c05d8c89779785bd49ddf083734a0d9e5b2a1768081d07491a59b5&amp;scene=21#wechat_redirect" target="_blank" data-itemshowtype="0" data-linktype="2" hasload="1" style="-webkit-tap-highlight-color: rgba(0, 0, 0, 0);max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;">k8s进阶课程推荐：打造独当一面的 Kubernetes 运维、开发工程师</a></span></h2><p><span style="font-size: 14px;">扫描下面的二维码(或微信搜索 k8s技术圈)关注我们的微信公众帐号，在微信公众帐号中回复 <strong>加群</strong> 即可加入到我们的 kubernetes 讨论群里面共同学习。</span></p><p style="max-width: 100%;min-height: 1em;font-family: -apple-system-font, system-ui, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif;letter-spacing: 0.544px;white-space: normal;background-color: rgb(255, 255, 255);text-align: center;box-sizing: border-box !important;overflow-wrap: break-word !important;"><img class="rich_pages " data-copyright="0" data-ratio="0.5555555555555556" data-s="300,640" data-type="png" data-w="900" data-src="https://mmbiz.qpic.cn/mmbiz_png/z9BgVMEm7YsiacqkyJTdQiaAGtia83tkk4UOQdU1ibEgyN7mHPyNaITUOYs4TobFuSeJe45qz0yO4dA4ZTq9njnxeQ/640?wx_fmt=png" style="box-sizing: border-box !important;overflow-wrap: break-word !important;visibility: visible !important;width: 677px !important;"  /></p><p><br  /></p>
                </div>
                <script nonce="1977935583" type="text/javascript">
                    var first_sceen__time = (+new Date());

                    if ("" == 1 && document.getElementById('js_content')) {
                        document.getElementById('js_content').addEventListener("selectstart",function(e){ e.preventDefault(); });
                    }

                    
                    (function(){
                        if (navigator.userAgent.indexOf("WindowsWechat") != -1){
                            var link = document.createElement('link');
                            var head = document.getElementsByTagName('head')[0];
                            link.rel = 'stylesheet';
                            link.type = 'text/css';
                            link.href = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/style/page/appmsg_new/winwx46b604.css";
                            head.appendChild(link);
                        }
                    })();
                </script>

                
  <div class="ct_mpda_wrp" id="js_sponsor_ad_area" style="display: none;"></div>


                
                <div class="read-more__area" id="js_more_read_area" style="display:none;">
                    
                </div>

                
                                <div class="reward_area tc reward_area_primary" id="js_preview_reward_author" style="display:none;">
                    <div class="reward-avatar" style="display: none;" id="js_preview_reward_author_avatar">
                        <img src="" alt="" id="js_preview_reward_author_head">
                    </div>
                    
                                        <div class="reward-author" id="js_preview_reward_author_name"></div>
                                        <p class="reward_tips" id="js_preview_reward_author_wording" style="display:none;"></p>
                    <p class="reward_button_wrp">
                    
                      <span class="reward_pop_panel">
                        <img src="https://res.wx.qq.com/mpres/zh_CN/htmledition/pages/home/index/pic_mp_app4290ba.png" alt="">
                        <strong>扫一扫下载订阅号助手，用手机发文章</strong>
                      </span>
                        <a class="reward_button" id='js_preview_reward_author_link' href="##"><span id="js_preview_reward_link_text">赞赏</span></a>
                    </p>
                </div>

                <div class="reward_qrcode_area reward_area tc" id="js_preview_reward_qrcode" style="display:none;">
                    <p class="tips_global">长按二维码向我转账</p>
                    <p id="js_preview_reward_ios_wording" class="reward_tips" style="display:none;"></p>
                    <span class="reward_qrcode_img_wrp"><img class="reward_qrcode_img" src="//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/pic/appmsg/pic_reward_qrcode.2x42f400.png"></span>
                    <p class="tips_global">受苹果公司新规定影响，微信 iOS 版的赞赏功能被关闭，可通过二维码转账支持公众号。</p>
                </div>
                                
                
                            </div>
                                        
                                

                        
            <ul id="js_hotspot_area" class="article_extend_area"></ul>


            

            


<div class="rich_media_tool" id="js_toobar3">

            <div id="js_read_area3" class="media_tool_meta tips_global_primary meta_primary" style="display:none;">

    <span id="readTxt">阅读</span>

    <span id="readNum3"></span>
    </div>


    <span style="display:none;" class="media_tool_meta meta_extra meta_praise" id="like_old">
        <i class="icon_praise_gray"></i><span class="praise_num" id="likeNum_old"></span>
    </span>

      
    <span class="media_tool_meta meta_extra meta_share" style="display: none;">
      <button class="share_btn" id="js_share_btn">分享</button>
    </span>
    <span style="display:none;" class="media_tool_meta meta_extra meta_like" id="like3">
        <button class="like_btn" id="js_like_btn"> 
          <span id="js_like_wording"> 在看</span><span class="like_num" id="likeNum3"></span>
        </button>
    </span>

      <div style="display:none">
        <div class="weui-mask"></div>
        <div class="weui-dialog weui-dialog_haokan">
        <div class="weui-dialog__hd"><strong class="weui-dialog__title" id="educate_title">已同步到看一看</strong></div>
        <div class="weui-dialog__bd">
        </div>
        <div class="weui-dialog__ft" id="educate_btn" style="display: none">
          <a href="javascript:;" class="weui-dialog__btn weui-dialog__btn_default" id="js_cancel">取消</a>
          <a href="javascript:;" class="weui-dialog__btn weui-dialog__btn_primary" id="js_confirm">发送</a>
        </div>
        <div class="weui-dialog__ft" id="educate_btn2" style="display: none">
          <a href="javascript:;" class="weui-dialog__btn weui-dialog__btn_primary" id="js_acknowledge">我知道了</a>
        </div>
        </div>
    </div>

    
    </div>

<div style="display: none;" id="js_like_educate" class="like_skin_primary like_comment_primary_pop">
  <div id="js_like_educate_wrapper" class="like_comment_primary_wrp like_comment_primary_pos_top">
    <div class="like_comment_primary_inner">
      <div class="like_comment_primary_hd">
        <h4 class="like_comment_primary_title">
          朋友会在“发现-看一看”看到你“在看”的内容        </h4>
        <button id="js_b_like_confirm" class="like_comment_primary_btn">确定</button>
      </div>
      <div class="like_comment_primary_bd">
        <img class="pic_haokan" src="//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/pic/appmsg/pic_haokan4516f7.png">
      </div>
    </div>
  </div>
  <div class="like_comment_primary_mask" id="js_mask_3"></div>
</div>

  
  <div class="like_comment_wrp" id="js_a_like_comment" style="display: none;">
    <div class="like_comment_inner">
      <div class="like_comment_hd" style="display:none" id="js_like_title"></div>
      <div class="like_comment_bd">
        <div class="tips_global_primary like_comment_tips">
          <i class="weui-icon-success"></i><i class="icon-success-primary"></i>已同步到看一看<a href="javascript:;" class="like_comment_share_link" id="js_a_like_comment_share">写下你的想法</a>
        </div>
      </div>
      <div class="like_comment_ft">
        <span id="js_a_like_comment_msg" class="like_comment_msg" style="visibility:hidden">最多200字，当前共<span id="js_a_like_current_cnt"></span>字</span>
        <button class="like_comment_btn" disabled="disabled" id="js_a_like_confirm">发送</button>
      </div>
    </div>
  </div>

<div id="js_like_toast" style="display: none;">
  <div class="weui-mask_transparent"></div>
  <div class="weui-toast">
    <i class="weui-icon-success-no-circle weui-icon_toast"></i>
    <p class="weui-toast__content" id="js_toast_msg">已发送</p>
  </div>
</div>


<div style="display: none;" id="js_b_comment_panel" class="like_comment_primary_pop">
  <div class="like_comment_primary_wrp" id="js_b_wrp">
    <div class="like_comment_primary_inner">
      <div class="like_comment_primary_hd">
        <h4 class="like_comment_primary_title">
          朋友将在看一看看到        </h4>
        <button id="js_b_like_confirm" class="like_comment_primary_btn">确定</button>
      </div>
      <div id='js_b_comment_text_first' class="like_comment_primary_bd">
        <span class="tips_global_primary">
          写下你的想法...        </span>
      </div>
    </div>
  </div>
  <div class="like_comment_primary_mask" id="js_mask_1"></div>
</div>

<div style="display: none;" id="js_b_comment_final">
  <div class="like_comment_primary_wrp editing" id="js_comment_wrp">
    <div class="like_comment_primary_inner">
      <div class="like_comment_primary_hd">
        <button class="like_comment_primary_cancel" id="js_b_comment_cancel">取消</button>
        <h4 class="like_comment_primary_title"> 发布到看一看 </h4>
        <button class="like_comment_primary_btn" id="js_b_comment_confirm">确定</button>
      </div>
      <div class="like_comment_primary_bd">
        <textarea class="like_comment_textarea weui-textarea" placeholder="写下你的想法..." id="js_b_comment_text_second"></textarea>
      </div>
      <span class="like_comment_msg" id="js_b_like_comment_msg" style="visibility: hidden;">最多200字，当前共<span id="js_b_like_current_cnt"></span>字</span>
    </div>
  </div>
  <div class="like_comment_primary_mask" id="js_mask_2"></div>
</div>

<div id="js_loading" style=" display: none;">
    <div class="weui-mask_transparent"></div>
    <div class="weui-toast">
        <i class="weui-loading weui-icon_toast"></i>
        <p class="weui-toast__content">发送中</p>
    </div>
</div>



                      </div>
        </div>

        <div class="rich_media_area_primary sougou" id="sg_tj" style="display:none"></div>


        
        <div class="rich_media_area_extra">
          <div class="rich_media_area_extra_inner">
              
              <div id="js_share_appmsg">
              </div>

              
  
      <div class="mpda_bottom_container" id="js_bottom_ad_area"></div>
                
              <div id="js_iframetest" style="display:none;"></div>
                            
              <div class="rich_media_extra rich_media_extra_discuss" id="js_cmt_container" style="display:none">
                

                
                <div class="discuss_mod" id="js_friend_cmt_area" style="display:none">
                  
                  
                  
                </div>

                                <div class="discuss_mod" id="js_cmt_area" style="display:none">
                </div>
                              </div>
          </div>
        </div>

        
        <div id="js_pc_qr_code" class="qr_code_pc_outer" style="display:none;">
            <div class="qr_code_pc_inner">
                <div class="qr_code_pc">
                    <img id="js_pc_qr_code_img" class="qr_code_pc_img">
                    <p>微信扫一扫<br>关注该公众号</p>
                </div>
            </div>
        </div>
    </div>
</div>



<div id="js_pc_weapp_code" class="weui-desktop-popover weui-desktop-popover_pos-up-center weui-desktop-popover_img-text" style="display: none;">
    <div class="weui-desktop-popover__content">
        <div class="weui-desktop-popover__desc">
            <img id="js_pc_weapp_code_img">
            微信扫一扫<br>使用小程序<span id="js_pc_weapp_code_des"></span>        </div>
    </div>
</div>
<div id="js_minipro_dialog" style="display:none;">
    <div class="weui-mask"></div>
    <div class="weui-dialog weui-dialog_link">
        <div class="weui-dialog__bd" id="js_minipro_dialog_name"></div>
        <div class="weui-dialog__ft">
            <a id="js_minipro_dialog_cancel" href="javascript:void(0);" class="weui-dialog__btn weui-dialog__btn_default">取消</a>
            <a id="js_minipro_dialog_ok" href="javascript:void(0);" class="weui-dialog__btn weui-dialog__btn_primary">允许</a>
        </div>
    </div>
</div>
<div id="js_link_dialog" style="display:none;">
    <div class="weui-mask"></div>
    <div class="weui-dialog weui-dialog_link">
        <div class="weui-dialog__bd" id="js_link_dialog_name">即将打开一个新页面</div>
        <div class="weui-dialog__ft">
            <a id="js_link_dialog_cancel" href="javascript:void(0);" class="weui-dialog__btn weui-dialog__btn_default">取消</a>
            <a id="js_link_dialog_ok" href="javascript:void(0);" class="weui-dialog__btn weui-dialog__btn_primary">允许</a>
        </div>
    </div>
</div>

        
        <script nonce="1977935583" type="text/javascript">
  (function () {
    
    var ajaxWhiteList = {
      1: { 
        reg: /^https?:\/\/mp\.weixin\.qq\.com\/mp\/appmsg_like/,
        times: 0
      },
      2: { 
        reg: /^https?:\/\/mp\.weixin\.qq\.com\/mp\/appmsg_comment((\?|&)[^=]*?=[^&]*?)*?(\?|&)action=likecomment/,
        times: 0
      },
      3: { 
        reg: /^https?:\/\/mp\.weixin\.qq\.com\/mp\/appmsg_comment((\?|&)[^=]*?=[^&]*?)*?(\?|&)action=addcomment/,
        times: 0
      },
      4: { 
        reg: /^https?:\/\/mp\.weixin\.qq\.com\/mp\/authorreward/,
        times: 0
      }
      
      
      
      
    };

    
    if (!performance || !performance.getEntries) return;

    
    var hasReported = false;

    
    var reportResLoadTime = function () {
      
      if (hasReported) return;

      
      var notSupport = false;

      
      var ajaxEntries = [];

      
      var entries = performance.getEntries().map(function (entry) {
        if (typeof entry !== 'object') {
          notSupport = true;
        } else if (entry.entryType === undefined) {
          notSupport = true;
        } else if (
          entry.entryType !== 'navigation' &&
          entry.entryType !== 'resource'
        ) {
          
          return null;
        } else if (entry.initiatorType === undefined) {
          notSupport = true;
        } else if (entry.initiatorType === 'xmlhttprequest') {
          
          if (entry.name === undefined || entry.duration === undefined) {
            notSupport = true;
          } else {
            for (var scene in ajaxWhiteList) {
              if (Object.prototype.hasOwnProperty.call(ajaxWhiteList, scene)) {
                var ajaxItem = ajaxWhiteList[scene];
                if (ajaxItem.times < 10 && ajaxItem.reg.test(entry.name)) {
                  ajaxEntries.push({
                    scene: scene,
                    protocol: entry.nextHopProtocol,
                    is_https: isHttps(entry),
                    time: entry.duration
                  });
                  ajaxItem.times++;
                }
              }
            }
          }
          return null;
        }
        return entry;
      });

      
      if (notSupport || ajaxEntries.length === 0) return;

      
      var data = {
        stat_list: ajaxEntries
      };
      var img = new Image();
      img.src = 'https://mp.weixin.qq.com/mp/timereport?data=' + JSON.stringify(data);
      hasReported = true;
    };

    
    window.addEventListener('beforeunload', reportResLoadTime, false);
    window.addEventListener('unload', reportResLoadTime, false);

    function isHttps(entry) {
      if (/^https/.test(entry.name)) return 1;
      else return 0;
    }
  })();
</script>
        <script nonce="1977935583">
    var __DEBUGINFO = {
        debug_js : "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/debug/console42f400.js",
        safe_js : "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/safe/moonsafe42f400.js",
        res_list: []
    };
</script>

<script nonce="1977935583" type="text/javascript">
(function() {
	var totalCount = 0,
			finishCount = 0;

	function _loadVConsolePlugin() {
		window.vConsole = new window.VConsole();
		while (window.vConsolePlugins.length > 0) {
			var p = window.vConsolePlugins.shift();
			window.vConsole.addPlugin(p);
		}
	}
	
	function _addVConsole(uri, cb) {
		totalCount++;
		var node = document.createElement('SCRIPT');
		node.type = 'text/javascript';
		node.src = uri;
		node.setAttribute('nonce', '1977935583');
		if (cb) {
			node.onload = cb;
		}
		document.getElementsByTagName('head')[0].appendChild(node);
	}
	if (
		(document.cookie && document.cookie.indexOf('vconsole_open=1') > -1)
		|| location.href.indexOf('vconsole=1') > -1
	) {
		window.vConsolePlugins = [];
		_addVConsole('//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/vconsole/3.2.2/vconsole.min440203.js', function() {
			
			_addVConsole('//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/vconsole/plugin/vconsole-mpopt/1.0.1/vconsole-mpopt42f400.js', _loadVConsolePlugin);
		});
	}

  
  
  
  
  
  
  
  

  

})();
</script>
        
<script nonce="1977935583" type="text/javascript">
if (location.href.match(/fontScale=\d+/) && /(iPhone|iPad|iPod|iOS)/i.test(navigator.userAgent)) {

    var m=location.href.match(/fontScale=(\d*)/);
    var map={
        "94":1,
        "100":2,
        "109":3,
        "119":4,
        "131":5,
    }
    if(m&&m[1]&&map[m[1]]){
        document.getElementsByTagName("body")[0].className+=" appmsg_skin_fontscale_"+map[m[1]];
        console.log("appmsg_skin_fontscale_:"+map[m[1]]);
    }
}
</script>
<script nonce="1977935583" type="text/javascript">
function _typeof(e){
return e&&"undefined"!=typeof Symbol&&e.constructor===Symbol?"symbol":typeof e;
}
!function(e){
if("object"===("undefined"==typeof module?"undefined":_typeof(module)))module.exports=e;else{
if(window.__second_open__)return;
var t="1562229368",n="1562212007",s="2019-07-04";
e(t,n,s,document.getElementById("publish_time"));
}
}(function(e,t,n,s){
var i="",o=86400,f=new Date(1e3*e),r=1*t,l=n||"";
f.setHours(0),f.setMinutes(0),f.setSeconds(0);
var c=f.getTime()/1e3;
f.setDate(1),f.setMonth(0);
var u=f.getTime()/1e3;
if(r>=c)i="今天";else if(r>=c-o)i="昨天";else if(r>=c-2*o)i="前天";else if(r>=c-3*o)i="3天前";else if(r>=c-4*o)i="4天前";else if(r>=c-5*o)i="5天前";else if(r>=c-6*o)i="6天前";else if(r>=c-14*o)i="1周前";else if(r>=u){
var d=l.split("-");
i="%s月%s日".replace("%s",parseInt(d[1],10)).replace("%s",parseInt(d[2],10));
}else i=l;
s&&(s.innerText=i,setTimeout(function(){
s.onclick=function(){
s.innerText=l;
};
},10));
});
</script>
<script nonce="1977935583" type="text/javascript">

if (!window.console) window.console = { log: function() {} };

if (typeof getComputedStyle == 'undefined') {
    if (document.body.currentStyle) {
        window.getComputedStyle = function(el) {
            return el.currentStyle;
        }
    } else {
        window.getComputedStyle = {};
    }
}
(function(){
    window.__zoom = 1;
    
    (function(){
        var validArr = ","+([0.875, 1, 1.125, 1.25, 1.375]).join(",")+",";
        var match = window.location.href.match(/winzoom=(\d+(?:\.\d+)?)/);
        if (match && match[1]) {
            var winzoom = parseFloat(match[1]);
            if (validArr.indexOf(","+winzoom+",")>=0) {
                window.__zoom = winzoom;
            }
        }
    })();

    var ua = navigator.userAgent.toLowerCase();
    var re = new RegExp("msie ([0-9]+[\.0-9]*)");
    var version;
    if (re.exec(ua) != null) {
        version = parseInt(RegExp.$1);
    }
    var isIE = false;
    if (typeof version != 'undefined' && version >= 6 && version <= 9) {
        isIE = true;
    }
    var getMaxWith=function(){
        var container = document.getElementById('img-content');
        var max_width = container.offsetWidth;
        var container_padding = 0;
        var container_style = getComputedStyle(container);
        container_padding = parseFloat(container_style.paddingLeft) + parseFloat(container_style.paddingRight);
        max_width -= container_padding;
        if (!max_width) {
            max_width = window.innerWidth - 30;      
        }
        return max_width;
    };
    var getParentWidth = function(dom){
        var parent_width = 0;
        var parent = dom.parentNode;
        var outerWidth = 0;
        while (true) {
            if(!parent||parent.nodeType!=1) break;
            var parent_style = getComputedStyle(parent);
            if (!parent_style) break;
            parent_width = parent.clientWidth - parseFloat(parent_style.paddingLeft) - parseFloat(parent_style.paddingRight) - outerWidth;
            if (parent_width > 0) break;
            outerWidth += parseFloat(parent_style.paddingLeft) + parseFloat(parent_style.paddingRight) + parseFloat(parent_style.marginLeft) + parseFloat(parent_style.marginRight) + parseFloat(parent_style.borderLeftWidth) + parseFloat(parent_style.borderRightWidth);
            parent = parent.parentNode;
        }
        return parent_width;
    }
    var getOuterW=function(dom){
        var style=getComputedStyle(dom),
            w=0;
        if(!!style){
            w = parseFloat(style.paddingLeft) + parseFloat(style.paddingRight) + parseFloat(style.borderLeftWidth) + parseFloat(style.borderRightWidth);
        }
        return w;
    };
    var getOuterH =function(dom){
        var style=getComputedStyle(dom),
            h=0;
        if(!!style){
            h = parseFloat(style.paddingTop) + parseFloat(style.paddingBottom) + parseFloat(style.borderTopWidth) + parseFloat(style.borderBottomWidth);
        }
        return h;
    };
    var insertAfter = function(dom,afterDom){
        var _p = afterDom.parentNode;
        if(!_p){
            return;
        }
        if(_p.lastChild === afterDom){
            _p.appendChild(dom);
        }else{
            _p.insertBefore(dom,afterDom.nextSibling);
        }
    };
    var getQuery = function(name,url){
        
        var u  = arguments[1] || window.location.search,
            reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)"),
            r = u.substr(u.indexOf("\?")+1).match(reg);
        return r!=null?r[2]:"";
    };

    
    function setImgSize(item, widthNum, widthUnit, ratio, breakParentWidth) {
        setTimeout(function () {
            var img_padding_border = getOuterW(item) || 0;
            var img_padding_border_top_bottom = getOuterH(item) || 0;
            
            if (widthNum > getParentWidth(item) && !breakParentWidth) {
                widthNum = getParentWidth(item);
            }

            height = (widthNum - img_padding_border) * ratio + img_padding_border_top_bottom;

            if (isIE) {
                var url = item.getAttribute('data-src');
                item.src = url;
            } else {
                if(parseFloat(widthNum, 10) > 40 && height > 40 && breakParentWidth) {
                    item.className += ' img_loading';
                }
                item.src = "data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg==";
            }
            item.style.cssText += ";width: " + widthNum + widthUnit + " !important;";
            item.style.cssText += ";height: " + height + widthUnit + " !important;";
        }, 10);
    }

    (function(){
        var images = document.getElementsByTagName('img');
        var length = images.length;
        var max_width = getMaxWith();
        for (var i = 0; i < length; ++i) {
            if (window.__second_open__ && images[i].getAttribute('__sec_open_place_holder__')) {
                continue;
            }
            var imageItem = images[i];
            var src_ = imageItem.getAttribute('data-src');
            var realSrc = imageItem.getAttribute('src');
            if (!src_ || realSrc) continue;
            
            var originWidth = imageItem.getAttribute('data-w');
            var ratio_ = 1 * imageItem.getAttribute('data-ratio');

            var height = 100;
            if (ratio_ && ratio_ > 0) {
                var parent_width = getParentWidth(imageItem) || max_width;
                var initWidth = imageItem.style.width || imageItem.getAttribute('width') || originWidth || parent_width;
                initWidth = parseFloat(initWidth, 10) > max_width ? max_width : initWidth;
                
                if (initWidth) {
                    imageItem.setAttribute('_width', !isNaN(initWidth * 1) ? initWidth + 'px' : initWidth);
                }
                
                if (typeof initWidth === 'string' && initWidth.indexOf('%') !== -1) {
                    initWidth = parseFloat(initWidth.replace('%', ''), 10) / 100 * parent_width;
                }
                
                if (initWidth === 'auto') {
                    initWidth = originWidth;
                }

                var res = /^(\d+(?:\.\d+)?)([a-zA-Z%]+)?$/.exec(initWidth);
                var widthNum = res && res.length >= 2 ? res[1] : 0;
                var widthUnit = res && res.length >= 3 && res[2] ? res[2] : 'px';

                
                setImgSize(imageItem, widthNum, widthUnit, ratio_, true);
                
                (function (item, widthNumber, unit, ratio) {
                    setTimeout(function () {
                        setImgSize(item, widthNumber, unit, ratio, false);
                    });
                })(imageItem, widthNum, widthUnit, ratio_);
            } else {
                imageItem.style.cssText += ";visibility: hidden !important;";
            }

        }
    })();
    window.__videoDefaultRatio=16/9;
    window.__getVideoWh = function(dom){
        var max_width = getMaxWith(),
            width = max_width,
            ratio_ = dom.getAttribute('data-ratio')*1,
            arr = [4/3, 16/9],
            ret = arr[0],
            abs = Math.abs(ret - ratio_);
        if (!ratio_) { 
            if (dom.getAttribute("data-mpvid")) { 
                ratio_ = 16/9;
            } else { 
                ratio_ = 4/3;
            }
        } else { 
            for (var j = 1, jl = arr.length; j < jl; j++) {
                var _abs = Math.abs(arr[j] - ratio_);
                if (_abs < abs) {
                    abs = _abs;
                    ret = arr[j];
                }
            }
            ratio_ = ret;
        }

        var parent_width = getParentWidth(dom)||max_width,
            width = width > parent_width ? parent_width : width,
            outerW = getOuterW(dom)||0,
            outerH = getOuterH(dom)||0,
            videoW = width - outerW,
            videoH = videoW/ratio_,
            height = videoH + outerH;
        return {w:Math.ceil(width),h:Math.ceil(height),vh:videoH,vw:videoW,ratio:ratio_};
    };


    (function(){
        var iframe = document.getElementsByTagName('iframe');
        for (var i=0,il=iframe.length;i<il;i++) {
            if (window.__second_open__ && iframe[i].getAttribute('__sec_open_place_holder__')) {
                continue;
            }
            var a = iframe[i];
            var src_ = a.getAttribute('src')||a.getAttribute('data-src')||"";
            if(!/^http(s)*\:\/\/v\.qq\.com\/iframe\/(preview|player)\.html\?/.test(src_)
                && !/^http(s)*\:\/\/mp\.weixin\.qq\.com\/mp\/readtemplate\?t=pages\/video_player_tmpl/.test(src_)
            ){
                continue;
            }
            var vid = getQuery("vid",src_);
            if(!vid){
                continue;
            }
            vid=vid.replace(/^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g,"");
            a.removeAttribute('src');
            a.style.display = "none";
            var obj = window.__getVideoWh(a),
                videoPlaceHolderSpan = document.createElement('span'),
                videoPlayerIconSpan = document.createElement('span'),
                mydiv = document.createElement('img');

            videoPlaceHolderSpan.className = "js_img_loading db";
            videoPlaceHolderSpan.setAttribute("data-vid", vid);
            

            videoPlayerIconSpan.className = 'wx_video_context db'; 
            videoPlayerIconSpan.style.display = 'none';
            videoPlayerIconSpan.innerHTML = '<span class="wx_video_thumb_primary"></span><button class="wx_video_play_btn">播放</button><span class="wx_video_mask"></span>'; 

            mydiv.className = "img_loading";

            mydiv.src="data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg==";
            
            
            videoPlaceHolderSpan.style.cssText = "width: " + obj.w + "px !important;";
            mydiv.style.cssText += ";width: " + obj.w + "px";
            videoPlaceHolderSpan.appendChild(videoPlayerIconSpan);
            videoPlaceHolderSpan.appendChild(mydiv);
            insertAfter(videoPlaceHolderSpan, a); 

            a.style.cssText += ";width: " + obj.w + "px !important;";
            a.setAttribute("width",obj.w);
            if(window.__zoom!=1){
                a.style.display = "block";
                videoPlaceHolderSpan.style.display = "none";
                a.setAttribute("_ratio",obj.ratio);
                a.setAttribute("_vid",vid);
            }else{
                videoPlaceHolderSpan.style.cssText += "height: " + obj.h + "px !important;";
                mydiv.style.cssText += "height: " + obj.h + "px !important;";
                a.style.cssText += "height: " + obj.h + "px !important;";
                a.setAttribute("height",obj.h);
            }
            a.setAttribute("data-vh",obj.vh);
            a.setAttribute("data-vw",obj.vw);
            if(a.getAttribute("data-mpvid")){
                a.setAttribute("data-src",location.protocol+"//mp.weixin.qq.com/mp/readtemplate?t=pages/video_player_tmpl&auto=0&vid="+vid);
            }else{
                a.setAttribute("data-src",location.protocol+"//v.qq.com/iframe/player.html?vid="+ vid + "&width="+obj.vw+"&height="+obj.vh+"&auto=0");
            }
        }
    })();

    (function(){
        if(window.__zoom!=1){
            if (!window.__second_open__) {
                document.getElementById('page-content').style.zoom = window.__zoom;
                var a = document.getElementById('activity-name');
                var b = document.getElementById('meta_content');
                if(!!a){
                    a.style.zoom = 1/window.__zoom;
                }
                if(!!b){
                    b.style.zoom = 1/window.__zoom;
                }
            }
            var images = document.getElementsByTagName('img');
            for (var i = 0,il=images.length;i<il;i++) {
                if (window.__second_open__ && images[i].getAttribute('__sec_open_place_holder__')) {
                    continue;
                }
                images[i].style.zoom = 1/window.__zoom;
            }
            var iframe = document.getElementsByTagName('iframe');
            for (var i = 0,il=iframe.length;i<il;i++) {
                if (window.__second_open__ && iframe[i].getAttribute('__sec_open_place_holder__')) {
                    continue;
                }
                var a = iframe[i];
                a.style.zoom = 1/window.__zoom;
                var src_ = a.getAttribute('data-src')||"";
                if(!/http(s)*\:\/\/v\.qq\.com\/iframe\/(preview|player)\.html\?/.test(src_)){
                    continue;
                }
                var ratio = a.getAttribute("_ratio");
                var vid = a.getAttribute("_vid");
                a.removeAttribute("_ratio");
                a.removeAttribute("_vid");
                var vw = a.offsetWidth - (getOuterW(a)||0);
                var vh = vw/ratio;
                var h = vh + (getOuterH(a)||0)
                a.style.cssText += "height: " + h + "px !important;"
                a.setAttribute("height",h);
                a.setAttribute("data-src",location.protocol+"//v.qq.com/iframe/player.html?vid="+ vid + "&width="+vw+"&height="+vh+"&auto=0");
                a.style.display = "none";
                var parent = a.parentNode;
                if(!parent){
                    continue;
                }
                for(var j=0,jl=parent.children.length;j<jl;j++){
                    var child = parent.children[j];
                    if(child.className.indexOf("img_loading")>=0 && child.getAttribute("data-vid")==vid){
                        child.style.cssText += "height: " + h + "px !important;";
                        child.style.display = "";
                    }
                }
            }
        }
    })();
})();
</script>
<script nonce="1977935583" type="text/javascript">
    var new_appmsg = 1;
    var item_show_type = "0";
    var can_see_complaint = "0";
    var not_in_mm_css = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/style/page/appmsg_new/not_in_mm46b604.css";
    var windowwx_css = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/style/page/appmsg_new/winwx46b604.css";
    var article_improve_combo_css = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/style/page/appmsg_new/combo46e3e3.css";
    var tid = "";
    var aid = "";
    var clientversion = "";
    var appuin = ""||"MzU4MjQ0MTU4Ng==";

    var source = "";
    var ascene = "";
    var subscene = "";
    var sessionid = ""||"svr_4b2fbed5bf3";
    var abtest_cookie = "";

    var scene = 75;

    var itemidx = "";
    var appmsg_token   = "";
    var _copyright_stat = "0";
    var _ori_article_type = "";

    var is_follow = "";
    var nickname = "k8s技术圈";
    var appmsg_type = "9";
    var ct = "1562212007";
    var user_name = "gh_d6dd87b6ceb4";
    var user_name_new = "";
    var fakeid   = "";
    var version   = "";
    var is_limit_user   = "0";
    var round_head_img = "http://mmbiz.qpic.cn/mmbiz_png/z9BgVMEm7YvjtUXlclDSDibwlzlLMc91hwkaPKf1quWzkAeBUoK1J7F5xZqablItviciclvWL4ZGiabr7oIDPszZ1A/0?wx_fmt=png";
    var hd_head_img = "http://wx.qlogo.cn/mmhead/Q3auHgzwzM6rCeCHwQhxPNQNSXP5uZmTz1a2sVXQ7by9p0aSYVftlA/0"||"";
    var ori_head_img_url = "http://wx.qlogo.cn/mmhead/Q3auHgzwzM6rCeCHwQhxPNQNSXP5uZmTz1a2sVXQ7by9p0aSYVftlA/132";
    var msg_title = "Istio多集群管理方案详解";
    var msg_desc = "随着Kubernetes成为云原生领域应用编排的标准，传统企业及新型互联网企业都在逐步将应用容器化及云化。为";
    var msg_cdn_url = "http://mmbiz.qpic.cn/mmbiz_jpg/z9BgVMEm7YsbvT2Aia1aQOf56e01MiaTXArjGwT5mqic3Xxowdo449R78cISjFCiaiaHjMWUVlttUXpibBWY6YKp6Z7A/0?wx_fmt=jpeg"; 
    var cdn_url_1_1  = "https://mmbiz.qlogo.cn/mmbiz_jpg/z9BgVMEm7YsbvT2Aia1aQOf56e01MiaTXAjpmNURmHDjnicRnxlq8DBepoYrepIRHfpSEnEoBxvl9W1iam0fpRuJnw/0?wx_fmt=jpeg"; 
    var cdn_url_235_1 = "https://mmbiz.qlogo.cn/mmbiz_jpg/z9BgVMEm7YsbvT2Aia1aQOf56e01MiaTXArjGwT5mqic3Xxowdo449R78cISjFCiaiaHjMWUVlttUXpibBWY6YKp6Z7A/0?wx_fmt=jpeg"; 
    
    var msg_link = "http://mp.weixin.qq.com/s?__biz=MzU4MjQ0MTU4Ng==&amp;mid=2247484319&amp;idx=1&amp;sn=cb013878dad4a29a9a354d12a5135e4b&amp;chksm=fdb90c82cace859441501a3df68d058d726bdf221f6e66e5831fb9dc606fe7f0eeaf62213235#rd"; 
    var user_uin = "0"*1;
    var msg_source_url = '';
    var img_format = 'jpeg';
    var srcid = '';
    var req_id = '0416FsJw1QAukAfDm14XXrLf';
    var networkType;
    var appmsgid = '' || ''|| "2247484319";
    var comment_id = "882812679168098310" || "882812679168098310" * 1;
    var comment_enabled = "" * 1;
    var is_need_reward = "0" * 1;
    var is_https_res = ("" * 1) && (location.protocol == "https:");
    var msg_daily_idx = "1" || "";
    var profileReportInfo = "" || "";

    var devicetype = "";
    var source_encode_biz = "";
    var source_username = "";
    
    var reprint_ticket = "";
    var source_mid = "";
    var source_idx = "";
    var source_biz = "";
    var author_id = "";

    
    var optimizing_flag = "0" * 1;
    var ad_abtest_padding = "0" * 1;

    var show_comment = "0";
    var __appmsgCgiData = {
        wxa_product : ""*1,
        wxa_cps : ""*1,
        show_msg_voice: "0"*1,
        can_use_page : "0"*1,
        is_wxg_stuff_uin : "0"*1,
        card_pos : "",
        copyright_stat : "0",
        source_biz : "",
        hd_head_img : "http://wx.qlogo.cn/mmhead/Q3auHgzwzM6rCeCHwQhxPNQNSXP5uZmTz1a2sVXQ7by9p0aSYVftlA/0"||(window.location.protocol+"//"+window.location.host + "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/pic/appmsg/pic_rumor_link.2x42f400.jpg")
    };
    var _empty_v = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/pic/pages/voice/empty42f400.mp3";

    var copyright_stat = "0" * 1;

    var pay_fee = "" * 1;
    var pay_timestamp = "";
    var need_pay = "" * 1;

    var need_report_cost = "0" * 1;
    var use_tx_video_player = "0" * 1;
    var appmsg_fe_filter = "contenteditable";

    var friend_read_source = "" || "";
    var friend_read_version = "" || "";
    var friend_read_class_id = "" || "";

    var is_only_read = "1" * 1;
    var read_num = "" * 1;
    var like_num = "" * 1;
    var liked = "" == 'true' ? true : false;
    var is_temp_url = "" ? 1 : 0;
    var send_time = "";
    var icon_emotion_switch = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/appmsg/emotion/icon_emotion_switch46b604.svg";
    var icon_emotion_switch_active = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/appmsg/emotion/icon_emotion_switch_active46b604.svg";
    var icon_emotion_switch_primary = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/appmsg/emotion/icon_emotion_switch_primary46b604.svg";
    var icon_emotion_switch_active_primary = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/appmsg/emotion/icon_emotion_switch_active_primary46b604.svg";
    var icon_loading_white = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/common/icon_loading_white42f400.gif";
    var icon_audio_unread = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/appmsg/audio/icon_audio_unread42f400.png";
    var icon_qqmusic_default = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/appmsg/qqmusic/icon_qqmusic_default.2x42f400.png";
    var icon_qqmusic_source = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/appmsg/qqmusic/icon_qqmusic_source42f400.png";
    var icon_kugou_source = "//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/appmsg/kugou/icon_kugou_source42f400.png";

    var topic_default_img = '//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/appmsg/topic/pic_book_thumb.2x42f400.png';
    var comment_edit_icon = '//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/appmsg_new/icon_edit42f400.png';
    var comment_loading_img = '//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/icon/common/icon_loading_white42f400.gif';

    var voice_in_appmsg = {
        "1":"1"
            };

    var reprint_style = ''*1;
    var wxa_img_alert = "" != 'false';

    var img_popup = 0;

    
    var more_read_type = '0'*1;

    
    
    
    

    
    var weapp_sn_arr_json = "" || "";

    
    var ban_scene = "0" * 1;

    var svr_time = "1562229368" * 1;
    
    var is_transfer_msg = ""*1||0;

    var malicious_title_reason_id = "0" * 1; 
    var malicious_content_type = "0" * 1; 

    
    var modify_time = "";

    
    var isprofileblock = "0";

    
    var hotspotInfoList = [
                    ];

    var jumpInfo = [
                    ];

        window.wxtoken = "777";
        
    
    
    
    
    window.is_login = '0' * 1; 

    window.__moon_initcallback = function(){
        if(!!window.__initCatch){
            window.__initCatch({
                idkey : 27611+2,
                startKey : 0,
                limit : 128,
                badjsId: 43,
                reportOpt : {
                    uin : uin,
                    biz : biz,
                    mid : mid,
                    idx : idx,
                    sn  : sn
                },
                extInfo : {
                    network_rate : 0.01,    
                    badjs_rate: 0.1 
                }
            });
        }
    }
        
    var title ="k8s技术圈";

    var is_new_msg=true;
    
    

    var is_wash = '' * 1;
    var show_top_bar  = '0' * 1;
    var topbarEnable = false;
    var enterid=""*1||0;

    var defaultAvatarUrl = '//res.wx.qq.com/mmbizwap/zh_CN/htmledition/images/pic/common/avatar_default46e3e2.svg';
</script>

<script nonce="1977935583" type="text/javascript">
(function(_g){
    _g.appmsg_like_type = "2" * 1 ? "2" * 1 : 1;
    
    _g.clientversion = "";
    _g.passparam = ""; 
    if(!_g.msg_link) {
      _g.msg_link = "http://mp.weixin.qq.com/s?__biz=MzU4MjQ0MTU4Ng==&amp;mid=2247484319&amp;idx=1&amp;sn=cb013878dad4a29a9a354d12a5135e4b&amp;chksm=fdb90c82cace859441501a3df68d058d726bdf221f6e66e5831fb9dc606fe7f0eeaf62213235#rd";
    }
    _g.appmsg_type = "9"; 
    _g.devicetype = ""; 
})(window);

</script>


        <script nonce="1977935583">window.__moon_host = 'res.wx.qq.com';window.__moon_mainjs = 'appmsg/index.js';window.moon_map = {"pages/iframe_communicate.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/iframe_communicate42f400.js","new_video/player.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/new_video/player.html433882.js","biz_wap/jsapi/log.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/jsapi/log4673d5.js","biz_wap/zepto/touch.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/zepto/touch42f400.js","biz_wap/zepto/event.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/zepto/event42f400.js","biz_wap/zepto/zepto.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/zepto/zepto440203.js","page/pages/video.css":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/style/page/pages/video.css46b604.js","a/tpl/smallbanner_msg_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/smallbanner_msg_tpl.html42f400.js","a/tpl/smallbanner_info_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/smallbanner_info_tpl.html44c2e3.js","a/tpl/banner_info_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/banner_info_tpl.html42f400.js","a/tpl/promote_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/promote_tpl.html42f400.js","a/tpl/smallcard_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/smallcard_tpl.html42f400.js","a/tpl/info_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/info_tpl.html42f400.js","a/tpl/cardticket_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/cardticket_tpl.html42f400.js","a/tpl/banner_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/banner_tpl.html46e9c3.js","a/tpl/sponsor_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/sponsor_tpl.html42f400.js","a/tpl/new_cpc_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/new_cpc_tpl.html45178d.js","appmsg/emotion/caret.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/emotion/caret42f400.js","pages/audition_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/audition_tpl.html44ccb4.js","biz_wap/utils/localstorage.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/localstorage42f400.js","appmsg/friend_comment_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/friend_comment_tpl.html42f400.js","appmsg/comment_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/comment_tpl.html45a3cd.js","biz_wap/utils/fakehash.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/fakehash42f400.js","appmsg/comment_report.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/comment_report4690d8.js","a/appdialog_confirm.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/appdialog_confirm.html42f400.js","widget/wx_profile_dialog_primary.css":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/style/widget/wx_profile_dialog_primary.css42f400.js","new_video/player.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/new_video/player46e9c3.js","a/tpl/mpda_bottom_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/mpda_bottom_tpl.html450c68.js","a/tpl/crt_size_map.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/crt_size_map4602fc.js","biz_wap/jsapi/cardticket.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/jsapi/cardticket42f400.js","biz_common/jquery.md5.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/jquery.md542f400.js","biz_common/utils/emoji_panel_data.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/utils/emoji_panel_data42f400.js","appmsg/emotion/textarea.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/emotion/textarea42f400.js","appmsg/emotion/nav.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/emotion/nav42f400.js","appmsg/emotion/common.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/emotion/common42f400.js","appmsg/emotion/slide.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/emotion/slide42f400.js","appmsg/emotion/dom.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/emotion/dom42f400.js","pages/player_tips.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/player_tips44ccb4.js","pages/music_report_conf.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/music_report_conf42f400.js","pages/report.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/report4629a1.js","pages/player_adaptor.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/player_adaptor42f400.js","pages/music_player.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/music_player42f400.js","biz_common/utils/emoji_data.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/utils/emoji_data45112f.js","appmsg/more_read_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/more_read_tpl.html42f400.js","appmsg/i18n.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/i18n45e203.js","appmsg/retry_ajax.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/retry_ajax451cc4.js","complain/tips.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/complain/tips42f400.js","pages/loadscript.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/loadscript42f400.js","biz_wap/utils/ajax_load_js.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/ajax_load_js42f400.js","appmsg/comment.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/comment46b604.js","appmsg/reward_entry.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/reward_entry46ef12.js","a/ios.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/ios42f400.js","a/android.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/android457bcb.js","a/profile.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/profile455ab4.js","a/app_card.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/app_card455a33.js","a/sponsor.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/sponsor4576f8.js","a/tpl/cpc_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/cpc_tpl.html450c68.js","a/appdialog_confirm.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/appdialog_confirm464cef.js","biz_common/dom/offset.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/dom/offset4690d8.js","a/video.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/video46e9c3.js","a/tpl/crt_tpl_manager.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/tpl/crt_tpl_manager450d79.js","a/cpc_a_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/cpc_a_tpl.html450c68.js","a/sponsor_a_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/sponsor_a_tpl.html42f400.js","a/a_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/a_tpl.html450c68.js","a/mpshop.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/mpshop42f400.js","a/wxopen_card.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/wxopen_card42f400.js","a/card.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/card42f400.js","biz_wap/utils/position.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/position42f400.js","a/a_report.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/a_report4402ec.js","biz_wap/utils/show_time.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/show_time4543c6.js","biz_common/utils/get_para_list.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/utils/get_para_list42f400.js","biz_wap/utils/openUrl.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/openUrl4402ec.js","a/a_sign.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/a_sign452c49.js","appmsg/my_comment_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/my_comment_tpl.html46b604.js","appmsg/cmt_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/cmt_tpl.html46b604.js","sougou/a_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/sougou/a_tpl.html42f400.js","appmsg/emotion/emotion.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/emotion/emotion46b604.js","biz_common/base64.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/base6442f400.js","biz_common/utils/report.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/utils/report42f400.js","appmsg/articleReport.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/articleReport42f400.js","biz_wap/jsapi/leaveReport.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/jsapi/leaveReport45a434.js","biz_wap/utils/hand_up_state.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/hand_up_state42f400.js","biz_wap/utils/storage.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/storage42f400.js","biz_common/utils/http.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/utils/http42f400.js","biz_common/utils/cookie.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/utils/cookie42f400.js","appmsg/topic_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/topic_tpl.html42f400.js","question_answer/appmsg_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/question_answer/appmsg_tpl.html45112f.js","pages/weapp_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/weapp_tpl.html42f400.js","biz_common/utils/monitor.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/utils/monitor42f400.js","pages/voice_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/voice_tpl.html42f400.js","pages/kugoumusic_ctrl.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/kugoumusic_ctrl42f400.js","pages/qqmusic_ctrl.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/qqmusic_ctrl44df7a.js","pages/voice_component.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/voice_component44df7a.js","pages/qqmusic_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/qqmusic_tpl.html42f400.js","new_video/ctl.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/new_video/ctl4532b3.js","pages/utils.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/utils45112f.js","appmsg/open_url_with_webview.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/open_url_with_webview440203.js","appmsg/more_read.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/more_read4576f8.js","appmsg/like.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/like46c6ac.js","appmsg/share_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/share_tpl.html42f400.js","appmsg/appmsgext.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/appmsgext430b95.js","appmsg/img_copyright_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/img_copyright_tpl.html42f400.js","pages/video_ctrl.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/video_ctrl42f400.js","pages/create_txv.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/create_txv42f400.js","appmsg/comment_utils.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/comment_utils42f400.js","appmsg/reward_utils.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/reward_utils46e54d.js","biz_common/ui/imgonepx.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/ui/imgonepx42f400.js","appmsg/malicious_wording.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/malicious_wording42f400.js","biz_common/utils/wxgspeedsdk.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/utils/wxgspeedsdk42f400.js","pages/version4video.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/version4video46df38.js","a/a_config.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/a_config4690d8.js","a/a_utils.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/a_utils46e9c3.js","a/a.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/a46edff.js","rt/appmsg/getappmsgext.rt.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/rt/appmsg/getappmsgext.rt42f400.js","pages/video_communicate_adaptor.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/pages/video_communicate_adaptor433882.js","biz_wap/utils/ajax_wx.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/ajax_wx42f400.js","biz_common/utils/respTypes.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/utils/respTypes42f400.js","biz_wap/utils/log.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/log42f400.js","sougou/index.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/sougou/index42f400.js","biz_wap/safe/mutation_observer_report.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/safe/mutation_observer_report42f400.js","appmsg/fereport.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/fereport438bee.js","appmsg/fereport_without_localstorage.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/fereport_without_localstorage438bee.js","appmsg/report.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/report4576f8.js","appmsg/report_and_source.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/report_and_source450c68.js","appmsg/page_pos.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/page_pos46e4c3.js","appmsg/cdn_speed_report.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/cdn_speed_report42f400.js","appmsg/wxtopic.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/wxtopic42f400.js","question_answer/appmsg.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/question_answer/appmsg45112f.js","appmsg/weapp.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/weapp46a084.js","appmsg/weproduct.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/weproduct4576f8.js","appmsg/voicemsg.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/voicemsg42f400.js","appmsg/autoread.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/autoread42f400.js","appmsg/voice.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/voice42f400.js","appmsg/qqmusic.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/qqmusic43440b.js","appmsg/iframe.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/iframe46ea12.js","question_answer/utils.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/question_answer/utils45112f.js","appmsg/product.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/product4576f8.js","appmsg/review_image.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/review_image46a084.js","appmsg/outer_link.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/outer_link46a084.js","appmsg/copyright_report.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/copyright_report4576f8.js","appmsg/async.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/async46e3e3.js","biz_wap/ui/lazyload_img.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/ui/lazyload_img42f400.js","biz_common/log/jserr.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/log/jserr42f400.js","appmsg/share.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/share42f50a.js","appmsg/cdn_img_lib.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/cdn_img_lib42f400.js","appmsg/finance_communicate.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/finance_communicate453348.js","page/appmsg_new/not_in_mm.css":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/style/page/appmsg_new/not_in_mm.css46b604.js","page/appmsg_new/combo.css":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/style/page/appmsg_new/combo.css46e3e3.js","complain/localstorage.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/complain/localstorage42f400.js","common/utils.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/common/utils46e9c3.js","biz_wap/utils/wapsdk.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/wapsdk44c130.js","a/mpAdAsync.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/a/mpAdAsync464cef.js","biz_common/utils/url/parse.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/utils/url/parse440451.js","appmsg/appmsg_report.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/appmsg_report435ac2.js","biz_common/moment.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/moment42f400.js","biz_wap/jsapi/core.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/jsapi/core4673d5.js","biz_common/dom/event.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/dom/event445789.js","appmsg/test.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/test42f400.js","biz_wap/utils/mmversion.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/mmversion45fc7f.js","appmsg/max_age.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/max_age42f400.js","biz_common/dom/attr.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/dom/attr42f400.js","biz_wap/utils/ajax.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/ajax462fbb.js","appmsg/log.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/log42f400.js","biz_common/dom/class.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/dom/class42f400.js","biz_wap/utils/device.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_wap/utils/device42f400.js","appmsg/weapp_common.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/weapp_common42f400.js","biz_common/utils/string/html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/utils/string/html42f400.js","cps/tpl/list_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/cps/tpl/list_tpl.html42f400.js","cps/tpl/card_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/cps/tpl/card_tpl.html42f400.js","cps/tpl/banner_tpl.html.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/cps/tpl/banner_tpl.html42f400.js","biz_common/tmpl.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/biz_common/tmpl42f400.js","appmsg/index.js":"//res.wx.qq.com/mmbizwap/zh_CN/htmledition/js/appmsg/index466d08.js"};</script><script nonce="1977935583" type="text/javascript" id="moon_inline" > window.__mooninline=1; window.setTimeout(function() {  function __moonf__(){
if(!window.__moonhasinit){
window.__moonhasinit=!0,window.__moonclientlog=[],window.__wxgspeeds&&(window.__wxgspeeds.moonloadedtime=+new Date),
"object"!=typeof JSON&&(window.JSON={
stringify:function(){
return"";
},
parse:function(){
return{};
}
});
var e=function(){
function e(e){
try{
var o;
/(iPhone|iPad|iPod|iOS)/i.test(navigator.userAgent)?o="writeLog":/(Android)/i.test(navigator.userAgent)&&(o="log"),
o&&t(o,e);
}catch(n){
throw console.error(n),n;
}
}
function t(e,o){
var n,r,i={};
n=top!=window?top.window:window;
try{
r=n.WeixinJSBridge,i=n.document;
}catch(a){}
e&&r&&r.invoke?r.invoke(e,{
level:"info",
msg:"[WechatFe][moon]"+o
}):setTimeout(function(){
i.addEventListener?i.addEventListener("WeixinJSBridgeReady",function(){
t(e,o);
},!1):i.attachEvent&&(i.attachEvent("WeixinJSBridgeReady",function(){
t(e,o);
}),i.attachEvent("onWeixinJSBridgeReady",function(){
t(e,o);
}));
},0);
}
var n;
localStorage&&JSON.parse(localStorage.getItem("__WXLS__moonarg"))&&"fromls"==JSON.parse(localStorage.getItem("__WXLS__moonarg")).method&&(n=!0),
e(" moon init, moon_inline:"+window.__mooninline+", moonls:"+n),function(){
var e={},o={},t={};
e.COMBO_UNLOAD=0,e.COMBO_LOADING=1,e.COMBO_LOADED=2;
var n=function(e,t,n){
if(!o[e]){
o[e]=n;
for(var r=3;r--;)try{
moon.setItem(moon.prefix+e,n.toString()),moon.setItem(moon.prefix+e+"_ver",moon_map[e]);
break;
}catch(i){
moon.clear();
}
}
},r=window.alert;
window.__alertList=[],window.alert=function(e){
r(e),window.__alertList.push(e);
};
var i=function(e){
if(!e||!o[e])return null;
var n=o[e];
if("function"==typeof n&&!t[e]){
var a={},s={
exports:a
},c=n(i,a,s,r);
n=o[e]=c||s.exports,t[e]=!0;
}
if(".css"===e.substr(-4)){
var d=document.getElementById(e);
if(!d){
d=document.createElement("style"),d.id=e;
var _=/url\s*\(\s*\/(\"(?:[^\\\"\r\n\f]|\\[\s\S])*\"|'(?:[^\\'\n\r\f]|\\[\s\S])*'|[^)}]+)\s*\)/g,m=window.testenv_reshost||window.__moon_host||"res.wx.qq.com";
n=n.replace(_,"url(//"+m+"/$1)"),d.innerHTML=n,document.getElementsByTagName("head")[0].appendChild(d);
}
}
return n;
};
e.combo_status=e.COMBO_UNLOAD,e.run=function(){
var o=e.run.info,t=o&&o[0],n=o&&o[1];
if(t&&e.combo_status==e.COMBO_LOADED){
var r=i(t);
n&&n(r);
}
},e.use=function(o,t){
window.__wxgspeeds&&(window.__wxgspeeds.seajs_use_time=+new Date),e.run.info=[o,t],
e.run();
},window.define=n,window.seajs=e;
}(),function(){
if(window.__nonce_str){
var e=document.createElement;
document.createElement=function(o){
var t=e.apply(this,arguments);
return"object"==typeof o&&(o=o.toString()),"string"==typeof o&&"script"==o.toLowerCase()&&t.setAttribute("nonce",window.__nonce_str),
t;
};
}
window.addEventListener&&window.__DEBUGINFO&&Math.random()<.01&&window.addEventListener("load",function(){
var e=document.createElement("script");
e.src=__DEBUGINFO.safe_js,e.type="text/javascript",e.async=!0;
var o=document.head||document.getElementsByTagName("head")[0];
o.appendChild(e);
});
}(),function(){
function t(e){
return"[object Array]"===Object.prototype.toString.call(e);
}
function n(e){
return"[object Object]"===Object.prototype.toString.call(e);
}
function r(e){
var t=e.stack+" "+e.toString()||"";
try{
if(window.testenv_reshost){
var n="http(s)?://"+window.testenv_reshost,r=new RegExp(n,"g");
t=t.replace(r,"");
}else t=t.replace(/http(s)?:\/\/res\.wx\.qq\.com/g,"");
for(var r=/\/([^.]+)\/js\/(\S+?)\.js(\,|:)?/g;r.test(t);)t=t.replace(r,function(e,o,t,n){
return t+n;
});
}catch(e){
t=e.stack?e.stack:"";
}
var i=[];
for(o in u)u.hasOwnProperty(o)&&i.push(o+":"+u[o]);
return i.push("STK:"+t.replace(/\n/g,"")),i.join("|");
}
function i(e){
if(!e){
var o=window.onerror;
window.onerror=function(){},f=setTimeout(function(){
window.onerror=o,f=null;
},50);
}
}
function a(e,o,t){
if(!/^mp\.weixin\.qq\.com$/.test(location.hostname)){
var n=[];
t=t.replace(location.href,(location.origin||"")+(location.pathname||"")).replace("#wechat_redirect","").replace("#rd","").split("&");
for(var r=0,i=t.length;i>r;r++){
var a=t[r].split("=");
a[0]&&a[1]&&n.push(a[0]+"="+encodeURIComponent(a[1]));
}
var s=new window.Image;
return void(s.src=(o+n.join("&")).substr(0,1024));
}
var c;
if(window.ActiveXObject)try{
c=new ActiveXObject("Msxml2.XMLHTTP");
}catch(d){
try{
c=new ActiveXObject("Microsoft.XMLHTTP");
}catch(_){
c=!1;
}
}else window.XMLHttpRequest&&(c=new XMLHttpRequest);
c&&(c.open(e,o,!0),c.setRequestHeader("cache-control","no-cache"),c.setRequestHeader("Content-Type","application/x-www-form-urlencoded; charset=UTF-8"),
c.setRequestHeader("X-Requested-With","XMLHttpRequest"),c.send(t));
}
function s(e){
return function(o,t){
if("string"==typeof o)try{
o=new Function(o);
}catch(n){
throw n;
}
var r=[].slice.call(arguments,2),a=o;
return o=function(){
try{
return a.apply(this,r.length&&r||arguments);
}catch(e){
throw e.stack&&console&&console.error&&console.error("[TryCatch]"+e.stack),_&&window.__moon_report&&(window.__moon_report([{
offset:j,
log:"timeout_error;host:"+location.host,
e:e
}]),i(f)),e;
}
},e(o,t);
};
}
function c(e){
return function(o,t,n){
if("undefined"==typeof n)var n=!1;
var r=this,a=t||function(){};
return t=function(){
try{
return a.apply(r,arguments);
}catch(e){
throw e.stack&&console&&console.error&&console.error("[TryCatch]"+e.stack),_&&window.__moon_report&&(window.__moon_report([{
offset:x,
log:"listener_error;type:"+o+";host:"+location.host,
e:e
}]),i(f)),e;
}
},a.moon_lid=E,b[E]=t,E++,e.call(r,o,t,n);
};
}
function d(e){
return function(o,t,n){
if("undefined"==typeof n)var n=!1;
var r=this;
return t=b[t.moon_lid],e.call(r,o,t,n);
};
}
var _,m,l,w,u,p,f,h=/MicroMessenger/i.test(navigator.userAgent),g=/MPAPP/i.test(navigator.userAgent),v=window.define,y=0,x=2,O=4,j=9,D=10;
if(window.__initCatch=function(e){
_=e.idkey,m=e.startKey||0,l=e.limit,w=e.badjsId,u=e.reportOpt||"",p=e.extInfo||{},
p.rate=p.rate||.5;
},window.__moon_report=function(e,o){
var i=!1,s="";
try{
s=top.location.href;
}catch(c){
i=!0;
}
var d=.5;
if(p&&p.rate&&(d=p.rate),o&&"number"==typeof o&&(d=o),!(!/mp\.weixin\.qq\.com/.test(location.href)&&!/payapp\.weixin\.qq\.com/.test(location.href)||Math.random()>d||!h&&!g||top!=window&&!i&&!/mp\.weixin\.qq\.com/.test(s))&&(n(e)&&(e=[e]),
t(e)&&""!=_)){
var u="",f=[],v=[],y=[],x=[];
"number"!=typeof l&&(l=1/0);
for(var j=0;j<e.length;j++){
var D=e[j]||{};
if(!(D.offset>l||"number"!=typeof D.offset||D.offset==O&&p&&p.network_rate&&Math.random()>=p.network_rate)){
var b=1/0==l?m:m+D.offset;
f[j]="[moon]"+_+"_"+b+";"+D.log+";"+r(D.e||{})||"",v[j]=b,y[j]=1;
}
}
for(var E=0;E<v.length;E++)x[E]=_+"_"+v[E]+"_"+y[E],u=u+"&log"+E+"="+f[E];
if(x.length>0){
a("POST",location.protocol+"//mp.weixin.qq.com/mp/jsmonitor?","idkey="+x.join(";")+"&r="+Math.random()+"&lc="+f.length+u);
var d=1;
if(p&&p.badjs_rate&&(d=p.badjs_rate),w&&Math.random()<d){
u=u.replace(/uin\:(.)*\|biz\:(.)*\|mid\:(.)*\|idx\:(.)*\|sn\:(.)*\|/,"");
var S=new Image,I="https://badjs.weixinbridge.com/badjs?id="+w+"&level=4&from="+encodeURIComponent(location.host)+"&msg="+encodeURIComponent(u);
S.src=I.slice(0,1024);
}
}
}
},window.setTimeout=s(window.setTimeout),window.setInterval=s(window.setInterval),
Math.random()<.01&&window.Document&&window.HTMLElement){
var b={},E=0;
Document.prototype.addEventListener=c(Document.prototype.addEventListener),Document.prototype.removeEventListener=d(Document.prototype.removeEventListener),
HTMLElement.prototype.addEventListener=c(HTMLElement.prototype.addEventListener),
HTMLElement.prototype.removeEventListener=d(HTMLElement.prototype.removeEventListener);
}
var S=window.navigator.userAgent;
if((/ip(hone|ad|od)/i.test(S)||/android/i.test(S))&&!/windows phone/i.test(S)&&window.localStorage&&window.localStorage.setItem){
var I=window.localStorage.setItem,L=0;
window.localStorage.setItem=function(e,o){
if(!(L>=10))try{
I.call(window.localStorage,e,o);
}catch(t){
t.stack&&console&&console.error&&console.error("[TryCatch]"+t.stack),window.__moon_report([{
offset:D,
log:"localstorage_error;"+t.toString(),
e:t
}]),L++,L>=3&&window.moon&&window.moon.clear&&moon.clear();
}
};
}
window.seajs&&v&&(window.define=function(){
for(var o,t=[],n=arguments&&arguments[0],a=0,s=arguments.length;s>a;a++){
var c=o=arguments[a];
"function"==typeof o&&(o=function(){
try{
return c.apply(this,arguments);
}catch(o){
throw"string"==typeof n&&console.error("[TryCatch][DefineeErr]id:"+n),o.stack&&console&&console.error&&console.error("[TryCatch]"+o.stack),
_&&window.__moon_report&&(window.__moon_report([{
offset:y,
log:"define_error;id:"+n+";",
e:o
}]),i(f)),e(" [define_error]"+JSON.stringify(r(o))),o;
}
},o.toString=function(e){
return function(){
return e.toString();
};
}(arguments[a])),t.push(o);
}
return v.apply(this,t);
});
}(),function(o){
function t(e,o,t){
return window.__DEBUGINFO?(window.__DEBUGINFO.res_list||(window.__DEBUGINFO.res_list=[]),
window.__DEBUGINFO.res_list[e]?(window.__DEBUGINFO.res_list[e][o]=t,!0):!1):!1;
}
function n(e){
var o=new TextEncoder("utf-8").encode(e),t=crypto.subtle||crypto.webkitSubtle;
return t.digest("SHA-256",o).then(function(e){
return r(e);
});
}
function r(e){
for(var o=[],t=new DataView(e),n=0;n<t.byteLength;n+=4){
var r=t.getUint32(n),i=r.toString(16),a="00000000",s=(a+i).slice(-a.length);
o.push(s);
}
return o.join("");
}
function i(e,o,t){
if("object"==typeof e){
var n=Object.prototype.toString.call(e).replace(/^\[object (.+)\]$/,function(e,o){
return o;
});
if(t=t||e,"Array"==n){
for(var r=0,i=e.length;i>r;++r)if(o.call(t,e[r],r,e)===!1)return;
}else{
if("Object"!==n&&a!=e)throw"unsupport type";
if(a==e){
for(var r=e.length-1;r>=0;r--){
var s=a.key(r),c=a.getItem(s);
if(o.call(t,c,s,e)===!1)return;
}
return;
}
for(var r in e)if(e.hasOwnProperty(r)&&o.call(t,e[r],r,e)===!1)return;
}
}
}
var a=o.localStorage,s=document.head||document.getElementsByTagName("head")[0],c=1,d=11,_=12,m=13,l=window.__allowLoadResFromMp?1:2,w=window.__allowLoadResFromMp?1:0,u=l+w,p=window.testenv_reshost||window.__moon_host||"res.wx.qq.com",f=new RegExp("^(http(s)?:)?//"+p);
window.__loadAllResFromMp&&(p="mp.weixin.qq.com",l=0,u=l+w);
var h={
prefix:"__MOON__",
loaded:[],
unload:[],
clearSample:!0,
hit_num:0,
mod_num:0,
version:1003,
cacheData:{
js_mod_num:0,
js_hit_num:0,
js_not_hit_num:0,
js_expired_num:0,
css_mod_num:0,
css_hit_num:0,
css_not_hit_num:0,
css_expired_num:0
},
init:function(){
h.loaded=[],h.unload=[];
var e,t,r;
if(window.no_moon_ls&&(h.clearSample=!0),a){
var s="_moon_ver_key_",c=a.getItem(s);
c!=h.version&&(h.clear(),a.setItem(s,h.version));
}
if((-1!=location.search.indexOf("no_moon1=1")||-1!=location.search.indexOf("no_lshttps=1"))&&h.clear(),
a){
var d=1*a.getItem(h.prefix+"clean_time"),_=+new Date;
if(_-d>=1296e6){
h.clear();
try{
!!a&&a.setItem(h.prefix+"clean_time",+new Date);
}catch(m){}
}
}
i(moon_map,function(i,s){
if(t=h.prefix+s,r=!!i&&i.replace(f,""),e=!!a&&a.getItem(t),version=!!a&&(a.getItem(t+"_ver")||"").replace(f,""),
h.mod_num++,r&&-1!=r.indexOf(".css")?h.cacheData.css_mod_num++:r&&-1!=r.indexOf(".js")&&h.cacheData.js_mod_num++,
h.clearSample||!e||r!=version)h.unload.push(r.replace(f,"")),r&&-1!=r.indexOf(".css")?e?r!=version&&h.cacheData.css_expired_num++:h.cacheData.css_not_hit_num++:r&&-1!=r.indexOf(".js")&&(e?r!=version&&h.cacheData.js_expired_num++:h.cacheData.js_not_hit_num++);else{
if("https:"==location.protocol&&window.moon_hash_map&&window.moon_hash_map[s]&&window.crypto)try{
n(e).then(function(e){
window.moon_hash_map[s]!=e&&console.log(s);
});
}catch(c){}
try{
var d="//# sourceURL="+s+"\n//@ sourceURL="+s;
o.eval.call(o,'define("'+s+'",[],'+e+")"+d),h.hit_num++,r&&-1!=r.indexOf(".css")?h.cacheData.css_hit_num++:r&&-1!=r.indexOf(".js")&&h.cacheData.js_hit_num++;
}catch(c){
h.unload.push(r.replace(f,""));
}
}
}),h.load(h.genUrl());
},
genUrl:function(){
var e=h.unload;
if(!e||e.length<=0)return[];
var o,t,n="",r=[],i={},a=-1!=location.search.indexOf("no_moon2=1"),s="//"+p;
-1!=location.href.indexOf("moon_debug2=1")&&(s="//mp.weixin.qq.com");
for(var c=0,d=e.length;d>c;++c){
/^\/(.*?)\//.test(e[c]);
var _=/^\/(.*?)\//.exec(e[c]);
_.length<2||!_[1]||(t=_[1],n=i[t],n?(o=n+","+e[c],o.length>1e3||a?(r.push(n+"?v="+h.version),
n=location.protocol+s+e[c],i[t]=n):(n=o,i[t]=n)):(n=location.protocol+s+e[c],i[t]=n));
}
for(var m in i)i.hasOwnProperty(m)&&r.push(i[m]);
return r;
},
load:function(e){
if(window.__wxgspeeds&&(window.__wxgspeeds.mod_num=h.mod_num,window.__wxgspeeds.hit_num=h.hit_num),
!e||e.length<=0)return seajs.combo_status=seajs.COMBO_LOADED,seajs.run(),console.debug&&console.debug("[moon] load js complete, all in cache, cost time : 0ms, total count : "+h.mod_num+", hit num: "+h.hit_num),
void window.__moonclientlog.push("[moon] load js complete, all in cache, cost time : 0ms, total count : "+h.mod_num+", hit num: "+h.hit_num);
seajs.combo_status=seajs.COMBO_LOADING;
var o=0,t=+new Date;
window.__wxgspeeds&&(window.__wxgspeeds.combo_times=[],window.__wxgspeeds.combo_times.push(t)),
i(e,function(n){
h.request(n,u,function(){
if(window.__wxgspeeds&&window.__wxgspeeds.combo_times.push(+new Date),o++,o==e.length){
var n=+new Date-t;
window.__wxgspeeds&&(window.__wxgspeeds.mod_downloadtime=n),seajs.combo_status=seajs.COMBO_LOADED,
seajs.run(),console.debug&&console.debug("[moon] load js complete, url num : "+e.length+", total mod count : "+h.mod_num+", hit num: "+h.hit_num+", use time : "+n+"ms"),
window.__moonclientlog.push("[moon] load js complete, url num : "+e.length+", total mod count : "+h.mod_num+", hit num: "+h.hit_num+", use time : "+n+"ms");
}
});
});
},
request:function(o,n,r){
if(o){
n=n||0,o.indexOf("mp.weixin.qq.com")>-1&&((new Image).src=location.protocol+"//mp.weixin.qq.com/mp/jsmonitor?idkey=27613_32_1&r="+Math.random(),
window.__moon_report([{
offset:_,
log:"load_script_from_mp: "+o
}],1));
var i=-1;
window.__DEBUGINFO&&(__DEBUGINFO.res_list||(__DEBUGINFO.res_list=[]),__DEBUGINFO.res_list.push({
type:"js",
status:"pendding",
start:+new Date,
end:0,
url:o
}),i=__DEBUGINFO.res_list.length-1),-1!=location.search.indexOf("no_lshttps=1")&&(o=o.replace("http://","https://"));
var a=document.createElement("script");
a.src=o,a.type="text/javascript",a.async=!0,a.down_time=+new Date,a.onerror=function(s){
t(i,"status","error"),t(i,"end",+new Date);
var _=new Error(s);
if(n>=0)if(w>n){
var l=o.replace("res.wx.qq.com","mp.weixin.qq.com");
h.request(l,n,r);
}else h.request(o,n,r);else window.__moon_report&&window.__moon_report([{
offset:c,
log:"load_script_error: "+o,
e:_
}],1);
if(n==w-1&&window.__moon_report([{
offset:d,
log:"load_script_error: "+o,
e:_
}],1),-1==n){
var u="ua: "+window.navigator.userAgent+", time="+(+new Date-a.down_time)+", load_script_error -1 : "+o;
window.__moon_report([{
offset:m,
log:u
}],1);
}
window.__moonclientlog.push("moon load js error : "+o+", error -> "+_.toString()),
e("moon_request_error url:"+o);
},"undefined"!=typeof moon_crossorigin&&moon_crossorigin&&a.setAttribute("crossorigin",!0),
a.onload=a.onreadystatechange=function(){
t(i,"status","loaded"),t(i,"end",+new Date),!a||a.readyState&&!/loaded|complete/.test(a.readyState)||(t(i,"status","200"),
a.onload=a.onreadystatechange=null,"function"==typeof r&&r());
},n--,s.appendChild(a),e("moon_request url:"+o+" retry:"+n);
}
},
setItem:function(e,o){
!!a&&a.setItem(e,o);
},
clear:function(){
a&&(i(a,function(e,o){
~o.indexOf(h.prefix)&&a.removeItem(o);
}),console.debug&&console.debug("[moon] clear"));
},
idkeyReport:function(e,o,t){
t=t||1;
var n=e+"_"+o+"_"+t;
(new Image).src="/mp/jsmonitor?idkey="+n+"&r="+Math.random();
}
};
seajs&&seajs.use&&"string"==typeof window.__moon_mainjs&&seajs.use(window.__moon_mainjs),
window.moon=h;
}(window),function(){
try{
Math.random()<1;
}catch(e){}
}(),window.moon.init();
};
e(),!!window.__moon_initcallback&&window.__moon_initcallback(),window.__wxgspeeds&&(window.__wxgspeeds.moonendtime=+new Date);
}
}
__moonf__(); }, 25);</script><script nonce="1977935583" type="text/javascript">
    var real_show_page_time = +new Date();
    if (!!window.addEventListener){
        window.addEventListener("load", function(){
            window.onload_endtime = +new Date();
        });
    }

    
</script>

    </body>
</html>

`
