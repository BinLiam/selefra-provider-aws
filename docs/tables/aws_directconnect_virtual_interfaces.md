# Table: aws_directconnect_virtual_interfaces

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| amazon_side_asn | big_int | X | √ |  | 
| auth_key | string | X | √ |  | 
| customer_address | string | X | √ |  | 
| virtual_interface_type | string | X | √ |  | 
| vlan | big_int | X | √ |  | 
| amazon_address | string | X | √ |  | 
| bgp_peers | json | X | √ |  | 
| virtual_interface_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| connection_id | string | X | √ |  | 
| customer_router_config | string | X | √ |  | 
| jumbo_frame_capable | bool | X | √ |  | 
| arn | string | √ | √ |  | 
| aws_logical_device_id | string | X | √ |  | 
| direct_connect_gateway_id | string | X | √ |  | 
| mtu | big_int | X | √ |  | 
| region | string | X | √ |  | 
| site_link_enabled | bool | X | √ |  | 
| virtual_gateway_id | string | X | √ |  | 
| address_family | string | X | √ |  | 
| asn | big_int | X | √ |  | 
| aws_device_v2 | string | X | √ |  | 
| location | string | X | √ |  | 
| virtual_interface_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| owner_account | string | X | √ |  | 
| route_filter_prefixes | json | X | √ |  | 
| tags | json | X | √ |  | 
| virtual_interface_state | string | X | √ |  | 


