package mockdbconn

import "github.com/elkcityhazard/andrew-mccall-go/internal/models"

type MockDBRepo struct {
}

func (mdbr *MockDBRepo) InsertUser(u *models.User) (int64, error) {
	return 1, nil
}
func (mdbr *MockDBRepo) GetUserByEmail(email string) (*models.User, error) {
	u := &models.User{}
	return u, nil
}
func (mdbr *MockDBRepo) GetUserByID(id int64) (*models.User, error) {
	return &models.User{}, nil

}
func (mdbr *MockDBRepo) InsertActivationToken(at *models.ActivationToken) (int64, error) {
	return 1, nil

}
func (mdbr *MockDBRepo) GetActivationToken(token string) (*models.User, *models.ActivationToken, error) {
	return &models.User{}, &models.ActivationToken{}, nil

}
func (mdbr *MockDBRepo) ActivateUser(*models.User) (int64, error) {
	return 1, nil

}

func (mdbr *MockDBRepo) InsertEditorContent(*models.Content) (int64, error) {
	return 1, nil

}
func (mdbr *MockDBRepo) GetBlogPost(postkey string) (*models.Content, error) {
	return &models.Content{}, nil

}
func (mdbr *MockDBRepo) GetBlogPostByID(id int64) (*models.Content, error) {
	return &models.Content{}, nil

}
func (mdbr *MockDBRepo) GetPaginatedPosts(userID int64, offset, size int) ([]*models.Content, error) {
	return []*models.Content{}, nil

}
func (mdbr *MockDBRepo) GetNextPrevPost(post *models.Content, increment bool) (*models.Content, error) {
	return &models.Content{}, nil
}
func (mdbr *MockDBRepo) ListPosts(limit, offset int) ([]*models.Content, error) {
	return []*models.Content{}, nil

}
func (mdbr *MockDBRepo) UpdatePost(*models.Content) (int64, error) { // returns version
	return 1, nil
}
func (mdbr *MockDBRepo) DeletePostById(id, userID int64) (int64, error) {
	return 1, nil

}

func (mdbr *MockDBRepo) GetTotalCount(table string) (int, error) {
	return 10, nil

}

// messages
func (mdbr *MockDBRepo) InsertMessage(*models.ContactMsg) (int64, error) {
	return 1, nil
}

// categories
func (mdbr *MockDBRepo) InsertCategory(*models.Category) (int64, error) {
	return 1, nil
}
func (mdbr *MockDBRepo) InsertCategoryPostJoin(*models.CategoryPostJoin) (int64, error) {
	return 1, nil
}
func (mdbr *MockDBRepo) ListCategories() ([]*models.Category, error) {
	var cat = []*models.Category{}
	return cat, nil
}
func (mdbr *MockDBRepo) GetCategoryByPostID(postID int64) (*models.Category, error) {
	cat := models.NewCategory()
	cat.ID = 1

	return cat, nil
}

// Resume
func (mdbr *MockDBRepo) InsertResume(rme *models.Resume) (int64, error) {
	return 1, nil
}
func (mdbr *MockDBRepo) GetResumeById(userID int64) (*models.Resume, error) {
	return &models.Resume{}, nil
}
func (mdbr *MockDBRepo) GetResumeObjective(resumeID int64) (*models.Objective, error) {
	return &models.Objective{}, nil
}
func (mdbr *MockDBRepo) GetResumeContactDetails(resumeID int64) (*models.ContactDetail, error) {
	return &models.ContactDetail{}, nil
}
func (mdbr *MockDBRepo) GetResumeSocialMedia(resumeID int64) (*models.SocialMediaList, error) {
	return &models.SocialMediaList{}, nil
}
func (mdbr *MockDBRepo) GetAwardItems(resumeID int64) (*models.AwardsList, error) {
	return &models.AwardsList{}, nil
}
func (mdbr *MockDBRepo) GetSkillItems(resumeID int64) (*models.SkillList, error) {
	return &models.SkillList{}, nil
}
func (mdbr *MockDBRepo) GetEmploymentList(resumeID int64) (*models.EmploymentList, error) {
	return &models.EmploymentList{}, nil
}
func (mdbr *MockDBRepo) GetEducationList(resumeID int64) (*models.EducationList, error) {
	return &models.EducationList{}, nil
}
func (mdbr *MockDBRepo) GetReferenceList(resumeID int64) (*models.ReferenceList, error) {
	return &models.ReferenceList{}, nil
}
