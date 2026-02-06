package repository

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"toychart/model"
	"toychart/utils"

	"gorm.io/gorm"
)

type ToyRepository interface {
	Create(Toy *model.Toy) error
	GetById(id string) (*model.Toy, error)
	GetByToyNameAndSet(ToyName, set string) bool
	SearchToyList(keyword, set, order string) ([]*model.Toy, error)
	GetAllToys() ([]*model.Toy, error)
	Update(Toy *model.Toy) error
	Delete(id string) error
}

type toyRepository struct {
	db *gorm.DB
}

func NewToyRepository(db *gorm.DB) ToyRepository {
	return &toyRepository{db: db}
}

func (r *toyRepository) Create(Toy *model.Toy) error {
	result := r.db.Create(Toy)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *toyRepository) GetById(id string) (*model.Toy, error) {
	var Toy model.Toy
	result := r.db.First(&Toy, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &Toy, nil
}

func (r *toyRepository) GetByToyNameAndSet(ToyName, set string) bool {
	var auth model.Toy
	result := r.db.
		Where("name = ? AND setName = ?", ToyName, set).
		First(&auth)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
	}
	return true
}

func (r *toyRepository) SearchToyList(keyword, set, order string) ([]*model.Toy, error) {
	var Toys []*model.Toy

	query := r.db

	if keyword != "" {
		if strings.Contains(keyword, "Pokemon") {
			keyword = utils.CapitalizeFirst(keyword)
			query = query.Where("setName LIKE ?", "%"+keyword+"%")
		} else {
			re := regexp.MustCompile(`\d+`)
			numbers := re.FindAllString(keyword, -1)

			namePart := keyword
			for _, num := range numbers {
				namePart = strings.ReplaceAll(namePart, num, "")
			}
			namePart = strings.TrimSpace(namePart)

			if namePart != "" {
				query = query.Where("name LIKE ?", "%"+namePart+"%")
			}

			for _, num := range numbers {
				query = query.Where("name REGEXP ?", fmt.Sprintf(`#\w*%s\b`, num))
			}
		}
	}

	if set != "" {
		query = query.Where("setName = ?", set)
	}

	if order != "" {
		switch order {
		case "no_asc":
			query = query.Where("name LIKE ?", "%#%").
				Order("CAST(SUBSTRING_INDEX(name, '#', -1) AS UNSIGNED) ASC")
		case "no_desc":
			query = query.Where("name LIKE ?", "%#%").
				Order("CAST(SUBSTRING_INDEX(name, '#', -1) AS UNSIGNED) DESC")
		case "price_asc":
			query = query.Order("CAST(REPLACE(ungrade, '$', '') AS DECIMAL(10,2)) ASC")
		case "price_desc":
			query = query.Order("CAST(REPLACE(ungrade, '$', '') AS DECIMAL(10,2)) DESC")
		case "trend":
		default:
			// "relevance"
		}

	}

	result := query.Find(&Toys)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return Toys, nil
}

func (r *toyRepository) GetAllToys() ([]*model.Toy, error) {
	var Toy []*model.Toy
	result := r.db.
		Order("createdDateTime ASC").
		Find(&Toy)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return Toy, nil
}

func (r *toyRepository) Update(Toy *model.Toy) error {
	result := r.db.Save(Toy)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *toyRepository) Delete(id string) error {
	result := r.db.Model(&model.Toy{}).Where("id = ?", id).Update("status", false)
	return result.Error
}
