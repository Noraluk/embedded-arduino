#ifndef ENV_H
#define ENV_H

#include "Arduino.h"

// Create env header and copy this to the target and remove _EXAMPLE
class ENV_EXAMPLE
{
public:
	String getWifiSSID();
	String getWifiPassword();
};

#endif