package helper

import (
	"api-pdf/modelo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func GrayColor() color.Color {
	return color.Color{
		Red:   220,
		Green: 220,
		Blue:  220,
	}
}

func CalcRowHeight(m pdf.Maroto, text string, textProp props.Text, gridWidth, colWidth uint) float64 {
	pdfM := m.(*pdf.PdfMaroto)
	percent := float64(colWidth) / float64(gridWidth)
	pageWidth, _ := pdfM.Pdf.GetPageSize()
	left, _, right, _ := pdfM.Pdf.GetMargins()
	width := (pageWidth - right - left) * percent
	lines := 1.0
	if !textProp.Extrapolate {
		lines = float64(pdfM.TextHelper.GetLinesQuantity(text, textProp, width))
	}
	fontHeight := textProp.Size/pdfM.Font.GetScaleFactor() + textProp.VerticalPadding
	return lines*fontHeight + textProp.Top
}

func CrearCarpeta(rutaCarpeta string) {
	// Verificar si la carpeta existe
	_, err := os.Stat(rutaCarpeta)
	if os.IsNotExist(err) {
		// La carpeta no existe, crearla
		err := os.MkdirAll(rutaCarpeta, 0755)
		if err != nil {
			fmt.Printf("Error al crear la carpeta: %s\n", err)
			return
		}
		fmt.Println("Carpeta creada exitosamente.")
	} else if err != nil {
		// Ocurrió un error al verificar la existencia de la carpeta
		fmt.Printf("Error al verificar la carpeta: %s\n", err)
		return
	} else {
		// La carpeta ya existe
		fmt.Println("La carpeta ya existe.")
	}
}

func TernarioSiNo(valor string, evaluar string) string {
	if strings.ToLower(valor) == strings.ToLower(evaluar) {
		return strings.ToUpper(valor)
	}
	return ""
}

func LeerArchivo(ruta string) ([]modelo.CorteCsj, error) {
	// Abrir el archivo JSON
	file, err := os.Open(ruta)
	if err != nil {
		return nil, fmt.Errorf("Error al abrir el archivo: %v", err)
	}

	defer file.Close()

	// Leer el contenido del archivo
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error al leer el archivo: %v", err)
	}

	// Decodificar el contenido en la estructura de datos
	var cortecsjs []modelo.CorteCsj
	err = json.Unmarshal(content, &cortecsjs)
	if err != nil {
		return nil, fmt.Errorf("Error al decodificar el JSON: %v", err)
	}

	// Acceder a los datos leídos
	return cortecsjs, nil
}
