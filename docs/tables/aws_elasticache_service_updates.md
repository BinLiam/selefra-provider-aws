# Table: aws_elasticache_service_updates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| auto_update_after_recommended_apply_by_date | bool | X | √ |  | 
| engine_version | string | X | √ |  | 
| estimated_update_time | string | X | √ |  | 
| service_update_type | string | X | √ |  | 
| region | string | X | √ |  | 
| service_update_end_date | timestamp | X | √ |  | 
| service_update_name | string | X | √ |  | 
| service_update_status | string | X | √ |  | 
| engine | string | X | √ |  | 
| service_update_release_date | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| service_update_description | string | X | √ |  | 
| service_update_recommended_apply_by_date | timestamp | X | √ |  | 
| service_update_severity | string | X | √ |  | 


