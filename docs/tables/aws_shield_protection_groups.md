# Table: aws_shield_protection_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aggregation | string | X | √ |  | 
| protection_group_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| members | string_array | X | √ |  | 
| pattern | string | X | √ |  | 
| protection_group_arn | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| tags | json | X | √ |  | 


