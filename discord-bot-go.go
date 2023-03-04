package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "syscall"

    "github.com/bwmarrin/discordgo"
)

func main() {
    // Create a new Discord session using the bot token.
    dg, err := discordgo.New("Bot " + "YOUR_BOT_TOKEN")
    if err != nil {
        fmt.Println("Error creating Discord session: ", err)
        return
    }

    // Load custom commands from file.
    commands, err := loadCommandsFromFile("commands.txt")
    if err != nil {
        fmt.Println("Error loading custom commands: ", err)
    }

    // Register the messageCreate func as a callback for the messageCreate events.
    dg.AddHandler(messageCreate)

    // Open a websocket connection to Discord and begin listening.
    err = dg.Open()
    if err != nil {
        fmt.Println("Error opening Discord session: ", err)
        return
    }

    // Wait here until CTRL-C or other term signal is received.
    fmt.Println("Bot is now running. Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Cleanly close down the Discord session.
    dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Ignore messages sent by the bot itself.
    if m.Author.ID == s.State.User.ID {
        return
    }

    // If the message is "ping", respond with "Pong!" and the response time.
    if m.Content == "ping" {
        response := "Pong!"
        s.ChannelMessageSend(m.ChannelID, response)
        return
    }

    // Check if the message matches a custom command and respond.
    for command, response := range commands {
        if m.Content == command {
            s.ChannelMessageSend(m.ChannelID, response)
            return
        }
    }

    // If the message starts with "weather ", get the weather information and respond.
    if len(m.Content) > 8 && m.Content[:8] == "weather " {
        location := m.Content[8:]
        weather, err := getWeather(location)
        if err != nil {
            s.ChannelMessageSend(m.ChannelID, "Error getting weather information.")
            return
        }
        response := fmt.Sprintf("The weather in %s is %s with a temperature of %s", location, weather.Condition.Text, weather.Condition.Temp)
        s.ChannelMessageSend(m.ChannelID, response)
        return
    }
}

type Weather struct {
    Condition struct {
        Text string `json:"text"`
        Temp string `json:"temperature"`
    } `json:"condition"`
}

func getWeather(location string) (*Weather, error) {
    url := fmt.Sprintf("https://wttr.in/%s?format=%%C{weather}\\n%%C{temperature}", location)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    weather := &Weather{}
    err = json.NewDecoder(resp.Body).Decode(weather)
    if err != nil {
        return nil, err
    }
    return weather, nil
}

func loadCommandsFromFile(filename string) (map[string]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer
