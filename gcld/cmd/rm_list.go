package cmd

type CMD string

const (
	LoginUser CMD = "login_user"
	Reconnect CMD = "reconnect"

	BuildingGetMainCityInfo CMD = "building@getMainCityInfo"
	BuildingGetBuildingInfo CMD = "building@getBuildingInfo"

	PushPlayer      CMD = "push@player"
	PushChat        CMD = "push@chat"
	PushBuilding    CMD = "push@building"
	PushNotice      CMD = "push@notice"
	PushRightNotice CMD = "push@rightNotice"
	PushWorldReward CMD = "push@worldReward"

	ChariotForgeSp   CMD = "chariot@forgeSp"
	ChariotGetBpInfo CMD = "chariot@getBpInfo"
	ChariotforgeBp   CMD = "chariot@forgeBp"
	ChariotGet       CMD = "chariot@getInfo"

	PlayerGetPlayerInfo CMD = "player@getPlayerInfo"
	PlayerGetPlayerList CMD = "player@getPlayerList"
	PlayerGame          CMD = "player@game"

	TaskFinishTask CMD = "task@finishTask"

	BattleBattlePermit    CMD = "battle@battlePermit"
	BattleBattlePrepare   CMD = "battle@battlePrepare"
	BattleGetPowerInfo    CMD = "battle@getPowerInfo"
	BattleSwitchPowerInfo CMD = "battle@switchPowerInfo"
)
