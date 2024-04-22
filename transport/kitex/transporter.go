package kitex

import (
	"github.com/dobyte/due/transport/kitex/v2/gate"
	"github.com/dobyte/due/transport/kitex/v2/internal/server"
	"github.com/dobyte/due/transport/kitex/v2/node"
	"github.com/dobyte/due/v2/core/endpoint"
	"github.com/dobyte/due/v2/registry"
	"github.com/dobyte/due/v2/transport"
	"sync"
)

type Transporter struct {
	opts *options
	once sync.Once
	//builder *client.Builder
	once1 sync.Once
	cli1  transport.GateClient
	once2 sync.Once
	cli2  transport.NodeClient
}

func NewTransporter(opts ...Option) *Transporter {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	return &Transporter{opts: o}
}

// SetDefaultDiscovery 设置默认的服务发现组件
func (t *Transporter) SetDefaultDiscovery(discovery registry.Discovery) {
	//if t.opts.client.Discovery == nil {
	//	t.opts.client.Discovery = discovery
	//}
}

// NewGateServer 新建网关服务器
func (t *Transporter) NewGateServer(provider transport.GateProvider) (transport.Server, error) {
	return gate.NewServer(provider, &t.opts.server)
}

// NewNodeServer 新建节点服务器
func (t *Transporter) NewNodeServer(provider transport.NodeProvider) (transport.Server, error) {
	return node.NewServer(provider, &t.opts.server)
}

// NewServiceServer 新建微服务服务器
func (t *Transporter) NewServiceServer() (transport.Server, error) {
	return server.NewServer(&t.opts.server)
}

// NewGateClient 新建网关客户端
func (t *Transporter) NewGateClient(ep *endpoint.Endpoint) (transport.GateClient, error) {
	//t.once.Do(func() {
	//	t.builder = client.NewBuilder(&t.opts.client)
	//})
	//
	//cli, err := t.builder.Build(ep.Target())
	//if err != nil {
	//	return nil, err
	//}

	t.once1.Do(func() {
		t.cli1, _ = gate.NewClient(ep.Address())
	})

	return t.cli1, nil
}

// NewNodeClient 新建节点客户端
func (t *Transporter) NewNodeClient(ep *endpoint.Endpoint) (transport.NodeClient, error) {
	//t.once.Do(func() {
	//	t.builder = client.NewBuilder(&t.opts.client)
	//})
	//
	//cli, err := t.builder.Build(ep.Target())
	//if err != nil {
	//	return nil, err
	//}

	//return node.NewClient(ep.Address())

	t.once2.Do(func() {
		t.cli2, _ = node.NewClient(ep.Address())
	})

	return t.cli2, nil
}

// NewServiceClient 新建微服务客户端
func (t *Transporter) NewServiceClient(target string) (transport.ServiceClient, error) {
	//t.once.Do(func() {
	//	t.builder = client.NewBuilder(&t.opts.client)
	//})
	//
	//cli, err := t.builder.Build(target)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return client.NewClient(cli), nil

	return nil, nil
}