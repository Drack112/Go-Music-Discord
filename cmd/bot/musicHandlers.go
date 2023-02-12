package bot

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/bwmarrin/discordgo"

    player "github.com/Drack112/Go-Music-Bot/cmd/youtube/player"
    request "github.com/Drack112/Go-Music-Bot/cmd/youtube/request"
)

func (bot *Bot) startPlayer(interact *discordgo.InteractionCreate, channel *discordgo.VoiceState) {
    musicPlayer := bot.musicPlayers[interact.GuildID]
    musicPlayer.Started = true
    var err error
    musicPlayer.VoiceConnection, err = bot.ChannelVoiceJoin(channel.GuildID, channel.ChannelID, false, false)
    if err != nil {
        bot.ChannelMessageSend(interact.ChannelID, "Error while joining voice channel: "+err.Error())
    } else {
        musicPlayer.Start()
    }
    musicPlayer.Disconnect()
    musicPlayer.Started = false
    delete(bot.musicPlayers, interact.GuildID)
}

func (bot *Bot) play(interact *discordgo.InteractionCreate, messageChannel chan<- string) {

    var query string = interact.ApplicationCommandData().Options[0].StringValue()
    var nowPlaying chan bool = make(chan bool)

    invokingMemberChannel, err := bot.State.VoiceState(interact.GuildID, interact.Member.User.ID)
    if err != nil {
        message := "You are not currently joined to a voice channel! Please join a voice channel to play music."
        messageChannel <- message
        return
    }

    req, err := request.New(query, nowPlaying)
    if err != nil {
        message := "Could not add request to queue: " + err.Error()
        messageChannel <- message
        return
    }
    message := interact.Member.User.Username + " requested: [`" + req.Title + "`](" + req.ReqURL + ")"
    messageChannel <- message

    bot.mu.Lock()
    defer bot.mu.Unlock()
    if bot.musicPlayers[interact.GuildID] == nil {
        bot.musicPlayers[interact.GuildID] = player.New()
    }
    musicPlayer := bot.musicPlayers[interact.GuildID]
    musicPlayer.AddToQueue(req)
    go func() {
        nowPlaying := <-nowPlaying
        if nowPlaying {
            bot.ChannelMessageSend(interact.ChannelID, "**Now Playing:** `"+req.Title+"`")
        } else {
            bot.ChannelMessageSend(interact.ChannelID, "**Error Playing:** `"+req.Title+"`; *skipping song*")
        }
    }()
    bot.ChannelMessageSend(interact.ChannelID, "*Added to Queue:* `"+req.Title+"`")
    if !musicPlayer.Started {
        go bot.startPlayer(interact, invokingMemberChannel)
    }

}

func (bot *Bot) stop(interact *discordgo.InteractionCreate, messageChannel chan<- string) {

    if bot.musicPlayers[interact.GuildID] == nil {
        message := "I'm not playing any music right now!"
        messageChannel <- message
        return
    }
    musicPlayer := bot.musicPlayers[interact.GuildID]

    musicPlayer.Stop <- true
    message := "Music stopped"
    messageChannel <- message
}

func (bot *Bot) skip(interact *discordgo.InteractionCreate, messageChannel chan<- string) {

    if bot.musicPlayers[interact.GuildID] == nil {
        message := "I'm not playing any music right now!"
        messageChannel <- message
        return
    }
    musicPlayer := bot.musicPlayers[interact.GuildID]

    musicPlayer.Next <- true
    message := "Skipped song"
    messageChannel <- message
}

func (bot *Bot) queue(interact *discordgo.InteractionCreate, messageChannel chan<- string) {

    if bot.musicPlayers[interact.GuildID] == nil {
        message := "I'm not playing any music right now!"
        messageChannel <- message
        return
    }
    musicPlayer := bot.musicPlayers[interact.GuildID]

    go func() {

        var builder strings.Builder

        builder.WriteString("`1.` **`" + musicPlayer.CurrentSong.Title + "`** - *Now Playing*\n")
        for i := 0; i < len(musicPlayer.Queue); i++ {
            builder.WriteString("`" + strconv.Itoa(i+2) + ".` `" + musicPlayer.Queue[i].Title + "`\n")
        }
        _, err := bot.ChannelMessageSend(interact.ChannelID, builder.String())
        if err != nil {
            fmt.Println("Error sending channel message: ", err)
        }

    }()

    message := "__Song Queue__"
    messageChannel <- message
}
