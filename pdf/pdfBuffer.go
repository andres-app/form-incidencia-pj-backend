package pdf

import (
	"api-pdf/helper"
	"api-pdf/modelo"
	"bytes"
	"strconv"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"

	"github.com/johnfercher/maroto/pkg/props"
)

func CrearPdf(pdfData modelo.Data) (*bytes.Buffer, error) {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)

	builHeading(m)
	builBody(m, pdfData)

	bufferPdf, err := m.Output()
	if err != nil {
		return nil, err
	}

	return &bufferPdf, nil
}

func builHeading(m pdf.Maroto) {

	m.RegisterHeader(func() {

		m.SetBorder(true)

		m.Row(25, func() {
			m.Col(3, func() {
				err := m.FileImage("pdf/logo.png", props.Rect{
					Center:  true,
					Percent: 75,
				})

				if err != nil {
					m.Text("Logo", props.Text{
						Top:  1,
						Size: 12,
						// Style: consts.Bold,
						Align: consts.Center,
					})
				}
			})

			m.Col(6, func() {
				m.Text("PODER JUDICIAL", props.Text{
					Top:   5,
					Style: consts.Bold,
					Align: consts.Center,
					Size:  16,
				})
				m.Text("FORMULARIO DE INCIDENCIA", props.Text{
					Top:   15,
					Align: consts.Center,
					Size:  13,
				})
			})

			m.Col(3, func() {
				m.Text("Formulario N°:", props.Text{
					Top:   1,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: darkPurpleColor(),
				})
				m.Text("P-01-2023-GI-GG-PJ-F-01", props.Text{
					Top:   7,
					Align: consts.Center,
				})

				m.Text("Versión: ", props.Text{
					Top:   13,
					Style: consts.Bold,
					Right: 10,
					Align: consts.Middle,
				})
				m.Text("05", props.Text{
					Top:   13,
					Left:  10,
					Align: consts.Middle,
				})

				m.Text("Fecha: ", props.Text{
					Top:   19,
					Style: consts.Bold,
					Right: 15,
					Align: consts.Middle,
				})
				m.Text("08/06/2023", props.Text{
					Top:   19,
					Left:  15,
					Align: consts.Middle,
				})

			})

		})

		m.SetBorder(false)

		m.Line(8, props.Line{
			Color: color.NewWhite(),
		})

		m.SetBorder(true)

	})

}

