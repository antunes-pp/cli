package user

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/antunes-pp/cli/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUserUseCase_Execute(t *testing.T) {
	expect := assert.New(t)

	inputRegisterUser := mocks.NewMockInputRegisterUser(t)
	httpClient := mocks.NewMockHttpClient(t)

	uc := NewRegisterUserUserCase(httpClient)

	inputRegisterUser.EXPECT().GetName().Return("Gabriel")
	inputRegisterUser.EXPECT().GetEmail().Return("gabriel@gmail.com")
	inputRegisterUser.EXPECT().GetSquads().Return([]string{
		"recommendation",
		"storehome",
	})

	message := "created"

	res := http.Response{
		Status:     "200 OK",
		StatusCode: http.StatusCreated,
		Body:       io.NopCloser(strings.NewReader(message)),
	}

	httpClient.EXPECT().Do(mock.AnythingOfType("*http.Request")).Return(&res, nil)

	err := uc.Execute(inputRegisterUser, true)

	expect.Nil(err)
}
