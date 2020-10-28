package image

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/phpdi/ant/util"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"math"
	"os"
	"strings"
)

//图片缩略
type ImageAbbreviation struct {
	maxWidth  float64
	maxHeight float64
}

func NewImageAbbreviation(maxWidth, maxHeight float64) *ImageAbbreviation {
	return &ImageAbbreviation{
		maxWidth:  maxWidth,
		maxHeight: maxHeight,
	}
}

// 计算图片缩放后的尺寸
func (this *ImageAbbreviation) calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(this.maxWidth/float64(srcWidth), this.maxHeight/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// 生成缩略图
func (this *ImageAbbreviation) MakeThumbnail(imagePath, savePath string) (err error) {

	var (
		arr []string
	)
	file, _ := os.Open(imagePath)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := this.calculateRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// 调用resize库进行图片缩放
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	if err = this.createDir(savePath); err != nil {
		return
	}

	if arr = strings.Split(savePath, "."); len(arr) < 2 {
		return
	}

	suffix := strings.ToLower(arr[len(arr)-1])

	// 需要保存的文件
	imgfile, _ := os.Create(savePath)
	defer imgfile.Close()

	switch suffix {
	case "jpg", "jpeg":
		if err = jpeg.Encode(imgfile, m, nil); err != nil {
			return err
		}
	case "png":
		// 以PNG格式保存文件
		if err = png.Encode(imgfile, m); err != nil {
			return err
		}

	default:

	}

	return nil
}

//生成目录
func (this *ImageAbbreviation) createDir(filePath string) (err error) {
	var (
		dir string
		ok  bool
	)

	arr := strings.Split(filePath, "/")

	dir = strings.Join(arr[0:len(arr)-1], "/")

	if ok, err = util.PathExists(dir); err != nil {
		return
	}
	//今日目录不存在创建
	if !ok {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return
		}
	}

	return

}
