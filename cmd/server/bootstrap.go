package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/xfrr/gophersys/internal/queries"
	"github.com/xfrr/gophersys/pkg/bus"
	"github.com/xfrr/gophersys/pkg/env"
	"github.com/xfrr/gophersys/pkg/logger"
	"github.com/xfrr/gophersys/web"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	gophergrpc "github.com/xfrr/gophersys/grpc"
	gophercmd "github.com/xfrr/gophersys/internal/commands"
	gophermgo "github.com/xfrr/gophersys/internal/storage/mongodb"
)

var (
	// grpc server env vars
	serverPort    = env.GetStr("GOPHERS_GRPC_SERVER_PORT", "50051")
	serverCert    = env.GetStr("GOPHERS_GRPC_SERVER_CERT", "")
	serverCertKey = env.GetStr("GOPHERS_GRPC_SERVER_CERT_KEY", "")
	serverVersion = env.GetStr("GOPHERS_GRPC_SERVER_VERSION", "v1.0.0")

	// database env vars
	databaseHost = env.GetStr("GOPHERS_DATABASE_HOST", "localhost")
	databasePort = env.GetStr("GOPHERS_DATABASE_PORT", "27017")
	databaseUser = env.GetStr("GOPHERS_DATABASE_USER", "root")
	databasePass = env.GetStr("GOPHERS_DATABASE_PASS", "higopher!")

	// mongo env vars
	mongoConnTimeout = env.GetDuration("GOPHERS_MONGO_CONN_TIMEOUT", "5s")

	// web app
	webAppEnabled = env.GetBool("GOPHERS_WEB_APP_ENABLED", true)
	webAppPort    = env.GetStr("GOPHERS_WEB_APP_PORT", "8080")

	// logger
	logLevel = env.GetStr("GOPHERS_LOG_LEVEL", "debug")
)

type Container struct {
	logger logger.Logger
}

func NewContainer() *Container {
	return &Container{
		logger: newLogger("bootstrap", logLevel),
	}
}

func (c *Container) Start(ctx context.Context) error {
	c.logger.Info().Msg("setting up the server...")

	client := c.connectMongoDB(ctx)
	repo := gophermgo.NewRepository(client)
	cmdbus := newCommandBus(repo)
	querybus := newQueryBus(repo)
	grpcServer := c.newGRPCServer(serverPort, cmdbus, querybus)

	err := c.executeMigrations(ctx, repo, c.logger)
	if err != nil {
		c.logger.Fatal().Err(err).Msg("failed to execute migrations")
	}

	if webAppEnabled {
		go startWebApp(cmdbus, querybus)
	}

	return c.startGRPCServer(ctx, grpcServer, c.logger)
}

func newCommandBus(repo *gophermgo.MongoDBRepository) *bus.InMemoryMessageBus {
	logr := newLogger("commandBus", logLevel)
	return gophercmd.NewBus(repo, logr)
}

func newQueryBus(repo *gophermgo.MongoDBRepository) *bus.InMemoryMessageBus {
	logr := newLogger("queryBus", logLevel)
	return queries.NewBus(repo, logr)
}

func (c *Container) connectMongoDB(ctx context.Context) (client *mongo.Client) {
	var err error

	ctx, cancel := context.WithTimeout(ctx, mongoConnTimeout)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		databaseUser,
		databasePass,
		databaseHost,
		databasePort,
	)

	if client, err = mongo.Connect(ctx,
		options.Client().ApplyURI(uri),
	); err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return client
}

func (c *Container) executeMigrations(ctx context.Context, repo *gophermgo.MongoDBRepository, logger logger.Logger) error {
	err := repo.RunMigrations(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to run mongodb migrations")
	}

	logger.Debug().Msg("mongodb migrations executed successfully")
	return nil
}

func (c *Container) newGRPCServer(port string, cmdbus bus.Bus, querybus bus.Bus) *gophergrpc.Server {
	var opts []grpc.ServerOption
	svc := gophergrpc.NewService(cmdbus, querybus)

	// create credentials if cert files are provided
	if serverCert != "" && serverCertKey != "" {
		creds, err := credentials.NewServerTLSFromFile(serverCert, serverCertKey)
		if err != nil {
			panic(err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	return gophergrpc.NewServer(port, svc, opts...)
}

func (c *Container) startGRPCServer(ctx context.Context, grpcServer *gophergrpc.Server, logger logger.Logger) error {
	logger.Info().
		Str("port", serverPort).
		Str("version", serverVersion).
		Msg("starting gRPC server")

	defer func() {
		c.logger.Info().
			Str("port", serverPort).
			Str("version", serverVersion).
			Msg("stopping gRPC server gracefully")
		grpcServer.GracefulStop()
	}()

	errCh := make(chan error)
	go func() {
		errCh <- grpcServer.Start()
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return nil
	}
}

func startWebApp(cmdbus bus.Bus, querybus bus.Bus) {
	logr := newLogger("webApp", logLevel)
	logr.Info().
		Str("port", webAppPort).
		Msg("starting web app")

	webApp, err := web.NewApp(cmdbus, querybus, logr)
	if err != nil {
		logr.Fatal().Err(err).Msg("failed to create web app")
	}

	err = webApp.ListenAndServe(webAppPort)
	if err != nil {
		logr.Fatal().Err(err).Msg("failed to start web app")
	}
}

func newLogger(svcname string, level string) logger.Logger {
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		lvl = zerolog.DebugLevel
	}
	return logger.NewLogger(svcname, lvl)
}
