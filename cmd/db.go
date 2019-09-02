package main

import (
	"github.com/matvoy/dates/models"
)

func (a *App) initDB() {
	a.db.AutoMigrate(
		&models.Date{},
		&models.Birth{},
		&models.Death{},
		&models.Event{},
		&models.DeathLink{},
		&models.BirthLink{},
		&models.EventLink{},
	)
	a.db.Model(&models.Birth{}).AddForeignKey("date_id", "dates(id)", "CASCADE", "CASCADE")
	a.db.Model(&models.Death{}).AddForeignKey("date_id", "dates(id)", "CASCADE", "CASCADE")
	a.db.Model(&models.Event{}).AddForeignKey("date_id", "dates(id)", "CASCADE", "CASCADE")
	a.db.Model(&models.BirthLink{}).AddForeignKey("birth_id", "births(id)", "CASCADE", "CASCADE")
	a.db.Model(&models.DeathLink{}).AddForeignKey("death_id", "deaths(id)", "CASCADE", "CASCADE")
	a.db.Model(&models.EventLink{}).AddForeignKey("event_id", "events(id)", "CASCADE", "CASCADE")
}

func (a *App) processDate(date *models.Date) error {
	if err := a.db.Delete(models.Date{}, "date_string=?", date.DateString).Error; err != nil {
		return err
	}
	if err := a.db.Create(date).Error; err != nil {
		return err
	}
	if err := a.db.Model(&models.Birth{}).Where("date_id = ?", date.ID).Count(&date.Count.Births).Error; err != nil {
		return err
	}
	if err := a.db.Model(&models.Death{}).Where("date_id = ?", date.ID).Count(&date.Count.Deaths).Error; err != nil {
		return err
	}
	if err := a.db.Model(&models.Event{}).Where("date_id = ?", date.ID).Count(&date.Count.Events).Error; err != nil {
		return err
	}
	return nil
}
