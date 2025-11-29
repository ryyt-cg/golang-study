package ecode

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type EcodeRepositorier interface {
	FindAll() ([]Ecode, error)
	FindById(id string) (*Ecode, error)
	Update(ecode *Ecode) (*Ecode, error)
	Insert(ecode *Ecode) (*Ecode, error)
	DeleteById(id string) (string, error)
}

// EcodeRepository searches error code in the database
type EcodeRepository struct {
	pg *gorm.DB
}

func NewEcodeRepository(pg *gorm.DB) *EcodeRepository {
	return &EcodeRepository{
		pg: pg,
	}
}

func (repository *EcodeRepository) FindAll() ([]Ecode, error) {
	log.Info().Msg("retrieve all error codes")
	var ecodes []Ecode
	err := repository.pg.Find(&ecodes).Error
	if err != nil {
		log.Error().Err(err).Msg("fails to retrieve all error codes")
		return nil, err
	}

	return ecodes, nil
}

// SELECT * FROM users WHERE id = 23;
// Get by primary key if it were a non-integer type
// db.First(&user, "id = ?", "string_primary_key")
// SELECT * FROM users WHERE id = 'string_primary_key';
func (repository *EcodeRepository) FindById(id string) (*Ecode, error) {
	log.Debug().Str("id", id).Msg("search error code by id")

	var ecode Ecode
	err := repository.pg.First(&ecode, "id = ?", id).Error
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("fails to find error code by id")
		return nil, err
	}

	return &ecode, nil
}

func (repository *EcodeRepository) Update(ecode *Ecode) (*Ecode, error) {
	log.Debug().Any("ecode", ecode).Msg("update error code")

	// Omit the column name from update...
	err := repository.pg.Omit("created_at").Save(&ecode).Error
	if err != nil {
		log.Error().Err(err).Msg("fails to update id")
		return nil, err
	}

	return ecode, err
}

func (repository *EcodeRepository) Insert(ecode *Ecode) (*Ecode, error) {
	log.Debug().Any("ecode", ecode).Msg("insert a new ecode")

	err := repository.pg.Create(&ecode).Error
	if err != nil {
		log.Error().Err(err).Msg("fails to insert new ecode")
		return nil, err
	}
	return ecode, err
}

func (repository *EcodeRepository) DeleteById(id string) (string, error) {
	log.Debug().Str("id", id).Msg("delete a ecode by id")

	err := repository.pg.Delete(&Ecode{}, id).Error
	if err != nil {
		log.Error().Str("id", id).Msg("fails to delete ecode by id")
		return "", err
	}
	return id, err
}
