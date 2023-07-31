# Table: aws_ram_resources

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| resource_group_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| status_message | string | X | √ |  | 
| arn | string | √ | √ |  | 
| resource_region_scope | string | X | √ |  | 
| resource_share_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


