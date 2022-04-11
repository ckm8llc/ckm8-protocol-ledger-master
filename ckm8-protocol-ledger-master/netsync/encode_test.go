// +build unit

package netsync

import (
	"testing"

	"github.com/ckm8token/ckm8/dispatcher"

	"github.com/stretchr/testify/assert"
	"github.com/ckm8token/ckm8/common"
)

func TestMessageEncoding(t *testing.T) {
	assert := assert.New(t)

	dataReq := dispatcher.DataRequest{ChannelID: common.ChannelIDBlock, Entries: []string{"A0"}}

	b, err := encodeMessage(dataReq)
	assert.Nil(err)

	raw, err := decodeMessage(b)
	dataReq2 := raw.(dispatcher.DataRequest)
	assert.Nil(err)
	assert.Equal(common.ChannelIDBlock, dataReq.ChannelID)
	assert.Equal(1, len(dataReq2.Entries))
	assert.Equal("A0", dataReq2.Entries[0])
}
