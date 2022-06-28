package presentation

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Salaton/formula-one/pkg/presentation/graph"
	"github.com/Salaton/formula-one/pkg/presentation/graph/generated"
	"github.com/Salaton/formula-one/pkg/presentation/logger"
	"github.com/Salaton/formula-one/pkg/usecase"
	"github.com/Salaton/formula-one/pkg/usecase/raceschedule"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func Router(ctx context.Context) *mux.Router {
	r := mux.NewRouter()

	// Initialise the usecase
	raceSchedule := raceschedule.NewRaceScheduleImplementation()
	useCase := usecase.NewFormulaOneUseCase(raceSchedule)
	r.Path("/ide").HandlerFunc(playground.Handler("GraphQL Playground", "/graphql"))

	// Graphql endpoint
	authR := r.Path("/graphql").Subrouter()
	authR.Methods(
		http.MethodPost,
		http.MethodGet,
	).HandlerFunc(GQLGenHandler(ctx, *useCase))

	return r
}

type Options struct {
	Port int
}

// Create a server
func CreateServer(ctx context.Context, opts Options) *http.Server {
	router := Router(ctx)
	address := fmt.Sprintf(":%v", opts.Port)
	srv := &http.Server{
		Addr:         address,
		Handler:      router,
		IdleTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}

	return srv
}

func InitializeServer(ctx context.Context, port int) error {
	srv := CreateServer(ctx, Options{Port: port})
	logger := logger.NewLogger()

	// Create a channel that will be used to listen for cancellation signales
	done := make(chan os.Signal, 1)

	// Relay signals to the channel
	// We listen for signals, capture them if they happen and are then sent to the channel
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log(zerolog.FatalLevel, "an error occured while starting the server: %v", err)
		}
	}()
	logger.Log(zerolog.InfoLevel, "Server running on port: %v", port)

	// This listens for output from the `done` channel created.
	// If there's any output, it'll begin the process of shutting down the server
	<-done

	cancelContext, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer func() {
		// closes databases, etc
		cancel()
	}()

	err := srv.Shutdown(cancelContext)
	logger.Log(zerolog.InfoLevel, "server shutting down ...")
	if err != nil {
		logger.Log(zerolog.FatalLevel, "server shutdown failed: %w", err)
		return err
	}
	logger.Log(zerolog.InfoLevel, "server exited properly")

	return nil
}

func GQLGenHandler(ctx context.Context, formulaOneUsecase usecase.FormulaOne) http.HandlerFunc {
	resolver := graph.NewResolver(ctx, formulaOneUsecase)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	return func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	}
}
