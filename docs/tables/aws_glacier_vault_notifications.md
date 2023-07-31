# Table: aws_glacier_vault_notifications

## Primary Keys 

```
vault_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| vault_arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_glacier_vaults_selefra_id | string | X | X | fk to aws_glacier_vaults.selefra_id | 
| events | string_array | X | √ |  | 
| sns_topic | string | X | √ |  | 
| account_id | string | X | √ |  | 


