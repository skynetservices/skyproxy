SkyProxy
========

SkyProxy generates Hipache configurations directly into Redis based on SkyDNS changes.


___
SkyProxy uses SkyDNS to discover available web instances and watches for changes in availability of those instances.  When changes are detected, SkyProxy will send configuration changes to Redis, which will automatically re-configure Hipache.


