# Table: aws_quicksight_ingestions

## Primary Keys 

```
arn, data_set_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| ingestion_status | string | X | √ |  | 
| ingestion_id | string | X | √ |  | 
| queue_info | json | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| error_info | json | X | √ |  | 
| ingestion_size_in_bytes | big_int | X | √ |  | 
| ingestion_time_in_seconds | big_int | X | √ |  | 
| request_type | string | X | √ |  | 
| data_set_arn | string | X | √ |  | 
| aws_quicksight_data_sets_selefra_id | string | X | X | fk to aws_quicksight_data_sets.selefra_id | 
| account_id | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| request_source | string | X | √ |  | 
| row_info | json | X | √ |  | 


