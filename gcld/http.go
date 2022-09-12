package gcld

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"sockets-proxy/util"
	"strings"
)

func (bot *Bot) OnHttpRequestEvent(request *http.Request) {
	//url := request.URL.String()
	//reqReader := bufio.NewReader(request.Body)
	//reqBody, _ := io.ReadAll(reqReader)
	//log.Log.Printf("[%s] %s\nRequestBody:%s", request.Method, url, string(reqBody))
}
func (bot *Bot) OnHttpResponseEvent(response *http.Response) {
	url := response.Request.URL.String()
	var reader io.Reader
	if strings.Contains(url, "gateway.action") {
		reader = bufio.NewReader(response.Body)
		body, _ := io.ReadAll(reader)
		if strings.HasPrefix(hex.EncodeToString(body), "789") {
			body = util.ZlibUnCompress(body)
			str := string(body)
			if strings.Contains(str, "blackNames") {
				bot.playerPlayerinfo.UpdateRec(body)
				bot.playerPlayerinfo.Update1()
				fmt.Println()
			}
			//else if strings.Contains(str, "playerList") {
			//	//util.Byte2Struct(body, &bot.playerPlayerlist.Rec2)
			//} else if strings.Contains(str, "totalOutput") {
			//	//util.Byte2Struct(body, &bot.buildingMaincityinfo.Rec2)
			//} else {
			//	//log.Log.Printf("ResponseBody:%s", str)
			//}
		}
	}
}
