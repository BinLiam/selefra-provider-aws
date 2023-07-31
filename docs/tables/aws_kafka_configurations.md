# Table: aws_kafka_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_time | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| description | string | X | √ |  | 
| kafka_versions | string_array | X | √ |  | 
| latest_revision | json | X | √ |  | 
| name | string | X | √ |  | 


