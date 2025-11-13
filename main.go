package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	lipgloss_styles "github.com/aditya-sutar-45/gosh/lipgloss_styles"
)

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
		lipgloss_styles.RenderHeader(dir, currentUser.Username)

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
