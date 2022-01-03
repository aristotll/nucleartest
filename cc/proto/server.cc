#include <echo_service.pb.h>

class MyEchoService : public EchoService {
    public:
    virtual void Echo(::PROTOBUF_NAMESPACE_ID::RpcController* controller,
                       const ::EchoRequest* request,
                       ::EchoResponse* response,
                       ::google::protobuf::Closure* done) {
                           std::cout << request->msg() << std::endl;
                           response->set_msg(std::string("I have received: " + request->msg() + std::string('\n')));
                           done->Run();
                       }
}

int main(int argc, char const *argv[]) { 
    MyEchoService myEchoService;
    
    return 0; 
}
