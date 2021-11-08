#include <QCommandLineParser>
#include <QDate>
#include <QIcon>

#include <QQmlApplicationEngine>
#include <QQmlContext>

#ifdef Q_OS_ANDROID
#include <QGuiApplication>
#include <MauiKit/Core/mauiandroid.h>
#else
#include <QApplication>
#endif

#include <MauiKit/Core/mauiapp.h>

#include <KI18n/KLocalizedString>

#include "../klash_version.h"

#include <clash.h>

#define KLASH_URI "org.maui.klash"

int main(int argc, char *argv[])
{
    QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);
    QCoreApplication::setAttribute(Qt::AA_UseHighDpiPixmaps, true);

#ifdef Q_OS_ANDROID
    QGuiApplication app(argc, argv);
    if (!MAUIAndroid::checkRunTimePermissions({"android.permission.WRITE_EXTERNAL_STORAGE"}))
        return -1;
#else
    QApplication app(argc, argv);
#endif

    app.setOrganizationName("Maui");
    app.setWindowIcon(QIcon(":/assets/klash.svg"));

    MauiApp::instance ()->setIconName ("qrc:/assets/klash.svg");

    KLocalizedString::setApplicationDomain("klash");
    KAboutData about(QStringLiteral("klash"), i18n("Klash"), KLASH_VERSION_STRING, i18n("A Clash frontend."), KAboutLicense::LGPL_V3, i18n("© 2021-%1 Maui Development Team", QString::number(QDate::currentDate().year())), QString(GIT_BRANCH) + "/" + QString(GIT_COMMIT_HASH));
    about.addAuthor(i18n("GFWK"), i18n("Developer"), QStringLiteral("gofuwako@gmail.com"));
    about.setHomepage("https://mauikit.org");
    about.setProductName("maui/klash");
    about.setBugAddress("https://invent.kde.org/maui/klash/-/issues");
    about.setOrganizationDomain(KLASH_URI);
    about.setProgramLogo(app.windowIcon());

    KAboutData::setApplicationData(about);

    QCommandLineParser parser;

    about.setupCommandLine(&parser);
    parser.process(app);

    about.processCommandLine(&parser);
    const QStringList args = parser.positionalArguments();

    QQmlApplicationEngine engine;
    const QUrl url(QStringLiteral("qrc:/main.qml"));

    engine.load(url);

    return app.exec();
}
