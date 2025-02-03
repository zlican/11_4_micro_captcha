package handler

import (
	"context"
	"encoding/json"
	"image/color"

	"getCaptcha/model"
	pb "getCaptcha/proto"

	"github.com/afocus/captcha"
)

var cap *captcha.Captcha

type GetCaptcha struct{}

func (e *GetCaptcha) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	cap = captcha.New()
	if err := cap.SetFont("./conf/comic.ttf"); err != nil { //路径是看main.go文件，而不是调用函数位置
		panic(err.Error())
	}
	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{5, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	img, str := cap.Create(4, captcha.NUM)

	//存储图片str到redis数据库
	err := model.SaveImgCode(str, req.Uuid)
	if err != nil {
		return err
	}
	//将生成的图片序列化
	imgBuf, _ := json.Marshal(img)

	//将imgBuf传出
	rsp.Img = imgBuf
	return nil
}
