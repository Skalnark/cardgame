#! /bin/bash
go mod tidy
go build -o bin/card_game && 
chmod +x bin/card_game &&
bin/card_game