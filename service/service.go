package service

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Service struct {
	Client *resty.Client
}

func NewService() Service {
	client := resty.New()
	fmt.Println("Start service")
	return Service{
		Client: client,
	}
}
