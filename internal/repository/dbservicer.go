package repository

import "github.com/elkcityhazard/andrew-mccall-go/internal/models"

type DBServicer interface {
	InsertUser(u *models.User) (int64, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int64) (*models.User, error)
	InsertActivationToken(at *models.ActivationToken) (int64, error)
	GetActivationToken(token string) (*models.User, *models.ActivationToken, error)
	ActivateUser(*models.User) (int64, error)

	InsertEditorContent(*models.Content) (int64, error)
	GetBlogPost(postkey string) (*models.Content, error)
	GetBlogPostByID(id int64) (*models.Content, error)
	GetPaginatedPosts(userID int64, offset, size int) ([]*models.Content, error)
	GetNextPrevPost(post *models.Content, increment bool) (*models.Content, error)
	ListPosts(limit, offset int) ([]*models.Content, error)
	UpdatePost(*models.Content) (int64, error) // returns version
	DeletePostById(id, userID int64) (int64, error)

	GetTotalCount(table string) (int, error)

	//messages
	InsertMessage(*models.ContactMsg) (int64, error)

	//categories
	InsertCategory(*models.Category) (int64, error)
	InsertCategoryPostJoin(*models.CategoryPostJoin) (int64, error)
	ListCategories() ([]*models.Category, error)
	GetCategoryByPostID(postID int64) (*models.Category, error)

	//	Resume
	InsertResume(rme *models.Resume) (int64, error)
}
