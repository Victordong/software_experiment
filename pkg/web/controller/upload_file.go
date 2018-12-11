package controller

import (
	"auto_fertilizer_back/pkg/web/plugin"
	"github.com/disintegration/imaging"
	"image"
	"mime/multipart"
	"strings"
	"time"
)

var AllowTags []string

func init() {
	AllowTags = []string{"jpg", "jpeg", "png", "JPG", "JPEG", "PNG"}
}

func JudgeTag(filename string) (string, bool) {
	parts := strings.Split(filename, ".")
	tag := parts[len(parts)-1]
	for _, AllowTag := range AllowTags {
		if tag == AllowTag {
			return tag, true
		}
	}
	return "", false
}
func Upload_image(file *multipart.FileHeader, path string) (string, error) {
	tag, ifAllow := JudgeTag(file.Filename)
	if !ifAllow {
		return "", plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: "图片格式不允许",
		}
	}
	file.Filename = time.Now().Format("20060102150405") + "." + tag
	src, err := file.Open()
	if err != nil {
		return "", plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: err.Error(),
		}
	}
	defer src.Close()
	cropImage, _, err := image.Decode(src)
	bonds := cropImage.Bounds()
	imgX := bonds.Dx()
	imgY := bonds.Dy()
	if imgX > imgY {
		imgY = imgY / imgX * 320
		imgX = 320
	} else {
		imgX = imgX / imgY * 320
		imgY = 320
	}
	cropImage = imaging.Resize(cropImage, imgX, imgY, imaging.Lanczos)
	err = imaging.Save(cropImage, "./assets/"+path+file.Filename)
	if err != nil {
		return "", plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: err.Error(),
		}
	}
	return "/static/" + path + file.Filename, nil
}
