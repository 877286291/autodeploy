package utils

import (
	"bytes"
	"fmt"
	"github.com/aurora/autodeploy/pkg/setting"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

func Connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	keyboardInteractiveChallenge := func(
		user,
		instruction string,
		questions []string,
		echos []bool,
	) (answers []string, err error) {
		if len(questions) == 0 {
			return []string{}, nil
		}
		return []string{password}, nil
	}
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	auth = append(auth, ssh.KeyboardInteractive(keyboardInteractiveChallenge))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

// FTP上执行shell命令
func FtpExecShell(command string) {
	session, err := Connect("root", "P1cc@xfzy", setting.FtpHost, 22)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	_ = session.Run(command)
}

// 分区上执行shell命令
func LparExecShell(ip, command string) (stdoutBuf bytes.Buffer, err error) {
	var session *ssh.Session
	session, err = Connect("root", "P1cc@xfzy", ip, 22)
	if err != nil {
		return
	}
	defer session.Close()
	session.Stdout = &stdoutBuf
	err = session.Run(command)
	return
}
