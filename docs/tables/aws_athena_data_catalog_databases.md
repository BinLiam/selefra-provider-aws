# Table: aws_athena_data_catalog_databases

## Primary Keys 

```
data_catalog_arn, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_athena_data_catalogs_selefra_id | string | X | X | fk to aws_athena_data_catalogs.selefra_id | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| parameters | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| data_catalog_arn | string | X | √ |  | 


