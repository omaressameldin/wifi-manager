# Wifi Manager
A golang cli to manage wifi networks, using [nmcli](https://developer.gnome.org/NetworkManager/stable/nmcli.html)

## How to install
- Download package from [releases](https://github.com/omaressameldin/wifi-manager/releases/)
- mv package to **/usr/local/bin** `cp wifi-manager /usr/local/bin`

## How to use
- **There are 5 commands available in the app:**
  - help:
      - for showing useful info about available commands`
      - **example:** `wifi-manager`
  ![wifi-manager](./readme-gifs/wifi-manager.gif)
  - list:
    - for listing available wifi netwokrs
    - **example:** `wifi-manager list --name --bars`
  ![wifi-manager-list](./readme-gifs/wifi-manager-list.gif)
  - info:
      - for showing all info about one of saved wifis
      - **example:** `wifi-manager info --main --password`
  ![wifi-manager-info](./readme-gifs/wifi-manager-info.gif)
  - con:
      - for connecting to a saved or a new network
      - **example:** `wifi-manager con`
  ![wifi-manager-con](./readme-gifs/wifi-manager-con.gif)
  - delete:
      - for deleting a saved network
      - **example:** `wifi-manager delete`
  ![wifi-manager-delete](./readme-gifs/wifi-manager-delete.gif)
  - restart:
      - for restarting network manager
      - **example:** `wifi-manager restart`
  ![wifi-manager-restart](./readme-gifs/wifi-manager-restart.gif)


## packages used
- [os/exec](https://golang.org/pkg/os/exec/)
- [Cobra](https://github.com/spf13/cobra#flags)
- [Promptui](https://github.com/manifoldco/promptui)
- [Spinner](https://github.com/briandowns/spinner)