package handler

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	pb "user/proto"

	model "user/model"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	"github.com/alibabacloud-go/tea/tea"
)

type User struct{}

var phoneNumber string
var Code string

func (e *User) SendSmsCode(ctx context.Context, req *pb.SendSmsCodeRequest, rsp *pb.SendSmsCodeResponse) error {
	phoneNumber = req.PhoneNumber

	boolS := model.CheckImg(req.Uuid, req.CaptCode)
	if !boolS {
		rsp.Message = "captchaCode false; 图片验证码错误,请重新输入"
		rsp.Ok = "0"
	} else {
		rsp.Message = "captchaCode success; 获取成功，请注意短信 "
		rsp.Ok = "1"
		_main(tea.StringSlice(os.Args[1:]))
	}
	return nil
}

func CreateClient() (_result *openapi.Client, _err error) {
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
	// 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &openapi.Client{}
	_result, _err = openapi.NewClient(config)
	return _result, _err
}

func CreateApiInfo() (_result *openapi.Params) {
	params := &openapi.Params{
		// 接口名称
		Action: tea.String("SendSms"),
		// 接口版本
		Version: tea.String("2017-05-25"),
		// 接口协议
		Protocol: tea.String("HTTPS"),
		// 接口 HTTP 方法
		Method:   tea.String("POST"),
		AuthType: tea.String("AK"),
		Style:    tea.String("RPC"),
		// 接口 PATH
		Pathname: tea.String("/"),
		// 接口请求体内容格式
		ReqBodyType: tea.String("json"),
		// 接口响应体内容格式
		BodyType: tea.String("json"),
	}
	_result = params
	return _result
}

func _main(args []*string) (_err error) {

	// 设置随机数种子，确保每次运行结果不同
	rand.Seed(time.Now().UnixNano())
	// 生成 0 到 9999 之间的随机数
	randomNumber := rand.Intn(10000)
	// 将随机数格式化为四位数，不足四位数前面补零，并赋值给 code 变量
	Code = fmt.Sprintf("%04d", randomNumber)
	err := model.SavePhoneCode(Code, phoneNumber)
	if err != nil {
		fmt.Println("redis save false", err)
	}

	client, _err := CreateClient()
	if _err != nil {
		return _err
	}

	params := CreateApiInfo()
	// query params
	queries := map[string]interface{}{}
	queries["PhoneNumbers"] = tea.String(phoneNumber)
	queries["SignName"] = tea.String("zlican博客")
	queries["TemplateCode"] = tea.String("SMS_474916099")
	queries["TemplateParam"] = tea.String("{\"code\": " + Code + "}")
	queries["SmsUpExtendCode"] = nil
	queries["OutId"] = nil
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
	}
	// 复制代码运行请自行打印 API 的返回值
	// 返回值实际为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
	_, _err = client.CallApi(params, request, runtime)
	if _err != nil {
		return _err
	}
	return _err
}
