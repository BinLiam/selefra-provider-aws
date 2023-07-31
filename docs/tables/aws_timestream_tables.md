# Table: aws_timestream_tables

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| database_name | string | X | √ |  | 
| table_name | string | X | √ |  | 
| table_status | string | X | √ |  | 
| aws_timestream_databases_selefra_id | string | X | X | fk to aws_timestream_databases.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| magnetic_store_write_properties | json | X | √ |  | 
| retention_properties | json | X | √ |  | 


