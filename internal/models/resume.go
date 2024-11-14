package models

import "time"

type Resume struct {
	ID       int64
	UserID   int64
	JobTitle string
	ContactDetail
	SocialMediaList
	Objective
	SkillList
	EmploymentList
	EducationList
	AwardsList
	ReferenceList
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

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
	ReferenceList []ReferenceItem
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
	Year             time.Time
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
	Awards    []AwardItem
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
	Education []EducationItem
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
	DegreeYear      time.Time
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
	From             time.Time
	To               time.Time
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
	Employers []EmploymentListItem
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
	Items     []SkillListItem
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
	Duration    time.Time
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
	SocialMediaListItems []SocialMediaListItems
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
