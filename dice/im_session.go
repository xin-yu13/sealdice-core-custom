package dice

import (
	"encoding/json"
	"fmt"
	"github.com/fy0/procs"
	"github.com/sacOO7/gowebsocket"
	"math/rand"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"
)

type Sender struct {
	Age      int32  `json:"age"`
	Card     string `json:"card"`
	Nickname string `json:"nickname"`
	Role     string `json:"owner"`
	UserId   int64  `json:"user_id"`
}

type Message struct {
	MessageId     int64  `json:"message_id"`
	MessageType   string `json:"message_type"` // Group
	Sender        Sender `json:"sender"`       // 发送者
	RawMessage    string `json:"raw_message"`
	Message       string `json:"message"` // 消息内容
	Time          int64  `json:"time"`    // 发送时间
	MetaEventType string `json:"meta_event_type"`
	GroupId       int64  `json:"group_id"`     // 群号
	PostType      string `json:"post_type"`    // 上报类型，如group、notice
	RequestType   string `json:"request_type"` // 请求类型，如group
	SubType       string `json:"sub_type"`     // 子类型，如add invite
	Flag          string `json:"flag"`         // 请求 flag, 在调用处理请求的 API 时需要传入
	NoticeType    string `json:"notice_type"`
	UserId        int64  `json:"user_id"`
	SelfId        int64  `json:"self_id"`
	Duration      int64  `json:"duration"`

	Data *struct {
		// 个人信息
		Nickname string `json:"nickname"`
		UserId   int64  `json:"user_id"`

		// 群信息
		GroupId         int64  `json:"group_id"`          // 群号
		GroupCreateTime uint32 `json:"group_create_time"` // 群号
		MemberCount     int64  `json:"member_count"`
		GroupName       string `json:"group_name"`
	} `json:"data"`
	Retcode int64 `json:"retcode"`
	//Status string `json:"status"`
	Echo int `json:"echo"`
}

type PlayerInfo struct {
	UserId int64  `yaml:"userId"`
	Name   string // 玩家昵称
	//ValueNumMap    map[string]int64  `yaml:"valueNumMap"`
	//ValueStrMap    map[string]string `yaml:"valueStrMap"`
	RpToday        int    `yaml:"rpToday"`
	RpTime         string `yaml:"rpTime"`
	LastUpdateTime int64  `yaml:"lastUpdateTime"`
	InGroup        bool   `yaml:"inGroup"`

	// level int 权限
	DiceSideNum    int                  `yaml:"diceSideNum"` // 面数，为0时等同于d100
	TempValueAlias *map[string][]string `yaml:"-"`

	ValueMap     map[string]VMValue `yaml:"-"`
	ValueMapTemp map[string]VMValue `yaml:"-"`
}

type ServiceAtItem struct {
	Active           bool                  `json:"active" yaml:"active"` // 需要能记住配置，故有此选项
	ActivatedExtList []*ExtInfo            `yaml:"activatedExtList"`     // 当前群开启的扩展列表
	Players          map[int64]*PlayerInfo // 群员信息

	LogCurName string `yaml:"logCurFile"`
	LogOn      bool   `yaml:"logOn"`
	GroupId    int64  `yaml:"groupId"`
	GroupName  string `yaml:"groupName"`

	ValueMap     map[string]VMValue `yaml:"-"`
	CocRuleIndex int                `yaml:"cocRuleIndex"`

	// http://www.antagonistes.com/files/CoC%20CheatSheet.pdf
	//RuleCriticalSuccessValue *int64 // 大成功值，1默认
	//RuleFumbleValue *int64 // 大失败值 96默认
}

type PlayerVariablesItem struct {
	Loaded              bool               `yaml:"-"`
	ValueMap            map[string]VMValue `yaml:"-"`
	LastSyncToCloudTime int64              `yaml:"lastSyncToCloudTime"`
	LastUsedTime        int64              `yaml:"lastUsedTime"`
}

