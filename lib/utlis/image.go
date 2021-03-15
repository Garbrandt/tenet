package utlis

import (
	"encoding/base64"
	"image"
	"strings"

	"github.com/disintegration/imaging"
)

//图片处理
const (
	MethodThumbnailShortCropCenter int = 11 //按短边缩放，居中裁剪
	MethodThumbnailLongFillShort   int = 12 //按长边缩放，缩略填充
	MethodFixedWidthAutoHeight     int = 21 //固定宽度，高度自适应
	MethodFixedHeightAutoWidth     int = 22 //固定高度，宽度自适应
)

//图片处理
//img:图片对象
//width:需要将图片处理成的宽度
//height:需要将图片处理成的高度
//method:图片处理方式，默认按短边缩放居中裁剪
//is_zoom:是否允许图片放大
func Process(img image.Image, width, height, method int, isZoom bool) image.Image {
	//获取需要处理的图片的原图宽高函数
	ImgInfo := func(img image.Image) (int, int) {
		Max := img.Bounds().Max
		Min := img.Bounds().Min
		w := Max.X - Min.X
		h := Max.Y - Min.Y
		return w, h
	}
	//计算原图宽高，然后根据比例获取长边
	w, h := ImgInfo(img)

	switch method {
	case MethodThumbnailShortCropCenter:
		//禁止图片放大
		if isZoom == false && (width > w || height > h) {
			img = imaging.Fit(img, width, height, imaging.Linear)
		} else {
			img = imaging.Fill(img, width, height, imaging.Center, imaging.Linear)
		}
		//imaging.Thumbnail(img, width, height,imaging.Linear)
	case MethodThumbnailLongFillShort:

		if float64(w)/float64(h) > float64(width)/float64(height) {
			//禁止图片放大
			if isZoom == false && width > w {
				width = w
			}
			img = imaging.Resize(img, width, 0, imaging.Linear) //需要填充高
		} else {
			//禁止图片放大
			if isZoom == false && height > h {
				height = h
			}
			img = imaging.Resize(img, 0, height, imaging.Linear) //需要填充宽
		}

		//创建白色的背景图片
		b, _ := base64.StdEncoding.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAB3RJTUUH4QICBRURmsv9YQAAAA9JREFUCB0BBAD7/wD///8F/gL+A30ZxgAAAABJRU5ErkJggg==")
		background, _, _ := image.Decode(strings.NewReader(string(b)))
		//设置背景图片大小
		background = imaging.Resize(background, width, height, imaging.NearestNeighbor)
		//获取经过处理了一回之后的图片宽高
		w, h = ImgInfo(img)
		img = imaging.Paste(background, img, image.Point{(width - w) / 2, (height - h) / 2})
	case MethodFixedWidthAutoHeight:
		//禁止图片放大
		if isZoom == false && width > w {
			width = w
		}
		img = imaging.Resize(img, width, 0, imaging.Linear)
	case MethodFixedHeightAutoWidth:
		//禁止图片放大
		if isZoom == false && height > h {
			height = h
		}
		img = imaging.Resize(img, 0, height, imaging.Linear)
	default:
		//禁止图片放大
		if isZoom == false && (width > w || height > h) {
			img = imaging.Fit(img, width, height, imaging.Linear)
		} else {
			img = imaging.Fill(img, width, height, imaging.Center, imaging.Linear)
		}
	}
	return img
}
