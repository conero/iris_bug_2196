package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func main() {

	app := iris.New()
	//app.Use(iris.Compression)
	const sxKey = "jkkkkl"

	sess := sessions.New(sessions.Config{
		Cookie: "test-jc",
		// CookieSecureTLS: true,
		Expires:      -1,
		AllowReclaim: true,
	})

	app.Use(sess.Handler())
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("Hello <strong>%s</strong>!", "World")
	})

	//session
	app.Get("/add", func(ctx iris.Context) {
		sx := sessions.Get(ctx)

		count, _ := sx.GetInt64(sxKey)

		sx.Set(sxKey, count+1)
		ctx.JSON(map[string]any{
			"type": "set",
			"count": count,
		})
	})

	//
	app.Get("/get", func(ctx iris.Context) {
		sx := sessions.Get(ctx)

		count, _ := sx.GetInt64(sxKey)
		ctx.JSON(map[string]any{
			"type": "get",
			"count": count,
		})
	})

	app.Listen(":8081")

}
