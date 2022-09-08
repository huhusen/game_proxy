package gcld

type CMD string

const (
	LoginUser               CMD = "login_user"
	BuildingGetMainCityInfo CMD = "building@getMainCityInfo"
	PushPlayer              CMD = "push@player"
	PushChat                CMD = "push@chat"
)
