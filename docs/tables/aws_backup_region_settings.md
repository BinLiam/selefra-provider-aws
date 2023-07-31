# Table: aws_backup_region_settings

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| resource_type_management_preference | json | X | √ |  | 
| resource_type_opt_in_preference | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 


