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
	PastEmployment
	EducationList
	AwardsList
	References
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

type Reference struct {
	ID           int64
	userID       int64
	ReferencesID int64
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	JobTitle     string
	Organization string
	Type         string
	Address1     string
	Address2     string
	City         string
	State        string
	Zipcode      string
	Description  string
}

type References struct {
	ID            int64
	UserID        int64
	ResumeID      int64
	ReferenceList []Reference
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Version       int
}

type Award struct {
	ID               int64
	UserID           int64
	ResumeID         int64
	Title            string
	OrganizationName string
	Year             time.Time
	Summary          string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Version          int
}

type AwardsList struct {
	ID        int64
	UserID    int64
	ResumeID  int64
	Awards    []Award
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

type EducationList struct {
	ID        int64
	UserID    int64
	ResumeID  int64
	Education []Education
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
}

type Education struct {
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

type Employer struct {
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

type PastEmployment struct {
	ID        int64
	UserID    int64
	ResumeID  int64
	Employers []Employer
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
	Content     string
	Duration    time.Duration
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
	SocialMedias []SocialMedia
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Version      int
}

type SocialMedia struct {
	ID          int64
	UserID      int64
	ResumeID    int64
	CompanyName string
	UserName    string
	WebAddress  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Version     int
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
