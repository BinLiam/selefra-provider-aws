# Table: aws_sagemaker_training_jobs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| algorithm_specification | json | X | √ |  | 
| stopping_condition | json | X | √ |  | 
| checkpoint_config | json | X | √ |  | 
| profiler_rule_configurations | json | X | √ |  | 
| secondary_status_transitions | json | X | √ |  | 
| vpc_config | json | X | √ |  | 
| secondary_status | string | X | √ |  | 
| training_job_name | string | X | √ |  | 
| hyper_parameters | json | X | √ |  | 
| retry_strategy | json | X | √ |  | 
| warm_pool_status | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| profiling_status | string | X | √ |  | 
| model_artifacts | json | X | √ |  | 
| training_job_status | string | X | √ |  | 
| enable_inter_container_traffic_encryption | bool | X | √ |  | 
| enable_network_isolation | bool | X | √ |  | 
| failure_reason | string | X | √ |  | 
| final_metric_data_list | json | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| training_end_time | timestamp | X | √ |  | 
| training_time_in_seconds | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_time | timestamp | X | √ |  | 
| debug_rule_evaluation_statuses | json | X | √ |  | 
| enable_managed_spot_training | bool | X | √ |  | 
| input_data_config | json | X | √ |  | 
| output_data_config | json | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ | `The tags associated with the model.` | 
| auto_ml_job_arn | string | X | √ |  | 
| labeling_job_arn | string | X | √ |  | 
| profiler_config | json | X | √ |  | 
| training_start_time | timestamp | X | √ |  | 
| resource_config | json | X | √ |  | 
| debug_hook_config | json | X | √ |  | 
| profiler_rule_evaluation_statuses | json | X | √ |  | 
| experiment_config | json | X | √ |  | 
| tensor_board_output_config | json | X | √ |  | 
| tuning_job_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| training_job_arn | string | X | √ |  | 
| billable_time_in_seconds | big_int | X | √ |  | 
| debug_rule_configurations | json | X | √ |  | 
| environment | json | X | √ |  | 
| role_arn | string | X | √ |  | 


