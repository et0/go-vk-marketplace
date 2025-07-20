# Go VK marketplace

**Ссылка на описание задачи:** https://github.com/et0/go-vk-marketplace/blob/master/docs/task.md

## API Endpoints

### Регистрация пользователя
```
POST /signup
Content-Type: application/json

{
  "username": "testuser",
  "password": "testpassword"
}
```

### Авторизация пользователя
```
POST /login
Content-Type: application/json

{
  "username": "testuser",
  "password": "testpassword"
}
```

### Создание объявления (требуется авторизация)
```
POST /ads
Content-Type: application/json
Authorization: Bearer <your_token>

{
  "title": "Test Ads",
  "description": "This is a test ads",
  "image_url": "https://example.com/image.jpg",
  "price": 999
}
```

### Получение списка объявлений 
```
GET /ads
GET /ads?page=1&limit=10&sort_by=price&order=asc&min_price=50&max_price=200
```

## Проверка работы

1. Зарегистрируйте пользователя:
``` bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"testpassword"}'
```

2. Авторизуйтесь:
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"testpassword"}'
```

3. Создайте объявление (используйте токен из предыдущего шага):
```bash
curl -X POST http://localhost:8080/ads \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{"title":"Test Ads","description":"This is a test ads","image_url":"https://example.com/image.jpg","price":999}'
```

4. Получите список объявлений:
```bash
curl http://localhost:8080/ads
```

