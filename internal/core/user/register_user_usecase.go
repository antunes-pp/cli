package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/antunes-pp/cli/internal/core"
	"github.com/antunes-pp/cli/internal/core/port"
)

type RegisterUserUseCase struct {
	httpClient port.HttpClient
}

func NewRegisterUserUserCase(client port.HttpClient) *RegisterUserUseCase {
	r := RegisterUserUseCase{
		httpClient: client,
	}

	return &r
}

func (it *RegisterUserUseCase) Execute(input port.InputRegisterUser, isDev bool) error {
	user := NewUser(input.GetName(), input.GetEmail(), input.GetSquads())

	log.Println(user)

	if err := user.Validate(); err != nil {
		return err
	}

	j, _ := json.Marshal(&user)

	var url string

	if isDev {
		url = core.MS_QA + "/user"
	} else {
		url = core.MS_PROD + "/user"
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := it.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()

		if err != nil {
			log.Println(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusCreated {
		return errors.New(fmt.Sprintf("⚠️ Status code should be %d but it is %d", http.StatusCreated, res.StatusCode))
	}

	return nil
}
