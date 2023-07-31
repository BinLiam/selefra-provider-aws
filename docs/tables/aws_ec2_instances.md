# Table: aws_ec2_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| launch_time | timestamp | X | √ |  | 
| metadata_options | json | X | √ |  | 
| network_interfaces | json | X | √ |  | 
| ramdisk_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| state_transition_reason_time | timestamp | X | √ |  | 
| client_token | string | X | √ |  | 
| public_ip_address | string | X | √ |  | 
| root_device_name | string | X | √ |  | 
| tags | json | X | √ |  | 
| capacity_reservation_id | string | X | √ |  | 
| ipv6_address | string | X | √ |  | 
| key_name | string | X | √ |  | 
| source_dest_check | bool | X | √ |  | 
| state_transition_reason | string | X | √ |  | 
| hibernation_options | json | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| platform_details | string | X | √ |  | 
| public_dns_name | string | X | √ |  | 
| enclave_options | json | X | √ |  | 
| instance_lifecycle | string | X | √ |  | 
| sriov_net_support | string | X | √ |  | 
| iam_instance_profile | json | X | √ |  | 
| monitoring | json | X | √ |  | 
| usage_operation | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| block_device_mappings | json | X | √ |  | 
| placement | json | X | √ |  | 
| private_dns_name_options | json | X | √ |  | 
| state_reason | json | X | √ |  | 
| hypervisor | string | X | √ |  | 
| maintenance_options | json | X | √ |  | 
| product_codes | json | X | √ |  | 
| region | string | X | √ |  | 
| instance_type | string | X | √ |  | 
| private_ip_address | string | X | √ |  | 
| tpm_support | string | X | √ |  | 
| image_id | string | X | √ |  | 
| platform | string | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| root_device_type | string | X | √ |  | 
| ami_launch_index | big_int | X | √ |  | 
| architecture | string | X | √ |  | 
| cpu_options | json | X | √ |  | 
| elastic_inference_accelerator_associations | json | X | √ |  | 
| virtualization_type | string | X | √ |  | 
| boot_mode | string | X | √ |  | 
| ebs_optimized | bool | X | √ |  | 
| kernel_id | string | X | √ |  | 
| security_groups | json | X | √ |  | 
| spot_instance_request_id | string | X | √ |  | 
| instance_id | string | X | √ |  | 
| licenses | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| elastic_gpu_associations | json | X | √ |  | 
| state | json | X | √ |  | 
| capacity_reservation_specification | json | X | √ |  | 
| ena_support | bool | X | √ |  | 
| usage_operation_update_time | timestamp | X | √ |  | 
| arn | string | √ | √ |  | 


