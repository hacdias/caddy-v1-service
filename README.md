# service - a caddy plugin

[![Build](https://img.shields.io/travis/bruhs/caddy-service.svg?style=flat-square)](https://travis-ci.org/bruhs/caddy-service)
[![community](https://img.shields.io/badge/community-forum-ff69b4.svg?style=flat-square)](https://forum.caddyserver.com)

This plugin uses [github.com/kardianos/service](https://github.com/kardianos/service) to create services on Caddy. To work, it implements two flags on Caddy executable:

+ ```-name```, which corresponds to the name of the service (default: ```Caddy```).
+ ```-action```, which can be either install, uninstall, start, stop or restart.

Right now, this plugin seems to work well on most systems that kardianos/service supports, which are Windows XP+, Linux/(systemd | Upstart | SysV), and OSX/Launchd.
