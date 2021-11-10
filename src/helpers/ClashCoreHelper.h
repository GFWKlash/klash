#ifndef CLASH_CORE_HELPER
#define CLASH_CORE_HELPER

#include <QtCore>

class ClashCoreHelper : public QObject {
    Q_OBJECT
private:
    bool m_isCoreRunning;
};

#endif
