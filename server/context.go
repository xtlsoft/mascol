package server

import (
	"github.com/xtlsoft/mascol/channel"
)

// Context is the server context
type Context struct {
	Channels map[string]*channel.Channel
}

// GetChannel returns a channel
func (ctx *Context) GetChannel(name string) *channel.Channel {
	val, ok := ctx.Channels[name]
	if ok {
		return val
	}
	ch := new(channel.Channel)
	ctx.Channels[name] = ch
	return ch
}
