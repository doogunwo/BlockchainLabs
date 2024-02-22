#include <string>
#include <unordered_map>

class Node {

    private:

        std::string nodeId;
        std::string ipAddr;
        int port;
        std::unordered_map<std::string,std::string> hashTable;

    public:

        Node(const std::string& id, const std::string& ip, int port) : nodeId(id), ipAddress(ip), port(port) {};

        void store(const std::string& key, const std::string& value){
            hashTable[key] = value;
        }

        std::string findValue(const std::string& key){
            auto it = hashTable.find(key);
            if(it != hashTable.end()){
                return it->second;
            }
            return NULL;
        }
};