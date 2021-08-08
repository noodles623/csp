package store

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/noodles623/csp/objects"
)

type AssetStore interface {
	Get(ctx context.Context, in *objects.GetRequest) (*objects.Asset, error)
	List(ctx context.Context, in *objects.ListRequest) ([]*objects.Asset, error)
	Create(ctx context.Context, in *objects.CreateRequest) error
	Delete(ctx context.Context, in *objects.DeleteRequest) error
}

func init() {
	rand.Seed(time.Now().UTC().Unix())
}

func GenerateUniqueID() string {
	word := []byte("0987654321")
	rand.Shuffle(len(word), func(i, j int) {
		word[i], word[j] = word[j], word[i]
	})
	now := time.Now().UTC()
	return fmt.Sprintf("%010v-%010v-%s", now.Unix(), now.Nanosecond(), string(word))
}
