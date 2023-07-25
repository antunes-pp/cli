package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antunes-pp/cli/internal/core"
	"github.com/antunes-pp/cli/internal/core/port"
	"io"
	"log"
	"net/http"
)

type RegisterUserUseCase struct{}

func NewRegisterUserUserCase() *RegisterUserUseCase {
	r := RegisterUserUseCase{}

	return &r
}

func (r *RegisterUserUseCase) Execute(input port.InputRegisterUser, isDev bool) error {
	user := NewUser(input.GetName(), input.GetEmail(), input.GetSquads())

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

	client := http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

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
