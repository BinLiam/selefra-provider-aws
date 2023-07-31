# Table: aws_ram_resource_share_associations

## Primary Keys 

```
associated_entity, resource_share_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| association_type | string | X | √ |  | 
| external | bool | X | √ |  | 
| resource_share_arn | string | X | √ |  | 
| resource_share_name | string | X | √ |  | 
| status | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| associated_entity | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| status_message | string | X | √ |  | 
| account_id | string | X | √ |  | 


