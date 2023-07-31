# Table: aws_route53_domains

## Primary Keys 

```
account_id, domain_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| admin_privacy | bool | X | √ |  | 
| dnssec_keys | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| dns_sec | string | X | √ |  | 
| admin_contact | json | X | √ |  | 
| status_list | string_array | X | √ |  | 
| who_is_server | string | X | √ |  | 
| transfer_lock | bool | X | √ |  | 
| tech_contact | json | X | √ |  | 
| tech_privacy | bool | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| expiration_date | timestamp | X | √ |  | 
| nameservers | json | X | √ |  | 
| registrant_contact | json | X | √ |  | 
| reseller | string | X | √ |  | 
| abuse_contact_email | string | X | √ |  | 
| auto_renew | bool | X | √ |  | 
| registrant_privacy | bool | X | √ |  | 
| registrar_name | string | X | √ |  | 
| abuse_contact_phone | string | X | √ |  | 
| registrar_url | string | X | √ |  | 
| registry_domain_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| domain_name | string | X | √ |  | 
| updated_date | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ | `A list of tags` | 


