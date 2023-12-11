package main

import (
	"fmt"
	"github.com/creack/pty"
	"github.com/gliderlabs/ssh"
	ssh2 "golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
	"unsafe"
)

func ReadPrivateKeyFromFile(path string) (ssh.Signer, error) {
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	key, err := ssh2.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

func main() {
	s := &ssh.Server{
		Addr:            "127.0.0.1:2222",
		Handler:         handleSession,
		PasswordHandler: handleAuthentication,
		IdleTimeout:     60 * time.Second,
	}

	key, err := ReadPrivateKeyFromFile("/home/user/.ssh/id_rsa")
	if err != nil {
		panic(err)
	}
	s.AddHostKey(key)

	log.Printf("[+]Starting SSH server on address: %v\n", s.Addr)

	log.Fatal(s.ListenAndServe())
}

// Called if a new ssh session was created
func handleSession(s ssh.Session) {
	cmd := exec.Command("bash")
	ptyReq, winCh, isPty := s.Pty()
	if isPty {
		cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
		f, err := pty.Start(cmd)
		if err != nil {
			panic(err)
		}
		go func() {
			for win := range winCh {
				setWinsize(f, win.Width, win.Height)
			}
		}()
		go func() {
			io.Copy(f, s) // stdin
		}()
		io.Copy(s, f) // stdout
		cmd.Wait()
	} else {
		io.WriteString(s, "No PTY requested.\n")
		s.Exit(1)
	}
}

// Return true to accept password and false to deny
func handleAuthentication(ctx ssh.Context, passwd string) bool {

	if ctx.User() != "Artyom" || passwd != "somethingGreat" {
		// Deny
		return false
	}

	fmt.Printf("User: %s,Password: %s, Address: %s \n", ctx.User(), passwd, ctx.RemoteAddr().String())

	// Accept
	return true
}