type ConnectInfoItem struct {
	//Password            string              `yaml:"password" json:"password"`
	Socket              *gowebsocket.Socket `yaml:"-" json:"-"`
	Id                  string              `yaml:"id" json:"id"` // uuid
	Nickname            string              `yaml:"nickname" json:"nickname"`
	State               int                 `yaml:"state" json:"state"` // 状态 0 断开 1已连接
	UserId              int64               `yaml:"userId" json:"userId"`
	GroupNum            int64               `yaml:"groupNum" json:"groupNum"`                       // 拥有群数
	CmdExecutedNum      int64               `yaml:"cmdExecutedNum" json:"cmdExecutedNum"`           // 指令执行次数
	CmdExecutedLastTime int64               `yaml:"cmdExecutedLastTime" json:"cmdExecutedLastTime"` // 指令执行次数
	OnlineTotalTime     int64               `yaml:"onlineTotalTime" json:"onlineTotalTime"`         // 在线时长
	ConnectUrl          string              `yaml:"connectUrl" json:"connectUrl"`                   // 连接地址

	Platform          string `yaml:"platform" json:"platform"`                   // 平台，如QQ等
	RelWorkDir        string `yaml:"relWorkDir" json:"relWorkDir"`               // 工作目录
	Enable            bool   `yaml:"enable" json:"enable"`                       // 是否启用
	Type              string `yaml:"type" json:"type"`                           // 协议类型，如onebot、koishi等
	UseInPackGoCqhttp bool   `yaml:"useInPackGoCqhttp" json:"useInPackGoCqhttp"` // 是否使用内置的gocqhttp

	InPackGoCqHttpProcess            *procs.Process `yaml:"-" json:"-"`
	InPackGoCqHttpLoginSuccess       bool           `yaml:"-" json:"inPackGoCqHttpLoginSuccess"`   // 是否登录成功
	InPackGoCqHttpLoginSucceeded     bool           `yaml:"inPackGoCqHttpLoginSucceeded" json:"-"` // 是否登录成功过
	InPackGoCqHttpRunning            bool           `yaml:"-" json:"inPackGoCqHttpRunning"`        // 是否仍在运行
	InPackGoCqHttpQrcodeReady        bool           `yaml:"-" json:"inPackGoCqHttpQrcodeReady"`    // 二维码已就绪
	InPackGoCqHttpNeedQrCode         bool           `yaml:"-" json:"inPackGoCqHttpNeedQrCode"`     // 是否需要二维码
	InPackGoCqHttpQrcodeData         []byte         `yaml:"-" json:"-"`                            // 二维码数据
	InPackGoCqHttpLoginDeviceLockUrl string         `yaml:"-" json:"inPackGoCqHttpLoginDeviceLockUrl"`
	DiceServing                      bool           `yaml:"-"`
}

type IMSession struct {
	Conns          []*ConnectInfoItem             `yaml:"connections"`
	Parent         *Dice                          `yaml:"-"`
	ServiceAt      map[int64]*ServiceAtItem       `json:"serviceAt" yaml:"serviceAt"`
	PlayerVarsData map[int64]*PlayerVariablesItem `yaml:"PlayerVarsData"`
	//CommandIndex int64                    `yaml:"-"`
	//GroupId int64 `json:"group_id"`
}

type MsgContext struct {
	MessageType     string
	Session         *IMSession
	Group           *ServiceAtItem
	Player          *PlayerInfo
	Dice            *Dice
	IsCurGroupBotOn bool
	Socket          *gowebsocket.Socket
	conn            *ConnectInfoItem
}

