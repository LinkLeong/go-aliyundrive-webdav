#!/bin/bash
#
#           CasaOS Installer Script
#
#   GitHub: https://github.com/IceWhaleTech/CasaOS
#   Issues: https://github.com/IceWhaleTech/CasaOS/issues
#   Requires: bash, mv, rm, tr, type, grep, sed, curl/wget, tar
#
#   This script installs CasaOS to your path.
#   Usage:
#
#   	$ curl -fsSL https://get.icewhale.io/casaos.sh | bash
#   	  or
#   	$ wget -qO- https://get.icewhale.io/casaos.sh | bash
#
#   In automated environments, you may want to run as root.
#   If using curl, we recommend using the -fsSL flags.
#
#   This should work on Mac, Linux, and BSD systems. Please
#   open an issue if you notice any bugs.
#

clear

echo '
   _____                 ____   _____ 
  / ____|               / __ \ / ____|
 | |     __ _ ___  __ _| |  | | (___  
 | |    / _` / __|/ _` | |  | |\___ \ 
 | |___| (_| \__ \ (_| | |__| |____) |
  \_____\__,_|___/\__,_|\____/|_____/ 
                                      
   --- Made by IceWhale with YOU ---
'

###############################################################################
# Golbals                                                                     #
###############################################################################
readonly MINIMUM_DISK_SIZE_GB="5"
readonly MINIMUM_MEMORY="400"
readonly CASA_PATH=/casaOS/server

readonly physical_memory=$(LC_ALL=C free -m | awk '/Mem:/ { print $2 }')
readonly disk_size_bytes=$(LC_ALL=C df --output=size / | tail -n1)
readonly disk_size_gb=$((${disk_size_bytes} / 1024 / 1024))
readonly casa_bin="casaos"
port=80
install_path="/usr/local/bin"
###############################################################################
# Helpers                                                                     #
###############################################################################

#######################################
# Custom printing function
# Globals:
#   None
# Arguments:
#   $1 0:OK   1:FAILED
#   message
# Returns:
#   None
#######################################

show() {
    local color=("$@") output grey green red reset
    if [[ -t 0 || -t 1 ]]; then
        output='\e[0m\r\e[J' grey='\e[90m' green='\e[32m' red='\e[31m' reset='\e[0m'
    fi
    local left="${grey}[$reset" right="$grey]$reset"
    local ok="$left$green  OK  $right " failed="$left${red}FAILED$right " info="$left$green INFO $right "
    # Print color array from index $1
    Print() {
        [[ $1 == 1 ]]
        for ((i = $1; i < ${#color[@]}; i++)); do
            output+=${color[$i]}
        done
        echo -ne "$output$reset"
    }

    if (($1 == 0)); then
        output+=$ok
        color+=('\n')
        Print 1

    elif (($1 == 1)); then
        output+=$failed
        color+=('\n')
        Print 1

    elif (($1 == 2)); then
        output+=$info
        color+=('\n')
        Print 1
    fi
}

#######################################
# Check whether the specified port is occupied
# Globals:
#   None
# Arguments:
#   $1 port number
# Returns:
#   None
#######################################

function check_port() {
    ss -tlp | grep $1\ 
}

function get_ipaddr() {
    hostname -I | awk '{print $1}'
}

###############################################################################
# Main logic                                                                  #
###############################################################################

# Exit path for non-root executions
# if (($UID)); then
#     show 1 "Root privileges required, please run this script with "sudo"."
#     exit 1
# fi

#Check memory
if [[ "${physical_memory}" -lt "${MINIMUM_MEMORY}" ]]; then
    show 1 "requires atleast 1GB physical memory."
    exit 1
fi

#Check Disk
if [[ "${disk_size_gb}" -lt "${MINIMUM_DISK_SIZE_GB}" ]]; then
    show 1 "requires atleast ${MINIMUM_DISK_SIZE_GB}GB disk space (Disk space on / is ${disk_size_gb}GB)."
    exit 1
fi

#Check Docker
if [[ -x "$(command -v docker)" ]]; then
    show 0 "Docker already installed."
else
    if [[ -r /etc/os-release ]]; then
        lsb_dist="$(. /etc/os-release && echo "$ID")"
    fi
    if [[ $lsb_dist == "openwrt" ]]; then
        show 1 "Openwrt, Please install docker manually."
        exit 1
    else
        show 0 "Docker will be installed automatically."
        curl -fsSL https://get.docker.com | bash
        if [ $? -ne 0 ]; then
            show 1 "Installation failed, please try again."
            exit 1
        else
            show 0 "Docker Successfully installed."
        fi
    fi
fi

#Create CasaOS directory
create_directory() {
    ((EUID)) && sudo_cmd="sudo"
    #mkdir /casaOS
    #mkdir /casaOS/server
    #mkdir /casaOS/server/user
    $sudo_cmd mkdir -p $CASA_PATH
    $sudo_cmd mkdir -p /casaOS/logs/server
    $sudo_cmd mkdir -p /casaOS/util/shell
}

#Create Service And Start Service
gen_service() {
    ((EUID)) && sudo_cmd="sudo"
    show 2 "Try stop CasaOS system service."
    $sudo_cmd systemctl stop casaos.service # Stop before generation
    show 2 "Create system service for CasaOS."
    $sudo_cmd tee $1 >/dev/null <<EOF
				[Unit]
				Description=CasaOS Service
				StartLimitIntervalSec=0

				[Service]
				Type=simple
				Restart=always
				RestartSec=1
				User=root
				ExecStart=$install_path/$casa_bin -c $CASA_PATH/conf/conf.ini

				[Install]
				WantedBy=multi-user.target
EOF
    show 0 "CasaOS service Successfully created."

    #Check Port
    if [ -n "$(check_port :http)" ]; then
        for PORT in {81..65536}; do
            if [ ! -n "$(check_port :$PORT)" ]; then
                port=$PORT
                break
            fi
        done
    fi

    #replace port
    $sudo_cmd sed -i "s/^HttpPort =.*/HttpPort = $port/g" $CASA_PATH/conf/conf.ini

    show 2 "Create a system startup service for CasaOS."

    $sudo_cmd systemctl daemon-reload
    $sudo_cmd systemctl enable casaos

    show 2 "Start CasaOS service."
    $sudo_cmd systemctl start casaos

    PIDS=$(ps -ef | grep casaos | grep -v grep | awk '{print $2}')
    if [[ "$PIDS" != "" ]]; then
        echo " "
        echo " "
        echo "  CasaOS running at:"
        if [[ "$port" -eq "80" ]]; then
            echo "  http://$(get_ipaddr)"
        else
            echo "  http://$(get_ipaddr):$port"
        fi
        echo " "
        echo " "
    else
        show 1 "CasaOS start failed."
    fi

    #$sudo_cmd systemctl status casaos
}

#install Casa
install_casa() {
    trap 'show 1 "error $? in command: $BASH_COMMAND"; trap ERR; return 1' ERR
    target_os="unsupported"
    target_arch="unknown"

    # Fall back to /usr/bin if necessary
    if [[ ! -d $install_path ]]; then
        install_path="/usr/bin"
    fi

    # Not every platform has or needs sudo (https://termux.com/linux.html)
    ((EUID)) && sudo_cmd="sudo"

    #########################
    # Which OS and version? #
    #########################
    casa_tmp_folder="casaos"

    casa_dl_ext=".tar.gz"

    # NOTE: `uname -m` is more accurate and universal than `arch`
    # See https://en.wikipedia.org/wiki/Uname
    unamem="$(uname -m)"
    case $unamem in
    *aarch64*)
        target_arch="arm64"
        ;;
    *64*)
        target_arch="amd64"
        ;;
    *86*)
        target_arch="386"
        ;;
    *armv5*)
        target_arch="armv5"
        ;;
    *armv6*)
        target_arch="armv6"
        ;;
    *armv7*)
        target_arch="armv7"
        ;;
    *)
        show 1 "Aborted, unsupported or unknown architecture: $unamem"
        return 2
        ;;
    esac

    unameu="$(tr '[:lower:]' '[:upper:]' <<<$(uname))"
    if [[ $unameu == *DARWIN* ]]; then
        target_os="darwin"
    elif [[ $unameu == *LINUX* ]]; then
        target_os="linux"
    elif [[ $unameu == *FREEBSD* ]]; then
        target_os="freebsd"
    elif [[ $unameu == *NETBSD* ]]; then
        target_os="netbsd"
    elif [[ $unameu == *OPENBSD* ]]; then
        target_os="openbsd"
    else
        show 1 "Aborted, unsupported or unknown OS: $uname"
        return 6
    fi

    ########################
    # Download and extract #
    ########################
    show 2 "Downloading CasaOS for $target_os/$target_arch..."
    if type -p curl >/dev/null 2>&1; then
        net_getter="curl -fsSL"
    elif type -p wget >/dev/null 2>&1; then
        net_getter="wget -qO-"
    else
        show 1 "Aborted, could not find curl or wget"
        return 7
    fi

    casa_file="${target_os}-$target_arch-casaos$casa_dl_ext"
    casa_tag="$(${net_getter} https://api.github.com/repos/IceWhaleTech/CasaOS/releases/latest | grep -o '"tag_name": ".*"' | sed 's/"//g' | sed 's/tag_name: //g')"
    casa_url="https://github.com/IceWhaleTech/CasaOS/releases/download/$casa_tag/$casa_file"
    show 2 "$casa_url"

    # Use $PREFIX for compatibility with Termux on Android
    $sudo_cmd rm -rf "$PREFIX/tmp/$casa_file"

    ${net_getter} "$casa_url" >"$PREFIX/tmp/$casa_file"

    show 2 "Extracting..."
    case "$casa_file" in
    *.zip) $sudo_cmd unzip -o "$PREFIX/tmp/$casa_file" -d "$PREFIX/tmp/" ;;
    *.tar.gz) $sudo_cmd tar -xzf "$PREFIX/tmp/$casa_file" -C "$PREFIX/tmp/" ;;
    esac

    $sudo_cmd chmod +x "$PREFIX/tmp/$casa_tmp_folder/$casa_bin"

    show 2 "Putting CasaOS in $install_path (may require password)"
    $sudo_cmd mv -f "$PREFIX/tmp/$casa_tmp_folder/$casa_bin" "$install_path/"

    show 2 "Putting CasaOS Configuration file in $CASA_PATH (may require password)"

    #check conf and shell folder
    local casa_conf_path=$CASA_PATH/conf
    local casa_shell_path=$CASA_PATH/shell
    if [[ -d $casa_conf_path ]]; then
        $sudo_cmd rm -rf $casa_conf_path
    fi

    if [[ -d $casa_shell_path ]]; then
        $sudo_cmd rm -rf $casa_shell_path
    fi

    $sudo_cmd mv -f "$PREFIX/tmp/$casa_tmp_folder/"* "$CASA_PATH/"

    # remove tmp files
    $sudo_cmd rm -rf $PREFIX/tmp/$casa_tmp_folder

    if type -p $casa_bin >/dev/null 2>&1; then
        show 0 "CasaOS Successfully installed."
        trap ERR
        systemd_folder=/usr/lib/systemd/system/casaos.service
        if [ ! -d "/usr/lib/systemd/system" ]; then
            systemd_folder=/lib/systemd/system/casaos.service
            if [ ! -d "/lib/systemd/system" ]; then
                systemd_folder=/etc/systemd/system/casaos.service
            fi
        fi
        gen_service $systemd_folder
        return 0
    else
        show 1 "Something went wrong, CasaOS is not in your path"
        trap ERR
        return 1
    fi
}

create_directory
install_casa
