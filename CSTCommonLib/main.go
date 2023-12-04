package cstcommonlib

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// DetailedResponse is a struct that contains the success status, message,
// data, and status code of a response. It is used to send detailed responses
// throughout the services.
type DetailedResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

// NewDetailedResponse creates a new DetailedResponse with the given data.
func NewDetailedResponse(data interface{}) *DetailedResponse {
	return &DetailedResponse{
		Success: true,
		Message: "OK",
		Data:    data,
		Code:    200,
	}
}

// OK sends a DetailedResponse with the success status, message, data, and
// status code set to the given values.
func (d *DetailedResponse) OK(ctx *gin.Context) {
    ctx.JSON(d.Code, d)
}

// BadRequest sends a DetailedResponse with the success status set to false,
// message set to "Bad Request", and status code set to 400.
func (d *DetailedResponse) BadRequest(ctx *gin.Context, error error) {
    d.Success = false
    d.Message = "Bad Request: " + error.Error()
    d.Code = 400
    ctx.JSON(d.Code, d)
}

// Unauthorized sends a DetailedResponse with the success status set to false,
// message set to "Unauthorized", and status code set to 401.
func (d *DetailedResponse) Unauthorized(ctx *gin.Context, error error) {
    d.Success = false
    d.Message = "Unauthorized: " + error.Error()
    d.Code = 401
    ctx.JSON(d.Code, d)
}

// Forbidden sends a DetailedResponse with the success status set to false,
// message set to "Forbidden", and status code set to 403.
func (d *DetailedResponse) Forbidden(ctx *gin.Context, error error) {
    d.Success = false
    d.Message = "Forbidden: " + error.Error()
    d.Code = 403
    ctx.JSON(d.Code, d)
}

// NotFound sends a DetailedResponse with the success status set to false,
// message set to "Not Found", and status code set to 404.
func (d *DetailedResponse) NotFound(ctx *gin.Context, error error) {
    d.Success = false
    d.Message = "Not Found: " + error.Error()
    d.Code = 404
    ctx.JSON(d.Code, d)
}

// InternalServerError sends a DetailedResponse with the success status set to false,
// message set to "Internal Server Error", and status code set to 500.
func (d *DetailedResponse) InternalServerError(ctx *gin.Context, error error) {
    d.Success = false
    d.Message = "Internal Server Error: " + error.Error()
    d.Code = 500
    ctx.JSON(d.Code, d)
}

// NotImplemented sends a DetailedResponse with the success status set to false,
// message set to "Not Implemented", and status code set to 501.
func (d *DetailedResponse) NotImplemented(ctx *gin.Context, error error) {
    d.Success = false
    d.Message = "Not Implemented: " + error.Error()
    d.Code = 501
    ctx.JSON(d.Code, d)
}

// BoolResponse is a struct that contains the success status, message,
// and status code of a response. It is used to send simple respense
// throughout the services and tpically only says that the request
// was successful or not.
type BoolResponse struct {
    Success bool `json:"success"`
    Code    int  `json:"code"`
    message string `json:"message"`
}

// NewBoolResponse creates a new BoolResponse with the given success status.
func NewBoolResponse() *BoolResponse {
    return &BoolResponse{
        Success: true,
        Code:    200,
    }
}

// OK sends a BoolResponse with the success status and status code set to 
// the given values.
func (b *BoolResponse) OK(ctx *gin.Context) {
    ctx.JSON(b.Code, b)
}

// BadRequest sends a BoolResponse with the success status set to false
// and status code set to 400.
func (b *BoolResponse) BadRequest(ctx *gin.Context, error error) {
    b.Success = false
    b.Code = 400
    ctx.JSON(b.Code, b)
}

// Unauthorized sends a BoolResponse with the success status set to false
// and status code set to 401.
func (b *BoolResponse) Unauthorized(ctx *gin.Context, error error) {
    b.Success = false
    b.Code = 401
    ctx.JSON(b.Code, b)
}

// Forbidden sends a BoolResponse with the success status set to false
// and status code set to 403.
func (b *BoolResponse) Forbidden(ctx *gin.Context, error error) {
    b.Success = false
    b.Code = 403
    ctx.JSON(b.Code, b)
}

