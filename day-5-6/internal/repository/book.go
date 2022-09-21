package repository

import (
	"alterra-agmc-day-5-6/internal/dto"
	"alterra-agmc-day-5-6/internal/model"
	"context"
	"strings"

	"gorm.io/gorm"
)

type Book interface {
	Create(ctx context.Context, data model.Book) error
	Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Book, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.Book, error)
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type book struct {
	Db *gorm.DB
}

func NewBook(db *gorm.DB) *book {
	return &book{
		db,
	}
}

func (p *book) Create(ctx context.Context, data model.Book) error {
	return p.Db.WithContext(ctx).Model(&model.Book{}).Create(&data).Error
}

func (p *book) Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Book, *dto.PaginationInfo, error) {
	var books []model.Book
	var count int64

	query := p.Db.WithContext(ctx).Model(&model.Book{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?  ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&books).Error

	return books, dto.CheckInfoPagination(paginate, count), err
}

func (p *book) FindByID(ctx context.Context, ID uint) (model.Book, error) {

	var data model.Book
	err := p.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	return data, err
}

func (p *book) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := p.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.Book{}).Updates(data).Error
	return err

}

func (p *book) Delete(ctx context.Context, ID uint) error {

	err := p.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.Book{}).Error
	return err
}
