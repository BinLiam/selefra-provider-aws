# Table: aws_ec2_hosts

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| host_recovery | string | X | √ |  | 
| host_reservation_id | string | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| available_capacity | json | X | √ |  | 
| client_token | string | X | √ |  | 
| instances | json | X | √ |  | 
| allows_multiple_instance_types | string | X | √ |  | 
| host_properties | json | X | √ |  | 
| availability_zone_id | string | X | √ |  | 
| release_time | timestamp | X | √ |  | 
| allocation_time | timestamp | X | √ |  | 
| availability_zone | string | X | √ |  | 
| member_of_service_linked_resource_group | bool | X | √ |  | 
| state | string | X | √ |  | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| auto_placement | string | X | √ |  | 
| host_id | string | X | √ |  | 


