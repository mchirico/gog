
# Commands:
```bash
docker-compose build
docker-compose up

```


```bash
curl -X POST \
   http://localhost:8082/rpc \
   -H 'cache-control: no-cache' \
   -H 'content-type: application/json' \
   -d '{
   "method": "JSONServer.GiveBookDetail",
   "params": [{
   "Id": "1234"
   }],
   "id": "1"
}'
```
