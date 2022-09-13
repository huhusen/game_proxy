var Plugin = {
    Name: function () {
        return "【推送】聊天"
    },
    Command: function () {
        return "push@chat"
    },
    Receive: function (type, data) {
        console.log(type)
        console.log(data)
        return 0
    }
};
