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

# IPSec Policy Protocol List
```
all   egp      ggp  icmp       igmp     ipsec-ah    ipv6-frag   ipv6-route  ospf  rdp   sctp  udp       vrrp   
dccp  encap    gre  icmpv6     ipencap  ipsec-esp   ipv6-nonxt  iso-tp4     pim   rspf  st    udp-lite  xns-idp
ddp   etherip  hmp  idpr-cmtp  ipip     ipv6-encap  ipv6-opts   l2tp        pup   rsvp  tcp   vmtp      xtp
```

# IPSec Policy Action

```
discard  encrypt  none
```

# IPSec Policy Action Level
```
require unique use
```

# IPSec Policy IPsec Protocol
```
ah esp
```

