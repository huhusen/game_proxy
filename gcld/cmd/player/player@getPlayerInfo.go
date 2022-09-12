package player

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"sockets-proxy/gcld/cmd"
)

type GetPlayerInfo struct {
	cmd.Command
	Rec2 struct {
		BlackNames []interface{} `json:"blackNames"`
		EndTime    int           `json:"endTime"`
		Player     struct {
			PlayerID       int    `json:"playerId"`
			PlayerLv       int    `json:"playerLv"`
			PlayerName     string `json:"playerName"`
			UserID         string `json:"userId"`
			Yx             string `json:"yx"`
			ForceID        int    `json:"forceId"`
			Pic            int    `json:"pic"`
			Gold           int    `json:"gold"`
			VipLv          int    `json:"vipLv"`
			Pkey           string `json:"pkey"`
			Pkey2          string `json:"pkey2"`
			Copper         int    `json:"copper"`
			CopperMax      int    `json:"copperMax"`
			CopperOutput   int    `json:"copperOutput"`
			Food           int64  `json:"food"`
			FoodMax        int    `json:"foodMax"`
			FoodOutput     int    `json:"foodOutput"`
			Wood           int    `json:"wood"`
			WoodMax        int    `json:"woodMax"`
			WoodOutput     int    `json:"woodOutput"`
			Iron           int    `json:"iron"`
			IronMax        int    `json:"ironMax"`
			IronOutput     int    `json:"ironOutput"`
			Exp            int    `json:"exp"`
			ExpNeed        int    `json:"expNeed"`
			Forces         int    `json:"forces"`
			ForcesMax      int    `json:"forcesMax"`
			InPveBattle    bool   `json:"inPveBattle"`
			InOccupyBattle bool   `json:"inOccupyBattle"`
			ServerTime     string `json:"serverTime"`
		} `json:"player"`
		SsTime      int64 `json:"ssTime"`
		ChargeItems struct {
			BuildingCd struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buildingCd"`
			RefreshGeneralCd struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"refreshGeneralCd"`
			StoreBuyItem struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"storeBuyItem"`
			RefreshStoreCd struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"refreshStoreCd"`
			BuyStorehouseSize struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyStorehouseSize"`
			TechCd struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"techCd"`
			MoveCd struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"moveCd"`
			GoldRecruit struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"goldRecruit"`
			ResourceMode1 struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"resourceMode1"`
			ResourceMode2 struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"resourceMode2"`
			ResourceMode3 struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"resourceMode3"`
			PoliticsEvent struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"politicsEvent"`
			LoadGem struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"loadGem"`
			UnloadGem struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"unloadGem"`
			Shaodaozi struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"shaodaozi"`
			Nverhong struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"nverhong"`
			Zhuyeqing struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"zhuyeqing"`
			IncenseGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"incenseGold"`
			BlackMarketGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"blackMarketGold"`
			NationalRankBatBuyNumGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"nationalRankBatBuyNumGold"`
			BuyBonusArmyGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyBonusArmyGold"`
			QuenchingGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"quenchingGold"`
			SlashSlave struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"slashSlave"`
			SlaveFreedom struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"slaveFreedom"`
			BuyWeaponItem struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyWeaponItem"`
			BuyPowerExtraGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyPowerExtraGold"`
			FreeConsGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"freeConsGold"`
			ConsDrawingGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"consDrawingGold"`
			KfDqGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"kfDqGold"`
			YoudiGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"youdiGold"`
			JiebingGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"jiebingGold"`
			ChujiGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"chujiGold"`
			CreateBattleTeam struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"createBattleTeam"`
			CopperInvest struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"copperInvest"`
			BattleTeamInspire struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"battleTeamInspire"`
			CoverManzuShoumaiCd struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"coverManzuShoumaiCd"`
			TeamOrder struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"teamOrder"`
			UpdateLash struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"updateLash"`
			UseGoldOrder struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"useGoldOrder"`
			KfzbBuySup struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"kfzbBuySup"`
			RecoverGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"recoverGold"`
			Irongive struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"irongive"`
			GemDamo struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"gemDamo"`
			GemJinglian struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"gemJinglian"`
			BuyCrack struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyCrack"`
			FarmQuickFinish struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"farmQuickFinish"`
			BuyGoldDice struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyGoldDice"`
			AncientBuyTime struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"ancientBuyTime"`
			DiamondShopComplementGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"diamondShopComplementGold"`
			EvokeGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"evokeGold"`
			QuickBattle struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"quickBattle"`
			BuyCatapult struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyCatapult"`
			LightGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"lightGold"`
			CardAllOpenGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"cardAllOpenGold"`
			SilkEvent struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"silkEvent"`
			GeneralRewardBuyFood struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"generalRewardBuyFood"`
			GeneralRewardBuyAllFood struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"generalRewardBuyAllFood"`
			BuySilkGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buySilkGold"`
			BloodFarmGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"bloodFarmGold"`
			BuyBlackSmithNum struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyBlackSmithNum"`
			BuyGrab struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyGrab"`
			AncientBuyKey struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"ancientBuyKey"`
			AncientBuyTime2 struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"ancientBuyTime2"`
			MoonCakeBuyIronWishTimes struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"moonCakeBuyIronWishTimes"`
			MulNationInviteBuy struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"mulNationInviteBuy"`
			BuyCostReduce struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyCostReduce"`
			GeneralOpenBox struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"generalOpenBox"`
			FeatRankRewardCost struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"featRankRewardCost"`
			BeastSummon struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"beastSummon"`
			YcbwGoldBuyMaterial struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"ycbwGoldBuyMaterial"`
			GoldBuyCoal struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"goldBuyCoal"`
			YcbwGoldOutput struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"ycbwGoldOutput"`
			YcbwGoldHammer struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"ycbwGoldHammer"`
			YcbwGoldReduceCd struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"ycbwGoldReduceCd"`
			YcbwGoldBuyProb struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"ycbwGoldBuyProb"`
			BuyChariotBlueprint struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyChariotBlueprint"`
			FeatGold struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"featGold"`
			GodEnchancer struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"godEnchancer"`
			GoldReduceCoal struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"goldReduceCoal"`
			ArrowBoatGoldDispatch struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"arrowBoatGoldDispatch"`
			BuyChangbanHp struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyChangbanHp"`
			BuyChangbanWin struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyChangbanWin"`
			XyGate struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"xyGate"`
			XyHelper struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"xyHelper"`
			XySearch struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"xySearch"`
			GoldBuyStrategicsPosition struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"goldBuyStrategicsPosition"`
			TreasureSprintBuy struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"treasureSprintBuy"`
			TreasureSprintAllBuy struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"treasureSprintAllBuy"`
			GoldBuyJYQJ struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"goldBuyJYQJ"`
			GoldBuyCLPT struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"goldBuyCLPT"`
			GoldBuyXWPW struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"goldBuyXWPW"`
			GoldWaterMirror struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"goldWaterMirror"`
			BuyGangCard struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyGangCard"`
			GemMineBuyBomb struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"gemMineBuyBomb"`
			GemMineAllBomb struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"gemMineAllBomb"`
			Meirenactive struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"meirenactive"`
			Meirenupgrade struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"meirenupgrade"`
			PoemChallenge struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"poemChallenge"`
			PoemWakeUp struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"poemWakeUp"`
			FstBuyPrisoner struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"fstBuyPrisoner"`
			DiplomacyBuy struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"diplomacyBuy"`
			LtcjBuy struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"ltcjBuy"`
			ChariotRft struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"chariotRft"`
			BuyBattleFlag struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"buyBattleFlag"`
			QuickGoHome struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"quickGoHome"`
			LongzhongTicket struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"longzhongTicket"`
			LongzhongAskHelp struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"longzhongAskHelp"`
			LongzhongFixup struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"longzhongFixup"`
			LongzhongQuick struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"longzhongQuick"`
			UnionEvokeMeat struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"unionEvokeMeat"`
			GoldEpicRoad struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"goldEpicRoad"`
			GoldWestSand struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"goldWestSand"`
			GoldEpicKillTime struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"goldEpicKillTime"`
			HuanggaiEvokeBuff struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"huanggaiEvokeBuff"`
			LvmengEvokeBuff struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"lvmengEvokeBuff"`
			WhitecTrader struct {
				Alert int  `json:"alert"`
				Show  bool `json:"show"`
				Cost  int  `json:"cost"`
				VipLv int  `json:"vipLv"`
			} `json:"whitecTrader"`
		} `json:"chargeItems"`
		CurTask struct {
			Tasks []struct {
				Type             int    `json:"type"`
				State            int    `json:"state"`
				TaskID           int    `json:"taskId"`
				Group            int    `json:"group"`
				Index            int    `json:"index"`
				TaskName         string `json:"taskName"`
				IntroShort       string `json:"introShort"`
				IntroLong        string `json:"introLong"`
				ProcessStr       string `json:"processStr"`
				RequestCompleted bool   `json:"requestCompleted"`
				MarkTrace        string `json:"markTrace"`
				IosMarktrace     string `json:"iosMarktrace"`
				NewTrace         string `json:"newTrace"`
				AreaID           int    `json:"areaId"`
				Pic              string `json:"pic"`
				Plot             string `json:"plot"`
			} `json:"tasks"`
		} `json:"curTask"`
		TryTasks struct {
		} `json:"tryTasks"`
		ProtectTasks struct {
		} `json:"protectTasks"`
		AddictionURL                  string        `json:"addictionURL"`
		RedirectURL                   string        `json:"redirectUrl"`
		BattleRewardTimes             int           `json:"battleRewardTimes"`
		FreeQuechingTimes             int           `json:"freeQuechingTimes"`
		QuechingTip                   int           `json:"quechingTip"`
		PowerID                       int           `json:"powerId"`
		HasSalary                     bool          `json:"hasSalary"`
		HasOfficialBuildingOthers     bool          `json:"hasOfficialBuildingOthers"`
		HasOfficerBuildingApply       bool          `json:"hasOfficerBuildingApply"`
		AutoBat                       bool          `json:"autoBat"`
		BatReward                     bool          `json:"batReward"`
		OfficialNew                   []interface{} `json:"officialNew"`
		OpenLegion                    int           `json:"openLegion"`
		Is2Th                         int64         `json:"is2th"`
		Function                      string        `json:"function"`
		HaveDayGift                   bool          `json:"haveDayGift"`
		HasNewMail                    bool          `json:"hasNewMail"`
		MarketCanBuyNum               int           `json:"marketCanBuyNum"`
		HasPurpleEquip                bool          `json:"hasPurpleEquip"`
		HasGift                       string        `json:"hasGift"`
		GuideID                       int           `json:"guideId"`
		DisplayTech                   int           `json:"displayTech"`
		IsTrainning                   bool          `json:"isTrainning"`
		IsTrainningOver               bool          `json:"isTrainningOver"`
		TotalExp                      int           `json:"totalExp"`
		CountryRewardNum              int           `json:"countryRewardNum"`
		HasVipPrivilege               bool          `json:"hasVipPrivilege"`
		Activity51                    bool          `json:"activity51"`
		HasGoldOrder                  bool          `json:"hasGoldOrder"`
		Antiaddiction                 bool          `json:"antiaddiction"`
		HavePayActivity               int           `json:"havePayActivity"`
		HaveDragonActivity            int           `json:"haveDragonActivity"`
		HaveIronActivity              int           `json:"haveIronActivity"`
		HaveTicketActivity            int           `json:"haveTicketActivity"`
		HaveDstqActivity              int           `json:"haveDstqActivity"`
		HaveSlaveActivity             int           `json:"haveSlaveActivity"`
		SlaveActivityBuff             int           `json:"slaveActivityBuff"`
		HaveMidAutumnActivity         int           `json:"haveMidAutumnActivity"`
		HaveNationalDayActivity       int           `json:"haveNationalDayActivity"`
		HaveResourceAddittionActivity int           `json:"haveResourceAddittionActivity"`
		HaveIronRewardActivity        int           `json:"haveIronRewardActivity"`
		IronIncenseEffect             int           `json:"ironIncenseEffect"`
		WeaponEffectCd                int           `json:"weaponEffectCd"`
		HaveQuenchingActivity         bool          `json:"haveQuenchingActivity"`
		HaveXiLianActivity            int           `json:"haveXiLianActivity"`
		HaveLoginRewardActivity       int           `json:"haveLoginRewardActivity"`
		HaveIronGiveActivity          int           `json:"haveIronGiveActivity"`
		HaveChristmasDayActivity      int           `json:"haveChristmasDayActivity"`
		HaveWishActivity              int           `json:"haveWishActivity"`
		HaveBeastActivity             int           `json:"haveBeastActivity"`
		HaveBaiNianActivity           int           `json:"haveBaiNianActivity"`
		HaveBaiNianBuff               int           `json:"haveBaiNianBuff"`
		BaiNianBuffCd                 int           `json:"baiNianBuffCd"`
		HaveRedPaperActivity          int           `json:"haveRedPaperActivity"`
		HaveLanternActivity           int           `json:"haveLanternActivity"`
		HaveGemMineActivity           bool          `json:"haveGemMineActivity"`
		HaveFeast                     int           `json:"haveFeast"`
		ChangeBat                     int           `json:"changeBat"`
		LiShangWangLai                bool          `json:"liShangWangLai"`
		MaxNationLv                   int           `json:"maxNationLv"`
		Pin                           int           `json:"pin"`
		AutoBattleTechGain            bool          `json:"autoBattleTechGain"`
	}
}

func NewGetPlayerInfo() *GetPlayerInfo {
	p := &GetPlayerInfo{}
	p.Zh = "【玩家】获取用户信息"
	p.Cmd = cmd.PlayerGetPlayerInfo
	p.Rec = &p.Rec2
	return p
}
func (c *GetPlayerInfo) Update() {
	mapstructure.Decode(c.Map(), &c.Rec2)
	fmt.Println()
}
