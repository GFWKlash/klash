import subprocess
import datetime
import os

def get_version():
    with open('./go.mod') as file:
        for line in file.readlines():
            if "clash" in line and "klash" not in line:
                return line.split("-")[-1].strip()[:6]
    return "unknown"

def get_full_version():
    with open('./go.mod') as file:
        for line in file.readlines():
            if "clash" in line and "klash" not in line:
                return line.split(" ")[-1].strip()

def build_clash(version):
    build_time = datetime.datetime.now().strftime("%Y-%m-%d-%H%M")
    command = f"""go build -trimpath -ldflags '-X "github.com/Dreamacro/clash/constant.Version={version}" \
-X "github.com/Dreamacro/clash/constant.BuildTime={build_time}"' \
-buildmode=c-archive -o clash.a """
    subprocess.check_output(command, shell=True)


def run():
    version = get_version()
    print("current clash version:", version)
    build_clash(version)
    print("build static library complete!")
    print("done")


if __name__ == "__main__":
    run()