// NotFound sends a BoolResponse with the success status set to false
// and status code set to 404.
func (b *BoolResponse) NotFound(ctx *gin.Context, error error) {
    b.Success = false
    b.Code = 404
    ctx.JSON(b.Code, b)
}

// InternalServerError sends a BoolResponse with the success status set to false
// and status code set to 500.
func (b *BoolResponse) InternalServerError(ctx *gin.Context, error error) {
    b.Success = false
    b.Code = 500
    ctx.JSON(b.Code, b)
}

// NotImplemented sends a BoolResponse with the success status set to false
// and status code set to 501.
func (b *BoolResponse) NotImplemented(ctx *gin.Context, error error) {
    b.Success = false
    b.Code = 501
    ctx.JSON(b.Code, b)
}

// ConstructHmacAuthHeader creates a header map with the Bearer HMACEncodedToken
// as the value for the Authorization key. It also has the signature and timestamp
// keys. Then it sets the content type to application/json and returns the map.
// It uses the endpoint and method to create the signature, and the secret key
// is pulled from the environment variable SECRET_KEY.
func ConstructHmacAuthHeader(endpoint, method string) (map[string]string, error) {
	// Get the secret key from the environment variable
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
        return nil, fmt.Errorf("SECRET_KEY environment variable not set")
    }

	// Current timestamp in seconds
	timestamp := time.Now().Unix()

	// Construct the message to be signed
    message := fmt.Sprintf("%s\n%s\n%d", method, endpoint, timestamp)
    println(message)

    // Calculate the HMAC signature
    hmacSignature := calculateHMAC(message, secretKey)

    // Create the header map
    headers := map[string]string{
        "Signature":      hmacSignature,
        "Timestamp":      strconv.FormatInt(timestamp, 10),
        "Content-Type":   "application/json",
    }

    return headers, nil
}

// ValidateHmac validates the HMAC signature in the Authorization header
// of a request. It checks if the received signature matches the expected
// signature calculated using the given method, endpoint, and body.
func ValidateHmac(ctx *gin.Context) error {
    // Get the secret key from the environment variable
    method := ctx.Request.Method
    endpoint := ctx.Request.URL.Path
    receivedSignature := ctx.GetHeader("Signature")
    println(receivedSignature)
    receivedTime, err := strconv.ParseInt(ctx.GetHeader("Timestamp"), 10, 64)
    if err != nil {
        return fmt.Errorf("Error parsing timestamp: %s", err)
    }


    secretKey := os.Getenv("SECRET_KEY")
    if secretKey == "" {
        return fmt.Errorf("SECRET_KEY environment variable not set")
    }

    // Convert the received signature to bytes
    receivedSignatureBytes, err := base64.StdEncoding.DecodeString(receivedSignature)
    if err != nil {
        return fmt.Errorf("Error decoding received signature: %s", err)
    }

    // format the message to be validated with the received time
    message := fmt.Sprintf("%s\n%s\n%d", method, endpoint, receivedTime)
    hmacSignature := calculateHMAC(message, secretKey)
    expectedSignatureBytes, err := base64.StdEncoding.DecodeString(hmacSignature)
    if err != nil {
        return fmt.Errorf("Error decoding expected signature: %s", err)
    }

    // Compare the received signature with the expected signature
    if !hmac.Equal(receivedSignatureBytes, expectedSignatureBytes) {
        return fmt.Errorf("Received signature does not match expected signature")
    }

    return nil
}

// calculateHMAC calculates the HMAC signature for the given message and key
func calculateHMAC(message, key string) string {
    keyBytes := []byte(key)
    messageBytes := []byte(message)

    hmacInstance := hmac.New(sha256.New, keyBytes)
    hmacInstance.Write(messageBytes)
    signature := base64.StdEncoding.EncodeToString(hmacInstance.Sum(nil))

    return signature
}

// CreateIDSuffix creates a suffix for the ID of a resource. it creates a
// random string of 16 characters and returns it.
func CreateIDSuffix() (string, error) {
    // Create a random string of 16 characters
    randomString := make([]byte, 16)
    _, err := rand.Read(randomString)
    if err != nil {
        return "", fmt.Errorf("Error generating random string: %s", err)
    }

    return base64.StdEncoding.EncodeToString(randomString), nil
}
