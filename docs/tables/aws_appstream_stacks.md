# Table: aws_appstream_stacks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| access_endpoints | json | X | √ |  | 
| arn | string | √ | √ |  | 
| stack_errors | json | X | √ |  | 
| application_settings | json | X | √ |  | 
| feedback_url | string | X | √ |  | 
| redirect_url | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| embed_host_domains | string_array | X | √ |  | 
| storage_connectors | json | X | √ |  | 
| streaming_experience_settings | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_time | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| display_name | string | X | √ |  | 
| user_settings | json | X | √ |  | 


