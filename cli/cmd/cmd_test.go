package cmd

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	fmt.Println("------TEST of Registering------")
	registCmd.Flags().Set("username", "user")
	registCmd.Flags().Set("password", "123")
	registCmd.Flags().Set("email", "111@mail.com")
	registCmd.Flags().Set("telephone", "12345678901")
	registCmd.Run(registCmd, nil)
}

func TestLogin(t *testing.T) {
	fmt.Println("------TEST of Login------")
	logoutCmd.Run(logoutCmd, nil)
	loginCmd.Flags().Set("username", "user")
	loginCmd.Flags().Set("password", "123")
	loginCmd.Run(loginCmd, nil)
}

func TestLogout(t *testing.T) {
	fmt.Println("------TEST of Logout------")
	loginCmd.Flags().Set("username", "user")
	loginCmd.Flags().Set("password", "123")
	loginCmd.Run(loginCmd, nil)
	logoutCmd.Run(logoutCmd, nil)
}

func TestListAllUsers(t *testing.T) {
	fmt.Println("------TEST of Listing all users------")
	lsuCmd.Run(lsuCmd, nil)
}

func TestCreateNewMeeting(t *testing.T) {
	fmt.Println("------TEST of Creating a new meeting------")
	cmCmd.Flags().Set("title", "testMeeting")
	cmCmd.Flags().Set("participant", "testUser0,testUser1")
	cmCmd.Flags().Set("start", "2017-10-20T19:00")
	cmCmd.Flags().Set("end", "22017-10-20T20:00")
	cmCmd.Run(cmCmd, nil)
}

func TestListAllMeetings(t *testing.T) {
	fmt.Println("------TEST of Listing all meetings------")
	lsmCmd.Run(lsmCmd, nil)
}

func TestClearMeetings(t *testing.T) {
	fmt.Println("------TEST of clearing all meetings------")
	clearCmd.Run(clearCmd, nil)
}
