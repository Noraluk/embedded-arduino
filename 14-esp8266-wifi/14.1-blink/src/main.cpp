#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <ESP8266WebServer.h>
#include <ESP8266mDNS.h>

#include <uri/UriBraces.h>

#include "env.h"

ESP8266WebServer server(80);
ENV env;

String WIFI_SSID = env.getWifiSSID();
String WIFI_PASSWORD = env.getWifiPassword();

void setup()
{
  pinMode(LED_BUILTIN, OUTPUT);

  Serial.begin(115200);
  WiFi.mode(WIFI_STA);
  WiFi.begin(WIFI_SSID, WIFI_PASSWORD);
  Serial.println("");

  // Wait for connection
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.print("Connected to ");
  Serial.println(WIFI_SSID);
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());

  if (MDNS.begin("esp8266"))
  {
    Serial.println("MDNS responder started");
  }

  server.on(UriBraces("/led/{}"), []()
            {
              String pressStr = server.pathArg(0);
              int press = pressStr.toInt();
              digitalWrite(LED_BUILTIN, press);
              server.send(200, "application/json", "{\"message\" : \"success\"}");
            });

  server.begin();
  Serial.println("HTTP server started");
}

void loop()
{
  server.handleClient();
}