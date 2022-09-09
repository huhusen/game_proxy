package gcld

import (
	"encoding/hex"
	"fmt"
	"sockets-proxy/util"
	"testing"
)

func TestLogin(t *testing.T) {
	buf, err := hex.DecodeString("0000002f63686172696f74406765744270496e666f0000000000000000000000000000000000000069643d3326627049643d31")
	if err != nil {
		panic(err)
	}
	cmd := NewSendData(buf)

	fmt.Println(cmd.String())
	cmd.Print()
}
func TestRecv(t *testing.T) {
	buf, err := hex.DecodeString("000000607075736840706c6179657200000000000000000000000000000000000000000000000000789cab564a4c2ec9cccf53b2aa562a2e492c4955b232d6514a492c4904899416a48085aa9532128b831373128b2a95ac4a8a4a536b810000578614da")
	if err != nil {
		panic(err)
	}
	r := NewRecvData(buf)
	r.Print()

}
func TestNewRequest(t *testing.T) {
	//0000004c6c6f67696e5f757365720000000000000000000000000000000000000000000000000000757365726b65793d3539636264353463363265363539353030636234633566323231313032623237
	NewChariotForgeSpInfoRequest("1", "1").Hex()
	NewChariotForgeSpInfoRequest("1", "2").Hex()
	NewChariotForgeSpInfoRequest("1", "3").Hex()
	NewChariotGetBpInfoRequest("1", "3").Hex()
	NewChariotforgeBpInfoRequest("1").Hex()
	NewChariotGetInfoRequest("3", "")
}
func TestNewResp(t *testing.T) {
	//0000004c6c6f67696e5f757365720000000000000000000000000000000000000000000000000000757365726b65793d3539636264353463363265363539353030636234633566323231313032623237
	data := []byte(`{"action":{"state":1,"data":{"sessionId":"6B4EA2C4526A9C37B4498A0436F9E37A"}}}`)
	m := util.NewMap(data)
	fmt.Println(m)
}
