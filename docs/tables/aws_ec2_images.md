# Table: aws_ec2_images

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ena_support | bool | X | √ |  | 
| hypervisor | string | X | √ |  | 
| region | string | X | √ |  | 
| kernel_id | string | X | √ |  | 
| block_device_mappings | json | X | √ |  | 
| creation_date | string | X | √ |  | 
| image_type | string | X | √ |  | 
| imds_support | string | X | √ |  | 
| platform_details | string | X | √ |  | 
| root_device_type | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| architecture | string | X | √ |  | 
| description | string | X | √ |  | 
| image_location | string | X | √ |  | 
| account_id | string | X | √ |  | 
| public | bool | X | √ |  | 
| state_reason | json | X | √ |  | 
| virtualization_type | string | X | √ |  | 
| platform | string | X | √ |  | 
| product_codes | json | X | √ |  | 
| root_device_name | string | X | √ |  | 
| sriov_net_support | string | X | √ |  | 
| image_id | string | X | √ |  | 
| image_owner_alias | string | X | √ |  | 
| ramdisk_id | string | X | √ |  | 
| usage_operation | string | X | √ |  | 
| deprecation_time | string | X | √ |  | 
| state | string | X | √ |  | 
| tpm_support | string | X | √ |  | 
| arn | string | √ | √ |  | 
| boot_mode | string | X | √ |  | 
| name | string | X | √ |  | 
| owner_id | string | X | √ |  | 


