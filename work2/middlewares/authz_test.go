// middlewares/authz_test.go

package middlewares

//import (
//	"work2/auth"
//	"work2/controllers"
//	"work2/handlers"
//	"work2/repository/sqlitedb"
//	"work2/models"
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//	"work2/services"
//
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//)
//
////type MockRepo struct{}
//
//func TestAuthzNoHeader(t *testing.T) {
//	//router := gin.Default()
//	//router.Use(Authz())
//	//
//	//router.GET("/api/protected/profile", controllers.Profile)
//	//
//	//w := httptest.NewRecorder()
//	//req, _ := http.NewRequest("GET", "/api/protected/profile", nil)
//	//router.ServeHTTP(w, req)
//	//
//	//assert.Equal(t, 403, w.Code)
//	//
//
//
//	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
//	w := httptest.NewRecorder()
//	//var m *services.MockRepo
//	//tSer := services.TaskNewService(m)
//	//uSer := services.UserNewService(m)
//	//s := handlers.NewServer(tSer, uSer)
//	//func (m foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	//	fmt.Fprintln(w, "Some text")
//	//}
//	//router.ServeHTTP(w, req)
//	//s := &handlers.Server{
//	//	ts: m, us: m,
//	//}
//	s := &handlers.Server{ts: *services.MockRepo, us: *services.UserMockRepo}
//	s.GetTasks(w, req)
//	assert.Equal(t, 403, w.Code)
//}
//
//func TestAuthzInvalidTokenFormat(t *testing.T) {
//	router := gin.Default()
//	router.Use(Authz())
//
//	router.GET("/api/protected/profile", controllers.Profile)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/api/protected/profile", nil)
//	req.Header.Add("Authorization", "test")
//
//	router.ServeHTTP(w, req)
//
//	assert.Equal(t, 400, w.Code)
//}
//
//func TestAuthzInvalidToken(t *testing.T) {
//	invalidToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
//	router := gin.Default()
//	router.Use(Authz())
//
//	router.GET("/api/protected/profile", controllers.Profile)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/api/protected/profile", nil)
//	req.Header.Add("Authorization", invalidToken)
//
//	router.ServeHTTP(w, req)
//
//	assert.Equal(t, 401, w.Code)
//}
//
//func TestValidToken(t *testing.T) {
//	var response models.User
//
//	err := sqlitedb.InitDatabase()
//	assert.NoError(t, err)
//
//	err = sqlitedb.GlobalDB.AutoMigrate(&models.User{})
//	assert.NoError(t, err)

//	user := models.User{
//		Email:    "test@email.com",
//		Password: "secret",
//		Name:     "Test User",
//	}
//
//	jwtWrapper := auth.JwtWrapper{
//		SecretKey:       "verysecretkey",
//		Issuer:          "AuthService",
//		ExpirationHours: 24,
//	}
//
//	token, err := jwtWrapper.GenerateToken(user.Email)
//	assert.NoError(t, err)
//
//	err = user.HashPassword(user.Password)
//	assert.NoError(t, err)
//
//	result := sqlitedb.GlobalDB.Create(&user)
//	assert.NoError(t, result.Error)
//
//	router := gin.Default()
//	router.Use(Authz())
//
//	router.GET("/api/protected/profile", controllers.Profile)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/api/protected/profile", nil)
//	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
//
//	router.ServeHTTP(w, req)
//
//	err = json.Unmarshal(w.Body.Bytes(), &response)
//	assert.NoError(t, err)
//
//	assert.Equal(t, 200, w.Code)
//	assert.Equal(t, "test@email.com", response.Email)
//	assert.Equal(t, "Test User", response.Name)
//
//	sqlitedb.GlobalDB.Unscoped().Where("email = ?", user.Email).Delete(&models.User{})
//}
