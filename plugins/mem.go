// arquivo: plugins/cpu.go
package plugins

import (
	"fmt"
	"errors"

	"github.com/raffs/go-labs/sistema"
)

type Mem struct {
	So string
}

func (m *Mem) Iniciar() {
	fmt.Println("Incializando o plugin de memória", m.So)
}

func (m *Mem) Coletar() (err error) {
	usoMem := sistema.MemUso()

	if usoMem < 0 {
		err = errors.New("Falha ao coletar a memória")
		return err
	}

	fmt.Printf("Mémória: %f\n", usoMem)
	return nil
}

func (m *Mem) Describe() {
	fmt.Printf("Plugin para monitorar memória no sistema %s\n", m.So)
}