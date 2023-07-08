package models

import (
	"time"
)

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "テストコメント1",
		CreatedAt: time.Now(),
	}

	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "テストコメント2",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "テスト記事1",
		Contents:    "テスト記事1の内容",
		UserName:    "テストユーザー1",
		NiceNum:     10,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}

	Article2 = Article{
		ID:        2,
		Title:     "テスト記事2",
		Contents:  "テスト記事2の内容",
		UserName:  "テストユーザー2",
		NiceNum:   20,
		CreatedAt: time.Now(),
	}
)
