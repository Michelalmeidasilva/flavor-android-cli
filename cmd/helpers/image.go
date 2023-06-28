package helpers

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

type ImageAndroid struct {
	size             uint
	nameDrawablePath string
	nameMipMapPath   string
}

func ResizeImage(originalImagePath string, androidPathResources string, images []ImageAndroid) {
	if originalImagePath == "" {
		return
	}

	originalFile, err := os.Open(originalImagePath)
	if err != nil {
		fmt.Println("Erro ao abrir a imagem original no path informado", err)
		return
	}
	defer originalFile.Close()

	originalImage, _, err := image.Decode(originalFile)
	if err != nil {
		fmt.Println("Erro ao decodificar a imagem original:", err)
		return
	}

	for _, image := range images {
		resizedImage := resize.Resize(image.size, image.size, originalImage, resize.Lanczos3)

		CreateImage(resizedImage, image.size, "logo.png", androidPathResources+image.nameDrawablePath)
		CreateImage(resizedImage, image.size, "launcher_ic.png", androidPathResources+image.nameMipMapPath)
	}

	fmt.Println("Images created at following path:" + Message(androidPathResources, "yellow") + "\n")

}

func CreateImage(value image.Image, size uint, imageName string, path string) {
	copyImagePath := fmt.Sprintf(path + "/" + imageName)
	copyFile, err := os.Create(copyImagePath)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo da cópia:", err)
		return
	}
	defer copyFile.Close()

	err = png.Encode(copyFile, value)
	if err != nil {
		fmt.Println("Erro ao codificar a cópia em PNG:", err)
		return
	}
}

func CreateAndroidImages(originalImage string, androidPath string) {

	images := []ImageAndroid{
		{size: 48, nameDrawablePath: "drawable-mdpi", nameMipMapPath: "mipmap-mdpi"},
		{size: 72, nameDrawablePath: "drawable-hdpi", nameMipMapPath: "mipmap-hdpi"},
		{size: 96, nameDrawablePath: "drawable-xhdpi", nameMipMapPath: "mipmap-xhdpi"},
		{size: 144, nameDrawablePath: "drawable-xxhdpi", nameMipMapPath: "mipmap-xxhdpi"},
		{size: 240, nameDrawablePath: "drawable-xxxhdpi", nameMipMapPath: "mipmap-xxxhdpi"},
	}

	ResizeImage(originalImage, androidPath, images)
}
