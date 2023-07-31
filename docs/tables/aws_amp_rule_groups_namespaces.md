# Table: aws_amp_rule_groups_namespaces

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| modified_at | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| data | byte_array | X | √ |  | 
| name | string | X | √ |  | 
| status | json | X | √ |  | 
| workspace_arn | string | X | √ |  | 
| aws_amp_workspaces_selefra_id | string | X | X | fk to aws_amp_workspaces.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


