# Table: aws_quicksight_data_sources

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| status | string | X | √ |  | 
| vpc_connection_properties | json | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| type | string | X | √ |  | 
| ssl_properties | json | X | √ |  | 
| data_source_id | string | X | √ |  | 
| error_info | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_time | timestamp | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| secret_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| alternate_data_source_parameters | json | X | √ |  | 


