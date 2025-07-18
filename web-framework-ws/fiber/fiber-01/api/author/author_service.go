package author

import (
	"github.com/rs/zerolog/log"
)

type Servicer interface {
	getAuthorByID(int) (*Author, error)
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (service *Service) getAuthorByID(ID int) (*Author, error) {
	author := &Author{
		ID:      ID,
		Name:    "Lucas Boyce",
		Bio:     "Lucas Boyce is a software engineer with a passion for building scalable applications.",
		Website: "https://lucasboyce.dev",
	}

	log.Debug().Any("author", author).Msg("getAuthorByID")
	return author, nil
}
