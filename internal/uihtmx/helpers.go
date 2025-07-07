package uihtmx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/swierq/golang/internal/uihtmx/ui/layout"
)

func renderTempl(w http.ResponseWriter, component templ.Component, menu layout.Menu, tile string, description string, fullPage bool) error {
	if component == nil {
		return fmt.Errorf("component is nil")
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	page := layout.WithBase(component, menu, tile, description, fullPage)
	err := page.Render(context.Background(), w)
	if err != nil {
		return fmt.Errorf("error rendering component: %v", err)
	}
	return nil
}

func RenderPage(w http.ResponseWriter, component templ.Component, menu layout.Menu, tile string, description string) error {
	return renderTempl(w, component, menu, tile, description, true)
}

func RenderPartial(w http.ResponseWriter, component templ.Component, tile string, description string) error {
	menu := layout.Menu{Items: []layout.MenuItem{}}
	return renderTempl(w, component, menu, tile, description, false)
}
