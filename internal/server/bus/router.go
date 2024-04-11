package bus

import (
	"context"
	"fmt"

	"github.com/krivenkov/order/internal/model/user"
	"github.com/krivenkov/order/internal/server/bus/user_handler"
	"github.com/krivenkov/pkg/bus"
	"github.com/krivenkov/pkg/bus/builder"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type router struct {
	cfg    Config
	busCfg builder.Config
	cg     *bus.ConsumerGroup
}

func newRouter(cfg Config, busCfg builder.Config, lc fx.Lifecycle) *router {
	cg := bus.NewConsumerGroup()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				_ = cg.Consume(context.Background())
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return cg.Close(ctx)
		},
	})

	return &router{
		cfg:    cfg,
		busCfg: busCfg,
		cg:     cg,
	}
}

func registerRoutes(
	logger *zap.Logger,
	r *router,
	updateUserHandler *user_handler.Handler,
) error {
	var (
		consumer bus.ClientConsumer
		err      error
	)

	consumer, err = builder.NewConsumer[user.User](
		r.busCfg, logger, user.UpdateUserTopic, r.cfg.WorkerID, updateUserHandler,
	)
	if err != nil {
		return fmt.Errorf("create consumer PublishAuthor: %w", err)
	}

	r.cg.Add(consumer)

	return nil
}
