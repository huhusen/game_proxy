package cmd

type CMD string

const (
	LoginUser CMD = "login_user"

	BuildingGetMainCityInfo CMD = "building@getMainCityInfo"

	PushPlayer   CMD = "push@player"
	PushChat     CMD = "push@chat"
	PushBuilding CMD = "push@building"
	PushNotice   CMD = "push@notice"

	ChariotForgeSp   CMD = "chariot@forgeSp"
	ChariotGetBpInfo CMD = "chariot@getBpInfo"
	ChariotforgeBp   CMD = "chariot@forgeBp"
	ChariotGet       CMD = "chariot@getInfo"

	PlayerGetPlayerInfo CMD = "player@getPlayerInfo"
	PlayerGetPlayerList CMD = "player@getPlayerList"
)
