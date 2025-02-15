package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/timut2/avito_test_task/internal/handlers"
	"github.com/timut2/avito_test_task/internal/middleware"
	"github.com/timut2/avito_test_task/internal/repository"
	"github.com/timut2/avito_test_task/internal/service"
)

func main() {
	InitLogging()
	ctx := context.Background()

	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	))

	if err != nil {
		log.Fatalf("Failed to connect to database: %v ", err)
	} else {
		log.Println("connected to db")
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	infoRepo := repository.NewInfoRepository(db)
	infoService := service.NewInfoService(infoRepo)
	infoHandler := handlers.NewInfoHandler(infoService)

	sendRepo := repository.NewSendRepository(db, userRepo)
	sendService := service.NewSendService(sendRepo)
	sendHandler := handlers.NewSendHandler(sendService)

	itemRepo := repository.NewItemRepository(db)
	purchaseRepo := repository.NewPurchaseRepository(db, userRepo, itemRepo)
	buyService := service.NewBuyService(purchaseRepo)
	buyHandler := handlers.NewBuyHandler(buyService)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/auth", authHandler.Login)
	mux.HandleFunc("/api/info", infoHandler.Get)
	mux.HandleFunc("/api/sendCoin", sendHandler.Send)
	mux.HandleFunc("/api/buy/", buyHandler.Buy)
	handler := middleware.JWTMiddleware(mux)

	server := &http.Server{Addr: ":8080", Handler: handler}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}

func InitLogging() {
	handler := slog.Handler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	handler = NewHandlerMiddleware(handler)
	slog.SetDefault(slog.New(handler))
}

type HandlerMiddlware struct {
	next slog.Handler
}

func (h *HandlerMiddlware) Enabled(ctx context.Context, rec slog.Level) bool {
	return h.next.Enabled(ctx, rec)
}

func NewHandlerMiddleware(next slog.Handler) *HandlerMiddlware {
	return &HandlerMiddlware{next: next}
}

type logCtx struct {
	UserID  int
	Phone   string
	Gate    string
	Message string
}

type keyType int

const key = keyType(0)

func (h *HandlerMiddlware) Handle(ctx context.Context, rec slog.Record) error {
	if c, ok := ctx.Value(key).(logCtx); ok {
		if c.UserID != 0 {
			rec.Add("userID", c.UserID)
		}
		if c.Phone != "" {
			rec.Add("phone", c.Phone)
		}
		if c.Gate != "" {
			rec.Add("sms_gate", c.Gate)
		}
		if c.Message != "" {
			rec.Add("message", c.Message)
		}
	}
	return h.next.Handle(ctx, rec)
}

func (h *HandlerMiddlware) WithGroup(name string) slog.Handler {
	return &HandlerMiddlware{next: h.next.WithGroup(name)} // не забыть обернуть, но осторожно
}
func (h *HandlerMiddlware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &HandlerMiddlware{next: h.next.WithAttrs(attrs)} // не забыть обернуть, но осторожно
}

func WithLogUserID(ctx context.Context, userID int) context.Context {
	if c, ok := ctx.Value(key).(logCtx); ok {
		c.UserID = userID
		return context.WithValue(ctx, key, c)
	}
	return context.WithValue(ctx, key, logCtx{UserID: userID})
}
