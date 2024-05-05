package pwd

import (
    "os"
)


type DirectoryRetriever interface {
    Getwd() (string, error)
}


type DefaultDirectoryRetriever struct{}

func (dr DefaultDirectoryRetriever) Getwd() (string, error) {
    return os.Getwd()
}


func CurrentWorkingDirectory(retriever DirectoryRetriever) (string, error) {
    return retriever.Getwd()
}

