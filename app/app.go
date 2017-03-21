package app

type config struct{}

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) AddMiddleware() {}

func (a *App) RegisterView() {}

func (a *App) Run() {}
