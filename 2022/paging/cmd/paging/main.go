package main

import (
	"context"
	"github.com/crnkofe/blog/2022/paging/pkg/sql"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	ctx := context.Background()
	err := sql.Init()
	if err != nil {
		log.Fatal().Msg("could not connect to DB. Make sure the sample DB is running")
	}
	defer func() {
		err := sql.Close()
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
	for {
		pagedComputers, err := comp.GetPaged(ctx, lastID, 10)
		if err != nil {
			log.Info().Err(err).Msg("failed getting computers")
		}

		if len(pagedComputers) == 0 {
			break
		}
		lastID = pagedComputers[len(pagedComputers)-1].ID
		log.Info().Msgf("Paged computers - page: %d", page)
		for _, comp := range pagedComputers {
			log.Info().Msg(comp.String())
		}
		page++
	}
}
