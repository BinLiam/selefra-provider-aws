# Table: aws_ec2_ebs_volumes

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| snapshot_id | string | X | √ |  | 
| volume_id | string | X | √ |  | 
| volume_type | string | X | √ |  | 
| arn | string | √ | √ |  | 
| region | string | X | √ |  | 
| state | string | X | √ |  | 
| tags | json | X | √ |  | 
| throughput | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| size | big_int | X | √ |  | 
| availability_zone | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| multi_attach_enabled | bool | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| iops | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| attachments | json | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| encrypted | bool | X | √ |  | 
| fast_restored | bool | X | √ |  | 


