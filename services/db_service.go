package services

import (
	"quote-server/types"
	"quote-server/utils"
)

type DBService interface {
	InsertUser(user types.UserModel) error
	GetUserById(id int32) (types.UserModel, error)
}

type DbServiceImpl struct {
	DBUtil utils.DBClient
}

func NewDbService(dbUtil utils.DBClient) DBService {
	return &DbServiceImpl{DBUtil: dbUtil}
}

func (s *DbServiceImpl) InsertUser(user types.UserModel) error {
	query := "INSERT INTO users (username, email) VALUES ($1, $2)"
	err := s.DBUtil.Exec(query, user.Username, user.Email)
	return err
}

func (s *DbServiceImpl) GetUserById(id int32) (types.UserModel, error) {
	query := "SELECT id, username, email FROM users WHERE id = $1"
	row := s.DBUtil.Query(query, id)
	var result types.UserModel
	err := row.Scan(&result.ID, &result.Username, &result.Email)
	if err != nil {
		return result, err
	}
	return result, nil
}
