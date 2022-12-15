package domain

type Router interface {
	RegisterRoutes()
	CreateNewURL(str string) string
}
