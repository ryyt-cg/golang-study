package ecode

import (
	"error-code-lookup/pkg/test"
	"testing"

	"gopkg.in/go-playground/assert.v1"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/suite"
)

type EcodeRepoTestSuite struct {
	suite.Suite
	postgresql      *embeddedpostgres.EmbeddedPostgres
	ecodeRepository *EcodeRepository
}

// This will run before the tests in the suite are run
func (suite *EcodeRepoTestSuite) SetupSuite() {
	suite.postgresql = test.PgStart(suite.T(), "../test/migrations")
	suite.ecodeRepository = getEcodeRepository(suite.T())
}

func (suite *EcodeRepoTestSuite) TearDownSuite() {
	err := suite.postgresql.Stop()
	if err != nil {
		suite.T().Fatal(err)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestEcodeRepoTestSuite(t *testing.T) {
	suite.Run(t, new(EcodeRepoTestSuite))
}

func getEcodeRepository(t *testing.T) *EcodeRepository {
	db, err := test.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return NewEcodeRepository(db)
}

func (suite *EcodeRepoTestSuite) Test_FindById() {
	var testCases = []struct {
		input    string
		expected Ecode
	}{
		{"dip1", Ecode{
			ID:          "dip1",
			Description: "connection refuse",
		}},
		{"dip2", Ecode{
			ID:          "dip2",
			Description: "connection timeout",
		}},
	}

	for _, testCase := range testCases {
		ecode, _ := suite.ecodeRepository.FindById(testCase.input)
		assert.Equal(suite.T(), testCase.expected.ID, ecode.ID)
		assert.Equal(suite.T(), testCase.expected.Description, ecode.Description)
	}
}

func (suite *EcodeRepoTestSuite) Test_FindById_NotFound() {
	var testCases = []struct {
		input    string
		expected Ecode
	}{
		{"dip101", Ecode{}},
	}

	for _, testCase := range testCases {
		ecode, err := suite.ecodeRepository.FindById(testCase.input)
		assert.Equal(suite.T(), err, nil)
		assert.Equal(suite.T(), ecode, nil)
	}
}

func (suite *EcodeRepoTestSuite) Test_UpdateEcode() {
	var backups = []struct {
		input    string
		expected Ecode
	}{
		{"dip1", Ecode{
			ID:          "dip1",
			Description: "connection refuse",
		}},
		{"dip2", Ecode{
			ID:          "dip2",
			Description: "connection timeout",
		}},
	}

	var testCases = []struct {
		input    string
		expected Ecode
	}{
		{"DIP-201", Ecode{
			ID:          "DIP-201",
			Description: "RAMs connection refused",
		}},
		{"DIP-202", Ecode{
			ID:          "DIP-202",
			Description: "auth validation failed",
		}},
	}

	for _, testCase := range testCases {
		suite.ecodeRepository.Update(&testCase.expected)
		ecode, _ := suite.ecodeRepository.FindById(testCase.input)
		assert.Equal(suite.T(), testCase.expected.ID, ecode.ID)
		assert.Equal(suite.T(), testCase.expected.Description, ecode.Description)
	}

	// backup the data to original state
	for _, backup := range backups {
		suite.ecodeRepository.Update(&backup.expected)
	}
}

func (suite *EcodeRepoTestSuite) Test_InsertEcode() {
	var testCases = []struct {
		input    string
		expected Ecode
	}{
		{"dip-50001", Ecode{
			ID:          "dip-50001",
			Description: "user errors",
		}},
		{"dip-50002", Ecode{
			ID:          "dip-50002",
			Description: "sqs connection timeout",
		}},
	}

	for _, testCase := range testCases {
		suite.ecodeRepository.Insert(&testCase.expected)
		ecode, _ := suite.ecodeRepository.FindById(testCase.input)
		assert.Equal(suite.T(), testCase.expected.ID, ecode.ID)
		assert.Equal(suite.T(), testCase.expected.Description, ecode.Description)
	}
}
