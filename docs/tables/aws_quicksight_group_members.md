# Table: aws_quicksight_group_members

## Primary Keys 

```
user_arn, group_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | X | √ |  | 
| member_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| user_arn | string | X | √ |  | 
| group_arn | string | X | √ |  | 
| aws_quicksight_groups_selefra_id | string | X | X | fk to aws_quicksight_groups.selefra_id | 


