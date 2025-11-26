package usecase

import (
	"github.com/go-vgo/robotgo"
)

type ClickUseCase struct{}

func NewClickUseCase() *ClickUseCase {
	return &ClickUseCase{}
}

func (uc *ClickUseCase) Execute(x, y int) {

	robotgo.Move(x, y) // Move o mouse para (x, y)
	//time.Sleep(100 * time.Millisecond) // Pequeno delay opcional
	robotgo.Click() // Clica com o botão esquerdo (segundo argumento: duplo clique?)

}
