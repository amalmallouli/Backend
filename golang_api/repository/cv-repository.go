package repository

import (
	"github.com/ydhnwb/golang_api/entity"
	"gorm.io/gorm"
)

//CVRepository is a ....
type CVRepository interface {
	InsertCV(b entity.CV) entity.CV
	UpdateCV(b entity.CV) entity.CV
	DeleteCV(b entity.CV)
	AllCV() []entity.CV
	FindCVByID(cvID uint64) entity.CV
}

type cvConnection struct {
	connection *gorm.DB
}

//NewCVRepository creates an instance CVRepository
func NewCVRepository(dbConn *gorm.DB) CVRepository {
	return &cvConnection{
		connection: dbConn,
	}
}

func (db *cvConnection) InsertCV(b entity.CV) entity.CV {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *cvConnection) UpdateCV(b entity.CV) entity.CV {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *cvConnection) DeleteCV(b entity.CV) {
	db.connection.Delete(&b)
}

func (db *cvConnection) FindCVByID(cvID uint64) entity.CV {
	var cv entity.CV
	db.connection.Preload("User").Find(&cv, cvID)
	return cv
}

func (db *cvConnection) AllCV() []entity.CV {
	var cvs []entity.CV
	db.connection.Preload("User").Find(&cvs)
	return cvs
}
