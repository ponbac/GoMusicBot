package cmd

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ponbac/GoMusicBot/framework"
	"github.com/robertkrimen/otto"
)

var vm = otto.New()

func EvalCommand(ctx framework.Context) {
	if ctx.Conf.OwnerId != ctx.User.ID {
		return
	}
	js := strings.Join(ctx.Args, " ")
	vm.Set("ctx", ctx)
	vm.Set("getGuild", Guild)
	vm.Set("getEmojis", Emojis)
	vm.Set("format", Format)
	val, err := vm.Run(js)
	if err != nil {
		ctx.Reply(err.Error())
		return
	}
	if val.IsNull() {
		return
	}
	ctx.Reply("`" + val.String() + "`")
}

func Format(input string, entities []interface{}) string {
	return fmt.Sprintf(input, entities...)
}

func Guild(ctx framework.Context, id string) *discordgo.Guild {
	guild, err := ctx.Discord.State.Guild(id)
	if err != nil {
		fmt.Println("Error getting guild,", err)
		return nil
	}
	return guild
}

func Emojis(ctx framework.Context, guildId string) []string {
	guild := Guild(ctx, guildId)
	arr := make([]string, 0)
	for _, emoj := range guild.Emojis {
		arr = append(arr, "<:"+emoj.Name+":"+emoj.ID+">")
	}
	return arr
}
