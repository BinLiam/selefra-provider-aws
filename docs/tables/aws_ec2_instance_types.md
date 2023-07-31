# Table: aws_ec2_instance_types

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| burstable_performance_supported | bool | X | √ |  | 
| hypervisor | string | X | √ |  | 
| memory_info | json | X | √ |  | 
| network_info | json | X | √ |  | 
| supported_boot_modes | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| current_generation | bool | X | √ |  | 
| dedicated_hosts_supported | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| auto_recovery_supported | bool | X | √ |  | 
| fpga_info | json | X | √ |  | 
| gpu_info | json | X | √ |  | 
| instance_storage_supported | bool | X | √ |  | 
| instance_type | string | X | √ |  | 
| v_cpu_info | json | X | √ |  | 
| placement_group_info | json | X | √ |  | 
| processor_info | json | X | √ |  | 
| inference_accelerator_info | json | X | √ |  | 
| instance_storage_info | json | X | √ |  | 
| supported_virtualization_types | string_array | X | √ |  | 
| free_tier_eligible | bool | X | √ |  | 
| hibernation_supported | bool | X | √ |  | 
| supported_root_device_types | string_array | X | √ |  | 
| bare_metal | bool | X | √ |  | 
| ebs_info | json | X | √ |  | 
| supported_usage_classes | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


