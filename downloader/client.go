package downloader

import (
	"github.com/anacrolix/torrent"
	"log"
)

//FilesDir is where all complete files are stored,
// not to be confused with videos/
const FileDir = "files/"

//Client is a pointer to
var Client *torrent.Client

func init() {

	//TODO torrent should download to own dir and move files once completed
	config := torrent.Config{DataDir: FileDir}
	var err error
	Client, err = torrent.NewClient(&config)

	//TODO pehaps allow client to be nil for the whole app
	if err != nil {
		log.Panic(err)
	}

	//TODO implement proper shutdown, and call Close as part of it
	//defer Client.Close()

	//TODO also possibly part of Client.Close()
	//Client.WaitAll()
}
