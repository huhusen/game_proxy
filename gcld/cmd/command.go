package cmd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sockets-proxy/util"
	"sockets-proxy/util/binaryext"
	"sockets-proxy/util/log"
)

type Command struct {
	Cmd  CMD
	Zh   string
	Send interface{}
	Rec  interface{}
}

func (c *Command) UpdateRec(data []byte) {
	util.Byte2Struct(data, &c.Rec)
	s, _ := json.Marshal(c.Rec)
	log.Log.Infof("接收%s数据,Data:\n%s", c.Zh, string(s))
}
func (c *Command) UpdateSend(data string) {
	c.Send = data
	log.Log.Infof("发送%s数据,Data:%s", c.Zh, data)
}
func (c *Command) Name() string {
	return c.Zh
}
func (c *Command) CMD() CMD {
	return c.Cmd
}
func (c *Command) Map() map[string]interface{} {
	return c.Rec.(map[string]interface{})
}

func (c *Command) PacketData(body string) []byte {
	bs := []byte(body)
	writer := binaryext.NewByteArray()
	writer.WriteInt(32 + 4 + len(bs))
	writer.Write([]byte(c.Cmd))
	writer.WriteInt(0)
	writer.Write(util.NewBuf(32 - len(c.Cmd)))
	writer.Write(bs)
	return writer.Data()
}
func (l *Command) Data() []byte {
	return nil
}
func (l *Command) Hex() {
	d := l.Data()
	fmt.Println(hex.EncodeToString(d))
	cmd := NewSendData(d)
	cmd.Print(l.Name())
}
