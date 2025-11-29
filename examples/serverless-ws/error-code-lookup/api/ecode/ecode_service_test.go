package ecode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetEcodeById(t *testing.T) {
	ecodeMock := MockEcodeRepositorier{}
	ecode := &Ecode{
		ID:          "DIP-001",
		Description: "connection refused",
	}
	ecodeMock.On("FindById", "DIP-001").Return(ecode, nil)

	ecodeService := NewEcodeService(&ecodeMock)
	result, _ := ecodeService.GetEcodeById("DIP-001")
	ecodeMock.AssertExpectations(t)
	ecodeMock.AssertNumberOfCalls(t, "FindById", 1)

	assert.Equal(t, ecode.ID, result.ID)
	assert.Equal(t, ecode.Description, result.Description)
}

func Test_UpdateEcode(t *testing.T) {
	ecodeMock := MockEcodeRepositorier{}
	ecode := &Ecode{
		ID:          "DIP-001",
		Description: "DNS configuration issue",
	}
	ecodeMock.On("Update", ecode).Return(ecode, nil)

	ecodeService := NewEcodeService(&ecodeMock)
	result, _ := ecodeService.Update(ecode)
	ecodeMock.AssertExpectations(t)
	ecodeMock.AssertNumberOfCalls(t, "Update", 1)

	assert.Equal(t, ecode.ID, result.ID)
	assert.Equal(t, ecode.Description, result.Description)
}

func Test_InsertEcode(t *testing.T) {
	ecodeMock := MockEcodeRepositorier{}
	ecode := &Ecode{
		ID:          "DIP-0015",
		Description: "VPN server issue",
	}
	ecodeMock.On("Insert", ecode).Return(ecode, nil)

	ecodeService := NewEcodeService(&ecodeMock)
	result, _ := ecodeService.Insert(ecode)
	ecodeMock.AssertExpectations(t)
	ecodeMock.AssertNumberOfCalls(t, "Insert", 1)

	assert.Equal(t, ecode.ID, result.ID)
	assert.Equal(t, ecode.Description, result.Description)
}
