package models

import "time"

// Resume holds all resume elements together
type Resume struct {
	ID              int64
	UserID          int64
	JobTitle        string           // parsed
	ContactDetail   *ContactDetail   // parsed
	SocialMediaList *SocialMediaList // parsed
	Objective       *Objective       // parsed
	SkillList       *SkillList       // parsed
	EmploymentList  *EmploymentList  // parsed
	EducationList   *EducationList   // parsed
	AwardsList      *AwardsList
	ReferenceList   *ReferenceList
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Version         int
}

// NewResume returns a resume  model with created_at. updated_at, and version pre-populated
func NewResume() *Resume {
	return &Resume{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type ReferenceItem struct {
	ID              int64
	ReferenceListID int64
	FirstName       string
	LastName        string
	Email           string
	PhoneNumber     string
	JobTitle        string
	Organization    string
	Type            string
	Address1        string
	Address2        string
	City            string
	State           string
	Zipcode         string
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Version         int
}

func NewReferenceItem() *ReferenceItem {
	return &ReferenceItem{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type ReferenceList struct {
	ID            int64
	ResumeID      int64
	ReferenceList []*ReferenceItem
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Version       int
}

func NewReferenceList() *ReferenceList {
	return &ReferenceList{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type AwardItem struct {
	ID               int64
	AwardListID      int64
	Title            string
	OrganizationName string
	Year             int
	Content          string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Version          int
}

func NewAwardItem() *AwardItem {
	return &AwardItem{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type AwardsList struct {
	ID        int64
	ResumeID  int64
	Awards    []*AwardItem
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

func NewAwardsList() *AwardsList {
	return &AwardsList{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type EducationList struct {
	ID        int64
	ResumeID  int64
	Education []*EducationItem
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

func NewEducationList() *EducationList {
	return &EducationList{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type EducationItem struct {
	ID              int64
	EducationListID int64
	Name            string
	DegreeYear      int
	Degree          string
	Address1        string
	Address2        string
	City            string
	State           string
	Zipcode         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Version         int
}

func NewEducationItem() *EducationItem {
	return &EducationItem{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type EmploymentListItem struct {
	ID               int64
	EmploymentListID int64
	Title            string
	From             int
	To               int
	JobTitle         string
	Summary          string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Version          int
}

func NewEmploymentListItem() *EmploymentListItem {
	return &EmploymentListItem{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type EmploymentList struct {
	ID        int64
	ResumeID  int64
	Employers []*EmploymentListItem
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

func NewEmploymentList() *EmploymentList {
	return &EmploymentList{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type SkillList struct {
	ID        int64
	ResumeID  int64
	Title     string
	Items     []*SkillListItem
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

func NewSkillList() *SkillList {
	return &SkillList{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type SkillListItem struct {
	ID          int64
	SKillListID int64
	Title       string
	Content     string
	Duration    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Version     int
}

func NewSkillListItem() *SkillListItem {
	return &SkillListItem{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type Objective struct {
	ID        int64
	ResumeID  int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

func NewObjective() *Objective {
	return &Objective{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type SocialMediaList struct {
	ID                   int64
	ResumeID             int64
	SocialMediaListItems []*SocialMediaListItems
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Version              int
}

func NewSocialMediaList() *SocialMediaList {
	return &SocialMediaList{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type SocialMediaListItems struct {
	ID                int64
	SocialMediaListID int64
	CompanyName       string
	UserName          string
	WebAddress        string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Version           int
}

func NewSocialMediaListItems() *SocialMediaListItems {
	return &SocialMediaListItems{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}

type ContactDetail struct {
	ID           int64
	ResumeID     int64
	Firstname    string
	Lastname     string
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Zipcode      string
	Email        string
	PhoneNumber  string
	WebAddress   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Version      int
}

func NewContactDetail() *ContactDetail {
	return &ContactDetail{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
}
