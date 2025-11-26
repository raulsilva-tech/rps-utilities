package usecase

import (
	"io"
	"net/http"
	"time"

	"github.com/icholy/digest"
)

type SendGETRequestUseCase struct{}

type SendGETRequestUseCaseOutput struct {
	HTTPCode     int    `json:"code"`
	HTTPResponse string `json:"response"`
}

func NewSendGETRequestUseCase() *SendGETRequestUseCase {
	return &SendGETRequestUseCase{}
}

func (uc *SendGETRequestUseCase) Execute(url, user, password, auth string) (SendGETRequestUseCaseOutput, error) {

	var client *http.Client

	switch auth {
	case "digest":
		client = &http.Client{
			Transport: &digest.Transport{
				Username: user,
				Password: password,
			},
			Timeout: time.Second * 5,
		}
	default:
		client = &http.Client{}
	}

	resp, err := client.Get(url)
	if err != nil {
		return SendGETRequestUseCaseOutput{0, err.Error()}, err
	}

	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("Error reading response:", err)
		return SendGETRequestUseCaseOutput{0, err.Error()}, err
	}

	// fmt.Println("Response Status:", resp.Status)
	// fmt.Println("Response Body:", string(body))
	return SendGETRequestUseCaseOutput{resp.StatusCode, string(body)}, nil
}
