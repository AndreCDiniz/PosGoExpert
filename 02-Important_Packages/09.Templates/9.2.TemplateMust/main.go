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
	curso2 := Curso{"Go", 60}
	t := template.Must(template.New("CursoTemplate").Parse("\nCurso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, curso2)
	if err != nil {
		panic(err)
	}
}
