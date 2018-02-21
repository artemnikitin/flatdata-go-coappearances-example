package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/artemnikitin/flatdata-go-coappearances-example/coappearances"
	"github.com/heremaps/flatdata/flatdata-go/flatdata"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	g, err := coappearances.OpenGraphArchive(flatdata.NewFileResourceStorage("flatdata/Graph.archive"))
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	fmt.Println(prettyPrint(g.ToString()))
}

func prettyPrint(s string) string {
	out := bytes.Buffer{}
	err := json.Indent(&out, []byte(s), "", "\t")
	if err != nil {
		return ""
	}
	return out.String()
}
