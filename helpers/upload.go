package helpers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

// Función para detectar el tipo de imagen por su cabecera
func detectImageType(base64Image string) (string, error) {
	// Limpiar la cadena base64 para eliminar el prefijo "data:image/...;base64,"
	segments := strings.Split(base64Image, ",")
	if len(segments) < 2 {
		return "", fmt.Errorf("la cadena base64 no es válida")
	}

	// Obtener el prefijo de tipo de imagen (ejemplo: image/png, image/jpeg, etc.)
	imageDataPrefix := segments[0]
	if strings.Contains(imageDataPrefix, "image/png") {
		return "png", nil
	} else if strings.Contains(imageDataPrefix, "image/jpeg") {
		return "jpeg", nil
	} else if strings.Contains(imageDataPrefix, "image/gif") {
		return "gif", nil
	}
	return "", fmt.Errorf("Tipo de imagen no soportado")
}

// Función para subir la imagen
func UploadImage(base64Image string) (string, error) {
	// Detectar el tipo de imagen
	format, err := detectImageType(base64Image)
	if err != nil {
		return "", err
	}

	// Limpiar la cadena base64 para eliminar el prefijo de los datos
	segments := strings.Split(base64Image, ",")
	decodedImage, err := base64.StdEncoding.DecodeString(segments[1])
	if err != nil {
		return "", fmt.Errorf("error al decodificar la imagen base64: %v", err)
	}

	// Crear el nombre de archivo y la ruta donde guardar la imagen
	fileName := fmt.Sprintf("%d.%s", time.Now().Unix(), format) // Usar el formato detectado (png, jpeg, etc.)
	filePath := filepath.Join("storage", "products", fileName)

	// Guardar el archivo en el sistema de archivos
	err = ioutil.WriteFile(filePath, decodedImage, 0644)
	if err != nil {
		return "", fmt.Errorf("error al guardar la imagen: %v", err)
	}

	// Retornar la ruta del archivo guardado
	return fileName, nil
}
