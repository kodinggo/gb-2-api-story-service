package helper

import (
	"fmt"

	"github.com/kodinggo/gb-2-api-story-service/internal/config"
)

func GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		config.GetDbUser(),
		config.GetDbPassword(),
		config.GetDbHost(),
		config.GetDbPort(),
		config.GetDbName(),
	)
}
