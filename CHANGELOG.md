## v0.3.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29).
- Recompiled plugin with Go version `1.21`.

## v0.2.1 [2023-05-05]

_Enhancements_

- Updated to [steampipe-plugin-sdk v5.4.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v541-2023-05-05)

## v0.2.0 [2023-02-15]

_What's new?_

- License Change: MPL 2.0 -> Apache 2.0

_Enhancements_

- Added `uuid` column to `vsphere_vm` table. [#19](https://github.com/theapsgroup/steampipe-plugin-vsphere/issues/19)
- Updated `steampipe-plugin-sdk` to `v5.1.2` [#22](https://github.com/theapsgroup/steampipe-plugin-vsphere/issues/22)

_Bug fixes_

- Fixed bug where `VSPHERE_ALLOW_UNVERIFIED_SSL` environment variable wasn't utilised. - Thanks @mattschleder

## v0.1.3 [2022-11-10]

_Enhancements_
- Added `moref` column to `vsphere_datastore` table.
- Added `moref`, `product` columns to `vsphere_host` table.
- Added `moref` column to `vsphere_network` table.
- Added `moref`, `guest_full_name`, `hardware`, `host_moref`, `storage_consumed` & `devices` columns to `vpshere_vm` table.

_Credits_
- [@AnyKeyNL](https://github.com/AnykeyNL) for adding additional columns :)

## v0.1.2 [2022-10-21]

_Enhancements_
- Updated `vsphere_vm` to include new column `power` to indicated power state of the VM. - Thanks [@AnyKeyNL](https://github.com/AnykeyNL) 

## v0.1.1 [2022-10-08]

_Enhancements_
- Upgraded golang to version 1.19
- Upgraded steampipe sdk version to v4.1.7

## v0.1.0 [2022-05-05]

_Enhancements_
- Upgraded golang to version 1.18
- Upgraded steampipe sdk version to v3.1.0

## v0.0.2 [2021-11-29]

_Enhancements_
- Upgraded to golang version 1.17
- Upgraded steampipe sdk version to v1.8.2
- Upgraded vsphere govmomi sdk to v0.27.2

## v0.0.1 [2021-10-29]

_What's new?_

- Initial release
- Query vm's, datastores, networks and hosts
