package main

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:        "github.com/ServiceWeaver/weaver/examples/collatz/Even",
		Iface:       reflect.TypeOf((*Even)(nil)).Elem(),
		New:         func() any { return &even{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any { return even_local_stub{impl: impl.(Even), tracer: tracer} },
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return even_client_stub{stub: stub, doMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/collatz/Even", Method: "Do"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return even_server_stub{impl: impl.(Even), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:        "github.com/ServiceWeaver/weaver/examples/collatz/Odd",
		Iface:       reflect.TypeOf((*Odd)(nil)).Elem(),
		New:         func() any { return &odd{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any { return odd_local_stub{impl: impl.(Odd), tracer: tracer} },
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return odd_client_stub{stub: stub, doMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/collatz/Odd", Method: "Do"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return odd_server_stub{impl: impl.(Odd), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type even_local_stub struct {
	impl   Even
	tracer trace.Tracer
}

func (s even_local_stub) Do(ctx context.Context, a0 int) (r0 int, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Even.Do", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Do(ctx, a0)
}

type odd_local_stub struct {
	impl   Odd
	tracer trace.Tracer
}

func (s odd_local_stub) Do(ctx context.Context, a0 int) (r0 int, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Odd.Do", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Do(ctx, a0)
}

// Client stub implementations.

type even_client_stub struct {
	stub      codegen.Stub
	doMetrics *codegen.MethodMetrics
}

func (s even_client_stub) Do(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	start := time.Now()
	s.doMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Even.Do", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = codegen.JoinErrors(weaver.SystemError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.doMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.doMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	s.doMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = codegen.JoinErrors(weaver.SystemError, err)
		return
	}
	s.doMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Int()
	err = dec.Error()
	return
}

type odd_client_stub struct {
	stub      codegen.Stub
	doMetrics *codegen.MethodMetrics
}

func (s odd_client_stub) Do(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	start := time.Now()
	s.doMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Odd.Do", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = codegen.JoinErrors(weaver.SystemError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.doMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.doMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	s.doMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = codegen.JoinErrors(weaver.SystemError, err)
		return
	}
	s.doMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Int()
	err = dec.Error()
	return
}

// Server stub implementations.

type even_server_stub struct {
	impl    Even
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s even_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Do":
		return s.do
	default:
		return nil
	}
}

func (s even_server_stub) do(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Do(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type odd_server_stub struct {
	impl    Odd
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s odd_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Do":
		return s.do
	default:
		return nil
	}
}

func (s odd_server_stub) do(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Do(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}
