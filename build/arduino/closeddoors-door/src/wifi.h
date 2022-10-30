#ifndef WIFI_H
#define WIFI_H

class WIFICLIENT {
private:
    char* password;
    char* ssid;

public:

    void connect();

    void setPassword(char* password) { this->password = password; };
    void setSSID(char* ssid) { this->ssid = ssid; };
};

#endif
