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
	cg, err := coappearances.OpenGraphArchive(flatdata.NewFileResourceStorage("flatdata/Graph.archive"))
	if err != nil {
		fmt.Println(err)
	}
	defer cg.Close()

	showBasicInfo(cg)
	showData(cg)
	//showCoappearances(cg)
}

func showBasicInfo(g *coappearances.GraphArchive) {
	b := g.StringsRawData.GetValue()
	fmt.Println("Title:", get(b, int(g.MetaInstance.Get().GetTitleRef())))
	fmt.Println("Author:", get(b, int(g.MetaInstance.Get().GetAuthorRef())))
	fmt.Println("Total characters:", g.VerticesVector.GetSize())
	fmt.Println()
}

func showData(g *coappearances.GraphArchive) {
	b := g.StringsRawData.GetValue()
	for i := 0; i < g.VerticesVector.GetSize(); i++ {
		char := g.VerticesVector.Get(i)
		fmt.Println("Character:", get(b, int(char.GetNameRef())))
		data := g.VerticesDataMultivector.Get(i)
		for _, v := range data {
			d1, ok := v.(*coappearances.Nickname)
			if ok {
				fmt.Println("Nickname:", get(b, int(d1.GetRef())))
			}
			d2, ok := v.(*coappearances.Description)
			if ok {
				fmt.Println("Description:", get(b, int(d2.GetRef())))
			}
			d3, ok := v.(*coappearances.UnaryRelation)
			if ok {
				fmt.Println("Relation:")
				fmt.Println("kind:", get(b, int(d3.GetKindRef())))
				fmt.Println("to:", get(b, int(g.VerticesVector.Get(int(d3.GetToRef())).GetNameRef())))
			}
			d4, ok := v.(*coappearances.BinaryRelation)
			if ok {
				fmt.Println("Relation:")
				fmt.Println("kind:", get(b, int(d4.GetKindRef())))
				fmt.Println("to:", get(b, int(g.VerticesVector.Get(int(d4.GetToARef())).GetNameRef())))
				fmt.Println("to:", get(b, int(g.VerticesVector.Get(int(d4.GetToBRef())).GetNameRef())))
			}
		}
		fmt.Println("====================")
	}
}

func showCoappearances(g *coappearances.GraphArchive) {
	b := g.StringsRawData.GetValue()
	fmt.Println("Coappearances:", g.EdgesVector.GetSize())
	// Skip the last edge since it is a sentinel
	for i := 0; i+1 < g.EdgesVector.GetSize(); i++ {
		edge := g.EdgesVector.Get(i)
		fmt.Println(fmt.Sprintf("%s meets %s %d times in chapters:",
			get(b, int(g.VerticesVector.Get(int(edge.GetARef())).GetNameRef())),
			get(b, int(g.VerticesVector.Get(int(edge.GetBRef())).GetNameRef())),
			edge.GetCount()))
		// The end of the chapters assigned to this edge is the first chapter from the next edge.
		// This is a typical trick when storing ranges. That's why a sentinel was added to edges.
		chapterBegin := edge.GetFirstChapterRef()
		chapterSize := g.EdgesVector.Get(i+1).GetFirstChapterRef() - chapterBegin
		chapters := g.ChaptersVector.GetSlice(int(chapterBegin), int(chapterSize), 1)
		for j := 0; j < len(chapters); j++ {
			fmt.Print(fmt.Sprintf("%d.%d, ", chapters[j].GetMajor(), chapters[j].GetMinor()))
		}
		fmt.Println()
	}
}

func get(data []byte, pos int) string {
	result := make([]byte, 0)
	for i := pos; i < len(data); i++ {
		if data[i] == byte(0) {
			break
		}
		result = append(result, data[i])
	}
	return string(result)
}
