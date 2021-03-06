#include <QCommandLineParser>
#include <QDate>
#include <QIcon>

#include <QQmlApplicationEngine>
#include <QQmlContext>
#include <QDebug>

#ifdef Q_OS_ANDROID
#include <QGuiApplication>
#include <MauiKit/Core/mauiandroid.h>
#else
#include <QApplication>
#include <QSystemTrayIcon>
#include <QMenu>
#endif

#include <MauiKit/Core/mauiapp.h>

#include <KI18n/KLocalizedString>

#include "../klash_version.h"

#include <clash.h>
#include <iostream>

#include <ClashCoreHelper.h>
#include <KlashRedirectLog.h>

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
    app.setWindowIcon(QIcon(":/assets/logo.png"));

    MauiApp::instance ()->setIconName ("qrc:/assets/logo.png");

    // Prepare clash version
    char *clashVersionString = getClashVersion();
    QString qClashVersionString("clash ");
    qClashVersionString += clashVersionString;
    qDebug() << "[Clash] Core version " << clashVersionString;
    free(clashVersionString);

    KLocalizedString::setApplicationDomain("klash");
    KAboutData about(QStringLiteral("klash"), i18n("Klash"), KLASH_VERSION_STRING, i18n("A Clash frontend."), KAboutLicense::LGPL_V3, i18n("© 2021-%1 Maui Development Team", QString::number(QDate::currentDate().year())), QString(GIT_BRANCH) + "/" + QString(GIT_COMMIT_HASH));
    about.addAuthor(i18n("GFWK"), i18n("Developer"), QStringLiteral("gofuwako@gmail.com"));
    about.setHomepage("https://mauikit.org");
    about.setProductName("maui/klash");
    about.setBugAddress("https://github.com/GFWKlash/klash/issues");
    about.setOrganizationDomain(KLASH_URI);
    about.setProgramLogo(app.windowIcon());
    about.setOtherText(qClashVersionString);

    KAboutData::setApplicationData(about);

    QCommandLineParser parser;

    about.setupCommandLine(&parser);
    parser.process(app);

    about.processCommandLine(&parser);
    const QStringList args = parser.positionalArguments();

    // Redirect logs
    KlashRedirectLog logRedirector;
    logRedirector.enableRedirect();

    // Init clash core
    bool clashRunning = false;
    initClashCore();
    // std::cout << "[Clash] Returned " << run(0, 1) << std::endl;

    QQmlApplicationEngine engine;

    // Daemonize the backend and notifier
#if defined(Q_OS_ANDROID) || defined(Q_OS_IOS)
    // TODO: On Android we may need a service
#else
    // Add menus in systray
    QMenu* menu = new QMenu;
    auto toggleClashCore = menu->addAction(QIcon(":/assets/logo.png"), i18n("Run Clash core"));
    QObject::connect(toggleClashCore, &QAction::triggered, toggleClashCore, [toggleClashCore, &clashRunning]() {
        // TODO: Toggle and update state
        if (clashRunning) {
            toggleClashCore->setText(i18n("Run Clash"));
            qDebug() << "Stopping Clash core";
        } else {
            toggleClashCore->setText(i18n("Stop Clash"));
            qDebug() << "Starting Clash core";
        }
        clashRunning = !clashRunning;
    });

    // TODO: Add profile menu
    //          - profiles
    QMenu* profileMenu = menu->addMenu(QIcon::fromTheme("document-preview"), i18n("Profile"));
    /* auto addProfileAction = */
    profileMenu->addAction(QIcon::fromTheme("document-new"), i18n("Add"));
    /* auto importProfileAction = */
    profileMenu->addAction(QIcon::fromTheme("document-import"), i18n("Import"));
    profileMenu->addSeparator();
    // TODO: profiles

    menu->addSeparator();

    auto configure = menu->addAction(QIcon::fromTheme("configure"), i18n("Configure..."));
    QObject::connect(configure, &QAction::triggered, configure, [&engine]() {
        const QUrl url("qrc:/main.qml");
        engine.load(url);
        // TODO: How to keep only one instance
    });
    menu->addAction(QIcon::fromTheme("application-exit"), i18n("Quit"), [](){qApp->quit();});
    // TODO: Add profile quick switch

    // Set systray
    QSystemTrayIcon systray;
    systray.setToolTip("Klash");
    systray.setIcon(QIcon(":/assets/logo.png"));
    systray.setVisible(true);
    systray.setContextMenu(menu);

    // Prevent app from closing
    app.setQuitOnLastWindowClosed(false);
#endif

    const QUrl url("qrc:/main.qml");
    engine.load(url);

    return app.exec();
}
