package gcld

type CMD string

const (
	LoginUser           CMD = "login_user"
	BuildingGetMainCity CMD = "building@getMainCityInfo"
	PushPlayer          CMD = "push@player"
	PushChat            CMD = "push@chat"
	ChariotForgeSp      CMD = "chariot@forgeSp"
	ChariotGetBp        CMD = "chariot@getBpInfo"
	ChariotforgeBp      CMD = "chariot@forgeBp"
	ChariotGet          CMD = "chariot@getInfo"
)

func ToCMD(c CMD) (zh string) {
	switch c {
	case LoginUser:
		zh = "登录"
	case BuildingGetMainCity:
		zh = "获取建筑主城信息"
	case PushPlayer:
		zh = "广播"
	case PushChat:
		zh = "聊天"
	case ChariotForgeSp:
		zh = "战车升级"
	case ChariotGetBp:
		zh = "获取战车部件信息"
	case ChariotforgeBp:
		zh = "战车部件改造"
	case ChariotGet:
		zh = "获取战车信息"
	default:
		zh = string(c)
	}
	return
}
