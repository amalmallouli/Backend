package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/config"
	"github.com/ydhnwb/golang_api/controller"
	"github.com/ydhnwb/golang_api/middleware"
	"github.com/ydhnwb/golang_api/repository"
	"github.com/ydhnwb/golang_api/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	cvRepository repository.CVRepository = repository.NewCVRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	cvService    service.CVService       = service.NewCVService(cvRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	cvController controller.CVController = controller.NewCVController(cvService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	cvRoutes := r.Group("api/cvs", middleware.AuthorizeJWT(jwtService))
	{
		cvRoutes.GET("/", cvController.All)
		cvRoutes.POST("/", cvController.Insert)
		cvRoutes.GET("/:id", cvController.FindByID)
		cvRoutes.PUT("/:id", cvController.Update)
		cvRoutes.DELETE("/:id", cvController.Delete)
	}

	r.Run()
}
