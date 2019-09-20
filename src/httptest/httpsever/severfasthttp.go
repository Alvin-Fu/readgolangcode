package httpsever

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

// index 页
func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome")
}

// 简单路由页
func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello")
}

// 获取GET请求json数据
// 使用 ctx.QueryArgs() 方法
// Peek类似与python中dict的pop方法，取某个键对应的值
func TestGet(ctx *fasthttp.RequestCtx) {
	values := ctx.QueryArgs()
	fmt.Fprint(ctx, string(values.Peek("abc"))) // 不加string返回的byte数组

}

// 获取post的请求json数据
// 这里就有点坑是，查了很多网页说可以用 ctx.PostArgs() 取post的参数，返现不行，返回空
// 后来用 ctx.FormValue() 取表单数据就好了，难道是版本升级的问题？
// ctx.PostBody() 在上传文件的时候比较有用
func TestPost(ctx *fasthttp.RequestCtx) {
	//postValues := ctx.PostArgs()
	//fmt.Fprint(ctx, string(postValues))

	// 获取表单数据
	fmt.Fprint(ctx, string(ctx.FormValue("abc")))

	// 这两行可以获取PostBody数据，在上传数据文件的时候有用
	postBody := ctx.PostBody()
	fmt.Fprint(ctx, string(postBody))
}
