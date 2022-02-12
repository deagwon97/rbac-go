package dblayer

import "rbac/content/models"

type DBLayer interface {
	GetAllContents(int, int, string) (models.ContentList, error)
	GetContent(int) (models.Content, error)
	AddContent(models.ContentData) (models.Content, error)
	UpdateContent(int, models.ContentData) (models.Content, error)
	DeleteContent(int) (models.Content, error)
}
