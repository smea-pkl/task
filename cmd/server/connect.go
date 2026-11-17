package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	usersv1 "pkl/backend/gen/pkl/users/v1"
	usersv1connect "pkl/backend/gen/pkl/users/v1/usersv1connect"

	"connectrpc.com/connect"
)

// connectServer mengimplementasikan usersv1connect.UserServiceHandler
// menggunakan method gRPC yang sudah ada di UserServer.
type connectServer struct {
	*UserServer
}

func (s *connectServer) Login(ctx context.Context, req *connect.Request[usersv1.LoginRequest]) (*connect.Response[usersv1.LoginResponse], error) {
	resp, err := s.UserServer.Login(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse[usersv1.LoginResponse](resp), nil
}

func (s *connectServer) Register(ctx context.Context, req *connect.Request[usersv1.RegisterRequest]) (*connect.Response[usersv1.RegisterResponse], error) {
	resp, err := s.UserServer.Register(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse[usersv1.RegisterResponse](resp), nil
}

// withCORS menambahkan header CORS agar frontend (Vite :5173) bisa mengakses server ini.
func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Connect-Protocol-Version, Connect-Timeout-Ms, Grpc-Timeout, X-Grpc-Web, X-User-Agent")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// authInterceptor memvalidasi token JWT untuk RPC selain Login dan Register.
func authInterceptor(cfg Config) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			switch req.Spec().Procedure {
			case usersv1connect.UserServiceLoginProcedure, usersv1connect.UserServiceRegisterProcedure:
				return next(ctx, req)
			}

			token := ExtractToken(req.Header().Get("Authorization"))
			if token == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("token tidak ada"))
			}
			if _, err := cfg.Parse(token); err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}
			return next(ctx, req)
		}
	}
}

// startConnectServer mendengarkan HTTP di addr dan melayani RPC pkl.users.v1.UserService.
// Frontend menggunakan @connectrpc/connect-web ke alamat ini.
func startConnectServer(srv *UserServer, addr string) error {
	path, handler := usersv1connect.NewUserServiceHandler(
		&connectServer{UserServer: srv},
		connect.WithInterceptors(authInterceptor(srv.jwtCfg)),
	)
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	httpSrv := &http.Server{
		Addr:    addr,
		Handler: withCORS(mux),
	}
	log.Printf("connect server listening on http://%s", addr)
	return httpSrv.ListenAndServe()
}
