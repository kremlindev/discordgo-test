package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	prefix = "go!"
	token  = "ODcwNjA2MzM1Mjg3Mzg2MTIy.GjL-9E.-2Hu8W7xAAnUv1OWCaHxPiN5YyPKst8BrRjwpI"
)

func main() {
	fmt.Println("Proje Başarıyla başlatıldı.")
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.AddHandler(messageCreate)
	client.Identify.Intents = discordgo.IntentsGuildMessages
	err = client.Open()
	if err != nil {
		fmt.Println("Başlatılırken hata oluştu: ", err)
		return
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	client.Close()
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	switch strings.ToLower(m.Content) {
	case "sa":
		s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+">, Aleykümselam. Hoşgeldin.")
	case prefix + "react":
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       "Tepki Ekleme İşlemi",
			Description: "**Başarılı**\nBaşarıyla Mesajınıza 📀 tepkisi eklendi!",
			Color:       12430073,
		})
		s.MessageReactionAdd(m.ChannelID, m.Reference().MessageID, "📀")
	}
}
