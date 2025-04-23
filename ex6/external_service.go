package ex6

import (
	"io"
	"net/http"
	"strconv"
)

//go:generate mockgen -source=external_service.go -destination=mock.go
type ServiceInterface interface {
	GetData(id int) (string, error)
}

type Service struct {
	baseUrl string
}

func NewService(baseUrl string) *Service {
	return &Service{
		baseUrl: baseUrl,
	}
}

func (s *Service) GetData(id int) (string, error) {
	resp, err := http.Get(s.baseUrl + "?id=" + strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
