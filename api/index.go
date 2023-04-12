package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<a href=\"https://store.steampowered.com/account/ackgift/1414230BF24B4657?redeemer=\'https://store.steampowered.com/account/ackgift/1414230BF24B4657?redeemer=\">happy birthday bud</a>")
}
