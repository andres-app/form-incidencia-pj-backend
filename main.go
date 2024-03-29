package main

import (
	"api-pdf/helper"
	"api-pdf/modelo"
	"api-pdf/pdf"
	"fmt"
	_ "image/jpeg" // Importa el formato JPEG
	_ "image/png"  // Importa el formato PNG
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar las variables de entorno desde el archivo .env
	godotenv.Load()

	// Obtener el valor de la variable de entorno GO_PORT
	var go_port string = os.Getenv("GO_PORT")
	var ruta_log string = os.Getenv("RUTA_LOG")

	// Crear archivo log
	f, err := os.Create(ruta_log)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	router := gin.Default()

	// Middleware para el LOGGER
	router.Use(gin.Logger())

	// Middleware para CORS
	router.Use(cors.Default())

	// Middleware para manejar el error 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Ruta no encontrada."})
	})

	// Rutas
	router.GET("/", func(c *gin.Context) {
		log.Println("Endpoint ping")

		htmlResponse := `
		<!DOCTYPE html>
		<html lang="es">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Mi Servidor Go</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					padding: 20px;
				}
				.container {
					max-width: 600px;
					margin: 0 auto;
					text-align: center;
				}
				.message {
					font-size: 24px;
					color: #333;
				}
			</style>
		</head>
		<body>
		<div>
		<h1 style="text-align: center;">Bienvenido a mi Servidor Go</h1>
		<p style="text-align: center;">&iexcl;<strong>Hola el servicio esta corriendo correctamente&nbsp;</strong>desde mi servidor Go!</p>
		<p><img style="display: block; margin-left: auto; margin-right: auto;" src="https://openupthecloud.com/wp-content/uploads/2020/01/Golang.png" width="384" height="215" /></p>
		</div>
		</body>
		</html>
		`

		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlResponse))
	})
	router.POST("/pdf", handlePDFRequestGin)
	router.GET("/cortecsj", handleListaCsj)

	router.Run(go_port)
}

func handlePDFRequestGin(c *gin.Context) {
	var data modelo.Data

	if err := c.BindJSON(&data); err != nil {
		log.Println("No se pudo parsear el body, " + err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No se pudo parsear el body"})
		return
	}

	helper.CrearCarpeta("tmp/")

	var count int
	for _, item := range data.Imagenes {
		base64String := "data:image/" + item.Extension + ";base64," + item.Base64String

		imageData := helper.ExtractImageData(base64String)
		if imageData == nil {
			fmt.Println("No se pudo extraer la imagen base64")
			return
		}

		imageType := helper.ExtractImageType(base64String)
		if imageType == "" {
			fmt.Println("No se pudo determinar el tipo de imagen")
			return
		}

		count++
		err := helper.SaveImage(imageData, imageType, "tmp/"+strconv.Itoa(count)+"output."+item.Extension)
		if err != nil {
			fmt.Println("Error al guardar la imagen:", err)
			return
		}
	}

	pdfBytes, err := pdf.CrearPdf(data)
	if err != nil {
		fmt.Println("Error en crear el pdf:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	pdfMemory, err := ioutil.ReadAll(pdfBytes)
	if err != nil {
		fmt.Println("Error en leer la imagen:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Envía el flujo de bytes del PDF como respuesta
	c.Header("Content-Type", "application/pdf")
	// c.Header("Content-Disposition", "attachment; filename=pdf.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfMemory)
}

func handleListaCsj(c *gin.Context) {
	cortecsjs, err := helper.LeerArchivo("cortecsj.json")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, cortecsjs)
}
