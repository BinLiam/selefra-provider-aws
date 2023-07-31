# Table: aws_ec2_transit_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| tags | json | X | √ |  | 
| transit_gateway_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_time | timestamp | X | √ |  | 
| options | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| state | string | X | √ |  | 
| transit_gateway_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 


