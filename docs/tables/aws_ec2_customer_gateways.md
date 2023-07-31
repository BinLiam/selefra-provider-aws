# Table: aws_ec2_customer_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| device_name | string | X | √ |  | 
| ip_address | string | X | √ |  | 
| bgp_asn | string | X | √ |  | 
| certificate_arn | string | X | √ |  | 
| customer_gateway_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| state | string | X | √ |  | 
| tags | json | X | √ |  | 
| type | string | X | √ |  | 


