# Table: aws_frauddetector_batch_predictions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| detector_name | string | X | √ |  | 
| failure_reason | string | X | √ |  | 
| account_id | string | X | √ |  | 
| completion_time | string | X | √ |  | 
| job_id | string | X | √ |  | 
| status | string | X | √ |  | 
| detector_version | string | X | √ |  | 
| input_path | string | X | √ |  | 
| start_time | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| total_records_count | big_int | X | √ |  | 
| arn | string | √ | √ |  | 
| event_type_name | string | X | √ |  | 
| iam_role_arn | string | X | √ |  | 
| last_heartbeat_time | string | X | √ |  | 
| output_path | string | X | √ |  | 
| processed_records_count | big_int | X | √ |  | 


