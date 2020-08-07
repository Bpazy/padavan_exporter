package collector

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"strconv"
)

func getContent(sshClient *ssh.Client, path string) string {
	return execCommand(sshClient, "cat "+path)
}

func execCommand(sshClient *ssh.Client, command string) string {
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatalf("create ssh session failed: %+v", err)
	}
	defer session.Close()

	rsp, err := session.CombinedOutput(command)
	if err != nil {
		log.Fatalf("execute ssh command failed: %+v", err)
	}
	return string(rsp)
}

func mustParseFloat(fs string) float64 {
	float, err := strconv.ParseFloat(fs, 32)
	if err != nil {
		log.Printf("%+v", err)
		return 0
	}
	return float
}
