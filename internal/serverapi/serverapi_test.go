package serverapi

import (
	"context"
	"net"
	"testing"

	shortenerv1 "github.com/dkshi/shortener/gen"
	"github.com/dkshi/shortener/internal/service"
	mock_service "github.com/dkshi/shortener/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

const buffSize = 1024 * 1024

func TestServerAPI_CreateShortURL(t *testing.T) {
	type mockBehavior func(s *mock_service.MockURL, originalURL string)

	testTable := []struct {
		name             string
		inputOriginalURL string
		mockBehavior     mockBehavior
		expectedShortURL string
		expectedError    error
	}{
		{
			name:             "OK",
			inputOriginalURL: "http://vk.com/dkshii",
			mockBehavior: func(s *mock_service.MockURL, originalURL string) {
				s.EXPECT().CreateShortURL(originalURL).Return("abcDeFghki", nil)
			},
			expectedShortURL: "abcDeFghki",
			expectedError:    nil,
		},
		{
			name:             "Service returned error",
			inputOriginalURL: "http://vk.com/dkshii",
			mockBehavior: func(s *mock_service.MockURL, originalURL string) {
				s.EXPECT().CreateShortURL(originalURL).Return("", status.Errorf(codes.Internal, "impossible error"))
			},
			expectedShortURL: "",
			expectedError:    status.Errorf(codes.Internal, "impossible error"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// Creating mock service
			urlMockService := mock_service.NewMockURL(c)
			testCase.mockBehavior(urlMockService, testCase.inputOriginalURL)
			service := &service.Service{URL: urlMockService}

			// Creating and starting to listen ServerAPI
			lis := bufconn.Listen(buffSize)
			grpcServer := grpc.NewServer()
			defer grpcServer.Stop()

			serverAPI := NewServerAPI(service)
			serverAPI.RegisterServer(grpcServer)

			go func() {
				if err := grpcServer.Serve(lis); err != nil {
					logrus.Fatalf("error serving server: %s", err.Error())
				}
			}()

			// Creating client
			ctx := context.Background()
			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
					return lis.Dial()
				}))
			if err != nil {
				logrus.Fatalf("error dial bufnet: %s", err.Error())
			}
			defer conn.Close()
			client := shortenerv1.NewShortenerClient(conn)

			// Getting response from server
			resp, err := client.CreateShortURL(ctx, &shortenerv1.OriginalURL{
				Url: testCase.inputOriginalURL,
			})

			// Checking if response is nil
			var respURL string
			if resp != nil {
				respURL = resp.Url
			}

			// Assert
			assert.Equal(t, respURL, testCase.expectedShortURL)
			assert.Equal(t, err, testCase.expectedError)
		})
	}
}

func TestServerAPI_GetOriginalURL(t *testing.T) {
	type mockBehavior func(s *mock_service.MockURL, shortURL string)

	testTable := []struct {
		name                string
		inputShortURL       string
		mockBehavior        mockBehavior
		expectedOriginalURL string
		expectedError       error
	}{
		{
			name:          "OK",
			inputShortURL: "abcDeFghki",
			mockBehavior: func(s *mock_service.MockURL, shortURL string) {
				s.EXPECT().GetOriginalURL(shortURL).Return("http://vk.com/dkshii", nil)
			},
			expectedOriginalURL: "http://vk.com/dkshii",
			expectedError:       nil,
		},
		{
			name:          "Service returned error",
			inputShortURL: "abcDeFghki",
			mockBehavior: func(s *mock_service.MockURL, originalURL string) {
				s.EXPECT().GetOriginalURL(originalURL).Return("", status.Errorf(codes.Internal, "impossible error"))
			},
			expectedOriginalURL: "",
			expectedError:       status.Errorf(codes.Internal, "impossible error"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// Creating mock service
			urlMockService := mock_service.NewMockURL(c)
			testCase.mockBehavior(urlMockService, testCase.inputShortURL)
			service := &service.Service{URL: urlMockService}

			// Creating and starting to listen ServerAPI
			lis := bufconn.Listen(buffSize)
			grpcServer := grpc.NewServer()
			defer grpcServer.Stop()

			serverAPI := NewServerAPI(service)
			serverAPI.RegisterServer(grpcServer)

			go func() {
				if err := grpcServer.Serve(lis); err != nil {
					logrus.Fatalf("error serving server: %s", err.Error())
				}
			}()

			// Creating client
			ctx := context.Background()
			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
					return lis.Dial()
				}))
			if err != nil {
				logrus.Fatalf("error dial bufnet: %s", err.Error())
			}
			defer conn.Close()
			client := shortenerv1.NewShortenerClient(conn)

			// Getting response from server
			resp, err := client.GetOriginalURL(ctx, &shortenerv1.ShortURL{
				Url: testCase.inputShortURL,
			})

			// Checking if response is nil
			var respURL string
			if resp != nil {
				respURL = resp.Url
			}

			// Assert
			assert.Equal(t, respURL, testCase.expectedOriginalURL)
			assert.Equal(t, err, testCase.expectedError)
		})
	}
}
