package main

import (
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

	b := g.StringsRawData.GetValue()
	fmt.Println("Title:", string(get(b, int(g.MetaInstance.Get().GetTitleRef()), byte(0))))
	fmt.Println("Author:", string(get(b, int(g.MetaInstance.Get().GetAuthorRef()), byte(0))))
	fmt.Println("Total characters:", g.VerticesVector.GetSize())

}

func get(data []byte, pos int, s byte) []byte {
	result := make([]byte, 0)
	for i := pos; i < len(data); i++ {
		if data[i] == s {
			break
		}
		result = append(result, data[i])
	}
	return result
}
