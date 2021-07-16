package service

import "github.com/vmandic/gin-gonic-crash-course/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{videos: []entity.Video{
		{
			Title:       "Video1",
			Description: "Desc1",
			URL:         "www.web.com",
			Author: entity.Person{
				FirstName: "Jack",
				LastName:  "Jones",
				Age:       22,
				Email:     "mail@mail.com",
			},
		},
	}}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
