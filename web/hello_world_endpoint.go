package web

import (
	"fmt"
	"net/http"
	"time"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "♣ Hello, world!\nIt's %s", time.Now())
}
