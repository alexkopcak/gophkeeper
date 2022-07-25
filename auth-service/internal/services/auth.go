package services

import (
	"context"

	"github.com/alexkopcak/gophkeeper/auth-service/internal/db"
	"github.com/alexkopcak/gophkeeper/auth-service/internal/models"
	"github.com/alexkopcak/gophkeeper/auth-service/internal/pb"
	"github.com/alexkopcak/gophkeeper/auth-service/internal/utils"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	Handler *db.Handler
	Jwt     *utils.JwtWraper
}

func (s *AuthServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if res := s.Handler.DB.Where(
		&models.User{Name: in.UserName}).First(&user); res.Error == nil {
		return &pb.RegisterResponse{
			Token:  "",
			UserID: 0,
			Error:  "user already exists",
		}, nil
	}

	var err error
	user.Name = in.UserName
	user.Password, err = utils.HashPassword(in.Password)

	if err != nil {
		return nil, err
	}

	tx := s.Handler.DB.Create(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	token, err := s.Jwt.GenerateToken(&user)

	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Token:  token,
		UserID: user.Id,
		Error:  "",
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if res := s.Handler.DB.Where(&models.User{Name: in.UserName}).First(&user); res.Error != nil {
		return &pb.LoginResponse{
			Token:  "",
			UserID: 0,
			Error:  "user not found",
		}, nil
	}

	match := utils.CheckPasswordHash(in.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Token:  "",
			UserID: 0,
			Error:  "user not found",
		}, nil
	}

	token, err := s.Jwt.GenerateToken(&user)

	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token:  token,
		UserID: user.Id,
		Error:  "",
	}, nil
}

func (s *AuthServer) Verify(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	claims, err := s.Jwt.ValidateToken(in.Token)

	if err != nil {
		return &pb.VerifyResponse{
			Error: "bad request",
		}, err
	}

	return &pb.VerifyResponse{
		UserID: claims.IdUser,
		Error:  "",
	}, nil

}
