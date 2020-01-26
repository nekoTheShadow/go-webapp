package theasaurus

type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
