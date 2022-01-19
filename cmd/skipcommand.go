package cmd

import (
	"github.com/ponbac/GoMusicBot/framework"
)

func SkipCommand(ctx framework.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("Not in a voice channel! To make the bot join one, use `!join`.")
		return
	}
	sess.Stop()
	ctx.Reply("Skipped song!")
}
