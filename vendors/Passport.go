// Package vendors ..
package vendors

import (
	"fmt"
	"server/config"
	"server/models"
	"time"

	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// AuthClientSecret Secret String
var AuthClientSecret string

// AuthClientID ...
var AuthClientID uint

// CreateToken ...
func CreateToken(userid uint) (string, error) {
	var err error

	// Generate Access Token
	os.Setenv("ACCESS_SECRET", AuthClientSecret)
	atClaims := jwt.MapClaims{}
	atClaims["authorizes"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 8640).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(AuthClientSecret))
	if err != nil {
		return "", err
	}

	authToken := models.AuthTokens{
		UserID:   userid,
		ClientID: AuthClientID,
		Token:    token,
	}
	config.DB.Create(&authToken)

	return token, nil
}

// SetupPassport ...
func SetupPassport() {
	/**
	* Setup passport method
	* Is about check the auth clients have or not
	* If Not have make new one
	 */
	var authClients models.AuthClients
	if err := config.DB.First(&authClients).Error; err != nil {
		// Register New One
		MakeNewClient()
		return
	}

	FetchAuthClientSecret()

}

// MakeNewClient ...
func MakeNewClient() {
	// Make New Clients
	Name := "Server"
	Secret := randSeq(50)
	Active := 1

	// Register Client
	authClient := models.AuthClients{
		Name:   Name,
		Secret: Secret,
		Active: Active,
	}

	// Create The Client with the DB
	config.DB.Create(&authClient)

	FetchAuthClientSecret()

}

// FetchAuthClientSecret ...
func FetchAuthClientSecret() {
	var authClient models.AuthClients
	config.DB.Where("name = ?", "Server").First(&authClient)

	AuthClientSecret = authClient.Name
	AuthClientID = authClient.ID
}

// VerifyToken ...
func VerifyToken(tok string) (uint, error) {
	var authToken models.AuthTokens
	if err := config.DB.Where("token = ?", tok).First(&authToken).Error; err != nil {
		return 0, err
	}

	userID := authToken.UserID

	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

// CreateAdmin ..
func CreateAdmin() {
	var rolesCount int64
	config.DB.Model(&models.Roles{}).Where("title = ?", "Admin").Count(&rolesCount)

	if rolesCount == 0 {
		config.DB.Create(&models.Roles{
			Title:  "Admin",
			Pages:  "",
			Scopes: "",
		})

		var role models.Roles
		config.DB.Where("title = ?", "Admin").Find(&role)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("00962"), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("error" + err.Error())
		}

		config.DB.Create(&models.User{
			Name:     "Admin",
			Phone:    "00962",
			Password: string(hashedPassword),
			RolesID:  role.ID,
		})

	}
}
