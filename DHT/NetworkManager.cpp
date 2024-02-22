#include <asio.hpp>
#include <iostream>
#include <memory>

class NetworkMenager {
    private:
        asio.io_context ioContext;
        std::unique_ptr<asio::ip::tcp::socket> socket;
    public:
        NetworkMenager() : socket(std::make_unique<asio::ip::tcp::socket>(ioContext)) {}

        void start(unsigned short port){
            try{
                asio::ip::tcp::acceptor acceptor(ioContext, asio::ip::tcp::endpoint(asio::ip::tcp::v4(), port));
                acceptor.acceptor(*socket);
                acceptor.accept(*socket);
                std::cout << "Connection established on port " << port << std::endl;
            }
            catch (std::exception& e){
                std::cerr << "Exception: " << e.what() << std::endl;
            }
        }

};