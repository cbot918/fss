package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type User struct {
	Username string
	Password string
}

var users []User

func init() {
	users = append(users, User{"yale", "12345"})
	users = append(users, User{"node", "12345"})
}

func main() {
	var err error
	var welcomeMessage string
	var input string
	var username string
	var password string

	reader := bufio.NewReader(os.Stdin)

	welcomeMessage += "welcome to fss"
	fmt.Println(welcomeMessage)
	for {

		fmt.Print("username: ")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		username = input[:len(input)-1]

		fmt.Print("password: ")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		password = input[:len(input)-1]

		userExists := false
		for _, user := range users {
			if username == user.Username && password != user.Password {
				fmt.Println("password incorrect")
				break
			} else if username == user.Username && password == user.Password {
				userExists = true
			}
		}

		if !userExists {
			fmt.Println("username not exists, create with > fss signup")
			break
		}
	}

	GameService()

}

func GameService() {
	fmt.Println("in middle of 8c ...")
}
