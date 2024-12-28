package templates

import (
	"embed"
	"fmt"
	"html"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/repository"
	"github.com/elkcityhazard/andrew-mccall-go/internal/repository/sqldbconn"
)

//go:embed templates
var templateDir embed.FS

var (
	templatePageDir    = "templates/pages"
	templateLayoutDir  = "templates/layouts"
	templatePartialDir = "templates/partials"
)

// templateRepo is a new sqldbconn to fetch the user and utilize any other dbservicers
var templateRepo repository.DBServicer

// SetTemplateSQLDbRepo gives the template package access to the sql db connection
func SetTemplateSQLDbRepo(sq *sqldbconn.SQLDbConn) {
	templateRepo = sq
}

var tFuncs = template.FuncMap{
	"parseEditorContent": parseQuillContent,
	"unsafeHTML":         unsafeHTML,
	"humanDate":          humanDate,
	"calculateLimit":     calculateLimit,
	"calculateOffset":    calculateOffset,
	"toLower":            strings.ToLower,
	"formatPluralYear":   formatPluralYear,
}

func formatPluralYear(year int) string {

	if year < 0 {
		year = -year
	}

	if year == 1 {
		return fmt.Sprintf("%d year", year)
	}
	return fmt.Sprintf("%d years", year)
}

func calculateLimit(limit, offset, count int, increment bool) int {
	switch true {
	case limit < 0:
		return 10
	case limit >= count:
		return 10
	default:
		return limit
	}
}

func calculateOffset(limit, offset, count int, increment bool) int {

	var off int

	if increment {
		if offset+limit >= count {
			off = offset
		} else if offset+limit < 0 {
			off = 0
		} else {
			off = offset + limit
		}
	}

	if !increment {
		if offset-limit < 0 {
			off = 0
		} else if offset-limit >= count {
			off = offset
		} else {
			off = offset - limit
		}
	}

	return off
}

func humanDate(datetime time.Time) string {
	return datetime.Format("Jan 02, 2006")
}

func unsafeHTML(content string) string {
	return html.EscapeString(content)
}

func parseQuillContent(content string) template.HTML {

	return template.HTML(content)

}

func GetTemplateDir() *embed.FS {
	return &templateDir
}

func BuildTemplateCache() map[string]*template.Template {

	var tc = make(map[string]*template.Template)

	pages, err := templateDir.ReadDir(templatePageDir)

	if err != nil {
		log.Panic(err)
	}

	for _, page := range pages {
		tmpl, err := template.New(page.Name()).Funcs(tFuncs).ParseFS(templateDir, fmt.Sprintf("%s/%s", templatePageDir, page.Name()))

		if err != nil {
			log.Panic(err)
		}

		layouts, err := templateDir.ReadDir(templateLayoutDir)

		if err != nil {
			log.Panic(err)
		}

		if len(layouts) > 0 {
			for _, l := range layouts {
				tmpl, err = tmpl.ParseFS(templateDir, fmt.Sprintf("%s/%s", templateLayoutDir, l.Name()))

				if err != nil {
					log.Panic(err)
				}
			}
		}

		partials, err := templateDir.ReadDir(templatePartialDir)

		if err != nil {
			log.Panic(err)
		}

		if len(partials) > 0 {
			for _, p := range partials {
				tmpl, err = tmpl.ParseFS(templateDir, fmt.Sprintf("%s/%s", templatePartialDir, p.Name()))

				if err != nil {
					log.Panic(err)
				}
			}
		}
		tc[page.Name()] = tmpl
	}
	return tc
}
