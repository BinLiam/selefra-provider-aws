# Table: aws_workspaces_workspaces

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| computer_name | string | X | √ |  | 
| modification_states | json | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| directory_id | string | X | √ |  | 
| error_code | string | X | √ |  | 
| ip_address | string | X | √ |  | 
| volume_encryption_key | string | X | √ |  | 
| workspace_id | string | X | √ |  | 
| related_workspaces | json | X | √ |  | 
| root_volume_encryption_enabled | bool | X | √ |  | 
| user_name | string | X | √ |  | 
| user_volume_encryption_enabled | bool | X | √ |  | 
| workspace_properties | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| bundle_id | string | X | √ |  | 
| error_message | string | X | √ |  | 
| state | string | X | √ |  | 
| subnet_id | string | X | √ |  | 


