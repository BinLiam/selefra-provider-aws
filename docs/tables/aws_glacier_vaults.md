# Table: aws_glacier_vaults

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| creation_date | string | X | √ |  | 
| last_inventory_date | string | X | √ |  | 
| number_of_archives | big_int | X | √ |  | 
| size_in_bytes | big_int | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| vault_arn | string | X | √ |  | 
| vault_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


