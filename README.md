# rssagg

![rss](https://github.com/user-attachments/assets/e9ea0b7f-d586-4b42-a919-94081fa375fd)

> This project was based in [FreeCodeCamp tutorial on Golang](https://www.youtube.com/watch?v=un6ZyFkqFKo)

Description
-

Web Server(**RSS Feed Agregator**) which allows clients to perform the following operations:
-  Add RSS feeds to be collection
-  Follow and unfollow RSS feeds that other users have added
-  Fetch all of the latest posts from the RSS feeds they follow
-  Add users, feeds
-  Delete feeds, feed follows
-  Get users, feeds, feed follows and posts

I used Thunder Client Extension in VS Code to perform the REST Requests above.

Install
-

-  go version go 1.21.3
-  go install github.com/pressly/goose/v3/cmd/goose@latest
-  and other tools which you can find in **Boot.Dev** tutorial


Setup .env file
-

```bash
PORT=8080
DB_URL=
```
Setup DB and generate queries
-

```bash
goose postgres <YOUR_DB_URL> up
sqlc generate
```

Build and start the server
-

```bash
go build && ./rssagg
```
