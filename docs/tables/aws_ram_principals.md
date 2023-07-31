# Table: aws_ram_principals

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_time | timestamp | X | √ |  | 
| external | bool | X | √ |  | 
| id | string | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| resource_share_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


