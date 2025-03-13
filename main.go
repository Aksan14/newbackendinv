package main

import (
	"context"
	"fmt"
	"godesaapps/config"
	"godesaapps/controller"
	"godesaapps/repository"
	"godesaapps/service"
	"godesaapps/util"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("DesaApps Runn...")

	mysql, err := config.ConnectToDatabase()
	util.SentPanicIfError(err)

	userRepository := repository.NewUserRepositoryImpl(mysql)
	wargaRepository := repository.NewWargaRepository(mysql)

	userService := service.NewUserServiceImpl(userRepository, mysql)
	wargaService := service.NewWargaService(wargaRepository)

	userController := controller.NewUserControllerImpl(userService)
	wargaController := controller.NewWargaController(wargaService)

	router := httprouter.New()

	// user
	router.POST("/api/user/sign-up", userController.CreateUser)
	router.POST("/api/user/login", userController.LoginUser)
	router.GET("/api/user/me", VerifyJWT(userController.GetUserInfo))
	router.POST("/api/user/forgot-password", userController.ForgotPassword)
	router.POST("/api/user/reset-password", userController.ResetPassword)

	//halaman didalamnya bagi akses berdasarkan role
	router.GET("/api/user/dashboard-bendahara", VerifyRole(userController.DashboardBendahara, "ROLE001"))
	router.GET("/api/user/dashboard-sekretaris", VerifyRole(userController.DashboardSekretaris, "ROLE002"))
	

	// warga
	router.POST("/api/warga/register", wargaController.RegisterWarga)

	handler := corsMiddleware(router)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.AppPort),
		Handler: handler,
	}

	//serverr
	errServer := server.ListenAndServe()
	util.SentPanicIfError(errServer)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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

func VerifyRole(next httprouter.Handle, allowedRole string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims := &service.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid atau token expired", http.StatusUnauthorized)
			return
		}

		// Debug
		// fmt.Printf("Decoded Token Claims: %+v\n", claims)

		//cekkk
		if claims.RoleId != allowedRole {
			http.Error(w, "Forbidden: tidak dapat mengakses ini role tidak sesuai ", http.StatusForbidden)
			return
		}

		// Lanjutkan ke handler jika role sesuai
		ctx := context.WithValue(r.Context(), nikadminKey, claims.Nikadmin)
		next(w, r.WithContext(ctx), ps)
	}
}































type contextKey string

const nikadminKey contextKey = "nikadmin"