//go:build acceptance

package starsturn_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AcceptanceSuite struct {
	suite.Suite

	cmd  exec.Cmd
	pid  int
	port int
}

func (s *AcceptanceSuite) SetupTest() {
	var err error
	s.port, err = freeport.GetFreePort()
	if err != nil {
		log.Fatalf("could not find port %v", err)
	}

	s.cmd = *exec.Command("./stars-turn", "-port", fmt.Sprintf("%d", s.port))
	err = s.cmd.Start()
	if err != nil {
		log.Fatalf("could not start process %v", err)
	}

	// This is necessary as we need to wait for the process to start up before it can start listening
	time.Sleep(50 * time.Millisecond)
}

func (s *AcceptanceSuite) AfterTest() {
	s.cmd.Process.Kill()
}

func TestAcceptance(t *testing.T) {
	suite.Run(t, new(AcceptanceSuite))
}

func (s *AcceptanceSuite) TestProcessRunsContinually() {
	proc, _ := os.FindProcess(s.pid)
	assert.Zero(s.T(), proc.Pid)
}

func (s *AcceptanceSuite) TestRespondsToHealthCheck() {
	res, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:%d/health", s.port))
	require.NoError(s.T(), err)
	require.Equal(s.T(), 200, res.StatusCode)
}
