package cosmodb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"

	mock_metrics "github.com/Azure/ARO-RP/pkg/util/mocks/metrics"
)

type failingRoundTripper struct {
}

func (f *failingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("it broke %s", "bad")
}

func TestTracerRoundTripperRoundTrip(t *testing.T) {
	gc := gomock.NewController(t)
	m := mock_metrics.NewMockInterface(gc)
	tripper := &tracerRoundTripper{
		log: logrus.NewEntry(logrus.StandardLogger()),
		m:   m,
		tr:  &failingRoundTripper{},
	}
	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	_, err := tripper.RoundTrip(req)
	if err != nil {
		t.Errorf("tracerRoundTripper.RoundTrip() error = %v", err)
	}
}
