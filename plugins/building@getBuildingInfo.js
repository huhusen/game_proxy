var Plugin = {
    Name: function () {
        return "【建筑】获取建筑信息"
    },
    Command: function () {
        return "building@getBuildingInfo"
    },
    Receive: function (type, data) {
        console.log(type)
        console.log(data)
        return 0
    }
};
