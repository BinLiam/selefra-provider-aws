# Table: aws_backup_protected_resource

## Primary Keys 

```
resource_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| resource_arn | string | √ | √ |  | 
| resource_type | string | X | √ |  | 
| last_backup_time | timestamp | X | √ |  | 


