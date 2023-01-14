package model

import (
	"errors"
	e "go-rest-api/entity"
)

var albums = []e.Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbumById(albumID int) (e.Album, error) {
	for _, album := range albums {
		if album.ID == int64(albumID) {
			return album, nil
		}
	}
	return e.Album{}, errors.New("no album")
}

func GetAllAlbums() ([]e.Album, error) {
	return albums, nil
}

func CountAlbums() (int, error) {
	return len(albums), nil
}

func InsertAlbum(album *e.Album) error {
	albums = append(albums, *album)
	return nil
}

func DeleteAlbum(albumID int) error {
	for idx, album := range albums {
		if album.ID == int64(albumID) {
			albums[idx] = albums[len(albums)-1]
			albums = albums[:len(albums)-1]
			return nil
		}
	}
	return errors.New("no album")
}

func UpdateAlbum(renewAlbum *e.Album) error {
	for idx, album := range albums {
		if album.ID == renewAlbum.ID {
			albums[idx] = e.Album{
				ID:     renewAlbum.ID,
				Title:  renewAlbum.Title,
				Artist: renewAlbum.Artist,
				Price:  renewAlbum.Price,
			}
			return nil
		}
	}
	return nil
}
