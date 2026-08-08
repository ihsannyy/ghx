package ghcli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type GHUserAPI struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GHEmailAPI struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

var sshUserRegex = regexp.MustCompile(`Hi ([a-zA-Z0-9-]+)!`)

func IsGHInstalled() bool {
	_, err := exec.LookPath("gh")
	return err == nil
}

func IsGitInstalled() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

func LoginWithToken(token string, host string) error {
	if host == "" {
		host = "github.com"
	}
	cmd := exec.Command("gh", "auth", "login", "--hostname", host, "--with-token")
	cmd.Stdin = strings.NewReader(token)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("gh auth login failed: %v (%s)", err, strings.TrimSpace(stderr.String()))
	}
	return nil
}

func GetActiveGHToken(host string) (string, error) {
	if host == "" {
		host = "github.com"
	}
	cmd := exec.Command("gh", "auth", "token", "--hostname", host)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(stdout.String()), nil
}

func GetActiveGHUsername(host string) (string, error) {
	if host == "" {
		host = "github.com"
	}
	token, err := GetActiveGHToken(host)
	if err != nil || token == "" {
		return "", fmt.Errorf("no active gh token found")
	}
	u, _, _, err := FetchUserInfo(token, host)
	return u, err
}

func TestSSHConnection(host string) (string, error) {
	if host == "" {
		host = "github.com"
	}
	target := "git@" + host
	cmd := exec.Command("ssh", "-T", "-o", "StrictHostKeyChecking=accept-new", "-o", "ConnectTimeout=5", target)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	_ = cmd.Run()

	out := stderr.String()
	matches := sshUserRegex.FindStringSubmatch(out)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", fmt.Errorf("ssh authentication not active or failed")
}

func FetchUserInfo(token string, host string) (username string, name string, email string, err error) {
	if host == "" {
		host = "github.com"
	}

	cmd := exec.Command("gh", "api", "user")
	cmd.Env = append(os.Environ(), "GH_TOKEN="+token)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", "", "", fmt.Errorf("failed to fetch GitHub user info: %v (%s)", err, strings.TrimSpace(stderr.String()))
	}

	var user GHUserAPI
	if err := json.Unmarshal(stdout.Bytes(), &user); err != nil {
		return "", "", "", fmt.Errorf("failed to parse GitHub user response: %v", err)
	}

	username = user.Login
	name = user.Name
	if name == "" {
		name = username
	}
	email = user.Email

	if email == "" {
		cmdEmails := exec.Command("gh", "api", "user/emails")
		cmdEmails.Env = append(os.Environ(), "GH_TOKEN="+token)
		var stdoutEmails bytes.Buffer
		cmdEmails.Stdout = &stdoutEmails
		if err := cmdEmails.Run(); err == nil {
			var emails []GHEmailAPI
			if err := json.Unmarshal(stdoutEmails.Bytes(), &emails); err == nil {
				for _, e := range emails {
					if e.Primary && e.Verified {
						email = e.Email
						break
					}
				}
				if email == "" && len(emails) > 0 {
					email = emails[0].Email
				}
			}
		}
	}

	if email == "" {
		email = username + "@users.noreply.github.com"
	}

	return username, name, email, nil
}

func SetGitUser(name string, email string) error {
	if name != "" {
		cmdName := exec.Command("git", "config", "--global", "user.name", name)
		if err := cmdName.Run(); err != nil {
			return fmt.Errorf("failed to set git user.name: %v", err)
		}
	}
	if email != "" {
		cmdEmail := exec.Command("git", "config", "--global", "user.email", email)
		if err := cmdEmail.Run(); err != nil {
			return fmt.Errorf("failed to set git user.email: %v", err)
		}
	}
	return nil
}

func GetGitUser() (name string, email string, err error) {
	cmdName := exec.Command("git", "config", "--global", "user.name")
	var stdoutName bytes.Buffer
	cmdName.Stdout = &stdoutName
	_ = cmdName.Run()
	name = strings.TrimSpace(stdoutName.String())

	cmdEmail := exec.Command("git", "config", "--global", "user.email")
	var stdoutEmail bytes.Buffer
	cmdEmail.Stdout = &stdoutEmail
	_ = cmdEmail.Run()
	email = strings.TrimSpace(stdoutEmail.String())

	return name, email, nil
}
