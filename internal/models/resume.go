package models

import "time"

type Resume struct {
	ID       int64
	UserID   int64
	JobTitle string
	ContactDetail
	SocialMedias
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
}

type ReferenceList struct {
	ID            int64
	UserID        int64
	ResumeID      int64
	ReferenceList []ReferenceItem
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Version       int
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

type AwardsList struct {
	ID        int64
	UserID    int64
	ResumeID  int64
	Awards    []AwardItem
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

type EducationList struct {
	ID        int64
	ResumeID  int64
	Education []EducationItem
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

type EducationItem struct {
	ID        int64
	UserID    int64
	ResumeID  int64
	Name      string
	Degree    string
	Address1  string
	Address2  string
	City      string
	State     string
	Zipcode   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

type EmploymentListItem struct {
	ID               int64
	UserID           int64
	PastEmploymentID int64
	Title            string
	From             time.Time
	To               time.Time
	JobTitle         string
	Summary          string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Version          int
}

type EmploymentList struct {
	ID        int64
	UserID    int64
	ResumeID  int64
	Employers []EmploymentListItem
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

type SkillList struct {
	ID        int64
	UserID    int64
	ResumeID  int64
	Title     string
	Items     []SkillListItem
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

type SkillListItem struct {
	ID          int64
	UserID      int64
	SKillListID int64
	Title       string
	Content     string
	Duration    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Version     int
}

type Objective struct {
	ID        int64
	UserID    int64
	ResumeID  int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

type SocialMedias struct {
	ID           int64
	UserID       int64
	ResumeID     int64
	SocialMedias []SocialMedia
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Version      int
}

type SocialMedia struct {
	ID              int64
	UserID          int64
	SocialMediaList int64
	CompanyName     string
	UserName        string
	WebAddress      string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Version         int
}

type ContactDetail struct {
	ID           int64
	UserID       int64
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
