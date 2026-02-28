```sh
go tool github.com/vektra/mockery/v2 \
  --config ./.mockery.yaml \
  --name BankKeeper \
  --dir precompiles/common/ \
  --output precompiles/common/mocks
```
