// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package execinfrapb

import (
	"context"
	"net"
	"time"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/rpc"
	"github.com/cockroachdb/cockroach/pkg/settings/cluster"
	"github.com/cockroachdb/cockroach/pkg/util"
	"github.com/cockroachdb/cockroach/pkg/util/hlc"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/cockroach/pkg/util/netutil"
	"github.com/cockroachdb/cockroach/pkg/util/stop"
	"github.com/cockroachdb/cockroach/pkg/util/tracing"
	"github.com/cockroachdb/cockroach/pkg/util/uuid"
)

// CallbackMetadataSource is a utility struct that implements the MetadataSource
// interface by calling a provided callback.
type CallbackMetadataSource struct {
	DrainMetaCb func(context.Context) []ProducerMetadata
}

// DrainMeta is part of the MetadataSource interface.
func (s CallbackMetadataSource) DrainMeta(ctx context.Context) []ProducerMetadata {
	return s.DrainMetaCb(ctx)
}

func newInsecureRPCContext(stopper *stop.Stopper) *rpc.Context {
	return rpc.NewContext(
		log.AmbientContext{Tracer: tracing.NewTracer()},
		&base.Config{Insecure: true},
		hlc.NewClock(hlc.UnixNano, time.Nanosecond),
		stopper,
		&cluster.MakeTestingClusterSettings().Version,
	)
}

// StartMockDistSQLServer starts a MockDistSQLServer and returns the address on
// which it's listening.
func StartMockDistSQLServer(
	clock *hlc.Clock, stopper *stop.Stopper, nodeID roachpb.NodeID,
) (uuid.UUID, *MockDistSQLServer, net.Addr, error) {
	rpcContext := newInsecureRPCContext(stopper)
	rpcContext.NodeID.Set(context.TODO(), nodeID)
	server := rpc.NewServer(rpcContext)
	mock := newMockDistSQLServer()
	RegisterDistSQLServer(server, mock)
	ln, err := netutil.ListenAndServeGRPC(stopper, server, util.IsolatedTestAddr)
	if err != nil {
		return uuid.Nil, nil, nil, err
	}
	return rpcContext.ClusterID.Get(), mock, ln.Addr(), nil
}

// MockDistSQLServer implements the DistSQLServer (gRPC) interface and allows
// clients to control the inbound streams.
type MockDistSQLServer struct {
	InboundStreams   chan InboundStreamNotification
	RunSyncFlowCalls chan RunSyncFlowCall
}

// InboundStreamNotification is the MockDistSQLServer's way to tell its clients
// that a new gRPC call has arrived and thus a stream has arrived. The rpc
// handler is blocked until Donec is signaled.
type InboundStreamNotification struct {
	Stream DistSQL_FlowStreamServer
	Donec  chan<- error
}

// RunSyncFlowCall is the MockDistSQLServer's way to tell its clients that a
// RunSyncFlowCall has arrived. The rpc handler is blocked until Donec is
// signaled.
type RunSyncFlowCall struct {
	Stream DistSQL_RunSyncFlowServer
	Donec  chan<- error
}

// MockDistSQLServer implements the DistSQLServer interface.
var _ DistSQLServer = &MockDistSQLServer{}

func newMockDistSQLServer() *MockDistSQLServer {
	return &MockDistSQLServer{
		InboundStreams:   make(chan InboundStreamNotification),
		RunSyncFlowCalls: make(chan RunSyncFlowCall),
	}
}

// RunSyncFlow is part of the DistSQLServer interface.
func (ds *MockDistSQLServer) RunSyncFlow(stream DistSQL_RunSyncFlowServer) error {
	donec := make(chan error)
	ds.RunSyncFlowCalls <- RunSyncFlowCall{Stream: stream, Donec: donec}
	return <-donec
}

// SetupFlow is part of the DistSQLServer interface.
func (ds *MockDistSQLServer) SetupFlow(
	_ context.Context, req *SetupFlowRequest,
) (*SimpleResponse, error) {
	return nil, nil
}

// FlowStream is part of the DistSQLServer interface.
func (ds *MockDistSQLServer) FlowStream(stream DistSQL_FlowStreamServer) error {
	donec := make(chan error)
	ds.InboundStreams <- InboundStreamNotification{Stream: stream, Donec: donec}
	return <-donec
}
