package service

const (
	SERVICE_NAME = "HelloService"
)

type HelloService interface {
	Hello(request string, reponse *string) error
}
