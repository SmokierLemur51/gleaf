package utils


// image resizing, python would probably be better suited for this though fr fr

// func Resize(imgFile string, height, width int) error {
// 	file, err := os.Open(fmt.Sprintf("/static/img/%s", imgFile))
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		return err
// 	}

// 	resizedImage := image.NewRGBA(image.Rect(0, 0, height, width))
// 	draw.CatmullRom.Scale(resizedImage, resizedImage.Bounds(), img, img.Bounds(), draw.Over, nil)

// 	return nil
// }
