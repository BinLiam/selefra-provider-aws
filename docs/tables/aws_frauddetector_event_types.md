# Table: aws_frauddetector_event_types

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| labels | string_array | X | √ |  | 
| entity_types | string_array | X | √ |  | 
| event_ingestion | string | X | √ |  | 
| event_variables | string_array | X | √ |  | 
| ingested_event_statistics | json | X | √ |  | 
| account_id | string | X | √ |  | 
| created_time | string | X | √ |  | 
| description | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| last_updated_time | string | X | √ |  | 
| tags | json | X | √ |  | 
| name | string | X | √ |  | 
| region | string | X | √ |  | 


