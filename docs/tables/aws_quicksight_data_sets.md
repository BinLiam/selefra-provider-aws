# Table: aws_quicksight_data_sets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_updated_time | timestamp | X | √ |  | 
| row_level_permission_data_set | json | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| column_level_permission_rules_applied | bool | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| data_set_id | string | X | √ |  | 
| import_mode | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| row_level_permission_tag_configuration_applied | bool | X | √ |  | 


