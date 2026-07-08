package collector

import (
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

var (
	clientMu  sync.Mutex
	client    *ssh.Client
	sshTarget string
	sshConfig *ssh.ClientConfig
)

// SetSSHConfig stores the SSH connection parameters for later use.
func SetSSHConfig(target, user, pass string) {
	sshTarget = target
	sshConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(pass)},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func connect() (*ssh.Client, error) {
	c, err := ssh.Dial("tcp", sshTarget, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("ssh dial: %w", err)
	}
	clientMu.Lock()
	if client != nil {
		client.Close()
	}
	client = c
	clientMu.Unlock()
	return c, nil
}

// ExecCommand runs a command on the router via SSH, with automatic reconnection on failure.
func ExecCommand(command string) (string, error) {
	return execCommandWithRetry(command, true)
}

// GetContent reads a file from the router via SSH (equivalent to "cat <path>").
func GetContent(path string) (string, error) {
	return ExecCommand("cat " + path)
}

func execCommandWithRetry(command string, allowRetry bool) (string, error) {
	clientMu.Lock()
	c := client
	clientMu.Unlock()

	if c == nil {
		var err error
		c, err = connect()
		if err != nil {
			return "", err
		}
	}

	session, err := c.NewSession()
	if err != nil && allowRetry {
		// Connection appears dead — reset and retry once.
		shutdownClient()
		return execCommandWithRetry(command, false)
	}
	if err != nil {
		return "", fmt.Errorf("create ssh session: %w", err)
	}
	defer session.Close()

	rsp, err := session.CombinedOutput(command)
	if err != nil {
		return "", fmt.Errorf("execute command: %w", err)
	}
	return string(rsp), nil
}

// WaitForConnection blocks until the first SSH connection succeeds, retrying every 5s.
func WaitForConnection() {
	for {
		_, err := connect()
		if err == nil {
			log.Infof("SSH connected to %s", sshTarget)
			return
		}
		log.Warnf("SSH connection to %s failed (retrying in 5s): %v", sshTarget, err)
		time.Sleep(5 * time.Second)
	}
}

func shutdownClient() {
	clientMu.Lock()
	defer clientMu.Unlock()
	if client != nil {
		client.Close()
		client = nil
	}
}
