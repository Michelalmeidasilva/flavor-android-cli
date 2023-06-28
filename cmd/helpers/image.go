package helpers

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

type ImageParameters struct {
	size             uint
	nameDrawablePath string
	nameMipMapPath   string
}

func ResizeImage(iconPath string, androidPathResources string) {
	if iconPath == "" {
		return
	}

	// Abrir o arquivo da imagem de entrada original
	originalFile, err := os.Open(iconPath)
	if err != nil {
		fmt.Println("Erro ao abrir a imagem original no path informado", err)
		return
	}
	defer originalFile.Close()

	// Decodificar a imagem original
	originalImage, _, err := image.Decode(originalFile)
	if err != nil {
		fmt.Println("Erro ao decodificar a imagem original:", err)
		return
	}

	// Tamanhos desejados para as cópias
	images := []ImageParameters{
		{size: 48, nameDrawablePath: "drawable-mdpi", nameMipMapPath: "mipmap-mdpi"},
		{size: 72, nameDrawablePath: "drawable-hdpi", nameMipMapPath: "mipmap-hdpi"},
		{size: 96, nameDrawablePath: "drawable-xhdpi", nameMipMapPath: "mipmap-xhdpi"},
		{size: 144, nameDrawablePath: "drawable-xxhdpi", nameMipMapPath: "mipmap-xxhdpi"},
		{size: 240, nameDrawablePath: "drawable-xxxhdpi", nameMipMapPath: "mipmap-xxxhdpi"},
	}
	for _, image := range images {
		// Redimensionar a imagem original para o tamanho desejado
		resizedImage := resize.Resize(image.size, image.size, originalImage, resize.Lanczos3)

		CreateImage(resizedImage, image.size, "logo.png", androidPathResources+image.nameDrawablePath)
		CreateImage(resizedImage, image.size, "launcher_ic.png", androidPathResources+image.nameMipMapPath)
	}
}

func CreateImage(value image.Image, size uint, imageName string, path string) {
	// Criar o arquivo de saída para a cópia
	copyImagePath := fmt.Sprintf(path + "/" + imageName)
	copyFile, err := os.Create(copyImagePath)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo da cópia:", err)
		return
	}
	defer copyFile.Close()

	// Codificar a cópia em PNG e escrever no arquivo de saída
	err = png.Encode(copyFile, value)
	if err != nil {
		fmt.Println("Erro ao codificar a cópia em PNG:", err)
		return
	}

	fmt.Println("Cópia redimensionada para", size, "salva em:", copyImagePath)

}
