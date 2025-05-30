package whiteCatApi

import (
	"testing"
)

func TestName(t *testing.T) {
	client := NewCatClient()
	client.GetCinemaSeat("43612784620885AE80CA13A3034DFD9E")
}
