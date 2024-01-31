package repository

type FileshareRepository interface {
	DownloadFile(string) error
}
