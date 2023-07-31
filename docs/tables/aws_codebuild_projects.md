# Table: aws_codebuild_projects

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| file_system_locations | json | X | √ |  | 
| tags | json | X | √ |  | 
| badge | json | X | √ |  | 
| secondary_source_versions | json | X | √ |  | 
| queued_timeout_in_minutes | big_int | X | √ |  | 
| public_project_alias | string | X | √ |  | 
| service_role | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| description | string | X | √ |  | 
| cache | json | X | √ |  | 
| logs_config | json | X | √ |  | 
| secondary_sources | json | X | √ |  | 
| source_version | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| project_visibility | string | X | √ |  | 
| source | json | X | √ |  | 
| artifacts | json | X | √ |  | 
| created | timestamp | X | √ |  | 
| environment | json | X | √ |  | 
| name | string | X | √ |  | 
| vpc_config | json | X | √ |  | 
| webhook | json | X | √ |  | 
| concurrent_build_limit | big_int | X | √ |  | 
| encryption_key | string | X | √ |  | 
| secondary_artifacts | json | X | √ |  | 
| timeout_in_minutes | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| build_batch_config | json | X | √ |  | 
| resource_access_role | string | X | √ |  | 
| last_modified | timestamp | X | √ |  | 


