CREATE TABLE IF NOT EXISTS NEWSPOST_TBL(
    news_id SERIAL PRIMARY KEY ,
    title TEXT,
    description TEXT,
    createdAt TIMESTAMP,
    updatedAt TIMESTAMP
);