# Микросервис для работы с балансом пользователей

## Запуск
В директории balance выполнить</br>

`cd main`</br>
`docker-compose up`</br>

Сервер будет доступен по адресу http://localhost:8000
## Запросы к API
Postman коллекция: https://www.getpostman.com/collections/3f1e8fd92c9861a52ed3
### Метод начисления средств на баланс
```
curl --location --request POST 'http://localhost:8000/receipt' \
--data-raw '{
    "user_id": "1",
    "income": 20.0,
    "source_id": 10,
    "comment": "Happy Birthday"
}'
```
Response
`"Succesful"`
### Метод получения баланса пользователя
`curl --location --request GET 'http://localhost:8080/1/balance'`</br>

Response </br>
```
{
  "balance":28
}
 ```
 ### Метод резервирования средств с основного баланса на отдельном счете
 
 ```
 curl --location --request POST 'http://localhost:8000/reserve' \
--data-raw '{
    "user_id": "1",
    "service_id": 5,
    "order_id": 4,
    "price": 2,
    "comment": "Coffee"
}'
```

Response</br>
`"Reservation successful"`
###  Метод признания выручки
Cписывает из резерва деньги, добавляет данные в отчет для бухгалтерии</br>
```
curl --location --request POST 'http://localhost:8000/accept' \
--data-raw '{
    "user_id": "1",
    "service_id": 5,
    "order_id": 5,
    "price": 2
}'
```

Response</br>
`"Reservation verified"`
### Метод получения месячного отчета для бухгалтерии
```
curl --location --request POST 'http://localhost:8080/report' \
--data-raw '{
    "year": "2022",
    "month": "11"
}'
```

Response</br>
`"reports/2022-11.csv"`</br>

Дикректория reports находится в корне проекта
### Метод получения списка транзакций пльзователя
`curl --location --request GET 'http://localhost:8080/1/transactions?limit=10&offset=1&sort='time''`</br>

Response
```
"2022-11-13 19:44:22: Income 20 from 10 with comment: 'Happy Birthday' ",
"2022-11-13 19:55:15: Debited 2 to 5 with comment: 'Coffee' "
```
