// arquivo: plugins/cpu.go
package plugins

import (
	"fmt"
	"errors"

	"github.com/raffs/go-labs/sistema"
)

type Cpu struct {
	So string
}

func (c *Cpu) Iniciar() {
	fmt.Println("Incializando o plugin de CPU", c.So)
}

func (c *Cpu) Coletar() (err error) {
	usoCpu := sistema.CpuUso()

	if usoCpu < 0 {
		err = errors.New("Falha ao coletar a cpu")
		return err
	}

	fmt.Printf("CPU: %f\n", usoCpu)
	return nil
}

func (c *Cpu) Describe() {
	fmt.Printf("Plugin para monitorar CPU no sistema %s\n", c.So)
}