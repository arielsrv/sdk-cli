package services

type TreeService interface {
	WalkDir(sourceDir string, pattern string, name string) error
}
