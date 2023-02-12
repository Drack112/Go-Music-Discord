package bot

import (
    "fmt"
    "time"

    "github.com/bwmarrin/discordgo"
)

func (bot *Bot) onReady(session *discordgo.Session, event *discordgo.Ready) {
    bot.createCommands()
}

func (bot *Bot) onReadyReset(session *discordgo.Session, event *discordgo.Ready) {
    bot.deleteCommands()
    bot.createCommands()
}

func (bot *Bot) createCommands() {
    commands := []*discordgo.ApplicationCommand{
        {
            Name:        "play",
            Description: "Play a song",
            Options: []*discordgo.ApplicationCommandOption{
                {
                    Type:        discordgo.ApplicationCommandOptionString,
                    Name:        "query",
                    Description: "Youtube search query, Youtube video ID, or URL to song",
                    Required:    true,
                },
            },
        },
        {
            Name:        "stop",
            Description: "Stop playing music and disconnect",
        },
        {
            Name:        "skip",
            Description: "Skip the current song in queue",
        },
        {
            Name:        "queue",
            Description: "List all songs in queue",
        },
    }
    for _, command := range commands {
        _, err := bot.ApplicationCommandCreate(bot.State.User.ID, "", command)
        if err != nil {
            fmt.Println("Error while registering commands: ", err)
        }
    }
}

func (bot *Bot) interactionHandler(session *discordgo.Session, i *discordgo.InteractionCreate) {
    switch i.Type {
    case discordgo.InteractionApplicationCommand:
        bot.commandHandler(i)
    }
}

func (bot *Bot) commandHandler(interact *discordgo.InteractionCreate) {
    err := bot.InteractionRespond(interact.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
    })
    if err != nil {
        fmt.Println("Error while sending interaction response: ", err)
        return
    }
    if bot.LogLevel > 1 {
        now := time.Now()
        fmt.Println(now, interact.ApplicationCommandData().Name, "command used by", interact.Member.User.Username)
    }
    messageChannel := make(chan string)
    switch interact.ApplicationCommandData().Name {
    case "play":
        go bot.play(interact, messageChannel)
    case "stop":
        go bot.stop(interact, messageChannel)
    case "skip":
        go bot.skip(interact, messageChannel)
    case "queue":
        go bot.queue(interact, messageChannel)
    }
    message := <-messageChannel
    _, err = bot.InteractionResponseEdit(interact.Interaction, &discordgo.WebhookEdit{
        Content: &message,
    })

    if err != nil {
        fmt.Println("Error while updating interaction response: ", err)
    }
}

func (bot *Bot) deleteCommands() {
    commands, err := bot.ApplicationCommands(bot.State.User.ID, "")
    if err != nil {
        fmt.Println("Error while getting commands: ", err)
        return
    }
    for _, command := range commands {
        err = bot.ApplicationCommandDelete(bot.State.User.ID, "", command.ID)
        if err != nil {
            fmt.Println("Error while clearing commands: ", err)
        }
    }
}
