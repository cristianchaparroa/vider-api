package api

import (
	"github.com/kataras/iris"
	"github.com/clarifai/clarifai-go"
	"os"
	"fmt"
	"io"
)


func Tags(ctx *iris.Context) {
	image, err := ctx.FormFile("image")
	//defer os.Remove(image.Filename)

	if err != nil{
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
	src ,err := image.Open()
	if err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
	defer src.Close()

	dst, err := os.Create(image.Filename)
	if err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}

	cf, ok := ctx.Get("cf").(*clarifai.Client)
	if !ok {
		panic(err)
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}


	fmt.Printf("%+v \n",cf)
	urls := []string{image.Filename}
	fmt.Printf("%+v \n",urls)
	resp, err := cf.Tag(clarifai.TagRequest{URLs: urls})
	if err != nil {
		ctx.JSON(500, err)
	}
	fmt.Printf("%+v \n", resp)
	ctx.JSON(200,resp)
}

