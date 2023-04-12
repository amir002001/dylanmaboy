package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "https://store.steampowered.com/account/ackgift/1414230BF24B4657?redeemer=")
}
