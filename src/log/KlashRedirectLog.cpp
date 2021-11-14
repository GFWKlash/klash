#include <stddef.h>
#include <stdint.h>
#include <malloc.h>

#include <QDebug>
#include <QLoggingCategory>

#include "KlashRedirectLog.h"
#include "clash.h"

Q_DECLARE_LOGGING_CATEGORY(KLASH_CORE)
Q_LOGGING_CATEGORY(KLASH_CORE, "klash.core")

void klash_log_info_impl(const char *payload)
{
    qCInfo(KLASH_CORE) << payload;
}

void klash_log_error_impl(const char *payload)
{
    qCCritical(KLASH_CORE) << payload;
}

void klash_log_warn_impl(const char *payload)
{
    qCWarning(KLASH_CORE) << payload;
}

void klash_log_debug_impl(const char *payload)
{
    qCDebug(KLASH_CORE) << payload;
}

void klash_log_verbose_impl(const char *payload)
{
    qCDebug(KLASH_CORE) << payload;
}


KlashRedirectLog::KlashRedirectLog()
{
}

KlashRedirectLog::~KlashRedirectLog()
{
    disableRedirect();
}

void KlashRedirectLog::enableRedirect()
{
    klash_set_log_info(klash_log_info_impl);
    klash_set_log_error(klash_log_error_impl);
    klash_set_log_warn(klash_log_warn_impl);
    klash_set_log_debug(klash_log_debug_impl);
    klash_set_log_verbose(klash_log_verbose_impl);

    redirectLogToKlash();
}

void KlashRedirectLog::disableRedirect()
{
    stopRedirectLogToKlash();

    klash_reset_log_info();
    klash_reset_log_error();
    klash_reset_log_warn();
    klash_reset_log_debug();
    klash_reset_log_verbose();
}
