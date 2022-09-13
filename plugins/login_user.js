var Plugin = {
    Name: function () {
        return "【登录】用户登录"
    },
    Command: function () {
        return "login_user"
    },
    Receive: function (type, data) {
        console.log(type)
        console.log(data)
        return 0
    }
};
