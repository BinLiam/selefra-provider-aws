# Table: aws_glacier_vault_access_policies

## Primary Keys 

```
vault_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_glacier_vaults_selefra_id | string | X | X | fk to aws_glacier_vaults.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| policy | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| vault_arn | string | √ | √ |  | 


