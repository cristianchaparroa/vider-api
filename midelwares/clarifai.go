package midelwares

import (
	"github.com/kataras/iris"
	"github.com/clarifai/clarifai-go"
)



func ClarifaiMiddleware(cf *clarifai.Client ) iris.HandlerFunc {
	_,err := cf.Info()
	if err != nil {
		panic(err)
	}

	return func(ctx *iris.Context) {
		ctx.Set("cf", cf)
		ctx.Next()
	}
}