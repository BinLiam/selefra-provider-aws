# Table: aws_iot_security_profiles

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| behaviors | json | X | √ |  | 
| security_profile_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| version | big_int | X | √ |  | 
| targets | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| additional_metrics_to_retain_v2 | json | X | √ |  | 
| alert_targets | json | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| security_profile_name | string | X | √ |  | 
| region | string | X | √ |  | 
| additional_metrics_to_retain | string_array | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| security_profile_description | string | X | √ |  | 
| result_metadata | json | X | √ |  | 


