package main

import (
	"net/http"
	"strings"

	"github.com/alwindoss/morse"
	"github.com/gin-gonic/gin"
)

func decodeString2Morse(palabbra string) (string, error) {

	h := morse.NewHacker()

	morseCode, err := h.Decode(strings.NewReader(palabbra))
	if err != nil {
		return "", err
	}
	return string(morseCode), nil
}

func transalate2Human(palabbra string) (string, error) {

	h := morse.NewHacker()

	morseCode, err := h.Encode(strings.NewReader(palabbra))
	if err != nil {
		return "", err
	}
	return string(morseCode), nil
}

func getMorse(context *gin.Context) {
	id := context.Param("mor")
	word, err := decodeString2Morse(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Is not possible to change"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"To morse": word})
}

func getString(context *gin.Context) {
	id := context.Param("mor")
	word, err := transalate2Human(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Is not possible to change"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"To String": word})
}
func main() {
	router := gin.Default()
	router.POST("/fromMorse/:mor", getMorse)
	router.POST("/fromString/:mor", getString)
	router.Run()
}
