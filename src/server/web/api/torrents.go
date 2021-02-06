package api

import (
	"net/http"

	"server/log"
	"server/torr"
	"server/torr/state"
	"server/web/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

//Action: add, get, rem, list, drop
type torrReqJS struct {
	requestI
	Link     string `json:"link,omitempty"`
	Hash     string `json:"hash,omitempty"`
	Title    string `json:"title,omitempty"`
	Poster   string `json:"poster,omitempty"`
	Data     string `json:"data,omitempty"`
	SaveToDB bool   `json:"save_to_db,omitempty"`
}

func torrents(c *gin.Context) {
	var req torrReqJS
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Status(http.StatusBadRequest)
	switch req.Action {
	case "add":
		{
			addTorrent(req, c)
		}
	case "get":
		{
			getTorrent(req, c)
		}
	case "rem":
		{
			remTorrent(req, c)
		}
	case "list":
		{
			listTorrent(req, c)
		}
	case "drop":
		{
			dropTorrent(req, c)
		}
	}
}

func addTorrent(req torrReqJS, c *gin.Context) {
	if req.Link == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("link is empty"))
		return
	}

	log.TLogln("add torrent", req.Link)
	torrSpec, err := utils.ParseLink(req.Link)
	if err != nil {
		log.TLogln("error add torrent:", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tor, err := torr.AddTorrent(torrSpec, req.Title, req.Poster, req.Data)
	if err != nil {
		log.TLogln("error add torrent:", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	go func() {
		if !tor.GotInfo() {
			log.TLogln("error add torrent:", "timeout connection torrent")
			return
		}

		if req.SaveToDB {
			torr.SaveTorrentToDB(tor)
		}

		savePlaylist(tor, c)
	}()

	c.JSON(200, tor.Status())
}

func getTorrent(req torrReqJS, c *gin.Context) {
	if req.Hash == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("hash is empty"))
		return
	}
	tor := torr.GetTorrent(req.Hash)

	if tor != nil {
		st := tor.Status()
		c.JSON(200, st)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func remTorrent(req torrReqJS, c *gin.Context) {
	if req.Hash == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("hash is empty"))
		return
	}
	delPlaylist(req.Hash)
	torr.RemTorrent(req.Hash)
	c.Status(200)
}

func listTorrent(req torrReqJS, c *gin.Context) {
	list := torr.ListTorrent()
	if list == nil {
		c.JSON(200, []*state.TorrentStatus{})
		return
	}
	var stats []*state.TorrentStatus
	for _, tr := range list {
		stats = append(stats, tr.Status())
	}
	c.JSON(200, stats)
}

func dropTorrent(req torrReqJS, c *gin.Context) {
	if req.Hash == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("hash is empty"))
		return
	}
	delPlaylist(req.Hash)
	torr.DropTorrent(req.Hash)
	c.Status(200)
}
