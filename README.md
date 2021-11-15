<h1 align="center">
  <img src="src/assets/logo.png" alt="Klash" width="200">
  <br>Klash<br>
</h1>

A cross-platform Clash GUI based on KDE MAUI kit.

# Build

We are currently recommending to build Klash on Arch Linux.

## Install Kirigami

Follow the instructions in https://develop.kde.org/docs/kirigami/introduction-getting_started/.

## Install mauikit

Install mauikit from the pacman package manger:

```bash
sudo pacman -S mauikit mauikit-filebrowsing
```

If the version provided by the package manager does not work, you need to build mauikit by yourself. Otherwise, jump to [Build Clash core](#Build Clash core).

## Build mauikit(optionally)

Fetch the code:

```bash
wget https://invent.kde.org/maui/mauikit/-/archive/v2.1/mauikit-v2.1.tar
wget https://invent.kde.org/maui/mauikit-filebrowsing/-/archive/v2.1/mauikit-filebrowsing-v2.1.tar
```

Extract the code:

```bash
tar xvf mauikit-v2.1.tar
tar xvf mauikit-filebrowsing-v2.1.tar
```

Build and install Maui Kit:

```bash
cd mauikit
make build && cd build
cmake ..
make
sudo make install
```

Return to the source code extraction folder.

Build and install Maui Kit File Browsing component:

```bash
cd mauikit-filebrowsing
make build && cd build
cmake ..
make
sudo make install
```

## Build Clash core

Golang 1.16 or higher is needed to build Clash core.

Enter `clash` subdirectory and execute:

```bash
python3 build_clash.py
```

There will be a `clash.a` and a `clash.h` in the directory.

## Build this project

Return to the project root directory. Build it:

```bash
make build && cd build
cmake ..
make
```

The executable is under `build/bin`. Run it in the `build` directory:

```bash
bin/klash
```
