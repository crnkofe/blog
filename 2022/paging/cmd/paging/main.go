package main

import (
	"context"
	"github.com/crnkofe/blog/2022/paging/pkg/sql"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// using number to make output less verbose
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	ctx := context.Background()
	err := sql.Init(ctx)
	if err != nil {
		log.Fatal().Msg("could not connect to DB. Make sure the sample DB is running")
	}
	defer func() {
		err := sql.Close(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("failed cleaning up SQL")
		}
	}()

	PrintAllComputers(ctx)
}

func PrintAllComputers(ctx context.Context) {
	comp := sql.Computer()

	lastID := 0
	page := 0
	limit := 10
	for {
		pagedComputers, err := comp.GetPaged(ctx, lastID, limit)
		if err != nil {
			log.Info().Err(err).Msg("failed getting computers")
		}

		log.Info().Msgf("Paged computers - page: %d", page)
		for _, comp := range pagedComputers {
			log.Info().Msg(comp.String())
		}

		if len(pagedComputers) < limit {
			break
		}
		lastID = pagedComputers[len(pagedComputers)-1].ID

		page++
	}
}
