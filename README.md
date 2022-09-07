## Sockets5代理

> 一般用于游戏的中间人操作

##### **使用方法**
```go
    //构建代理
    s := proxy.NewServer(1888)
    s.OnSocket5RequestEvent = func(message []byte) (out []byte) {
        log.Log.Println("send", string(message))
        if strings.Index(string(message), "123") != -1 {
            out = []byte("hack Send 123\n")
        }
        return
    }
    s.OnSocket5ResponseEvent = func(message []byte) (out []byte) {
        log.Log.Println("recv", string(message))
        if strings.Index(string(message), "123") != -1 {
            out = []byte("hack recv 123\n")
        }
        return
    }
	
	//启动代理
    go s.Start()

    for {
		//这个包注入了客户端对象，可以用来自定义包发送到客户端or服务端
        for k, proto := range s.Clients {
            r, ok := proto.(*proxy.Socket5)
            if ok {
                r.Send2Server([]byte("I am socket5 Client"))
                r.Send2client([]byte("I am socket5 Server"))
            }
            fmt.Println(k)
    
        }
        time.Sleep(time.Second * 10)
    }
```
TODO：

**利用它开发一款页游的脱机项目.敬请期待！**