package loader

import (
	"image/png"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func uniPath(relativ string) (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	base := filepath.Dir(filename)
	path := filepath.Join(base, "..", "..", relativ)

	return path, nil
}

func LoadFyneResource(relativ string, name string) fyne.Resource {

	path, err := uniPath(relativ)
	if err != nil {
		log.Printf("Ошибка генерации пути: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Ошибка чтения %v", path)
	}

	if len(data) == 0 {
		log.Printf("Прочитанный файл пуст: %v", data)
	}

	result := fyne.NewStaticResource(name, data)

	return result
}

func LoadFyneCanvasImage(relativ string) *canvas.Image {

	path, err := uniPath(relativ)
	if err != nil {
		log.Printf("Ошибка генерации пути: %v", err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("Ошибка открытия файла изображения: %v", err)
	}
	defer file.Close()

	image, err := png.Decode(file)
	if err != nil {
		log.Printf("Ошибка кодировки изображения: %v", err)
	}

	canvasImage := canvas.NewImageFromImage(image)

	return canvasImage
}
