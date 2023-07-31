# Table: aws_ram_resource_share_permissions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| default_version | bool | X | √ |  | 
| resource_type | string | X | √ |  | 
| status | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| is_resource_type_default | bool | X | √ |  | 
| version | string | X | √ |  | 
| aws_ram_resource_shares_selefra_id | string | X | X | fk to aws_ram_resource_shares.selefra_id | 
| arn | string | √ | √ |  | 
| name | string | X | √ |  | 
| permission | json | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 


