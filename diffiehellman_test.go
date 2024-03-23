package diffiehellman_test

import (
	"strconv"
	"testing"

	"github.com/erfanmomeniii/diffiehellman"
	"github.com/stretchr/testify/require"
)

func Test_TransportKey(t *testing.T) {
	cases := []struct {
		primeNumber int64
	}{
		{
			primeNumber: 10141,
		},
		{
			primeNumber: 10061,
		},
		{
			primeNumber: 10037,
		},
	}

	for _, c := range cases {
		clientA, _ := diffiehellman.New(c.primeNumber)
		clientB, _ := diffiehellman.New(c.primeNumber)

		t.Run(strconv.FormatInt(c.primeNumber, 10), func(t *testing.T) {
			require.Equal(t, clientA.GenerateTransportKey(clientB.GetPublicKey()), clientB.GenerateTransportKey(clientA.GetPublicKey()))
		})
	}
}
