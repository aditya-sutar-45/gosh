package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var dirStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	MarginRight(1)

var usernameStyle = lipgloss.NewStyle().
	Bold(false).
	Foreground(lipgloss.Color("#9c7ef7"))

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		currentUser, err := user.Current()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// for vanity purposes
		// fmt.Printf("%s\n %s > ", dir, currentUser.Username)
		usernameString := fmt.Sprintf("%s > ", currentUser.Username)

		usernameRender := usernameStyle.Render(usernameString)
		dirRender := dirStyle.Render(dir)

		fmt.Println(dirRender + usernameRender)

		// read input from keyboard
		inp, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = execInp(inp)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

var ErrNoPath = errors.New("path required")

func execInp(inp string) error {
	// removing newline
	inp = strings.TrimSuffix(inp, "\n")
	// getting the args
	args := strings.Split(inp, " ")

	command := args[0]
	switch command {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])
	case "ls":
		command = "lsd"
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(command, args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
