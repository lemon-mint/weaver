package generate

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
		Name:  "github.com/ServiceWeaver/weaver/weavertest/internal/generate/testApp",
		Iface: reflect.TypeOf((*testApp)(nil)).Elem(),
		New:   func() any { return &impl{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return testApp_local_stub{impl: impl.(testApp), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return testApp_client_stub{stub: stub, getMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/generate/testApp", Method: "Get"}), incPointerMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/generate/testApp", Method: "IncPointer"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return testApp_server_stub{impl: impl.(testApp), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type testApp_local_stub struct {
	impl   testApp
	tracer trace.Tracer
}

func (s testApp_local_stub) Get(ctx context.Context, a0 string, a1 behaviorType) (r0 int, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "generate.testApp.Get", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Get(ctx, a0, a1)
}

func (s testApp_local_stub) IncPointer(ctx context.Context, a0 *int) (r0 *int, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "generate.testApp.IncPointer", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.IncPointer(ctx, a0)
}

// Client stub implementations.

type testApp_client_stub struct {
	stub              codegen.Stub
	getMetrics        *codegen.MethodMetrics
	incPointerMetrics *codegen.MethodMetrics
}

func (s testApp_client_stub) Get(ctx context.Context, a0 string, a1 behaviorType) (r0 int, err error) {
	// Update metrics.
	start := time.Now()
	s.getMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "generate.testApp.Get", trace.WithSpanKind(trace.SpanKindClient))
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
			s.getMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	enc.Int((int)(a1))
	var shardKey uint64

	// Call the remote method.
	s.getMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = codegen.JoinErrors(weaver.SystemError, err)
		return
	}
	s.getMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Int()
	err = dec.Error()
	return
}

func (s testApp_client_stub) IncPointer(ctx context.Context, a0 *int) (r0 *int, err error) {
	// Update metrics.
	start := time.Now()
	s.incPointerMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "generate.testApp.IncPointer", trace.WithSpanKind(trace.SpanKindClient))
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
			s.incPointerMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.incPointerMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += serviceweaver_size_ptr_int_98a2a745(a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_ptr_int_98a2a745(enc, a0)
	var shardKey uint64

	// Call the remote method.
	s.incPointerMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	if err != nil {
		err = codegen.JoinErrors(weaver.SystemError, err)
		return
	}
	s.incPointerMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_int_98a2a745(dec)
	err = dec.Error()
	return
}

// Server stub implementations.

type testApp_server_stub struct {
	impl    testApp
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s testApp_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Get":
		return s.get
	case "IncPointer":
		return s.incPointer
	default:
		return nil
	}
}

func (s testApp_server_stub) get(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 behaviorType
	*(*int)(&a1) = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Get(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s testApp_server_stub) incPointer(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 *int
	a0 = serviceweaver_dec_ptr_int_98a2a745(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.IncPointer(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_int_98a2a745(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Encoding/decoding implementations.

func serviceweaver_enc_ptr_int_98a2a745(enc *codegen.Encoder, arg *int) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		enc.Int(*arg)
	}
}

func serviceweaver_dec_ptr_int_98a2a745(dec *codegen.Decoder) *int {
	if !dec.Bool() {
		return nil
	}
	var res int
	res = dec.Int()
	return &res
}

// Size implementations.

// serviceweaver_size_ptr_int_98a2a745 returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_ptr_int_98a2a745(x *int) int {
	if x == nil {
		return 1
	} else {
		return 1 + 8
	}
}
