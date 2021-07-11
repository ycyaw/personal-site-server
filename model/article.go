package model

import (
	"personal-site/log"
	"time"
)

// Article表结构
type Article struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Category    string    `json:"category"`
	Content     string    `json:"content"`
	Reading     int       `json:"reading"`
	ReleaseDate time.Time `json:"releaseDate"`
}

// 像数据库插入文章
func InsertArticle(title string, category string, content string) error {
	sql := "INSERT INTO article_t (title, category, content, reading, releaseDate) VALUES ($1, $2, $3, $4, $5)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		log.Warning(err.Error())
	}
	defer stmt.Close()

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	err = stmt.QueryRow(title, category, content, 0, nowTime).Err()
	if err != nil {
		log.Warning(err.Error())
	}

	return err
}

// 依据id查询表数据
func QueryRowArticle(id int) (Article, error) {
	// 保存查询的数据
	article := Article{}

	// 依据id查询数据
	sql := "SELECT * FROM article_t WHERE id = $1"

	// 填充数据
	err := Db.QueryRow(sql, id).
		Scan(&article.Id, &article.Title, &article.Category, &article.Content, &article.ReleaseDate)

	return article, err
}

// 查询最近的20条数据
func LatestArticle() ([]Article, error) {
	var articles []Article

	rows, err := Db.Query("SELECT * FROM article_t ORDER BY releaseDate DESC Limit 20")
	if err != nil {
		log.Warning(err.Error())
	}

	for rows.Next() {
		article := Article{}

		err = rows.Scan(&article.Id, &article.Title, &article.Category, &article.Content, &article.Reading, &article.ReleaseDate)
		if err != nil {
			log.Warning(err.Error())
		}

		articles = append(articles, article)
	}

	return articles, err
}
