package model

import "time"

// -------------------- chapter_assignments --------------------

type ChapterAssignment struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement"`

	ChapterID int64 `gorm:"column:chapter_id;not null"`
	UserID    int64 `gorm:"column:user_id;not null"`
	Role      int   `gorm:"column:role;not null"`

	AssignedBy *int64 `gorm:"column:assigned_by"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (ChapterAssignment) TableName() string { return "chapter_assignments" }

// -------------------- chapter_work_images --------------------

type ChapterWorkImage struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement"`

	ChapterID int64  `gorm:"column:chapter_id;not null"`
	Role      int16  `gorm:"column:role;not null"`
	ImageURL  string `gorm:"column:image_url;type:text;not null"`

	OrderIndex int `gorm:"column:order_index;not null"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (ChapterWorkImage) TableName() string { return "chapter_work_images" }

// -------------------- chapter_work_logs --------------------

type ChapterWorkLog struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement"`

	ChapterID  int64   `gorm:"column:chapter_id;not null"`
	OrderIndex int     `gorm:"column:order_index;not null"`
	Role       *int16  `gorm:"column:role"`
	ActionType string  `gorm:"column:action_type;type:varchar(16);not null"`
	UserID     *int64  `gorm:"column:user_id"`
	Note       *string `gorm:"column:note;type:text"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (ChapterWorkLog) TableName() string { return "chapter_work_logs" }

// -------------------- chapter_works --------------------

type ChapterWork struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement"`

	ChapterID int64 `gorm:"column:chapter_id;not null"`

	TranslationText *string `gorm:"column:translation_text;type:text"`
	ProofreadText   *string `gorm:"column:proofread_text;type:text"`

	Stage int16 `gorm:"column:stage;not null;default:1"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (ChapterWork) TableName() string { return "chapter_works" }

// -------------------- messages --------------------

type Message struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement"`

	UserID   int64  `gorm:"column:user_id;not null"`
	SenderID *int64 `gorm:"column:sender_id"`

	Type int16 `gorm:"column:type;not null"`

	Title   *string `gorm:"column:title;type:varchar(128)"`
	Content *string `gorm:"column:content;type:text"`

	RefType *string `gorm:"column:ref_type;type:varchar(32)"`
	RefID   *int64  `gorm:"column:ref_id"`

	IsRead bool `gorm:"column:is_read;not null;default:false"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Message) TableName() string { return "messages" }

// -------------------- project_assignments --------------------

type ProjectAssignment struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement"`

	ProjectID int64 `gorm:"column:project_id;not null"`
	UserID    int64 `gorm:"column:user_id;not null"`
	Role      int   `gorm:"column:role;not null"`

	AssignedBy *int64 `gorm:"column:assigned_by"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (ProjectAssignment) TableName() string { return "project_assignments" }

// -------------------- project_chapters --------------------

type ProjectChapter struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement"`

	ProjectID   int64  `gorm:"column:project_id;not null"`
	ChapterName string `gorm:"column:chapter_name;type:varchar(128);not null"`
	IsVisible   bool   `gorm:"column:is_visible;not null;default:false"`
	OrderIndex  int    `gorm:"column:order_index;not null;default:1"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (ProjectChapter) TableName() string { return "project_chapters" }

// -------------------- project_tags --------------------

type ProjectTag struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement"`

	ProjectID int64 `gorm:"column:project_id;not null"`
	TagID     int64 `gorm:"column:tag_id;not null"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (ProjectTag) TableName() string { return "project_tags" }

// -------------------- projects --------------------

type Project struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement"`

	CoverURL *string `gorm:"column:cover_url;type:varchar(255)"`

	Name           string  `gorm:"column:name;type:varchar(128);not null"`
	TranslatedName *string `gorm:"column:translated_name;type:varchar(128)"`
	Author         *string `gorm:"column:author;type:varchar(128)"`
	SourceSite     *string `gorm:"column:source_site;type:varchar(128)"`

	Description            *string `gorm:"column:description;type:text"`
	TranslationDescription *string `gorm:"column:translation_description;type:text"`

	Status int `gorm:"column:status;not null;default:1"`

	SourceMax      int `gorm:"column:source_max;not null;default:1"`
	TranslatorMax  int `gorm:"column:translator_max;not null;default:1"`
	ProofreaderMax int `gorm:"column:proofreader_max;not null;default:1"`
	TypesetterMax  int `gorm:"column:typesetter_max;not null;default:1"`
	SupervisorMax  int `gorm:"column:supervisor_max;not null;default:1"`
	AdminMax       int `gorm:"column:admin_max;not null;default:1"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Project) TableName() string { return "projects" }

// -------------------- tags --------------------

type Tag struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement"`

	Name  string `gorm:"column:name;type:varchar(64);not null;unique"`
	Color string `gorm:"column:color;type:varchar(7);not null;default:#FFFFFF"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Tag) TableName() string { return "tags" }

// -------------------- users --------------------

type User struct {
	ID        int64   `gorm:"column:id;primaryKey;autoIncrement"`
	AvatarURL *string `gorm:"column:avatar_url;type:varchar(255)"`
	Nickname  string  `gorm:"column:nickname;type:varchar(64);not null"`

	Role     int     `gorm:"column:role;not null;default:1"`
	QQOpenID *string `gorm:"column:qq_openid;type:varchar(64);unique"`
	Status   int     `gorm:"column:status;not null;default:1"`

	Username     *string `gorm:"column:username;type:varchar(64);unique"`
	PasswordHash *string `gorm:"column:password_hash;type:text"`

	RegisteredAt time.Time  `gorm:"column:registered_at;autoCreateTime"`
	LastLoginAt  *time.Time `gorm:"column:last_login_at"`
}

func (User) TableName() string { return "users" }
