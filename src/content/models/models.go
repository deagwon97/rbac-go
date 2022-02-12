package models

type Content struct {
	ID        int    `gorm:"column:content_id" json:"content_id"`
	Title     string `gorm:"column:title"      json:"title"`
	Summary   string `gorm:"column:summary"    json:"summary"`
	Content   string `gorm:"column:content"    json:"content"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
	User      int    `gorm:"column:user"       json:"user"`
}

func (Content) TableName() string {
	// gorm에서 호출하는 테이블 명  커스텀
	// 기본값 Content -> contents
	return "content_content"
}

type ContentData struct {
	ID        int    `gorm:"column:content_id" json:"-"`
	Title     string `gorm:"column:title"      json:"title"`
	Summary   string `gorm:"column:summary"    json:"summary"`
	Content   string `gorm:"column:content"    json:"content"`
	UpdatedAt string `gorm:"column:updated_at" json:"-"`
}

func (ContentData) TableName() string {
	return "content_content"
}

type ContentItem struct {
	ID      int    `gorm:"column:content_id" json:"content_id"`
	Title   string `gorm:"column:title"      json:"title"`
	Summary string `gorm:"column:summary"    json:"summary"`
}

func (ContentItem) TableName() string {
	return "content_content"
}

type ContentList struct {
	Count        int           `json:"count"`
	NextPage     string        `json:"next"`
	PreviousPage string        `json:"previous"`
	Results      []ContentItem `json:"results"`
}
