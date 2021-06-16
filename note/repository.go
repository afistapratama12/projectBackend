package note

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Note, error)
	FindAllByUser(userID string) ([]Note, error)
	FIndByID(ID string) (Note, error)
	Create(note Note) (Note, error)
	Update(ID string, dataUpdate map[string]interface{}) (Note, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Note, error) {
	var notes []Note

	if err := r.db.Where("deleted = ?", "0").Find(&notes).Error; err != nil {
		return notes, err
	}

	return notes, nil

}
func (r *repository) FindAllByUser(userID string) ([]Note, error) {
	var notes []Note

	if err := r.db.Where("user_id = ? AND deleted = ?", userID, "0").Find(&notes).Error; err != nil {
		return notes, err
	}

	return notes, nil
}
func (r *repository) FIndByID(ID string) (Note, error) {
	var note Note

	if err := r.db.Where("id = ?", ID).Find(&note).Error; err != nil {
		return note, err
	}

	return note, nil
}
func (r *repository) Create(note Note) (Note, error) {
	if err := r.db.Create(&note).Error; err != nil {
		return note, err
	}

	return note, nil

}
func (r *repository) Update(ID string, dataUpdate map[string]interface{}) (Note, error) {
	var note Note

	if err := r.db.Model(&note).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return note, err
	}

	if err := r.db.Where("id = ?", ID).Find(&note).Error; err != nil {
		return note, err
	}

	return note, nil
}
