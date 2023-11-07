package dht

import (
	"context"
	"encoding/hex"
	"testing"
	"time"

	"github.com/anacrolix/dht/v2/bep44"
	"github.com/anacrolix/torrent/bencode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TBD54566975/did-dht-method/config"
	"github.com/TBD54566975/did-dht-method/internal/util"
)

func TestGetPutDHT(t *testing.T) {
	d, err := NewDHT(config.GetDefaultBootstrapPeers())
	require.NoError(t, err)

	pubKey, privKey, err := util.GenerateKeypair()
	require.NoError(t, err)

	put := &bep44.Put{
		V:   []byte("hello dht"),
		K:   (*[32]byte)(pubKey),
		Seq: time.Now().UnixMilli() / 1000,
	}
	put.Sign(privKey)

	id, err := d.Put(context.Background(), *put)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	got, err := d.Get(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, got)

	var payload string
	err = bencode.Unmarshal(got.V, &payload)
	require.NoError(t, err)

	assert.Equal(t, string(put.V.([]byte)), payload)
}

func TestKnownVector(t *testing.T) {
	pubKey := "796f7457532cd39697f4fccd1a2d7074e6c1f6c59e6ecf5dc16c8ecd6e3fea6c"
	privKey := "3077903f62fbcff4bdbae9b5129b01b78ab87f68b8b3e3d332f14ca13ad53464796f7457532cd39697f4fccd1a2d7074e6c1f6c59e6ecf5dc16c8ecd6e3fea6c"
	pubKeyBytes, _ := hex.DecodeString(pubKey)
	privKeyBytes, _ := hex.DecodeString(privKey)

	put := &bep44.Put{
		V:   []byte("Hello World!"),
		K:   (*[32]byte)(pubKeyBytes),
		Seq: 1,
	}
	put.Sign(privKeyBytes)

	assert.Equal(t, "48656c6c6f20576f726c6421", hex.EncodeToString(put.V.([]byte)))
	assert.Equal(t, "c1dc657a17f54ca51933b17b7370b87faae10c7edd560fd4baad543869e30e8154c510f4d0b0d94d1e683891b06a07cecd9f0be325fe8f8a0466fe38011b2d0a", hex.EncodeToString(put.Sig[:]))
	assert.Equal(t, "796f7457532cd39697f4fccd1a2d7074e6c1f6c59e6ecf5dc16c8ecd6e3fea6c", hex.EncodeToString(put.K[:]))
}
