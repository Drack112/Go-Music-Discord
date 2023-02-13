package bot

import (
    "github.com/bwmarrin/discordgo"
)

func (bot *Bot) Avatar(interact *discordgo.InteractionCreate, messageChannel chan<- *discordgo.MessageEmbed) {

    avatarUrl := interact.Member.AvatarURL("4096")
    avatarUsername := interact.Member.User.Username

    messageChannel <- &discordgo.MessageEmbed{
        Title:       "🖼️ " + "Click here to download the Pic",
        Description: "It's really great! Nice.",
        URL:         avatarUrl,
        Author: &discordgo.MessageEmbedAuthor{
            Name:    avatarUsername + " 📸",
            IconURL: avatarUrl,
        },
        Image: &discordgo.MessageEmbedImage{
            URL: avatarUrl,
        },
        Color: 0xfff8f7,
    }
}
