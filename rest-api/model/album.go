package model

import (
	"database/sql"
	"errors"
	"go-rest-api/db"
	e "go-rest-api/entity"
)

func GetAlbumById(albumID int) (album e.Album, err error) {
	const query = `SELECT * FROM albums WHERE Id = ?`
	err = db.GetInstance().DBPool().QueryRow(query, albumID).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	if err == sql.ErrNoRows {
		err = errors.New("album is not found")
	}
	return
}

func GetAllAlbums() (albums []e.Album, err error) {
	var album e.Album

	rows, err := db.GetInstance().DBPool().Query(`SELECT * FROM albums order by id desc`)
	// rows, err := db.GetInstance().DBPool().Query(`SELECT * FROM albums order by id desc offset $1 limit $2`, pagination.Offset, pagination.ItemInPage) // param -> pagination e.ResponsePagination
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return
		}
		albums = append(albums, album)
	}

	err = rows.Err()
	return
}

func CountAlbums() (countAlbums int, err error) {
	err = db.GetInstance().DBPool().QueryRow("SELECT count(*) FROM albums").Scan(&countAlbums)

	if err == sql.ErrNoRows {
		countAlbums = 0
		err = nil
	}
	return
}

func InsertAlbum(album *e.Album) error {
	const query = `INSERT INTO albums (title, artist, price) VALUES (?, ?, ?)`
	tx, err := db.GetInstance().DBPool().Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, album.Title, album.Artist, album.Price)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}

func UpdateAlbum(album *e.Album) error {

	const query = `UPDATE albums SET title = ?, artist = ?, price = ? WHERE id = ?`
	tx, err := db.GetInstance().DBPool().Begin()
	if err != nil {
		return err
	}

	res, err := tx.Exec(query, album.Title, album.Artist, album.Price, album.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return errors.New("data is not found")
	}

	if rowsAffected > 1 {
		tx.Rollback()
		return errors.New("Strange behaviour. Total affected is : " + string(rowsAffected))
	}

	tx.Commit()

	return nil

}

func DeleteAlbum(albumID int) error {
	const query = `DELETE FROM albums WHERE id = ?`
	tx, err := db.GetInstance().DBPool().Begin()
	if err != nil {
		return err
	}

	res, err := tx.Exec(query, albumID)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return errors.New("data is not found")
	}

	if rowsAffected > 1 {
		return errors.New("Strange behaviour. Total affected : " + string(rowsAffected))
	}

	tx.Commit()
	return nil

}
