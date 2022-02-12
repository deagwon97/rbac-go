package rest

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"

	"net/http"
	"rbac/account/dblayer"
	"rbac/account/models"
	"rbac/database"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateToken(userID int, exp int64, secret string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = exp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
func CreateTokens(userID int) (string, string, error) {
	var exp int64

	os.Setenv("SECRET", "you need to set secret")
	accessSecret := "access" + os.Getenv("SECRET")
	exp = time.Now().Add(time.Hour * 2).Unix()
	accessToken, err := CreateToken(userID, exp, accessSecret)

	refreshSecret := "refresh" + os.Getenv("SECRET")
	exp = time.Now().Add(time.Hour * 24 * 14).Unix()
	refreshToken, err := CreateToken(userID, exp, refreshSecret)

	return accessToken, refreshToken, err
}

func NullStr2Str(str string) (nullStr sql.NullString) {
	if str == "" {
		nullStr.String = ""
		nullStr.Valid = false
	} else {
		nullStr.String = str
		nullStr.Valid = true
	}
	return nullStr
}

// 패스워드 sha256 암호화
func EncodePassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	md := hash.Sum(nil)
	encodedPassword := hex.EncodeToString(md)
	return encodedPassword
}

type Handler struct {
	db dblayer.DBLayer
}

// @BasePath /

// Go API godoc
// @Summary User 생성
// @Schemes
// @Tags Account
// @Accept json
// @Produce json
// @Param data body models.AddUserData true  "회원가입 정보"
// @Success 200 {object} models.LoginResult
// @Router /account [post]
func (h *Handler) AddUser(c *gin.Context) {

	// Data Parse
	var userData models.AddUserData
	err := c.ShouldBindJSON(&userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 사용자 데이터 저장
	var user models.User
	user.LoginID = userData.LoginID
	encodedPassword := EncodePassword(userData.Password)
	user.Password = encodedPassword
	user.Name = NullStr2Str(userData.Name)
	user.Email = NullStr2Str(userData.Email)
	user, err = h.db.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// login
	// token 생성
	accessToken, refreshToken, err := CreateTokens(user.ID)
	var loginResult models.LoginResult
	loginResult.UserID = user.ID
	loginResult.AccessToken = accessToken
	loginResult.RefreshToken = refreshToken

	c.JSON(http.StatusOK, loginResult)
	return
}

// Go API godoc
// @Summary 로그인
// @Schemes
// @Tags Account
// @Accept json
// @Produce json
// @Param data body models.LoginRequest true  "로그인 정보"
// @Success 200 {object} models.LoginResult "access token & refresh token"
// @Router /account/login [post]
func (h *Handler) Login(c *gin.Context) {

	// Data Parse
	var loginRquest models.LoginRequest
	err := c.ShouldBindJSON(&loginRquest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 비밀번호 조회
	var user models.User
	user, err = h.db.GetPassword(loginRquest.LoginID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "존재하지 않는 아이디입니다."})
		return
	}
	encodedPassword := EncodePassword(loginRquest.Password)
	if user.Password == encodedPassword {
		// login
		// token 생성
		var accessToken string
		var refreshToken string
		accessToken, refreshToken, err = CreateTokens(user.ID)
		var loginResult models.LoginResult
		loginResult.UserID = user.ID
		loginResult.AccessToken = accessToken
		loginResult.RefreshToken = refreshToken

		c.JSON(http.StatusOK, loginResult)
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "비밀번호가 존재하지 않습니다."})
		return
	}
}

func DecodeToken(tokenString string, secret string) (bool, jwt.MapClaims, error) {

	Claims := jwt.MapClaims{}

	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ErrUnexpectedSigningMethod := errors.New("unexpected signing method")
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(secret), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, Claims, key)
	var valid bool
	valid = token.Valid
	return valid, Claims, err
}

type accessToken struct {
	AccessToken string `json:"access_token"`
}

// Go API godoc
// @Summary access token 인증
// @Schemes
// @Tags Account
// @Accept json
// @Produce json
// @Param data body accessToken true  "Access Token"
// @Success 200 {object} bool "유효성 검증 결과"
// @Router /account/valid [post]
func (h *Handler) IsValid(c *gin.Context) {

	// Data Parse
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	var data map[string]string
	json.Unmarshal([]byte(value), &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "access_token 항목은 필수입니다."})
		return
	}

	var accesToken string
	accesToken = data["access_token"]
	var valid bool
	accessSecret := "access" + os.Getenv("SECRET")
	valid, _, err = DecodeToken(accesToken, accessSecret)

	if valid == true {
		c.JSON(http.StatusOK, valid)
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 access_token 입니다."})
		return
	}
}

type refreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

// Go API godoc
// @Summary access token 재발급
// @Schemes
// @Tags Account
// @Accept json
// @Produce json
// @Param data body refreshToken true  "Refresh Token"
// @Success 200 {object} string "access token"
// @Router /account/renew [post]
func (h *Handler) RenewToken(c *gin.Context) {

	// Data Parse
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	var data map[string]string
	json.Unmarshal([]byte(value), &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "access_token 항목은 필수입니다."})
		return
	}

	var refreshToken string
	refreshToken = data["refresh_token"]
	var valid bool
	var atClaims jwt.MapClaims
	refreshSecret := "refresh" + os.Getenv("SECRET")
	valid, atClaims, err = DecodeToken(refreshToken, refreshSecret)

	userID, _ := atClaims["user_id"].(int)

	if valid == true {
		var exp int64
		accessSecret := "refresh" + os.Getenv("SECRET")
		exp = time.Now().Add(time.Hour * 2).Unix()
		accessToken, _ := CreateToken(userID, exp, accessSecret)
		c.JSON(http.StatusOK, accessToken)
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 refresh_token 입니다."})
		return
	}
}

type HandlerInterface interface {
	AddUser(c *gin.Context)
	Login(c *gin.Context)
	IsValid(c *gin.Context)
	RenewToken(c *gin.Context)
}

// HandlerInterface의 생성자
func NewHandler() (HandlerInterface, error) {
	dsn := database.DataSource
	// DBORM 초기화
	db, err := dblayer.NewORM("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}
