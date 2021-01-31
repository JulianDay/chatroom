1.代码结构
--cmd
	--chatroom		聊天室server可执行程序
	--client		聊天室client可执行程序
--config			服务器配置
--dfa				屏蔽字dfa算法
--logic				聊天室逻辑
	--lastmsg.go	历史聊天记录
	--message.go 	消息结构定义
	--messagedispather.go 消息分发器
	--player.go		玩家
	--popular.go	流行词统计
	--room.go		聊天室
	--sensitive.go  屏蔽字
--server
	--客户端服务器交互消息
--utls
	--公共函数
	
2.lastmsg历史消息
	使用ring.Ring环形链表保存最近50条历史消息
	
3.popular 流行词统计
	使用{秒,{word:count}}的结构统计玩家输入的单词,在get/add的会触发删除过期的word_count(一秒内不会触发两次)

4.屏蔽字
	使用dfa算法替换屏蔽字

5.引用
	引用了websocket解决客户端和服务器之间通信
	require nhooyr.io/websocket v1.8.6

6.运行
	启动服务器
	cd cmd/chatroom
	chatroom.exe
	cd cmd/client
	client.exe
	
7.单元测试
	go test -v sensitive_test.go sensitive.go
	=== RUN   TestFilterSensitive
	--- PASS: TestFilterSensitive (0.00s)
	PASS
	ok      command-line-arguments  0.485s
	
	go test -v popular_test.go popular.go
	=== RUN   TestPopularWords
	--- PASS: TestPopularWords (7.00s)
    popular_test.go:14: [1] a:3
        [2] b:2
        [3] c:1

    popular_test.go:18: [1] a:4
        [2] b:3
        [3] c:1

    popular_test.go:23: [1] a:4
        [2] b:3
        [3] c:2

    popular_test.go:23: [1] a:4
        [2] b:3
        [3] c:3

    popular_test.go:23: [1] a:4
        [2] c:4
        [3] b:3

    popular_test.go:23: [1] c:5
        [2] a:4
        [3] b:3

    popular_test.go:23: [1] c:5
        [2] a:1
        [3] b:1

	PASS
	ok      command-line-arguments  7.341s