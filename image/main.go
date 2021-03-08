package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"
)

func main() {
	Handle()
}
func Circle(str string) *CirCleMask {
	file, err := os.Create("newCircle.png")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	imageFile, err := os.Open(str)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer imageFile.Close()

	srcImg , _ := png.Decode(imageFile)
	w := srcImg.Bounds().Max.X - srcImg.Bounds().Min.X
	h := srcImg.Bounds().Max.Y - srcImg.Bounds().Min.Y
	d := w
	fmt.Printf("w = %d, h = %d", w,h)
	if w > h {
		d = h
	}
	dstImg := NewCircleMask(srcImg,image.Point{d/2,d/2},d)
	return dstImg
	//png.Encode(file,dstImg)
}
func Handle() {
	//target, err := imaging.Open("/Users/lt/Desktop/123.png")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	bg, err := imaging.Open("/Users/lt/Desktop/bg.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	circle := Circle("/Users/lt/Desktop/123.png")

	dst := imaging.Resize(circle, 157,157, imaging.Lanczos)

	result := imaging.Overlay(bg, dst, image.Pt(239,516),1)

	result = imaging.Overlay(result,dst,image.Pt(379,516),1)
	writeOnImage(result)
	filename := "/Users/lt/work/priv/go/learn/image/key.jpg"
	err = imaging.Save(result,filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func writeOnImage(target *image.NRGBA) {
	c := freetype.NewContext()
	c.SetDPI(256)
	c.SetClip(target.Bounds())
	c.SetDst(target)
	c.SetHinting(font.HintingFull)

	c.SetSrc(image.NewUniform(color.RGBA{255,255,255,180}))
	c.SetFontSize(10)
	fontFam, err := getFontFamily()
	if err != nil {
		fmt.Println("get font family error")
	}
	c.SetFont(fontFam)

	width := len("习近平")
	fmt.Println(width)
	pt := freetype.Pt(215 + (150 - width/3 * 40),695)
	_, err = c.DrawString("习近平", pt)
	if err != nil {
		fmt.Printf("draw error: %v \n", err)
	}
	pt = freetype.Pt(368,695)
	_, err = c.DrawString("&", pt)
	if err != nil {
		fmt.Printf("draw error: %v \n", err)
	}
	pt = freetype.Pt(397,695)
	_, err = c.DrawString("彭丽媛", pt)
	if err != nil {
		fmt.Printf("draw error: %v \n", err)
	}
}
func getFontFamily()(*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile("/Users/lt/work/priv/go/learn/image/方正准圆.TTF")
	if err != nil {
		fmt.Println("read file error:", err)
		return &truetype.Font{}, err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("parse font error:", err)
		return &truetype.Font{}, err
	}

	return f, err
}

type  CirCleMask struct {
	image image.Image
	point image.Point
	diameter int
}
func NewCircleMask(img image.Image, p image.Point, d int) *CirCleMask {
	return &CirCleMask{
		image: img,
		point: p,
		diameter: d,
	}
}
func (c *CirCleMask)ColorModel() color.Model {
	return c.image.ColorModel()
}
func (c *CirCleMask)Bounds() image.Rectangle {
	return image.Rect(c.point.X-c.diameter/2, c.point.Y- c.diameter/2, c.point.X + c.diameter/2,c.point.Y + c.diameter/2)
}
func (c *CirCleMask)At(x,y int) color.Color {
	d := c.diameter
	dis := math.Sqrt(math.Pow(float64(x-d/2),2) + math.Pow(float64(y -d/2),2) )
	fmt.Println("dis = ", dis)
	fmt.Println("d/2 = ", float64(d)/2)
	if dis > float64(d)/2 {
		return c.image.ColorModel().Convert(color.RGBA{255,255,255,0})
	} else {
		return c.image.At( x, y)
	}
	//xx, yy, rr := float64(x-c.point.X)+0.5, float64(y-c.point.Y)+0.5, float64(c.diameter/2)
	//if xx*xx+yy*yy < rr*rr {
	//	return c.image.ColorModel().Convert(color.RGBA{255,255,255,0})
	//}
	//	return c.image.At(c.point.X + x,c.point.Y + y)
}