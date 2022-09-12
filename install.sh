#!/bin/sh

set -e

target="StatixLabs/kvcopy"
version="0.0.5"
exe_name="kvcopy"

owner=""
repo=""
githubUrl=""

get_arch() {
    a=$(uname -m)
    case ${a} in
        "x86_64" | "amd64" )
            echo "amd64"
        ;;
        "i386" | "i486" | "i586")
            echo "386"
        ;;
        "aarch64" | "arm64" | "arm")
            echo "arm64"
        ;;
        "mips64el")
            echo "mips64el"
        ;;
        "mips64")
            echo "mips64"
        ;;
        "mips")
            echo "mips"
        ;;
        *)
            echo ${NIL}
        ;;
    esac
}

get_os(){
    # darwin: Darwin
    echo $(uname -s | awk '{print tolower($0)}')
}

args=(`echo $target | tr '/' ' '`)

if [ ${#args[@]} -ne 2 ]; then
    echo "ERROR: invalid params for repo '$1'"
    echo "ERROR: the argument should be format like 'owner/repo'"
    exit 1
else
    owner=${args[0]}
    repo=${args[1]}
fi

if [ -z "$exe_name" ]; then
    exe_name=$repo
    echo "INFO: file name is not specified, use '$repo'"
    echo "INFO: if you want to specify the name of the executable, set flag --exe=name"
fi

if [ -z "$githubUrl" ]; then
    githubUrl="https://github.com"
fi

downloadFolder="${HOME}/Downloads"
mkdir -p ${downloadFolder} # make sure download folder exists
os=$(get_os)
arch=$(get_arch)
file_name="${exe_name}_${version}_${os}_${arch}.tar.gz" # the file name should be download
downloaded_file="${downloadFolder}/${file_name}" # the file path should be download
executable_folder="$HOME/bin" # Eventually, the executable file will be placed here
asset_uri="${githubUrl}/${owner}/${repo}/releases/download/v${version}/${file_name}"

echo "[1/3] Download ${asset_uri} to ${downloadFolder}"
rm -f ${downloaded_file}
curl --fail --location --output "${downloaded_file}" "${asset_uri}"

echo "[2/3] Install ${exe_name} to the ${executable_folder}"
tar -xz -f ${downloaded_file} -C ${downloadFolder}
chmod +x ${downloadFolder}/${exe_name}
cp ${downloadFolder}/${exe_name} ${executable_folder}/${exe_name}

echo "[3/3] Set environment variables"
echo "${exe_name} was installed successfully to ${executable_folder}/${exe_name}"
if command -v $exe_name --help >/dev/null; then
    echo "Run '$exe_name --help' to get started"
else
    echo "Manually add the directory to your \$HOME/.bash_profile (or similar)"
    echo "  export PATH=${executable_folder}:\$PATH"
    echo "Run '$exe_name --help' to get started"
fi

exit 0
