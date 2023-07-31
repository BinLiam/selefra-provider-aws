# Table: aws_frauddetector_batch_imports

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| iam_role_arn | string | X | √ |  | 
| job_id | string | X | √ |  | 
| processed_records_count | big_int | X | √ |  | 
| status | string | X | √ |  | 
| region | string | X | √ |  | 
| completion_time | string | X | √ |  | 
| event_type_name | string | X | √ |  | 
| failure_reason | string | X | √ |  | 
| output_path | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| input_path | string | X | √ |  | 
| start_time | string | X | √ |  | 
| total_records_count | big_int | X | √ |  | 
| failed_records_count | big_int | X | √ |  | 


