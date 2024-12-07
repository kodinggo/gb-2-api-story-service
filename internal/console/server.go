package console

import (
	"fmt"
	"log"
	"sync"

	"github.com/kodinggo/gb-2-api-comment-service/pb/comment_service"
	"github.com/kodinggo/gb-2-api-story-service/db"
	"github.com/kodinggo/gb-2-api-story-service/internal/config"
	handlerHttp "github.com/kodinggo/gb-2-api-story-service/internal/delivery/http"
	"github.com/kodinggo/gb-2-api-story-service/internal/repository"
	"github.com/kodinggo/gb-2-api-story-service/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	grpcCommentClient :=initgRPCCommentClient()
	storyUsecase := usecase.NewStoryUsecase(storyRepo, grpcCommentClient,categoryRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	e := echo.New()

	handlerHttp.NewStoryHandler(e, storyUsecase)
	handlerHttp.NewCategoryHandler(e, categoryUsecase)

	var wg sync.WaitGroup
	errCh := make(chan error, 2)
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println(config.CommentgRPCHost())
		fmt.Println("database connected ready to use")
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
func initgRPCCommentClient() comment_service.CommentServiceClient {
	// connect to grpc server without credentials
	conn, err := grpc.NewClient(config.CommentgRPCHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("failed to open connection grpc server, error %v", err)
	}
	// init grpc client as package dependency from grpc-server repository
	return comment_service.NewCommentServiceClient(conn)
}
