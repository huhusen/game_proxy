var Plugin = {
    Name: function () {
        return "【推送】右边通知"
    },
    Command: function () {
        return "push@rightNotice"
    },
    Receive: function (type, data) {
        console.log(type)
        console.log(data)
        return 0
    }
};
