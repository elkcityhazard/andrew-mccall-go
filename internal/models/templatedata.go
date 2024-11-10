package models

import "github.com/elkcityhazard/andrew-mccall-go/internal/forms"

type TemplateData struct {
	SiteTitle       string
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	BoolMap         map[string]bool
	Data            map[string]any
	Form            *forms.Form
	CSRFToken       string
	IsAuthenticated bool
	Flash           string
	Warning         string
	Error           string
}

type MenuItem struct {
	URL   string
	Title string
}

func (td *TemplateData) PopulateAdminMenu() {

	var menu []*MenuItem

	menu = append(menu, &MenuItem{
		URL:   "/admin",
		Title: "Admin Home",
	})

	menu = append(menu, &MenuItem{
		URL:   "/admin/compose",
		Title: "Compose New",
	})

	if td.Data == nil {
		td.Data = make(map[string]any)
	}

	td.Data["AdminMenu"] = menu

}
