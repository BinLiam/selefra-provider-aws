# Table: aws_db_proxies

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| db_proxy_arn | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created_date | timestamp | X | √ |  | 
| endpoint | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| vpc_subnet_ids | string_array | X | √ |  | 
| db_proxy_name | string | X | √ |  | 
| debug_logging | bool | X | √ |  | 
| engine_family | string | X | √ |  | 
| updated_date | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ |  | 
| auth | json | X | √ |  | 
| idle_client_timeout | big_int | X | √ |  | 
| require_tls | bool | X | √ |  | 
| status | string | X | √ |  | 
| vpc_security_group_ids | string_array | X | √ |  | 
| account_id | string | X | √ |  | 


