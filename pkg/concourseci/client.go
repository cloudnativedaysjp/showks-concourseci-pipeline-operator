package concourseci

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
)

type ConcourseCIClientInterface interface {
	SetPipeline(target string, pipeline string, manifest string) error
	DestroyPipeline(target string, pipeline string) error
}

func NewClient() ConcourseCIClientInterface {
	return &ConcourseCIClient{}
}

type ConcourseCIClient struct {
}

func (c *ConcourseCIClient) SetPipeline(target string, pipeline string, manifest string) error {
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
	args := []string{
		"-t", target,
		"destroy-pipeline",
		"-p", pipeline,
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
