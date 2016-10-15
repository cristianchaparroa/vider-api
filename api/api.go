package api

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/logger"
	"github.com/clarifai/clarifai-go"


	"github.com/cristianchaparroa/vider-api/midelwares"
)

type viderAPI struct{
	Iris *iris.Framework
}

func NewViderAPI(cf * clarifai.Client) (*viderAPI,error){
	api := iris.New()
	api.Use(midelwares.ClarifaiMiddleware(cf))

	loggerConfig := logger.Config{Method:true, Path:true}
	logger := logger.New(loggerConfig)
	api.Use(logger)

	return &viderAPI{Iris:api},nil
}

// Start, runs the api
func (a *viderAPI) Start(){
	a.Iris.Listen(":8090")
}

// Close the api
func (a *viderAPI) Close(){
	a.Iris.Close()
}

// Setup the endpoints
func (a *viderAPI) Setup() {

	a.Iris.Post("/api/tags",Tags)
}