# Table: aws_glacier_vault_lock_policies

## Primary Keys 

```
vault_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_date | string | X | √ |  | 
| policy | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| vault_arn | string | √ | √ |  | 
| aws_glacier_vaults_selefra_id | string | X | X | fk to aws_glacier_vaults.selefra_id | 
| expiration_date | string | X | √ |  | 
| state | string | X | √ |  | 
| result_metadata | json | X | √ |  | 


