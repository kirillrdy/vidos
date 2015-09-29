package downloader

import (
	"github.com/anacrolix/torrent"
	"log"
)

const FileDir = "files/"

var Client *torrent.Client

func init() {

	//TODO In future use this
	//user, err := user.Current()
	//if err != nil {
	//	log.Panic(err)
	//}

	//config := torrent.Config{DataDir: user.HomeDir}
	config := torrent.Config{DataDir: FileDir}
	var err error
	Client, err = torrent.NewClient(&config)

	//TODO pehaps allow client to be nil for the whole app
	if err != nil {
		log.Panic(err)
	}

	//TODO implement proper shutdown, and call Close as part of it
	//defer Client.Close()

	//torrentFile, err := Client.AddMagnet("magnet:?xt=urn:btih:ZOCMZQIPFFW7OLLMIC5HUB6BPCSDEOQU")

	////Return errors to http client
	//if err != nil {
	//	log.Println(err)
	//}

	//<-torrentFile.GotInfo()
	//torrentFile.DownloadAll()

	//go func() {
	//	for {
	//		c.WriteStatus(os.Stdout)
	//		time.Sleep(1 * time.Second)
	//	}
	//}()

	//TODO also possibly part of Client.Close()
	//Client.WaitAll()
	log.Print("torrent downloader init finished")

}
