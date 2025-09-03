# level_1 task_22

## Быстрая проверка

В данном задании было реализовано небольшое unit тестирование

```
go build -o ntpcli ./cmd
./ntpcli
```

## Флаги

- -server - флаг для установки нужного ntp сервера (по умолчанию стоит time.google.com)
  пример: ./ntpcli -server=pool.ntp.org
