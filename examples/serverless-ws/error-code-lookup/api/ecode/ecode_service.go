package ecode

import (
	"errors"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Ecode Service
type EcodeServicer interface {
	GetAllEcodes() ([]Response, error)
	GetEcodeById(id string) (*Response, error)
	Update(ecode *Ecode) (*Response, error)
	Insert(ecode *Ecode) (*Response, error)
	DeleteById(id string) (*Response, error)
}

type EcodeService struct {
	repository EcodeRepositorier
}

func NewEcodeService(repository EcodeRepositorier) *EcodeService {
	return &EcodeService{repository: repository}
}

func (service *EcodeService) GetAllEcodes() ([]Response, error) {
	log.Debug().Msg("retrieve all ecodes")
	ecodes, err := service.repository.FindAll()
	if err != nil {
		log.Error().Err(err).Msg("fails to retrieve all error codes")
		return nil, err
	}

	return ecodes[0].ToEcodeResponses(ecodes), nil
}

func (service *EcodeService) GetEcodeById(id string) (*Response, error) {
	log.Debug().Str("id", id).Msg("retrieve error code by id")
	ecode, err := service.repository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Str("id", id).Msg("record not found")
		} else {
			log.Error().Err(err).Str("id", id).Msg("fails to retrieve ecode by id")
		}
		return nil, err
	}

	return ecode.ToEcodeResponse(ecode), nil
}

func (service *EcodeService) Update(ecode *Ecode) (*Response, error) {
	log.Debug().Any("ecode", ecode).Msg("update ecode")
	updatedEcode, err := service.repository.Update(ecode)

	if err != nil {
		log.Error().Err(err).Msg("update ecode failed")
		return nil, err
	}

	return ecode.ToEcodeResponse(updatedEcode), err
}

func (service *EcodeService) Insert(ecode *Ecode) (*Response, error) {
	log.Debug().Any("ecode", ecode).Msg("insert ecode")
	updatedEcode, err := service.repository.Insert(ecode)

	if err != nil {
		log.Error().Err(err).Msg("insert ecode failed")
		return nil, err
	}

	return ecode.ToEcodeResponse(updatedEcode), err
}

func (service *EcodeService) DeleteById(id string) (*Response, error) {
	log.Debug().Str("id", id).Msg("delete ecode id")

	_, err := service.repository.DeleteById(id)

	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("delete ecode by id failed")
		return nil, err
	}

	return nil, err
}
