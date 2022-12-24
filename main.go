package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	oosConfig "wemade_project/config"
	receipt_controller "wemade_project/controller/receipt"
	oos_valid "wemade_project/controller/validators"
	receipt_model "wemade_project/model/receipt"
	receipt_router "wemade_project/router/receipt"
	receipt_service "wemade_project/service/receipt"

	user_model "wemade_project/model/user"
)

var (
	server *gin.Engine
	ctx context.Context
	mongoClient *mongo.Client

	config *oosConfig.Config

	//Menu
	menuCollection *mongo.Collection	
	menuModel receipt_model.MenuCollection
	menuService receipt_service.MenuService
	menuController receipt_controller.MenuController
	menuRouter receipt_router.MenuRoute

	//User
	userCollection *mongo.Collection	
	userModel user_model.UserCollection
)

//init 함수
func init() {
	//Timeline 
	

	var configErr error //컨피그 처리용 로컬 변수
	config, configErr = oosConfig.GetConfig()
	if (configErr != nil) {
		panic(configErr)
	}
	
	ctx = context.TODO()
	
	//Connect MongoDB
	var mongoErr error //몽고 처리용 로컬 변수
	mongoConnect := options.Client().ApplyURI(config.DB.Host).SetAuth(options.Credential{AuthSource : config.DB.Name, Username: config.DB.User, Password: config.DB.Pw})
	mongoClient, mongoErr = mongo.Connect(ctx, mongoConnect)

	if mongoErr != nil {
		panic(mongoErr)
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	//Mongo DB connect
	// fmt.Println("MongoDB successfully connected...")

	// Collections
	mongoDB := mongoClient.Database("sso")
	menuCollection = mongoDB.Collection("menu")
	userCollection = mongoDB.Collection("users")

	//Menu
	menuModel = receipt_model.InitWithSelf(menuCollection, ctx);
	menuService = receipt_service.InitWithSelf(menuModel);
	menuController = receipt_controller.InitWithSelf(menuService)
	menuRouter = receipt_router.InitWithSelf(menuController)

	//User
	userModel = user_model.InitWithSelf(userCollection, ctx)

	//Add Validator
	oos_valid.RegValidator4MenuEvent()

	//Gin Server
	server = gin.Default()
	server.Use(gin.Logger())

	/*
		ginEngin := gin.Default()
	ginEngin.Use(gin.Logger())
	ginEngin.Use(gin.Recovery())
	ginEngin.Use(router.CORS())
	*/
}

//main
func main() {
	startGinServer();
}

//func startGinServer(config *oosConfig.Config ) {
func startGinServer( ) {
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{config.Origin}
	// corsConfig.AllowCredentials = true
	
	//라우터 등록 부분

	router := server.Group("/api")
	{
		router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Success"})
	})
	}
	

	//Menu Route Setup
	menuRouter.InitWithRoute(server)
	

	
	log.Fatal(server.Run(":" + strconv.Itoa(config.Server.Port)))
}