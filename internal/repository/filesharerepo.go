package repository

type FileshareRepoStruct struct {
	Name string `json:"name"`
}

type FileshareRepository interface {
	DownloadFile(string) error
	DecodeFile(string) (*string, error)
	ListFile() error
}
