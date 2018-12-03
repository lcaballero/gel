package gel

import (
	"log"
	"os"
)

type Resolver func(string) string
type FileReader func(string) ([]byte, error)

type Inserter struct {
	resolver Resolver
	reader   FileReader
}

func NewInserter(res Resolver, reader FileReader) Inserter {
	return Inserter{
		resolver: res,
		reader:   reader,
	}
}

func (r Inserter) Include(file string) View {
	fqname := r.resolver(file)
	_, err := os.Stat(fqname)
	if os.IsNotExist(err) {
		log.Println("file doesn't exist", fqname)
		return None()
	}
	if err != nil {
		log.Println(err)
		return None()
	}

	b, err := r.reader(fqname)
	if err != nil {
		log.Println("didn't file file to include", file, fqname)
		return None()
	}
	return Text(string(b))
}
