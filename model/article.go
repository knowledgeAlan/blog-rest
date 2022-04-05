package model

import (
	"blog-rest/dbManage"
	"blog-rest/utils"
	"fmt"
)

/**

{
    "title": "Hello World",
    "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
    "author": "John",
}
**/

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
}

// add a new article
func AddArticle(title string, content string,author string)(*article,error) {

	id := utils.GenSonyflake()
	article := article{ID:1, Title: title, Content: content,Author:author}
	sqlString := "insert article(id, title, content, author) VALUE (?,?,?,?)"
	result, error :=dbManage.DB.Exec(sqlString,id,title,content,author)
	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Println("result:",result)
	return &article,nil
}
