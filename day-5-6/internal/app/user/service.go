package user

import (
	"context"

	"alterra-agmc-day-5-6/internal/dto"
	"alterra-agmc-day-5-6/internal/factory"
	"alterra-agmc-day-5-6/internal/model"
	"alterra-agmc-day-5-6/internal/repository"
	"alterra-agmc-day-5-6/pkg/constant"
	res "alterra-agmc-day-5-6/pkg/util/response"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[dto.UserResponse], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*dto.UserResponse, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(ctx context.Context, ID uint) (*model.User, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[dto.UserResponse], error) {

	users, info, err := s.UserRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var datas []dto.UserResponse

	for _, user := range users {

		datas = append(datas, dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})

	}

	result := new(dto.SearchGetResponse[dto.UserResponse])
	result.Datas = datas
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*dto.UserResponse, error) {
	var result *dto.UserResponse

	data, err := s.UserRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.UserResponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}

	return result, nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	var result *dto.UserResponse

	_, err := s.UserRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data = make(map[string]interface{})

	if payload.Name != nil {
		data["name"] = payload.Name
	}
	if payload.Email != nil {
		data["email"] = payload.Email
	}
	if payload.Password != nil {
		data["password"] = payload.Password
	}

	err = s.UserRepository.Update(ctx, ID, data)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = ID
	result.Name = *payload.Name
	result.Email = *payload.Email

	return result, nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.User, error) {

	data, err := s.UserRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.UserRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil

}
