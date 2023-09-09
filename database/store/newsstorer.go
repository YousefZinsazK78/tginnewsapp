package store

import (
	"context"
	"errors"
	"log"

	"github.com/yousefzinsazk78/go_news_app/database"
	"github.com/yousefzinsazk78/go_news_app/models"
)

type NewsStorer interface {
	InsertNews(*models.Post) error
	GetNews() ([]*models.Post, error)
	// GetNewsByTitle() ([]*models.Post, error)
}

type NewsStore struct {
	database.Database
}

func NewNewsStore(database database.Database) *NewsStore {
	return &NewsStore{
		Database: database,
	}
}

func (n *NewsStore) InsertNews(newsPost *models.Post) error {
	stmt, err := n.DB.PrepareContext(context.Background(), "INSERT INTO NEWSPOST_TBL (title,description, createdAt) VALUES ($1,$2,$3);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(newsPost.Title, newsPost.Description, newsPost.CreatedAt)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("insert operation not successfull!")
	}
	return nil
}

func (n *NewsStore) GetNews() ([]*models.Post, error) {
	rows, err := n.DB.Query("SELECT * FROM public.newspost_tbl;")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var newsPosts []*models.Post
	for rows.Next() {
		var newsPost models.Post
		err := rows.Scan(&newsPost.NEWSID, &newsPost.Title, &newsPost.Description, &newsPost.CreatedAt, &newsPost.UpdatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		newsPosts = append(newsPosts, &newsPost)
	}
	log.Println(newsPosts)
	return newsPosts, nil
}
