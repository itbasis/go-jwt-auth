package internal

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	itbasisTestUtils "github.com/itbasis/go-test-utils/v2"
	"go.uber.org/zap"
)

var (
	LogInterceptorOpts = []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}
)

// nolint:cyclop,varnamelen,gomnd
func InterceptorLogger() logging.Logger {
	return logging.LoggerFunc(
		func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
			f := make([]zap.Field, 0, len(fields)/2)

			for i := 0; i < len(fields); i += 2 {
				key := fields[i]
				value := fields[i+1]

				switch value := value.(type) {
				case string:
					f = append(f, zap.String(key.(string), value))
				case int:
					f = append(f, zap.Int(key.(string), value))
				case bool:
					f = append(f, zap.Bool(key.(string), value))
				default:
					f = append(f, zap.Any(key.(string), value))
				}
			}

			logger := itbasisTestUtils.TestLogger.WithOptions(zap.AddCallerSkip(1)).With(f...)

			switch lvl {
			case logging.LevelDebug:
				logger.Debug(msg)
			case logging.LevelInfo:
				logger.Info(msg)
			case logging.LevelWarn:
				logger.Warn(msg)
			case logging.LevelError:
				logger.Error(msg)
			default:
				panic(fmt.Sprintf("unknown level %v", lvl))
			}
		},
	)
}
