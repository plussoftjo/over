// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"
	"server/vendors"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// StoreUserRoles ..
func StoreUserRoles(c *gin.Context) {
	var role models.Roles
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"role": role,
	})
}

// IndexUserRoles ...
func IndexUserRoles(c *gin.Context) {
	var roles []models.Roles
	config.DB.Find(&roles)

	c.JSON(200, gin.H{
		"roles": roles,
	})
}

// UpdateUserRole ...
func UpdateUserRole(c *gin.Context) {
	var role models.Roles
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&role).Update(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var roles []models.Roles
	config.DB.Find(&roles)

	c.JSON(200, gin.H{
		"role":  role,
		"roles": roles,
	})
}

// DeleteUserRole ...
func DeleteUserRole(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Roles{}, ID)
	var roles []models.Roles
	config.DB.Find(&roles)
	c.JSON(200, gin.H{
		"roles": roles,
	})
}

// HR Controller ...

// StoreEmployeeType ..
type StoreEmployeeType struct {
	User models.User `json:"user"`
}

// StoreEmployee ..
func StoreEmployee(c *gin.Context) {
	var employeeType StoreEmployeeType
	if err := c.ShouldBindJSON(&employeeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := employeeType.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// StoreEmployee

	token, err := vendors.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var users []models.User
	config.DB.Preload("Roles").Where("roles_id != ?", 0).Find(&users)
	config.DB.Preload("Roles").Where("id = ?", user.ID).First(&user)

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token, "users": users})
}

// IndexEmployee ..
func IndexEmployee(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Roles").Where("roles_id != ?", 0).Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// DeleteEmployee ...
func DeleteEmployee(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.User{}, ID)
	var users []models.User
	config.DB.Preload("Roles").Where("roles_id != ?", 0).Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}

// UpdateEmployee ...
func UpdateEmployee(c *gin.Context) {
	var employeeType StoreEmployeeType
	if err := c.ShouldBindJSON(&employeeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := employeeType.User

	if user.Password == "" {
		config.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(models.User{
			Name:    user.Name,
			Phone:   user.Phone,
			RolesID: user.RolesID,
		})
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		config.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(models.User{
			Name:     user.Name,
			Phone:    user.Phone,
			RolesID:  user.RolesID,
			Password: string(hashedPassword),
		})
	}

	var users []models.User
	config.DB.Preload("Roles").Where("roles_id != ?", 0).Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}
