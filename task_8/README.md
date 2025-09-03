# level_2 task_22

## Быстрый запуск

В данном задании было реализовано небольшое unit тестирование

```
go build -o ntpcli ./cmd
./ntpcli
```

## Флаги

- -server - флаг для установки нужного ntp сервера (по умолчанию стоит time.google.com)
  пример: ./ntpcli -server=pool.ntp.org
