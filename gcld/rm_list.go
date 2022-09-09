package gcld

type CMD string

const (
	LoginUser               CMD = "login_user"
	BuildingGetMainCityInfo CMD = "building@getMainCityInfo"
	PushPlayer              CMD = "push@player"
	PushChat                CMD = "push@chat"
)

func ToCMD(c CMD) (zh string) {
	switch c {
	case LoginUser:
		zh = "登录"
	case BuildingGetMainCityInfo:
		zh = "建筑"
	case PushPlayer:
		zh = "广播"
	case PushChat:
		zh = "聊天"
	default:
		zh = string(c)
	}
	return
}
