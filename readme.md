## service-location [![CircleCI](https://img.shields.io/circleci/project/github/ilgooz/service-location.svg)](https://github.com/ilgooz/service-location) [![codecov](https://codecov.io/gh/ilgooz/service-location/branch/master/graph/badge.svg)](https://codecov.io/gh/ilgooz/service-location)
A MESG service to find locations of IP addresses.

_This product includes GeoLite2 data created by MaxMind, available from [http://www.maxmind.com](http://www.maxmind.com)._


```bash
mesg-core service deploy https://github.com/ilgooz/service-location
```



# Tasks

## locate

Task key: `locate`



### Inputs

| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **ip** | `String` | IP address to find location for. |


### Outputs

##### error

Output key: `error`



| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **message** | `String` |  |

##### location

Output key: `location`



| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **city** | `String` |  |
| **continent** | `String` |  |
| **continentCode** | `String` |  |
| **country** | `String` |  |
| **countryISOCode** | `String` |  |
| **latitude** | `Number` |  |
| **longitude** | `Number` |  |
| **postalCode** | `String` |  |
| **timeZone** | `String` |  |




