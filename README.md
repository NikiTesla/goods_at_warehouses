# lamoda_test
Тестовое задание для Lamoda Tech. Реализован API с помощью протокола JSON-RPC для работы с товарами на разных складах.

**Что можно делать:**
1. Создавать товар, задавать имя, размер, уникальный код и общее количество товаров на складах (**Goods.Create**)
2. Создавать склад, задавать имя и доступность склада в текущий момент (**Warehouses.Add**)
3. Распределять товары по складам из общего источника (**Goods.Add**)
4. Резерировать некоторое количество товаров на складе, для доставки (**Goods.Reserve**)
5. Отменять резерв некоторого количества товаров (**Goods.CancelReservation**)
6. Узнавать количество доступных для резерва (доставки) товаров на конкретном складе (**Warehouses.GetAmount**)

**Для запуска можно использовать Makefile (make _команда из списка ниже_):**
1. _all_ - собирает бинарный файл из cmd/main.go, запускает контейнер с БД (должен быть предварительно создан) и запускает сервис
2. _run_ - запускает контейнер с БД (должен быть предварительно создан) и запускает сервис без сохранения собранного бинарника
3. _docker_ - удаляет в случае наличия существующий образ приложения, создает новый образ и запускает docker compose
4. _migration-up_ ARGS="_version_" - выполняет миграцию базы данных в докере указанной версии
5. _migration-down_ ARGS="_version_" - отменяет миграцию базы данных в докере указанной версии

Также в проекте предоставлен файл **my_own_postman.py** для тестирования функциональности приложения, 
вызова разных процедур с удобным вводом передаваемых данных

**База данных**
В качестве хранилища данных используется контейнер в docker с postres, который настраивается в файле docker-compose.yml

**В базе данных _lamoda_ созданы три таблицы для хранения информации о:**
1. Товарах (**goods**) - id товара, имя, размер, уникальный код и количество товаров, досутпное в общем источнике
2. Складах (**warehouses**) - id склада, имя и доступность склада
3. Информация о распределении товаров на складах (**warehouse_goods**) - id склада, code товара, количество доступных товаров и количество зарезервированных товаров

## Goods.Create:
Создает новые виды товаров и размещает их в "сортировочном центре" для дальнейшего распределения по складам.
Возвращает список успешно созданных товаров, ошибки создания логгирует в консоль.

Принимает слайс (список) json-структур (поля: name, size, code, amount):
```
args = [{
  "name": "Coffee",
  "size": 1.2,
  "code": 123,
  "amount": 100
},
{
  "name": "Water",
  "size": 5,
  "code": 51,
  "amount": 40
}]
```

## Warehouses.Create
Создает новые склады, с указанием их имени и доступности.
Возвращает список успешно созданных складов, ошибки создания логгирует в консоль.

Принимает слайс (список) json-структур (поля: name, availability):
```
args = [{
  "name": "Main",
  "availability": true
},
{
  "name": "Old",
  "availability": false
}]
```

## Goods.Addreturns 
Перемещает некоторые количества товаров с goodCode из сортировочного центра на склады с id warehouseID,
проверяя доступность складов и наличие необходимого количества товаров в сортировочном центре.
Возвращает список успешно транспортированных товаров, ошибки транспортировки логгируются в консоль

Принимает слайс (список)  json-структур (поля: goodCode, warehouseID, amount):
```
args = [{
        "goodCode": 123,
        "warehouseID": 1,
        "amount": 50
    },
    {
        "goodCode": 51,
        "warehouseID": 1,
        "amount": 10
    }]
```

## Goods.Reserve
Резервирует некоторые количества товаров goodCode на складах с warehouseID
проверяя доступность складов и наличие необходимых количеств товаров на этих складах.
Возвращает список успешно зарезервированных товаров, ошибки резервирования логгируются в консоль

Принимает слайс (список)  json-структур (поля: goodCode, warehouseID, amount):
```
args = [{
        "goodCode": 123,
        "warehouseID": 1,
        "amount": 20
    },
    {
        "goodCode": 51,
        "warehouseID": 1,
        "amount": 10
    }]
```

## Goods.CancelReservation
Отменяет резервирование некоторых количеств товаров goodCode на складах с warehouseID
проверяя доступность складов и наличие необходимых количеств товаров на этих складах.
Возвращает список успешно отмененных резервов товаров, ошибки отмены логгируются в консоль

Принимает слайс (список)  json-структур (поля: goodCode, warehouseID, amount):
```
args = [{
        "goodCode": 123,
        "warehouseID": 1,
        "amount": 15
    },
    {
        "goodCode": 51,
        "warehouseID": 1,
        "amount": 5
    }]
```

## Warehouses.GetAmount
Возвращает количество товаров goodCode, доступных для резервирования на складе с warehouseID.

Принимает json-структуру (поля: goodCode, warehouseID):
```
args = {
  "goodCode": 123,
  "warehouseID": 1
}
// возращает 45 при последовательном вызове методов с представленными данными
```
