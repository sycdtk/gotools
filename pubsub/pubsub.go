package pubsub

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

//WaitGroup封装结构
type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(fn func(argvs ...interface{}), argvs ...interface{}) {
	w.Add(1)
	go func() {
		fn(argvs...)
		w.Done()
	}()
}

//订阅通道
type Channel struct {
	Name      string           //通道名称
	Clients   *sync.Map        //订阅客户端集合
	WaitGroup WaitGroupWrapper //WaitGroup封装结构
	exitFlag  int32            //退出标志
}

//添加通道的订阅客户端
func (c *Channel) AddClient(client *Client) {
	c.Clients.Store(client.Id, client)
}

//移除通道的订阅客户端
func (c *Channel) RemoveClient(client *Client) {
	c.Clients.Delete(client.Id)
}

//向通道客户端发送消息
func (c *Channel) Notify(message string) bool {

	c.Clients.Range(func(key, value interface{}) bool {
		c.notifyMsg(value.(*Client), message)
		return true
	})
	return true
}

//通道上注册的客户端数量
func (c *Channel) Size() int64 {
	size := int64(0)
	c.Clients.Range(func(key, value interface{}) bool {
		size++
		return true
	})
	return size
}

//通道销毁
func (c *Channel) Exit() {
	if atomic.CompareAndSwapInt32(&c.exitFlag, 0, 1) {
		c.WaitGroup.Wait()
	}
}

//channel触发消息
func (c *Channel) notifyMsg(client *Client, message string) {
	c.WaitGroup.Wrap(func(argvs ...interface{}) {
		fmt.Println(argvs[0].(*Client).Id, argvs[1].(string))
	}, client, message)
}

//订阅客户端
type Client struct {
	Id string
	Ip string
}

//发布服务端
type Server struct {
	Channels *sync.Map //channel名称(topic)/channel
}

//构建发布服务
func NewServer() *Server {
	return &Server{Channels: &sync.Map{}}
}

//订阅操作
func (s *Server) Subscribe(client *Client, channelName string) {
	if v, ok := s.Channels.Load(channelName); ok {
		v.(*Channel).AddClient(client)
	} else {
		newChannel := &Channel{Name: channelName, Clients: &sync.Map{}}
		newChannel.AddClient(client)
		s.Channels.Store(channelName, newChannel)
	}
}

//取消订阅
func (s *Server) Unsubscribe(client *Client, channelName string) {
	if v, ok := s.Channels.Load(channelName); ok {
		channel := v.(*Channel)
		channel.RemoveClient(client) //移除订阅客户端
		if channel.Size() == 0 {     //销毁通道
			channel.Exit()
			s.Channels.Delete(channelName)
		}
	}
}

//发布消息
func (s *Server) PublishMessage(channelName, message string) (bool, error) {
	if v, ok := s.Channels.Load(channelName); ok {
		channel := v.(*Channel)
		channel.Notify(message)
		channel.WaitGroup.Wait()
	} else {
		return false, errors.New("channelName不存在!")
	}
	return true, nil
}
