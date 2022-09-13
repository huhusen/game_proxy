package gcld

import (
	"encoding/hex"
	"fmt"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto"
	"os"
	"sockets-proxy/gcld/cmd"
	"sockets-proxy/gcld/cmd/chariot"
	"sockets-proxy/gcld/cmd/player"
	"sockets-proxy/util"
	"testing"
)

func TestLogin(t *testing.T) {
	buf, err := hex.DecodeString("0000002f63686172696f74406765744270496e666f0000000000000000000000000000000000000069643d3326627049643d31")
	if err != nil {
		panic(err)
	}
	cmd := cmd.NewSendData(buf)

	fmt.Println(cmd.String())
}
func TestRecv(t *testing.T) {
	//buf, err := hex.DecodeString("000000607075736840706c6179657200000000000000000000000000000000000000000000000000789cab564a4c2ec9cccf53b2aa562a2e492c4955b232d6514a492c4904899416a48085aa9532128b831373128b2a95ac4a8a4a536b810000578614da")
	//if err != nil {
	//	panic(err)
	//}
	//r := cmd.NewRecvData(buf)
	//r.Print(r.Command)

}
func TestNewRequest(t *testing.T) {
	//0000004c6c6f67696e5f757365720000000000000000000000000000000000000000000000000000757365726b65793d3539636264353463363265363539353030636234633566323231313032623237

	chariot.NewForgeSpRequest("1", "1").Hex()

	chariot.NewForgeSpRequest("1", "2").Hex()
	chariot.NewForgeSpRequest("1", "3").Hex()
	chariot.NewForgeSpRequest("1", "3").Hex()
	chariot.NewForgeBpInfoRequest("1").Hex()
	chariot.NewGetInfoRequest("3", "")
}
func TestNewResp(t *testing.T) {
	//0000004c6c6f67696e5f757365720000000000000000000000000000000000000000000000000000757365726b65793d3539636264353463363265363539353030636234633566323231313032623237
	//data := []byte(`{"action":{"state":1,"data":{"sessionId":"6B4EA2C4526A9C37B4498A0436F9E37A"}}}`)
	//m := util.NewMap(data)
	//fmt.Println(m)
}
func TestJson(t *testing.T) {
	s := `{"action":{"state":1,"data":{"blackNames":[],"endTime":0,"player":{"playerId":16661,"playerLv":17,"playerName":"1","userId":"120336","yx":"gcld","forceId":2,"pic":4,"gold":6566000,"vipLv":1,"pkey":"36043a58b0ce74e909311475d61aac4e","pkey2":"1662946501044_36043a58b0ce74e909311475d61aac4e","copper":2000000826,"copperMax":22000,"copperOutput":545,"food":6000000000,"foodMax":10000,"foodOutput":15,"wood":2000004888,"woodMax":10000,"woodOutput":35,"iron":2000000000,"ironMax":10000,"ironOutput":0,"exp":25794,"expNeed":1000,"forces":0,"forcesMax":0,"inPveBattle":false,"inOccupyBattle":false,"serverTime":"2022/09/12/09/35/01"},"ssTime":1662946501045,"chargeItems":{"buildingCd":{"alert":1,"show":true,"cost":1,"vipLv":0},"refreshGeneralCd":{"alert":1,"show":true,"cost":1,"vipLv":0},"storeBuyItem":{"alert":2,"show":true,"cost":0,"vipLv":0},"refreshStoreCd":{"alert":1,"show":true,"cost":2,"vipLv":0},"buyStorehouseSize":{"alert":2,"show":true,"cost":90,"vipLv":0},"techCd":{"alert":1,"show":true,"cost":1,"vipLv":0},"moveCd":{"alert":2,"show":true,"cost":2,"vipLv":0},"goldRecruit":{"alert":1,"show":true,"cost":1,"vipLv":0},"resourceMode1":{"alert":2,"show":true,"cost":10,"vipLv":0},"resourceMode2":{"alert":2,"show":true,"cost":20,"vipLv":0},"resourceMode3":{"alert":2,"show":true,"cost":50,"vipLv":0},"politicsEvent":{"alert":1,"show":true,"cost":0,"vipLv":0},"loadGem":{"alert":2,"show":true,"cost":0,"vipLv":0},"unloadGem":{"alert":1,"show":true,"cost":0,"vipLv":0},"shaodaozi":{"alert":2,"show":true,"cost":1,"vipLv":0},"nverhong":{"alert":2,"show":true,"cost":3,"vipLv":0},"zhuyeqing":{"alert":2,"show":true,"cost":5,"vipLv":0},"incenseGold":{"alert":1,"show":true,"cost":0,"vipLv":0},"blackMarketGold":{"alert":1,"show":true,"cost":1,"vipLv":0},"nationalRankBatBuyNumGold":{"alert":2,"show":true,"cost":50,"vipLv":0},"buyBonusArmyGold":{"alert":1,"show":true,"cost":20,"vipLv":0},"quenchingGold":{"alert":1,"show":true,"cost":5,"vipLv":0},"slashSlave":{"alert":2,"show":true,"cost":2,"vipLv":0},"slaveFreedom":{"alert":1,"show":true,"cost":1,"vipLv":0},"buyWeaponItem":{"alert":1,"show":true,"cost":5,"vipLv":0},"buyPowerExtraGold":{"alert":2,"show":true,"cost":100,"vipLv":0},"freeConsGold":{"alert":1,"show":true,"cost":5,"vipLv":0},"consDrawingGold":{"alert":1,"show":true,"cost":1,"vipLv":0},"kfDqGold":{"alert":1,"show":true,"cost":8,"vipLv":0},"youdiGold":{"alert":2,"show":true,"cost":1,"vipLv":0},"jiebingGold":{"alert":1,"show":true,"cost":20,"vipLv":0},"chujiGold":{"alert":2,"show":true,"cost":1,"vipLv":0},"createBattleTeam":{"alert":1,"show":true,"cost":5,"vipLv":0},"copperInvest":{"alert":1,"show":true,"cost":1,"vipLv":0},"battleTeamInspire":{"alert":1,"show":true,"cost":10,"vipLv":0},"coverManzuShoumaiCd":{"alert":1,"show":true,"cost":1,"vipLv":0},"teamOrder":{"alert":1,"show":true,"cost":1,"vipLv":0},"updateLash":{"alert":2,"show":true,"cost":0,"vipLv":0},"useGoldOrder":{"alert":1,"show":true,"cost":1000,"vipLv":0},"kfzbBuySup":{"alert":1,"show":true,"cost":2,"vipLv":0},"recoverGold":{"alert":1,"show":true,"cost":1,"vipLv":0},"irongive":{"alert":1,"show":true,"cost":10,"vipLv":0},"gemDamo":{"alert":1,"show":true,"cost":5,"vipLv":0},"gemJinglian":{"alert":1,"show":true,"cost":10,"vipLv":0},"buyCrack":{"alert":1,"show":true,"cost":50,"vipLv":0},"farmQuickFinish":{"alert":1,"show":true,"cost":1,"vipLv":0},"buyGoldDice":{"alert":1,"show":true,"cost":40,"vipLv":0},"ancientBuyTime":{"alert":1,"show":true,"cost":100,"vipLv":0},"diamondShopComplementGold":{"alert":1,"show":true,"cost":20,"vipLv":0},"evokeGold":{"alert":1,"show":true,"cost":0,"vipLv":0},"quickBattle":{"alert":1,"show":true,"cost":10,"vipLv":0},"buyCatapult":{"alert":1,"show":true,"cost":2000,"vipLv":0},"lightGold":{"alert":1,"show":true,"cost":0,"vipLv":0},"cardAllOpenGold":{"alert":1,"show":true,"cost":20,"vipLv":0},"silkEvent":{"alert":1,"show":true,"cost":0,"vipLv":0},"generalRewardBuyFood":{"alert":1,"show":true,"cost":3,"vipLv":0},"generalRewardBuyAllFood":{"alert":1,"show":true,"cost":0,"vipLv":0},"buySilkGold":{"alert":1,"show":true,"cost":0,"vipLv":0},"bloodFarmGold":{"alert":1,"show":true,"cost":1,"vipLv":0},"buyBlackSmithNum":{"alert":1,"show":true,"cost":5,"vipLv":0},"buyGrab":{"alert":1,"show":true,"cost":0,"vipLv":0},"ancientBuyKey":{"alert":1,"show":true,"cost":50,"vipLv":0},"ancientBuyTime2":{"alert":1,"show":true,"cost":100,"vipLv":0},"moonCakeBuyIronWishTimes":{"alert":1,"show":true,"cost":30,"vipLv":0},"mulNationInviteBuy":{"alert":1,"show":true,"cost":20,"vipLv":0},"buyCostReduce":{"alert":1,"show":true,"cost":5,"vipLv":0},"generalOpenBox":{"alert":1,"show":true,"cost":20,"vipLv":0},"featRankRewardCost":{"alert":1,"show":true,"cost":30,"vipLv":0},"beastSummon":{"alert":1,"show":true,"cost":100,"vipLv":0},"ycbwGoldBuyMaterial":{"alert":1,"show":true,"cost":5,"vipLv":0},"goldBuyCoal":{"alert":1,"show":true,"cost":150,"vipLv":0},"ycbwGoldOutput":{"alert":1,"show":true,"cost":100,"vipLv":0},"ycbwGoldHammer":{"alert":1,"show":true,"cost":100,"vipLv":0},"ycbwGoldReduceCd":{"alert":1,"show":true,"cost":100,"vipLv":0},"ycbwGoldBuyProb":{"alert":1,"show":true,"cost":50,"vipLv":0},"buyChariotBlueprint":{"alert":1,"show":true,"cost":200,"vipLv":0},"featGold":{"alert":1,"show":true,"cost":60,"vipLv":0},"godEnchancer":{"alert":1,"show":true,"cost":100,"vipLv":0},"goldReduceCoal":{"alert":1,"show":true,"cost":50,"vipLv":0},"arrowBoatGoldDispatch":{"alert":1,"show":true,"cost":100,"vipLv":0},"buyChangbanHp":{"alert":1,"show":true,"cost":10,"vipLv":0},"buyChangbanWin":{"alert":1,"show":true,"cost":100,"vipLv":0},"xyGate":{"alert":1,"show":true,"cost":75,"vipLv":0},"xyHelper":{"alert":1,"show":true,"cost":10,"vipLv":0},"xySearch":{"alert":1,"show":true,"cost":8,"vipLv":0},"goldBuyStrategicsPosition":{"alert":1,"show":true,"cost":10,"vipLv":0},"treasureSprintBuy":{"alert":1,"show":true,"cost":0,"vipLv":0},"treasureSprintAllBuy":{"alert":1,"show":true,"cost":0,"vipLv":0},"goldBuyJYQJ":{"alert":1,"show":true,"cost":600,"vipLv":0},"goldBuyCLPT":{"alert":1,"show":true,"cost":0,"vipLv":0},"goldBuyXWPW":{"alert":1,"show":true,"cost":50,"vipLv":0},"goldWaterMirror":{"alert":1,"show":true,"cost":0,"vipLv":0},"buyGangCard":{"alert":1,"show":true,"cost":0,"vipLv":0},"gemMineBuyBomb":{"alert":1,"show":true,"cost":0,"vipLv":0},"gemMineAllBomb":{"alert":1,"show":true,"cost":0,"vipLv":0},"meirenactive":{"alert":1,"show":true,"cost":0,"vipLv":0},"meirenupgrade":{"alert":1,"show":true,"cost":0,"vipLv":0},"poemChallenge":{"alert":1,"show":true,"cost":0,"vipLv":0},"poemWakeUp":{"alert":1,"show":true,"cost":0,"vipLv":0},"fstBuyPrisoner":{"alert":1,"show":true,"cost":0,"vipLv":0},"diplomacyBuy":{"alert":1,"show":true,"cost":0,"vipLv":0},"ltcjBuy":{"alert":1,"show":true,"cost":0,"vipLv":0},"chariotRft":{"alert":1,"show":true,"cost":0,"vipLv":0},"buyBattleFlag":{"alert":1,"show":true,"cost":0,"vipLv":0},"quickGoHome":{"alert":1,"show":true,"cost":5,"vipLv":0},"longzhongTicket":{"alert":1,"show":true,"cost":0,"vipLv":0},"longzhongAskHelp":{"alert":1,"show":true,"cost":0,"vipLv":0},"longzhongFixup":{"alert":1,"show":true,"cost":0,"vipLv":0},"longzhongQuick":{"alert":1,"show":true,"cost":0,"vipLv":0},"unionEvokeMeat":{"alert":1,"show":true,"cost":0,"vipLv":0},"goldEpicRoad":{"alert":1,"show":true,"cost":0,"vipLv":0},"goldWestSand":{"alert":1,"show":true,"cost":0,"vipLv":0},"goldEpicKillTime":{"alert":1,"show":true,"cost":0,"vipLv":0},"huanggaiEvokeBuff":{"alert":1,"show":true,"cost":0,"vipLv":0},"lvmengEvokeBuff":{"alert":1,"show":true,"cost":0,"vipLv":0},"whitecTrader":{"alert":1,"show":true,"cost":0,"vipLv":0}},"curTask":{"tasks":[{"type":1,"state":1,"taskId":9,"group":0,"index":0,"taskName":"消灭土匪","introShort":"击败盘踞木场<br>的土匪","introLong":"击败盘踞木场的土匪","processStr":"","requestCompleted":false,"markTrace":"0","iosMarktrace":"0","newTrace":"","areaId":0,"pic":"task1","plot":""}]},"tryTasks":{},"protectTasks":{},"addictionURL":"http://m.766kf.com","redirectUrl":"http://m.766kf.com/server2/denglu.php","battleRewardTimes":0,"freeQuechingTimes":0,"quechingTip":0,"powerId":1,"hasSalary":true,"hasOfficialBuildingOthers":false,"hasOfficerBuildingApply":false,"autoBat":false,"batReward":false,"officialNew":[],"openLegion":0,"is2th":-135596101049,"function":"10001100000000010000010001000000000000000010000000000000100000000000000000011100000000000000000000000000000000000000000000000000","haveDayGift":false,"hasNewMail":false,"marketCanBuyNum":0,"hasPurpleEquip":false,"hasGift":"0","guideId":2,"displayTech":0,"isTrainning":false,"isTrainningOver":false,"totalExp":0,"countryRewardNum":0,"hasVipPrivilege":false,"activity51":false,"hasGoldOrder":false,"antiaddiction":false,"havePayActivity":0,"haveDragonActivity":0,"haveIronActivity":0,"haveTicketActivity":0,"haveDstqActivity":0,"haveSlaveActivity":0,"slaveActivityBuff":0,"haveMidAutumnActivity":0,"haveNationalDayActivity":0,"haveResourceAddittionActivity":0,"haveIronRewardActivity":0,"ironIncenseEffect":0,"weaponEffectCd":0,"haveQuenchingActivity":false,"haveXiLianActivity":0,"haveLoginRewardActivity":1,"haveIronGiveActivity":0,"haveChristmasDayActivity":0,"haveWishActivity":0,"haveBeastActivity":0,"haveBaiNianActivity":0,"haveBaiNianBuff":0,"baiNianBuffCd":0,"haveRedPaperActivity":0,"haveLanternActivity":0,"haveGemMineActivity":true,"haveFeast":0,"changeBat":0,"liShangWangLai":false,"maxNationLv":1,"pin":12,"autoBattleTechGain":false}}}`
	playerinfo := player.GetPlayerInfo{}
	util.Byte2Struct([]byte(s), playerinfo.Rec)
	fmt.Println()
}
func TestJs(t *testing.T) {
	bs, _ := os.ReadFile("../plugins/login_user.js")
	n := otto.New()
	//n.Set("", func(call otto.FunctionCall) otto.Value {
	//	fmt.Printf("Hello, %s.\n", call.Argument(0).String())
	//	fmt.Printf("Hello2, %s.\n", call.Argument(1).String())
	//	result, _ := n.ToValue("123")
	//	return result
	//})
	n.Run(bs)

	v, _ := n.Call("Plugin.Receive", nil, "123", "456")

	println(v.String())
}
