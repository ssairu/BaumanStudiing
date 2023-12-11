package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"strings"
	"time"
)

type SSHCommand struct {
	Path   string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

type SSHClient struct {
	Config *ssh.ClientConfig
	Host   string
	Port   int
}

func (client *SSHClient) RunCommand(cmd *SSHCommand) error {
	var (
		session *ssh.Session
		err     error
	)

	if session, err = client.newSession(); err != nil {
		return err
	}
	defer session.Close()

	if err = client.prepareCommand(session, cmd); err != nil {
		return err
	}

	err = session.Run(cmd.Path)
	return err
}

func (client *SSHClient) prepareCommand(session *ssh.Session, cmd *SSHCommand) error {
	if cmd.Stdin != nil {
		stdin, err := session.StdinPipe()
		if err != nil {
			return fmt.Errorf("Unable to setup stdin for session: %v", err)
		}
		go io.Copy(stdin, cmd.Stdin)
	}

	if cmd.Stdout != nil {
		stdout, err := session.StdoutPipe()
		if err != nil {
			return fmt.Errorf("Unable to setup stdout for session: %v", err)
		}
		go io.Copy(cmd.Stdout, stdout)
	}

	if cmd.Stderr != nil {
		stderr, err := session.StderrPipe()
		if err != nil {
			return fmt.Errorf("Unable to setup stderr for session: %v", err)
		}
		go io.Copy(cmd.Stderr, stderr)
	}

	return nil
}

func (client *SSHClient) newSession() (*ssh.Session, error) {
	connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port), client.Config)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s", err)
	}

	session, err := connection.NewSession()
	if err != nil {
		return nil, fmt.Errorf("Failed to create session: %s", err)
	}

	modes := ssh.TerminalModes{
		// ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 100, modes); err != nil {
		session.Close()
		return nil, fmt.Errorf("request for pseudo terminal failed: %s", err)
	}

	return session, nil
}

func main() {
	sshConfig := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            "Artyom",
		Auth: []ssh.AuthMethod{
			ssh.Password("somethingGreat"),
		},
	}

	client := &SSHClient{
		Config: sshConfig,
		Host:   "127.0.0.1",
		Port:   2222,
	}

	startDir := ""
	in := bufio.NewReader(os.Stdin)
	for {
		time.Sleep(1000)
		fmt.Println("input command")

		command, _ := in.ReadString('\n')
		command = command[:len(command)-1]

		if len(startDir) != 0 && startDir[0] == '/' {
			startDir = startDir[1:]
		}

		commandPATH := command
		if len(startDir) != 0 {
			commandPATH = "cd " + startDir + " && " + command
		}

		words := strings.Fields(command)
		for i, word := range words {
			if "cd" == word {
				if i+1 == len(words) {
					startDir = ""
				} else if words[i+1] == "." {
					continue
				} else if words[i+1] == ".." {
					for j := len(startDir) - 1; len(startDir) != 0 && startDir[j] != '/'; j-- {
						startDir = startDir[:j]
					}
					if len(startDir) != 0 && startDir[len(startDir)-1] == '/' {
						startDir = startDir[:len(startDir)-1]
					}
				} else {
					startDir += "/" + words[i+1]
				}
			}
		}

		cmd := &SSHCommand{
			Path:   commandPATH,
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		fmt.Printf("Running command: %s\n", cmd.Path)
		if err := client.RunCommand(cmd); err != nil {
			fmt.Fprintf(os.Stderr, "command run error: %s\n", err)
			os.Exit(1)
		}
	}
}

