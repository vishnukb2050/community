package handlers

import (
	"api-gateway/config"
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HealthCheck endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "api-gateway",
		"time":    time.Now(),
	})
}

// Generic proxy function
func proxyRequest(c *gin.Context, targetURL string) {

	// Build target URL
	fullURL := targetURL + c.Request.URL.Path

	// Copy body
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Create request
	req, err := http.NewRequest(c.Request.Method, fullURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create proxy request"})
		return
	}

	// Copy headers
	req.Header = c.Request.Header.Clone()

	// Forward user ID if available
	if userID, exists := c.Get("user_id"); exists {
		req.Header.Set("X-User-ID", userID.(string))
	}

	// Make request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Service unavailable"})
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	// Copy response body
	c.Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}

// Service-specific proxy handlers
func ProxyToAuth(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.AuthServiceURL)
}

func ProxyToProfile(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.ProfileServiceURL)
}

func ProxyToExpense(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.ExpenseServiceURL)
}

func ProxyToReminder(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.ReminderServiceURL)
}

func ProxyToNotes(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.NotesServiceURL)
}

func ProxyToScanner(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.ScannerServiceURL)
}

func ProxyToDocument(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.DocumentServiceURL)
}

func ProxyToCalendar(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.CalendarServiceURL)
}

func ProxyToCommunity(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.CommunityServiceURL)
}

func ProxyToNotice(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.NoticeServiceURL)
}

func ProxyToPoll(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.PollServiceURL)
}

func ProxyToMeeting(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.MeetingServiceURL)
}

func ProxyToChat(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.ChatServiceURL)
}

func ProxyToNotification(c *gin.Context) {
	cfg := config.Load()
	proxyRequest(c, cfg.NotificationServiceURL)
}

// ProxyChatWebSocket proxies WebSocket connections to chat service
func ProxyChatWebSocket(c *gin.Context) {
	cfg := config.Load()

	// Upgrade HTTP connection to WebSocket
	clientConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer clientConn.Close()

	// Connect to chat service WebSocket
	chatWS := cfg.ChatServiceURL + "/ws/chat"
	chatWS = "ws" + chatWS[4:] // Convert http to ws

	serverConn, _, err := websocket.DefaultDialer.Dial(chatWS, nil)
	if err != nil {
		return
	}
	defer serverConn.Close()

	// Bidirectional proxy
	go func() {
		for {
			messageType, message, err := clientConn.ReadMessage()
			if err != nil {
				return
			}
			serverConn.WriteMessage(messageType, message)
		}
	}()

	for {
		messageType, message, err := serverConn.ReadMessage()
		if err != nil {
			return
		}
		clientConn.WriteMessage(messageType, message)
	}
}
