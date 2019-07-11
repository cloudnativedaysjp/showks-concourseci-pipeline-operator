package concourseci

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
)

type ConcourseCIClientInterface interface {
	Login() error
	SetPipeline(target string, pipeline string, manifest string) error
	DestroyPipeline(target string, pipeline string) error
}

func NewClient(url string, team string, username string, password string) ConcourseCIClientInterface {
	return &ConcourseCIClient{
		Url:      url,
		Team:     team,
		Username: username,
		Password: password,
	}
}

type ConcourseCIClient struct {
	Url      string
	Team     string
	Username string
	Password string
}

func (c *ConcourseCIClient) Login() error {
	args := []string{
		"-t", c.Team,
		"login",
		"-k",
		"-u", c.Username,
		"-p", c.Password,
		"-c", c.Url,
	}
	cmd := exec.Command("fly", args...)
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
func (c *ConcourseCIClient) SetPipeline(target string, pipeline string, manifest string) error {
	err := c.Login()
	if err != nil {
		return err
	}

	tmpfile, err := ioutil.TempFile("", "manifest")
	if err != nil {
		return err
	}

	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	tmpfile.Write([]byte(manifest))
	args := []string{
		"-t", target,
		"set-pipeline",
		"-n",
		"-p", pipeline,
		"-c", tmpfile.Name(),
	}
	cmd := exec.Command("fly", args...)
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (c *ConcourseCIClient) DestroyPipeline(target string, pipeline string) error {
	err := c.Login()
	if err != nil {
		return err
	}

	args := []string{
		"-t", target,
		"destroy-pipeline",
		"-p", pipeline,
	}
	cmd := exec.Command("fly", args...)
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
