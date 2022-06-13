### Build  
building binary, ready to use application can be located in bin folder:  
make build

### Usage  
Get help:  
go-couchbase-cli kv --help  
go-couchbase-cli kv get --help  
go-couchbase-cli kv upsert --help  
go-couchbase-cli bucket --help  

Upsert with string value:
Апсёрт строкого значения:
go-couchbase-cli kv --dsn "couchbase://couchbase.local" -u test_v1_cache -p "password123" -b "test_v1_cache" upsert -k kus3 --ttl 15m -v "just string"  

Operations insert, upsert, replace are supported:  
Поддерживаются операциию insert, upsert, replace:  
go-couchbase-cli kv --dsn "couchbase://couchbase.local" -u test_v1_cahce -p "password123" -b "test_v1_cache" upsert --transcoder json -k kus3 --ttl 15m -v '{"first_name": "M", "last_name": "L"}'  

Writing raw json value:  
Запись значения с роу джсона:  
go-couchbase-cli kv --dsn "couchbase://couchbase.local" -u test_v1_cache -p "password123" -b "test_v1_cache" upsert --transcoder raw_json -k kus3 --ttl 15m -v '{"first_name": "M", "last_name": "L"}'  

Receving json value:  
Получения значения с джсоном:  
go-couchbase-cli kv --dsn "couchbase://couchbase.local" -u test_v1_cache -p "password123" -b "test_v1_cache" get --transcoder json kus1 kus2 kus3

Removing by key:  
Удаление по ключу:  
go-couchbase-cli kv --dsn "couchbase://couchbase.local" -u test_v1_cache -p "password123" -b "test_v1_cache" remove kus1 kus2 kus3

Pinging nodes of cluster. If domain couchbase.local will be resolved into few ip addresses, each ip address will be tested individually: 
Пинг нод в кластере. Если домен резолвится в несколько ip адресов, каждый из них будет протестирован отдельно:  
go-couchbase-cli bucket --dsn "couchbase://couchbase.local" -u test_v1_cache -p "password123" -b "test_v1_cache" ping -n 10 --pause 100ms
Getting node diagnostics output:  
go-couchbase-cli bucket --dsn "couchbase://couchbase.local" -u test_v1_cache -p "password123" -b "test_v1_cache" diagnostics


### TODO
- Refactoring 
- Figure out how to get where key is located  
- Add interactive interface  
- Add stdin interface for importing data 
- export latency data into csv  
- Add settings for durability level  
- Commands for running benchmark  
