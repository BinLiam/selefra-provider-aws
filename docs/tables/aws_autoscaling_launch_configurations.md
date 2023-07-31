# Table: aws_autoscaling_launch_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| classic_link_vpc_security_groups | string_array | X | √ |  | 
| kernel_id | string | X | √ |  | 
| key_name | string | X | √ |  | 
| placement_tenancy | string | X | √ |  | 
| ramdisk_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created_time | timestamp | X | √ |  | 
| launch_configuration_name | string | X | √ |  | 
| associate_public_ip_address | bool | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| user_data | string | X | √ |  | 
| account_id | string | X | √ |  | 
| image_id | string | X | √ |  | 
| block_device_mappings | json | X | √ |  | 
| instance_monitoring | json | X | √ |  | 
| launch_configuration_arn | string | X | √ |  | 
| instance_type | string | X | √ |  | 
| ebs_optimized | bool | X | √ |  | 
| iam_instance_profile | string | X | √ |  | 
| metadata_options | json | X | √ |  | 
| spot_price | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| classic_link_vpc_id | string | X | √ |  | 


