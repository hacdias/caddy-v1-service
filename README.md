# service - a caddy plugin

[![Build](https://img.shields.io/travis/bruhs/caddy-service.svg?style=flat-square)](https://travis-ci.org/bruhs/caddy-service)
[![community](https://img.shields.io/badge/community-forum-ff69b4.svg?style=flat-square)](https://forum.caddyserver.com)

This plugin uses [github.com/kardianos/service](https://github.com/kardianos/service) to create services on Caddy. To work, it implements two flags on Caddy executable:

+ ```-name```, which corresponds to the name of the service (default: ```Caddy```).
+ ```-service```, which can be either install, uninstall, start, stop or restart.

Right now, this plugin seems to work well on most systems that kardianos/service supports, which are Windows XP+, Linux/(systemd | Upstart | SysV), and OSX/Launchd.

**This plugin wasn't widely tested. We count on the community to let us know if you find any eventual bugs and tell us what we can do to improve its functionality.**

### Install a Caddy service:

```
caddy -service install [-name optionalName]
```

Notice that if you install the service with a name that is not the default's, you will need to specify it everytime you use one of the other commands using the flag `-name`.

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
