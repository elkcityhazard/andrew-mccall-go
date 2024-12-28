package handlers

import (
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetPost(w http.ResponseWriter, r *http.Request) {

	// since this route starts with /blog we can extract the route key pretty easily to fetch by sluga

	routeKey := r.URL.Path[len("/blog"):]

	post, err := hr.conn.GetBlogPost(routeKey)

	if err != nil {
		hr.app.MsgChan <- err.Error()
		render.RenderTemplate(w, r, "404.gohtml", &models.TemplateData{})
		return
	}

	user, err := hr.conn.GetUserByID(post.UserId)

	if err != nil {
		returnErr(w, err)
		return
	}
	post.User = user

	cat, err := hr.conn.GetCategoryByPostID(post.ID)

	if err != nil {
		cat = &models.Category{}
	}

	post.Category = cat

	prevPost, err := hr.conn.GetNextPrevPost(post, false)

	if err != nil {
		prevPost = &models.Content{}
	}

	nextPost, err := hr.conn.GetNextPrevPost(post, true)

	if err != nil {
		nextPost = &models.Content{}

	}

	var stringMap = map[string]string{}

	stringMap["PageTitle"] = post.Title
	stringMap["PageDescription"] = post.Description

	var data = map[string]interface{}{}
	data["Post"] = post
	data["Category"] = cat
	data["PrevPost"] = prevPost
	data["NextPost"] = nextPost

	render.RenderTemplate(w, r, "single-post.gohtml", &models.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})

}
