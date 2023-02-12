package main

import (
    "flag"
    "fmt"
    "log"

    "github.com/Drack112/Go-Music-Bot/cmd/bot"
    "github.com/Drack112/Go-Music-Bot/config"
    "github.com/joho/godotenv"
)

func main() {

    defaultTokenValue := ""

    err := godotenv.Load()

    if err != nil {
        log.Fatal("Error loading .env file")
    }

    config, err := config.Load()

    if err != nil {
        log.Fatal("Error loading config file")
    }

    cfg := config
    defaultTokenValue = cfg.Token

    tFlag := flag.String("t", defaultTokenValue, "Discord API Token")
    rFlag := flag.Bool("r", false, "Reset all bot commands")
    lFlag := flag.String("l", "ERROR", "Logging level")
    flag.Parse()

    token := *tFlag
    resetCommands := *rFlag
    logLevel := *lFlag

    if token == "" {
        fmt.Println("No token provided. Please set DISCORD_TOKEN environment variable, or use '-t' option to set your Discord API token.")
        return
    }

    houseband := bot.New(token)

    switch logLevel {
    case "ERROR":
        houseband.LogLevel = 0
    case "WARN":
        houseband.LogLevel = 1
    case "INFO":
        houseband.LogLevel = 2
    case "DEBUG":
        houseband.LogLevel = 3
    default:
        fmt.Println("Unknown LogLevel. Please set LogLevel to ERROR, WARN, INFO, or DEBUG.")
        return
    }

    houseband.Run(resetCommands)
}
