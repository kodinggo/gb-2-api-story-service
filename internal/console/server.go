package console

import (
	"sync"

	"github.com/kodinggo/gb-2-api-story-service/db"
	handlerHttp "github.com/kodinggo/gb-2-api-story-service/internal/delivery/http"
	"github.com/kodinggo/gb-2-api-story-service/internal/repository"
	"github.com/kodinggo/gb-2-api-story-service/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCMD)
}

var serverCMD = &cobra.Command{
	Use:   "httpsrv",
	Short: "Start server",
	Run:   httpServer,
}

func httpServer(cmd *cobra.Command, args []string) {
	mysql := db.NewMysql()
	defer mysql.Close()

	storyRepo := repository.NewStoryRepo(mysql)
	categoryRepo := repository.NewCategoryRepo(mysql)

	storyUsecase := usecase.NewStoryUsecase(storyRepo, categoryRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	e := echo.New()

	handlerHttp.NewStoryHandler(e, storyUsecase)
	handlerHttp.NewCategoryHandler(e, categoryUsecase)

	var wg sync.WaitGroup
	errCh := make(chan error, 2)
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := e.Start(":3000")
		if err != nil {
			errCh <- err
		}
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			logrus.Error(err.Error())
		}
	}

}
