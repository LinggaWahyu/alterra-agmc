package book

import (
	"alterra-agmc-day-5-6/internal/dto"
	"alterra-agmc-day-5-6/internal/factory"
	"alterra-agmc-day-5-6/internal/model"
	"alterra-agmc-day-5-6/internal/repository"
	"alterra-agmc-day-5-6/pkg/constant"

	res "alterra-agmc-day-5-6/pkg/util/response"
	"context"
)

type service struct {
	BookRepository repository.Book
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Book], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Book, error)
	Create(ctx context.Context, payload *dto.CreateBookRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateBookRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.Book, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		BookRepository: f.BookRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Book], error) {

	Books, info, err := s.BookRepository.Find(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.Book])
	result.Datas = Books
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Book, error) {

	data, err := s.BookRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateBookRequest) (string, error) {

	var book = model.Book{
		Name:        payload.Name,
		Stock:       payload.Stock,
		Description: payload.Description,
	}

	err := s.BookRepository.Create(ctx, book)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateBookRequest) (string, error) {

	var data = make(map[string]interface{})

	if payload.Name != nil {
		data["name"] = payload.Name
	}
	if payload.Stock != nil {
		data["stock"] = payload.Stock
	}
	if payload.Description != nil {
		data["description"] = payload.Description
	}

	err := s.BookRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.Book, error) {

	data, err := s.BookRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.BookRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil

}
