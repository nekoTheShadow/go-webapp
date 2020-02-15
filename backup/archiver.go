package backup

type Archiver interface {
	Archive(src, dst string) error
}
