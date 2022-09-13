var Plugin = {
    Name: function () {
        return "重连"
    },
    Command: function () {
        return "reconnect"
    },
    Receive: function (type, data) {
        console.log(type)
        console.log(data)
        return 0
    }
};
