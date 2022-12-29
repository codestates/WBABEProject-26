package main

import (
	"context"
	"log"
	"strconv"

	"wemade_project/docs"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	oosConfig "wemade_project/config"
	menu_controller "wemade_project/controller/menu"
	order_list_controller "wemade_project/controller/order_list"
	rating_controller "wemade_project/controller/rating"
	user_controller "wemade_project/controller/user"
	oos_valid "wemade_project/controller/validators"
	menu_model "wemade_project/model/menu"
	order_list_model "wemade_project/model/order"
	rating_model "wemade_project/model/rating"
	user_model "wemade_project/model/user"
	menu_router "wemade_project/router/menu"
	order_list_router "wemade_project/router/order_list"
	rating_route "wemade_project/router/rating"
	menu_service "wemade_project/service/menu"
	order_list_service "wemade_project/service/order_list"
	rating_service "wemade_project/service/rating"
	user_service "wemade_project/service/user"

	user_route "wemade_project/router/user"
)

var (
	server *gin.Engine
	ctx context.Context
	mongoClient *mongo.Client

	config *oosConfig.Config

	//Menu
	menuCollection *mongo.Collection	
	menuModel menu_model.MenuCollection
	menuService menu_service.MenuService
	menuController menu_controller.MenuController
	menuRouter menu_router.MenuRoute

	//User
	userCollection *mongo.Collection	
	userModel user_model.UserCollection
	userService user_service.UserService
	userController user_controller.UserController
	userRouter user_route.UserRoute

	//OrderList
	orderListCollection *mongo.Collection
	orderListModel order_list_model.OrderListCollection
	orderListService order_list_service.OrderListService
	orderListController order_list_controller.OrderListController
	orderListRouter order_list_router.OrderListRoute

	//Rating
	ratingCollection *mongo.Collection
	ratingModel rating_model.RatingCollection
	ratingService rating_service.RatingService
	ratingController rating_controller.RatingController
	ratingRoute rating_route.RatingRoute
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
	orderListCollection = mongoDB.Collection("order_list")
	ratingCollection = mongoDB.Collection("rating")

	//Rating
	ratingModel = rating_model.InitWithSelf(ratingCollection, ctx)
	ratingService = rating_service.InitWithSelf(ratingModel)
	ratingController = rating_controller.InitWithSelf(ratingService)
	ratingRoute =rating_route.InitWithSelf(ratingController)

	//Menu
	menuModel = menu_model.InitWithSelf(menuCollection, ctx);
	menuService = menu_service.InitWithSelf(menuModel, ratingService);
	menuController = menu_controller.InitWithSelf(menuService)
	menuRouter = menu_router.InitWithSelf(menuController)

	//User
	userModel = user_model.InitWithSelf(userCollection, ctx)
	userService = user_service.InitWithSelf(userModel)
	userController = user_controller.InitWithSelf(userService)
	userRouter = user_route.InitWithSelf(userController)

	//OrderList
	orderListModel = order_list_model.InitWithSelf(orderListCollection, ctx)
	orderListService = order_list_service.InitWithSelf(orderListModel, menuService)
	orderListController = order_list_controller.InitWithSelf(orderListService)
	orderListRouter = order_list_router.InitWithSelf(orderListController)

	

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
	// router := server.Group("/api")
	// {
	// 	router.GET("/healthchecker", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Success"})
	// })
	// }

	//Menu Route Setup
	menuRouter.InitWithRoute(server)
	userRouter.InitWithRoute(server)
	orderListRouter.InitWithRoute(server)
	ratingRoute.InitWithRoute(server)

	//Swagger
	server.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Title = "Go Study Project : Delivery API"
	docs.SwaggerInfo.Description = "Delivery API Swagger "
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.InfoInstanceName = "111"
	
	
	log.Fatal(server.Run(":" + strconv.Itoa(config.Server.Port)))
}