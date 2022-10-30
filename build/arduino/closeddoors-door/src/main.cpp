#include <Arduino.h>
#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include "wifi.h"

//Wifi Settings
WiFiClient wificlient;
char *wifi_ssid = "";
char *wifi_password = "";

//Server Settings
HTTPClient httpClient;
const char *serverName = "http://192.168.0.179:8080/check";


void sendRequest(char* hash)
{
  if (WiFi.status() == WL_CONNECTED)
  {
    #ifdef DEBUG_CLOSEDDOORS
        Serial.printf("Sending request to '%s' with payload: '%s'\n", serverName, hash);
    #endif

    httpClient.begin(wificlient, serverName);
    httpClient.addHeader("Content-Type", "application/x-www-form-urlencoded");
    httpClient.POST(hash);
    httpClient.end();

    #ifdef DEBUG_CLOSEDDOORS
        Serial.println("Request sent");
    #endif
  }

}

void setup()
{
    Serial.begin(9600);

    //Init Pins
    pinMode(D1, OUTPUT);

    //StartWifi
    WIFICLIENT WIFICLIENT;
    WIFICLIENT.setSSID(wifi_ssid);
    WIFICLIENT.setPassword(wifi_password);
    WIFICLIENT.connect();

}

void loop()
{
    digitalWrite(D1, HIGH);
    sendRequest("hash=9d187647296e16f2d4fbc3a768d0ab658821dc05cbcefeecff017699410f03e9");
    delay(1000);
    digitalWrite(D1, LOW);
    delay(1000);
}
