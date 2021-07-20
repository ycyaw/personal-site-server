package model

import (
	"personal-site/log"
	"time"
)

// Article表结构
type Article struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Category    string    `json:"category"`
	Content     string    `json:"content"`
	Reading     int       `json:"reading"`
	ReleaseDate time.Time `json:"releaseDate"`
}

type ResponseArticle struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Content     string `json:"content"`
	Reading     int    `json:"reading"`
	ReleaseDate string `json:"releaseDate"`
}

// 将Article转换为ResponseArticle
func converArticle(article Article) ResponseArticle {
	// 转换数据封装
	response := ResponseArticle{
		Id:          article.Id,
		Title:       article.Title,
		Author:      article.Author,
		Category:    article.Category,
		Content:     article.Content,
		Reading:     article.Reading,
		ReleaseDate: article.ReleaseDate.Format("2006-01-02 15:04:05"),
	}

	return response
}

// 像数据库插入文章
func InsertArticle(title string, author string, category string, content string) error {
	sql := "INSERT INTO article_t (title, author, category, content, reading, releaseDate) VALUES ($1, $2, $3, $4, $5, $6)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		log.Warning(err.Error())
	}
	defer stmt.Close()

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	err = stmt.QueryRow(title, author, category, content, 0, nowTime).Err()
	if err != nil {
		log.Warning(err.Error())
	}

	return err
}

// 依据id查询表数据
func QueryArticleOfId(id int64) (ResponseArticle, error) {
	// 保存查询的数据
	article := Article{}

	// 依据id查询数据
	sql := "SELECT * FROM article_t WHERE id = $1"

	// 填充数据
	err := Db.QueryRow(sql, id).
		Scan(&article.Id, &article.Title, &article.Author, &article.Category, &article.Content, &article.Reading, &article.ReleaseDate)
	if err != nil {
		log.Info(err.Error())
	}

	// 转换数据封装
	responseArticle := converArticle(article)

	// 将阅读次数+1
	err = Db.QueryRow("UPDATE article_t SET reading = reading + 1 WHERE id = $1", id).Err()
	if err != nil {
		log.Info(err.Error())
	}

	return responseArticle, err
}

// 依据类别查询表数据
func QueryArticleByCategory(category string) ([]ResponseArticle, error) {
	var articles []ResponseArticle

	// 依据类别查询最新的文章，条数为20
	rows, err := Db.Query("SELECT * FROM article_t WHERE category = $1 ORDER BY releaseDate DESC Limit 20", category)
	if err != nil {
		log.Warning(err.Error())
	}
	defer rows.Close()

	// 将数据放入数组中
	for rows.Next() {
		article := Article{}

		err = rows.Scan(&article.Id, &article.Title, &article.Author, &article.Category, &article.Content, &article.Reading, &article.ReleaseDate)
		if err != nil {
			log.Warning(err.Error())
		}

		// 转换数据封装
		responseArticle := converArticle(article)

		articles = append(articles, responseArticle)
	}

	return articles, err
}

// 依据文章标题关键字查询
func QueryArticleOfTitle(title string) ([]ResponseArticle, error) {
	var articles []ResponseArticle

	// 通过文章关键字查询最新的文章
	rows, err := Db.Query("SELECT * FROM article_t WHERE title LIKE '%' || $1 || '%' ORDER BY releaseDate DESC", title)
	if err != nil {
		log.Warning(err.Error())
	}
	defer rows.Close()

	// 将数据放入数组中
	for rows.Next() {
		article := Article{}

		err = rows.Scan(&article.Id, &article.Title, &article.Author, &article.Category, &article.Content, &article.Reading, &article.ReleaseDate)
		if err != nil {
			log.Warning(err.Error())
		}

		// 转换数据封装
		responseArticle := converArticle(article)

		articles = append(articles, responseArticle)
	}

	return articles, err
}

// 查询最近的20条数据
func QueryArticleOfLatest() ([]ResponseArticle, error) {
	var articles []ResponseArticle

	rows, err := Db.Query("SELECT * FROM article_t ORDER BY releaseDate DESC Limit 20")
	if err != nil {
		log.Warning(err.Error())
	}
	defer rows.Close()

	// 将数据放入数组中
	for rows.Next() {
		article := Article{}

		err = rows.Scan(&article.Id, &article.Title, &article.Author, &article.Category, &article.Content, &article.Reading, &article.ReleaseDate)
		if err != nil {
			log.Warning(err.Error())
		}

		// 转换数据封装
		responseArticle := converArticle(article)

		articles = append(articles, responseArticle)
	}

	return articles, err
}

// 通过作者姓名查询文章
func QueryArticleOfName(name string) ([]ResponseArticle, error) {
	// 保存查询的数据
	var responseArticles []ResponseArticle

	// 依据姓名查询数据
	rows, err := Db.Query("SELECT * FROM article_t WHERE author = $1 ORDER BY releaseDate DESC", name)
	if err != nil {
		log.Warning(err.Error())
	}
	defer rows.Close()

	// 填充数据
	for rows.Next() {
		article := Article{}

		// 填充数据
		err = rows.Scan(&article.Id, &article.Title, &article.Author, &article.Category, &article.Content, &article.Reading, &article.ReleaseDate)
		if err != nil {
			log.Warning(err.Error())
		}

		// 转换数据封装，并添加到数组中
		responseArticles = append(responseArticles, converArticle(article))
	}

	return responseArticles, err
}

// 更新指定Id的文章内容
func UpdateArticle(id string, title string, category string, content string) error {
	stmt, err := Db.Prepare("UPDATE article_t SET title = $1, category = $2, content = $3  WHERE id = $4")
	if err != nil {
		log.Warning(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, category, content, id)
	if err != nil {
		log.Warning(err.Error())
	}

	return err
}

// 通过Id删除指定文章
func DeleteArticleOfId(id string, name string) error {
	stmt, err := Db.Prepare("DELETE FROM article_t WHERE id = $1 AND author = $2")
	if err != nil {
		log.Warning(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name)
	if err != nil {
		log.Warning(err.Error())
	}

	return err
}
