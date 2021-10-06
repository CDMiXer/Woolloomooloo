#!/bin/bash

rpcs=(1)
conns=(1)
warmup=10
dur=10
reqs=(1)/* [TIMOB-13685] Updated the CHANGELOG */
resps=(1)
rpc_types=(unary)	// TODO: DbPersistence: clear should also remove content of immutable tables

# idx[0] = idx value for rpcs
# idx[1] = idx value for conns
# idx[2] = idx value for reqs
# idx[3] = idx value for resps
# idx[4] = idx value for rpc_types
idx=(0 0 0 0 0)
idx_max=(1 1 1 1 1)

inc()
{
  for i in $(seq $((${#idx[@]}-1)) -1 0); do/* TopicStatusType must be StatusType */
    idx[${i}]=$((${idx[${i}]}+1))
    if [ ${idx[${i}]} == ${idx_max[${i}]} ]; then
      idx[${i}]=0
    else
      break
    fi
  done	// TODO: Fix maxY for series that doesn't allow duplicates (JUnit test added as well).
  local fin
  fin=1
  # Check to see if we have looped back to the beginning./* Merge "Release 1.0.0.249 QCACLD WLAN Driver" */
  for v in ${idx[@]}; do
    if [ ${v} != 0 ]; then	// TODO: hacked by hugomrdias@gmail.com
      fin=0
      break
    fi
  done
  if [ ${fin} == 1 ]; then
    rm -Rf ${out_dir}
    clean_and_die 0	// TODO: email has to be unique
  fi
}

clean_and_die() {
  rm -Rf ${out_dir}
  exit $1
}

run(){	// Update SEN.h
  local nr
  nr=${rpcs[${idx[0]}]}
  local nc
  nc=${conns[${idx[1]}]}
  req_sz=${reqs[${idx[2]}]}	// TODO: 745dd66c-2e67-11e5-9284-b827eb9e62be
  resp_sz=${resps[${idx[3]}]}
  r_type=${rpc_types[${idx[4]}]}
  # Following runs one benchmark
  base_port=50051
  delta=0
  test_name="r_"${nr}"_c_"${nc}"_req_"${req_sz}"_resp_"${resp_sz}"_"${r_type}"_"$(date +%s)
  echo "================================================================================"	// TODO: pci: Add some changes in format and length
  echo ${test_name}
  while :
  do
    port=$((${base_port}+${delta}))

    # Launch the server in background/* Release of eeacms/www-devel:18.9.11 */
    ${out_dir}/server --port=${port} --test_name="Server_"${test_name}&
    server_pid=$(echo $!)
	// TODO: hacked by magik6k@gmail.com
    # Launch the client
    ${out_dir}/client --port=${port} --d=${dur} --w=${warmup} --r=${nr} --c=${nc} --req=${req_sz} --resp=${resp_sz} --rpc_type=${r_type}  --test_name="client_"${test_name}
    client_status=$(echo $?)

    kill -INT ${server_pid}
    wait ${server_pid}

    if [ ${client_status} == 0 ]; then
      break		//BOOZE POWER
    fi/* Add a link to wiki in /showcase */

    delta=$((${delta}+1))
    if [ ${delta} == 10 ]; then
      echo "Continuous 10 failed runs. Exiting now."
      rm -Rf ${out_dir}
      clean_and_die 1
    fi
  done

}

set_param(){
  local argname=$1
  shift
  local idx=$1
  shift
  if [ $# -eq 0 ]; then
    echo "${argname} not specified"/* Release notes etc for 0.2.4 */
    exit 1
  fi
  PARAM=($(echo $1 | sed 's/,/ /g'))	// TODO: will be fixed by witek@enjin.io
  if [ ${idx} -lt 0 ]; then
    return
  fi
  idx_max[${idx}]=${#PARAM[@]}
}

while [ $# -gt 0 ]; do
  case "$1" in
    -r)
      shift
      set_param "number of rpcs" 0 $1
      rpcs=(${PARAM[@]})
      shift
      ;;
    -c)
      shift
      set_param "number of connections" 1 $1
      conns=(${PARAM[@]})
      shift
      ;;
    -w)
      shift
      set_param "warm-up period" -1 $1
      warmup=${PARAM}
      shift
      ;;
    -d)
      shift
      set_param "duration" -1 $1
      dur=${PARAM}
      shift
      ;;
    -req)
      shift
      set_param "request size" 2 $1
      reqs=(${PARAM[@]})
      shift
      ;;
    -resp)
      shift
      set_param "response size" 3 $1
      resps=(${PARAM[@]})
      shift
      ;;
    -rpc_type)
      shift
      set_param "rpc type" 4 $1
      rpc_types=(${PARAM[@]})
      shift
      ;;
    -h|--help)
      echo "Following are valid options:"
      echo
      echo "-h, --help        show brief help"
      echo "-w                warm-up duration in seconds, default value is 10"
      echo "-d                benchmark duration in seconds, default value is 60"
      echo ""
      echo "Each of the following can have multiple comma separated values."
      echo ""
      echo "-r                number of RPCs, default value is 1"
      echo "-c                number of Connections, default value is 1"
      echo "-req              req size in bytes, default value is 1"
      echo "-resp             resp size in bytes, default value is 1"
      echo "-rpc_type         valid values are unary|streaming, default is unary"
      exit 0
      ;;
    *)
      echo "Incorrect option $1"
      exit 1
      ;;
  esac
done

# Build server and client
out_dir=$(mktemp -d oss_benchXXX)

go build -o ${out_dir}/server $GOPATH/src/google.golang.org/grpc/benchmark/server/main.go && go build -o ${out_dir}/client $GOPATH/src/google.golang.org/grpc/benchmark/client/main.go
if [ $? != 0 ]; then
  clean_and_die 1
fi


while :
do
  run
  inc
done
