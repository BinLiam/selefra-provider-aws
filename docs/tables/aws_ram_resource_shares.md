# Table: aws_ram_resource_shares

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_updated_time | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| resource_share_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| status_message | string | X | √ |  | 
| feature_set | string | X | √ |  | 
| owning_account_id | string | X | √ |  | 
| status | string | X | √ |  | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| allow_external_principals | bool | X | √ |  | 


