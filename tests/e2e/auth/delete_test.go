package authtest

import (
	"net/http"
	"strconv"
	"testing"

	model "github.com/gateway-address/internal/models"
	testsetup "github.com/gateway-address/tests/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeleteTestSuite struct {
	suite.Suite
	URL  string
	user *model.User
}

func (s *DeleteTestSuite) SetupSuite() {
	s.URL = testsetup.AuthURL
	s.user = testsetup.UserGenerator()
}

func (s *DeleteTestSuite) TestDeleteUser() {
	url := s.URL + strconv.Itoa(s.user.ID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		s.T().Errorf("Error creating request: %v", err)
		return
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		s.T().Errorf("Error sending request: %v", err)
		return
	}
	defer res.Body.Close()

	assert.Equal(s.T(), http.StatusOK, res.StatusCode, "should return 200 for successful deletion")
}

//	func (s *DeleteTestSuite) TearDownDeleteUser() {
//		db, err := GetDB()
//		if err != nil {
//			s.T().Errorf("GetDB: %v", err)
//			return
//		}
//
//		repo := repository.NewRepositorySqlite(db)
//		_, err = repo.GetByID(LoginDetails.ID)
//		if err == nil {
//			return
//		}
//
//		_, err = db.Exec("DELETE FROM users WHERE user_id = $1", LoginDetails.ID)
//		if err != nil {
//			s.T().Errorf("Error deleting user: %v", err)
//		}
//	}
func TestDeleteTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}
