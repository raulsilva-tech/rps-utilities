package usecase

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/icholy/digest"
)

type FBIGetFaceUseCase struct{}

func NewFBIGetFaceUseCase() *FBIGetFaceUseCase {
	return &FBIGetFaceUseCase{}
}

func (uc *FBIGetFaceUseCase) Execute(host string, port int, user, password, url string, timeout int) (Output, error) {

	//LIGA CAPTURA

	//requere usuario cadastrar face
	// http://192.168.15.4/cgi-bin/accessControl.cgi?action=captureCmd&type=1&heartbeat=5&timeout=30
	if url == "" {
		url = "/cgi-bin/accessControl.cgi?action=captureCmd&type=1&heartbeat=5&timeout=" + fmt.Sprintf("%d", timeout)
	}
	output, err := digestRequest(host, port, user, password, url)
	if err != nil {
		return output, err
	}

	fullUrl := fmt.Sprintf("http://%s:%d/cgi-bin/snapManager.cgi?action=attachFileProc&Flags[0]=Event&Events=[CitizenPictureCompare]", host, port)
	resp, err := openStream(fullUrl, user, password, timeout)
	if err != nil {
		return Output{-1, err.Error()}, err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return Output{resp.StatusCode, "Erro ao interpretar Content-Type: " + err.Error()}, err
	}

	boundary, ok := params["boundary"]
	if !ok {
		return Output{resp.StatusCode, "Boundary não encontrado no Content-Type"}, err
	}

	//PROCESSAR SNAPMANAGER PARA OBTER A IMAGEM DA CAPTURA
	reader := multipart.NewReader(resp.Body, boundary)

	var msg string
	for {

		part, err := reader.NextPart()
		if err != nil {
			if err == io.EOF || errors.Is(err, io.EOF) {
				msg = "Stream encerrado"
				break
			}
			continue
		}

		if finished, err := handlePart(part); finished {

			if err != nil {
				return Output{-1, err.Error()}, err
			}
			return Output{resp.StatusCode, "Imagem capturada"}, err
		}

	}

	return Output{0, msg}, nil
}
func handlePart(part *multipart.Part) (bool, error) {

	if part.Header.Get("Content-Type") != "image/jpeg" {
		// Aqui pode fechar, porque vai continuar o loop
		part.Close()
		return false, nil
	}

	file, err := os.Create("fbi_snapshot.jpg")
	if err != nil {
		return true, err
	}
	defer file.Close()

	err = readJPEG(part, file)
	if err != nil {
		return true, err
	}

	// fmt.Println("Imagem salva: fbi_snapshot.jpg")
	return true, nil
}

func readJPEG(r io.Reader, w io.Writer) error {
	buf := make([]byte, 8192)
	var last byte

	for {
		n, err := r.Read(buf)
		if n > 0 {
			w.Write(buf[:n])

			for i := 0; i < n; i++ {
				if last == 0xFF && buf[i] == 0xD9 {
					return nil // fim do JPEG
				}
				last = buf[i]
			}
		}

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func openStream(url, user, password string, timeout int) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "multipart/x-mixed-replace")

	client := &http.Client{
		Transport: &digest.Transport{
			Username: user,
			Password: password,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				DisableKeepAlives: false,
			},
		},
		Timeout: time.Duration(timeout) * time.Second, // stream contínuo
	}

	return client.Do(req)
}
