# Table: aws_appstream_fleets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| session_script_s3_location | json | X | √ |  | 
| vpc_config | json | X | √ |  | 
| account_id | string | X | √ |  | 
| state | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| max_concurrent_sessions | big_int | X | √ |  | 
| idle_disconnect_timeout_in_seconds | big_int | X | √ |  | 
| disconnect_timeout_in_seconds | big_int | X | √ |  | 
| display_name | string | X | √ |  | 
| name | string | X | √ |  | 
| domain_join_info | json | X | √ |  | 
| iam_role_arn | string | X | √ |  | 
| stream_view | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| fleet_type | string | X | √ |  | 
| platform | string | X | √ |  | 
| usb_device_filter_strings | string_array | X | √ |  | 
| max_user_duration_in_seconds | big_int | X | √ |  | 
| compute_capacity_status | json | X | √ |  | 
| enable_default_internet_access | bool | X | √ |  | 
| fleet_errors | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| image_name | string | X | √ |  | 
| instance_type | string | X | √ |  | 
| description | string | X | √ |  | 
| image_arn | string | X | √ |  | 


