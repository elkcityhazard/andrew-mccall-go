package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetEditCompose(w http.ResponseWriter, r *http.Request) {
	paramKey, err := strconv.ParseInt(GetField(r, 0), 10, 32)

	if err != nil {
		panic(err)
	}

	post, err := hr.conn.GetBlogPostByID(paramKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var f = forms.New(url.Values{})

	f.Set("title", post.Title)
	f.Set("slug", post.Slug)
	f.Set("description", post.Description)
	f.Set("content", post.Content)
	stringMap := make(map[string]string)

	stringMap["Title"] = fmt.Sprintf("Edit: %s", post.Title)
	stringMap["Method"] = "PUT"
	stringMap["Action"] = fmt.Sprintf("/admin/compose/edit/%d", post.ID)

	var data = map[string]interface{}{}

	data["EditorContent"] = post.Delta

	var intMap = map[string]int{}

	intMap["PostID"] = int(post.ID)

	render.RenderTemplate(w, r, "compose.gohtml", &models.TemplateData{
		StringMap: stringMap,
		IntMap:    intMap,
		Data:      data,
		Form:      f,
	})

}
