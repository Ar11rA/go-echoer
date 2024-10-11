// log_service.go
package services

import (
	"fmt"
	"os"
	"quote-server/types"
	"quote-server/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogService interface {
	InsertLog(appID, logs string) error
	GetLogsByApplicationID(appID string, limit int) ([]types.Log, error)
}

// LogService is responsible for log operations
type LogServiceImpl struct {
	mongoUtil utils.MongoDBClient
}

// NewLogService creates a new instance of LogService
func NewLogService(mongoUtil utils.MongoDBClient) LogService {
	return &LogServiceImpl{mongoUtil: mongoUtil}
}

// InsertLog inserts a new log into the logs collection
func (s *LogServiceImpl) InsertLog(appID, logs string) error {
	logEntry := types.Log{
		ApplicationID: appID,
		Logs:          logs,
		Timestamp:     time.Now(),
	}

	err := s.mongoUtil.InsertOne(os.Getenv("MONGODB_COLLECTION"), logEntry)
	if err != nil {
		return fmt.Errorf("failed to insert log: %w", err)
	}
	return nil
}

func (s *LogServiceImpl) GetLogsByApplicationID(appID string, limit int) ([]types.Log, error) {
	filter := bson.M{"application_id": appID}
	logEntries, err := s.mongoUtil.FindByFilter(os.Getenv("MONGODB_COLLECTION"), filter, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}

	var logs []types.Log
	for _, entry := range logEntries {
		log := types.Log{
			ApplicationID: entry["application_id"].(string),               // Cast to string
			Logs:          entry["logs"].(string),                         // Cast to string
			Timestamp:     entry["timestamp"].(primitive.DateTime).Time(), // Cast to time.Time
		}
		logs = append(logs, log)
	}

	return logs, nil
}
