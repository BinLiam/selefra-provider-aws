# Table: aws_lambda_function_event_source_mappings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| queues | string_array | X | √ |  | 
| parallelization_factor | big_int | X | √ |  | 
| topics | string_array | X | √ |  | 
| uuid | string | X | √ |  | 
| event_source_arn | string | X | √ |  | 
| maximum_batching_window_in_seconds | big_int | X | √ |  | 
| state_transition_reason | string | X | √ |  | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| maximum_retry_attempts | big_int | X | √ |  | 
| self_managed_event_source | json | X | √ |  | 
| tumbling_window_in_seconds | big_int | X | √ |  | 
| last_modified | timestamp | X | √ |  | 
| scaling_config | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| bisect_batch_on_function_error | bool | X | √ |  | 
| function_response_types | string_array | X | √ |  | 
| function_arn | string | X | √ |  | 
| last_processing_result | string | X | √ |  | 
| starting_position | string | X | √ |  | 
| batch_size | big_int | X | √ |  | 
| destination_config | json | X | √ |  | 
| source_access_configurations | json | X | √ |  | 
| amazon_managed_kafka_event_source_config | json | X | √ |  | 
| filter_criteria | json | X | √ |  | 
| starting_position_timestamp | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| maximum_record_age_in_seconds | big_int | X | √ |  | 
| self_managed_kafka_event_source_config | json | X | √ |  | 


