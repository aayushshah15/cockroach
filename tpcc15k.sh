set -e
export EBSTPCC15="aayushs-ebs-tpcc";

build/builder.sh mkrelease amd64-linux-gnu bin/workload;
build/builder.sh mkrelease amd64-linux-gnu;

roachprod create $EBSTPCC15 --local-ssd=false --aws-machine-type=c5.4xlarge --nodes=16 --clouds=aws;

echo "putting workload and cockroach binaries onto their respective nodes..."
roachprod put $EBSTPCC15:16 bin.docker_amd64/workload workload;
roachprod put $EBSTPCC15:1-15 cockroach-linux-2.6.32-gnu-amd64 ./cockroach;

echo "starting the cockroach nodes....."
roachprod start $EBSTPCC15:1-15 --args="--cache=0.25 --max-sql-memory=0.4" --racks=5

echo "configuring the cluster for fast import..."
roachprod sql $EBSTPCC15:1 -e "alter range default configure zone using num_replicas=1; SET CLUSTER SETTING schemachanger.backfiller.max_buffer_size = '5 GiB'; SET CLUSTER SETTING kv.snapshot_recovery.max_rate = '128 MiB'; SET CLUSTER SETTING kv.snapshot_rebalance.max_rate = '128 MiB';";

echo "importing..."
roachprod run $EBSTPCC15:16 './workload fixtures import tpcc --warehouses=15000 --partitions=5 {pgurl:1}' 2>&1 | tee tpcc15kimport.out;

echo "done"
