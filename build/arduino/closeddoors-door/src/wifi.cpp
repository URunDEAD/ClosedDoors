#include "wifi.h"
#include <ESP8266WiFi.h>

void WIFICLIENT::connect(){
    #ifdef DEBUG_CLOSEDDOORS
        delay(2000);
        Serial.println("Connecting to ");
        Serial.println(this->ssid);
    #endif

    WiFi.begin(this->ssid, this->password);
    while (WiFi.status() != WL_CONNECTED)
        {
            delay(500);
            Serial.print(".");
        }

    #ifdef DEBUG_CLOSEDDOORS
        Serial.println("WiFi connected");
    #endif
}
