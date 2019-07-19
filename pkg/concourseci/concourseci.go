package concourseci

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/exec"
)

type ConcourseCIClientInterface interface {
	Login() error
	SetPipeline(target string, pipeline string, manifest string) error
	DestroyPipeline(target string, pipeline string) error
}

func NewClient(url string, target, team string, username string, password string) ConcourseCIClientInterface {
	flyPath := "fly"
	if os.Getenv("CONCOURSECI_FLY_PATH") != "" {
		flyPath = os.Getenv("CONCOURSECI_FLY_PATH")
	}

	return &ConcourseCIClient{
		Url:      url,
		Target:   target,
		Team:     team,
		Username: username,
		Password: password,
		FlyPath: flyPath,
	}
}

type ConcourseCIClient struct {
	Url      string
	Target   string
	Team     string
	Username string
	Password string
	FlyPath string
}

func (c *ConcourseCIClient) Login() error {
	args := []string{
		"-t", c.Target,
		"login",
		"-k",
		"-u", c.Username,
		"-p", c.Password,
		"-c", c.Url,
		"-n", c.Team,
	}
	cmd := exec.Command(c.FlyPath, args...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		return errors.Wrapf(err, "Failed to login: %s", out)
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
	cmd := exec.Command(c.FlyPath, args...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		return errors.Wrapf(err, "Failed to SetPipeline: %s", out)
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
		"-n",
		"-p", pipeline,
	}
	cmd := exec.Command(c.FlyPath, args...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		return errors.Wrapf(err, "Failed to DestroyPipeline: %s", out)
	}

	return nil
}
