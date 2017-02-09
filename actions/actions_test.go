package actions

import (
	"math/rand"
	"testing"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/learn/models"
	"github.com/markbates/pop"
	"github.com/markbates/pop/nulls"
	"github.com/markbates/willie"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var seed *rand.Rand

func init() {
	rand.Seed(time.Now().UnixNano())
	seed = rand.New(rand.NewSource(rand.Int63()))
}

type ActionSuite struct {
	suite.Suite
	*require.Assertions
	DB     *pop.Connection
	Willie *willie.Willie
	app    *buffalo.App
}

func TestActionSuite(t *testing.T) {
	as := new(ActionSuite)
	as.app = App()
	c, err := pop.Connect("test")
	if err != nil {
		t.Fatal(err)
	}
	as.DB = c
	suite.Run(t, as)
}

func (as *ActionSuite) HTML(u string, args ...interface{}) *willie.Request {
	return as.Willie.Request(u, args...)
}

func (as *ActionSuite) JSON(u string, args ...interface{}) *willie.JSON {
	return as.Willie.JSON(u, args...)
}

func (as *ActionSuite) SetupTest() {
	as.DB.MigrateReset("../migrations")
	as.Assertions = require.New(as.T())
	as.Willie = willie.New(as.app)
}

func (as *ActionSuite) Login() (*models.User, error) {
	u, err := createUser(as.DB)
	if err != nil {
		return u, err
	}
	scu := func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			c.Set("current_user", u)
			c.Set("current_user_id", u.ID)
			return next(c)
		}
	}
	as.app.Middleware.Replace(setCurrentUser, scu)
	return u, nil
}

func publicCourse(tx *pop.Connection) (*models.Course, error) {
	c := &models.Course{
		Title:       "My Public Course",
		Description: "Some Description",
		Price:       5000,
		Status:      "public",
	}
	err := tx.Create(c)
	return c, err
}

func privateCourse(tx *pop.Connection) (*models.Course, error) {
	c := &models.Course{
		Title:       "My Private Course",
		Description: "Some Description",
		Price:       5000,
		Status:      "private",
	}
	err := tx.Create(c)
	return c, err
}

func createUser(tx *pop.Connection) (*models.User, error) {
	u := &models.User{
		Name:       "Homer Simpson",
		Email:      nulls.NewString("homer@simpson.com"),
		Provider:   "github",
		ProviderID: "12345",
		AvatarUrl:  nulls.NewString(""),
	}
	err := tx.Create(u)
	return u, err
}
