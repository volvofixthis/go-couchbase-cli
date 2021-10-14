### Сборка  
Собранный бинарник будет лежать в bin:  
make build

### Использование  
Справка:  
go-couchbase-cli kv --help  
go-couchbase-cli kv get --help  
go-couchbase-cli kv upsert --help  
go-couchbase-cli bucket --help  

Апсёрт строкового значения:  
go-couchbase-cli kv --dsn "couchbase://cb2-test.dev.rbc.ru" -u test_v1_billing_in_app_cache -p "eiZoh7chae0ohte5" -b "test_v1_billing_in_app_cache" upsert -k kus3 --ttl 15m -v "just string"  

Запись объекта в базу, поддерживаются операции insert, upsert, replace:  
go-couchbase-cli kv --dsn "couchbase://cb2-test.dev.rbc.ru" -u test_v1_billing_in_app_cache -p "eiZoh7chae0ohte5" -b "test_v1_billing_in_app_cache" upsert --transcoder json -k kus3 --ttl 15m -v '{"first_name": "M", "last_name": "L"}'  

Запись raw json в базу:  
go-couchbase-cli kv --dsn "couchbase://cb2-test.dev.rbc.ru" -u test_v1_billing_in_app_cache -p "eiZoh7chae0ohte5" -b "test_v1_billing_in_app_cache" upsert --transcoder raw_json -k kus3 --ttl 15m -v '{"first_name": "M", "last_name": "L"}'  

Получаем так же, передаём обязательно нужный транскодер:  
go-couchbase-cli kv --dsn "couchbase://cb2-test.dev.rbc.ru" -u test_v1_billing_in_app_cache -p "eiZoh7chae0ohte5" -b "test_v1_billing_in_app_cache" get --transcoder json kus1 kus2 kus3

Удаляем ключи:  
go-couchbase-cli kv --dsn "couchbase://cb2-test.dev.rbc.ru" -u test_v1_billing_in_app_cache -p "eiZoh7chae0ohte5" -b "test_v1_billing_in_app_cache" remove kus1 kus2 kus3

Пингуем ноды:  
go-couchbase-cli bucket --dsn "couchbase://cb2-test.dev.rbc.ru" -u test_v1_billing_in_app_cache -p "eiZoh7chae0ohte5" -b "test_v1_billing_in_app_cache" ping -n 10 --pause 100ms
Получаем диагностику нод:  
go-couchbase-cli bucket --dsn "couchbase://cb2-test.dev.rbc.ru" -u test_v1_billing_in_app_cache -p "eiZoh7chae0ohte5" -b "test_v1_billing_in_app_cache" diagnostics


### TODO
- Причесать приложение  
- Научить получать информацию о расположении ключа  
- Подумать над интерактивным интерфейсом  
- Подумать над stdin интерфейсов  
- Посмотреть что там с контекстом в новых версиях gocb  
- Экспорт данных о латенси нод в csv  
- Добавить флаги для включения уровня дьюрабилити  
- добавить команды для выполнения бенчмарка  