func (s *IMSession) Serve(index int) int {
	log := s.Parent.Logger
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	disconnected := make(chan int, 1)

	conn := s.Conns[index]
	session := s

	socket := gowebsocket.New(conn.ConnectUrl)
	conn.Socket = &socket

	socket.OnConnected = func(socket gowebsocket.Socket) {
		fmt.Println("onebot 连接成功")
		log.Info("onebot 连接成功")
		//  {"data":{"nickname":"闃斧鐗岃�佽檸鏈�","user_id":1001},"retcode":0,"status":"ok"}
		conn.GetLoginInfo()
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Info("Recieved connect error: ", err)
		fmt.Println("连接失败")
		disconnected <- 2
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		msg := new(Message)
		err := json.Unmarshal([]byte(message), msg)

		if err == nil {
			// 心跳包，忽略
			if msg.MetaEventType == "heartbeat" {
				return
			}
			if msg.MetaEventType == "heartbeat" {
				return
			}

			// 获得用户信息
			if msg.Echo == -1 {
				conn.UserId = msg.Data.UserId
				conn.Nickname = msg.Data.Nickname

				log.Debug("骰子信息已刷新")
				return
			}

			// 获得群信息
			if msg.Echo == -2 {
				if msg.Data != nil {
					group := session.ServiceAt[msg.Data.GroupId]
					if group != nil {
						group.GroupName = msg.Data.GroupName
						group.GroupId = msg.Data.GroupId
					}
					log.Debug("群信息刷新: ", msg.Data.GroupName)
				}
				return
			}

			// 处理加群请求
			if msg.PostType == "request" && msg.RequestType == "group" && msg.SubType == "invite" {
				// {"comment":"","flag":"111","group_id":222,"post_type":"request","request_type":"group","self_id":333,"sub_type":"invite","time":1646782195,"user_id":444}
				log.Infof("收到加群邀请: 群组(%d) 邀请人:%d", msg.GroupId, msg.UserId)
				time.Sleep(time.Duration((0.8 + rand.Float64()) * float64(time.Second)))
				SetGroupAddRequest(conn.Socket, msg.Flag, msg.SubType, true, "")
				return
			}

			// 入群后自动开启
			if msg.PostType == "notice" && msg.NoticeType == "group_increase" {
				//{"group_id":111,"notice_type":"group_increase","operator_id":0,"post_type":"notice","self_id":333,"sub_type":"approve","time":1646782012,"user_id":333}
				if msg.UserId == msg.SelfId {
					// 判断进群的人是自己，自动启动
					SetBotOnAtGroup(session, msg)
					// fmt.Sprintf("<%s>已经就绪。可通过.help查看指令列表", conn.Nickname)
					ctx := &MsgContext{Session: session, Dice: session.Parent, Socket: conn.Socket}
					replyGroupRaw(ctx, msg.GroupId, DiceFormatTmpl(ctx, "核心:骰子进群"), "")
					log.Infof("加入群组: (%d)", msg.GroupId)
				}
				return
			}

			if msg.PostType == "notice" && msg.NoticeType == "group_decrease" && msg.SubType == "kick_me" {
				// 被踢
				//  {"group_id":111,"notice_type":"group_decrease","operator_id":222,"post_type":"notice","self_id":333,"sub_type":"kick_me","time":1646689414 ,"user_id":333}
				if msg.UserId == msg.SelfId {
					log.Infof("被踢出群: 在群组(%d)中被踢出，操作者:(%d)", msg.GroupId, msg.UserId)
				}
				return
			}

			if msg.PostType == "notice" && msg.NoticeType == "group_ban" && msg.SubType == "ban" {
				// 禁言
				// {"duration":600,"group_id":111,"notice_type":"group_ban","operator_id":222,"post_type":"notice","self_id":333,"sub_type":"ban","time":1646689567,"user_id":333}
				if msg.UserId == msg.SelfId {
					log.Infof("被禁言: 在群组(%d)中被禁言，时长%d秒，操作者:(%d)", msg.GroupId, msg.Duration, msg.UserId)
				}
				return
			}

			// 处理命令
			if msg.MessageType == "group" || msg.MessageType == "private" {
				mctx := &MsgContext{}
				mctx.Dice = session.Parent
				mctx.MessageType = msg.MessageType
				mctx.Session = session
				mctx.Socket = conn.Socket
				mctx.conn = conn
				var cmdLst []string

				// 兼容模式检查
				if s.Parent.CommandCompatibleMode {
					for k := range session.Parent.CmdMap {
						cmdLst = append(cmdLst, k)
					}

					sa := session.ServiceAt[msg.GroupId]
					if sa != nil && sa.Active {
						for _, i := range sa.ActivatedExtList {
							for k := range i.CmdMap {
								cmdLst = append(cmdLst, k)
							}
						}
					}
					sort.Sort(ByLength(cmdLst))
				}

				// 收到信息回调
				sa := session.ServiceAt[msg.GroupId]
				mctx.Group = sa
				mctx.Player = GetPlayerInfoBySender(session, msg)
				mctx.IsCurGroupBotOn = IsCurGroupBotOn(session, msg)

				// 收到群 test(1111) 内 XX(222) 的消息: 好看 (1232611291)
				if msg.MessageType == "group" {
					log.Infof("收到群(%d)内<%s>(%d)的消息: %s", msg.GroupId, msg.Sender.Nickname, msg.Sender.UserId, msg.Message)
				}
				if msg.MessageType == "private" {
					log.Infof("收到<%s>(%d)的私聊消息: %s", msg.Sender.Nickname, msg.Sender.UserId, msg.Message)
				}

				if sa != nil && sa.Active {
					for _, i := range sa.ActivatedExtList {
						if i.OnMessageReceived != nil {
							i.OnMessageReceived(mctx, msg)
						}
					}
				}

				msgInfo := CommandParse(msg.Message, s.Parent.CommandCompatibleMode, cmdLst)

				if msgInfo != nil {
					f := func() {
						defer func() {
							if r := recover(); r != nil {
								//  + fmt.Sprintf("%s", r)
								log.Error(r)
								ReplyToSender(mctx, msg, DiceFormatTmpl(mctx, "核心:骰子崩溃"))
							}
						}()
						session.commandSolve(mctx, msg, msgInfo)
					}
					go f()
					//session.commandSolve(mctx, msg, msgInfo)
				} else {
					//text := fmt.Sprintf("信息 来自群%d - %s(%d)：%s", msg.GroupId, msg.Sender.Nickname, msg.Sender.UserId, msg.Message);
					//replyGroup(Socket, 22, text)
				}
				//}
				//}
			} else {
				fmt.Println("Recieved message " + message)
			}
		} else {
			log.Error("error" + err.Error())
		}
	}

	socket.OnBinaryMessage = func(data []byte, socket gowebsocket.Socket) {
		log.Debug("Recieved binary data ", data)
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		log.Debug("Recieved ping " + data)
	}

	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		log.Debug("Recieved pong " + data)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Info("Disconnected from server ")
		disconnected <- 1
	}

	socket.Connect()
	defer func() {
		fmt.Println("socket close")
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("关闭连接时遭遇异常")
					//core.GetLogger().Error(r)
				}
			}()

			// 可能耗时好久
			socket.Close()
		}()
	}()

	for {
		select {
		case <-interrupt:
			log.Info("interrupt")
			disconnected <- 0
			return 0
		case val := <-disconnected:
			return val
		}
	}
}

