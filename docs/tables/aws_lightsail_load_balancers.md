# Table: aws_lightsail_load_balancers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| instance_health_summary | json | X | √ |  | 
| tls_certificate_summaries | json | X | √ |  | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 
| dns_name | string | X | √ |  | 
| health_check_path | string | X | √ |  | 
| https_redirection_enabled | bool | X | √ |  | 
| resource_type | string | X | √ |  | 
| state | string | X | √ |  | 
| support_code | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| configuration_options | json | X | √ |  | 
| instance_port | big_int | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| location | json | X | √ |  | 
| name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| protocol | string | X | √ |  | 
| public_ports | int_array | X | √ |  | 
| tls_policy_name | string | X | √ |  | 


