package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"pkl/backend/database"
	usersv1 "pkl/backend/gen/pkl/users/v1"
	"pkl/backend/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// UserServer mengimplementasikan pkl.users.v1.UserService.
type UserServer struct {
	usersv1.UnimplementedUserServiceServer
	db     *gorm.DB
	jwtCfg Config
}

// Login memvalidasi kredensial pengguna terhadap database.
func (s *UserServer) Login(ctx context.Context, req *usersv1.LoginRequest) (*usersv1.LoginResponse, error) {
	var user models.User
	if err := s.db.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error; err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "username atau password salah")
	}

	token, err := s.jwtCfg.Generate(user.ID, user.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "gagal membuat token: %v", err)
	}

	return &usersv1.LoginResponse{Token: token}, nil
}

// Register mendaftarkan pengguna baru dan menyimpannya ke database.
func (s *UserServer) Register(ctx context.Context, req *usersv1.RegisterRequest) (*usersv1.RegisterResponse, error) {
	var existing models.User
	if err := s.db.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "username '%s' sudah terdaftar", req.Username)
	}

	user := models.User{
		Nama:     req.Nama,
		Username: req.Username,
		Password: req.Password,
	}
	if err := s.db.Create(&user).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "gagal menyimpan pengguna: %v", err)
	}

	return &usersv1.RegisterResponse{
		Nama:     req.Nama,
		Username: req.Username,
	}, nil
}

// GetProfile mengambil data profil dari database (bukan dari token).
// Data user tetap diambil dari DB supaya potensi existing data tetap bisa
// dicentang/diubah (mis. nama, saldo), sementara keamanan dijamin JWT.
func (s *UserServer) GetProfile(ctx context.Context, req *usersv1.GetProfileRequest) (*usersv1.GetProfileResponse, error) {
	claims, ok := ctx.Value(claimsKey{}).(*Claims)
	if !ok || claims == nil {
		return nil, status.Errorf(codes.Unauthenticated, "token tidak valid")
	}

	var user models.User
	if err := s.db.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "user tidak ditemukan")
	}

	return &usersv1.GetProfileResponse{
		UserId:   uint32(user.ID),
		Nama:     user.Nama,
		Username: user.Username,
	}, nil
}

// loadDotEnv membaca file .env (jika ada) dan mengisi environment variable
// yang belum diset. Mencoba beberapa lokasi karena working directory bisa
// berbeda (project root atau cmd/server).
func loadDotEnv() {
	for _, path := range []string{".env", "../../.env", "../.env"} {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			line = strings.TrimPrefix(line, "export ")
			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				continue
			}
			key := strings.TrimSpace(parts[0])
			val := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
			if os.Getenv(key) == "" {
				_ = os.Setenv(key, val)
			}
		}
	}
}

func main() {
	loadDotEnv()

	dbFile := os.Getenv("DATABASE_FILE")
	if dbFile == "" {
		dbFile = "../../app.db"
	}

	db, err := database.Init(dbFile)
	if err != nil {
		log.Fatalf("gagal membuka database: %v", err)
	}

	jwtCfg := NewConfig()
	log.Printf("jwt: issuer=%s ttl=%s", jwtCfg.Issuer, jwtCfg.TTL)

	srv := &UserServer{db: db, jwtCfg: jwtCfg}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	// Registrasikan service gRPC menggunakan helper yang di-generate.
	usersv1.RegisterUserServiceServer(grpcServer, srv)

	reflection.Register(grpcServer)

	fmt.Println("Jalan")

	// Server HTTP/Connect (frontend @connectrpc/connect-web)
	go func() {
		if err := startConnectServer(srv, ":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("connect server: %v", err)
		}
	}()

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
