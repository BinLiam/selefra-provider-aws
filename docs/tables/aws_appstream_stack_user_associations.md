# Table: aws_appstream_stack_user_associations

## Primary Keys 

```
account_id, region, stack_name, user_name, authentication_type
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| user_name | string | X | √ |  | 
| send_email_notification | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_appstream_stacks_selefra_id | string | X | X | fk to aws_appstream_stacks.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| authentication_type | string | X | √ |  | 
| stack_name | string | X | √ |  | 


