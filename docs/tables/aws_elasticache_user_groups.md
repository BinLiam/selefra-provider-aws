# Table: aws_elasticache_user_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| engine | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| pending_changes | json | X | √ |  | 
| replication_groups | string_array | X | √ |  | 
| status | string | X | √ |  | 
| user_group_id | string | X | √ |  | 
| user_ids | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


