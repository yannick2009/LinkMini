package utils_test

import (
	"linkmini/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashURL(t *testing.T) {
	t.Run(
		"generate hash",
		func(t *testing.T) {
			hash := utils.GenerateHash()

			assert.IsType(t, "string", hash)
			assert.Len(t, hash, utils.HashLength)
		},
	)
}
