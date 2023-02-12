package bot

import (
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"

    player "github.com/Drack112/Go-Music-Bot/cmd/youtube/player"
    "github.com/bwmarrin/discordgo"
)

type Bot struct {
    *discordgo.Session

    mu           sync.Mutex
    musicPlayers map[string]*player.MusicPlayer
}

func New(token string) *Bot {
    session, err := discordgo.New("Bot " + token)
    if err != nil {
        fmt.Println("Error creating Discord session: ", err)
        return nil
    }
    return &Bot{session, sync.Mutex{}, make(map[string]*player.MusicPlayer)}
}

func (bot *Bot) Run(resetCommands bool) {
    if resetCommands {
        bot.AddHandler(bot.onReadyReset)
    } else {
        bot.AddHandler(bot.onReady)
    }
    bot.AddHandler(bot.interactionHandler)
    bot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
    err := bot.Open()
    if err != nil {
        fmt.Println("Cant open a discord session: ", err)
        return
    }
    defer bot.Close()

    fmt.Println(" 🎧🙇 Music Bot is now running")
    fmt.Println("Press CTRL-C to exit.")
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
    <-signals
}
