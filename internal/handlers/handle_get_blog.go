package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetBlog(w http.ResponseWriter, r *http.Request) {

	limitParam, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limitParam = 10
	}
	offsetParam, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offsetParam = 0
	}

	posts, err := hr.conn.ListPosts(limitParam, offsetParam)

	if err != nil {
		returnErr(w, err)
		return
	}

	count, err := hr.conn.GetTotalCount("posts")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		returnErr(w, err)
		return
	}
	intMap := make(map[string]int)

	intMap["Limit"] = limitParam
	intMap["Count"] = count
	intMap["Offset"] = offsetParam

	data := make(map[string]interface{})
	data["Posts"] = posts
	data["Count"] = count

	stringMap := make(map[string]string)

	stringMap["PageTitle"] = "Blog"

	if len(posts) < 1 {

		if offsetParam-limitParam < 0 {
			http.Redirect(w, r.WithContext(r.Context()), fmt.Sprintf("/blog?limit=%d&offset=%d", limitParam, 0), http.StatusSeeOther)
		} else {
			http.Redirect(w, r.WithContext(r.Context()), fmt.Sprintf("/blog?limit=%d&offset=%d", limitParam, offsetParam-limitParam), http.StatusSeeOther)
		}
		return
	}

	render.RenderTemplate(w, r, "blog.gohtml", &models.TemplateData{
		StringMap: stringMap,
		IntMap:    intMap,
		Data:      data,
		Form:      forms.New(nil),
	})

}
