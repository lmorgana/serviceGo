package tests

import (
	"github.com/lamoda/gonkey/runner"
	"net/http/httptest"
	"testing"
)

func Test_API(t *testing.T) {

	srv := httptest.NewServer(nil)
	defer srv.Close()

	runner.RunWithTesting(t, &runner.RunWithTestingParams{
		Server:   srv,
		TestsDir: "./cases",
	})
}
