# Table: aws_elasticache_users

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status | string | X | √ |  | 
| user_group_ids | string_array | X | √ |  | 
| user_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| minimum_engine_version | string | X | √ |  | 
| access_string | string | X | √ |  | 
| authentication | json | X | √ |  | 
| engine | string | X | √ |  | 
| user_name | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 


