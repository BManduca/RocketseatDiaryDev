package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/BManduca/RocketseatDiaryDev/tree/main/EventWeekGoReact/internal/api"
	"github.com/BManduca/RocketseatDiaryDev/tree/main/EventWeekGoReact/internal/store/pgstore"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	// criando um pool de conexões
	pool, err := pgxpool.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("WSBM_DATABASE_USER"),
		os.Getenv("WSBM_DATABASE_PASSWORD"),
		os.Getenv("WSBM_DATABASE_HOST"),
		os.Getenv("WSBM_DATABASE_PORT"),
		os.Getenv("WSBM_DATABASE_NAME"),
	))
	if err != nil {
		panic(err)
	}

	// quando a função main retornar, será fechado o pool de conexão
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	// definindo meu handler
	handler := api.NewHandler(pgstore.New(pool))
	// chamando de maneira assincrona a função para iniciar o meu server http

	go func() {
		if err := http.ListenAndServe(":8080", handler); err != nil {
			// caso o servidor seja fechado, não resultano assim em uma falha
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	// canal para receber sinais do meun OS e basicamente ele vai etsar aguardando ou recebendo um CTRL + C (interrupt)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	// bloquenado minha função main, ate receber um sinal e dar um exit.
	<-quit
}
