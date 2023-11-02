package datastruct

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	DocumentID   uint    `gorm:"uniqueIndex:document_id" json:"document_id,omitempty"`
	UserRefer    uint    `json:"user_refer,omitempty"`
	DocumentPath *string `gorm:"column:document_path" json:"document_path,omitempty"`
}
