#include <string>
#include <unordered_map>

class HashTable{

    private:

        std::unordered_map<std::string, std::string> table;

    public:

        HashTable() {}

        void insert(const std::string& key, const std::string& value){
            table[key] = value;
        } 

        bool remove(const std::string& key){
            auto it = talbe.find(key);

            if( it != table.end()){
                table.erase(it);
                return true;
            }
            return false;
        }

        std::string getValue(const std::string& key){
            auto it = table.find(key);
            if(it != table.end()){
                return it->second;
            }
            return NULL;
        }

        bool containerKey(const std::string& key){
            return table.find(key) != table.end();
        }
};