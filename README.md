# Mikrotik provider for Terraform 

## Intro

This is a terraform provider for managing resources on your RouterOS device. To see what resources and data sources are supported, please see the [documentation](https://registry.terraform.io/providers/ddelnano/mikrotik/latest/docs) on the terraform registry.

## Support

You can discuss any issues you have or feature requests in [Discord](https://discord.gg/ZpNq8ez).

## Donations

If you get value out this project and want to show your support you can find me on [patreon](https://www.patreon.com/ddelnano).

## Contributing

### Dependencies
- RouterOS. See which versions are supported by what is tested in [CI](.github/workflows/continuous-integration.yml)
- Terraform 0.12+

### Testing

The provider is tested with Terraform's acceptance testing framework. As long as you have a RouterOS device you should be able to run them. Please be aware it will create resources on your device! Code that is accepted by the project will not be destructive for anything existing on your router but be careful when changing test code!

In order to run the tests you will need to set the following environment variables:
```bash
export MIKROTIK_HOST=router-hostname:8728
export MIKROTIK_USER=username
# Please be aware this will put your password in your bash history and is not safe
export MIKROTIK_PASSWORD=password
```

After those environment variables are set you can run the tests with the following command:
```bash
make testacc
```

# IPSec

## Proposal

## Authentication Algorithms (can select multiple)
```
md5  null  sha1  sha256  sha512
```

### Encryption Algorithms (can select multiple)
```
3des  aes-128-cbc  aes-128-ctr  aes-128-gcm  aes-192-cbc  aes-192-ctr  aes-192-gcm  aes-256-cbc
aes-256-ctr  aes-256-gcm  blowfish  camellia-128  camellia-192  camellia-256  des  null  twofish
```

### PFS Group (select ony one)
```
ec2n155   ec2n185   ecp256    ecp384    ecp521    modp768   modp1024
modp1536  modp2048  modp3072  modp4096  modp6144  modp8192  none
```

## Identity

### Authentication Method (select only one)
```
digital-signature  eap  eap-radius  pre-shared-key  pre-shared-key-xauth  rsa-key  rsa-signature-hybrid
```

### EAP Method (select only one)
```
eap-mschapv2  eap-peap  eap-tls  eap-ttls
```

### Generate Policy
```
no  port-override  port-strict
```

### Matched By
```
certificate  remote-id
```

## Policy

### Protocol List (can select multiple)
```
all   egp      ggp  icmp       igmp     ipsec-ah    ipv6-frag   ipv6-route  ospf  rdp   sctp  udp       vrrp   
dccp  encap    gre  icmpv6     ipencap  ipsec-esp   ipv6-nonxt  iso-tp4     pim   rspf  st    udp-lite  xns-idp
ddp   etherip  hmp  idpr-cmtp  ipip     ipv6-encap  ipv6-opts   l2tp        pup   rsvp  tcp   vmtp      xtp
```

### Action (select only one)

```
discard  encrypt  none
```

### Action Level (select only one)
```
require unique use
```

### IPsec Protocol (select only one)
```
ah esp
```

