package usecase

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/icholy/digest"
)

type SendGETStreamRequestUseCase struct{}

type SendGETStreamRequestUseCaseOutput struct {
	HTTPCode     int    `json:"code"`
	HTTPResponse string `json:"response"`
}

func NewSendGETStreamRequestUseCase() *SendGETRequestUseCase {
	return &SendGETRequestUseCase{}
}

func (uc *SendGETStreamRequestUseCase) Execute(url, user, password, auth string) (SendGETStreamRequestUseCaseOutput, error) {

	var client *http.Client

	switch auth {
	case "digest":
		client = &http.Client{
			Transport: &digest.Transport{
				Username: user,
				Password: password,
			},
			Timeout: time.Second * 0,
		}
	default:
		client = &http.Client{}
	}

	// Faça a requisição GET
	resp, err := client.Get(url)
	if err != nil {
		return SendGETStreamRequestUseCaseOutput{0, "TESTE" + err.Error()}, err
	}
	defer resp.Body.Close()

	fmt.Println("Conectado! Aguardando eventos...")

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	emptyLines := 0
	event := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(">>", line)

		if strings.TrimSpace(line) == "" {
			emptyLines++
			if emptyLines == 2 {
				fmt.Println("Evento completo:", event)
				event = make(map[string]string)
				emptyLines = 0
			}
			continue
		}

		emptyLines = 0

		// Se linha tem formato "key: value", adiciona ao evento
		if parts := strings.SplitN(line, ": ", 2); len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			event[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return SendGETStreamRequestUseCaseOutput{0,  err.Error()}, err
	}

	// fmt.Println("Response Status:", resp.Status)
	// fmt.Println("Response Body:", string(body))
	return SendGETStreamRequestUseCaseOutput{resp.StatusCode, ""}, nil
}
