var Plugin = {
    Name: function () {
        return "【玩家】游戏信息"
    },
    Command: function () {
        return "player@game"
    },
    Receive: function (type, data) {
        console.log(type)
        console.log(data)
        return 0
    }
};