func (s *IMSession) commandSolve(ctx *MsgContext, msg *Message, cmdArgs *CmdArgs) {
	// 设置AmIBeMentioned
	cmdArgs.AmIBeMentioned = false
	for _, i := range cmdArgs.At {
		if i.UserId == ctx.conn.UserId {
			cmdArgs.AmIBeMentioned = true
			break
		}
	}

	// 设置临时变量
	if ctx.Player != nil {
		VarSetValue(ctx, "$t玩家", &VMValue{VMTypeString, fmt.Sprintf("<%s>", ctx.Player.Name)})
		VarSetValue(ctx, "$tQQ昵称", &VMValue{VMTypeString, fmt.Sprintf("<%s>", msg.Sender.Nickname)})
		VarSetValue(ctx, "$t个人骰子面数", &VMValue{VMTypeInt64, ctx.Player.DiceSideNum})
		VarSetValue(ctx, "$tQQ", &VMValue{VMTypeInt64, msg.Sender.UserId})
		VarSetValue(ctx, "$t骰子帐号", &VMValue{VMTypeInt64, ctx.conn.UserId})
		VarSetValue(ctx, "$t骰子昵称", &VMValue{VMTypeInt64, ctx.conn.Nickname})
		// 注: 未来将私聊视为空群吧
	}

	tryItemSolve := func(item *CmdItemInfo) bool {
		if item != nil {
			ret := item.Solve(ctx, msg, cmdArgs)
			if ret.Success {
				return true
			}
		}
		return false
	}

	sa := ctx.Group
	if sa != nil && sa.Active {
		for _, i := range sa.ActivatedExtList {
			if i.OnCommandReceived != nil {
				i.OnCommandReceived(ctx, msg, cmdArgs)
			}
		}
	}

	item := ctx.Session.Parent.CmdMap[cmdArgs.Command]
	if tryItemSolve(item) {
		return
	}

	if sa != nil && sa.Active {
		for _, i := range sa.ActivatedExtList {
			item := i.CmdMap[cmdArgs.Command]
			if tryItemSolve(item) {
				return
			}
		}
	}

	if msg.MessageType == "private" {
		for _, i := range ctx.Dice.ExtList {
			if i.ActiveOnPrivate {
				item := i.CmdMap[cmdArgs.Command]
				if tryItemSolve(item) {
					return
				}
			}
		}
	}
}
