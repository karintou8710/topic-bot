package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	stopBot = make(chan bool)
	topics  = []string{
		"お菓子作り",
		"今日の天気",
		"出身地",
		"好きな本",
		"アルバイト",
		"お金",
		"趣味",
		"旅行",
		"温泉",
		"自作PC",
		"マイエディター",
		"vim VS nano",
		"野球 VS サッカー",
		"恋愛",
		"登山",
		"雪遊び",
		"スマブラ",
		"マイクラ",
		"インターン",
	}
	CLIENT_ID = ""
	TOKEN     = ""
)

func main() {
	rand.Seed(time.Now().UnixNano()) // シード値の設定
	loadEnv()                        // 環境変数の初期化

	discord, err := discordgo.New()
	checkErr(err)
	discord.Token = TOKEN

	discord.AddHandler(onMessageCreate)

	err = discord.Open()
	checkErr(err)
	defer discord.Close()

	fmt.Println("Listening...")
	<-stopBot

}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// BOT自身なら無視する
	if m.Author.ID == CLIENT_ID {
		return
	}

	if m.Message.Content == "!topic" {
		index := rand.Intn(len(topics))
		message := fmt.Sprintf("話題： %s", topics[index])
		_, err := s.ChannelMessageSend(m.ChannelID, message) // 指定のチャンネルに送信
		checkErr(err)
		return
	}
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	TOKEN = os.Getenv("TOKEN")
	CLIENT_ID = os.Getenv("CLIENT_ID")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
