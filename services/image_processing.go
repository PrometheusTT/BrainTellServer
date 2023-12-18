package services

import (
	"BrainTellServer/ao"
	"BrainTellServer/do"
	"BrainTellServer/utils"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type ModelParam struct {
	Args string          `json:"args"`
	User UserVerifyParam `json:"user"`
}

func (param *ModelParam) String() string {
	jsonres, err := json.Marshal(param)
	if err != nil {
		return ""
	}
	return string(jsonres)
}

func (param *ModelParam) FromJsonString(jsonstr string) (utils.RequestParam, error) {
	if err := json.Unmarshal([]byte(jsonstr), param); err != nil {
		return nil, err
	}
	return param, nil
}

type CMDParam struct {
	CmdName string          `json:"cmdname"`
	Args    string          `json:"args"`
	User    UserVerifyParam `json:"user"`
}

func (param *CMDParam) String() string {
	jsonres, err := json.Marshal(param)
	if err != nil {
		return ""
	}
	return string(jsonres)
}

func (param *CMDParam) FromJsonString(jsonstr string) (utils.RequestParam, error) {
	if err := json.Unmarshal([]byte(jsonstr), param); err != nil {
		return nil, err
	}
	return param, nil
}

func GetPluginList(w http.ResponseWriter, r *http.Request) {
	var p UserVerifyParam
	_, err := utils.DecodeFromHttp(r, &p)
	if err != nil {
		utils.EncodeToHttp(w, 500, err.Error())
		return
	}

	if _, err := ao.Login(&do.UserInfo{
		Name:   p.Name,
		Passwd: p.Passwd,
	}); err != nil {
		utils.EncodeToHttp(w, 401, err.Error())
		return
	}

	str, err := utils.GetPluginList()
	if err != nil {
		utils.EncodeToHttp(w, 501, err.Error())
		return
	}
	utils.EncodeToHttp(w, 200, str)
}
func GetImage(w http.ResponseWriter, r *http.Request) {
	var p CMDParam
	param, err := utils.DecodeFromHttp(r, &p)
	if err != nil {
		utils.EncodeToHttp(w, 500, err.Error())
		return
	}

	_, ok := param.(*CMDParam)
	if !ok {
		log.WithFields(log.Fields{
			"event": "ConvertImage",
			"desc":  "param.(*ConvertImage) failed",
		}).Warnf("%v\n", err)
		utils.EncodeToHttp(w, 500, "")
		return
	}

	if _, err := ao.Login(&do.UserInfo{
		Name:   p.User.Name,
		Passwd: p.User.Passwd,
	}); err != nil {
		utils.EncodeToHttp(w, 401, err.Error())
		return
	}
	p.Args = strings.Replace(p.Args, "\\", "", -1)
	utils.SendFileNoDelete(w, 200, p.Args)
}
func GetImageProcessingList(w http.ResponseWriter, r *http.Request) {
	var p UserVerifyParam
	_, err := utils.DecodeFromHttp(r, &p)
	if err != nil {
		utils.EncodeToHttp(w, 500, err.Error())
		return
	}

	if _, err := ao.Login(&do.UserInfo{
		Name:   p.Name,
		Passwd: p.Passwd,
	}); err != nil {
		utils.EncodeToHttp(w, 401, err.Error())
		return
	}
	str, err := utils.ListImageFiles()
	fmt.Printf(str)
	if err != nil {
		utils.EncodeToHttp(w, 501, err.Error())
		return
	}
	utils.EncodeToHttp(w, 200, str)
}

func ConvertImage(w http.ResponseWriter, r *http.Request) {
	var p CMDParam
	param, err := utils.DecodeFromHttp(r, &p)
	if err != nil {
		utils.EncodeToHttp(w, 500, err.Error())
		return
	}

	_, ok := param.(*CMDParam)
	if !ok {
		log.WithFields(log.Fields{
			"event": "ConvertImage",
			"desc":  "param.(*ConvertImage) failed",
		}).Warnf("%v\n", err)
		utils.EncodeToHttp(w, 500, "")
		return
	}

	if _, err := ao.Login(&do.UserInfo{
		Name:   p.User.Name,
		Passwd: p.User.Passwd,
	}); err != nil {
		utils.EncodeToHttp(w, 401, err.Error())
		return
	}
	p.Args = strings.Replace(p.Args, "\\", "", -1)

	out, err := utils.ConvertImageType(p.Args)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "ConvertImage",
			"desc":  "param.(*ConvertImage) failed",
			"error": err.Error(),
			"Args":  p.Args,
		}).Warnf("%v\n", err)
		utils.EncodeToHttp(w, 501, err.Error())
		return
	}
	utils.SendFileNoDelete(w, 200, out)
}

func CommandLine(w http.ResponseWriter, r *http.Request) {
	var p CMDParam
	param, err := utils.DecodeFromHttp(r, &p)
	if err != nil {
		utils.EncodeToHttp(w, 500, err.Error())
		return
	}

	_, ok := param.(*CMDParam)
	if !ok {
		log.WithFields(log.Fields{
			"event": "ConvertImage",
			"desc":  "param.(*ConvertImage) failed",
		}).Warnf("%v\n", err)
		utils.EncodeToHttp(w, 500, "")
		return
	}

	if _, err := ao.Login(&do.UserInfo{
		Name:   p.User.Name,
		Passwd: p.User.Passwd,
	}); err != nil {
		utils.EncodeToHttp(w, 401, err.Error())
		return
	}

	p.Args = strings.Replace(p.Args, "\\", "", -1)
	out, err := utils.V3DCommandLine(p.CmdName, p.Args)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "V3DCommand",
			"desc":  "param.(*execCmd) failed",
		}).Warnf("%v\n", err)
		utils.EncodeToHttp(w, 501, err.Error())
		return
	}
	time.Sleep(time.Duration(5) * time.Second)
	utils.SendFileNoDelete(w, 200, out)

}

func CallModel(w http.ResponseWriter, r *http.Request) {
	var p ModelParam
	param, err := utils.DecodeFromHttp(r, &p)
	if err != nil {
		utils.EncodeToHttp(w, 500, err.Error())
		return
	}

	_, ok := param.(*ModelParam)
	if !ok {
		log.WithFields(log.Fields{
			"event": "callModel",
			"desc":  "param.(*callModel) failed",
		}).Warnf("%v\n", err)
		utils.EncodeToHttp(w, 500, "")
		return
	}

	if _, err := ao.Login(&do.UserInfo{
		Name:   p.User.Name,
		Passwd: p.User.Passwd,
	}); err != nil {
		utils.EncodeToHttp(w, 401, err.Error())
		return
	}

	p.Args = strings.Replace(p.Args, "\\", "", -1)
	out, err := utils.CallPredictService(p.Args)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "V3DCommand",
			"desc":  "param.(*execCmd) failed",
		}).Warnf("%v\n", err)
		utils.EncodeToHttp(w, 501, err.Error())
		return
	}
	time.Sleep(time.Duration(5) * time.Second)
	utils.SendFileNoDelete(w, 200, out)

}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// 写入要发送到浏览器的文本内容
	_, err := fmt.Fprint(w, "Hello, Browser!")
	if err != nil {
		return
	}

}
