grpc_json_transcoder example project based on grpc-bridge example project. 

1, Clone project

    cd envoy/examples
    git clone https://github.com/tangxinfa/grpc_json_transcoder_example grpc_json_transcoder
    cd grpc_json_transcoder

2, Build proto

    sudo pip install grpcio requests
    cd service
    ./script/gen
    ./script/gen_kv.pb
    cd -

3, Build go service

    script/bootstrap
    ./script/build

4, Startup docker-compose

    docker-compose up --build
    
5, Do request

    docker-compose exec python /client/client.py set foo bar
    docker-compose exec python /client/client.py get foo
    docker-compose exec python /client/client.py count


