package rest

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	ce "rbac-go/common/error"
	"time"

	"net/http"
	"rbac-go/account/dblayer"
	"rbac-go/account/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateToken(userID int, exp int64, secret string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["Authorized"] = true
	atClaims["UserID"] = userID
	atClaims["Exp"] = exp
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
	accessToken, _ := CreateToken(userID, exp, accessSecret)

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
	accessToken, refreshToken, _ := CreateTokens(user.ID)
	var loginResult models.LoginResult
	loginResult.UserID = user.ID
	loginResult.AccessToken = accessToken
	loginResult.RefreshToken = refreshToken

	c.JSON(http.StatusOK, loginResult)
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
		accessToken, refreshToken, _ = CreateTokens(user.ID)
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

	valid := token.Valid
	return valid, Claims, err
}

type AccessToken struct {
	AccessToken string `json:"AccessToken"`
}

// Go API godoc
// @Summary access token 인증
// @Schemes
// @Tags Account
// @Accept json
// @Produce json
// @Param data body AccessToken true  "Access Token"
// @Success 200 {object} bool "유효성 검증 결과"
// @Router /account/valid [post]
func (h *Handler) IsValid(c *gin.Context) {

	// Data Parse
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	var data map[string]string
	json.Unmarshal([]byte(value), &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AccessToken 항목은 필수입니다."})
		return
	}

	accesToken := data["accesToken"]
	var valid bool
	accessSecret := "access" + os.Getenv("SECRET")
	valid, _, _ = DecodeToken(accesToken, accessSecret)

	if valid {
		c.JSON(http.StatusOK, valid)
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 access_token 입니다."})
		return
	}
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

// Go API godoc
// @Summary access token 재발급
// @Schemes
// @Tags Account
// @Accept json
// @Produce json
// @Param data body RefreshToken true  "Refresh Token"
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

	refreshToken := data["refreshToken"]
	refreshSecret := "refresh" + os.Getenv("SECRET")
	valid, atClaims, _ := DecodeToken(refreshToken, refreshSecret)

	userID, _ := atClaims["user_id"].(int)

	if valid {
		var exp int64
		accessSecret := "refresh" + os.Getenv("SECRET")
		exp = time.Now().Add(time.Hour * 2).Unix()
		accessToken, _ := CreateToken(userID, exp, accessSecret)
		c.JSON(http.StatusOK, accessToken)
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 refreshToken 입니다."})
		return
	}
}

type UserIDList struct {
	IDList []int `json:"IDList"`
}

// @Summary 사용자 이름 목록 조회
// @Tags Account
// @Accept json
// @Produce json
// @Param data body UserIDList true "Data"
// @Success 200 {object} dblayer.UserIDName
// @Router /account/name/list [post]
func (h *Handler) GetUserListName(c *gin.Context) {
	var userIDList UserIDList

	err := c.ShouldBindJSON(&userIDList)
	if ce.GinError(c, err) {
		return
	}

	userIDName, err := h.db.GetUserListName(userIDList.IDList)
	if ce.GinError(c, err) {
		return
	}
	c.JSON(http.StatusOK, userIDName)
}

type HandlerInterface interface {
	AddUser(c *gin.Context)
	Login(c *gin.Context)
	IsValid(c *gin.Context)
	RenewToken(c *gin.Context)
	GetUserListName(c *gin.Context)
}

func NewHandler() (HandlerInterface, error) {

	// DBORM 초기화
	db, err := dblayer.NewORM()
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}
