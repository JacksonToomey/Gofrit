package app

import "github.com/jacksontoomey/gofrit/middleware"

// App is the main web application
type App struct {
}

// NewApp creates a pointer to a new application
func NewApp() *App {
	return new(App)
}

// RegisterMiddleWare adds the middleware to the app
func (app *App) RegisterMiddleWare(m *middleware.Middleware) {}
