package main

import (
	"fmt"
	"godesaapps/config"
	"godesaapps/controller"
	"godesaapps/repository"
	"godesaapps/service"
	"godesaapps/util"
	"net/http"
	"os"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("DesaApps Runn...")

	db, err := config.ConnectToDatabase()
	util.SentPanicIfError(err)

	userRepository := repository.NewUserRepositoryImpl(db)
	wargaRepository := repository.NewWargaRepository(db)
	contentRepository := repository.NewWebsiteContentRepository(db)

	userService := service.NewUserServiceImpl(userRepository, db)
	wargaService := service.NewWargaService(wargaRepository)
	contentService := service.NewWebsiteContentService(contentRepository)

	userController := controller.NewUserControllerImpl(userService)
	wargaController := controller.NewWargaController(wargaService)
	contentController := controller.NewWebsiteContentController(contentService)

	router := httprouter.New()

	// User route
	router.POST("/api/user/sign-up", userController.CreateUser)
	router.POST("/api/user/login", userController.LoginUser)
	router.GET("/api/user/me", VerifyJWT(userController.GetUserInfo))
	router.POST("/api/user/forgot-password", userController.ForgotPassword)
	router.POST("/api/user/reset-password", userController.ResetPassword)
	router.GET("/api/users", userController.GetAllUsers)
	router.DELETE("/api/deleteusers/:id", userController.DeleteUserHandler)

	// Warga route
	router.POST("/api/warga/register", wargaController.RegisterWarga)
	router.POST("/api/warga", wargaController.InsertDataWarga)
	router.GET("/api/warga", wargaController.GetAllWarga)
	router.PUT("/api/warga/:id", wargaController.UpdateWarga)
	router.DELETE("/api/warga/:id", wargaController.DeleteWarga)

	// Website content route
	router.GET("/api/content", contentController.GetContent)
	router.PUT("/api/content", contentController.UpdateContent)
	router.ServeFiles("/kontenwebsite/*filepath", http.Dir("./kontenwebsite"))

	//dashboard statss
	dashboardRepository := repository.NewDashboardRepository(db)
	dashboardService := service.NewDashboardService(dashboardRepository)
	dashboardController := controller.NewDashboardController(dashboardService)

	router.GET("/api/dashboard/stats", dashboardController.GetStats)


	handler := corsMiddleware(router)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.AppPort),
		Handler: handler,
	}

	errServer := server.ListenAndServe()
	util.SentPanicIfError(errServer)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGIN"), ",")

		for _, o := range allowedOrigins {
			if strings.TrimSpace(o) == origin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}


func VerifyJWT(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid Token Format", http.StatusUnauthorized)
			return
		}

		claims := &service.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or Expired Token", http.StatusUnauthorized)
			return
		}

		r.Header.Set("User-Email", claims.Email)
		next(w, r, ps)
	}
}
