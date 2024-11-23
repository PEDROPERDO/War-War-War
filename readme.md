```bash
docker-compose up
```

:one: Lab 01

```bash
docker build -t simpletensor:latest War

docker run -p 8501:8501 simpletensor:latest

curl -d '{"instances":[[5], [10]]}' \
-H "Content-Type: application/json" \
-X POST http://localhost:8501/v1/models/SimpleReg:predict
```