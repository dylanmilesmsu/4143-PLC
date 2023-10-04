package main

import (
	"fmt"

	"github.com/dylanmilesmsu/img_mod/Colors"
	"github.com/dylanmilesmsu/img_mod/GetPic"
	"github.com/dylanmilesmsu/img_mod/Grayscale"
	"github.com/dylanmilesmsu/img_mod/Text"
)

func main() {
	fmt.Print("Downloading image to ez.png...\n")
	GetPic.GetImageFromUrl("https://s3e8p5g8.rocketcdn.me/wp-content/uploads/2020/11/midwestern-state-university2.jpg", "ez.png")
	fmt.Print("Done\n")
	fmt.Print("Converting image to grayscale...\n")
	Grayscale.ConvertImageToGrayscale("ez.png", "gray.png")
	fmt.Print("Done\n")
	fmt.Print("Making image with the text \"COLORS\" on it....\n")
	Text.CreateImageFromText("COLORS", "text.png", 500, 500)
	fmt.Print("Done, saved to text.png\n")
	fmt.Print("Turning ez.png into pz.txt by color...\n")
	Colors.GetColorsFromImage("ez.png", "pz.txt")
	fmt.Print("Done.\n")
}
