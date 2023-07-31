# Table: aws_cloudwatch_alarms

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state_reason_data | string | X | √ |  | 
| treat_missing_data | string | X | √ |  | 
| actions_enabled | bool | X | √ |  | 
| evaluate_low_sample_count_percentile | string | X | √ |  | 
| evaluation_periods | big_int | X | √ |  | 
| evaluation_state | string | X | √ |  | 
| threshold | float | X | √ |  | 
| alarm_actions | string_array | X | √ |  | 
| state_updated_timestamp | timestamp | X | √ |  | 
| threshold_metric_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| alarm_description | string | X | √ |  | 
| dimensions | json | X | √ |  | 
| state_value | string | X | √ |  | 
| unit | string | X | √ |  | 
| region | string | X | √ |  | 
| alarm_configuration_updated_timestamp | timestamp | X | √ |  | 
| metric_name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| state_reason | string | X | √ |  | 
| alarm_arn | string | X | √ |  | 
| extended_statistic | string | X | √ |  | 
| insufficient_data_actions | string_array | X | √ |  | 
| period | big_int | X | √ |  | 
| statistic | string | X | √ |  | 
| account_id | string | X | √ |  | 
| alarm_name | string | X | √ |  | 
| comparison_operator | string | X | √ |  | 
| datapoints_to_alarm | big_int | X | √ |  | 
| ok_actions | string_array | X | √ |  | 
| state_transitioned_timestamp | timestamp | X | √ |  | 
| metrics | json | X | √ |  | 


