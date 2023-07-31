# Table: aws_eventbridge_endpoints

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state_reason | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| event_buses | json | X | √ |  | 
| state | string | X | √ |  | 
| endpoint_url | string | X | √ |  | 
| replication_config | json | X | √ |  | 
| description | string | X | √ |  | 
| endpoint_id | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| routing_config | json | X | √ |  | 
| arn | string | √ | √ |  | 
| last_modified_time | timestamp | X | √ |  | 


