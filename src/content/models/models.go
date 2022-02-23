package models

// CREATE TABLE `content_content` (
// 	`content_id` int NOT NULL AUTO_INCREMENT,
// 	`title` varchar(500) NOT NULL,
// 	`summary` varchar(1000) NOT NULL,
// 	`content` longtext,
// 	`created_at` datetime(6) NOT NULL,
// 	`updated_at` datetime(6) NOT NULL,
// 	`user` bigint NOT NULL,
// 	PRIMARY KEY (`content_id`),
// 	KEY `content_content_user_fe2e0079_fk_users_id` (`user`),
// 	CONSTRAINT `content_content_user_fe2e0079_fk_users_id` FOREIGN KEY (`user`) REFERENCES `users` (`id`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8;

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
