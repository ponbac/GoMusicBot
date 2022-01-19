package cmd

import (
	"github.com/ponbac/GoMusicBot/framework"
	"fmt"
)

func PlayCommand(ctx framework.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		sess = JoinCommand(ctx)
		if sess == nil {
			return
		}
	}
	AddCommand(ctx)
	if sess != nil {
		fmt.Println("sess is not nil")
		queue := sess.Queue
		if !queue.HasNext() {
			ctx.Reply("Queue is empty! Add songs with `!add`.")
			return
		}
		go queue.Start(sess, func(msg string) {
			ctx.Reply(msg)
		})
	} else {
		fmt.Println("sess is nil")
	}
}
