```sh
go tool github.com/vektra/mockery/v2 \
  --config ./.mockery.yaml \
  --name BankKeeper \
  --dir ./x/precisebank/types \
  --output ./x/precisebank/types/mocks \
  --filename MockBankKeeper.go
```

```sh
go tool github.com/vektra/mockery/v2 \
  --config ./.mockery.yaml \
  --name AccountKeeper \
  --dir ./x/precisebank/types \
  --output ./x/precisebank/types/mocks \
  --filename MockAccountKeeper.go
```
