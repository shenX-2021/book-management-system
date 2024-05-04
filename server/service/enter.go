package service

type ServiceStruct struct {
	BookService *BookService
	UserService *UserService
}

var Service = ServiceStruct{
	BookService: new(BookService),
	UserService: new(UserService),
}
