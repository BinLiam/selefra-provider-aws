# Table: aws_appstream_images

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| applications | json | X | √ |  | 
| description | string | X | √ |  | 
| image_builder_name | string | X | √ |  | 
| image_permissions | json | X | √ |  | 
| visibility | string | X | √ |  | 
| name | string | X | √ |  | 
| image_builder_supported | bool | X | √ |  | 
| state | string | X | √ |  | 
| platform | string | X | √ |  | 
| public_base_image_released_date | timestamp | X | √ |  | 
| state_change_reason | json | X | √ |  | 
| appstream_agent_version | string | X | √ |  | 
| arn | string | √ | √ |  | 
| base_image_arn | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| display_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| image_errors | json | X | √ |  | 
| account_id | string | X | √ |  | 


