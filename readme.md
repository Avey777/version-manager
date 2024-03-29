<p style="" align="center">
  <!-- <img src="https://github.com/moqsien/img_repo/raw/main/vm_header_photo_2.png" alt="Logo" width="720" height="240"> -->
  <img src="https://github.com/moqsien/img_repo/raw/main/vm_profile_image.png" alt="Logo" width="240" height="240">
</p>

[![Go Report Card](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=for-the-badge)](https://goreportcard.com/report/github.com/gvcgo/version-manager)
[![GitHub License](https://img.shields.io/github/license/gvcgo/version-manager?style=for-the-badge)](LICENSE)
[![GitHub Release](https://img.shields.io/github/v/release/gvcgo/version-manager?display_name=tag&style=for-the-badge)](https://github.com/gvcgo/version-manager/releases)
[![Discord](https://img.shields.io/discord/1191981003204477019?style=for-the-badge&logo=discord)](https://discord.gg/85c8ptYgb7)

[中文](https://github.com/gvcgo/version-manager/blob/main/docs/readmeCN.md) | [En](https://github.com/gvcgo/version-manager)

- [version-manager(vm)](#version-managervm)
- [features](#features)
- [vm versus vfox](#vm-versus-vfox)
- [Installation/Update](#installationupdate)
- [How to set a proxy?](#how-to-set-a-proxy)
- [Subcommands](#subcommands)
- [Related dirs](#related-dirs)
- [For Windows](#for-windows)
- [Contributors](#contributors)
- [Supplementary](#supplementary)
- [Todo-List](#todo-list)

------
<p id="1"></p>  

### version-manager(vm)

**vm** is a simple, cross-platform, and well-tested version manager for programming **languages** and **tools**. It is totally created for general purpose. You don't need any plugins, but just vm. Then everything can be managed.

Maybe you've already heard of **sdkman**, **gvm**, **nvm**, **pyenv**, **phpenv**, etc. However, never of these can manage multiple programming languages. **gvc** does the job, but it adopts a lot of other features. Most importantly, **gvc** provides free VPNs/Proxies, which makes its promotion become impossible in China. So, **vfox** is born. Indeed, **vfox** implies some promising features by introducing lua runtime as plugins just like neovim. Actually, **vfox** is not as perfect as what it promises. Therefore, **vm** comes.

------

<p id="2"></p>

### features

- Installs or uninstalls versions of sdk.
- Swithes between versions of sdk.
- Handles envs.
- Friendly to VSCoders or Neovimers.
- No plugins needed.
- More stable.
- Download files blazingly fast🚀🚀🚀 with multi-threads. See with command **vm use -h**.

------
<p id="3"></p> 

### vm versus vfox

| sdk | vm | vfox |
|-------|-------|-------|
| **java(jdk)** | ✅︎ | ✅︎ |
| **maven** | ✅︎ | ✅︎ |
| **gradle** | ✅︎ | ✅︎ |
| **kotlin** | ✅︎ | ✅︎ |
| **scala** | ✅︎ | ❌︎ |
| **python** | ✅︎ | ✅︎ |
| **miniconda** | ✅︎ | ❌︎ |
| **go** | ✅︎ | ✅︎ |
| **node** | ✅︎ | ✅︎ |
| **deno** | ✅︎ | ✅︎ |
| **bun** | ✅︎ | ❌︎ |
| **flutter(dart)** | ✅︎ | ✅︎ |
| **.net** | ✅︎ | ✅︎ |
| **zig** | ✅︎ | ✅︎ |
| **php** | ✅︎ | ❌︎ |
| **rust** | ✅︎ | ❌︎ |
| **cmdline-tool(android)** | ✅︎ | ❌︎ |
| **vlang** | ✅︎ | ❌︎ |
| **cygwin** | ✅︎ | ❌︎ |
| **msys2** | ✅︎ | ❌︎ |
| **julia** | ✅︎ | ❌︎ |
| **typst** | ✅︎ | ❌︎ |
| **gleam** | ✅︎ | ❌︎ |
| **git-for-windows** | ✅︎ | ❌︎ |
| **neovim** | ✅︎ | ❌︎ |
| **vscode** | ✅︎ | ❌︎ |
| **protobuf(protoc)** | ✅︎ | ❌︎ |
| **lazygit** | ✅︎ | ❌︎ |
| **kubectl** | ✅︎ | ❌︎ |

------

<p id="4"></p>  

### Installation/Update
- for **MacOS/Linux**(run the command below in terminal)
```bash
curl --proto '=https' --tlsv1.2 -sSf https://gvc.1710717.xyz/proxy/https://raw.githubusercontent.com/gvcgo/version-manager/main/scripts/install.sh | sh
```

- for **Windows**(run the command below in powershell)
```powershell
powershell -nop -c "iex(New-Object Net.WebClient).DownloadString('https://gvc.1710717.xyz/proxy/https://raw.githubusercontent.com/gvcgo/version-manager/main/scripts/install.ps1')"
```

------

<p id="5"></p> 

### How to set a proxy?

**Choose either proxy or reverse-proxy.**

- **proxy**
```bash
vm set-proxy <http://localhost:port or socks5://localhost:port>
```

- **reverse-proxy**

```bash
# reverse proxy <https://gvc.1710717.xyz/proxy/> is available for free.
vm set-reverse-proxy https://gvc.1710717.xyz/proxy/
```

- **enable downloading from mirror sites in China**.
```bash
vm use -mirror-in-china go@1.22.1
```

------

<p id="6"></p> 

### Subcommands

| subcommand | args | desc |
|-------|-------|-------|
| **list** | - | Shows what's supported. |
| **search** | sdk-name | Shows available versions for a sdk. |
| **use** | sdk-name@version | Installs/Swithes to the specific version of a sdk. |
| **local** | sdk-name | Shows installed versions of a sdk. |
| **uninstall** | sdk-name@version or sdk-name@all | Uninstalls versions for a sdk. |
| **clear-cache** | sdk-name | Clears the cached files for a sdk. |
| **set-reverse-proxy** | https://gvc.1710717.xyz/proxy/ | Sets a reverse-proxy for vm. |
| **set-proxy** | http or socks5( scheme://host:port ) | Sets a local proxy for vm. |
| **env** | --remove=false/true | Sets/Removes env manually. |
| **install-self** | - | Installs vm. |
| **version** | - | Shows version info of vm. |
| **completion** | - | Generate the autocompletion script for  for the specified shell.(bash, zsh, fish, or powershell) |

------

**demo for MacOS**

<!-- <a href="https://asciinema.org/a/647462" target="_blank"><img src="https://asciinema.org/a/647462.svg" /></a> -->
![demo](https://github.com/moqsien/img_repo/raw/main/vm.gif)

**demo for Windows**

![demo](https://github.com/moqsien/img_repo/raw/main/vm_win.gif)

**demo for linux**

![demo](https://github.com/moqsien/img_repo/raw/main/vm_linux.gif)

------

<p id="7"></p> 

### Related dirs

- **vm installation dir**
```bash
$HOME/.vm/
```

- **application installation dir**

Specified during installation of **vm**.
```bash
~ % ./vm install-self
Enter App Installation Dir["$Home/.vm/" by default]:
/Users/moqsien/.vm
```

------

<p id="8"></p> 

### For Windows

**Note**: If you are using vm on Windows11, you need to enable the **Developer Mode** as vm requires to create symbolic links. If you're on Windows10, and any creating-symbolic-links-failure occurs, you can try vm with **Admin Privilege**. To get **envs** take effect for windows, you may need to close the current powershell terminal and open a new one.

------
<p id="9"></p>  

### Contributors
> Thanks to following people who have contributed to this project.
<a href="https://github.com/gvcgo/version-manager/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=gvcgo/version-manager" />
</a>

------
<p id="10"></p> 

### Supplementary
**vm** is created to be a cross-platform command line tool. It will not try to include everything like asdf-vm(based on shell script, only support unix-like system) did. Because that will greatly increase the complexity and also reduce the possibility of cross-platform. And most of the time, programming languages and tools that we usually use have been covered by **vm**. **vm** will not try to include sdks that need to be compiled under a certain platform. Because each developer's development environment is different, it is impossible to ensure the completion of a compilation. So **vm** will only use pre-compiled binaries for installations. If you have any sdks or tools to recommand for version management, please raise an issue in [Issues](https://github.com/gvcgo/version-manager/issues).

------
<p id="11"></p> 

### Todo-List
- [ ] To manage package repo mirror sites in China.
