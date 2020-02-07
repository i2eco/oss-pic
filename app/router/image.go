package router

import (
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/goecology/oss-pic/app/pkg/conf"
	"github.com/goecology/oss-pic/app/pkg/mus"
	"go.uber.org/zap"
	"image"
	"image/color"
	"image/png"
	"strconv"
	"strings"
)

func Info(c *gin.Context) {
	url := strings.Split(c.Request.RequestURI, "/")
	length := len(url)
	if length < 3 {
		c.String(200, "%s", "params error")
		return
	}
	space := url[1]
	imageName := strings.Join(url[2:length-1], "/")
	scale := url[length-1]

	config, ok := conf.Conf.Image[space]
	if !ok {
		c.String(200, "%s", "space1 not exist")
		return
	}

	thumbnails_size := strings.Split(scale, "_")
	if len(thumbnails_size) != 2 {
		c.String(200, "%s", "scale error")
		return
	}

	thumbnails_width, _ := strconv.Atoi(thumbnails_size[0])
	thumbnails_height, _ := strconv.Atoi(thumbnails_size[1])

	img, err := imaging.Open(config.Path + "/" + imageName)
	if err != nil {
		mus.Logger.Error("img open error", zap.String("err", err.Error()))
		c.String(200, "%s", "open file error")
		return
	}
	thumb := imaging.Thumbnail(img, thumbnails_width, thumbnails_height, imaging.CatmullRom)
	dst := imaging.New(thumbnails_width, thumbnails_height, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumb, image.Pt(0, 0))

	header := c.Writer.Header()

	header.Add("Content-Type", "image/jpeg")

	c.Status(200)

	png.Encode(c.Writer, dst)
}
