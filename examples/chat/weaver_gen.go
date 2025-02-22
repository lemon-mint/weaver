package main

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/examples/chat/ImageScaler",
		Iface: reflect.TypeOf((*ImageScaler)(nil)).Elem(),
		New:   func() any { return &scaler{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return imageScaler_local_stub{impl: impl.(ImageScaler), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return imageScaler_client_stub{stub: stub, scaleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/chat/ImageScaler", Method: "Scale"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return imageScaler_server_stub{impl: impl.(ImageScaler), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/examples/chat/LocalCache",
		Iface: reflect.TypeOf((*LocalCache)(nil)).Elem(),
		New:   func() any { return &localCache{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return localCache_local_stub{impl: impl.(LocalCache), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return localCache_client_stub{stub: stub, getMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/chat/LocalCache", Method: "Get"}), putMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/chat/LocalCache", Method: "Put"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return localCache_server_stub{impl: impl.(LocalCache), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:     "github.com/ServiceWeaver/weaver/examples/chat/SQLStore",
		Iface:    reflect.TypeOf((*SQLStore)(nil)).Elem(),
		New:      func() any { return &sqlStore{} },
		ConfigFn: func(i any) any { return i.(*sqlStore).WithConfig.Config() },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return sQLStore_local_stub{impl: impl.(SQLStore), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return sQLStore_client_stub{stub: stub, createThreadMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/chat/SQLStore", Method: "CreateThread"}), createPostMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/chat/SQLStore", Method: "CreatePost"}), getFeedMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/chat/SQLStore", Method: "GetFeed"}), getImageMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/chat/SQLStore", Method: "GetImage"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return sQLStore_server_stub{impl: impl.(SQLStore), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type imageScaler_local_stub struct {
	impl   ImageScaler
	tracer trace.Tracer
}

func (s imageScaler_local_stub) Scale(ctx context.Context, a0 []byte, a1 int, a2 int) (r0 []byte, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.ImageScaler.Scale", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Scale(ctx, a0, a1, a2)
}

type localCache_local_stub struct {
	impl   LocalCache
	tracer trace.Tracer
}

func (s localCache_local_stub) Get(ctx context.Context, a0 string) (r0 string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.LocalCache.Get", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Get(ctx, a0)
}

func (s localCache_local_stub) Put(ctx context.Context, a0 string, a1 string) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.LocalCache.Put", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Put(ctx, a0, a1)
}

type sQLStore_local_stub struct {
	impl   SQLStore
	tracer trace.Tracer
}

func (s sQLStore_local_stub) CreateThread(ctx context.Context, a0 string, a1 time.Time, a2 []string, a3 string, a4 []byte) (r0 ThreadID, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.SQLStore.CreateThread", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CreateThread(ctx, a0, a1, a2, a3, a4)
}

func (s sQLStore_local_stub) CreatePost(ctx context.Context, a0 string, a1 time.Time, a2 ThreadID, a3 string) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.SQLStore.CreatePost", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CreatePost(ctx, a0, a1, a2, a3)
}

func (s sQLStore_local_stub) GetFeed(ctx context.Context, a0 string) (r0 []Thread, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.SQLStore.GetFeed", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetFeed(ctx, a0)
}

func (s sQLStore_local_stub) GetImage(ctx context.Context, a0 string, a1 ImageID) (r0 []byte, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.SQLStore.GetImage", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetImage(ctx, a0, a1)
}

// Client stub implementations.

type imageScaler_client_stub struct {
	stub         codegen.Stub
	scaleMetrics *codegen.MethodMetrics
}

func (s imageScaler_client_stub) Scale(ctx context.Context, a0 []byte, a1 int, a2 int) (r0 []byte, err error) {
	// Update metrics.
	start := time.Now()
	s.scaleMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.ImageScaler.Scale", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.scaleMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.scaleMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + (len(a0) * 1))
	size += 8
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_slice_byte_87461245(enc, a0)
	enc.Int(a1)
	enc.Int(a2)
	var shardKey uint64

	// Call the remote method.
	s.scaleMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.scaleMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_byte_87461245(dec)
	err = dec.Error()
	return
}

type localCache_client_stub struct {
	stub       codegen.Stub
	getMetrics *codegen.MethodMetrics
	putMetrics *codegen.MethodMetrics
}

func (s localCache_client_stub) Get(ctx context.Context, a0 string) (r0 string, err error) {
	// Update metrics.
	start := time.Now()
	s.getMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.LocalCache.Get", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
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
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.getMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.getMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

func (s localCache_client_stub) Put(ctx context.Context, a0 string, a1 string) (err error) {
	// Update metrics.
	start := time.Now()
	s.putMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.LocalCache.Put", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.putMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.putMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	size += (4 + len(a1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	enc.String(a1)
	var shardKey uint64

	// Call the remote method.
	s.putMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.putMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type sQLStore_client_stub struct {
	stub                codegen.Stub
	createThreadMetrics *codegen.MethodMetrics
	createPostMetrics   *codegen.MethodMetrics
	getFeedMetrics      *codegen.MethodMetrics
	getImageMetrics     *codegen.MethodMetrics
}

func (s sQLStore_client_stub) CreateThread(ctx context.Context, a0 string, a1 time.Time, a2 []string, a3 string, a4 []byte) (r0 ThreadID, err error) {
	// Update metrics.
	start := time.Now()
	s.createThreadMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.SQLStore.CreateThread", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.createThreadMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.createThreadMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	enc.String(a0)
	enc.EncodeBinaryMarshaler(&a1)
	serviceweaver_enc_slice_string_4af10117(enc, a2)
	enc.String(a3)
	serviceweaver_enc_slice_byte_87461245(enc, a4)
	var shardKey uint64

	// Call the remote method.
	s.createThreadMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.createThreadMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	*(*int64)(&r0) = dec.Int64()
	err = dec.Error()
	return
}

func (s sQLStore_client_stub) CreatePost(ctx context.Context, a0 string, a1 time.Time, a2 ThreadID, a3 string) (err error) {
	// Update metrics.
	start := time.Now()
	s.createPostMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.SQLStore.CreatePost", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.createPostMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.createPostMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	enc.String(a0)
	enc.EncodeBinaryMarshaler(&a1)
	enc.Int64((int64)(a2))
	enc.String(a3)
	var shardKey uint64

	// Call the remote method.
	s.createPostMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.createPostMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s sQLStore_client_stub) GetFeed(ctx context.Context, a0 string) (r0 []Thread, err error) {
	// Update metrics.
	start := time.Now()
	s.getFeedMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.SQLStore.GetFeed", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getFeedMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getFeedMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.getFeedMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.getFeedMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_Thread_511e1469(dec)
	err = dec.Error()
	return
}

func (s sQLStore_client_stub) GetImage(ctx context.Context, a0 string, a1 ImageID) (r0 []byte, err error) {
	// Update metrics.
	start := time.Now()
	s.getImageMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.SQLStore.GetImage", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getImageMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getImageMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	enc.Int64((int64)(a1))
	var shardKey uint64

	// Call the remote method.
	s.getImageMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 3, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.getImageMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_byte_87461245(dec)
	err = dec.Error()
	return
}

// Server stub implementations.

type imageScaler_server_stub struct {
	impl    ImageScaler
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s imageScaler_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Scale":
		return s.scale
	default:
		return nil
	}
}

func (s imageScaler_server_stub) scale(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 []byte
	a0 = serviceweaver_dec_slice_byte_87461245(dec)
	var a1 int
	a1 = dec.Int()
	var a2 int
	a2 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Scale(ctx, a0, a1, a2)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_byte_87461245(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type localCache_server_stub struct {
	impl    LocalCache
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s localCache_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Get":
		return s.get
	case "Put":
		return s.put
	default:
		return nil
	}
}

func (s localCache_server_stub) get(ctx context.Context, args []byte) (res []byte, err error) {
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

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Get(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s localCache_server_stub) put(ctx context.Context, args []byte) (res []byte, err error) {
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
	var a1 string
	a1 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Put(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type sQLStore_server_stub struct {
	impl    SQLStore
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s sQLStore_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "CreateThread":
		return s.createThread
	case "CreatePost":
		return s.createPost
	case "GetFeed":
		return s.getFeed
	case "GetImage":
		return s.getImage
	default:
		return nil
	}
}

func (s sQLStore_server_stub) createThread(ctx context.Context, args []byte) (res []byte, err error) {
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
	var a1 time.Time
	dec.DecodeBinaryUnmarshaler(&a1)
	var a2 []string
	a2 = serviceweaver_dec_slice_string_4af10117(dec)
	var a3 string
	a3 = dec.String()
	var a4 []byte
	a4 = serviceweaver_dec_slice_byte_87461245(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.CreateThread(ctx, a0, a1, a2, a3, a4)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int64((int64)(r0))
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s sQLStore_server_stub) createPost(ctx context.Context, args []byte) (res []byte, err error) {
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
	var a1 time.Time
	dec.DecodeBinaryUnmarshaler(&a1)
	var a2 ThreadID
	*(*int64)(&a2) = dec.Int64()
	var a3 string
	a3 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.CreatePost(ctx, a0, a1, a2, a3)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s sQLStore_server_stub) getFeed(ctx context.Context, args []byte) (res []byte, err error) {
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

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetFeed(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_Thread_511e1469(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s sQLStore_server_stub) getImage(ctx context.Context, args []byte) (res []byte, err error) {
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
	var a1 ImageID
	*(*int64)(&a1) = dec.Int64()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetImage(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_byte_87461245(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = &Post{}

func (x *Post) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Post.WeaverMarshal: nil receiver"))
	}
	enc.Int64((int64)(x.ID))
	enc.String(x.Creator)
	enc.EncodeBinaryMarshaler(&x.When)
	enc.String(x.Text)
	enc.Int64((int64)(x.ImageID))
}

func (x *Post) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Post.WeaverUnmarshal: nil receiver"))
	}
	*(*int64)(&x.ID) = dec.Int64()
	x.Creator = dec.String()
	dec.DecodeBinaryUnmarshaler(&x.When)
	x.Text = dec.String()
	*(*int64)(&x.ImageID) = dec.Int64()
}

var _ codegen.AutoMarshal = &Thread{}

func (x *Thread) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Thread.WeaverMarshal: nil receiver"))
	}
	enc.Int64((int64)(x.ID))
	serviceweaver_enc_slice_Post_29a9ee83(enc, x.Posts)
}

func (x *Thread) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Thread.WeaverUnmarshal: nil receiver"))
	}
	*(*int64)(&x.ID) = dec.Int64()
	x.Posts = serviceweaver_dec_slice_Post_29a9ee83(dec)
}

func serviceweaver_enc_slice_Post_29a9ee83(enc *codegen.Encoder, arg []Post) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_Post_29a9ee83(dec *codegen.Decoder) []Post {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]Post, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_string_4af10117(enc *codegen.Encoder, arg []string) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.String(arg[i])
	}
}

func serviceweaver_dec_slice_string_4af10117(dec *codegen.Decoder) []string {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = dec.String()
	}
	return res
}

func serviceweaver_enc_slice_byte_87461245(enc *codegen.Encoder, arg []byte) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.Byte(arg[i])
	}
}

func serviceweaver_dec_slice_byte_87461245(dec *codegen.Decoder) []byte {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = dec.Byte()
	}
	return res
}

func serviceweaver_enc_slice_Thread_511e1469(enc *codegen.Encoder, arg []Thread) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_Thread_511e1469(dec *codegen.Decoder) []Thread {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]Thread, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}
