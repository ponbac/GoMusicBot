package cmd

import (
	"fmt"
	"github.com/ponbac/GoMusicBot/framework"
)

func PlayCommand(ctx framework.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		sess = JoinVoiceChannel(ctx)
		if sess == nil {
			ctx.Reply("An error occured when joining voice channel!")
			return
		}
	}
	AddCommand(ctx)
	if sess != nil {
		queue := sess.Queue
		if !queue.HasNext() {
			ctx.Reply("Queue is empty! Add songs with `!add`.")
			return
		}
		go queue.Start(sess, func(msg string) {
			//ctx.Reply(msg)
			if msg == "Finished queue." {
				ctx.Discord.UpdateGameStatus(0, "absolutely nothing 😞")
			} else {
				ctx.Discord.UpdateListeningStatus("💿 " + msg + " 💿")
			}
		})
	} else {
		fmt.Println("PlayCommand: sess is nil")
	}
}

func JoinVoiceChannel(ctx framework.Context) *framework.Session {
	vc := ctx.GetVoiceChannel()
	if vc == nil {
		ctx.Reply("You must be in a voice channel to use the bot!")
		return nil
	}
	sess, err := ctx.Sessions.Join(ctx.Discord, ctx.Guild.ID, vc.ID, framework.JoinProperties{
		Muted:    false,
		Deafened: true,
	})
	if err != nil {
		ctx.Reply("An error occured!")
		return nil
	}
	return sess
}
