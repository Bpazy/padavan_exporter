package collector

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"strconv"
)

func mustGetContent(sshClient *ssh.Client, path string) string {
	rsp, err := execCommand(sshClient, "cat "+path)
	if err != nil {
		panic(err)
	}
	return rsp
}

func execCommand(sshClient *ssh.Client, command string) (string, error) {
	session, err := sshClient.NewSession()
	if err != nil {
		return "", fmt.Errorf("create ssh session failed: %+v", err)
	}
	defer session.Close()

	rsp, err := session.CombinedOutput(command)
	if err != nil {
		return "", fmt.Errorf("execute ssh command failed: %+v", err)
	}
	return string(rsp), nil
}

func mustParseFloat(fs string) float64 {
	float, err := strconv.ParseFloat(fs, 32)
	if err != nil {
		log.Printf("%+v", err)
		return 0
	}
	return float
}
