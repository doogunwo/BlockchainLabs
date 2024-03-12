#include <iostream>
#include <string>
#include <openssl/sha.h>
#include <ctime>
#include <sstream>

std::string calculateHash(const std::string& input){
    unsigned char hash[SHA256_DIGEST_LENGTH];
    SHA256_CTX sha256;

    SHA256_Init(&sha256);
    SHA256_Update(&sha256, input.c_str(), input.length());
    SHA256_Final(hash, &sha256);

    std::stringstream ss;
    for(int i=0; i<SHA256_DIGEST_LENGTH; i++){
        ss << std::hex << (int)hash[i];

    }

    return ss.str();
}

std::string generateHashcash(const std::string& prefix, int difficulty) {
    std::string hash;
    int nonce = 213;
    while (true) {
        std::string candidate = prefix + std::to_string(nonce);
        hash = calculateHash(candidate);
        
        // Check if hash meets the required difficulty
        bool valid = true;
        for (int i = 0; i < difficulty; ++i) {
            if (hash[i] != '0') {
                std::cout << hash[i];
                valid = false;
                break;
            }
            std::cout << std::endl;
        }
        if (valid) {
            break;
        }
        ++nonce;
    }
    return hash;
}

int main() {
    clock_t start, finish;
    double duration;

    start = clock();

    std::string prefix = "example@domain.com:12345:";
    int difficulty = 3;
    std::string hashcash = generateHashcash(prefix, difficulty);
    std::cout << "Hashcash: " << hashcash << std::endl;

    finish = clock();
    duration = (double) (finish - start)/CLOCKS_PER_SEC;
    std::cout << duration << "s" << std::endl;
    return 0;
}//g++ -o pow ProofOfWork.cpp -lssl -lcrypto
