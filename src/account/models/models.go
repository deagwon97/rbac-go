package models

/*
CREATE TABLE `users` (
	`id` bigint NOT NULL AUTO_INCREMENT,
	`password` varchar(128) NOT NULL,
	`login_id` varchar(32) NOT NULL,
	`name` varchar(32) NOT NULL,
	`email` varchar(128) DEFAULT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY `email` (`email`)
  ) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/

type User struct {
	ID       int    `gorm:"column:id"       json:"id"`
	LoginId  string `gorm:"column:login_id" json:"login_id"`
	Password string `gorm:"column:password" json:"password"`
	Name     string `gorm:"column:name"     json:"name"`
	Email    string `gorm:"column:email"    json:"email"`
}

func (User) TableName() string {
	return "users"
}
