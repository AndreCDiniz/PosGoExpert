package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

// template.Parse e template.Must têm maneiras diferentes de lidar com erros durante a criação de templates.
// A função template.Parse permite um controle mais refinado sobre a manipulação de erros, enquanto template.Must
// é mais adequada para situações em que você deseja que a execução seja interrompida se um template não puder ser
// carregado/interpretado corretamente.
func main() {
	curso1 := Curso{"Go", 40}
	tmp := template.New("CursoTemplate")
	temp, err := tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	if err != nil {
		panic(err)
	}
	err = temp.Execute(os.Stdout, curso1)
	if err != nil {
		panic(err)
	}
}
