# Table: aws_applicationautoscaling_policies

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| target_tracking_scaling_policy_configuration | json | X | √ |  | 
| policy_name | string | X | √ |  | 
| alarms | json | X | √ |  | 
| arn | string | √ | √ |  | 
| resource_id | string | X | √ |  | 
| region | string | X | √ |  | 
| service_namespace | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| policy_type | string | X | √ |  | 
| scalable_dimension | string | X | √ |  | 
| step_scaling_policy_configuration | json | X | √ |  | 
| account_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| policy_arn | string | X | √ |  | 


