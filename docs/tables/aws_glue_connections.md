# Table: aws_glue_connections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| physical_connection_requirements | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| connection_properties | json | X | √ |  | 
| description | string | X | √ |  | 
| last_updated_by | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| match_criteria | string_array | X | √ |  | 
| connection_type | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 


