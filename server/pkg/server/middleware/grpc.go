package middleware

import (
	"context"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LogMessageProducer() grpc_zap.MessageProducer {
	return func(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {

		fields := []zap.Field{
			zap.Error(err),
			zap.String("grpc.code", code.String()),
			duration,
		}

		st := status.Convert(err)
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *errdetails.DebugInfo:
				fields = append(fields, zap.String("err.debug.detail", t.Detail))
				fields = append(fields, zap.Strings("err.debug.stack", t.StackEntries))
			case *errdetails.ErrorInfo:
				fields = append(fields, zap.String("err.error.reason", t.Reason))
			case *errdetails.BadRequest:
				for _, violation := range t.GetFieldViolations() {
					fields = append(fields, zap.String("err.request.field", violation.GetField()))
					fields = append(fields, zap.String("err.request.violation", violation.GetDescription()))
				}
			}
		}

		ctxzap.Extract(ctx).Check(level, msg).Write(
			fields...
		)
	}
}
