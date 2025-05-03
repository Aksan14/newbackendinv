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


	adminRepo := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepo)
	adminController := controller.NewAdminController(adminService)
	router.ServeFiles("/pegawai/*filepath", http.Dir("./pegawai"))

	router.POST("/api/admin/create", adminController.CreateAdminFromPegawai)

	router.POST("/api/user/sign-up", userController.CreateUser)
	router.POST("/api/user/login", userController.LoginUser)
	router.GET("/api/user/me", VerifyJWT(userController.GetUserInfo))
	router.POST("/api/user/forgot-password", userController.ForgotPassword)
	router.POST("/api/user/reset-password", userController.ResetPassword)
	router.GET("/api/users", userController.GetAllUsers)
	router.DELETE("/api/deleteusers/:id", userController.DeleteUserHandler)

	router.POST("/api/warga/register", wargaController.RegisterWarga)
	router.POST("/api/warga", wargaController.InsertDataWarga)
	router.GET("/api/warga", wargaController.GetAllWarga)
	router.PUT("/api/warga/:id", wargaController.UpdateWarga)
	router.DELETE("/api/warga/:id", wargaController.DeleteWarga)

	router.GET("/api/content", contentController.GetContent)
	router.PUT("/api/content", contentController.UpdateContent)
	router.ServeFiles("/kontenwebsite/*filepath", http.Dir("./kontenwebsite"))

	dashboardRepository := repository.NewDashboardRepository(db)
	dashboardService := service.NewDashboardService(dashboardRepository)
	dashboardController := controller.NewDashboardController(dashboardService)

	router.GET("/api/dashboard/stats", dashboardController.GetStats)

	requestSuratRepo := repository.NewRequestSuratRepository(db)
	requestSuratService := service.NewRequestSuratService(requestSuratRepo)
	requestSuratController := controller.NewRequestSuratController(requestSuratService)

	router.GET("/api/request/warga/:nik", requestSuratController.FindWargaByNik)
	router.POST("/api/request/surat", requestSuratController.CreateRequestSurat)


	pegawaiRepo := repository.NewPegawaiRepository(db)
	pegawaiService := service.NewPegawaiService(pegawaiRepo)
	pegawaiController := controller.NewPegawaiController(pegawaiService)

	router.POST("/api/pegawai/create", pegawaiController.CreatePegawai)
	router.GET("/api/pegawai/getall", pegawaiController.GetAllPegawai)
	router.GET("/api/pegawai/getpegawaibyid/:id", pegawaiController.GetPegawaiByID)
	router.PUT("/api/pegawai/update/:id", pegawaiController.UpdatePegawai)
	router.DELETE("/api/pegawai/delete/:id", pegawaiController.DeletePegawai)
		

	handler := corsMiddleware(router)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.AppHost, config.AppPort),
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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, ngrok-skip-browser-warning")
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
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or Expired Token", http.StatusUnauthorized)
			return
		}

		r.Header.Set("User-Email", claims.Email)
		next(w, r, ps)
	}
}
