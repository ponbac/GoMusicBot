package cmd

import "github.com/ponbac/GoMusicBot/framework"

func CurrentCommand(ctx framework.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("Not in a voice channel! To make the bot join one, use `!join`.")
		return
	}
	current := sess.Queue.Current()
	if current == nil {
		ctx.Reply("The song queue is empty! Add a song with `!add`.")
		return
	}
	ctx.Reply("Currently playing `" + current.Title + "`.")
}
