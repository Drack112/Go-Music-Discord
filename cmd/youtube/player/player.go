package youtube

import (
    "fmt"

    request "github.com/Drack112/Go-Music-Bot/cmd/youtube/request"
    stream "github.com/Drack112/Go-Music-Bot/cmd/youtube/stream"
    "github.com/bwmarrin/discordgo"
)

type MusicPlayer struct {
    *discordgo.VoiceConnection
    Queue       []*request.Request
    CurrentSong *request.Request
    Started     bool
    Stop        chan bool
    Next        chan bool
}

func New() *MusicPlayer {
    return &MusicPlayer{&discordgo.VoiceConnection{}, make([]*request.Request, 0, 24), nil, false, make(chan bool), make(chan bool)}
}

// Main player loop
func (player *MusicPlayer) Start() {
    err := player.Speaking(true)
    if err != nil {
        fmt.Println("Couldn't set speaking: ", err)
        return
    }
    for player.Started && !player.isEmpty() {
        player.CurrentSong = player.nextSong()
        player.CurrentSong.NowPlaying()
        streamUrl, err := player.CurrentSong.GetStream()
        if err != nil {
            player.CurrentSong.Cancel()
            continue
        }
        player.play(streamUrl)
    }
    err = player.Speaking(false)
    if err != nil {
        fmt.Println("Couldn't stop speaking: ", err)
    }
}

func (player *MusicPlayer) play(url string) {
    song := stream.New(url)
    go song.Get()
    for {
        select {
        case opusBytes, ok := <-song.Audio:
            if !ok {
                return
            }
            player.OpusSend <- opusBytes
        case <-player.Next:
            return
        case <-player.Stop:
            player.Started = false
            return
        }
    }
}

func (player *MusicPlayer) AddToQueue(request *request.Request) {
    player.Queue = append(player.Queue, request)
}

func (player *MusicPlayer) nextSong() *request.Request {
    request := player.Queue[0]
    player.Queue = player.Queue[1:]
    return request
}

func (player *MusicPlayer) isEmpty() bool {
    return len(player.Queue) < 1
}
