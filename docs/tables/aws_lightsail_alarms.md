# Table: aws_lightsail_alarms

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| comparison_operator | string | X | √ |  | 
| period | big_int | X | √ |  | 
| threshold | float | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| notification_enabled | bool | X | √ |  | 
| notification_triggers | string_array | X | √ |  | 
| resource_type | string | X | √ |  | 
| state | string | X | √ |  | 
| arn | string | √ | √ |  | 
| contact_protocols | string_array | X | √ |  | 
| metric_name | string | X | √ |  | 
| monitored_resource_info | json | X | √ |  | 
| name | string | X | √ |  | 
| statistic | string | X | √ |  | 
| treat_missing_data | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| datapoints_to_alarm | big_int | X | √ |  | 
| evaluation_periods | big_int | X | √ |  | 
| location | json | X | √ |  | 
| support_code | string | X | √ |  | 
| unit | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


