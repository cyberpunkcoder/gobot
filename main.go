package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	//"github.com/bwmarrin/discordgo"
)

func init() {
	err := godotenv.Load()

	//Check if
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println(os.Getenv("DISCORDTOKEN"))
}

func main() {

}
