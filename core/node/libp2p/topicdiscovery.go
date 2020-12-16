package libp2p

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	disc "github.com/libp2p/go-libp2p-discovery"

	"github.com/ipfs/go-ipfs/core/node/helpers"
	"go.uber.org/fx"
)

type AdvertiseOnly struct {
	base discovery.Discovery
}

func (a *AdvertiseOnly) Advertise(ctx context.Context, ns string, opts ...discovery.Option) (time.Duration, error) {
	return a.base.Advertise(ctx, ns, opts...)
}

func (a *AdvertiseOnly) FindPeers(ctx context.Context, ns string, opts ...discovery.Option) (<-chan peer.AddrInfo, error) {
	c := make(chan peer.AddrInfo)
	close(c)
	return c, nil
}

func TopicDiscovery() interface{} {
	return func(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host, cr BaseIpfsRouting) (service discovery.Discovery, err error) {
		baseDisc := disc.NewRoutingDiscovery(cr)
		return &AdvertiseOnly{base: baseDisc}, nil
	}
}
