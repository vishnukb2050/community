package handlers

import (
	"auth-service/config"
	"auth-service/models"
	"auth-service/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SendOTP sends OTP to mobile number via MSG91
func SendOTP(c *gin.Context) {
	var req models.SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate mobile number (basic validation)
	if len(req.Mobile) < 10 || len(req.Mobile) > 15 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mobile number"})
		return
	}

	// Generate 6-digit OTP
	otp := services.GenerateOTP()

	// Store OTP in database
	otpRecord := models.OTPCode{
		Mobile:    req.Mobile,
		OTPCode:   otp,
		ExpiresAt: time.Now().Add(5 * time.Minute),
		Verified:  false,
	}

	if err := config.DB.Create(&otpRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
		return
	}

	// Send OTP via MSG91 (Async for dev bypass)
	cfg := config.Load()
	go services.SendSMS(cfg, req.Mobile, otp)

	c.JSON(http.StatusOK, gin.H{
		"message":            "OTP sent successfully (DEV MODE)",
		"otp":                otp, // Include OTP for testing bypass
		"expires_in_seconds": 300,
	})
}

// VerifyOTP verifies OTP and returns JWT token
func VerifyOTP(c *gin.Context) {
	var req models.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Find OTP record
	var otpRecord models.OTPCode
	if err := config.DB.Where("mobile = ? AND otp_code = ? AND verified = false AND expires_at > ?",
		req.Mobile, req.OTP, time.Now()).First(&otpRecord).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	// Mark OTP as verified
	otpRecord.Verified = true
	config.DB.Save(&otpRecord)

	// Check if user exists
	var user models.User
	result := config.DB.Where("mobile = ?", req.Mobile).First(&user)

	// If user doesn't exist, create new user
	if result.Error != nil {
		user = models.User{
			Mobile:            req.Mobile,
			AllowChatRequests: true,
			HidePhoneNumber:   false,
		}
		if err := config.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	// Generate JWT token
	cfg := config.Load()
	token, err := services.GenerateJWT(user.ID.String(), cfg.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Generate refresh token
	refreshToken, err := services.GenerateRefreshToken(user.ID.String(), cfg.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":     user.ID,
			"mobile": user.Mobile,
			"name":   user.Name,
			"email":  user.Email,
		},
	})
}

// RefreshToken generates a new access token from refresh token
func RefreshToken(c *gin.Context) {
	var req models.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	cfg := config.Load()
	userID, err := services.ValidateRefreshToken(req.RefreshToken, cfg.JWTSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Generate new access token
	token, err := services.GenerateJWT(userID, cfg.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Logout invalidates the current token
func Logout(c *gin.Context) {
	// Token invalidation can be implemented using Redis blacklist
	// For now, just return success (client should delete token)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

// HealthCheck endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "auth-service",
		"time":    time.Now(),
	})
}
