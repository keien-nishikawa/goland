package blog

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

