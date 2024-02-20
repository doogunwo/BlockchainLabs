#include <string>
#include <unordered_map>

class Node {
    private:
    std::string nodeId;
    std::string ipAddr;
    int port;
    std::unordered_map<std::string,std::string> hashTable;
    
};