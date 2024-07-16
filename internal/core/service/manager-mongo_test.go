package service

import (
	"go-hex-arch/internal/core/port/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionExampleDivide_Success(t *testing.T) {
	mck := &mocks.ManagerAccountRepository{}
	mckMS := &mocks.ManagerAccountService{}
	ma := NewManagerAccountMongoService(mck, mckMS)

	res, err := ma.FunctionExampleDivide(3, 3)
	assert.Equal(t, 1, res, "Dividing value of 3 and 3 is 1")
	assert.NoError(t, err, "There are no errors")
}

func TestFunctionExampleDivide_Fail(t *testing.T) {
	ma := NewManagerAccountMongoService(nil, nil)

	_, err := ma.FunctionExampleDivide(3, 0)
	assert.Error(t, err, "There are has errors")
}

func TestFunctionExampleDivide_Float_Success(t *testing.T) {
	ma := NewManagerAccountMongoService(nil, nil)

	res, err := ma.FunctionExampleDivide(200, 199)
	assert.Equal(t, 1, res, "Dividing value(rounded to the smallest) of 200 and 199 is 1")
	assert.NoError(t, err, "There are no errors")
}
