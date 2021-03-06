package check

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/robdimsdale/concourse-pipeline-resource/concourse"
	"github.com/robdimsdale/concourse-pipeline-resource/concourse/api"
	"github.com/robdimsdale/concourse-pipeline-resource/fly"
	"github.com/robdimsdale/concourse-pipeline-resource/logger"
	"github.com/robdimsdale/concourse-pipeline-resource/pipelinerunner"
)

type CheckCommand struct {
	logger        logger.Logger
	logFilePath   string
	binaryVersion string
	flyConn       fly.FlyConn
	apiClient     api.Client
}

func NewCheckCommand(
	binaryVersion string,
	logger logger.Logger,
	logFilePath string,
	flyConn fly.FlyConn,
	apiClient api.Client,
) *CheckCommand {
	return &CheckCommand{
		logger:        logger,
		logFilePath:   logFilePath,
		binaryVersion: binaryVersion,
		flyConn:       flyConn,
		apiClient:     apiClient,
	}
}

func (c *CheckCommand) Run(input concourse.CheckRequest) (concourse.CheckResponse, error) {
	logDir := filepath.Dir(c.logFilePath)
	existingLogFiles, err := filepath.Glob(filepath.Join(logDir, "concourse-pipeline-resource-check.log*"))
	if err != nil {
		// This is untested because the only error returned by filepath.Glob is a
		// malformed glob, and this glob is hard-coded to be correct.
		return nil, err
	}

	for _, f := range existingLogFiles {
		if filepath.Base(f) != filepath.Base(c.logFilePath) {
			c.logger.Debugf("Removing existing log file: %s\n", f)
			err := os.Remove(f)
			if err != nil {
				// This is untested because it is too hard to force os.Remove to return
				// an error.
				return nil, err
			}
		}
	}

	c.logger.Debugf("Received input: %+v\n", input)

	c.logger.Debugf("Performing login\n")

	_, err = c.flyConn.Login(
		input.Source.Target,
		input.Source.Username,
		input.Source.Password,
	)
	if err != nil {
		return concourse.CheckResponse{}, err
	}

	c.logger.Debugf("Login successful\n")

	pipelines, err := c.apiClient.Pipelines()
	if err != nil {
		return nil, err
	}

	c.logger.Debugf("Found pipelines: %+v\n", pipelines)

	gpFunc := func(index int, pipeline api.Pipeline) (string, error) {
		c.logger.Debugf("Getting pipeline: %s\n", pipeline.Name)
		outBytes, err := c.flyConn.GetPipeline(pipeline.Name)

		if err != nil {
			return "", err
		}

		return string(outBytes), nil
	}

	pipelinesContents, err := pipelinerunner.RunForAllPipelines(gpFunc, pipelines, c.logger)
	if err != nil {
		return nil, err
	}

	allContent := strings.Join(pipelinesContents, "")

	pipelinesChecksumString := fmt.Sprintf(
		"%x",
		md5.Sum([]byte(allContent)),
	)
	c.logger.Debugf("pipeline content checksum: %s\n", pipelinesChecksumString)

	out := concourse.CheckResponse{
		concourse.Version{
			PipelinesChecksum: pipelinesChecksumString,
		},
	}

	c.logger.Debugf("Returning output: %+v\n", out)

	return out, nil
}
