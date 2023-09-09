package handlers

import "github.com/yousefzinsazk78/go_news_app/database/store"

type handlers struct {
	store store.NewsStorer
}

func NewHandler(newsStore store.NewsStorer) handlers {
	return handlers{
		store: newsStore,
	}
}
