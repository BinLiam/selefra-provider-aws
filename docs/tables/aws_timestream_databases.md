# Table: aws_timestream_databases

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| database_name | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| last_updated_time | timestamp | X | √ |  | 
| table_count | big_int | X | √ |  | 