func builBody(m pdf.Maroto, info modelo.Data) {

	// INFORMACIÓN DEL SISTEMA
	m.SetBorder(true)

	m.SetBackgroundColor(helper.GrayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("INFORMACIÓN DEL SISTEMA", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(8, func() {
		m.Col(3, func() {
			m.Text("Nombre", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.NombreSistema, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text("Versión", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.VersionSistema, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})

	// INFORMACIÓN DEL USUARIO

	m.SetBackgroundColor(helper.GrayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("INFORMACIÓN DEL USUARIO", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(8, func() {
		m.Col(3, func() {
			m.Text("Nombres y Apellidos", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.UsuarioNombre, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text("Celular / Anexo", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.CelularAxeso, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})
	m.Row(8, func() {
		m.Col(3, func() {
			m.Text("Sede", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.Sede, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text("Cargo", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.Cargo, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})

	// REPORTADO POR

	m.SetBackgroundColor(helper.GrayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("REPORTADO POR", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(8, func() {
		m.Col(3, func() {
			m.Text("Nombres y Apellidos", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.PersonaReporte, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text("Celular", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.CelularPersona, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})

	// INCIDENCIA

	m.SetBackgroundColor(helper.GrayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("INCIDENCIA", props.Text{
				Top:   1,
				Size:  12,
				Right: 25,
				Style: consts.Bold,
				Align: consts.Middle,
			})
			m.Text(info.Fecha, props.Text{
				Top:   1,
				Size:  12,
				Left:  25,
				Style: consts.Bold,
				Align: consts.Middle,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	textMultiline := props.Text{
		Extrapolate: false,
		Top:         1,
		Left:        1,
		Size:        10,
		Style:       consts.Bold,
		Align:       consts.Left,
	}

	m.Row(helper.CalcRowHeight(m, info.Descripcion, textMultiline, 12, 12)+2, func() {
		m.Col(12, func() {
			m.Text(info.Descripcion, textMultiline)
		})
	})

	// DESCARTES

	m.SetBackgroundColor(helper.GrayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("DESCARTES", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Middle,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())
	m.Row(8, func() {
		m.Col(10, func() {
			m.Text("1. ¿El mismo incidente se reproduce en otro equipo? ", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})

		m.Col(1, func() {
			m.Text(helper.TernarioSiNo(info.PreguntaUno, "si"), props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(1, func() {
			m.Text(helper.TernarioSiNo(info.PreguntaUno, "no"), props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.Row(8, func() {
		m.Col(10, func() {
			m.Text("2. ¿El mismo incidente se reproduce con otros usuarios? ", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(1, func() {
			m.Text(helper.TernarioSiNo(info.PreguntaDos, "si"), props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(1, func() {
			m.Text(helper.TernarioSiNo(info.PreguntaDos, "no"), props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.Row(8, func() {
		m.Col(10, func() {
			m.Text("3. ¿El incidente ocurre solo con un expediente? ", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(1, func() {
			m.Text(helper.TernarioSiNo(info.PreguntaTres, "si"), props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(1, func() {
			m.Text(helper.TernarioSiNo(info.PreguntaTres, "no"), props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.Row(8, func() {
		m.Col(10, func() {
			m.Text("4. ¿Lo reportado ha sido validado por implantación? ", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(1, func() {
			m.Text(helper.TernarioSiNo(info.PreguntaCuatro, "si"), props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(1, func() {
			m.Text(helper.TernarioSiNo(info.PreguntaCuatro, "no"), props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.Row(8, func() {
		m.Col(10, func() {
			m.Text("5. ¿Se está utilizando la última versión de la aplicación desplegada en la corte? ", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(1, func() {
			m.Text(helper.TernarioSiNo(info.PreguntaCinco, "si"), props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(1, func() {
			m.Text(helper.TernarioSiNo(info.PreguntaCinco, "no"), props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	// FLUJO

	m.SetBackgroundColor(helper.GrayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("FLUJO", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Middle,
			})
		})
	})

	var count int
	for _, item := range info.Imagenes {
		count++

		m.SetBackgroundColor(color.NewWhite())

		textMultiline := props.Text{
			Extrapolate: false,
			Top:         6,
			Left:        1,
			Size:        10,
			Style:       consts.Bold,
			Align:       consts.Left,
		}

		m.Row(helper.CalcRowHeight(m, item.Descripcion, textMultiline, 12, 12)+2, func() {
			m.Col(12, func() {
				m.Text("Pantalla_"+helper.NumbertToString(count), props.Text{
					Top:   1,
					Left:  1,
					Size:  10,
					Style: consts.Bold,
					Align: consts.Left,
				})
				m.Text("(Obligatorio SIJ):", props.Text{
					Top:   1,
					Left:  21,
					Size:  10,
					Style: consts.Bold,
					Align: consts.Left,
					Color: color.Color{
						Red:   224,
						Green: 36,
						Blue:  37,
					},
				})
				m.Text("- Descripción breve de la captura de pantalla: ", props.Text{
					Top:   1,
					Left:  51,
					Size:  10,
					Style: consts.Bold,
					Align: consts.Left,
				})
				m.Text(item.Descripcion, textMultiline)
			})
		})

		m.Row(75, func() {
			m.Col(12, func() {
				err := m.FileImage("tmp/"+strconv.Itoa(count)+"output."+item.Extension, props.Rect{
					Center:  true,
					Percent: 95,
				})

				if err != nil {
					m.Text("Captura", props.Text{
						Top:  1,
						Size: 12,
						// Style: consts.Bold,
						Align: consts.Center,
					})
				}
			})
		})
	}

	// Parámetros de prueba de Desarrollo*
	var cortecsj modelo.CorteCsj

	cortecsjs, _ := helper.LeerArchivo("cortecsj.json")
	for _, item := range cortecsjs {
		if item.Id == info.IdCorteCsj {
			cortecsj = item
			break
		}
	}

	m.SetBackgroundColor(helper.GrayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("Parámetros de prueba de Desarrollo*", props.Text{
				Top:   1,
				Left:  1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(20, func() {
		m.Col(6, func() {
			m.Text("[Nombre del Servicio**]: "+cortecsj.NombreServicio, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[Nombre BD]: "+cortecsj.NombreDb, props.Text{
				Top:   7,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[User Id]: "+cortecsj.UserId, props.Text{
				Top:   13,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(6, func() {
			m.Text("[Password]: "+cortecsj.PasswordDb, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[IP]: "+cortecsj.Ip, props.Text{
				Top:   7,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[Puerto]: "+cortecsj.Puerto, props.Text{
				Top:   13,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})

	})

	// Parámetros de prueba de Calidad (Testing)*

	m.SetBackgroundColor(helper.GrayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("Parámetros de prueba de Calidad (Testing)*", props.Text{
				Top:   1,
				Left:  1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(20, func() {
		m.Col(6, func() {
			m.Text("[Nombre del Servicio**]: "+cortecsj.NombreServicioTest, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[Nombre BD]: "+cortecsj.NombreDbTest, props.Text{
				Top:   7,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[User Id]: "+cortecsj.UserIdTest, props.Text{
				Top:   13,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(6, func() {
			m.Text("[Password]: "+cortecsj.PasswordDbTest, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[IP]: "+cortecsj.IpTest, props.Text{
				Top:   7,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[Puerto]: "+cortecsj.PuertoTest, props.Text{
				Top:   13,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})

	})

	m.SetBorder(false)

}
