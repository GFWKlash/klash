#ifndef SYSTEM_PROXY_HELPER_H
#define SYSTEM_PROXY_HELPER_H

/* Abstract class for system proxy helper */
class SystemProxyHelper {
public:
    virtual void startSystemProxy() = 0;
    virtual void stopSystemProxy() = 0;
    virtual void resetSystemProxy() = 0;
};

#endif
