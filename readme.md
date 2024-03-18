<p style="" align="center">
  <!-- <img src="https://github.com/moqsien/img_repo/raw/main/vm_header_photo_2.png" alt="Logo" width="720" height="240"> -->
  <img src="https://github.com/moqsien/img_repo/raw/main/vm_profile_image.png" alt="Logo" width="240" height="240">
</p>

------
### version-manager(vm)

**vm** is a simple version manager for programming **languages** and **tools**. It is totally created for general purpose.
You don't need any plugins, just vm. Then you can manage nearly everything you want.

------
### What's supported?

- programming languages
  - **java**(jdk, maven, gradle)
  - **kotlin**
  - **scala**(coursier, scala)
  - **go**
  - **python**(miniconda, python)
  - **php**(php8.0+ only)
  - **javascript/typescript**(node, bun, deno)
  - **dart**(flutter, dart)
  - **julia**
  - **.net**(dotnet-sdk, c#)
  - **c/c++**(cygwin-installer, msys2-installer)
  - **rust**(rustup-init, rust)
  - **vlang**(v, v-analyzer)
  - **zig**(zig, zls)
  - **typst**(typst, typst-lsp)
  - **gleam**
- tools
  - **commandline-tools**(for android, latest version only)
  - **git**(for windows only)
  - **lazygit**(depends on git)
  - **protoc**(protobuf)
  - **gsudo**(for windows only)
  - **vscode**(latest version only)
  - **neovim**
  - **fd**
  - **fzf**
  - **ripgrep**
  - **tree-sitter**
  - **vhs**
  - **glow**

------

### Installation/Update
- for MacOS/Linux
```bash
curl --proto '=https' --tlsv1.2 -sSf https://gvc.1710717.xyz/proxy/https://raw.githubusercontent.com/gvcgo/version-manager/main/scripts/install.sh | sh
```

- for Windows
```powershell
powershell -nop -c "iex(New-Object Net.WebClient).DownloadString('https://gvc.1710717.xyz/proxy/https://raw.githubusercontent.com/gvcgo/version-manager/main/scripts/install.ps1')"
```

------

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
| **set-proxy** | http or socks5 | Sets a local proxy for vm. |
| **install-self** | - | Installs vm. |
| **version** | sdk-name | Shows version info of vm. |
------

**demo**

<script src="https://asciinema.org/a/647462.js" id="asciicast-647462" async="true"></script>

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

### Proxy
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
