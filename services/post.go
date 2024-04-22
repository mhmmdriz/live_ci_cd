package services

import (
	"ci_cd/database"
	"ci_cd/models"
	"context"
)

func GetAllPosts(ctx context.Context) ([]models.Post, error) {
	data, err := database.DB.QueryContext(ctx, "SELECT * FROM posts")

	if err != nil {
		return []models.Post{}, err
	}

	posts := []models.Post{}

	for data.Next() {
		post := models.Post{}

		err = data.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			return []models.Post{}, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func CreatePost(ctx context.Context, input models.PostRequest) (models.Post, error) {
	stmt, err := database.DB.PrepareContext(ctx, "INSERT INTO posts(title,content) VALUES(?,?)")

	if err != nil {
		return models.Post{}, err
	}

	res, err := stmt.ExecContext(ctx, input.Title, input.Content)

	if err != nil {
		return models.Post{}, err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		return models.Post{}, err
	}

	data, err := database.DB.QueryContext(ctx, "SELECT * FROM posts WHERE id=?", lastID)

	if err != nil {
		return models.Post{}, err
	}

	post := models.Post{}

	for data.Next() {
		err = data.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}
