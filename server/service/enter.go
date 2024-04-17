package service

type ServiceStruct struct {
	BookService *BookService
}

var Service = ServiceStruct{
	BookService: new(BookService),
}
