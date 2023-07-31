# Table: aws_autoscaling_scheduled_actions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| auto_scaling_group_name | string | X | √ |  | 
| scheduled_action_arn | string | X | √ |  | 
| time | timestamp | X | √ |  | 
| desired_capacity | big_int | X | √ |  | 
| end_time | timestamp | X | √ |  | 
| max_size | big_int | X | √ |  | 
| arn | string | √ | √ |  | 
| recurrence | string | X | √ |  | 
| scheduled_action_name | string | X | √ |  | 
| region | string | X | √ |  | 
| min_size | big_int | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| time_zone | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


