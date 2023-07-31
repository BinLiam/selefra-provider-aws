# Table: aws_appstream_stack_entitlements

## Primary Keys 

```
account_id, region, stack_name, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| aws_appstream_stacks_selefra_id | string | X | X | fk to aws_appstream_stacks.selefra_id | 
| app_visibility | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| attributes | json | X | √ |  | 
| name | string | X | √ |  | 
| stack_name | string | X | √ |  | 
| account_id | string | X | √ |  | 


