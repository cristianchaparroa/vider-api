package main

import (
	"github.com/joho/godotenv"
	"github.com/clarifai/clarifai-go"
	"os"
	"github.com/cristianchaparroa/vider-api/api"
)

func main(){
	// loading the clarifai credentials
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	cf := clarifai.NewClient(clientID,clientSecret)

	// check if exist some error with clarifai client
	_, err =  cf.Info()
	if err != nil {
		panic(err)
	}
	api ,err := api.NewViderAPI(cf)

	if err !=nil{
		panic(err)
	}
	api.Setup()
	api.Start()
	defer api.Close()
}
