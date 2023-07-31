# Table: aws_eventbridge_replays

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state_reason | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| event_last_replayed_time | timestamp | X | √ |  | 
| replay_name | string | X | √ |  | 
| event_start_time | timestamp | X | √ |  | 
| replay_end_time | timestamp | X | √ |  | 
| replay_start_time | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| event_end_time | timestamp | X | √ |  | 
| event_source_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 


