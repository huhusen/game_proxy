var Plugin = {
    Name: function () {
        return "【推送】建筑信息"
    },
    Command: function () {
        return "push@building"
    },
    Receive: function (type, data) {
        console.log(type)
        console.log(data)
        return 0
    }
};
