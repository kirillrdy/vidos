package downloader

import (
	"github.com/anacrolix/torrent"
	"github.com/kirillrdy/vidos/util"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var downloadsDir = util.VidosDataDirFor("downloads")

//FilesDir is where files go once they are downloaded
var FilesDir = util.VidosDataDirFor("files")

//Client is a pointer to
var Client *torrent.Client

func init() {

	config := torrent.NewDefaultClientConfig()
	config.DataDir = downloadsDir

	var err error
	Client, err = torrent.NewClient(config)

	//TODO pehaps allow client to be nil for the whole app
	if err != nil {
		log.Panic(err)
	}

	//TODO implement proper shutdown, and call Close as part of it
	//defer Client.Close()

	//TODO also possibly part of Client.Close()
	//Client.WaitAll()
	go downloadsFileMover()
}

func downloadsFileMover() {
	for {
		moveAllCompletedTorrentsToFiles()

		//TODO perhaps thats too much
		time.Sleep(1 * time.Minute)
	}
}

func allTorrentsCompleted() bool {

	allCompleted := true

	for _, torrent := range Client.Torrents() {
		allCompleted = torrent.Length() == torrent.BytesCompleted() && allCompleted
	}
	return allCompleted
}

//TODO reimplement so that you dont have to wait for ALL torrents to be completed
func moveAllCompletedTorrentsToFiles() {

	if len(Client.Torrents()) > 0 && allTorrentsCompleted() {
		for _, torrent := range Client.Torrents() {
			torrent.Drop()
		}

		items, err := ioutil.ReadDir(downloadsDir)
		util.LogError(err)
		for _, item := range items {
			origin := downloadsDir + string(os.PathSeparator) + item.Name()
			dest := FilesDir + string(os.PathSeparator) + item.Name()
			log.Printf("moving %#v to %#v", origin, dest)
			err := os.Rename(origin, dest)
			util.LogError(err)
		}

	}
}
