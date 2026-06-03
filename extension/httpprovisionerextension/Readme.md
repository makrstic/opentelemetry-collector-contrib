# HttpProvisioner Extension

## Intro
Provides automated way to sync config from the HTTP server. Extension will poll configured HTTP server, check for changes in the config based on the hostname and update local config.yaml. It will send os.exit(1) signal and leave to the OS to handle restart with the new config.

Check for changes will be done using md5 hash against local and remote config files.

