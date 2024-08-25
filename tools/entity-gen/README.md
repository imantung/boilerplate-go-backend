# Entity Generator

Generate entity/repo from database table

```bash
task gen-entity

# or 
rm -rf internal/generated/entity
go run ./tools/entity-gen   
```

## Table Convention

Code generation based assumption that table follow below conventions:
1. Tables only have a single primary key called `id` with data type `SERIAL` or `BIGSERIAL`
2. Audit columns (i.e. `created_at`, `updated_at`, and `deleted_at`) are mandatory. Generated entities don't include these columns
3. Hard delete is not allowed. The select operation doesn't show soft-deleted rows. 

## Supported data type

The code generation only support below data type:
 - `integer` --> `int`
 - `bigint` --> `int64`
 - `text` --> `string`
 - `timestamp` --> `time.Time`

Please add additional data type support at `convertToFieldType()` function. 

## Database Transaction

- Repo check flag to use database transcation and store the error in the `context.Context`
- `BEGIN` should be called before called the operation via `dbtxn.Begin()`
- Check the example at [clock_svc.go](../../internal/app/service/clock_svc.go)
- Find the library at [dbtxn](https://github.com/imantung/dbtxn)