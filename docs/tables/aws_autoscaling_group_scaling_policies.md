# Table: aws_autoscaling_group_scaling_policies

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policy_type | string | X | √ |  | 
| predictive_scaling_configuration | json | X | √ |  | 
| group_arn | string | X | √ |  | 
| aws_autoscaling_groups_selefra_id | string | X | X | fk to aws_autoscaling_groups.selefra_id | 
| adjustment_type | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| metric_aggregation_type | string | X | √ |  | 
| alarms | json | X | √ |  | 
| target_tracking_configuration | json | X | √ |  | 
| scaling_adjustment | big_int | X | √ |  | 
| step_adjustments | json | X | √ |  | 
| cooldown | big_int | X | √ |  | 
| estimated_instance_warmup | big_int | X | √ |  | 
| min_adjustment_step | big_int | X | √ |  | 
| policy_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| auto_scaling_group_name | string | X | √ |  | 
| min_adjustment_magnitude | big_int | X | √ |  | 
| policy_arn | string | X | √ |  | 


