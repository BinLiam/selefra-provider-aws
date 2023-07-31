# Table: aws_ec2_vpn_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| vpn_gateway_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| availability_zone | string | X | √ |  | 
| state | string | X | √ |  | 
| tags | json | X | √ |  | 
| type | string | X | √ |  | 
| vpc_attachments | json | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| amazon_side_asn | big_int | X | √ |  | 


