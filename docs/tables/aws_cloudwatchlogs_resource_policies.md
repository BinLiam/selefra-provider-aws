# Table: aws_cloudwatchlogs_resource_policies

## Primary Keys 

```
account_id, region, policy_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policy_document | json | X | √ |  | 
| policy_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| last_updated_time | big_int | X | √ |  | 


