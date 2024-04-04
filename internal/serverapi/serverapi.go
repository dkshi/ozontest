package serverapi

import (
	"context"

	shortenerv1 "github.com/dkshi/shortener/gen"
	"github.com/dkshi/shortener/internal/service"
	"google.golang.org/grpc"
)

type ServerAPI struct {
	shortenerv1.UnimplementedShortenerServer
	service *service.Service
}

func NewServerAPI(service *service.Service) *ServerAPI {
	return &ServerAPI{service: service}
}

func (s *ServerAPI) CreateShortURL(ctx context.Context, in *shortenerv1.OriginalURL) (*shortenerv1.ShortURL, error) {
	url, err := s.service.CreateShortURL(in.Url)

	if err != nil {
		return &shortenerv1.ShortURL{}, err
	}

	shortURL := &shortenerv1.ShortURL{}
	shortURL.Url = url

	return shortURL, nil
}

func (s *ServerAPI) GetOriginalURL(ctx context.Context, in *shortenerv1.ShortURL) (*shortenerv1.OriginalURL, error) {
	url, err := s.service.GetOriginalURL(in.Url)

	if err != nil {
		return &shortenerv1.OriginalURL{}, err
	}

	originalURL := &shortenerv1.OriginalURL{}
	originalURL.Url = url

	return originalURL, nil
}

func (s *ServerAPI) RegisterServer(gRPCServer *grpc.Server) {
	shortenerv1.RegisterShortenerServer(gRPCServer, s)
}
