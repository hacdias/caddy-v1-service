# service - a caddy plugin

[![Build](https://img.shields.io/travis/hacdias/caddy-service.svg?style=flat-square)](https://travis-ci.org/hacdias/caddy-service)
[![community](https://img.shields.io/badge/community-forum-ff69b4.svg?style=flat-square)](https://forum.caddyserver.com)

This plugin uses [github.com/kardianos/service](https://github.com/kardianos/service) to create services on Caddy. To work, it implements two flags on Caddy executable:

+ ```-name```, which corresponds to the name of the service (default: ```Caddy```).
+ ```-service```, which can be either install, uninstall, start, stop or restart.

Right now, this plugin seems to work well on most systems that kardianos/service supports, which are Windows XP+, Linux/(systemd | Upstart | SysV), and OSX/Launchd.

**This plugin wasn't widely tested. We count on the community to let us know if you find any eventual bugs and tell us what we can do to improve its functionality.**

### Install a Caddy service:

```
caddy -service install -conf /path/to/Caddyfile [-name optionalServiceName] [-option optionValue]
```

Notes:
1. You **must** set the `-conf` flag to your Caddyfile. Though, usually Caddy works by using the caddyfile in the directory which the caddy executable resides in, which is not necessarily the working directory. However, the service does not load the Caddyfile from the executable directory, and so must be specified using the `-conf` flag.
2. Notice that if you install the service with a name that is not the default's, you will need to specify it everytime you use one of the other commands using the flag `-name`.
3. You can install the service with default Caddy flag values (e.g. -conf MyCaddyfile)



### Uninstall a Caddy service:

```
caddy -service uninstall [-name optionalName]
```

### Start a Caddy service:

```
caddy -service start [-name optionalName]
```

### Stop a Caddy service:

```
caddy -service stop [-name optionalName]
```

### Restart a Caddy service:

```
caddy -service restart [-name optionalName]
```
