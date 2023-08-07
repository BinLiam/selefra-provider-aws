# Table: aws_amplify_app

## Primary Keys 

```
app_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| app_id | string | √ | √ |  | 
| arn | string | X | √ |  | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| update_time | timestamp | X | √ |  | 
| basic_auth_credentials | string | X | √ |  | 
| custom_headers | string | X | √ |  | 
| default_domain | string | X | √ |  | 
| enable_auto_branch_creation | bool | X | √ |  | 
| enable_basic_auth | bool | X | √ |  | 
| enable_branch_auto_build | bool | X | √ |  | 
| enable_branch_auto_deletion | bool | X | √ |  | 
| iam_service_role_arn | string | X | √ |  | 
| platform | string | X | √ |  | 
| repository | string | X | √ |  | 
| repository_clone_method | string | X | √ |  | 
| auto_branch_creation_config | json | X | √ |  | 
| auto_branch_creation_patterns | json | X | √ |  | 
| build_spec | json | X | √ |  | 
| custom_rules | json | X | √ |  | 
| environment_variables | json | X | √ |  | 
| production_branch | json | X | √ |  | 


