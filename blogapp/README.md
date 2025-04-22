blogapp/
├── backend/
│   ├── handlers/
│   │   └── blog.go
│   ├── models/
│   │   └── post.go
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── index.html
│   ├── css/
│   │   └── styles.css
│   └── js/
│       └── app.js
├── Dockerfile
└── docker-compose.yml


```
cd backend
sudo docker-compose up --build
```

frontned port 3000