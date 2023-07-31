# Table: aws_iam_policies

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| path | string | X | √ |  | 
| permissions_boundary_usage_count | big_int | X | √ |  | 
| update_date | timestamp | X | √ |  | 
| policy_id | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| id | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | X | √ |  | 
| attachment_count | big_int | X | √ |  | 
| default_version_id | string | X | √ |  | 
| description | string | X | √ |  | 
| is_attachable | bool | X | √ |  | 
| policy_version_list | json | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 


