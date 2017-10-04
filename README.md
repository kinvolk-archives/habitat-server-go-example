# habitat-server-go-example

This repo contains the [Habitat](https://github.com/habitat-sh/habitat) Go plan used in the [Habitat-Operator](https://github.com/kinvolk/habitat-operator). A Docker image of this plan compiled as a Habitat service can be found [here](https://hub.docker.com/r/kinvolk/bindgo-hab/).

To build:
* Make sure you have `hab` command-line interface tool for either [Mac OS X](https://api.bintray.com/content/habitat/stable/darwin/x86_64/hab-$latest-x86_64-darwin.zip?bt_package=hab-x86_64-darwin) or [Linux](https://api.bintray.com/content/habitat/stable/linux/x86_64/hab-$latest-x86_64-linux.tar.gz?bt_package=hab-x86_64-linux).
* Run `hab studio enter`.
* From the root of the repo, run `build`.
