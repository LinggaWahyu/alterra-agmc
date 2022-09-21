package repository

import (
	"alterra-agmc-day-5-6/internal/dto"
	"alterra-agmc-day-5-6/internal/model"
	"context"
	"strings"

	"gorm.io/gorm"
)

type User interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, p *dto.Pagination) ([]model.User, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.User, error)
	FindByEmail(ctx context.Context, email *string) (*model.User, error)
	Create(ctx context.Context, data model.User) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type user struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{
		db,
	}
}

func (r *user) FindAll(ctx context.Context, payload *dto.SearchGetRequest, p *dto.Pagination) ([]model.User, *dto.PaginationInfo, error) {
	var users []model.User
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.User{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? or lower(email) Like ? ", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(p)

	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, dto.CheckInfoPagination(p, count), err
}

func (r *user) FindByID(ctx context.Context, id uint) (model.User, error) {
	var user model.User
	err := r.Db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *user) FindByEmail(ctx context.Context, email *string) (*model.User, error) {
	conn := r.Db.WithContext(ctx)

	var data model.User
	err := conn.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (p *user) Create(ctx context.Context, data model.User) error {
	return p.Db.WithContext(ctx).Model(&model.User{}).Create(&data).Error
}

func (r *user) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.User{}).Updates(data).Error
	return err
}

func (r *user) Delete(ctx context.Context, ID uint) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.User{}).Error
	return err
}
