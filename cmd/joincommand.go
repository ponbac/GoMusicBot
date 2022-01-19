package cmd

import (
	"github.com/ponbac/GoMusicBot/framework"
)

func JoinCommand(ctx framework.Context) *framework.Session {
	if ctx.Sessions.GetByGuild(ctx.Guild.ID) != nil {
		ctx.Reply("Already connected! Use `!leave` for the bot to disconnect.")
		return nil
	}
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
	ctx.Reply("Joined <#" + sess.ChannelId + ">!")
	return sess
}
